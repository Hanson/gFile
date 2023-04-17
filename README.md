File tool for golang

# Install

`go get github.com/hanson/gFile`

# Usage

```
gFile.IsExists(path) // true or false
gFile.CreateDirIfNotExists(path, 0666) // create dir if not exists
gFile.IsDir(path) // return true if path is dir
gFile.IsFile(path) // return true if path is file
```