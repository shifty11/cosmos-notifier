#![allow(non_snake_case)]

use std::borrow::Borrow;
use std::fmt::Display;
use std::ops::Deref;
use std::rc::Rc;

use log::debug;
use log::Level;
use sycamore::futures::spawn_local;
use sycamore::prelude::*;
use tonic::Status;

use crate::services::grpc::{GrpcClient, LoginResponse};

mod services;

#[derive(Debug, Default, Clone)]
pub struct Services {
    pub grpc_client: RcSignal<GrpcClient>,
}

impl Services {
    pub fn new() -> Self {
        Self {
            grpc_client: create_rc_signal(GrpcClient::default()),
        }
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
}

impl AppState {
    pub fn new() -> Self {
        Self {
            auth_state: create_rc_signal(AuthState::LoggedOut),
        }
    }
}

#[component]
async fn InitComponent<G: Html>(cx: Scope<'_>) -> View<G> {
    let auth_state = use_context::<AppState>(cx).auth_state.get();
    let auth_state = match *auth_state {
        AuthState::LoggedOut => {
            debug!("Try to login");
            let mut grpc_client = use_context::<Services>(cx).grpc_client.modify();
            let response = grpc_client.login().await;
            match response {
                Ok(_) => {
                    debug!("Login successful");
                    let mut auth_state = use_context::<AppState>(cx).auth_state.modify();
                    *auth_state = AuthState::LoggedIn;
                }
                Err(_) => {}
            }
        }
        _ => {}
    };

    view!(cx,
        p { "Hello, Fetch!" }
        p { (use_context::<AppState>(cx).auth_state.get()) }
    )
}

#[component]
fn SubComponent<G: Html>(cx: Scope) -> View<G> {
    let app_state = use_context::<AppState>(cx);
    let text = create_selector(cx, || {
        debug!("auth_state changed: {}", app_state.auth_state.get());
        match *app_state.auth_state.get() {
            AuthState::LoggedOut => "LoggedOut",
            AuthState::LoggedIn => "LoggedIn",
            AuthState::LoggingIn => "LoggingIn",
        }
    });

    view!(cx,
        p { "Hello, SubComponent!" }
        p { (text.get()) }
    )
}

async fn timer_function() {
    // Do something here
    debug!("Timer function executed!");
}


#[component]
fn App<G: Html>(cx: Scope<'_>) -> View<G> {
    provide_context(cx, Services::new());
    provide_context(cx, AppState::new());

    view! {cx,
        h1 { "Hello, World!" }
        InitComponent()
        SubComponent()
    }
}

fn main() {
    console_error_panic_hook::set_once();
    console_log::init_with_level(Level::Debug).unwrap();
    debug!("Console log level set to debug");

    sycamore::render(|cx| view! { cx, App()});
}