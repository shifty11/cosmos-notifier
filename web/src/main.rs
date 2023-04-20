#![allow(non_snake_case)]

use std::fmt::Display;

use log::debug;
use log::Level;
use sycamore::futures::spawn_local_scoped;
use sycamore::prelude::*;
use sycamore::suspense::Suspense;
use sycamore_router::{navigate, HistoryIntegration, Route, Router};
use uuid::Uuid;

use crate::components::layout::LayoutWrapper;
use crate::components::messages::{create_error_msg_from_status, create_message, MessageOverlay};
use crate::config::keys;
use crate::pages::communication::page::Communication;
use crate::pages::home::page::Home;
use crate::pages::login::page::Login;
use crate::pages::overview::page::Overview;
use crate::pages::reminders::page::Reminders;
use crate::services::auth::AuthService;
use crate::services::grpc::GrpcClient;
use crate::services::grpc::User;

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
    Info = 1,
    Success,
    Warning,
    Error,
}

#[derive(Debug, Clone, PartialEq)]
pub struct InfoMsg {
    pub id: Uuid,
    pub title: String,
    pub message: String,
    pub level: InfoLevel,
    pub created_at: f64,
}

#[derive(Debug, Clone)]
pub struct AppState {
    auth_service: AuthService,
    pub auth_state: RcSignal<AuthState>,
    pub route: RcSignal<AppRoutes>,
    pub messages: RcSignal<Vec<RcSignal<InfoMsg>>>,
    pub user: RcSignal<Option<User>>,
}

impl AppState {
    pub fn new(auth_service: AuthService) -> Self {
        let auth_state = match auth_service.is_jwt_valid() {
            true => AuthState::LoggedIn,
            false => AuthState::LoggedOut,
        };
        Self {
            auth_service,
            auth_state: create_rc_signal(auth_state),
            route: create_rc_signal(AppRoutes::Overview),
            messages: create_rc_signal(vec![]),
            user: create_rc_signal(None),
        }
    }

    pub fn add_message(&self, title: String, message: String, level: InfoLevel) -> Uuid {
        let uuid = Uuid::new_v4();
        let created_at = js_sys::Date::now();
        self.messages.modify().push(create_rc_signal(InfoMsg {
            id: uuid,
            title,
            message,
            level,
            created_at,
        }));
        uuid
    }

    pub fn remove_message(&self, id: Uuid) {
        self.messages.modify().retain(|m| m.get().id != id);
    }

    pub fn logout(&self) {
        self.auth_service.logout();
        self.auth_state.set(AuthState::LoggedOut);
    }

    pub fn get_user_name(&self) -> String {
        match self.user.get().as_ref() {
            Some(user) => user.name.clone(),
            None => "Unknown".to_string(),
        }
    }

    pub fn get_user_avatar(&self) -> String {
        let user = self.user.get().as_ref().clone();
        if let Some(user) = user {
            if user.avatar != "" {
                return user.avatar;
            }
        }
        keys::DEVAULT_AVATAR_PATH.to_string()
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
            let app_state = use_context::<AppState>(cx);
            app_state.logout();
        }
    });
}

fn has_access_permission(auth_service: &AuthService, route: &AppRoutes) -> bool {
    let is_admin = auth_service.is_admin();
    let is_user = auth_service.is_user();
    match route {
        AppRoutes::Home => true,
        AppRoutes::Overview => is_user || is_admin,
        AppRoutes::Reminders => is_user || is_admin,
        AppRoutes::Communication => is_user || is_admin,
        AppRoutes::Login => true,
        AppRoutes::NotFound => true,
    }
}

fn activate_view<G: Html>(cx: Scope, route: &AppRoutes) -> View<G> {
    debug!("Route changed to: {:?}", route);
    let app_state = use_context::<AppState>(cx);
    let services = use_context::<Services>(cx);
    if has_access_permission(&services.auth_manager, route) {
        app_state.route.set(route.clone());
        match route {
            AppRoutes::Home => view!(cx, LayoutWrapper{Home {}}),
            AppRoutes::Overview => view!(cx, LayoutWrapper{Overview {}}),
            AppRoutes::Reminders => view!(cx, LayoutWrapper{Reminders {}}),
            AppRoutes::Communication => view!(cx, LayoutWrapper{Communication {}}),
            AppRoutes::Login => Login(cx),
            AppRoutes::NotFound => view! { cx, "404 Not Found"},
        }
    } else {
        app_state.route.set(AppRoutes::Login);
        create_message(
            cx,
            "Access denied".to_string(),
            "Please login to access this page",
            InfoLevel::Info,
        );
        Login(cx)
    }
}

async fn get_user_info(cx: Scope<'_>) {
    let app_state = use_context::<AppState>(cx);
    let services = use_context::<Services>(cx);
    let request = services.grpc_client.create_request({});
    let response = services
        .grpc_client
        .get_user_service()
        .get_user(request)
        .await
        .map(|res| res.into_inner());
    if let Ok(user) = response {
        *app_state.user.modify() = Some(user);
    } else {
        create_error_msg_from_status(cx, response.err().unwrap());
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
                create_message(
                    cx,
                    "Login success",
                    "Logged in successfully",
                    InfoLevel::Info,
                );
            }
            Err(e) => create_error_msg_from_status(cx, e),
        }
    }

    start_jwt_refresh_timer(cx.to_owned());

    view! {cx,
        div(class="flex min-h-screen") {
            MessageOverlay {}
            Router(
                integration=HistoryIntegration::new(),
                view=|cx, route: &ReadSignal<AppRoutes>| {
                    debug!("Router: create_effect");
                    create_effect(cx, move || {
                        let app_state = use_context::<AppState>(cx);
                        let auth_state = app_state.auth_state.get();
                        debug!("Auth state changed: {}", auth_state);
                        match auth_state.as_ref() {
                            AuthState::LoggedOut => navigate(AppRoutes::Login.to_string().as_str()),
                            AuthState::LoggedIn => {
                                spawn_local_scoped(cx, async move {
                                    get_user_info(cx).await;
                                });
                                navigate(AppRoutes::Overview.to_string().as_str())
                            },
                            AuthState::LoggingIn => {}
                        }
                    });
                    view! {cx, (
                            activate_view(cx, route.get().as_ref())
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
