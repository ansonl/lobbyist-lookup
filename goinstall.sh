#!/bin/sh


#setup Go lang
cd ~
wget http://golang.org/dl/go1.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.3.linux-amd64.tar.gz
mkdir $HOME/go

#set envvars
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PORT=80

mkdir -p $GOPATH/src/github.com/ansonl

#install git, hg, htop
sudo apt-get install git mercurial htop

#clone lobbyist-lookup
cd go/src/github.com/ansonl
git clone https://github.com/ansonl/lobbyist-lookup

#build and run
cd lobbyist-lookup
go get

sudo setcap 'cap_net_bind_service=+ep' /home/$USER/go/bin/lobbyist-lookup

nohup lobbyist-lookup &
