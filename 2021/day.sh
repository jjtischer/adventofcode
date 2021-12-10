#!/bin/bash

day=$(date +'%d')
if [ -z "$1" ];then
  mkdir $day
else
  day=$1
  mkdir $day
fi

cp main_template.go $day/main.go

curl -o "$day/input.txt" https://adventofcode.com/2021/day/$day/input -H "Cookie: session=$SESSIONID"
