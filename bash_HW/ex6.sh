#!/bin/bash

# Запишет кол-во строк в output1.txt
wc -l input.txt > output1.txt 2> error1.log

# Выведет ошибку в error2.log
wc -l inpu.txt > output2.txt 2> error2.log