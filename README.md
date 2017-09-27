# Introduction

This project contains code to migrate time to time the existing windows/MS access DB managing our diving materials to an internet web service.

The internet service will at least provide the same functionality and will be evolve with much more features on needs.

As we are going to migrate from windows MS Access to a remote web service, progressively, we need to keep functional, the existing DB and enhance it with a connectivity to prepare the transition.

For that reason, we are writing a golang windows binary from linux to get access to Access DB, under Windows.

Details are in [github project](https://github.com/clarsonneur/matos_api/projects]

# Getting started

Here are details to build the current golang program from a linux box:

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
    ```

## Windows machine requirements

On the windows machine, you need to install the appropriate `AccessDatabaseEngine` package aligned with your MS Access installation.

The binary is a 32 bits, so, you need to install the 32bits library as well.

For MS Access 2007, I installed the following:
https://www.microsoft.com/en-us/download/confirmation.aspx?id=13255

