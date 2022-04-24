#!/bin/bash

path=$(pwd)
folder=$(basename $path)

if [[ $folder == "scripts" ]];
then
    rm -rf ../public/honeypot/icons/
    rm -rf ../public/honeypot/apache-404.html
    rm -rf ../public/honeypot/index.html
    ./clean-config.sh
    echo "Done!"
fi

if [[ $folder == "boggart" ]];
then
    rm -rf public/honeypot/icons/
    rm -rf public/honeypot/apache-404.html
    rm -rf public/honeypot/index.html
    ./scripts/clean-config.sh
    echo "Done!"
fi