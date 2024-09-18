#!/bin/sh

RES=0

while IFS= read -r line || [ -n "$line" ]
do
    CONTAINER1=${line:0:$((${#line}/2))}
    CONTAINER2=${line:$((${#line}/2))}

    CAR=$(tr -dc "$CONTAINER2" <<< "$CONTAINER1")
    echo $CAR

    ASCII=$(printf '%d' "'$CAR")

    if [[ "$ASCII" -gt 96 ]]; then
        RES=$((RES + ASCII - 96))
    else
        RES=$((RES + ASCII - 38))
    fi
done

echo "Result: $RES"