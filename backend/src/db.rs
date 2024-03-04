use r2d2_sqlite::SqliteConnectionManager;
use r2d2::{Pool};

pub type SqlitePool = Pool<SqliteConnectionManager>;

pub fn init_pool(db_url: String) -> SqlitePool {
    let manager = SqliteConnectionManager::file(db_url);
    Pool::new(manager).expect("database pool")
}