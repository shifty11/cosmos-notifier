use log::{debug, error};
use sycamore::futures::spawn_local_scoped;
use sycamore::prelude::*;
use tonic::Status;

use crate::{AppState, InfoLevel};

#[derive(Debug, Clone, PartialEq)]
pub struct IndexedItem<T> {
    index: usize,
    item: T,
}

#[component]
pub fn ErrorOverlay<G: Html>(cx: Scope) -> View<G> {
    let app_state = use_context::<AppState>(cx);
    let messages = create_selector(cx, || {
        app_state
            .messages
            .get()
            .iter()
            .enumerate()
            .map(|(index, msg)| IndexedItem {
                index,
                item: msg.clone(),
            })
            .collect::<Vec<_>>()
    });

    view!(
        cx,
        Indexed(
            iterable = messages,
            view = move |cx, iItem| {
                let margin = format!("margin-bottom: {}rem;", 4 * iItem.index + 2);
                debug!("rendering error message: {:?}", iItem.index);
                let color = match iItem.item.get().level {
                    InfoLevel::Error => "bg-red-500",
                    InfoLevel::Info => "bg-blue-500",
                };
                view! { cx,
                    div(class=format!("fixed bottom-0 left-1/2 transform -translate-x-1/2 text-white p-3 rounded-md shadow-lg {color}"), style=margin) {
                        (iItem.item.get().msg)
                    }
                }
            },
        )
    )
}

pub fn create_message(cx: Scope, message: String, level: InfoLevel) {
    let app_state = use_context::<AppState>(cx);
    let uuid = app_state.add_message(message, level);
    create_effect(cx, move || {
        spawn_local_scoped(cx, async move {
            debug!("wait 10 seconds before removing message");
            gloo_timers::future::TimeoutFuture::new(1000 * 10).await;
            debug!("removing message");
            app_state.remove_message(uuid);
        });
    });
}

pub fn create_error_msg_from_status(cx: Scope, status: Status) {
    let msg = format!("{}: {}", status.code(), status.message());
    error!("Error during API call: {}", msg);
    create_message(cx, msg, InfoLevel::Error);
}
