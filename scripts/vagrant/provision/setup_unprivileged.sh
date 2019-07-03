#!/bin/bash 

# Use with Ubuntu 16.x+
set -xe

# Install Go
eval "$(${HOME}/provision/gimme.sh stable)"

# Setup GOPATH and binary path
echo '' >> ${HOME}/.bashrc
echo '# GOPATH' >> ${HOME}/.bashrc
echo 'export GOPATH=${HOME}/go' >> ${HOME}/.bashrc
echo '' >> ${HOME}/.bashrc
echo '# GOPATH bin' >> ${HOME}/.bashrc
echo 'export PATH=$PATH:${HOME}/go/bin' >> ${HOME}/.bashrc

# Install Kind
GO111MODULE="on" go get sigs.k8s.io/kind@v0.4.0
