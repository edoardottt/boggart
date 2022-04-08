#!/bin/bash

path=$(pwd)
folder=$(basename $path)

if [[ $folder == "scripts" ]];
then
    echo $folder
    rm -rf ../config.yaml
    touch ../config.yaml
fi

if [[ $folder == "boggart" ]];
then
    echo $folder
    rm -rf config.yaml
    touch config.yaml
fi