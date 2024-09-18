#!/bin/sh


SCORE=0

while IFS=' ' read -r letter1 letter2 || [ -n "$letter1" ]
do

    if [[ "$letter2" == "X" ]]; then
        case $letter1 in
            A) SCORE=$((SCORE + 3));;
            B) SCORE=$((SCORE + 1));;
            C) SCORE=$((SCORE + 2));;
        esac

    elif [[ "$letter2" == "Y" ]]; then
        case $letter1 in
            A) SCORE=$((SCORE + 4));;
            B) SCORE=$((SCORE + 5));;
            C) SCORE=$((SCORE + 6));;
        esac

    elif [[ "$letter2" == "Z" ]]; then
        case $letter1 in
            A) SCORE=$((SCORE + 8));;
            B) SCORE=$((SCORE + 9));;
            C) SCORE=$((SCORE + 7));;
        esac
    fi

done

echo $SCORE