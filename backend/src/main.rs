mod api;
mod db;

use rocket::fs::NamedFile;
use serde_derive::Deserialize;
use rocket::response::status::NotFound;
use std::path::PathBuf;
use std::process::exit;
use std::fs;
use toml;
use db::{init_pool};
#[derive(Deserialize)]

struct Data {
    config: Config,
}

#[derive(Deserialize)]
struct Config {
    database: String
}


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
    let config_file = "config/config.toml";
    let contents = match fs::read_to_string(config_file) {
        Ok(c) => c,
        Err(_) => {
            eprintln!("Could not read file `{}`", config_file);
            exit(1);
        }
    };
    let data: Data = match toml::from_str(&contents) {
        Ok(d) => d,
        Err(_) => {
            eprintln!("Unable to load data from `{}`", config_file);
            exit(1);
        }
    };

    let pool = init_pool(data.config.database);
    rocket::build()
        .mount("/", routes![index, static_files])
        .mount("/api", api::routes())
        .manage(pool)
}
