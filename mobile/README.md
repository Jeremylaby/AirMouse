# AirMouse – mobile

## Setting up for the first time

Requires Node.js `24.13.1` and Java `17` (pinned in [`.tool-versions`](../.tool-versions); with asdf just run `asdf install`), plus Android Studio (Android) and Xcode with CocoaPods (iOS).

```sh
cd mobile
npm install
```

## Running

```sh
npm run android   # build & run debug app on Android emulator/device
npm run ios       # build & run debug app on iOS simulator
npm start         # start Metro only (app already installed)
```

Native rebuild (`npm run android` / `npm run ios`) is only needed after adding a native library or changing `app.json`; otherwise `npm start` is enough.

## Custom npm registry

If your global `~/.npmrc` points to a company registry, you can use [direnv](https://direnv.net) to switch to the public npm registry only inside your study directory:

```sh
brew install direnv
echo 'eval "$(direnv hook zsh)"' >> ~/.zshrc
echo 'export NPM_CONFIG_REGISTRY=https://registry.npmjs.org/' > ~/path/to/Studies/.envrc
direnv allow ~/path/to/Studies
```

To check it: `npm config get registry` should print `https://registry.npmjs.org/`.
