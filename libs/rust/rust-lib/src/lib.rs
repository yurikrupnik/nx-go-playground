#![allow(dead_code, unused_variables)]

pub fn rust_lib() -> String {
    "rust_lib".into()
}

struct Credentials {
    username: String,
    password: String,
}
enum Status {
    Connected,
    Interrupted,
}

fn connect_to_database() -> Status {
    return Status::Connected;
}
fn authenticate(creds: Credentials) {
    if let Status::Connected = connect_to_database() {
        login(creds);
    }
}

fn get_user() {
    // get user from db
}

fn login(creds: Credentials) {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(rust_lib(), "rust_lib".to_string());
    }
}
