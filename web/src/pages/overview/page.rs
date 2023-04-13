use sycamore::prelude::*;
use crate::AppRoutes;
use crate::components::sidebar::Sidebar;

#[component]
pub fn Overview<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        Sidebar()
        p {"Overview"}
        a(href=AppRoutes::Home) { "Home" }
    }
}