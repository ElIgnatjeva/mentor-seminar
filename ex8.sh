#!/bin/bash

THRESHOLD=80
usage=$(df / | awk 'NR==2 {print $5}' | sed 's/%//')
echo "Использование диска: $usage%"
if [ "$usage" -gt "$THRESHOLD" ]; then
    echo "ВНИМАНИЕ! Превышен лимит использования диска ${THRESHOLD}%!"
fi