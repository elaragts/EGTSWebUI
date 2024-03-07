mod router;
mod components;

use yew::prelude::*;
use yew_router::prelude::*;
use crate::components::organisms::navbar::Navbar;
use crate::router::{Route, switch};
pub enum Msg {
}

#[function_component(App)]
pub fn app() -> Html {
    html! {
        <BrowserRouter>
            <Navbar />
            <Switch<Route> render={switch} />
        </BrowserRouter>
    }
}