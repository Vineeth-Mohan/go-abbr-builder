#!/bin/bash

mv nohup.out out
echo "short form,expansion,occurance" > abbrevations.csv
sort out | uniq -c | sed 's/^[\t ]*\([0-9][0-9]*\) \(.*\)$/\2,\1/' > unique-abbr.csv
sort -t, -k 1,1 -k 3,3nr  unique-abbr.csv >> abbrevations.csv
rm unique-abbr.csv
