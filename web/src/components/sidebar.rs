use crate::AppRoutes;
use sycamore::prelude::*;

#[component]
pub fn Sidebar<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        div(class="fixed top-0 left-0 h-screen w-64 py-6 px-4 bg-gray-800 text-white") {
            ul(class="flex flex-col space-y-2") {
                // add more space between text and icon
                li(class="p-3 rounded-md hover:bg-gray-600") {
                    a(href=AppRoutes::Home, class="block text-base font-medium") {
                        i(class="fas fa-home") {}
                        span(class="ml-3") { "Home" }
                    }
                }
                li(class="p-3 rounded-md hover:bg-gray-600") {
                    a(href=AppRoutes::Overview, class="block text-base font-medium") {
                        i(class="fas fa-chart-line") {}
                        span(class="ml-3") { "Overview" }
                    }
                }
                li(class="p-3 rounded-md hover:bg-gray-600") {
                    a(href=AppRoutes::Reminders, class="block text-base font-medium") {
                        i(class="fas fa-bell") {}
                        span(class="ml-3") { "Reminders" }
                    }
                }
                li(class="p-3 rounded-md hover:bg-gray-600") {
                    a(href=AppRoutes::Communication, class="block text-base font-medium") {
                        i(class="fas fa-comments") {}
                        span(class="ml-3") { "Communication" }
                    }
                }

            }
        }
    }
}
