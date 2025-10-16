#!/bin/bash

read -p "Укажите директорию: " dir

if [ ! -d "$dir" ]; then
    echo "Ошибка: Директория '$dir' не существует"
    exit 1
fi

echo "Следующие файлы будут удалены:"
find "$dir" -type f -mtime +7 -ls
read -p "Продолжить удаление? (y/N): " confirm
if [[ $confirm == [yY] ]]; then
    find "$dir" -type f -mtime +7 -delete
    echo "Удаление завершено"
else
    echo "Удаление отменено"
fi