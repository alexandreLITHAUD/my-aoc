#!/bin/bash

SCORE=0

while IFS=' ' read -r letter1 letter2 || [ -n "$letter1" ]
do

    if [[ "$letter2" == "X" ]]; then
        SCORE=$((SCORE + 1))
        case $letter1 in
            A) SCORE=$((SCORE + 3)) ;;
            B) SCORE=$((SCORE + 0)) ;;
            C) SCORE=$((SCORE + 6)) ;;
        esac
    elif [[ "$letter2" == "Y" ]]; then
        SCORE=$((SCORE + 2))
        case $letter1 in
            A) SCORE=$((SCORE + 6)) ;;
            B) SCORE=$((SCORE + 3)) ;;
            C) SCORE=$((SCORE + 0)) ;;
        esac
    elif [[ "$letter2" == "Z" ]]; then
        SCORE=$((SCORE + 3))
        case $letter1 in
            A) SCORE=$((SCORE + 0)) ;;
            B) SCORE=$((SCORE + 6)) ;;
            C) SCORE=$((SCORE + 3)) ;;
        esac
    fi

    # echo "Current score: $SCORE"

done

echo $SCORE