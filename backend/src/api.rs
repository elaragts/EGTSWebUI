// Still in src/api.rs

// Import Rocket's macros
use rocket::{get, routes};
// src/api.rs

#[get("/hello")]
pub fn hello() -> &'static str {
    "Hello, API!"
}

pub fn routes() -> Vec<rocket::Route> {
    routes![hello]
}
