#!/bin/bash

path=$(pwd)
folder=$(basename $path)

if [[ $folder == "scripts" ]];
then
    echo $folder
    rm -rf ../config.yaml
    echo -e "# Same as examples/basic-raw.yaml\n$(cat ../examples/basic-raw.yaml)" > ../config.yaml
fi

if [[ $folder == "boggart" ]];
then
    echo $folder
    rm -rf config.yaml
    echo -e "# Same as examples/basic-raw.yaml\n$(cat examples/basic-raw.yaml)" > config.yaml
fi