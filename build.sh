#!/usr/bin/env bash
# pwd

echo "Compiling..."
go build  -o ksctl -ldflags "-w" main.go

echo "Compressing..."
upx ksctl
