#![allow(non_snake_case)]

use log::debug;
use log::Level;
use sycamore::prelude::*;

#[component]
fn App<G: Html>(cx: Scope) -> View<G> {
    view! {cx,
        h1 { "Hello, World!" }
    }
}

fn main() {
    console_error_panic_hook::set_once();
    console_log::init_with_level(Level::Debug).unwrap();
    debug!("Console log level set to debug");

    sycamore::render(|cx| view! { cx, App()});
}