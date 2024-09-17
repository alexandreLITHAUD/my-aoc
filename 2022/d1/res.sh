#!/bin/sh

RES=0
MAX=0

while read -r line; 
do
    if [[ -z "$line" ]]; then  # Check if line is empty
        if [[ "$RES" -gt "$MAX" ]]; then
            MAX=$RES  # Update MAX if RES is greater
        fi
        RES=0  # Reset RES
    else
        RES=$((RES + line))  # Add the line value to RES (assuming numeric)
    fi
done < input.txt

if [[ "$RES" -gt "$MAX" ]]; then
    MAX=$RES  # Update MAX if RES is greater
fi

echo $MAX