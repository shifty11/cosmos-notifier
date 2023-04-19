use sycamore::futures::spawn_local_scoped;
use sycamore::prelude::*;

use crate::components::error_overlay::create_message;
use crate::components::sidebar::Sidebar;
use crate::{AppRoutes, InfoLevel};

#[component]
pub fn Overview<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        Sidebar()
        p {"Overview"}
        a(href=AppRoutes::Home) { "Home" }
        button(on:click=move |_| {
            spawn_local_scoped(cx, async move {
                create_message(cx, "new error".to_string(), InfoLevel::Error);
                create_message(cx, "new info".to_string(), InfoLevel::Info);
            });
        }) { "Add error" }
    }
}
