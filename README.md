Build multiplatform apps with Go, HTML and CSS. Currently this platform only supports MacOS.

MacOS: 10.13.1 (High Sierra)
Latest version of Xcode (mandatory for Apple frameworks). You can install it from Apple store.

1. For running this GUI, please keep the files as the way they are arranged.
  - main folder
    - go files (.go)
    - text file recording scores (.txt). PLEASE DO NOT CHANGE THIS FILE! OR PROGRAM MAY CRASH!
    - resources
      - CSS
        - css files (.css)
      - images (.jpg)

 Caution: Please DO NOT change the score record file, or program may crash! If you accidentally
 changed it, please just delete this file.

2. Get dependent package
  go get -u github.com/murlokswarm/mac

3. In command line, cd to this directory (2048) and type:
  go Build
  ./2048GUI

That's it!
