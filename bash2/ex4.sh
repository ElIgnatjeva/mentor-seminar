#!/bin/bash

sleep 30

if ping -c 5 8.8.8.8 &> /dev/null; then
    echo "Гугл доступен"
else
    echo "Гугл недоступен"
fi