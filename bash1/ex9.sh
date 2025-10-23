#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Ошибка ввода. Правильный ввод: $0 команда [аргументы...]"
    exit 1
fi

"$@" &
PID=$!
echo "PID: $PID"