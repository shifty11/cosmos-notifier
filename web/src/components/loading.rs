use sycamore::prelude::*;

#[component]
pub fn LoadingSpinner<G: Html>(cx: Scope) -> View<G> {
    view!(
        cx,
        div(class="flex justify-center items-center h-full") {
            div(class="animate-spin rounded-full h-32 w-32 border-t-2 border-b-2 border-gray-900") {}
        }
    )
}
