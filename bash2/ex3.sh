#!/bin/bash

echo "$(date +%s%N | sha256sum | head -c$1)"