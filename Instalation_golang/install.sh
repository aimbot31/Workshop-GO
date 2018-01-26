#!/bin/bash

sudo wget https://dl.google.com/go/go1.9.2.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

echo "all done - try the command go"
