use std::fmt::Debug;

use grpc_web_client::Client;
use tonic::metadata::MetadataValue;
use tonic::Request;

use crate::services::auth::AuthManager;
use crate::services::grpc::auth_service_client::AuthServiceClient;
use crate::services::grpc::tracker_service_client::TrackerServiceClient;

tonic::include_proto!("cosmos_notifier_grpc");

#[derive(Debug, Clone)]
pub struct GrpcClient {
    endpoint_url: String,
    auth_manager: AuthManager,
}

impl Default for GrpcClient {
    fn default() -> Self {
        Self::new()
    }
}

impl GrpcClient {
    pub fn new() -> Self {
        Self {
            endpoint_url: env!("GRPC_ENDPOINT_URL").to_string(),
            auth_manager: AuthManager::new(),
        }
    }

    pub fn create_request<T>(&self, message: T) -> Request<T> {
        let token = self.auth_manager.get_access_token();
        let mut req = Request::new(message);
        if let Ok(token) = token {
            req.metadata_mut().insert(
                "authorization",
                MetadataValue::try_from(token).unwrap_or_else(|_| MetadataValue::from_static("")),
            );
        }
        req
    }

    pub fn get_auth_service(&self) -> AuthServiceClient<Client> {
        AuthServiceClient::new(Client::new(self.endpoint_url.clone()))
    }

    pub fn get_tracker_service(&self) -> TrackerServiceClient<Client> {
        TrackerServiceClient::new(Client::new(self.endpoint_url.clone()))
    }
}
