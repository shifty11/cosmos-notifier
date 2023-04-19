use std::rc::Rc;
use gloo_history::History;
use log::debug;
use sycamore::prelude::*;
use sycamore_router::Route;

use crate::{AppRoutes, AppState};

fn get_class_for_route(route: &AppRoutes, current_route: &AppRoutes) -> String {
    if route.to_string() == current_route.to_string(){
        "text-blue-500".to_string()
    } else {
        "".to_string()
    }
}

#[component]
pub fn Sidebar<G: Html>(cx: Scope) -> View<G> {
    let app_state = use_context::<AppState>(cx);

    view! {cx,
        div(class="fixed top-0 left-0 h-screen w-64 py-6 px-4 bg-gray-800 text-white") {
            ul(class="flex flex-col space-y-2") {
                li() {
                    a(href=AppRoutes::Home, class=format!("p-3 rounded-md hover:bg-gray-600 block text-base font-medium {}",
                        get_class_for_route(&AppRoutes::Home, app_state.route.get().as_ref()))) {
                        i(class="fas fa-home") {}
                        span(class="ml-3") { "Home" }
                    }
                }
                li() {
                    a(href=AppRoutes::Overview, class=format!("p-3 rounded-md hover:bg-gray-600 block text-base font-medium {}",
                    get_class_for_route(&AppRoutes::Overview, app_state.route.get().as_ref()))) {
                        i(class="fas fa-chart-line") {}
                        span(class="ml-3") { "Overview" }
                    }
                }
                li() {
                    a(href=AppRoutes::Reminders, class=format!("p-3 rounded-md hover:bg-gray-600 block text-base font-medium {}",
                    get_class_for_route(&AppRoutes::Reminders, app_state.route.get().as_ref()))) {
                        i(class="fas fa-bell") {}
                        span(class="ml-3") { "Reminders" }
                    }
                }
                li() {
                    a(href=AppRoutes::Communication, class=format!("p-3 rounded-md hover:bg-gray-600 block text-base font-medium {}",
                    get_class_for_route(&AppRoutes::Communication, app_state.route.get().as_ref()))) {
                        i(class="fas fa-user") {}
                        span(class="ml-3") { "Profile" }
                    }
                }
            }
        }
    }
}
