#!/bin/bash

for f in $1*; do
  if [ -f "$f" ]; then
    mv "$f" "$(echo "$f" | tr 'A-Z' 'a-z')"
  fi
done