#!/bin/bash

echo "bash: cp"
date +"%T.%N" 
cp -r $1 ../temp
date +"%T.%N" 
sudo rm -rf ../go

echo "go: ./cp"
date +"%T.%N" 
./cp -r $1 ../temp
date +"%T.%N" 
sudo rm -rf ../temp