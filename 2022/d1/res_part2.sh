#!/bin/sh

RES=0
MAX1=0
MAX2=0
MAX3=0

while read -r line; 
do
    if [[ -z "$line" ]]; then  # Check if line is empty
        if [[ "$RES" -gt "$MAX1" ]]; then
            MAX3=$MAX2
            MAX2=$MAX1
            MAX1=$RES  # Update MAX if RES is greater
        elif [[ "$RES" -gt "$MAX2" ]]; then
            MAX3=$MAX2
            MAX2=$RES  # Update MAX if RES is greater
        elif [[ "$RES" -gt "$MAX3" ]]; then
            MAX3=$RES  # Update MAX if RES is greater
        fi
        RES=0  # Reset RES
    else
        RES=$((RES + line))  # Add the line value to RES (assuming numeric)
    fi
done < input.txt

echo $((MAX1 + MAX2 + MAX3))