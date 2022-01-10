<div align="center">
    <h1 style="width: 100%; text-align: center;">Lancet</h1>
    <p style="font-size: 18px">
        Lancet is a comprehensive, efficient, and reusable util function library of go. Inspired by the java apache common package and lodash.js.
    </p>
<div align="center" style="text-align: center;">

![Go version](https://img.shields.io/badge/go-%3E%3D1.16<recommend>-9cf)
[![Release](https://img.shields.io/badge/release-1.1.8-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com//duke-git/lancet?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet)
[![Go Report Card](https://goreportcard.com/badge/github.com/duke-git/lancet)](https://goreportcard.com/report/github.com/duke-git/lancet)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/main/LICENSE)

</div>

English | [简体中文](./README_zh-CN.md)

</div>

### Feature

- 👏 Comprehensive, efficient and reusable.
- 💪 140+ common go util functions, support string, slice, datetime, net, crypt...
- 💅 Only depend on the go standard library.
- 🌍 Unit test for every exported function.

### Installation

```go
go get github.com/duke-git/lancet
```

### Usage
Lancet organizes the code into package structure, and you need to import the corresponding package name when use it. For example, if you use string-related functions,import the strutil package like below:

```go
import "github.com/duke-git/lancet/strutil"
```

### Example

Here takes the string function ReverseStr (reverse order string) as an example, and the strutil package needs to be imported.

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    s := "hello"
    rs := strutil.ReverseStr(s)
    fmt.Println(rs) //olleh
}
```

### API Documentation

#### 1. convertor contains some functions for data convertion

- Support conversion between commonly used data types.
- Usage: import "github.com/duke-git/lancet/cryptor"

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    s := "12.3"
    f, err := convertor.ToFloat(s)
    if err != nil {
        fmt.Errorf("error is %s", err.Error())
    }
    fmt.Println(f) // 12.3
}
```

- Function list:

```go
func ColorHexToRGB(colorHex string) (red, green, blue int) //convert color hex to color rgb
func ColorRGBToHex(red, green, blue int) string //convert color rgb to color hex
func ToBool(s string) (bool, error)  //convert string to a boolean
func ToBytes(data interface{}) ([]byte, error) //convert interface to bytes
func ToChar(s string) []string //convert string to char slice
func ToFloat(value interface{}) (float64, error) //convert value to float64, if input is not a float return 0.0 and error
func ToInt(value interface{}) (int64, error) //convert value to int64, if input is not a numeric format return 0 and error
func ToJson(value interface{}) (string, error) //convert value to a valid json string
func ToString(value interface{}) string //convert value to string
func StructToMap(value interface{}) (map[string]interface{}, error) //convert struct to map, only convert exported field, tag `json` should be set
```

#### 2. cryptor is for data encryption and decryption

- Support md5, hmac, aes, des, ras.
- Usage: import "github.com/duke-git/lancet/cryptor"

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))
    fmt.Println(string(decrypted)) // hello
}
```

- Function list:

```go
func AesEcbEncrypt(data, key []byte) []byte //AES ECB encrypt
func AesEcbDecrypt(encrypted, key []byte) []byte //AES ECB decrypt
func AesCbcEncrypt(data, key []byte) []byte //AES CBC encrypt
func AesCbcDecrypt(encrypted, key []byte) []byte //AES CBC decrypt
func AesCtrCrypt(data, key []byte) []byte //AES CTR encrypt / decrypt
func AesCfbEncrypt(data, key []byte) []byte //AES CFB encrypt
func AesCfbDecrypt(encrypted, key []byte) []byte //AES CFB decrypt
func AesOfbEncrypt(data, key []byte) []byte //AES OFB encrypt
func AesOfbDecrypt(data, key []byte) []byte //AES OFB decrypt
func Base64StdEncode(s string) string //base64 encode
func Base64StdDecode(s string) string //base64 decode
func DesCbcEncrypt(data, key []byte) []byte //DES CBC encrypt
func DesCbcDecrypt(encrypted, key []byte) []byte //DES CBC decrypt
func DesCtrCrypt(data, key []byte) []byte //DES CTR encrypt/decrypt
func DesCfbEncrypt(data, key []byte) []byte //DES CFB encrypt
func DesCfbDecrypt(encrypted, key []byte) []byte //DES CFB decrypt
func DesOfbEncrypt(data, key []byte) []byte //DES OFB encrypt
func DesOfbDecrypt(data, key []byte) []byte //DES OFB decrypt
func HmacMd5(data, key string) string //get hmac md5 value
func HmacSha1(data, key string) string //get hmac sha1 value
func HmacSha256(data, key string) string //get hmac sha256 value
func HmacSha512(data, key string) string //get hmac sha512 value
func Md5String(s string) string //return the md5 value of string
func Md5File(filename string) (string, error) //return the md5 value of file
func Sha1(data string) string //get sha1 value
func Sha256(data string) string //getsha256 value
func Sha512(data string) string //get sha512 value
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) //generate RSA pem file
func RsaEncrypt(data []byte, pubKeyFileName string) []byte //RSA encrypt
func RsaDecrypt(data []byte, privateKeyFileName string) []byte //RSA decrypt

