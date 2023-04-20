use std::error::Error;

use base64::engine::general_purpose::URL_SAFE_NO_PAD;
use base64::Engine;
use gloo_storage::{LocalStorage, Storage};
use grpc_web_client::Client;
use log::debug;
use serde::{Deserialize, Serialize};
use simple_error::bail;
use tonic::Status;
use wasm_bindgen::JsValue;

use crate::config::keys;
use crate::services::grpc::auth_service_client::AuthServiceClient;
use crate::services::grpc::dev_service_client::DevServiceClient;
use crate::services::grpc::{
    dev_login_request, DevLoginRequest, DiscordLoginRequest, LoginResponse,
    RefreshAccessTokenRequest, TelegramLoginRequest,
};

#[derive(Debug, Clone)]
enum UserType {
    Discord = dev_login_request::UserType::Discord as isize,
    Telegram = dev_login_request::UserType::Telegram as isize,
}

impl<'de> Deserialize<'de> for UserType {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        let s = String::deserialize(deserializer)?;

        match s.as_str() {
            "discord" => Ok(UserType::Discord),
            "telegram" => Ok(UserType::Telegram),
            _ => Err(serde::de::Error::unknown_variant(
                &s,
                &["discord", "telegram"],
            )),
        }
    }
}

impl Serialize for UserType {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        match self {
            UserType::Discord => serializer.serialize_str("discord"),
            UserType::Telegram => serializer.serialize_str("telegram"),
        }
    }
}

#[derive(Debug, Clone, Eq, PartialEq)]
enum Role {
    User = dev_login_request::Role::User as isize,
    Admin = dev_login_request::Role::Admin as isize,
}

impl<'de> Deserialize<'de> for Role {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        let s = String::deserialize(deserializer)?;

        match s.as_str() {
            "admin" => Ok(Role::Admin),
            "user" => Ok(Role::User),
            _ => Err(serde::de::Error::unknown_variant(&s, &["admin", "user"])),
        }
    }
}

impl Serialize for Role {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        match self {
            Role::Admin => serializer.serialize_str("admin"),
            Role::User => serializer.serialize_str("user"),
        }
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
struct Claims {
    exp: usize,
    user_id: i64,
    #[serde(rename = "type")]
    user_type: UserType,
    role: Role,
}

const IS_ABOUT_TO_EXPIRE: usize = 60 * 5; // seconds

impl Claims {
    fn is_expired(&self) -> bool {
        let now = (js_sys::Date::now() / 1000.0) as usize;
        now > self.exp
    }

    fn is_about_to_expire(&self) -> bool {
        let now = (js_sys::Date::now() / 1000.0) as usize;
        now > self.exp - IS_ABOUT_TO_EXPIRE
    }
}

#[derive(Debug, Clone)]
pub struct AuthService {
    endpoint_url: String,
}

impl Default for AuthService {
    fn default() -> Self {
        Self::new()
    }
}

impl AuthService {
    pub fn new() -> Self {
        Self {
            endpoint_url: env!("GRPC_ENDPOINT_URL").to_string(),
        }
    }

    fn auth_service(&self) -> AuthServiceClient<Client> {
        let client = Client::new(self.endpoint_url.clone());
        AuthServiceClient::new(client)
    }

    pub fn is_jwt_valid(&self) -> bool {
        match self.get_jwt_claims().ok() {
            Some(claims) => !claims.is_expired(),
            None => false,
        }
    }

    pub fn is_jwt_about_to_expire(&self) -> bool {
        match self.get_jwt_claims().ok() {
            Some(claims) => claims.is_about_to_expire(),
            None => true,
        }
    }

    fn get_jwt_claims(&self) -> Result<Claims, Box<dyn Error>> {
        match self.get_access_token() {
            Ok(token) => self.decode_jwt(&token),
            Err(err) => Err(Box::try_from(err).unwrap()),
        }
    }

    fn decode_jwt(&self, jwt: &str) -> Result<Claims, Box<dyn Error>> {
        let parts: Vec<&str> = jwt.split('.').collect();
        if parts.len() != 3 {
            bail!("Invalid token.")
        }

        let payload = self.decode_base64(parts[1])?;
        let claims: Claims = serde_json::from_str(&payload).map_err(|_| "Invalid payload.")?;

        Ok(claims)
    }

    fn decode_base64(&self, input: &str) -> Result<String, String> {
        let decoded_bytes = URL_SAFE_NO_PAD
            .decode(input)
            .map_err(|_| "Illegal base64 string.")?;
        let decoded_str = String::from_utf8(decoded_bytes).map_err(|_| "Invalid UTF-8 string.")?;
        Ok(decoded_str)
    }

    pub fn get_access_token(&self) -> gloo_storage::Result<String> {
        LocalStorage::get(keys::LS_KEY_ACCESS_TOKEN)
    }

