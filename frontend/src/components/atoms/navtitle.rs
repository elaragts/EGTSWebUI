use yew::prelude::*;
use yew_router::prelude::*;
use crate::router::Route;

#[function_component(NavTitle)]
pub fn navtitle() -> Html {
    html! {
        <Link<Route> to={Route::Home} classes="flex items-center"> // Changed classes to class as per your last message
                <img src="static/taiko.png" class="mr-3 h-10" alt="Logo" /> // You can adjust `mr-3` for the margin-right as you need
                <span class="text-2xl font-semibold whitespace-nowrap text-[#BBE1FA] font-[Taiko] self-center"> // Removed h-10 to allow text to define its own height
                    {"Taiko Public Server"}
                </span>
            </Link<Route>>
        }
}