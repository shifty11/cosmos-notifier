use log::debug;
use sycamore::futures::spawn_local_scoped;
use sycamore::prelude::*;

use crate::{AppRoutes, AppState, AuthState, Services};

// #[component]
// fn SubComponent<G: Html>(cx: Scope) -> View<G> {
//     let app_state = use_context::<AppState>(cx);
//     let text = create_selector(cx, || {
//         debug!("auth_state changed: {}", app_state.auth_state.get());
//         match *app_state.auth_state.get() {
//             AuthState::LoggedOut => "LoggedOut",
//             AuthState::LoggedIn => "LoggedIn",
//             AuthState::LoggingIn => "LoggingIn",
//         }
//     });
//
//     view!(cx,
//         p { "Hello, SubComponent!" }
//         p { (text.get()) }
//     )
// }

#[component]
pub async fn Login<G: Html>(cx: Scope<'_>) -> View<G> {
    let app_state = use_context::<AppState>(cx);

    view!(cx,
        p { "Hello, Init!" }
        p { (app_state.auth_state.get()) }
        button(on:click=move |_| {
            spawn_local_scoped(cx, async move {
                match *app_state.auth_state.get() {
                    AuthState::LoggedOut => {
                        debug!("Try to login");
                        let response = use_context::<Services>(cx).auth_manager.modify().login().await;
                        match response {
                            Ok(_) => {
                                debug!("Login successful");
                                let mut auth_state = use_context::<AppState>(cx).auth_state.modify();
                                *auth_state = AuthState::LoggedIn;
                            }
                            Err(status) => debug!("Login failed with error: {:?}", status),
                        }
                    }
                    _ => {}
                };
            });
        }) { "Login" }
        a(href=AppRoutes::Home) { "Home" }
    )
}
