# Rust developer notes for VS Code

This document is a collection of notes used in getting developers on Windows and OSX up and running on Rust with VS Code

## Getting started

This is an augmented getting started taken mostly from the Rust add-on for VS Code

1. Install [Rustup](https://rustup.rs/)
2. Install [VS Code Add In forRust](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust)
3. Use [Cargo](https://github.com/rust-lang/cargo/) to [create a new project](https://doc.rust-lang.org/book/ch01-03-hello-cargo.html)

Replace all code in main.rs with the following:

```Rust
use std::env;

fn main() {
    println!("Hello, world!");
    let label = "Curent Directory";
    let dir = env::current_dir().unwrap();
    println!("{} {}", label, dir.display());
}
```

## Running and Debugging

A number of steps and guidance have been taken from [here](http://thiago.rocks/view/20200512_vscode_with_rust).

1 Install this VS Code [add-in](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) to enable debugging
2 Create a launch.json file from the debug section
3 Paste this json into file:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "type": "lldb",
            "request": "launch",
            "name": "Debug",
            "program": "${workspaceFolder}/src/rust/hello-world/target/debug/hello-world",
            "windows": {
                "program": "${workspaceFolder}/src/rust/hello-world/target/debug/hello-world.exe"
            },
            "args": [],
            "cwd": "${workspaceFolder}",
            "stopOnEntry": false,
            "sourceLanguages": [
              "rust"
            ]
        }
    ]
}
```

Now place a break point on the first line and press run (F5) from the debug add in