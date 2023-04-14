#![allow(non_snake_case)]

use std::fmt::Display;

use log::debug;
use log::Level;
use sycamore::futures::spawn_local;
use sycamore::prelude::*;
use sycamore_router::{HistoryIntegration, Route, Router};

use crate::services::auth::AuthManager;
use crate::services::grpc::GrpcClient;

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
    pub grpc_client: RcSignal<GrpcClient>,
    pub auth_manager: RcSignal<AuthManager>,
}

impl Services {
    pub fn new() -> Self {
        Self {
            grpc_client: create_rc_signal(GrpcClient::default()),
            auth_manager: create_rc_signal(AuthManager::default()),
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
    pub fn new(auth_manager: &AuthManager) -> Self {
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
        let auth_client = AuthManager::new();
        debug!("is_jwt_valid: {}", auth_client.is_jwt_valid());
        if auth_client.is_jwt_about_to_expire() {
            auth_client.refresh_access_token().await;
        }
        start_jwt_refresh_timer();
    });
}

#[component]
pub fn App<G: Html>(cx: Scope) -> View<G> {
    let services = Services::new();
    provide_context(cx, services.clone());
    provide_context(
        cx,
        AppState::new(services.auth_manager.get_untracked().as_ref()),
    );

    start_jwt_refresh_timer();

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
