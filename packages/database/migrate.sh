#!/bin/bash

# Load .env
if [ -f .env ]; then
  export $(cat .env | xargs)
fi

# Migration commands
case "$1" in
  create)
    goose create "$2" sql
    ;;
  up)
    goose up
    ;;
  down)
    goose down
    ;;
  status)
    goose status
    ;;
  *)
    echo "Usage: $0 {create|up|down|status}"
    exit 1
esac
