use crate::components::sidebar::Sidebar;
use sycamore::prelude::*;

#[component]
pub fn Communication<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        Sidebar()
    }
}
