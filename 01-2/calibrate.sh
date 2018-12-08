#!/bin/bash
A=0
FF="freq"
FOUND=0
DC=0
rm -f $FF
echo 0 >> $FF # test case 1 - sepcial case 0
while [ $FOUND -eq 0 ]
do
echo "Loop: $DC"
    while read e
    do
        EXP="$A$e"
        #echo $EXP
        A=$(($EXP))
        #A=$(echo $EXP | bc)
        echo $A >> $FF
        C=$(fgrep -w $A $FF | wc -l)
        if [ $C -eq 2 ]
        then
            echo "Antwort: $A"
            FOUND=1
            break
        fi 
    done < "$1"
DC=$(($DC + 1))
#FOUND=1
done
#rm -f $FF
