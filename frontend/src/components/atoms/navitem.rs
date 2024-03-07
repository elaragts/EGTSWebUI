use yew::prelude::*;
use yew_router::prelude::*;
use crate::router::Route;

#[derive(Properties, PartialEq)]
pub struct Props {
    pub title: String,
    pub to: Route
}

#[function_component(NavItem)]
pub fn navitem(props: &Props) -> Html {
    html! {
            <Link<Route> to={props.to.clone()} classes="block py-2 px-3 text-[#BBE1FA] rounded md:bg-transparent md:p-0 hover:text-[#FF6000]">{&props.title}
        </Link<Route>>
    }
}