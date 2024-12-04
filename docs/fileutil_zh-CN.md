# Fileutil

fileutil 包支持文件基本操作。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/fileutil/file.go](https://github.com/duke-git/lancet/blob/v1/fileutil/file.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/fileutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [ClearFile](#ClearFile)
-   [CreateFile](#CreateFile)
-   [CreateDir](#CreateDir)
-   [CopyFile](#CopyFile)
-   [FileMode](#FileMode)
-   [MiMeType](#MiMeType)
-   [IsExist](#IsExist)
-   [IsLink](#IsLink)
-   [IsDir](#IsDir)画
-   [ListFileNames](#ListFileNames)
-   [RemoveFile](#RemoveFile)
-   [ReadFileToString](#ReadFileToString)
-   [ReadFileByLine](#ReadFileByLine)
-   [Zip](#Zip)
-   [UnZip](#UnZip)
-   [ZipAppendEntry](#ZipAppendEntry)
-   [CurrentPath](#CurrentPath)
-   [IsZipFile](#IsZipFile)
-   [FileSize](#FileSize)
-   [MTime](#MTime)
-   [Sha](#Sha)
-   [ReadCsvFile](#ReadCsvFile)
-   [WriteCsvFile](#WriteCsvFile)
-   [WriteMapsToCsv](#WriteMapsToCsv)
-   [WriteStringToFile](#WriteStringToFile)
-   [WriteBytesToFile](#WriteBytesToFile)
-   [ReadFile](#ReadFile)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="ClearFile">ClearFile</span>

<p>清空文件内容</p>

<b>函数签名:</b>

```go
func ClearFile(path string) error
```

<b>例子:</b>

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

### <span id="CreateDir">CreateDir</span>

<p>使用绝对路径创建嵌套目录，例如/a/, /a/b/</p>

<b>函数签名:</b>

```go
func CreateDir(absPath string) error
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    err := fileutil.CreateDir("/a/")
    fmt.Println(err)
}
```

### <span id="CreateFile">CreateFile</span>

<p>创建文件，创建成功返回true, 否则返回false</p>

<b>函数签名:</b>

```go
func CreateFile(path string) bool
```

<b>例子:</b>

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

<p>拷贝文件，会覆盖原有的拷贝文件</p>

<b>函数签名:</b>

```go
func CopyFile(srcFilePath string, dstFilePath string) error
```

<b>例子:</b>

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

<p>获取文件mode信息</p>

<b>函数签名:</b>

```go
func FileMode(path string) (fs.FileMode, error)
```

<b>例子:</b>

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

<p>获取文件mime类型, 'file'参数的类型必须是string或者*os.File</p>

<b>函数签名:</b>

```go
func MiMeType(file interface{}) string
```

<b>例子:</b>

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

<p>判断文件或目录是否存在</p>

<b>函数签名:</b>

```go
func IsExist(path string) bool
```

<b>例子:</b>

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

<p>判断文件是否是符号链接</p>

<b>函数签名:</b>

```go
func IsLink(path string) bool
```

<b>例子:</b>

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

<p>判断目录是否存在</p>

<b>函数签名:</b>

```go
func IsDir(path string) bool
```

<b>例子:</b>

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

<p>返回目录下所有文件名</p>

<b>函数签名:</b>

```go
func ListFileNames(path string) ([]string, error)
```

<b>例子:</b>

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

<p>删除文件</p>

<b>函数签名:</b>

```go
func RemoveFile(path string) error
```

<b>例子:</b>

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

<p>读取文件内容并返回字符串</p>

<b>函数签名:</b>

```go
func ReadFileToString(path string) (string, error)
```

<b>例子:</b>

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

<p>按行读取文件内容，返回字符串切片包含每一行</p>

<b>函数签名:</b>

```go
func ReadFileByLine(path string)([]string, error)
```

<b>例子:</b>

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

<p>zip压缩文件, fpath参数可以是文件或目录</p>

<b>函数签名:</b>

```go
func Zip(fpath string, destPath string) error
```

<b>例子:</b>

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

<p>zip解压缩文件并保存在目录中</p>

<b>函数签名:</b>

```go
func UnZip(zipFile string, destPath string) error
```

<b>例子:</b>

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

### <span id="ZipAppendEntry">ZipAppendEntry</span>

<p>通过将单个文件或目录追加到现有的zip文件</p>

<b>函数签名:</b>

```go
func ZipAppendEntry(fpath string, destPath string) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    err := fileutil.ZipAppendEntry("./test.txt", "./test.zip")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="CurrentPath">CurrentPath</span>

<p>返回当前位置的绝对路径。</p>

<b>函数签名:</b>

```go
func CurrentPath() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    absPath := CurrentPath()
    fmt.Println(absPath)
}
```

### <span id="IsZipFile">IsZipFile</span>

<p>判断文件是否是zip压缩文件。</p>

<b>函数签名:</b>

```go
func IsZipFile(filepath string) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    isZip := IsZipFile("./zipfile.zip")
    fmt.Println(isZip)
}
```

### <span id="FileSize">FileSize</span>

<p>返回文件字节大小。</p>

<b>函数签名:</b>

```go
func FileSize(path string) (int64, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
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

<p>返回文件修改时间(unix timestamp).</p>

<b>函数签名:</b>

```go
func MTime(filepath string) (int64, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
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

<p>返回文件sha值，参数`shaType` 应传值为: 1, 256，512.</p>

<b>函数签名:</b>

```go
func Sha(filepath string, shaType ...int) (string, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
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

<p>读取csv文件内容到切片</p>

<b>函数签名:</b>

```go
func ReadCsvFile(filepath string) ([][]string, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
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

### <span id="WriteCsvFile">WriteCsvFile</span>

<p>向csv文件写入内容。</p>

<b>函数签名:</b>

```go
func WriteCsvFile(filepath string, records [][]string, append bool) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    data := [][]string{
        {"Lili", "22", "female"},
        {"Jim", "21", "male"},
    }
    err := WriteCsvFile("./testdata/test2.csv", data, false)
    fmt.Println(err)

    content, _ := ReadCsvFile("./testdata/test2.csv")
    fmt.Println(content)

    // Output:
    // <nil>
    // [[Lili 22 female] [Jim 21 male]]
}
```

### <span id="WriteMapsToCsv">WriteMapsToCsv</span>

<p>将map切片写入csv文件中。</p>

<b>函数签名:</b>

```go
// filepath: CSV文件路径。
// records: 写入文件的map切片。map值必须为基本类型。会以map键的字母顺序写入。
// appendToExistingFile: 是否为追加写模式。
// delimiter: CSV文件分割符。
// headers: CSV文件表头顺序（需要与map key保持一致)，不指定时按字母排序。
func WriteMapsToCsv(filepath string, records []map[string]interface{}, appendToExistingFile bool, delimiter rune, headers ...[]string) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    fpath := "./test.csv"
    fileutil.CreateFile(fpath)

    f, _ := os.OpenFile(fpath, os.O_WRONLY|os.O_TRUNC, 0777)
    defer f.Close()

    records := []map[string]interface{}{
        {"Name": "Lili", "Age": "22", "Gender": "female"},
        {"Name": "Jim", "Age": "21", "Gender": "male"},
    }

    headers := []string{"Name", "Age", "Gender"}
    err := fileutil.WriteMapsToCsv(csvFilePath, records, false, ';', headers)

    if err != nil {
        log.Fatal(err)
    }

    content, err := fileutil.ReadCsvFile(csvFilePath, ';')

    fmt.Println(content)

    // Output:
    // [[Name Age Gender] [Lili 22 female] [Jim 21 male]]
}
```

### <span id="WriteBytesToFile">WriteBytesToFile</span>

<p>将bytes写入文件。</p>

<b>函数签名:</b>

```go
func WriteBytesToFile(filepath string, content []byte) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    filepath := "./bytes.txt"

    file, err := os.Create(filepath)
    if err != nil {
        return
    }

    defer file.Close()

    err = fileutil.WriteBytesToFile(filepath, []byte("hello"))
    if err != nil {
        return
    }

    content, err := fileutil.ReadFileToString(filepath)
    if err != nil {
        return
    }

    os.Remove(filepath)

    fmt.Println(content)

    // Output:
    // hello
}
```

### <span id="WriteStringToFile">WriteStringToFile</span>

<p>将字符串写入文件。</p>

<b>函数签名:</b>

```go
func WriteStringToFile(filepath string, content string, append bool) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    filepath := "./test.txt"

    file, err := os.Create(filepath)
    if err != nil {
        return
    }

    defer file.Close()

    err = fileutil.WriteStringToFile(filepath, "hello", true)
    if err != nil {
        return
    }

    content, err := fileutil.ReadFileToString(filepath)
    if err != nil {
        return
    }

    os.Remove(filepath)

    fmt.Println(content)

    // Output:
    // hello
}
```

### <span id="ReadFile">ReadFile</span>

<p>读取文件或者URL</p>

<b>函数签名:</b>

```go
func ReadFile(path string) (reader io.ReadCloser, closeFn func(), err error)
```

<b>示例:<span style="float:right;display:inline-block;"></span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
   reader, fn, err := ReadFile("https://httpbin.org/robots.txt")
    if err != nil {
        return
    }
    defer fn()

    dat, err := io.ReadAll(reader)
    if err != nil {
        return
    }
    fmt.Println(string(dat))
    // Output:
    // User-agent: *
    // Disallow: /deny
}
```
