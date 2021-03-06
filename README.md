# keyboard

## ubuntu dev env

```bash
sudo apt install gcc libc6-dev
sudo apt install libx11-dev xorg-dev libxtst-dev libpng++-dev
sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt install libxkbcommon-dev
sudo apt install xsel xclip
```

## Cross Compilation

### ubuntu install

```
sudo apt-get install gcc-multilib
sudo apt-get install gcc-mingw-w64
sudo apt install libz-mingw-w64-dev
```

### build command

```bash
./build_windows.sh main.go 1.0.0 keyboard_x86.exe
```

**need zlib1.dll to run**

http://www.winimage.com/zLibDll/
