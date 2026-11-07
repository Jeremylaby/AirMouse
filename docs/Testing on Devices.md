# Testing on Devices

Testing on emulators is basically useless for our app (no real motion sensors), so we test on real devices.

- Krzysztof Swędziął is responsible for iOS
- Stanisław Barycki is responsible for Android

All commands below are run from the `mobile/` directory. Finish the setup from [`mobile/README.md`](../mobile/README.md) first.

The first build (`npm run android` / `npm run ios`) compiles the native app, installs it on the phone and starts Metro. After that, as long as no native code changed, `npm start` is enough – just open the installed app on the phone.

## Android

You need Android Studio installed (it ships the Android SDK and `adb`). Make sure `adb` is in your `PATH`:

```sh
adb version
```

### 1. Enable Developer options

Go to **Settings > About phone > Software information** and tap **Build number** 7 times. After this, **Developer options** shows up in your settings.

> Menu names differ between vendors. The path above is for Samsung (One UI).

### 2. Disable Auto Blocker (Samsung only)

Go to **Settings > Security and privacy > Auto Blocker** and turn it off. Otherwise it blocks USB debugging and installing the app.

### 3. Connect the phone

You can connect through USB or through Wi-Fi.

#### Option A: USB

1. Go to **Settings > Developer options** and enable **USB debugging**.
2. Plug the phone in and accept the **Allow USB debugging?** prompt on the phone.
3. Check that the device is visible:

   ```sh
   adb devices
   ```

   It should be listed as `device` (not `unauthorized`).

#### Option B: Wi-Fi

The phone and your computer have to be on the same Wi-Fi network.

1. Go to **Settings > Developer options** and enable **Wireless debugging**.
2. Tap **Wireless debugging > Pair device with pairing code**. The phone shows an IP address with a port and a 6-digit code.
3. Pair the device:

   ```sh
   adb pair <ip>:<pairing-port>
   # Enter pairing code: <6-digit code>
   ```

4. Connect to the device. Use the IP and port shown on the main **Wireless debugging** screen – this port is **different** from the pairing port:

   ```sh
   adb connect <ip>:<port>
   ```

5. Check that the device is visible:

   ```sh
   adb devices
   ```

Pairing is needed only once. Later you only run `adb connect` (the port changes every time wireless debugging is turned on).

### 4. Forward the Metro port

```sh
adb reverse tcp:8081 tcp:8081
```

This forwards port `8081` on the phone to port `8081` on your computer, so the app can reach Metro at `localhost:8081` regardless of the network (works for both USB and Wi-Fi). Run it again every time the phone reconnects.

### 5. Build and run

First time, or after a native change:

```sh
npm run android
```

If more than one device is connected, pick one with `npm run android -- --device`.

Next time, when the app is already installed:

```sh
npm start
```

and open the app on the phone.

### Troubleshooting

- `unauthorized` in `adb devices` – unplug the phone, run `adb kill-server`, plug it back and accept the prompt.
- App shows "Unable to load script" / can't connect to Metro – check that Metro is running and run `adb reverse tcp:8081 tcp:8081` again.

## iOS

You need a Mac with Xcode and CocoaPods installed and an Apple ID. A free Apple ID is enough, but the app installed with it expires after 7 days and has to be reinstalled.

### 1. Add your Apple ID to Xcode

Open **Xcode > Settings > Accounts**, click **+** and sign in with your Apple ID.

### 2. Trust the computer

Plug the iPhone in with a cable, unlock it and tap **Trust** on the **Trust This Computer?** prompt.

### 3. Enable Developer Mode

Go to **Settings > Privacy & Security > Developer Mode**, turn it on and restart the phone. After the restart confirm **Turn On**.

> The option shows up only after the phone was connected to a Mac with Xcode at least once.

### 4. (Optional) Connect through Wi-Fi

After the first cable connection you can switch to Wi-Fi: open **Xcode > Window > Devices and Simulators**, select the iPhone and check **Connect via network**. The phone and the Mac have to be on the same Wi-Fi network.

### 5. Build and run

First time, or after a native change:

```sh
npm run ios -- --device
```

Pick your iPhone from the list. If the CLI asks for a development team, choose your Apple ID (Personal Team).

If signing fails, open the project in Xcode and set it manually:

```sh
xed ios
```

Select the `mobile` target > **Signing & Capabilities**, check **Automatically manage signing** and pick your team. With a free account the bundle identifier (`com.anonymous.mobile`) may already be taken – change it to something unique (e.g. `com.<yourname>.airmouse`), also in `app.json` under `ios.bundleIdentifier` so the next native rebuild doesn't overwrite it.

Next time, when the app is already installed:

```sh
npm start
```

and open the app on the phone.

### 6. Trust the developer certificate

On the first launch iOS blocks the app with **Untrusted Developer**. Go to **Settings > General > VPN & Device Management**, select your Apple ID and tap **Trust**.

### Metro connection

iOS has no equivalent of `adb reverse`, so the app connects to Metro over the network using your Mac's IP address. This means the iPhone and the Mac **must be on the same Wi-Fi network**, even when the phone is connected with a cable. When the app asks for **Local Network** access, allow it.

### Troubleshooting

- App can't connect to Metro – check that both devices are on the same network (eduroam and other public networks often block device-to-device traffic; use a phone hotspot or home Wi-Fi instead) and that Local Network access is allowed in **Settings > Privacy & Security > Local Network**.
- iPhone not listed in `npm run ios -- --device` – unlock the phone, reconnect the cable and check that Developer Mode is on.