    fn get_refresh_token(&self) -> gloo_storage::Result<String> {
        LocalStorage::get(keys::LS_KEY_REFRESH_TOKEN)
    }

    pub async fn refresh_access_token(&self) {
        debug!("refresh_access_token");
        let mut auth_service = self.auth_service();
        if let Ok(token) = self.get_refresh_token() {
            let req = RefreshAccessTokenRequest {
                refresh_token: token,
            };
            let resp = auth_service.refresh_access_token(req).await;
            match resp {
                Ok(resp) => {
                    debug!("set access token");
                    LocalStorage::set(keys::LS_KEY_ACCESS_TOKEN, resp.into_inner().access_token)
                        .unwrap();
                }
                Err(status) => {
                    if status.code() == tonic::Code::Unauthenticated {
                        debug!("refresh_access_token: Unauthenticated");
                        self.logout();
                    } else {
                        debug!("refresh_access_token: {}", status);
                    }
                }
            }
        }
    }

    pub async fn login(&mut self) -> Result<LoginResponse, Status> {
        let request = DevLoginRequest {
            user_id: 0,
            user_type: dev_login_request::UserType::Discord as i32,
            role: dev_login_request::Role::Admin as i32,
        };
        let client = Client::new(self.endpoint_url.clone());
        let response = DevServiceClient::new(client)
            .login(request)
            .await
            .map(|res| res.into_inner());
        self.save_login_response(response.clone().unwrap());
        response
    }

    pub fn logout(&self) {
        LocalStorage::delete(keys::LS_KEY_ACCESS_TOKEN);
        LocalStorage::delete(keys::LS_KEY_REFRESH_TOKEN);
    }

    fn save_login_response(&self, response: LoginResponse) {
        LocalStorage::set(keys::LS_KEY_ACCESS_TOKEN, response.access_token).unwrap();
        LocalStorage::set(keys::LS_KEY_REFRESH_TOKEN, response.refresh_token).unwrap();
    }

    pub async fn login_with_query_params(&self) -> Result<LoginResponse, Status> {
        if self.has_discord_login_query_params() {
            return self.login_with_discord_query_params().await;
        } else if self.has_telegram_login_query_params() {
            return self.login_with_telegram_query_params().await;
        }
        Err(Status::new(
            tonic::Code::InvalidArgument,
            "Invalid query params".to_string(),
        ))
    }

    async fn login_with_discord_query_params(&self) -> Result<LoginResponse, Status> {
        debug!("login_with_discord_query_params");
        let mut auth_service = self.auth_service();
        let code = self
            .get_query_params()
            .iter()
            .find(|params: &&(String, String)| params.0 == "code")
            .unwrap()
            .1
            .clone();
        let req = DiscordLoginRequest { code };
        let resp = auth_service
            .discord_login(req)
            .await
            .map(|res| res.into_inner());
        self.save_login_response(resp.clone()?);
        resp
    }

    async fn login_with_telegram_query_params(&self) -> Result<LoginResponse, Status> {
        // TODO: fix this
        debug!("login_with_telegram_query_params");
        let mut auth_service = self.auth_service();
        let hash = self
            .get_query_params()
            .iter()
            .find(|params: &&(String, String)| params.0 == "hash")
            .unwrap()
            .1
            .clone();
        let req = TelegramLoginRequest {
            user_id: 0,
            data_str: "".to_string(),
            username: "".to_string(),
            auth_date: 0,
            hash,
        };
        let resp = auth_service
            .telegram_login(req)
            .await
            .map(|res| res.into_inner());
        debug!("login_with_telegram_query_params: {:?}", resp);
        resp
    }

    pub fn has_login_query_params(&self) -> bool {
        self.has_telegram_login_query_params() || self.has_discord_login_query_params()
    }

    fn has_telegram_login_query_params(&self) -> bool {
        self.get_query_params()
            .iter()
            .map(|params: &(String, String)| params.0 == "hash")
            .any(|x| x)
    }

    fn has_discord_login_query_params(&self) -> bool {
        self.get_query_params()
            .iter()
            .map(|params: &(String, String)| params.0 == "code")
            .any(|x| x)
    }

    fn get_query_params(&self) -> Vec<(String, String)> {
        let location = gloo_utils::window().location();
        let search: Result<String, JsValue> = location.search();
        let mut params = Vec::new();
        for s in search.unwrap().trim_start_matches('?').split('&') {
            if s.is_empty() {
                continue;
            }
            let mut kv = s.split('=');
            let k = kv.next().unwrap();
            let v = kv.next().unwrap();
            params.push((k.to_string(), v.to_string()));
        }
        params
    }

    pub fn is_admin(&self) -> bool {
        self.get_jwt_claims()
            .map(|claims| claims.role == Role::Admin)
            .unwrap_or(false)
    }

    pub fn is_user(&self) -> bool {
        self.get_jwt_claims()
            .map(|claims| claims.role == Role::User)
            .unwrap_or(false)
    }
}
