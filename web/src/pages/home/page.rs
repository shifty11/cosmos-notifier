use crate::AppRoutes;
use sycamore::prelude::*;
use crate::components::sidebar::Sidebar;

#[component]
pub fn Home<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        Sidebar()
    }
}
