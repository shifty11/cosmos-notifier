use crate::AppRoutes;
use sycamore::prelude::*;

#[component]
pub fn Communication<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        p {"Home"}
        a(href=AppRoutes::Overview) { "Overview" }
    }
}
