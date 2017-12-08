# CMU Nucleome Browser
[![Go Report Card](https://goreportcard.com/badge/github.com/nbrowser/cnb)](https://goreportcard.com/report/github.com/nbrowser/cnb)
[![Licenses](https://img.shields.io/badge/license-gpl3-orange.svg)](https://    opensource.org/licenses/GPL-3.0)

## Introduction
*CMU Nucleome Browser* is an extendable genome browser desktop/web application for integrating all kinds of nucleome related data distributed on the internet or in your local machine.

Now it supports bigwig, bigbed and hic format files and as demonstration it has a simple panel module for 3d DNA structure.



## Demo
[Demo Website](http://genome.compbio.cs.cmu.edu:8080)

## Install
*CMU Nucleome Browser* is implemented in Go and JavaScript. It can be installed in OSX, windows, linux machines.

#### GO binaries for CMU Nucleome Browser
The pre-compiled GO binaries for mac, linux and windows OS can be downloaded from this [link](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current). Please note that there're binaries (32 bit or 64 bit) for two type of windows machines. A quick way to tell which binaries to download is to check the memory size of your machine. The 32 bit windows OS can not support more than 4Gb of RAM/memory.

#### install CMU Nucleome Browser from GitHub
If the binaries does not work on your machine, you can also get the source and binary files from this site via GO. (You have to get GO environment installed beforehead. To get GO installed, you can refer to [this link](https://golang.org/doc/install)).

`go get -u github.com/nbrowser/cnb`

## Run Application

`cnb app -i [google-sheet-id or excel filename]`
