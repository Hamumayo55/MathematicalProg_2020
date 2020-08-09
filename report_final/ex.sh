#!/bin/zsh
count=2100
for i in `seq 11`
do
count=`expr $count - 100`
python3 glpk.py ${count} >> f_x4.txt
done

