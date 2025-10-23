#!/bin/bash

read -p "Укажите директорию: " dir
tar -czvf $(date +"%Y-%m-%d").tar.gz $dir