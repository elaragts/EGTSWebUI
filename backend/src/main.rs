mod api;

use rocket::fs::NamedFile;
use rocket::response::status::NotFound;
use std::path::PathBuf;

//Import the rocket macros
#[macro_use]
extern crate rocket;

// Return the index file as a Rocket NamedFile
async fn get_index() -> Result<NamedFile, NotFound<String>> {
    NamedFile::open("../ui/dist/index.html")
        .await
        .map_err(|e| NotFound(e.to_string()))
}

//Create a route for any url that is a path from the /
#[get("/<path..>")]
async fn static_files(path: PathBuf) -> Result<NamedFile, NotFound<String>> {
    let path = PathBuf::from("../ui/dist").join(path);
    match NamedFile::open(path).await {
        Ok(f) => Ok(f),
        Err(_) => get_index().await,
    }
}

// Return the index when the url is /
#[get("/")]
async fn index() -> Result<NamedFile, NotFound<String>> {
    get_index().await
}
// Finlay start the web sever using the launch macro.
#[launch]
fn rocket() -> _ {
    // You must mount the static_files route
    rocket::build()
        .mount("/", routes![index, static_files])
        .mount("/api", api::routes())
}
