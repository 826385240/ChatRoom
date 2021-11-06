#!/bin/bash

clientnum=200

for i in $(seq 1 $clientnum)
do
    ./mainclient &
done
