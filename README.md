
# CLI Alpaca.tv

CLI client for Alpaca.tv

## Requirements

- git
- make
- go (v1.13 and higher)

## How to install

```bash
$ git clone https://github.com/alpaca-tv/cli.alpaca.tv.git /tmp/alpc
$ cd /tmp/alpc
$ make install
```

## How to use

Search for film

```bash
$ alpc -search "Матрица"
...
```

Search for series

```bash
# Single episode
$ alpc -search "Игра Престолов" -series -seasion 1 -episode 1
...
# All season
$ alpc -search "Игра Престолов" -series -seasion 1 -episode 0
...
# All season with specific parameters
$ alpc -search "Игра Престолов" -series -seasion 1 -episode 0 -voicecover "Дубляж" -quality 720
...
```
