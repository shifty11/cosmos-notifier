use sycamore::prelude::*;
use crate::AppRoutes;

#[component]
pub fn Home<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        p {"Home"}
        a(href=AppRoutes::Overview) { "Overview" }
        a(href=AppRoutes::Login) { "Login" }
    }
}