# CMU Nucleome Browser
[![Build Status](https://travis-ci.org/nbrowser/cnb.svg?branch=master)](https://travis-ci.org/nbrowser/cnb)
[![Go Report Card](https://goreportcard.com/badge/github.com/nbrowser/cnb)](https://goreportcard.com/report/github.com/nbrowser/cnb)
[![Releases](https://img.shields.io/github/release/nbrowser/cnb.svg)](https://github.com/nbrowser/cnb/releases)
[![Downloads](https://img.shields.io/github/downloads/nbrowser/cnb/latest/total.svg)]()
[![Licenses](https://img.shields.io/badge/license-gpl3-orange.svg)](https://opensource.org/licenses/GPL-3.0)
[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/linux/cnb)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/win64/cnb.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/mac/cnb)

## Introduction
*CMU Nucleome Browser* is an interactively browser web/desktop mixture application/platform with great extendibility for integrating distributed heterogeneous nucleome related data.
It is built with a communicatable multi [window/panel platform](https://github.com/nbrowser/sand) and  configurable , composable, reusable, event-driven panel modules for browsing different type or chromosomes coordinate related data. 
Currently , It mainly provides 3 types of panels. 
The first one is Genome Browser panel , which supports [bigwig](http://genome.ucsc.edu/goldenPath/help/bigWig.html), [bigbed](http://genome.ucsc.edu/goldenPath/help/bigBed.html) and [juicebox hic](https://github.com/theaidenlab/juicer/wiki/Data)  format files as tracks. 
In 3d structrue panel, it supports PDB style ascii text file format. Here's [an example file](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/data/structure_3.txt) for 3d structure input.
In image panel, it support most main stream image format. 
More types of panel and customized panel interface will be added in the future.




**Figure** CMU Nucleome Browser Design and Concepts
![Nbrowser](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/nbrowser.png)





## Web Application
[![Web Application](https://img.shields.io/badge/CMU-Nucleome--Browser-green.svg?style=for-the-badge)](http://genome.compbio.cs.cmu.edu:8080)
### Quick Start
#### Browse genome
:exclamation:  **Hint** right click brush region to update to selected coordinates. 
![How to Navigate](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/nav_500px.gif)
#### Select tracks
![How to Select Tracks](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/select_500px.gif)

#### Multi Panels (Example: DNA 3D structure) 
![DNA 3D structure](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/3d_500px.gif)

#### Multi Windows (Example: genome coordinates related image panel)
The *CMU Nucleome Browser* supports extend window system.

![FISH data](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/ext_500px.gif)

#### User Sessions Manager 
![Session & Google Sheet](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/session_500px.gif)

## Install Web Server/Desktop Application/Desktop Client
*CMU Nucleome Browser* is implemented in Go and JavaScript. It can be installed in OSX, windows, linux machines.

#### GO binaries for CMU Nucleome Browser
[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/linux/cnb)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/win64/cnb.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/mac/cnb)

The pre-compiled GO binaries for mac, linux and windows OS can be downloaded from this [link](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current). Please note that there're binaries (32 bit or 64 bit) for two type of windows machines. A quick way to tell which binaries to download is to check the memory size of your machine. The 32 bit windows OS can not support more than 4Gb of RAM/memory.

#### install CMU Nucleome Browser from GitHub
If the binaries does not work on your machine, you can also get the source and binary files from this site via GO. (You have to get GO environment installed beforehead. To get GO installed, you can refer to [this link](https://golang.org/doc/install)).

`go get -u github.com/nbrowser/cnb`

## Run Application

### Quick Start
#### Desktop Application/Client
To start a desktop application. Use command below.

`cnb app -i [google-sheet-id or excel filename]`

When first time to run the application, it will automatically generate a .cnb directory in your HOME directory, it might take a while to download electron into your .cnb directory. 
#### Web Server Application 
if you just want to start a web server instead of desktop application , you can add option "-m web". 

`cnb app -i [google-sheet-id or excel filename] -m web`

Then open "http://127.0.0.1:8080" in your web browser.
#### Input GSheet or Xls
[![google sheet example](https://img.shields.io/badge/example-gsheet-green.svg?style=flat)](https://docs.google.com/spreadsheets/d/1WaChnccn5iyKHm1ccZhRHSZmpzW7pPPrpGTSGoKTKO4/edit?usp=sharing)


## Powered by
[![d3js](https://img.shields.io/badge/javascript-d3js-yellow.svg?style=flat)](http://d3js.org)
[![golden-layout](https://img.shields.io/badge/javascript-golden--layout-green.svg?style=flat)](http://golden-layout.com)
[![astilectron](https://img.shields.io/badge/golang-astilectron-blue.svg?style=flat)](https://github.com/asticode/go-astilectron)
[![gonetics](https://img.shields.io/badge/golang-gonetics-red.svg?style=flat)](https://github.com/pbenner/gonetics)
[![imageserver](https://img.shields.io/badge/golang-imageserver-blue.svg?style=flat)](https://github.com/pierrre/imageserver)
