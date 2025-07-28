#!/bin/bash
set -e

CURRENT_DIR=$(cd "$(dirname "$0")" && pwd)
SERVICE_DIR=$(cd "$CURRENT_DIR"/.. && pwd)

image_prefix=
server=git-contributions-draw-server
tag=latest

cd "$SERVICE_DIR"/cmd && ./xbuild.sh linux amd64

git_hash=$(git rev-list -1 HEAD)
git_branch=$(git branch --show-current)

(
  set -x
  docker build -f "$SERVICE_DIR"/build/Dockerfile --platform linux/amd64 -t "$server:$tag" \
    --build-arg GIT_BRANCH="$git_branch" --build-arg GIT_HASH="$git_hash" "$SERVICE_DIR"
)

if [ X"" != X"$image_prefix" ]; then
  target="$image_prefix/$server:$tag"
  docker tag "$server:$tag" "$target"
  docker push "$target"
fi
