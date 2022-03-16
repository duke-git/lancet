# Fileutil
Package fileutil implements some basic functions for file operations.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/fileutil/file.go](https://github.com/duke-git/lancet/blob/main/fileutil/file.go)

<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/fileutil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [ClearFile](#ClearFile)
- [CreateFile](#CreateFile)
- [CopyFile](#CopyFile)
- [FileMode](#FileMode)
- [MiMeType](#MiMeType)
- [IsExist](#IsExist)
- [IsLink](#IsLink)
- [IsDir](#IsDir)
- [ListFileNames](#ListFileNames)
- [RemoveFile](#RemoveFile)
- [ReadFileToString](#ReadFileToString)
- [ReadFileByLine](#ReadFileByLine)
- [Zip](#Zip)

- [UnZip](#UnZip)

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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    isCreatedSucceed := fileutil.CreateFile("./test.txt")
    fmt.Println(isCreatedSucceed)
}
```


### <span id="CopyFile">CopyFile</span>
<p>Copy src file to dest file. If dest file exist will overwrite it.</p>

<b>Signature:</b>

```go
func CopyFile(srcFilePath string, dstFilePath string) error
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    err := fileutil.CopyFile("./test.txt", "./test_copy.txt")
	if err != nil {
		fmt.Println(err)
	}
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
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
    "github.com/duke-git/lancet/fileutil"
)

func main() {
	err := fileutil.Zip("./test.zip", "./unzip/test.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```





