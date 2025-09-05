# GTIN Checksum Calculator

Minimal command-line tool for calculating and validating GTIN checksums according to GS1 standard.

## Usage

```bash
./gtin
1=Calculate 2=Validate q=Quit: 1
Digits: 400638133315
Checksum: 2
Complete: 4006381333152
```

## Features

- Calculate GTIN checksums (any length)
- Validate existing GTINs
- Supports EAN-8, EAN-13, UPC-A, ITF-14, and custom formats
- Zero dependencies, single binary

## Install

Download binary from [releases](https://github.com/noahjeana/gtin-calculator/releases) or build:

```bash
go build gtin.go
```

## Cross-compile

```bash
GOOS=linux GOARCH=amd64 go build -o gtin-linux gtin.go
GOOS=windows GOARCH=amd64 go build -o gtin-windows.exe gtin.go  
GOOS=darwin GOARCH=amd64 go build -o gtin-macos gtin.go
```


