#!/usr/bin/env bash

set -e

echo '=> Resetting...'
rm -rf *.go go.mod go.sum .idea/ .git/

echo '=> Initializing Go module...'
go mod init github.com/mdwhatcott/tcr-environment-controller-kata

echo '=> Initializing git repository...'
git init
git add .
git commit -m "Initial commit."
git remote add origin git@github.com:mdwhatcott/tcr-environment-controller-kata.git
git push -f origin main

echo '=> Starting IDE...'
goland .

echo '=> Starting TCR Stopwatch...'
tcr-stopwatch

echo '=> Pushing final state...'
git push -f origin main

echo '=> Finished.'
