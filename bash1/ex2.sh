#!/bin/bash

read -p "Введите путь до файла:" file
if [ -f "$file" ]; then
echo "Файл найден!"
else
echo "Файл не найден."
fi