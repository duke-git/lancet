# Fileutil

Package fileutil implements some basic functions for file operations.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/fileutil/file.go](https://github.com/duke-git/lancet/blob/main/fileutil/file.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/fileutil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [ClearFile](#ClearFile)
-   [CreateFile](#CreateFile)
-   [CreateDir](#CreateDir)
-   [CopyFile](#CopyFile)
-   [CurrentPath](#CurrentPath)
-   [FileMode](#FileMode)
-   [MiMeType](#MiMeType)
-   [IsExist](#IsExist)
-   [IsLink](#IsLink)
-   [IsDir](#IsDir)
-   [ListFileNames](#ListFileNames)
-   [RemoveFile](#RemoveFile)
-   [ReadFileToString](#ReadFileToString)
-   [ReadFileByLine](#ReadFileByLine)
-   [Zip](#Zip)
-   [UnZip](#UnZip)
-   [IsZipFile](#IsZipFile)
-   [FileSize](#FileSize)
-   [MTime](#MTime)
-   [Sha](#Sha)
-   [ReadCsvFile](#ReadCsvFile)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="ClearFile">ClearFile</span>

<p>Clear the file content, write empty string to the file.</p>

<b>Signature:</b>

```go
func ClearFile(path string) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    err := fileutil.ClearFile("./test.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="CreateFile">CreateFile</span>

<p>Create file in path. return true if create succeed.</p>

<b>Signature:</b>

```go
func CreateFile(path string) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    isCreatedSucceed := fileutil.CreateFile("./test.txt")
    fmt.Println(isCreatedSucceed)
}
```

### <span id="CreateDir">CreateDir</span>

<p>Create directory in absolute path. param `absPath` like /a/, /a/b/.</p>

<b>Signature:</b>

```go
func CreateDir(absPath string) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    err := fileutil.CreateDir("/a/")
    fmt.Println(err)
}
```

### <span id="CopyFile">CopyFile</span>

<p>Copy src file to dest file. If dest file exist will overwrite it.</p>

<b>Signature:</b>

```go
func CopyFile(srcPath string, dstPath string) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    err := fileutil.CopyFile("./test.txt", "./test_copy.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="CurrentPath">CurrentPath</span>

<p>return current absolute path.</p>

<b>Signature:</b>

```go
func CurrentPath() string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    absPath := CurrentPath()
    fmt.Println(absPath)
}
```

### <span id="FileMode">FileMode</span>

<p>Return file mode infomation.</p>

<b>Signature:</b>

```go
func FileMode(path string) (fs.FileMode, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    mode, err := fileutil.FileMode("./test.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(mode)
}
```

### <span id="MiMeType">MiMeType</span>

<p>Get file mime type, 'file' param's type should be string or *os.File.</p>

<b>Signature:</b>

```go
func MiMeType(file any) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "os"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    type1 := fileutil.MiMeType("./test.txt")
    fmt.Println(type1) //text/plain; charset=utf-8

    f, _ := os.Open("./file.go")
    type2 := fileutil.MiMeType(f)
    fmt.Println(type2) //text/plain; charset=utf-8
}
```

### <span id="IsExist">IsExist</span>

<p>Checks if a file or directory exists.</p>

<b>Signature:</b>

```go
func IsExist(path string) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    fileutil.CreateFile("./test.txt")
    isFileExist := fileutil.IsExist("./test.txt")
    fmt.Println(isFileExist) //true
}
```

### <span id="IsLink">IsLink</span>

<p>Checks if a file is symbol link or not.</p>

<b>Signature:</b>

```go
func IsLink(path string) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    isLinkFile := fileutil.IsLink("./test.txt")
    fmt.Println(isLinkFile) //false
}
```

### <span id="IsDir">IsDir</span>

<p>Checks if the path is directy or not.</p>

<b>Signature:</b>

```go
func IsDir(path string) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    isDir := fileutil.IsDir("./")
    fmt.Println(isDir) //true

    isDir = fileutil.IsDir("./test.txt")
    fmt.Println(isDir) //false
}
```

### <span id="ListFileNames">ListFileNames</span>

<p>List all file names in given path.</p>

<b>Signature:</b>

```go
func ListFileNames(path string) ([]string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    fileNames, _ := fileutil.ListFileNames("./")
    fmt.Println(fileNames)
}
```

### <span id="RemoveFile">RemoveFile</span>

<p>Remove the file of path.</p>

<b>Signature:</b>

```go
func RemoveFile(path string) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    err := fileutil.RemoveFile("./test.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="ReadFileToString">ReadFileToString</span>

<p>Return string of file content.</p>

<b>Signature:</b>

```go
func ReadFileToString(path string) (string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "os"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    path := "./test.txt"
    fileutil.CreateFile(path)

    f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
    f.WriteString("hello world")

    content, _ := fileutil.ReadFileToString(path)
    fmt.Println(content) //hello world
}
```

### <span id="ReadFileByLine">ReadFileByLine</span>

<p>Read file line by line, and return slice of lines</p>

<b>Signature:</b>

```go
func ReadFileByLine(path string)([]string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "os"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    path := "./text.txt"
    fileutil.CreateFile(path)

    f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
    defer f.Close()
    f.WriteString("hello\nworld")

    contents, _ := fileutil.ReadFileByLine(path)
    fmt.Println(contents) //[]string{"hello", "world"}
}
```

### <span id="Zip">Zip</span>

<p>Create a zip file of fpath, fpath could be a file or a directory.</p>

<b>Signature:</b>

```go
func Zip(fpath string, destPath string) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    err := fileutil.Zip("./test.txt", "./test.zip")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="UnZip">UnZip</span>

<p>Unzip the file and save it to dest path.</p>

<b>Signature:</b>

```go
func UnZip(zipFile string, destPath string) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    err := fileutil.Zip("./test.zip", "./unzip/test.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="IsZipFile">IsZipFile</span>

<p>Checks if file is zip file or not.</p>

<b>Signature:</b>

```go
func IsZipFile(filepath string) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    isZip := fileutil.IsZipFile("./zipfile.zip")
    fmt.Println(isZip)
}
```

### <span id="FileSize">FileSize</span>

<p>Returns file size in bytes.</p>

<b>Signature:</b>

```go
func FileSize(path string) (int64, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    size, err := fileutil.FileSize("./testdata/test.txt")

    fmt.Println(size)
    fmt.Println(err)

    // Output:
    // 20
    // <nil>
}
```

### <span id="MTime">MTime</span>

<p>Returns file modified time(unix timestamp).</p>

<b>Signature:</b>

```go
func MTime(filepath string) (int64, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    mtime, err := fileutil.MTime("./testdata/test.txt")

    fmt.Println(mtime)
    fmt.Println(err)

    // Output:
    // 1682391110
    // <nil>
}
```

### <span id="Sha">Sha</span>

<p>returns file sha value, param `shaType` should be 1, 256 or 512.</p>

<b>Signature:</b>

```go
func Sha(filepath string, shaType ...int) (string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    sha1, err := fileutil.Sha("./testdata/test.txt", 1)
    sha256, _ := fileutil.Sha("./testdata/test.txt", 256)
    sha512, _ := fileutil.Sha("./testdata/test.txt", 512)

    fmt.Println(sha1)
    fmt.Println(sha256)
    fmt.Println(sha512)
    fmt.Println(err)

    // Output:
    // dda3cf10c5a6ff6c6659a497bf7261b287af2bc7
    // aa6d0a3fbc3442c228d606da09e0c1dc98c69a1cac3da1909199e0266171df35
    // d22aba2a1b7a2e2f512756255cc1c3708905646920cb1eb95e45b531ba74774dbbb89baebf1f716220eb9cf4908f1cfc5b2a01267704d9a59f59d77cab609870
    // <nil>
}
```

### <span id="ReadCsvFile">ReadCsvFile</span>

<p>Reads file content into slice.</p>

<b>Signature:</b>

```go
func ReadCsvFile(filepath string) ([][]string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/fileutil"
)

func main() {
    content, err := fileutil.ReadCsvFile("./testdata/test.csv")

    fmt.Println(content)
    fmt.Println(err)

    // Output:
    // [[Bob  12  male] [Duke  14  male] [Lucy  16  female]]
    // <nil>
}
```
