# PingAlive

Tired of random network hiccups and being unsure if you lost connection? PingAlive is the solution! This lightweight utility constantly checks your network connection, giving you live feedback on whether your connection is active or down.


https://github.com/user-attachments/assets/5b335293-242a-4c99-ada4-dd85ff1da0f2


## Usage

1. Download the latest release from the [releases page](https://github.com/YonLiud/pingalive/releases).

### Windows

2. Run the following command in the directory where you downloaded the executable. You can also double-click the executable to run it.

```cmd
pingalive-windows.exe
```
3. The program will start running and will display a *OK* message if the connection is active, or a *DOWN* message if the connection is down.

### Linux

> [!WARNING]
> Unfortunately, the app requires `sudo` to run, or you can grant ICMP privileges by running `sudo setcap cap_net_raw=eip ./pingalive-linux`. This is because the app requires raw socket access to send ICMP packets.

2. Run the following commands in the directory where you downloaded the executable, replacing `pingalive-linux` with the name of the executable if it's different. You can also run the command with `sudo` if you don't want to grant the app raw socket access. The app will still work, but you will need to run it with `sudo` every time you want to use it.
```bash
chmod +x pingalive-linux
sudo ./pingalive-linux
```
3. The program will start running and will display a *OK* message if the connection is active, or a *DOWN* message if the connection is down.

## Flags & Features

- **-h**: Display the help message contaning all the flags
- **--interval**: Set the interval in milliseconds between each ping. Default is 250 milliseconds.
- **--ips**: Set the IPs to ping. Default is `8.8.8.8`. You can specify multiple IPs by separating them with a comma. like so:
```bash
./pingalive --ips 8.8.8.8,192.168.1.1,google.com
```
- **--spinner**: choose spinner style. Default is `dots`. but there is `clock` as well.
> [!NOTE]
> If you don't use any nerd fonts, the spinner will not display correctly. and **YOU MUST USE `--spinner clock`**.
> Which looks cooler anyway.
- **--msg**: Set the message to send with the ping. Default is `ping`.

## Examples

- Run the app with a 1-second interval and check connections to facebook and wikipedia.
```bash
./pingalive-linux --interval 1000 --ips facebook.com,wikipedia.org
```

- Run the app with a 100 milliseconds interval and check connections to vercel with a custom message.
```bash
./pingalive-linux --interval 100 --ips vercel.com --msg "Hello, Vercel!"
```

## Contributions

Any and all contributions are welcome! Feel free to fork the repository and submit a pull request with your changes!

> [!TIP]
> Feel free to add even MORE spinners! Just add them to the `spinners` switch case in `main.go` and submit a PR!
> I am not really a front-end guy, so design changes are more than just welcome!
