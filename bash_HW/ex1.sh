#!/bin/bash

	# •	Задание: Напишите Bash-скрипт, который выполняет следующие действия:
        # •	Создаёт список всех файлов в текущей директории, указывая их тип (файл, каталог и т.д.).
        # •	Проверяет наличие определённого файла, переданного как аргумент скрипта, и выводит сообщение о его наличии или отсутствии.
        # •	Использует цикл for для вывода информации о каждом файле: его имя и права доступа.

command=$1

if [ "$command" -eq '1' ]; then
    for filename in *; do
        if [ -f "$filename" ]; then
            echo "$filename - файл"
        elif [ -d "$filename" ]; then
            echo "$filename - директория"
        fi
    done
elif [ "$command" -eq '2' ]; then
    file="$(ls | grep $2)"
    if [ "$file" ]; then
        echo "Файл существует"
    else 
        echo "Файл не найден"
    fi
elif [ "$command" -eq '3' ]; then
    for filename in *; do
        if [ -f "$filename" ]; then
            echo "$filename $(ls -l $filename | cut -f 1 -d ' ')"
        fi
    done
fi
