use sycamore::prelude::*;
use crate::AppRoutes;

#[component]
pub fn Overview<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        p {"Overview"}
        a(href=AppRoutes::Home) { "Home" }
    }
}