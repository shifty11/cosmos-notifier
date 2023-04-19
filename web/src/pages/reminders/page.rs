use crate::components::sidebar::Sidebar;
use sycamore::prelude::*;

#[component]
pub fn Reminders<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        Sidebar()
    }
}
