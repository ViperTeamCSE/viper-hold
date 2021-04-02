# Rust developer notes for VS Code

This document is a collection of notes used in getting developers on Windows and OS X up and running on Rust with VS Code. **PLEASE NOTE: If you are on OS X, you will need to install both [XCode](https://developer.apple.com/xcode/) and the XCode [command line tools](https://developer.apple.com/download/more/?=xcode)**

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

## Building, Testing, and Linting

Use of tasks in VS Code can accomplish building, testing, and linting. While you are free to use your own, an example tasks.json file is located below:

```json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "Run Cargo Build Full",
            "type": "shell",
            "command": "cargo build",
            "presentation": {
              "reveal": "always",
              "panel": "new"
            }
        },
        {
            "label": "Run Cargo Tests Full",
            "type": "shell",
            "command": "cargo test",
            "presentation": {
            "reveal": "always",
            "panel": "new"
            },
        },
        {
            "label": "Run Cargo Format Full",
            "type": "shell",
            "command": "cargo fmt",
            "presentation": {
              "reveal": "always",
              "panel": "new"
            }
        }
    ]
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
            "name": "Debug Hello World",
            "program": "${workspaceFolder}/target/debug/hello-world",
            "windows": {
                "program": "${workspaceFolder}/target/debug/hello-world.exe"
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
