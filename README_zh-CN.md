# Lancet
<p style="font-size: 18px">
    lancet（柳叶刀）是一个全面、高效、可复用的go语言工具函数库。 lancet受到了java apache common包和lodash.js的启发。
</p>

![Go version](https://img.shields.io/badge/go-%3E%3D1.16<recommend>-9cf)
[![Release](https://img.shields.io/badge/release-1.2.1-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com//duke-git/lancet?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet)
[![Go Report Card](https://goreportcard.com/badge/github.com/duke-git/lancet)](https://goreportcard.com/report/github.com/duke-git/lancet)
[![test](https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/duke-git/lancet/actions/workflows/codecov.yml)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/main/LICENSE)


简体中文 | [English](./README.md)


## 特性

- 👏 全面、高效、可复用
- 💪 160+常用go工具函数，支持string、slice、datetime、net、crypt...
- 💅 只依赖go标准库
- 🌍 所有导出函数单元测试覆盖率100%

## 安装

```go
go get github.com/duke-git/lancet
```

## 用法

lancet是以包的结构组织代码的，使用时需要导入相应的包名。例如：如果使用字符串相关函数，需要导入strutil包:

```go
import "github.com/duke-git/lancet/strutil"
```

## 例子

此处以字符串工具函数ReverseStr（逆序字符串）为例，需要导入strutil包:

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

## API文档
### 1. convertor数据转换包

- 转换函数支持常用数据类型之间的转换
- 导入包：import "github.com/duke-git/lancet/convertor"

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

- 函数列表：

```go
func ColorHexToRGB(colorHex string) (red, green, blue int) //颜色值16进制转rgb
func ColorRGBToHex(red, green, blue int) string //颜色值rgb转16进制
func ToBool(s string) (bool, error)  //字符串转成Bool
func ToBytes(data interface{}) ([]byte, error) //interface转成byte slice
func ToChar(s string) []string //字符串转成字符slice
func ToFloat(value interface{}) (float64, error) //interface转成float64
func ToInt(value interface{}) (int64, error) //interface转成int64
func ToJson(value interface{}) (string, error) //interface转成json string
func ToString(value interface{}) string //interface转成string
func StructToMap(value interface{}) (map[string]interface{}, error) //struct串转成map, 需要设置struct tag `json`
```

### 2. cryptor加解密包

- 加密函数支持md5, hmac, aes, des, ras
- 导入包：import "github.com/duke-git/lancet/cryptor"

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

- 函数列表：

```go
func AesEcbEncrypt(data, key []byte) []byte //AES ECB模式加密
func AesEcbDecrypt(encrypted, key []byte) []byte //AES ECB模式解密
func AesCbcEncrypt(data, key []byte) []byte //AES CBC模式加密
func AesCbcDecrypt(encrypted, key []byte) []byte //AES CBC模式解密
func AesCtrCrypt(data, key []byte) []byte //AES CTR模式加密/解密
func AesCfbEncrypt(data, key []byte) []byte //AES CFB模式加密
func AesCfbDecrypt(encrypted, key []byte) []byte //AES CFB模式解密
func AesOfbEncrypt(data, key []byte) []byte //AES OFB模式加密
func AesOfbDecrypt(data, key []byte) []byte //AES OFB模式解密
func Base64StdEncode(s string) string //base64编码
func Base64StdDecode(s string) string //base64解码
func DesCbcEncrypt(data, key []byte) []byte //DES CBC模式加密
func DesCbcDecrypt(encrypted, key []byte) []byte //DES CBC模式解密
func DesCtrCrypt(data, key []byte) []byte //DES CTR模式加密/解密
func DesCfbEncrypt(data, key []byte) []byte //DES CFB模式加密
func DesCfbDecrypt(encrypted, key []byte) []byte //DES CFB模式解密
func DesOfbEncrypt(data, key []byte) []byte //DES OFB模式加密
func DesOfbDecrypt(data, key []byte) []byte //DES OFB模式解密
func HmacMd5(data, key string) string //获取hmac md5值
func HmacSha1(data, key string) string //获取hmac sha1值
func HmacSha256(data, key string) string //获取hmac sha256值
func HmacSha512(data, key string) string //获取hmac sha512值
func Md5String(s string) string //获取字符串md5值
func Md5File(filename string) (string, error) //获取文件md5值
func Sha1(data string) string //获取sha1值
func Sha256(data string) string //获取sha256值
func Sha512(data string) string //获取sha512值
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) //生成RSA私钥文件
func RsaEncrypt(data []byte, pubKeyFileName string) []byte //RSA加密
func RsaDecrypt(data []byte, privateKeyFileName string) []byte //RSA解密

```

### 3. datetime日期时间处理包

- 处理日期时间
- 导入包：import "github.com/duke-git/lancet/datetime"

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

- 函数列表：

```go
func AddDay(t time.Time, day int64) time.Time //加减天数
func AddHour(t time.Time, hour int64) time.Time //加减小时数
func AddMinute(t time.Time, minute int64) time.Time //加减分钟数
func GetNowDate() string  //获取当天日期 格式yyyy-mm-dd
func GetNowTime() string //获取当前时间 格式hh:mm:ss
func GetNowDateTime() string //获取当前日期时间 格式yyyy-mm-dd hh:mm:ss
func GetZeroHourTimestamp() int64 //获取当天零时时间戳（00:00)
func GetNightTimestamp() int64 //获取当天23时时间戳（23:59)
func FormatTimeToStr(t time.Time, format string) string //时间格式化字符串
func FormatStrToTime(str, format string) time.Time //字符串转换成时间
```

### 4. fileutil文件处理包

- 文件处理常用函数
- 导入包：import "github.com/duke-git/lancet/fileutil"

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

- 函数列表：

```go
func ClearFile(path string) error //清空文件内容
func CreateFile(path string) bool //创建文件
func FileMode(path string) (fs.FileMode, error) //返回文件mode信息
func MiMeType(file interface{}) string //返回文件mime类型
func IsExist(path string) bool  //判断文件/目录是否存在
func IsDir(path string) bool //判断是否为目录
func IsLink(path string) bool //检查文件是否为符号链接文件
func RemoveFile(path string) error //删除文件
func CopyFile(srcFilePath string, dstFilePath string) error //复制文件
func ListFileNames(path string) ([]string, error) //列出目录下所有文件名称
func ReadFileToString(path string) (string, error) //读取文件内容为字符串
func ReadFileByLine(path string)([]string, error) //按行读取文件内容
func Zip(fpath string, destPath string) error //压缩文件fpath参数可以是文件或目录，destPath是压缩后目标文件
func UnZip(zipFile string, destPath string) error //解压文件，并将文件存储在destPath目录中
```

### 5. formatter格式化处理包

- 格式化相关处理函数
- 导入包：import "github.com/duke-git/lancet/formatter"

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

- 函数列表：

```go
func Comma(v interface{}, symbol string) string  //用逗号每隔3位分割数字/字符串
```

### 6. function包可以控制函数执行，支持部分函数式编程

- 控制函数执行，支持部分函数式编程
- 导入包：import "github.com/duke-git/lancet/function"

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
func After(n int, fn interface{}) func(args ...interface{}) []reflect.Value  //创建一个函数, 只有在运行了n次之后才有效果
func Before(n int, fn interface{}) func(args ...interface{}) []reflect.Value  //创建一个函数,调用不超过n次。 当n已经达到时，最后一个函数调用的结果将被记住并返回
func (f Fn) Curry(i interface{}) func(...interface{}) interface{}  //函数柯里化
func Compose(fnList ...func(...interface{}) interface{}) func(...interface{}) interface{}  //从右至左组合函数
func Delay(delay time.Duration, fn interface{}, args ...interface{})  //延迟调用函数
func Schedule(d time.Duration, fn interface{}, args ...interface{}) chan bool //每隔duration时间调用函数, 关闭返回通道可以停止调用
func (w *Watcher) Start() //开时watcher
func (w *Watcher) Stop() //开时watcher
func (w *Watcher) Reset() {} //重置代码watcher
func (w *Watcher) GetElapsedTime() time.Duration //get code excution elapsed time.
```

### 7. netutil网络处理包

- 处理ip, http请求相关函数
- 导入包：import "github.com/duke-git/lancet/netutil"
- http方法params参数顺序：header, query string, body, httpclient

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

- 函数列表：

```go
func GetInternalIp() string //获取内部ip
func GetPublicIpInfo() (*PublicIpInfo, error) //获取公共ip信息: country, region, isp, city, lat, lon, ip
func IsPublicIP(IP net.IP) bool //判断ip是否为公共ip
func HttpGet(url string, params ...interface{}) (*http.Response, error) //http get请求
func HttpPost(url string, params ...interface{}) (*http.Response, error) //http post请求
func HttpPut(url string, params ...interface{}) (*http.Response, error) //http put请求
func HttpDelete(url string, params ...interface{}) (*http.Response, error) //http delete请求
func HttpPatch(url string, params ...interface{}) (*http.Response, error) //http patch请求
func ConvertMapToQueryString(param map[string]interface{}) string //将map转换成url query string
func ParseHttpResponse(resp *http.Response, obj interface{}) error //将http响应解码成特定interface
```

### 8. random随机数处理包

- 生成和处理随机数
- 导入包：import "github.com/duke-git/lancet/random"

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

- 函数列表：

```go
func RandBytes(length int) []byte //生成随机[]byte
func RandInt(min, max int) int //生成随机int
func RandString(length int) string //生成随机string
```

### 9. retry重试执行函数

- 重试执行函数直到函数成功或被context停止
- 默认重试次数5, 默认执行间隔3秒.
- Usage: import "github.com/duke-git/lancet/retry".

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/retry"
)

func main() {
    var number int
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryDuration(time.Microsecond*50))

    fmt.Println(number) //3
}
```

- Function list:

```go
type RetryFunc func() error //要重试执行的函数
func RetryTimes(n uint) //设置重试次数，默认5次
func RetryDuration(d time.Duration) //设置重试间隔时间，默认3秒
func Context(ctx context.Context) //context config
func Retry(retryFunc RetryFunc, opts ...Option) error //重试函数
```

### 10. slice切片操作包

- 切片操作相关函数
- 导入包：import "github.com/duke-git/lancet/slice"
- 由于go目前对范型支持不稳定，slice处理函数参数和返回值大部分为interface{}, 待范型特性稳定后，会重构相关函数

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

- 函数列表：

```go
func Contain[T comparable](slice []T, value T) bool //判断slice是否包含value
func ContainSubSlice[T comparable](slice, subslice []T) bool  //判断slice是否包含subslice
func Chunk[T any](slice []T, size int) [][]T //均分slice
func Difference[T comparable](slice1, slice2 []T) []T //返回切片，其元素在slice1中，不在slice2中
func DeleteByIndex[T any](slice []T, start int, end ...int) []T //删除切片中start到end位置的值(不包含end)
func Drop[T any](slice []T, n int) []T //创建一个新切片，当n大于0时删除原切片前n个元素，当n小于0时删除原切片后n个元素
func Every[T any](slice []T, fn func(index int, t T) bool) bool //slice中所有元素都符合函数条件时返回true, 否则返回false. 函数签名：func(int, t T) bool
func None[T any](slice []T, fn func(index int, t T) bool) bool //slice中所有元素都不符合函数条件时返回true, 否则返回false. 函数签名：func(int, T) bool
func Find[T any](slice []T, fn func(index int, t T) bool) (*T, bool)//查找slice中第一个符合条件的元素，函数签名：func(int, T) bool
func Filter [T any] (slice []T, fn func(index int, t T) bool) []T //过滤slice
func FlattenDeep(slice interface{}) interface{} //将slice递归为一维切片
func ForEach [T any] (slice []T, fn func(index int, t T)) //遍历切片，在每个元素上执行函数
func IntSlice(slice interface{}) ([]int, error) //转成int切片
func InterfaceSlice(slice interface{}) []interface{} //转成interface{}切片
func Intersection[T comparable](slices ...[]T) []T //slice交集，去重
func InsertByIndex[T any](slice []T, index int, value interface{}) []T //在切片中index位置插入value
func Map [T any, U any] (slice []T, fn func(index int, t T) U) []U //遍历切片
func Reverse[T any](slice []T) //反转切片
func Reduce[T any](slice []T, fn func(index int, t1, t2 T) T, initial T) T //切片reduce操作
func Shuffle[T any](slice []T) []T  //创建一个被打乱值的切片
func Some[T any](slice []T, fn func(index int, t T) bool) bool //slice中任意一个元素都符合函数条件时返回true, 否则返回false.
func SortByField(slice interface{}, field string, sortType ...string) error //对struct切片进行排序
func StringSlice(slice interface{}) []string //转为string切片
func Unique[T comparable](slice []T) []T //去重切片
func Union[T comparable](slices ...[]T) []T //slice并集, 去重
func UpdateByIndex[T any](slice []T, index int, value T) []T //在切片中index位置更新value
func Without[T comparable](slice []T, values ...T) []T //slice去除values
func GroupBy[T any](slice []T, fn func(index int, t T) bool) ([]T, []T) //根据函数function的逻辑分slice为两组slice
func Count[T any](slice []T, fn func(index int, t T) bool) int //遍历slice的元素，返回所有匹配元素的计数
```

### 11. strutil字符串处理包

- 字符串操作相关函数
- 导入包：import "github.com/duke-git/lancet/strutil"

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

- 函数列表：

```go
func After(s, char string) string //截取字符串中char第一次出现之后的字符串
func AfterLast(s, char string) string //截取字符串中char最后一次出现之后的字符串
func Before(s, char string) string //截取字符串中char第一次出现之前的字符串
func BeforeLast(s, char string) string //截取字符串中char最后一次出现之前的字符串
func CamelCase(s string) string //字符串转为cameCase, "foo bar" -> "fooBar"
func Capitalize(s string) string //字符串转为Capitalize, "fOO" -> "Foo"
func IsString(v interface{}) bool //判断是否是字符串
func KebabCase(s string) string //字符串转为KebabCase, "foo_Bar" -> "foo-bar"
func LowerFirst(s string) string //字符串的第一个字母转为小写字母
func PadEnd(source string, size int, padStr string) string //字符串末尾填充size个字符
func PadStart(source string, size int, padStr string) string//字符串开头填充size个字符
func ReverseStr(s string) string //字符串逆序
func Wrap(str string, wrapWith string) string //包裹字符串 Wrap("abc", "*") -> *abc*.
func Unwrap(str string, wrapToken string) string //解包裹字符串 Wrap("*abc*", "*") -> abc.
func SnakeCase(s string) string //字符串转为SnakeCase, "fooBar" -> "foo_bar"
```

### 12. system系统包

- 包含一些操作系统，运行时，shell命令执行的函数.
- Usage: import "github.com/duke-git/lancet/system".

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/system"
)

func main() {
    envFoo := system.GetOsEnv("foo")
    fmt.Println(envFoo)
}
```

- Function list:

```go
func IsWindows() bool //判断操作系统是windows
func IsLinux() bool //判断操作系统是linux
func IsMac() bool //判断操作系统是macos
func GetOsEnv(key string) string //获取名称为key的环境变量
func SetOsEnv(key, value string) error //设置环境变量
func RemoveOsEnv(key string) error //删除指定key的环境变量
func CompareOsEnv(key, comparedEnv string) bool //获取名称为key的环境变量并和comparedEnv比较
func ExecCommand(command string) (stdout, stderr string, err error) //执行shell命令（/bin/bash)
```

### 13. validator验证器包

- 数据校验相关函数
- 导入包：import "github.com/duke-git/lancet/validator"

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

- 函数列表：

```go
func ContainChinese(s string) bool //判断字符串中是否含有中文字符
func IsAlpha(s string) bool //判断字符串是否只含有字母
func IsBase64(base64 string) bool //判断字符串是base64
func IsAllUpper(str string) bool //断字符串是否全是大写字母
func IsAllLower(str string) bool //断字符串是否全是小写字母
func ContainUpper(str string) bool //判断字符串是否包含大写字母
func ContainLower(str string) bool //判断字符串是否包含小写字母
func ContainLetter(str string) bool //判断字符串是否包含字母
func IsJSON(str string) bool //判断字符串是否是JSON
func IsChineseMobile(mobileNum string) bool //判断字符串是否是手机号
func IsChineseIdNum(id string) bool //判断字符串是否是身份证号
func IsChinesePhone(phone string) bool //判断字符串是否是座机电话号码
func IsCreditCard(creditCart string) bool //判断字符串是否是信用卡
func IsDns(dns string) bool //判断字符串是否是DNS
func IsEmail(email string) bool //判断字符串是否是邮箱
func IsEmptyString(s string) bool //判断字符串是否为空
func IsFloatStr(s string) bool //判断字符串是否可以转成float
func IsNumberStr(s string) bool //判断字符串是否可以转成数字
func IsRegexMatch(s, regex string) bool //判断字符串是否match正则表达式
func IsIntStr(s string) bool //判断字符串是否可以转成整数
func IsIp(ipstr string) bool //判断字符串是否是ip
func IsIpV4(ipstr string) bool //判断字符串是否是ipv4
func IsIpV6(ipstr string) bool //判断字符串是否是ipv6
func IsStrongPassword(password string, length int) bool //判断字符串是否是强密码（大小写字母+数字+特殊字符）
func IsWeakPassword(password string) bool //判断字符串是否是弱密码（只有字母或数字）
```
