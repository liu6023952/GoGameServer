#!/bin/sh 

sh stop.sh

sleep 1
./bin/connector -e local -s 1 &
sleep 1
./bin/connector -e local -s 2 &

sleep 1
./bin/game -e local -s 1 &
sleep 1
./bin/game -e local -s 2
