#!/bin/bash

# Напишите скрипт, который автоматически сортирует файлы в указанной директории:
# 	•	Файлы с расширениями .jpg, .png и .gif перемещает в папку Images.
# 	•	Документы (.txt, .pdf, .docx) — в папку Documents.
# 	•	Скрипт должен запускаться автоматически через cron каждую ночь и вести лог всех перемещений.

LOG_FILE="$(date +%Y_%m_%d)_move.log"

cd sort_dir

echo "[$(date +%H:%M:%S)] Начало сортировки" >> "$LOG_FILE"

find . \( -name "*.docx" -o -name "*.txt" -o -name "*.pdf" \) -exec mv {} Documents \; &> /dev/null
echo "[$(date +%H:%M:%S)] Документы перенесены в Documents" >> "$LOG_FILE"

find . \( -name "*.png" -o -name "*.jpg" -o -name "*.gif" \) -exec mv {} Images \; &> /dev/null
echo "[$(date +%H:%M:%S)] Изображения перенесены в Images" >> "$LOG_FILE"

echo "[$(date +%H:%M:%S)] Сортировка завершена" >> "$LOG_FILE"

# Задаем автоматический запуск через cron каждую полночь
# 0 0 * * * ex10.sh