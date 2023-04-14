use crate::components::sidebar::Sidebar;
use crate::AppRoutes;
use sycamore::prelude::*;

#[component]
pub fn Overview<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        Sidebar()
        p {"Overview"}
        a(href=AppRoutes::Home) { "Home" }
    }
}
