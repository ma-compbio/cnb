# ![CMU Nucleome Browser](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/image/logo_bar.png)
[![Build Status](https://travis-ci.org/nbrowser/cnb.svg?branch=master)](https://travis-ci.org/nbrowser/cnb)
[![Go Report Card](https://goreportcard.com/badge/github.com/nbrowser/cnb)](https://goreportcard.com/report/github.com/nbrowser/cnb)
[![Releases](https://img.shields.io/github/release/nbrowser/cnb.svg)](https://github.com/nbrowser/cnb/releases)
[![Licenses](https://img.shields.io/badge/license-gpl3-orange.svg)](https://opensource.org/licenses/GPL-3.0)
[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/linux/cnb)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/win64/cnb.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/mac/cnb)




## Introduction
*CMU Nucleome Browser* is an interactive browser with both web-based and desktop applications. It has great extendibility for integrating distributed heterogeneous large-scale datasets related to nucleome.
The browser is built with a multi [window/panel platform](https://github.com/nbrowser/sand) withe highly flexible and powerful communication among different panels. The panels are also modularized, configurable, reusable, and event-driven. 

Currently, it mainly provides three types of panels. 
The first one is Genome Browser panel, which supports [bigwig](http://genome.ucsc.edu/goldenPath/help/bigWig.html), [bigbed](http://genome.ucsc.edu/goldenPath/help/bigBed.html) and [juicebox hic](https://github.com/theaidenlab/juicer/wiki/Data)  format files as tracks. 
The second panel is for 3D structrue, which supports PDB style ascii text file format. Here is [an example file](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/data/structure_3.txt) for 3D structure input.
The third panel is image panel that supports most main stream image formats. 
More types of panel and customized panel interface will be implemented and incorporated in the future.

**Figure** CMU Nucleome Browser Design and Concepts
***Figure 1***
![N Browsers](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/image/pipeline.png)

***Figure 2***
![Nbrowser](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/nbrowser.png)


## Web Application
[![Web Application](https://img.shields.io/badge/CMU-Nucleome--Browser-green.svg?style=for-the-badge)](http://genome.compbio.cs.cmu.edu:8080)
### Quick Start
#### Browse genome
:exclamation: **hint** right click brush region to update the selected coordinates. 

![How to Navigate](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/nav_500px.gif)
#### Select tracks
:exclamation: **hint**  All panel modules are designed as configurable. You can click the config toggle button in the top-right corner, which will toggle between content view and config view.

![How to Select Tracks](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/select_500px.gif)

#### Multi Panels (Example: DNA 3D structure) 
:exclamation: **hint**  you can add panel to window in menu.

![DNA 3D structure](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/3d_500px.gif)

#### Multi Windows (Example: genome coordinates related image panel)
:exclamation: **hint**  The *CMU Nucleome Browser* supports multiple window system. Click *add window* button in the menu will pop up a new window, which would allow you to add more panels. These windows have the capabilities to communicate with each other, which means you can select any region in any window, and all panels will respond accordingly.

![FISH data](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/ext_500px.gif)

#### User Sessions Manager 
:exclamation: **hint** create a new Google sheet, and copy google sheet ID to CNB Web Application. You can save and load your sessions easily. You can also share your sessions with your colleagues by sharing this Google sheet with them.

![Session & Google Sheet](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/gifs/session_500px.gif)

## Install Web Server/Desktop Application/Desktop Client
*CMU Nucleome Browser* is implemented in Go and JavaScript. It can be installed in OSX, Windows, and Linux machines.

#### GO binaries for CMU Nucleome Browser
[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/linux/cnb)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/win64/cnb.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current/mac/cnb)

The pre-compiled GO binaries for Mac, Linux, and Windows OS can be downloaded from this [link](http://genome.compbio.cs.cmu.edu/~xiaopenz/cnb/current). Please note that there are binaries (32 bit or 64 bit) for two types of Windows machines. An easy way to tell which binaries to download is to check the memory size of your machine. The 32 bit Windows OS cannot support more than 4Gb of RAM/memory.

#### install CMU Nucleome Browser from GitHub
If the binaries does not work on your machine, you can also get the source and binary files from this site via GO. (You have to get GO environment installed first. To get GO installed, you can refer to [this link](https://golang.org/doc/install)).

`go get -u github.com/nbrowser/cnb`

## Run Application

### Quick Start
#### Desktop Application/Client
To start a desktop application. Use command below.

`cnb app -i [google-sheet-id or excel filename] -c creds.json`

:exclamation: **hint** A google sheet ID can be extracted from its URL. For example, the Google sheet ID in the URL https://docs.google.com/spreadsheets/d/abc1234567/edit#gid=0 is "abc1234567".

:exclamation: **hint** a creds.json file is needed for user session management and administration. To test you can use [this example](https://github.com/nbrowser/cnb/blob/master/creds.json). The cid and csecret only work on local 127.0.0.1:8080-8082. If you need to host a web server, you need to apply a new [oauth2 service](https://developers.google.com/identity/protocols/OAuth2) from Google.

:exclamation: **hint** When it is your first time to run the application, it will automatically generate a .cnb directory in your HOME directory, it might take a while to download electron into your .cnb directory. 
#### Web Server Application 
if you just want to start a web server instead of desktop application , you can add option "-m web". 

`cnb app -i [google-sheet-id or excel filename] -m web -c creds.json`

Then open "http://127.0.0.1:8080" in your web browser.
#### Input GSheet or Xls
[![google sheet example](https://img.shields.io/badge/example-gsheet-green.svg?style=flat)](https://docs.google.com/spreadsheets/d/1WaChnccn5iyKHm1ccZhRHSZmpzW7pPPrpGTSGoKTKO4/edit?usp=sharing)

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
