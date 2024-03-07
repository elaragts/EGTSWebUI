use yew_router::prelude::*;
use yew::prelude::*;
use crate::components::pages::home::Home;
#[derive(Debug, Clone, PartialEq, Routable)]
pub enum Route {
    #[at("/")]
    Home,
    #[at("/login")]
    Login,
    #[at("/register")]
    Register,
}

pub fn switch(route: Route) -> Html {
    match route {
        Route::Home => html! {<Home />},
        Route::Login => html! {<p>{"Login!!"}</p>},
        Route::Register => html! {<p>{"Register"}</p>},
    }
}