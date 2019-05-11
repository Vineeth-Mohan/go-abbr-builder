#!/bin/bash

sort out | uniq -c | sed 's/^[\t ]*\([0-9][0-9]*\) \(.*\)$/\2,\1/' 
#| sort -t , -k 1,1 > shorts.txt