# AirMouse – desktop

## Setting up for the first time

Requires macOS, Xcode Command Line Tools (C compiler used by robotgo) and Go `<version>` (pinned in [`.tool-versions`](../.tool-versions)).

```sh
xcode-select --install
```

Then install Go in one of these ways:

- **With asdf** (recommended, uses the pinned version):

  ```sh
  asdf plugin add golang   # first time only
  asdf install
  ```

- **Without asdf**: `brew install go` or the installer from [go.dev/dl](https://go.dev/dl/). Any recent Go version should work, but the pinned one is the tested one.

Check with `go version`.

Dependencies are downloaded automatically on the first `go run` / `go build`.

## Running

```sh
cd desktop
go run ./cmd/airmouse                     # build & run
go build -o bin/airmouse ./cmd/airmouse   # build binary to bin/
go test ./...                             # run unit tests (no real cursor movement)
```

You can also use make commands

```sh
make run
make build
make test
make tidy
```
## macOS permissions

The app needs **Accessibility** permission to move the cursor, click and press keys. Without it, nothing happens and no error is shown.

When running from a terminal, the permission belongs to the terminal app (Terminal, iTerm, VS Code etc.), not to the binary. Add it in *System Settings → Privacy & Security → Accessibility* and restart the terminal.

MacOs will ask you for permission if you run exmaple main code first time. There is posibility that you will have to add this permission to code editor you use too.

## Dependencies

```sh
go get <module>   # add or update a dependency
go mod tidy       # sync go.mod and go.sum with the imports in the code
```

Run `go mod tidy` after adding or removing imports: it adds missing dependencies, removes unused ones and updates checksums in `go.sum`. Commit both `go.mod` and `go.sum`.
