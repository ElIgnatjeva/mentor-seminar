#!/bin/bash

TOKEN=
CHAT_ID=

curl -s -X POST "https://api.telegram.org/bot$TOKEN/sendMessage" \
     -d chat_id="$CHAT_ID" \
     -d text="Файл изменён!"