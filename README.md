# Introduction

This project is a POC to sync up a local access (windows) collection of tables to mysql DB on a secured remote server (through SSH)

The program is built from golang and can be built from linux (cross compiling)

# Getting started

From a linux box:

## Requirements

- docker 1.8 or higher

## Steps

- Initialize your build environment:

   ```bash
   mkdir ~/src/go/src
   cd ~/src/go/src
   git clone https://github.com/clarsonneur/matos_sync
   # If you are running docker without sudo
   source build-env.sh ~/src/go
   # If you are running docker with sudo
   source build-env.sh --sudo ~/src/go
   build.sh
   ``̀
