#!/bin/bash

sh ex2.sh 2 3 &
pid1=$!
sh ex4.sh &
pid2=$!
echo "Запущенные команды: $pid1, $pid2"
wait $pid1 $pid2
echo "Команды завершены"