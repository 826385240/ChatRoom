#!/bin/bash

clientnum=1

for i in $(seq 1 $clientnum)
do
    ./mainclient $i &
done
