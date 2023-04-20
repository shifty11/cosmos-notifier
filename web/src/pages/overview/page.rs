use sycamore::futures::spawn_local_scoped;
use sycamore::prelude::*;

use crate::components::messages::create_message;
use crate::{AppRoutes, InfoLevel};

#[component]
pub fn Overview<G: Html>(cx: Scope) -> View<G> {
    let counter = create_signal(cx, 0);

    view! {cx,
        p {"Overview"}
        a(href=AppRoutes::Home) { "Home" }
        button(on:click=move |_| {
            spawn_local_scoped(cx, async move {
                counter.set(counter.get().as_ref() + 1);
                let counter_str = counter.get().as_ref().to_string();
                create_message(cx, "Error", counter_str.as_str(), InfoLevel::Error);
                create_message(cx, "Info", "New info", InfoLevel::Info);
            });
        }) { "Add error" }
    }
}
