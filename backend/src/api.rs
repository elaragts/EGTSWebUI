use rocket::{get, routes};
use rocket::State;
use rocket::serde::{json::Json, Deserialize, Serialize};
use db::SqlitePool;
use crate::db;

#[derive(Serialize, Deserialize)]
#[serde(crate = "rocket::serde")]
struct Leaderboard {
    name: String,
    best_score: i32,
    best_crown: i32,
    best_score_rank: i32
}


#[get("/hello")]
fn hello() -> &'static str {
    "Hello, API!"
}

#[get("/leaderboard?<song_id>&<difficulty>")]
fn leaderboard(pool: &State<SqlitePool>, song_id: i32, difficulty: i32) -> Json<Vec<Leaderboard>> {
    let conn = pool.get().expect("get db connection");
    let mut stmt = conn
        .prepare("SELECT ud.MyDonName, sbd.BestScore, sbd.BestCrown, sbd.BestScoreRank
FROM SongBestData sbd
         INNER JOIN UserData ud ON sbd.Baid = ud.Baid
         WHERE SongID = ?1
  AND Difficulty = ?2
ORDER BY sbd.BestScore DESC
LIMIT 10").expect("");

    let leaderboard: Vec<Leaderboard> = stmt
        .query_map(&[&song_id, &difficulty], |row| {
            Ok(Leaderboard {
                name: row.get(0)?,
                best_score: row.get(1)?,
                best_crown: row.get(2)?,
                best_score_rank: row.get(3)?
            })
        }).expect("query map")
        .map(|result| result.expect("result mapping"))
        .collect();


    Json(leaderboard)

}
pub fn routes() -> Vec<rocket::Route> {
    routes![hello, leaderboard]
}
