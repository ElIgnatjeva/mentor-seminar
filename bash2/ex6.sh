#!/bin/bash

STATE_FILE="$(basename "$1").state"
current_info="$(ls -lh $1 2>/dev/null)"
previous_info=$(cat "$STATE_FILE")

if [ "$current_info" != "$previous_info" ]; then
echo "Файл изменен"
fi

echo "$current_info" > "$STATE_FILE"