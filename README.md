# BEXfinder

Browsers Extensions Finder (BEXfinder) is a portable and cross-platform (Windows, Linux and MacOS) command-line tool to find out all web browsers (Google Chrome, Microsoft Edge, Brave Browser, Mozilla FireFox, Opera, etc.) extensions installed on system.

Currently Supported Browsers:
1. Google Chrome
2. Microsoft Edge
3. Brave Browser
4. Vivaldi
5. Mozilla FireFox
6. Opera Browser

More are coming...

![Main](https://i.ibb.co/fDHNx8F/main.png "Main")

> Please feel free to contribute to this project. If you have an idea or improvement issue a pull request!

#### Disclaimer
This tool is a PoC (Proof of Concept) and does not guarantee results.
This tool is only for academic purposes and testing  under controlled environments.
The author bears no responsibility for any misuse of the tool.


## To Run All
Example: On Linux, to retrive all browsers extensions found on the system, run:
```
./bexfinder all
```
![Example 1](https://i.ibb.co/wLjc2Mt/screenshot.png "Example 1")

## Checking Browser Extensions/Plugins Online
Example: On Linux system, to check if the extensions/plugins from the official websites:
```
./bexfinder chrome --check
```
![Example 2](https://i.ibb.co/8dFS0pT/check-online.png "Example 2")


### Development
To install the dependencies used by BEXfinder and run it:
```
go install
go run main.go
```

## Contact
ch4meleon@protonmail-dot-com


