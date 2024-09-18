#!/bin/sh

RES=0

while read -r line1 || [ -n "$line1" ]
do
    read -r line2 || [ -n "$line2" ]
    read -r line3 || [ -n "$line3" ]

    CAR=$(tr -dc "$line2" <<< "$line1")
    CAR=$(tr -dc "$line3" <<< "$CAR")

    ASCII=$(printf '%d' "'$CAR")

    if [[ "$ASCII" -gt 96 ]]; then
        RES=$((RES + ASCII - 96))
    else
        RES=$((RES + ASCII - 38))
    fi
done

echo "Result: $RES"