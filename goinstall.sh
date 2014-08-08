#!/bin/sh

cd ~

wget http://golang.org/dl/go1.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
mkdir $HOME/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
mkdir -p $GOPATH/src/github.com/ansonl
