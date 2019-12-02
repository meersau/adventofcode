#!/bin/sh
A=0
while read e
do
    EXP="$A$e"
    A=$(($EXP))
done <$1
echo $A
