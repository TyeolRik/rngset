#!/bin/bash

#go test -run TestDieharderSR__Keccak256__WELL512a -timeout 800m
#
#for i in $(seq 2 1 10)
#do
#    fileName=sr__Keccak256__WELL512a_${i}Block_2Participant_4Polls_100000000
#
#    time_start="$(date -u +%s)"
#    dieharder -a -g 202 -f ./generated/${fileName}.dat >> ./output/${fileName}.txt
#    time_end="$(date -u +%s)"
#    elapsed=$(($time_end-$time_start))
#    echo "$elapsed seconds costs for Dieharder -a -g 202" >> ./output/${fileName}.txt
#    echo "Test ends in ${elapsed} seconds!" | mail -s "Test${i} End" kino6147@gmail.com -A ./output/${fileName}.txt
#done

go test -run TestDieharderSR__Keccak256__WELL512a__Binaries -timeout 999m

# For Dieharder test
fileName=sr__Keccak256__WELL512a_6Block_2Participant_4Polls_2000000000
time_start="$(date -u +%s)"
dieharder -a -g 201 -f ./generated/${fileName}.bin >> ./output/${fileName}.txt
time_end="$(date -u +%s)"
elapsed=$(($time_end-$time_start))
echo "$elapsed seconds costs for Dieharder -a -g 201" >> ./output/${fileName}.txt
