#!/bin/bash

read -p "Укажите директорию: " dir

if [ ! -d "$dir" ]; then
    echo "Ошибка: Директория '$dir' не существует"
    exit 1
fi

cd $dir

PREFIX="backup_"
for filename in *; do
  if [ -f "$filename" ]; then
    mv "$filename" "${PREFIX}${filename}"
  fi
done