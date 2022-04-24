#!/bin/bash

path=$(pwd)
folder=$(basename $path)

if [[ $folder == "scripts" ]];
then
    cp ../examples/apache-httpd/apache.yaml ../config.yaml
    cp -r ../examples/apache-httpd/icons ../public/honeypot
    cp ../examples/apache-httpd/apache-404.html ../public/honeypot
    cp ../examples/apache-httpd/index.html ../public/honeypot
    echo "Done!"
fi

if [[ $folder == "boggart" ]];
then
    cp examples/apache-httpd/apache.yaml config.yaml
    cp -r examples/apache-httpd/icons public/honeypot
    cp examples/apache-httpd/apache-404.html public/honeypot
    cp examples/apache-httpd/index.html public/honeypot
    echo "Done!"
fi