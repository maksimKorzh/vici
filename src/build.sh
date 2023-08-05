#!/bin/bash
export GOOS=windows && go build -o vici.exe *.go
export GOOS=linux && go build -o vici *.go
