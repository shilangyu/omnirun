# omnirun

![](https://github.com/shilangyu/omnirun/workflows/ci/badge.svg)

Run your source code directly

```sh
go get github.com/shilangyu/omnirun
```

Or get from [releases](https://github.com/shilangyu/omnirun/releases)

---

### run

As arg:

```sh
omnirun main.cpp
```

From stdin:

```sh
echo "main.cpp" | omnirun -
```

### config

Edit config:

```sh
vim $(omnirun config)
```

- `exts`: array of extensions
- `run`: array of commands to run:
  - `$or_file`: path of the source file
