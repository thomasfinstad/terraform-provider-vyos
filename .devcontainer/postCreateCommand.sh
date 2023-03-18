#!/bin/bash -ex

###
# Apt packages
export DEBIAN_FRONTEND=noninteractive
sudo apt-get update

# Used to parse ISO download website
sudo apt-get -y install libxml2-utils

# User to format ISO download website
wget -O tidy.deb "https://github.com/htacg/tidy-html5/releases/download/5.9.14-next/tidy-5.9.14-Linux-64bit.deb"
sudo apt-get install "./tidy.deb"
rm -v "./tidy.deb"

rm -rf /var/lib/apt/lists/*

###
# Pre-commit / git tools

# Used for commit message formatting
pipx install argcomplete
pipx install Commitizen
register-python-argcomplete cz >> ~/.bashrc

# Used to improve commits before they are commited
pipx install pre-commit
pre-commit install --hook-type pre-commit --hook-type commit-msg --hook-type pre-push
