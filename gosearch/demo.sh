#!/usr/bin/env bash
#
# gosearch self-demo: builds the tool and runs a few searches over its own code.
# Usage:  ./demo.sh   (from the gosearch/ directory)
#
set -uo pipefail

echo "═══ gosearch — demo ═══"
echo "Building..."
go build -o gosearch . || { echo "build failed"; exit 1; }
echo "✓ built ./gosearch"
echo

run() {
  echo "──────────────────────────────────────────────────────────"
  echo "\$ ./gosearch $*"
  ./gosearch "$@"
  echo
}

run -ext .go "func New" .
run -i "todo" internal
run -r "^func \w+\(" .
run -c "string" .

rm -f gosearch
echo "✓ demo complete"
