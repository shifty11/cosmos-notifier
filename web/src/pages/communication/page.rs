use sycamore::prelude::*;
use crate::AppRoutes;

#[component]
pub fn Communication<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        p {"Home"}
        a(href=AppRoutes::Overview) { "Overview" }
    }
}