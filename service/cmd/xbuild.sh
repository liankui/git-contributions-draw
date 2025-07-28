#!/bin/bash
set -e

server=git-contributions-draw-server

case "$1" in
linux | windows | darwin) OS="$1" ;;
mac) OS="darwin" ;;
"") OS="" ;;
*) echo "illegal option: $1" && exit 1 ;;
esac

case "$2" in
amd64 | arm64) ARCH="$2" ;;
"") ARCH="" ;;
*) echo "illegal option: $2" && exit 1 ;;
esac

git_hash=$(git rev-list -1 HEAD)
git_branch=$(git branch --show-current)
build_time=$(date "+%Y-%m-%d %H:%M:%S %Z")

set -x
env CGO_ENABLED=0 GOOS="$OS" GOARCH="$ARCH" go build -o "$server" -ldflags "-s -w -X main.gitHash=$git_hash -X main.gitBranch=$git_branch -X 'main.buildTime=$build_time'"
