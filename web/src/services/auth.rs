use std::error::Error;

use base64::Engine;
use base64::engine::general_purpose::URL_SAFE_NO_PAD;
use gloo_storage::{LocalStorage, Storage};
use log::debug;
use serde::{Deserialize, Serialize};
use simple_error::bail;
use tonic::Status;
use crate::config::keys;

use crate::services::grpc::auth_service_client::AuthServiceClient;
use crate::services::grpc::{dev_login_request, DevLoginRequest, LoginResponse, RefreshAccessTokenRequest};
use crate::services::grpc::dev_service_client::DevServiceClient;

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
            _ => Err(serde::de::Error::unknown_variant(&s, &["discord", "telegram"])),
        }
    }
}

impl<'de> Serialize for UserType {
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

#[derive(Debug, Clone)]
enum Role {
    User = dev_login_request::Role::User as isize,
    Admin = dev_login_request::Role::Admin as isize
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

impl<'de> Serialize for Role {
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

const IS_ABOUT_TO_EXPIRE: usize = 60 * 5;    // seconds

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
pub struct AuthManager {
    endpoint_url: String,
}

impl Default for AuthManager {
    fn default() -> Self {
        Self::new()
    }
}

impl AuthManager {
    pub fn new() -> Self {
        Self {
            endpoint_url: env!("GRPC_ENDPOINT_URL").to_string(),
        }
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
            Err(err) => {
                debug!("get_jwt_claims: {}", err);
                Err(Box::try_from(err).unwrap())
            }
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
        let decoded_bytes = URL_SAFE_NO_PAD.decode(&input).map_err(|_| "Illegal base64 string.")?;
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
        let client = grpc_web_client::Client::new(self.endpoint_url.clone());
        let mut auth_service = AuthServiceClient::new(client);
        if let Ok(token) = self.get_refresh_token() {
            let req = RefreshAccessTokenRequest {
                refresh_token: token,
            };
            let resp = auth_service.refresh_access_token(req).await;
            match resp {
                Ok(resp) => {
                    debug!("set access token");
                    LocalStorage::set(keys::LS_KEY_ACCESS_TOKEN, resp.into_inner().access_token).unwrap();
                }
                Err(_) => {}
            }
        }
    }

    pub async fn login(&mut self) -> Result<LoginResponse, Status> {
        let request = DevLoginRequest {
            user_id: 0,
            user_type: dev_login_request::UserType::Discord as i32,
            role: dev_login_request::Role::Admin as i32,
        };
        let client = grpc_web_client::Client::new(self.endpoint_url.clone());
        let response = DevServiceClient::new(client).login(request).await.map(|res| res.into_inner());

        match response.clone() {
            Ok(result) => {
                LocalStorage::set(keys::LS_KEY_ACCESS_TOKEN, result.access_token).unwrap();
                LocalStorage::set(keys::LS_KEY_REFRESH_TOKEN, result.refresh_token).unwrap();
            }
            Err(_) => {}
        }
        response
    }
}
