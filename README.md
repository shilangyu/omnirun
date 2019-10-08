# omnirun

Run your source code directly

```sh
go get github.com/shilangyu/omnirun
```

---

### run

As arg:

```sh
omnirun main.cpp
```

From stdin:

```sh
"main.cpp" | omnirun
```

### config

Edit config:

```sh
vim $(omnirun config)
```

- `exts`: array of extensions
- `run`: array of commands to run:
  - `$or_file`: path of the source file