```

#### 3. datetime parse and format datetime

- Parse and format datetime
- Usage: import "github.com/duke-git/lancet/datetime"

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    s := datetime.FormatTimeToStr(now, "yyyy-mm-dd hh:mm:ss")
    fmt.Println(s) // 2021-11-24 11:16:55
}
```

- Function list:

```go
func AddDay(t time.Time, day int64) time.Time //add or sub days to time
func AddHour(t time.Time, hour int64) time.Time //add or sub hours to time
func AddMinute(t time.Time, minute int64) time.Time //add or sub minutes to time
func GetNowDate() string  //get current date, format is yyyy-mm-dd
func GetNowTime() string //get current time, format is hh:mm:ss
func GetNowDateTime() string //get current date and time, format is yyyy-mm-dd hh:mm:ss
func GetZeroHourTimestamp() int64 //return timestamp of zero hour (timestamp of 00:00)
func GetNightTimestamp() int64 //return timestamp of zero hour (timestamp of 23:59)
func FormatTimeToStr(t time.Time, format string) string //convert time to string
func FormatStrToTime(str, format string) time.Time //convert string to time
```

#### 4. fileutil basic functions for file operations

- Basic functions for file operations.
- Usage: import "github.com/duke-git/lancet/fileutil"

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/fileutil"
)

func main() {
    fmt.Println(fileutil.IsDir("./")) // true
}
```

- Function list：

```go
func ClearFile(path string) error //write empty string to path file
func CreateFile(path string) bool // create a file in path
func CopyFile(srcFilePath string, dstFilePath string) error //copy src file to dst file
func FileMode(path string) (fs.FileMode, error) //return file's mode and permission
func MiMeType(file interface{}) string //return file mime type, file should be string or *os.File
func IsExist(path string) bool  //checks if a file or directory exists
func IsLink(path string) bool //checks if a file is symbol link or not
func IsDir(path string) bool //checks if the path is directy or not
func ListFileNames(path string) ([]string, error) //return all file names in the path
func RemoveFile(path string) error //remove the path file
func ReadFileToString(path string) (string, error) //return string of file content
func ReadFileByLine(path string)([]string, error) //read file content by line
func Zip(fpath string, destPath string) error //create zip file, fpath could be a single file or a directory
func UnZip(zipFile string, destPath string) error //unzip the file and save it to destPath
```

#### 5. formatter is for data format

- Contain some formatting function
- Usage: import "github.com/duke-git/lancet/formatter"

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
     fmt.Println(formatter.Comma("12345", ""))   // "12,345"
     fmt.Println(formatter.Comma(12345.67, "¥")) // "¥12,345.67"
}
```

- Function list:

```go
func Comma(v interface{}, symbol string) string  //add comma to number by every 3 numbers from right. ahead by symbol char
```

#### 6. function can control the function execution and support functional programming

- Control function execution and support functional programming.
- Usage: import "github.com/duke-git/lancet/function"

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
    var print = func(s string) {
		fmt.Println(s)
	}
	function.Delay(2*time.Second, print, "hello world")
}
```

- Function list:

```go
func After(n int, fn interface{}) func(args ...interface{}) []reflect.Value  //creates a function that invokes func once it's called n or more times
func Before(n int, fn interface{}) func(args ...interface{}) []reflect.Value  //creates a function that invokes func once it's called less than n times
func (f Fn) Curry(i interface{}) func(...interface{}) interface{}  //make a curryed function
func Compose(fnList ...func(...interface{}) interface{}) func(...interface{}) interface{}  //compose the functions from right to left
func Delay(delay time.Duration, fn interface{}, args ...interface{})  //invoke function after delayed time
func Schedule(d time.Duration, fn interface{}, args ...interface{}) chan bool //invoke function every duration time, util close the returned bool chan
func (w *Watcher) Start() //start the watch timer.
func (w *Watcher) Stop() //stop the watch timer
func (w *Watcher) Reset() {} //reset the watch timer.
func (w *Watcher) GetElapsedTime() time.Duration //return time duration from watcher start to end.
```

#### 7. netutil is for net process

- Ip and http request method.
- Usage: import "github.com/duke-git/lancet/netutil".
- The Http function params order：url, header, query string, body, httpclient.

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/netutil"
)

func main() {
    url := "https://gutendex.com/books?"
    header := make(map[string]string)
    header["Content-Type"] = "application/json"
    queryParams := make(map[string]interface{})
    queryParams["ids"] = "1"

    resp, err := netutil.HttpGet(url, header, queryParams)
    if err != nil {
       log.Fatal(err)
    }

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response: ", resp.StatusCode, string(body))
}
```

