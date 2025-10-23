#!/bin/bash

TOKEN=
CHAT_ID=

while True; do
    response=$(curl -s "https://api.telegram.org/bot$TOKEN/getUpdates?offset=0")
    update=$(echo "$response" | jq -r '.result[-1]')
    update_id=$(echo "$update" | jq -r '.update_id')
    message_text=$(echo "$update" | jq -r '.message.text')
    chat_id=$(echo "$update" | jq -r '.message.chat.id')

    if [[ "$message_text" == "Дата" ]]; then
        current_date=$(date '+%d.%m.%Y')
        curl -s -X POST "https://api.telegram.org/bot$TOKEN/sendMessage" \
                            -d chat_id="$CHAT_ID" \
                            -d text="$current_date" > /dev/null
        break;
    fi    
    sleep 1
done