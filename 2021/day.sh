#!/bin/bash

# Run this script without args to setup current day, or pass in specific day 1,2 not 01, 02.
# It will:
#   - create a new directory for the day
#   - download the problem input
#   - copy the template file
#

day=$(date +'%d')
if [ -z "$1" ];then
  mkdir $day
else
  day=$1
  mkdir $day
fi

cp main_template.go $day/main.go

curl -o "$day/input.txt" https://adventofcode.com/2021/day/$day/input -H "Cookie: session=$SESSIONID"
