use yew::prelude::*;
use crate::router::Route;
use crate::components::atoms::navitem::NavItem;
use crate::components::atoms::navtitle::NavTitle;
#[function_component(Navbar)]
pub fn navbar() -> Html {

//    html! {
//     <nav class="sticky top-0 z-999 bg-[#7D0A0A] h-15"> // Replace with your desired navbar height
//         <div class="flex justify-between items-center mx-auto p-4 h-full">
//
//         </div>
//        <div class="hidden w-full md:block md:w-auto" id="navbar-default">
//       <ul class="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
//         <NavItem title="Home" to={Route::Home} />
//       </ul>
//     </div>
//     </nav>
// }
    html! {
    <nav class="flex flex-col text-center sm:flex-row sm:text-left sm:justify-between py-4 px-6 bg-[#0F4C75] shadow sm:items-baseline w-full">
  <div class="mb-2 sm:mb-0">
    <NavTitle />
    </div>
  <div class="text-xl flex flex-row p-4 mb:2 md:p-0 self-center space-x-4">
    <NavItem title="Leaderboard" to={Route::Home} />
    <NavItem title="Dashboard" to={Route::Home} />
    <NavItem title="Profile" to={Route::Login} />
  </div>
</nav>
    }
}

