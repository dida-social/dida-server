#!/usr/bin/env bash

on_interrupt() {
  echo "server stopped"
}

go run cmd/server/main.go
trap on_interrupt SIGINT