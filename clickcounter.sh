#!/bin/bash

count=0

function redraw {
    local str width height length
    
    width=$(tput cols)
    height=$(tput lines)
    str="$count"
    length=${#str}
    clear
    tput cup $((height / 2)) $(((width / 2) - (length / 2)))
    echo "$str"
}

function cleanup {
    tput cnorm
    echo "Clean up"
    rm temp_file
    clear
}

trap cleanup EXIT
trap redraw WINCH

tput civis
tput setab 0
tput setaf 3
tput bold
redraw

temp_file=$(mktemp)
xinput --test-xi2 --root 12 | grep --line-buffered 'RawButtonPress' >> temp_file & 

while true
do
    new_count=$(wc -l temp_file | awk '{print $1}')
    if [ $new_count != $count ]; then
        redraw
        count=$new_count
    fi
done
