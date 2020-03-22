# omnirun

[![](https://github.com/shilangyu/omnirun/workflows/ci/badge.svg)](https://github.com/shilangyu/omnirun/actions)

Want to quickly run a single source file without having to remember the compilation steps?

```sh
go get github.com/shilangyu/omnirun
```

Or get from [releases](https://github.com/shilangyu/omnirun/releases)

---

### ⇒ run

As arg:

```sh
omnirun main.cpp
omnirun main.c
omnirun main.rs
omnirun main.js
omnirun main.go
omnirun main.fish
omnirun main.py
```

From stdin:

```sh
echo "main.cpp" | omnirun -
```

---

### 🛠️ config

Edit config:

```sh
vim $(omnirun config)
```

- `exts`: array of extensions
- `run`: array of commands to run:
  - `$or_file`: path of the source file

---

### ✔ included runners

These are the presets, you can change them any time with [#config](#%f0%9f%9b%a0%ef%b8%8f-config)

| file type | compiler/interpreter used                                   |
| --------- | ----------------------------------------------------------- |
| `js`      | [node](https://nodejs.org/en/)                              |
| `py`      | [python](https://www.python.org/)                           |
| `rs`      | [rustc](https://doc.rust-lang.org/rustc/what-is-rustc.html) |
| `go`      | [go](https://golang.org/)                                   |
| `bf`      | [brainfuck](https://github.com/shilangyu/brainfuck)         |
| `c`       | [gcc](https://gcc.gnu.org/)                                 |
| `cpp`     | [g++](https://gcc.gnu.org/)                                 |
| `fish`    | [fish](http://fishshell.com/)                               |
| `ts`      | [ts-node](https://github.com/TypeStrong/ts-node)            |

Gladly accepting PRs with new presets
