use gloo_timers::future::TimeoutFuture;
use js_sys;
use log::debug;
use sycamore::futures::spawn_local_scoped;
use sycamore::motion::create_raf;
use sycamore::prelude::*;
use tonic::Status;

use crate::{AppState, InfoLevel};

#[derive(Debug, Clone, PartialEq)]
pub struct IndexedItem<T> {
    index: usize,
    item: T,
}

#[component]
pub fn MessageOverlay<G: Html>(cx: Scope) -> View<G> {
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
                debug!("Rendering message: {:?}", iItem.index);
                let style = format!("margin-bottom: {}rem;", 5 * iItem.index + 2);
                let color = match iItem.item.get().level {
                    InfoLevel::Info => "bg-blue-100 border-blue-500 text-blue-700",
                    InfoLevel::Success => "bg-green-100 border-green-500 text-green-700",
                    InfoLevel::Warning => "bg-yellow-100 border-yellow-500 text-yellow-700",
                    InfoLevel::Error => "bg-red-100 border-red-500 text-red-700",
                };
                let item = iItem.item.get();
                let title = item.title.clone();
                let message = item.message.clone();
                let created_at = item.created_at.clone();
                let id = item.id;

                let state = create_signal(cx, 1.0);
                let (_running, start, stop) = create_raf(cx, move || {
                    let elapsed = js_sys::Date::now() - created_at;
                    if elapsed > 7000.0 {
                        let new_state = 1.0 - (elapsed - 7000.0) / 3000.0;
                        state.set(new_state);
                        if new_state <= 0.0 {
                            app_state.remove_message(id);
                        }
                    }
                });
                start();

                view! { cx,
                    div(class=format!("fixed bottom-0 right-0 w-96 z-50 p-2 m-2 bg-white border-l-4 {}", color), style=format!("{} opacity: {}", style, state.get().as_ref())) {
                        h3(class="text-lg font-bold") { (title) }
                        p(class="text-sm") { (message) }
                        button(
                            class="absolute top-0 right-0 p-2 m-2",
                            on:click=move |_| {
                                spawn_local_scoped(cx, async move {
                                    stop();
                                    app_state.remove_message(id);
                                });
                            }
                        ) {
                            i(class="fas fa-times") {}
                        }
                    }
                }
            },
        )
    )
}

pub fn create_message(
    cx: Scope,
    title: impl Into<String>,
    message: impl Into<String>,
    level: InfoLevel,
) {
    let app_state = use_context::<AppState>(cx);
    let uuid = app_state.add_message(title.into(), message.into(), level);
    create_effect(cx, move || {
        spawn_local_scoped(cx, async move {
            debug!("wait 10 seconds before removing message");
            TimeoutFuture::new(1000 * 10).await;
            debug!("removing message");
            app_state.remove_message(uuid);
        });
    });
}

pub fn create_error_msg_from_status(cx: Scope, status: Status) {
    create_message(
        cx,
        status.code().to_string().as_str(),
        status.message(),
        InfoLevel::Error,
    );
}
