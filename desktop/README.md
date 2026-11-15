# AirMouse – desktop

## Installing (without Go)

Builds are published on [GitHub Releases](https://github.com/Jeremylaby/AirMouse/releases) as `airmouse-desktop-<version>-macos.zip`. The binary is universal, so it runs on both Apple Silicon and Intel Macs, and requires macOS 13 or newer.

1. Download the zip from the latest `desktop-v*` release and unzip it (double-click in Finder), or from a terminal:

   ```sh
   gh release download desktop-v<version> --repo Jeremylaby/AirMouse --pattern '*-macos.zip'
   unzip airmouse-desktop-<version>-macos.zip
   ```

2. Allow the app to run. It isn't signed with an Apple Developer ID (that needs a paid account), so Gatekeeper blocks files downloaded with a browser. Do one of these:
   - Run `./airmouse` once. macOS shows "Apple could not verify…". Click **Done**, go to _System Settings → Privacy & Security_, scroll down and click **Open Anyway** next to the `airmouse` message, then run it again.
   - Or remove the quarantine flag that the browser added:

     ```sh
     xattr -d com.apple.quarantine airmouse
     ```

   Files downloaded with `gh release download` or `curl` aren't quarantined, so this step isn't needed for them.

3. Run it from a terminal:

   ```sh
   ./airmouse
   ```

4. Grant the **Accessibility** permission when macOS asks for it (see [macOS permissions](#macos-permissions)).

To check which version you have, look at the first line the app prints.

## Releasing

The [`desktop-release.yml`](../.github/workflows/desktop-release.yml) workflow builds the app and creates a GitHub Release. The version comes from the tag, there is no version file. Release in one of two ways:

- Push a tag:

  ```sh
  git tag desktop-v0.2.0
  git push origin desktop-v0.2.0
  ```

- Or go to _Actions → Desktop – GitHub Release → Run workflow_, pick a branch and enter the version (e.g. `0.2.0`). The tag is created on the latest commit of that branch.

Versions follow `MAJOR.MINOR.PATCH`, optionally with a suffix like `0.2.0-beta.1`.

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
