#!bin/bash

date
go test -run TestDieharderSR__Kiss__WELL512 -timeout 500m
dieharder -a -g 202 -f sr__Kiss__WELL512_6Block_5000000.txt >> ./output/18polls_kiss_well512.txt
date