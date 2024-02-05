#!/bin/bash

set -e

run_set() {
    case $1 in
        1)
            A=1
            ;;
        2)
            A=9
            ;;
        3)
            A=17
            ;;
        4)
            A=25
            ;;
        *)
            A=0
    esac
            
    for ((i = $A; i < $A + 8; i++)); do
        challenge="src/ch$(printf "%02d" "$i").py"
        if [ -e $challenge ]; then 
            python $challenge
        fi
    done
}

run_set $1
