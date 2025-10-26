#!/bin/bash

# Менеджер резервного копирования: Создайте скрипт, который:
# 	•	Создаёт резервную копию всех файлов в указанной директории, добавляя к именам файлов текущую дату.
# 	•	Поддерживает сохранение логов операций в отдельный файл и уведомляет пользователя об успешном завершении
#      с указанием количества файлов.

BACKUP_DIR="new_dir"
DATE=$(date +%Y-%m-%d)
LOG_FILE="${DATE}_backup.log"

create_backup() {
    local count=0
    cd $BACKUP_DIR
    touch $LOG_FILE
    echo "[$(date +%H:%M:%S)] Начало резервного копирования" >> "$LOG_FILE"
    for filename in *; do
    if [ -f "$filename" ] && [ "$filename" != "$LOG_FILE" ]; then
        cp "$filename" "${filename}_${DATE}"
        echo "[$(date +%H:%M:%S)] Скопирован: $filename -> ${filename}_${DATE}" >> "$LOG_FILE"
        ((count++))
    fi
    done
    echo "[$(date +%H:%M:%S)] Всего скопировано: $count файлов" >> "$LOG_FILE"
    return $count
}

notify() {
    local count=$1
    echo "=========================================="
    echo "Скопировано файлов: $count"
    echo "Лог операции сохранен в: $LOG_FILE"
    echo "Директория: $BACKUP_DIR"
    echo "Время создания: $DATE"
    echo "=========================================="
}

main() {
    create_backup
    local file_count=$?
    notify $file_count
}

main
