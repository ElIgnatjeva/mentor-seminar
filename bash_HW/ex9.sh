#!/bin/bash

# Мониторинг системы: Напишите скрипт для мониторинга ресурсов системы:
# 	•	Сбор данных о загрузке процессора, использовании памяти и дискового пространства.
# 	•	Уведомление пользователя, если использование памяти превышает 80%, с выводом текущих процессов, 
#     потребляющих наибольшее количество ресурсов.

THRESHOLD=80
disk=$(df / | awk 'NR==2 {print $5}' | sed 's/%//')
echo "Использование дискового пространства: $disk%"
if [ "$disk" -gt "$THRESHOLD" ]; then
    echo "ВНИМАНИЕ! Превышен лимит использования диска ${THRESHOLD}%!"
fi

top_output=$(top -l 1)
cpu=$(echo "$top_output" | grep "CPU usage" | awk -v OFS=' ' '{print $3, $4 $5, $6 $7, $8}')
echo "Загрузка процессора: $cpu"


pagesize=$(vm_stat | grep "page size" | awk '{print $8}')
free_pages=$(vm_stat | grep "Pages free" | awk '{print $3}' | sed 's/\.//')
active_pages=$(vm_stat | grep "Pages active" | awk '{print $3}' | sed 's/\.//')
inactive_pages=$(vm_stat | grep "Pages inactive" | awk '{print $3}' | sed 's/\.//')
wired_pages=$(vm_stat | grep "Pages wired" | awk '{print $4}' | sed 's/\.//')

total_gb=$(( ($free_pages + $active_pages + $inactive_pages + $wired_pages) * $pagesize / 1024 / 1024 / 1024 ))
used_gb=$(( ($active_pages + $wired_pages) * $pagesize / 1024 / 1024 / 1024))
used_ram=$(( ($used_gb * 100) / $total_gb ))


echo "Total RAM: ${total_gb} GB"
echo "Used RAM: ${used_gb} GB"
if [ "$used_ram" -gt "$THRESHOLD" ]; then
    echo "ВНИМАНИЕ! Превышен лимит использования памяти ${THRESHOLD}%!"
    echo "Процессы, потребляющие наибольшее количество ресурсов"
    echo "$(ps aux -m | head -n 11)"
fi