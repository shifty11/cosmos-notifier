#![allow(non_snake_case)]

use std::fmt::Display;

use log::debug;
use log::Level;
use sycamore::futures::spawn_local_scoped;
use sycamore::prelude::*;
use sycamore::suspense::Suspense;
use sycamore_router::{HistoryIntegration, navigate, Route, Router};
use uuid::Uuid;

use crate::components::error_overlay::{
    create_error_msg_from_status, create_message, ErrorOverlay,
};
use crate::components::sidebar::SidebarWrapper;
use crate::pages::communication::page::Communication;
use crate::pages::home::page::Home;
use crate::pages::login::page::Login;
use crate::pages::overview::page::Overview;
use crate::pages::reminders::page::Reminders;
use crate::services::auth::AuthService;
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

#[repr(usize)]
#[derive(Debug, Clone, Eq, PartialEq)]
pub enum InfoLevel {
    Error = 1,
    Info,
}

#[derive(Debug, Clone, PartialEq)]
pub struct InfoMsg {
    pub id: Uuid,
    pub msg: String,
    pub level: InfoLevel,
}

#[derive(Debug, Clone)]
pub struct AppState {
    pub auth_state: RcSignal<AuthState>,
    pub route: RcSignal<AppRoutes>,
    pub messages: RcSignal<Vec<RcSignal<InfoMsg>>>,
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
            messages: create_rc_signal(vec![]),
        }
    }

    pub fn add_message(&self, msg: String, level: InfoLevel) -> Uuid {
        let uuid = Uuid::new_v4();
        self.messages.modify().push(create_rc_signal(InfoMsg {
            id: uuid,
            msg,
            level,
        }));
        uuid
    }

    pub fn remove_message(&self, id: Uuid) {
        self.messages.modify().retain(|m| m.get().id != id);
    }
}

fn start_jwt_refresh_timer(cx: Scope) {
    spawn_local_scoped(cx, async move {
        gloo_timers::future::TimeoutFuture::new(1000 * 60).await;
        let auth_client = AuthService::new();
        debug!("is_jwt_valid: {}", auth_client.is_jwt_valid());
        if auth_client.is_jwt_about_to_expire() {
            auth_client.refresh_access_token().await;
        }
        if auth_client.is_jwt_valid() {
            start_jwt_refresh_timer(cx.to_owned());
        } else {
            debug!("JWT is not valid anymore");
            auth_client.logout();
            let app_state = use_context::<AppState>(cx);
            app_state.auth_state.set(AuthState::LoggedOut);
        }
    });
}

fn get_active_view<G: Html>(cx: Scope, route: &AppRoutes) -> View<G> {
    let app_state = use_context::<AppState>(cx);
    app_state.route.set(route.clone());
    debug!("Route changed to: {:?}", route);
    match route {
        AppRoutes::Home => view!(cx, SidebarWrapper{Home {}}),
        AppRoutes::Overview => view!(cx, SidebarWrapper{Overview {}}),
        AppRoutes::Reminders => view!(cx, SidebarWrapper{Reminders {}}),
        AppRoutes::Communication => view!(cx, SidebarWrapper{Communication {}}),
        AppRoutes::Login => Login(cx),
        AppRoutes::NotFound => view! { cx, "404 Not Found"},
    }
}

#[component]
pub async fn App<G: Html>(cx: Scope<'_>) -> View<G> {
    let services = Services::new();
    let app_state = AppState::new(services.auth_manager.clone());

    provide_context(cx, services.clone());
    provide_context(cx, app_state.clone());

    if services.auth_manager.clone().has_login_query_params() {
        debug!("Logging in...");
        let resp = services
            .auth_manager
            .clone()
            .login_with_query_params()
            .await;
        match resp {
            Ok(_) => {
                app_state.auth_state.set(AuthState::LoggedIn);
                create_message(cx, "Logged in successfully".to_string(), InfoLevel::Info);
            }
            Err(e) => create_error_msg_from_status(cx, e),
        }
    }

    start_jwt_refresh_timer(cx.to_owned());

    view! {cx,
        div(class="flex min-h-screen") {
            ErrorOverlay {}
            Router(
                integration=HistoryIntegration::new(),
                view=|cx, route: &ReadSignal<AppRoutes>| {
                    create_effect(cx, move || {
                        let app_state = use_context::<AppState>(cx);
                        let auth_state = app_state.auth_state.get();
                        debug!("Auth state changed: {}", auth_state);
                        match auth_state.as_ref() {
                            AuthState::LoggedOut => navigate(AppRoutes::Login.to_string().as_str()),
                            AuthState::LoggedIn => navigate(AppRoutes::Overview.to_string().as_str()),
                            AuthState::LoggingIn => {}
                        }
                    });
                    view! {cx, (
                            get_active_view(cx, route.get().as_ref())
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

    sycamore::render(|cx| {
        view! { cx,
            Suspense(fallback=components::loading::LoadingSpinner(cx)) {
                App {}
            }
        }
    });
}
