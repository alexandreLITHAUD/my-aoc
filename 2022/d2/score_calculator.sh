#!/bin/bash

SCORE=0

while IFS=' ' read -r letter1 letter2
do
    # echo "First letter: $letter1, Second letter: $letter2"

    if [[ "$letter2" == "X" ]]; then
        SCORE=$((SCORE + 1))

        if [[ "$letter1" == "A" ]]; then
            SCORE=$((SCORE + 3))
        fi

        if [[ "$letter1" == "B" ]]; then
            SCORE=$((SCORE + 0))
        fi

        if [[ "$letter1" == "C" ]]; then
            SCORE=$((SCORE + 6))
        fi

    fi

    if [[ "$letter2" == "Y" ]]; then
        SCORE=$((SCORE + 2))

        if [[ "$letter1" == "A" ]]; then
            SCORE=$((SCORE + 6))
        fi

        if [[ "$letter1" == "B" ]]; then
            SCORE=$((SCORE + 3))
        fi 

        if [[ "$letter1" == "C" ]]; then
            SCORE=$((SCORE + 0))
        fi
    fi

    if [[ "$letter2" == "Z" ]]; then
        SCORE=$((SCORE + 3))

        if [[ "$letter1" == "A" ]]; then
            SCORE=$((SCORE + 0))
        fi 

        if [[ "$letter1" == "B" ]]; then
            SCORE=$((SCORE + 6))
        fi

        if [[ "$letter1" == "C" ]]; then
            SCORE=$((SCORE + 3))
        fi
    fi
done

echo $SCORE