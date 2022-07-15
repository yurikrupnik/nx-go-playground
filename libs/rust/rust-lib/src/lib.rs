#![allow(dead_code, unused_variables)]

extern crate core;

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

fn largest<T: std::cmp::PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];
    for item in list {
        if item > largest {
            largest = item
        };
    }
    largest
}

fn sum_loops(n: i32) -> i32 {
    let mut sum = 0;
    for i in 1..n {w
        sum += i;
    }
    sum
    // sum
}

fn add_one(n: i32) -> i32 {
    12 + n
}

fn connect_to_database() -> Status {
    Status::Connected
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
