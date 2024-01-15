#!/bin/bash

set -e

run_set() {
    case $1 in
        1)
            A=1
            B=9
            ;;
        2)
            A=9
            B=17
            ;;
        3)
            A=17
            B=25
            ;;
        *)
            A=0
            B=0
    esac
            
    for ((i = $A; i < $B; i++)); do
        challenge="src/ch$(printf "%02d" "$i").py"
        if [ -e $challenge ]; then 
            python $challenge
        fi
    done
}

run_set $1
