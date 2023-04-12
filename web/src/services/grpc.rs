use std::collections::HashMap;
use std::fmt::{Debug, Formatter};

use grpc_web_client::Client;
use jsonwebtoken::{decode, DecodingKey, TokenData, Validation};
use log::debug;
use serde::{Deserialize, Serialize};
use sycamore::rt::JsValue;
use tonic::{Request, Status};
use tonic::metadata::MetadataValue;

use std::time::{SystemTime, UNIX_EPOCH};
use crate::services::grpc::auth_service_client::AuthServiceClient;
use crate::services::grpc::dev_login_request::{Role, UserType};
use crate::services::grpc::dev_service_client::DevServiceClient;
use crate::services::grpc::tracker_service_client::TrackerServiceClient;

tonic::include_proto!("cosmos_notifier_grpc");

#[derive(Debug, Default, Clone)]
struct LocalStorage {
    items: HashMap<String, String>,
}

impl LocalStorage {
    fn new() -> Result<Self, JsValue> {
        let window = web_sys::window().ok_or(JsValue::from("no global `window` exists"))?;
        let local_storage = window.local_storage()?;
        let mut items = HashMap::new();

        if let Some(storage) = local_storage {
            for i in 0..storage.length()? {
                let key = storage.key(i)?;
                if let Some(key) = key {
                    let value = storage.get_item(&key)?;
                    if let Some(value) = value {
                        items.insert(key, value);
                    }
                }
            }
        }

        Ok(Self { items })
    }

    fn get(&self, key: &str) -> Option<&String> {
        self.items.get(key)
    }

    fn set(&mut self, key: &str, value: &str) {
        self.items.insert(key.to_string(), value.to_string());
        let window = web_sys::window().unwrap();
        let local_storage = window.local_storage().unwrap().unwrap();
        local_storage
            .set_item(key, value)
            .unwrap();
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
struct Claims {
    // sub: String, // Subject (typically a user ID)
    exp: usize,  // Expiration time (Unix timestamp)
}

impl Claims {
    fn is_expired(&self) -> bool {
        let now = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_secs() as usize;
        now > self.exp
    }
}

const ACCESS_TOKEN_KEY: &str = "access_token";
const REFRESH_TOKEN_KEY: &str = "refresh_token";

#[derive(Clone)]
pub struct GrpcClient {
    client: Client,
    local_storage: LocalStorage,
}

impl Default for GrpcClient {
    fn default() -> Self {
        Self::new("http://test.mydomain.com:8090".to_string())
    }
}

impl Debug for GrpcClient {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("GrpcClient")
            .field("local_storage", &self.local_storage)
            .finish()
    }
}

impl GrpcClient {
    pub(crate) fn new(url: String) -> Self {
        let client = Client::new(url);
        Self {
            client,
            local_storage: LocalStorage::new().unwrap(),
        }
    }

    pub fn create_request<T>(&self, message: T) -> Request<T> {
        let token = self.local_storage.get(ACCESS_TOKEN_KEY);
        let mut req = Request::new(message);
        if token.is_some() {
            req.metadata_mut().insert(
                "authorization",
                MetadataValue::try_from(token.unwrap()).unwrap_or_else(|_| MetadataValue::from_static(""))
            );
        }
        req
    }

    fn is_jwt_valid(&self, token: &str) -> bool {
        if let Some(claims) = self.decode_jwt(token) {
            !claims.is_expired()
        } else {
            false
        }
    }

    fn decode_jwt(&self, jwt: &str) -> Option<Claims> {
        let decoding_key = DecodingKey::from_secret(jwt.as_ref());
        match decode::<Claims>(jwt, &decoding_key, &Validation::default()) {
            Ok(TokenData { claims, .. }) => Some(claims),
            Err(_) => None,
        }
    }

    async fn try_to_refresh_access_token(&mut self) -> Option<&String> {
        if let Some(token) = self.local_storage.get(REFRESH_TOKEN_KEY) {
            let request = RefreshAccessTokenRequest {
                refresh_token: token.to_string(),
            };
            let response = self.get_auth_service().refresh_access_token(request).await.map(|res| res.into_inner());
            match response {
                Ok(result) => {
                    self.local_storage.set(ACCESS_TOKEN_KEY, &result.clone().access_token);
                    return self.local_storage.get(ACCESS_TOKEN_KEY);
                }
                Err(_) => {}
            }
        }
        None
    }

    fn get_dev_service(&self) -> DevServiceClient<Client> {
        DevServiceClient::new(self.client.clone())
    }

    pub fn get_auth_service(&self) -> AuthServiceClient<Client> {
        AuthServiceClient::new(self.client.clone())
    }

    pub fn get_tracker_service(&self) -> TrackerServiceClient<Client> {
        TrackerServiceClient::new(self.client.clone())
    }

    pub async fn login(&mut self) -> Result<LoginResponse, Status> {
        let request = self.create_request(DevLoginRequest {
            user_id: 0,
            user_type: UserType::Discord as i32,
            role: Role::Admin as i32,
        });
        let response = self.get_dev_service().login(request).await.map(|res| res.into_inner());

        match response.clone() {
            Ok(result) => {
                self.local_storage.set(ACCESS_TOKEN_KEY, &result.access_token);
                self.local_storage.set(REFRESH_TOKEN_KEY, &result.refresh_token);
            }
            Err(_) => {}
        }
        response
    }
}
