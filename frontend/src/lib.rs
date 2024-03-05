use yew::prelude::*;
pub struct App {
}
pub enum Msg {
}
impl Component for App {
    type Message = Msg;
    type Properties = ();

    //Create a new App
    fn create(_ctx: &Context<Self>) -> Self {
        App {
        }
    }

    fn update(&mut self, _ctx: &Context<Self>, msg: Self::Message) -> bool {
        false
    }

    fn view(&self, ctx: &Context<Self>) -> Html {
        //Creates The HTML that will show up in the browser.
        html! {
             <p class={"text-3xl underline"}>{"Hello World!s?ss"}</p>
        }
    }
}