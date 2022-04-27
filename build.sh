#!/bin/bash

rm -f guestbook
env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o guestbook

docker build -t guestbook:v6 .
