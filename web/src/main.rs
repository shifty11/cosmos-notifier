#![allow(non_snake_case)]

use std::fmt::Display;

use log::debug;
use log::Level;
use sycamore::futures::{spawn_local, spawn_local_scoped};
use sycamore::prelude::*;
use sycamore_router::{HistoryIntegration, Route, Router};
use tonic::Status;

use crate::services::auth::AuthService;
use crate::services::grpc::{GrpcClient, LoginResponse};

mod components;
mod config;
mod pages;
mod services;

#[derive(Route, Debug, Clone)]
pub enum AppRoutes {
    #[to("/")]
    Home,
    #[to("/overview")]
    Overview,
    #[to("/reminders")]
    Reminders,
    #[to("/communication")]
    Communication,
    #[to("/login")]
    Login,
    #[not_found]
    NotFound,
}

impl ToString for AppRoutes {
    fn to_string(&self) -> String {
        match self {
            AppRoutes::Home => "/".to_string(),
            AppRoutes::Overview => "/overview".to_string(),
            AppRoutes::Reminders => "/reminders".to_string(),
            AppRoutes::Communication => "/communication".to_string(),
            AppRoutes::Login => "/login".to_string(),
            AppRoutes::NotFound => "/404".to_string(),
        }
    }
}

#[derive(Debug, Clone)]
pub struct Services {
    pub grpc_client: GrpcClient,
    pub auth_manager: AuthService,
}

impl Services {
    pub fn new() -> Self {
        Self {
            grpc_client: GrpcClient::default(),
            auth_manager: AuthService::default(),
        }
    }
}

impl Default for Services {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Clone, PartialEq)]
pub enum AuthState {
    LoggedOut,
    LoggedIn,
    LoggingIn,
}

impl Display for AuthState {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            AuthState::LoggedOut => write!(f, "LoggedOut"),
            AuthState::LoggedIn => write!(f, "LoggedIn"),
            AuthState::LoggingIn => write!(f, "LoggingIn"),
        }
    }
}

#[derive(Debug, Clone)]
pub struct AppState {
    pub auth_state: RcSignal<AuthState>,
    pub route: RcSignal<AppRoutes>,
}

impl AppState {
    pub fn new(auth_manager: AuthService) -> Self {
        let auth_state = match auth_manager.is_jwt_valid() {
            true => AuthState::LoggedIn,
            false => AuthState::LoggedOut,
        };
        Self {
            auth_state: create_rc_signal(auth_state),
            route: create_rc_signal(AppRoutes::Overview),
        }
    }
}

fn start_jwt_refresh_timer() {
    spawn_local(async {
        gloo_timers::future::TimeoutFuture::new(1000 * 60).await;
        let auth_client = AuthService::new();
        debug!("is_jwt_valid: {}", auth_client.is_jwt_valid());
        if auth_client.is_jwt_about_to_expire() {
            auth_client.refresh_access_token().await;
        }
        start_jwt_refresh_timer();
    });
}

#[component]
pub async fn App<G: Html>(cx: Scope<'_>) -> View<G> {
    let services = Services::new();
    let app_state = AppState::new(services.auth_manager.clone());

    // TODO: make this cleaner
    if services.auth_manager.clone().has_login_query_params() {
        debug!("Logging in...");
        let resp = services.auth_manager.clone().login().await;
        match resp {
            Ok(_) => {
                app_state.auth_state.set(AuthState::LoggedIn);
                app_state.route.set(AppRoutes::Overview);
                debug!("Logged in successfully");
            }
            Err(_) => {}
        }
    }

    provide_context(cx, services.clone());
    provide_context(cx, app_state.clone());

    start_jwt_refresh_timer();  //TODO: make this with a scope and updating the state

    view! {cx,
        div(class="h-full w-full") {
            Router(
                integration=HistoryIntegration::new(),
                view=|cx, route: &ReadSignal<AppRoutes>| {
                    view! {cx, (
                            match route.get().as_ref() {
                                AppRoutes::Home => pages::home::page::Home(cx),
                                AppRoutes::Overview => pages::overview::page::Overview(cx),
                                AppRoutes::Reminders => pages::reminders::page::Reminders(cx),
                                AppRoutes::Communication => pages::communication::page::Communication(cx),
                                AppRoutes::Login => pages::login::page::Login(cx),
                                AppRoutes::NotFound => view! { cx, "404 Not Found"}
                            }
                        )
                    }
                }
            )
        }
    }
}

fn main() {
    console_error_panic_hook::set_once();
    console_log::init_with_level(Level::Debug).unwrap();
    debug!("Console log level set to debug");

    sycamore::render(|cx| view! { cx, App()});
}
