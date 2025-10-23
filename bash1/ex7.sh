#!/bin/bash

add() {
    local num1=$1
    local num2=$2
    local sum=$((num1 + num2))
    echo "$num1 + $num2 = $sum"
}
if [ $# -ne 2 ]; then
    echo "Ошибка ввода. Правильный ввод: $0 число1 число2"
    exit 1
fi
add "$1" "$2"