use yew::prelude::*;

#[function_component(Home)]
pub fn home() -> Html {
    html! {
        <p class={"text-3xl underline"}>{"Welcome home!"}</p>
    }
}