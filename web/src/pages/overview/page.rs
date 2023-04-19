use sycamore::futures::spawn_local_scoped;
use sycamore::prelude::*;

use crate::components::messages::create_message;
use crate::{AppRoutes, InfoLevel};

#[component]
pub fn Overview<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        p {"Overview"}
        a(href=AppRoutes::Home) { "Home" }
        button(on:click=move |_| {
            spawn_local_scoped(cx, async move {
                create_message(cx, "Error", "New error", InfoLevel::Error);
                create_message(cx, "Info", "New info", InfoLevel::Info);
            });
        }) { "Add error" }
    }
}
