# UkraineWarRussiansData
Contains script to search around some data

# How to run
Build:
```
go build -o searcher main.go 
```

Or use prebuilt binaries (e. g. searcher.exe from repo)

# How to use
Run from directory with xlsx files
```
searcher abc foo
```
It will find all rows, which contain`s cells with substrings `abc` and `foo`. You can path any number of substrings:

```
searcher abc
searcher abc foo
searcher abc foo bar
```

# How to get xlsx files
Find them in tg, or open Issue with ur contacts, or write me up (if you can)