- Function list:

```go
func GetInternalIp() string //get internal ip
func GetPublicIpInfo() (*PublicIpInfo, error) //get public ip info: country, region, isp, city, lat, lon, ip
func IsPublicIP(IP net.IP) bool //判断ip是否为公共ip
func HttpGet(url string, params ...interface{}) (*http.Response, error) //http get request
func HttpPost(url string, params ...interface{}) (*http.Response, error) //http post request
func HttpPut(url string, params ...interface{}) (*http.Response, error) //http put request
func HttpDelete(url string, params ...interface{}) (*http.Response, error) //http delete request
func HttpPatch(url string, params ...interface{}) (*http.Response, error) //http patch request
func ConvertMapToQueryString(param map[string]interface{}) string //convert map to url query string
func ParseHttpResponse(resp *http.Response, obj interface{}) error //decode http response to specified interface
```

#### 8. random is for rand string and int generation

- Generate random string and int.
- Usage: import "github.com/duke-git/lancet/random".

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/random"
)

func main() {
    randStr := random.RandString(6)
    fmt.Println(randStr)
}
```

- Function list:

```go
func RandBytes(length int) []byte //generate random []byte
func RandInt(min, max int) int //generate random int
func RandString(length int) string //generate random string
```

#### 9. slice is for process slice

- Contain function for process slice.
- Usage: import "github.com/duke-git/lancet/slice"
- Due to the unstable support of generic, most of the slice processing function parameter and return value is interface {}. After go generic is stable, the related functions will be refactored.

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/slice"
)

func main() {
    nums := []int{1, 4, 3, 4, 6, 7, 3}
    uniqueNums, _ := slice.IntSlice(slice.Unique(nums))
    fmt.Println(uniqueNums) //[1 4 3 6 7]
}
```

- Function list:

```go
func Contain[T comparable](slice []T, value T) bool //check if the value is in the slice or not
func Chunk[T any](slice []T, size int) [][]T //creates an slice of elements split into groups the length of size.
func ConvertSlice(originalSlice interface{}, newSliceType reflect.Type) interface{} //convert originalSlice to newSliceType
func Difference[T comparable](slice1, slice2 []T) []T //creates an slice of whose element not included in the other given slice
func DeleteByIndex(slice interface{}, start int, end ...int) (interface{}, error) //delete the element of slice from start index to end index - 1
func Drop(slice interface{}, n int) interface{} //creates a slice with `n` elements dropped from the beginning when n > 0, or `n` elements dropped from the ending when n < 0
func Every[T any](slice []T, fn func(index int, t T) bool) bool  //return true if all of the values in the slice pass the predicate function
func None[T any](slice []T, fn func(index int, t T) bool) bool // return true if all the values in the slice mismatch the criteria
func Filter [T any] (slice []T, fn func(index int, t T) bool) []T //filter slice, fn signature should be func(int, T) bool.
func Find[T any](slice []T, fn func(index int, t T) bool) (*T, bool) //iterates over elements of slice, returning the first one that passes a truth test on function.function signature should be func(index int, value interface{}) bool .
func FlattenDeep(slice interface{}) interface{} //flattens slice recursive
func ForEach [T any] (slice []T, fn func(index int, t T)) //iterates over elements of slice and invokes function for each element, fn signature should be func(int, T ).
func IntSlice(slice interface{}) ([]int, error) //convert value to int slice
func InterfaceSlice(slice interface{}) []interface{} //convert value to interface{} slice
func Intersection(slices ...interface{}) interface{} //creates a slice of unique values that included by all slices.
func InsertByIndex(slice interface{}, index int, value interface{}) (interface{}, error) //insert the element into slice at index.
func Map [T any, U any] (slice []T, fn func(index int, t T) U) []U //map lisce, fn signature should be func(int, T).
func Reverse[T any](slice []T)//revere slice
func Reduce[T any](slice []T, fn func(index int, t1, t2 T) T, initial T) T //reduce slice
func Shuffle[T any](slice []T) []T  //creates an slice of shuffled values
func SortByField(slice interface{}, field string, sortType ...string) error //sort struct slice by field
func Some[T any](slice []T, fn func(index int, t T) bool) bool //return true if any of the values in the list pass the predicate fn function
func StringSlice(slice interface{}) []string //convert value to string slice
func Unique(slice interface{}) interface{} //remove duplicate elements in slice
func Union(slices ...interface{}) interface{} //Union creates a slice of unique values, in order, from all given slices. using == for equality comparisons.
func UpdateByIndex(slice interface{}, index int, value interface{}) (interface{}, error) //update the slice element at index.
func Without(slice interface{}, values ...interface{}) interface{} //creates a slice excluding all given values
func GroupBy(slice, function interface{}) (interface{}, interface{}) // groups slice into two categories
func Count[T any](slice []T, fn func(index int, t T) bool) int // Count iterates over elements of slice, returns a count of all matched elements
```

