#!/bin/sh

RES=0

while IFS=, read -r line line2 || [ -n "$line" ]
do
    IFS='-' read -r first1 last1 <<< "$line"
    IFS='-' read -r first2 last2 <<< "$line2"

    if [[ "$first1" -le "$first2" && "$last1" -ge "$last2" ]]; then
        RES=$((RES + 1))
    elif [[ "$first2" -le "$first1" && "$last2" -ge "$last1" ]]; then
        RES=$((RES + 1))
    fi
done

echo "Result: $RES"