# ![CMU Nucleome Browser](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/image/logo_bar.png)
[![Build Status](https://travis-ci.org/nbrowser/cnb.svg?branch=master)](https://travis-ci.org/nbrowser/cnb)
[![Go Report Card](https://goreportcard.com/badge/github.com/nbrowser/cnb)](https://goreportcard.com/report/github.com/nbrowser/cnb)
[![Releases](https://img.shields.io/github/release/nbrowser/cnb.svg)](https://github.com/nbrowser/cnb/releases)
[![Licenses](https://img.shields.io/badge/license-gpl3-orange.svg)](https://opensource.org/licenses/GPL-3.0)

## Introduction

[*CMU Nucleome Browser* Introduction Slides](https://docs.google.com/presentation/d/1KQx4Si6pbCeM2dhqJBkRPuSmwr7tF_CpE-siZtbzqXg/edit?usp=sharing)

[*CMU Nucleome Browser* Website](http://genome.compbio.cs.cmu.edu:8080/static/)

**Figure** CMU Nucleome Browser Design and Concepts

***Figure 1***
![N Browsers](https://genome.compbio.cs.cmu.edu/static/image/pipeline.png)


## Web Application
[![Web Application](https://img.shields.io/badge/CMU-Nucleome--Browser-green.svg?style=for-the-badge)](https://genome.compbio.cs.cmu.edu/v1/pub.html)

## Start a Data Server
### GO binaries for CMU Nucleome Browser Data Server Tools
[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](https://genome.compbio.cs.cmu.edu/static/dist/current/linux/cnbData)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](https://genome.compbio.cs.cmu.edu/static/dist/current/win64/cnbData.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](https://genome.compbio.cs.cmu.edu/static/dist/current/mac/cnbData)

The pre-compiled GO binaries for Mac, Linux, and Windows OS can be downloaded from this [link](https://genome.compbio.cs.cmu.edu/static/dist/current). Please note that there are binaries (32 bit or 64 bit) for two types of Windows machines. An easy way to tell which binaries to download is to check the memory size of your machine. The 32 bit Windows OS cannot support more than 4Gb of RAM/memory.

### install CMU Nucleome Browser Data Server from GitHub
If the binaries does not work on your machine, you can also get the source and binary files from this site via GO. (You have to get GO environment installed first. To get GO installed, you can refer to [this link](https://golang.org/doc/install)).

`go get -u github.com/ma-compbio/cnb/...`

### Start a data server
To start a data server. Use command below in your local machine.

`cnbData start -i [google-sheet-id or excel filename]`

:exclamation: **hint** A google sheet ID can be extracted from its URL. For example, the Google sheet ID in the URL https://docs.google.com/spreadsheets/d/abc1234567/edit#gid=0 is "abc1234567".


:exclamation: **hint** When it is your first time to run the application, it will automatically generate a .cnbData directory in your HOME directory.

### Visualize your data
Then you should able to see your host data in our [web application](https://genome.compbio.cs.cmu.edu)

### Input GSheet or Xls
[![google sheet example](https://img.shields.io/badge/example-gsheet-green.svg?style=flat)](https://docs.google.com/spreadsheets/d/1DEvA94QkN1KZQT51IYOOcIvGL2Ux7Qwqe5IpE9Pe1N8/edit?usp=sharing)

## Powered by
### Our Lib
[![sand](https://img.shields.io/badge/go--js-sand-3d72c6.svg?style=flat)](https://github.com/nbrowser/sand)
[![data](https://img.shields.io/badge/golang-data-3d72c6.svg?style=flat)](https://github.com/nimezhu/data)
[![netio](https://img.shields.io/badge/golang-netio-3d72c6.svg?style=flat)](https://github.com/nimezhu/netio)
[![indexed](https://img.shields.io/badge/golang-indexed-3d72c6.svg?style=flat)](https://github.com/nimezhu/indexed)
### Other Lib
[![d3js](https://img.shields.io/badge/javascript-d3js-yellow.svg?style=flat)](http://d3js.org)
[![golden-layout](https://img.shields.io/badge/javascript-golden--layout-green.svg?style=flat)](http://golden-layout.com)
[![astilectron](https://img.shields.io/badge/golang-astilectron-blue.svg?style=flat)](https://github.com/asticode/go-astilectron)
[![gonetics](https://img.shields.io/badge/golang-gonetics-red.svg?style=flat)](https://github.com/pbenner/gonetics)
[![imageserver](https://img.shields.io/badge/golang-imageserver-blue.svg?style=flat)](https://github.com/pierrre/imageserver)
