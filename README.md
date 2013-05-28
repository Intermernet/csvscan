###csvscan is a utility for extracting fields from CSV files into separate files###

####Usage:####

`csvexporter.exe -in=somefile.csv`

`-frag` : Overwrite files (defaults to "false")
`-i` : Index of field to export (defaults to "1")
`-in` : CSV file to import (Must be supplied!)
`-ni` : Name Index (field to name files by - must be unique and filename safe!, defaults to line number)
`-out` : Output directory (Defaults to the current directory)
`-pre` : File prefix (defaults to nothing)
`-suf` : File suffix and extension (defaults to ".txt")

Usage example: 

####Example usage:####

`csvexporter -in=in.csv -out=outdirectory -pre=Invoice- -suf=-2013.txt -i=2 -ni=1 -frag=true`

####Downloads####
*~ 1.3MB*
- Linux ([64 bit tar.gz][3] | [32 bit tar.gz][4])
- Windows ([64 bit zip][5] | [32 bit zip][6])
- OSX ([64 bit zip][7] | [32 bit zip][8])

NOTE: This program has been tested on Windows 8 64bit, and Ubuntu 13.04

[3]: http://downloads.intermer.net/csvscan/bin/linux_amd64/csvscan_linux_amd64.tar.gz
[4]: http://downloads.intermer.net/csvscan/bin/linux_386/csvscan_linux_386.tar.gz
[5]: http://downloads.intermer.net/csvscan/bin/windows_amd64/csvscan_windows_amd64.zip
[6]: http://downloads.intermer.net/csvscan/bin/windows_386/csvscan_windows_386.zip
[7]: http://downloads.intermer.net/csvscan/bin/darwin_amd64/csvscan_darwin_amd64.zip
[8]: http://downloads.intermer.net/csvscan/bin/darwin_386/csvscan_darwin_386.zip
[8]: http://downloads.intermer.net/csvscan/bin/darwin_386/csvscan_darwin_386.zip