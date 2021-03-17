use std::env;

fn main() {
    let mut s = String::new();
    let hello = say_hello(&mut s);
    println!("{}", hello);
    let label = "Curent Directory";
    let dir = env::current_dir().unwrap();
    println!("{} {}", label, dir.display());
}

fn say_hello(s: &mut String) -> &str {
    s.push_str("Hello, world!!");
    &s[..]
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_hello() {
        let mut s = String::new();
        assert_eq!(say_hello(&mut s), "Hello, world!!");
    }
}
