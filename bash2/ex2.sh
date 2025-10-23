#!/bin/bash

sleep 20

if [ "$1" -gt "$2" ]; then
    echo "$1 больше $2"
elif [ "$1" -lt "$2" ]; then
    echo "$2 больше $1"
else
    echo "Числа равны"
fi