#### 10. strutil is for processing string

- Contain functions to precess string
- Usage: import "github.com/duke-git/lancet/strutil"

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := "Foo-Bar"
    camelCaseStr := strutil.CamelCase(str)
    fmt.Println(camelCaseStr) //fooBar
}
```

- Function list:

```go
func After(s, char string) string //create substring in source string after position when char first appear
func AfterLast(s, char string) string //create substring in source string after position when char last appear
func Before(s, char string) string //create substring in source string before position when char first appear
func BeforeLast(s, char string) string //create substring in source string before position when char last appear
func CamelCase(s string) string //covert string to camelCase string. "foo bar" -> "fooBar"
func Capitalize(s string) string //convert the first character of a string to upper case, "fOO" -> "Foo"
func IsString(v interface{}) bool //check if the value data type is string or not
func KebabCase(s string) string //covert string to kebab-case, "foo_Bar" -> "foo-bar"
func LowerFirst(s string) string //convert the first character of string to lower case
func PadEnd(source string, size int, padStr string) string //pads string on the right side if it's shorter than size
func PadStart(source string, size int, padStr string) string//pads string on the left side if it's shorter than size
func ReverseStr(s string) string //return string whose char order is reversed to the given string
func SnakeCase(s string) string //covert string to snake_case "fooBar" -> "foo_bar"
func Wrap(str string, wrapWith string) string //wrap a string with another string.
func Unwrap(str string, wrapToken string) string //unwrap a given string from anther string. will change str value
```

#### 11. validator is for data validation

- Contain function for data validation.
- Usage: import "github.com/duke-git/lancet/validator".

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/validator"
)

func main() {
    str := "Foo-Bar"
    isAlpha := validator.IsAlpha(str)
    fmt.Println(isAlpha) //false
}
```

- Function list:

```go
func ContainChinese(s string) bool //check if the string contain mandarin chinese
func IsAlpha(s string) bool //checks if the string contains only letters (a-zA-Z)
func IsBase64(base64 string) bool //check if the string is base64 string
func IsChineseMobile(mobileNum string) bool //check if the string is chinese mobile number
func IsChineseIdNum(id string) bool //check if the string is chinese id number
func IsChinesePhone(phone string) bool //check if the string is chinese phone number
func IsCreditCard(creditCart string) bool //check if the string is credit card
func IsDns(dns string) bool //check if the string is dns
func IsEmail(email string) bool //check if the string is a email address
func IsEmptyString(s string) bool //check if the string is empty
func IsFloatStr(s string) bool //check if the string can convert to a float
func IsNumberStr(s string) bool //check if the string can convert to a number
func IsRegexMatch(s, regex string) bool //check if the string match the regexp
func IsIntStr(s string) bool //check if the string can convert to a integer
func IsIp(ipstr string) bool //check if the string is a ip address
func IsIpV4(ipstr string) bool //check if the string is a ipv4 address
func IsIpV6(ipstr string) bool //check if the string is a ipv6 address
func IsStrongPassword(password string, length int) bool //check if the string is strong password (alpha(lower+upper) + number + special chars(!@#$%^&*()?><))
func IsWeakPassword(password string) bool //check if the string is weak password（only letter or only number or letter + number）
```
