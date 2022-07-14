#[derive(Clone)]
struct Person {
    name: String,
    // age: i32,
    // email: String,
    // password: String
}

fn main() {
    let user = Person {
        name: String::from("yuri"),
    };
    println!("Hello from {}", user.name);
}
