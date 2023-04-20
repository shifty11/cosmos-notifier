use sycamore::prelude::*;

use crate::{AppRoutes, AppState};

#[component]
pub fn Header<G: Html>(cx: Scope) -> View<G> {
    let app_state = use_context::<AppState>(cx);

    view!(cx,
        div(class="fixed w-full flex items-center justify-between h-14 text-white z-10") {
            div(class="flex items-center justify-start pl-4 pl-3 w-14 md:w-64 h-14 bg-blue-800 dark:bg-gray-800 border-none") {
                img(class="w-7 h-7 md:w-10 md:h-10 mr-2 rounded-md overflow-hidden", src=app_state.get_user_avatar())
                span(class="hidden md:block") { (app_state.get_user_name()) }
            }
            div(class="flex flex-grow justify-between items-center h-14 bg-blue-800 dark:bg-gray-800 header-right") {
                div(class="outline-none focus:outline-none") {}
                div(class="w-full pl-3 text-sm text-black outline-none focus:outline-none bg-transparent" ) {}
            }
            div(class="flex justify-between items-center h-14 bg-blue-800 dark:bg-gray-800 header-right") {
                ul(class="flex items-center") {
                    li {
                        button(aria-hidde="true", class="group w-9 h-9 transition-colors duration-200 rounded-full shadow-md bg-blue-200 hover:bg-blue-200 dark:bg-gray-50 dark:hover:bg-gray-200 text-gray-900 focus:outline-none") {
                            i(class="fas fa-bell text-lg") {}
                        }
                    }
                    li {
                        div(class="block w-px h-6 mx-3 bg-gray-400 dark:bg-gray-700") {}
                    }
                    li {
                        button(class="flex items-center mr-4 hover:text-blue-100", on:click=move |_| app_state.logout()) {
                            span(class="inline-flex mr-1") {
                                i(class="fas fa-sign-out-alt text-xl") {}
                            }
                            "Logout"
                        }
                    }
                }
            }
        }
    )
}

fn highlight_active_route(route: &AppRoutes, current_route: &AppRoutes) -> String {
    if route.to_string() == current_route.to_string() {
        "text-blue-500".to_string()
    } else {
        "".to_string()
    }
}

#[component]
pub fn Sidebar<G: Html>(cx: Scope) -> View<G> {
    let app_state = use_context::<AppState>(cx);

    let a_class = "relative flex flex-row items-center h-11 focus:outline-none hover:bg-blue-800 dark:hover:bg-gray-600 text-white-600 hover:text-white-800 border-l-4 border-transparent hover:border-blue-500 dark:hover:border-gray-800 pr-6";
    let span_icon_class = "inline-flex justify-center items-center ml-4";
    let span_text_class = "ml-2 text-sm tracking-wide truncate";

    view! { cx,
        div(class="fixed flex flex-col top-14 left-0 w-14 hover:w-64 md:w-64 bg-blue-900 dark:bg-gray-900 h-full text-white transition-all duration-300 border-none z-10 sidebar") {
            div(class="overflow-y-auto overflow-x-hidden flex flex-col justify-between flex-grow") {
                ul(class="flex flex-col py-4 space-y-1") {
                    li(class="px-5 hidden md:block") {
                        div(class="flex flex-row items-center h-8") {
                            div(class="text-sm font-light tracking-wide text-gray-400 uppercase") { "Main" }
                        }
                    }
                    li() {
                        a(href=AppRoutes::Home, class=format!("{} {}", a_class, highlight_active_route(&AppRoutes::Home, app_state.route.get().as_ref()))) {
                            span(class=span_icon_class) {
                                i(class="fas fa-home") {}
                            }
                            span(class=span_text_class) { "Home" }
                        }
                    }
                    li() {
                        a(href=AppRoutes::Overview, class=format!("{} {}", a_class, highlight_active_route(&AppRoutes::Overview, app_state.route.get().as_ref()))) {
                            span(class=span_icon_class) {
                                i(class="fas fa-chart-line") {}
                            }
                            span(class=span_text_class) { "Overview" }
                        }
                    }
                    li() {
                        a(href=AppRoutes::Reminders, class=format!("{} {}", a_class, highlight_active_route(&AppRoutes::Reminders, app_state.route.get().as_ref()))) {
                            span(class=span_icon_class) {
                                i(class="fas fa-bell") {}
                            }
                            span(class=span_text_class) { "Reminders" }
                        }
                    }
                    li() {
                        a(href=AppRoutes::Communication, class=format!("{} {}", a_class, highlight_active_route(&AppRoutes::Communication, app_state.route.get().as_ref()))) {
                            span(class=span_icon_class) {
                                i(class="fas fa-message") {}
                            }
                            span(class=span_text_class) { "Communication" }
                        }
                    }
                }
            }
        }
    }
}

#[component(inline_props)]
pub fn LayoutWrapper<'a, G: Html>(cx: Scope<'a>, children: Children<'a, G>) -> View<G> {
    let children = children.call(cx);
    view! { cx,
        div(class="min-h-screen flex flex-col flex-auto flex-shrink-0 antialiased bg-white dark:bg-gray-700 text-black dark:text-white") {
            Header{}
            Sidebar{}
            div(class="h-full ml-14 mt-14 mb-10 md:ml-64") {
                (children)
            }
        }
    }
}
