#!/bin/bash

start_time="$(date -u +%s)"
sleep 2
end_time="$(date -u +%s)"
elapsed=$(($end_time-$start_time))
echo $elapsed

start_time="$(date -u +%s)"
sleep 3
end_time="$(date -u +%s)"
elapsed=$(($end_time-$start_time))
echo $elapsed