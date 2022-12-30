<div align=center>
<a href="https://uvdream.github.io/lancet-docs/"><img src="./logo.png" width="200" height="200"/><a/>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-2.1.12-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com/duke-git/lancet/v2?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/duke-git/lancet/v2)](https://goreportcard.com/report/github.com/duke-git/lancet/v2)
[![test](https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/duke-git/lancet/actions/workflows/codecov.yml)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/main/LICENSE)

</div>

<div STYLE="page-break-after: always;"></div>
<p style="font-size: 18px">
    lancet（柳叶刀）是一个全面、高效、可复用的go语言工具函数库。 lancet受到了java apache common包和lodash.js的启发。
</p>

简体中文 | [English](./README.md)

## 特性

-   👏 全面、高效、可复用
-   💪 300+常用 go 工具函数，支持 string、slice、datetime、net、crypt...
-   💅 只依赖 go 标准库
-   🌍 所有导出函数单元测试覆盖率 100%

## 安装

### Note:

1. <b>使用go1.18及以上版本的用户，建议安装v2.x.x。 因为v2.x.x应用go1.18的泛型重写了大部分函数。</b>

```go
go get github.com/duke-git/lancet/v2 //安装v2最新版本v2.x.x
```

2. <b>使用go1.18以下版本的用户，必须安装v1.x.x。目前最新的v1版本是v1.3.5。</b>

```go
go get github.com/duke-git/lancet@v1.3.5 // 使用go1.18以下版本, 必须安装v1.x.x版本
```

## 用法

lancet 是以包的结构组织代码的，使用时需要导入相应的包名。例如：如果使用字符串相关函数，需要导入 strutil 包:

```go
import "github.com/duke-git/lancet/v2/strutil"
```

## 例子

此处以字符串工具函数 Reverse（逆序字符串）为例，需要导入 strutil 包:

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    s := "hello"
    rs := strutil.Reverse(s)
    fmt.Println(rs) //olleh
}
```

## 文档

### 1. algorithm包实现一些基本查找和排序算法。

```go
import "github.com/duke-git/lancet/v2/algorithm"
```

#### Function list:
- **<big>BubbleSort</big>** : 使用冒泡排序算法对切片进行排序。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#BubbleSort)] 
[[play](https://go.dev/play/p/GNdv7Jg2Taj)]
- **<big>CountSort</big>** : 使用计数排序算法对切片进行排序。不改变原数据。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#CountSort)] 
[[play](https://go.dev/play/p/tB-Umgm0DrP)]
- **<big>HeapSort</big>** : 使用堆排序算法对切片进行排序。 
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#HeapSort)] 
[[play](https://go.dev/play/p/u6Iwa1VZS_f)]
- **<big>InsertionSort</big>** : 使用插入排序算法对切片进行排序。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#InsertionSort)] 
[[play](https://go.dev/play/p/G5LJiWgJJW6)]
- **<big>MergeSort</big>** : 使用合并排序算法对切片进行排序。 
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#MergeSort)] 
[[play](https://go.dev/play/p/ydinn9YzUJn)]
- **<big>QuickSort</big>** : 使用快速排序算法对切片进行排序。 
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#QuickSort)] 
[[play](https://go.dev/play/p/7Y7c1Elk3ax)]
- **<big>SelectionSort</big>** : 使用选择排序算法对切片进行排序。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#SelectionSort)] 
[[play](https://go.dev/play/p/oXovbkekayS)]
- **<big>ShellSort</big>** : 使用希尔排序算法对切片进行排序。 
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#ShellSort)] 
[[play](https://go.dev/play/p/3ibkszpJEu3)]
- **<big>BinarySearch</big>** : 返回排序切片中目标值的索引，使用二分搜索（递归调用）。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#BinarySearch)]
[[play](https://go.dev/play/p/t6MeGiUSN47)]
- **<big>BinaryIterativeSearch</big>** :返回排序切片中目标值的索引，使用二分搜索（非递归）。 
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#BinaryIterativeSearch)]
[[play](https://go.dev/play/p/Anozfr8ZLH3)]
- **<big>LinearSearch</big>** : 基于传入的相等函数返回切片中目标值的索引。（线性查找）
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#LinearSearch)]
- **<big>LRUCache</big>** : 应用lru算法实现内存缓存. 
[[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#LRUCache)]
[[play](https://go.dev/play/p/-EZjgOURufP)]

### 2. concurrency包含一些支持并发编程的功能。例如：goroutine, channel, async 等。

```go
import "github.com/duke-git/lancet/v2/concurrency"
```

#### Function list:

-   [NewChannel](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#NewChannel)
-   [Bridge](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Bridge)
-   [FanIn](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#FanIn)
-   [Generate](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Generate)
-   [Or](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Or)
-   [OrDone](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#OrDone)
-   [Repeat](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Repeat)
-   [RepeatFn](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#RepeatFn)
-   [Take](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Take)
-   [Tee](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Tee)

### 3. condition包含一些用于条件判断的函数。

```go
import "github.com/duke-git/lancet/v2/condition"
```

#### Function list:

-   [Bool](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Bool)
-   [And](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#And)
-   [Or](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Or)
-   [Xor](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Xor)
-   [Nor](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Nor)
-   [Nand](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Nand)
-   [TernaryOperator](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#TernaryOperator)

### 4. convertor 转换器包支持一些常见的数据类型转换。

```go
import "github.com/duke-git/lancet/v2/convertor"
```

#### 函数列表:

-   [ColorHexToRGB](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ColorHexToRGB)
-   [ColorRGBToHex](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ColorRGBToHex)
-   [ToBool](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToBool)
-   [ToBytes](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToBytes)
-   [ToChar](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToChar)
-   [ToChannel](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToChannel)
-   [ToFloat](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToFloat)
-   [ToInt](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToInt)
-   [ToJson](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToJson)
-   [ToMap](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToMap)
-   [ToPointer](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToPointer)
-   [ToString](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToString)
-   [StructToMap](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#StructToMap)
-   [MapToSlice](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#MapToSlice)
-   [EncodeByte](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#EncodeByte)
-   [DecodeByte](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#DecodeByte)

### 5. cryptor 加密包支持数据加密和解密，获取md5，hash 值。支持base64, md5, hmac, aes, des, rsa。

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### 函数列表:

-   [AesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesEcbEncrypt)
-   [AesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesEcbDecrypt)
-   [AesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCbcEncrypt)
-   [AesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCbcDecrypt)
-   [AesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCtrCrypt)
-   [AesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCfbEncrypt)
-   [AesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCfbDecrypt)
-   [AesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesOfbEncrypt)
-   [AesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesOfbDecrypt)
-   [Base64StdEncode](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Base64StdEncode)
-   [Base64StdDecode](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Base64StdDecode)
-   [DesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesEcbEncrypt)
-   [DesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesEcbDecrypt)
-   [DesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCbcEncrypt)
-   [DesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCbcDecrypt)
-   [DesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCtrCrypt)
-   [DesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCfbEncrypt)
-   [DesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCfbDecrypt)
-   [DesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesOfbEncrypt)
-   [DesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesOfbDecrypt)
-   [HmacMd5](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacMd5)
-   [HmacSha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha1)
-   [HmacSha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha256)
-   [HmacSha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha512)
-   [Md5String](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Md5String)
-   [Md5File](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Md5File)
-   [Sha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha1)
-   [Sha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha256)
-   [Sha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha512)
-   [GenerateRsaKey](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#GenerateRsaKey)
-   [RsaEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#RsaEncrypt)
-   [RsaDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#RsaDecrypt)

### 6. datetime 日期时间处理包，格式化日期，比较日期。

```go
import "github.com/duke-git/lancet/v2/datetime"
```

#### 函数列表:

-   [AddDay](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddDay)
-   [AddHour](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddHour)
-   [AddMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddMinute)
-   [BeginOfMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMinute)
-   [BeginOfHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfHour)
-   [BeginOfDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfDay)
-   [BeginOfWeek](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfWeek)
-   [BeginOfMonth](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMonth)
-   [BeginOfYear](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfYear)
-   [EndOfMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMinute)
-   [EndOfHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfHour)
-   [EndOfDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfDay)
-   [EndOfWeek](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfWeek)
-   [EndOfMonth](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMonth)
-   [EndOfYear](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfYear)
-   [GetNowDate](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowDate)
-   [GetNowTime](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowTime)
-   [GetNowDateTime](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowDateTime)
-   [GetZeroHourTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetZeroHourTimestamp)
-   [GetNightTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNightTimestamp)
-   [FormatTimeToStr](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#FormatTimeToStr)
-   [FormatStrToTime](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#FormatStrToTime)
-   [NewUnix](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewUnix)
-   [NewUnixNow](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewUnixNow)
-   [NewFormat](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewFormat)
-   [NewISO8601](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewISO8601)
-   [ToUnix](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToUnix)
-   [ToFormat](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToFormat)
-   [ToFormatForTpl](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToFormatForTpl)
-   [ToIso8601](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToIso8601)

### 7. datastructure 包含一些普通的数据结构实现。例如：list, linklist, stack, queue, set, tree, graph.

```go
import list "github.com/duke-git/lancet/v2/datastructure/list"
import link "github.com/duke-git/lancet/v2/datastructure/link"
import stack "github.com/duke-git/lancet/v2/datastructure/stack"
import queue "github.com/duke-git/lancet/v2/datastructure/queue"
import set "github.com/duke-git/lancet/v2/datastructure/set"
import tree "github.com/duke-git/lancet/v2/datastructure/tree"
import heap "github.com/duke-git/lancet/v2/datastructure/heap"
import hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
```

#### Function list:

-   [List](https://github.com/duke-git/lancet/blob/main/docs/datastructure/list_zh-CN.md)
-   [Linklist](https://github.com/duke-git/lancet/blob/main/docs/datastructure/linklist_zh-CN.md)
-   [Stack](https://github.com/duke-git/lancet/blob/main/docs/datastructure/stack_zh-CN.md)
-   [Queue](https://github.com/duke-git/lancet/blob/main/docs/datastructure/queue_zh-CN.md)
-   [Set](https://github.com/duke-git/lancet/blob/main/docs/datastructure/set_zh-CN.md)
-   [Tree](https://github.com/duke-git/lancet/blob/main/docs/datastructure/tree_zh-CN.md)
-   [Heap](https://github.com/duke-git/lancet/blob/main/docs/datastructure/heap.md)
-   [HashMap](https://github.com/duke-git/lancet/blob/main/docs/datastructure/hashmap.md)

### 8. fileutil 包含文件基本操作。

```go
import "github.com/duke-git/lancet/v2/fileutil"
```

#### 函数列表：

-   [ClearFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ClearFile)
-   [CreateFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CreateFile)
-   [CreateDir](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CreateDir)
-   [CopyFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CopyFile)
-   [FileMode](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#FileMode)
-   [MiMeType](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#MiMeType)
-   [IsExist](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsExist)
-   [IsLink](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsLink)
-   [IsDir](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsDir)
-   [ListFileNames](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ListFileNames)
-   [RemoveFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#RemoveFile)
-   [ReadFileToString](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ReadFileToString)
-   [ReadFileByLine](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ReadFileByLine)
-   [Zip](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#Zip)
-   [UnZip](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#UnZip)

### 9. formatter 格式化器包含一些数据格式化处理方法。

```go
import "github.com/duke-git/lancet/v2/formatter"
```

#### 函数列表:

-   [Comma](https://github.com/duke-git/lancet/blob/main/docs/formatter_zh-CN.md#Comma)

### 10. function 函数包控制函数执行流程，包含部分函数式编程。

```go
import "github.com/duke-git/lancet/v2/function"
```

#### 函数列表:

-   [After](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#After)
-   [Before](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Before)
-   [Curry](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Curry)
-   [Compose](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Compose)
-   [Debounced](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Debounced)
-   [Delay](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Delay)
-   [Pipeline](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Pipeline)
-   [Watcher](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Watcher)

### 11. maputil包括一些操作map的函数.

```go
import "github.com/duke-git/lancet/v2/maputil"
```

#### 函数列表:

-   [ForEach](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#ForEach)
-   [Filter](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Filter)
-   [Intersect](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Intersect)
-   [Keys](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Keys)
-   [Merge](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Merge)
-   [Minus](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Minus)
-   [Values](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Values)
-   [IsDisjoint](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#IsDisjoint)

### 12. mathutil包实现了一些数学计算的函数。

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### Function list:

-   [Average](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Average)
-   [Exponent](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Exponent)
-   [Fibonacci](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Fibonacci)
-   [Factorial](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Factorial)
-   [Max](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Max)
-   [MaxBy](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#MaxBy)
-   [Min](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Min)
-   [MinBy](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#MinBy)
-   [Percent](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Percent)
-   [RoundToFloat](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#RoundToFloat)
-   [RoundToString](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#RoundToString)
-   [TruncRound](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#TruncRound)

### 13. netutil网络包支持获取ip地址，发送http请求。

```go
import "github.com/duke-git/lancet/v2/netutil"
```

#### 函数列表:

-   [ConvertMapToQueryString](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#ConvertMapToQueryString)
-   [GetInternalIp](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetInternalIp)
-   [EncodeUrl](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#EncodeUrl)
-   [GetIps](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetIps)
-   [GetMacAddrs](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetMacAddrs)
-   [GetPublicIpInfo](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetPublicIpInfo)
-   [GetRequestPublicIp](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetRequestPublicIp)
-   [IsPublicIP](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#IsPublicIP)
-   [IsInternalIP](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#IsInternalIP)
-   [HttpRequest](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpRequest)
-   [HttpClient](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpClient)
-   [SendRequest](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#SendRequest)
-   [DecodeResponse](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#DecodeResponse)
-   [StructToUrlValues](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#StructToUrlValues)
-   [HttpGet<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpGet)
-   [HttpDelete<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpDelete)
-   [HttpPost<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPost)
-   [HttpPut<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPut)
-   [HttpPatch<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPatch)
-   [ParseHttpResponse](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#ParseHttpResponse)

### 14. random随机数生成器包，可以生成随机[]bytes, int, string。

```go
import "github.com/duke-git/lancet/v2/random"
```

#### 函数列表:

-   [RandBytes](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandBytes)
-   [RandInt](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandInt)
-   [RandString](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandString)
-   [RandUpper](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandUpper)
-   [RandLower](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandLower)
-   [RandNumeral](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandNumeral)
-   [RandNumeralOrLetter](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandNumeralOrLetter)
-   [UUIdV4](https://github.com/duke-git/lancet/blob/main/docs/random.md#UUIdV4)

### 15. retry重试执行函数直到函数运行成功或被context cancel。

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### 函数列表:

-   [Context](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Context)
-   [Retry](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Retry)
-   [RetryFunc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryFunc)
-   [RetryDuration](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryDuration)
-   [RetryTimes](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryTimes)

### 16. slice包含操作切片的方法集合。

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### 函数列表:

-   [AppendIfAbsent](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#AppendIfAbsent)
-   [Contain](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Contain)
-   [ContainSubSlice](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ContainSubSlice)
-   [Chunk](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Chunk)
-   [Compact](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Compact)
-   [Concat](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Concat)
-   [Count](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Count)
-   [CountBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#CountBy)
-   [Difference](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Difference)
-   [DifferenceBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DifferenceBy)
-   [DifferenceWith](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DifferenceWith)
-   [DeleteAt](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DeleteAt)
-   [Drop](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Drop)
-   [Every](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Every)
-   [Filter](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Filter)
-   [Find](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Find)
-   [FindLast](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#FindLast)
-   [Flatten](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Flatten)
-   [FlattenDeep](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#FlattenDeep)
-   [ForEach](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ForEach)
-   [GroupBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#GroupBy)
-   [GroupWith](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#GroupWith)
-   [IntSlice<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IntSlice)
-   [InterfaceSlice<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#InterfaceSlice)
-   [Intersection](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Intersection)
-   [InsertAt](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#InsertAt)
-   [IndexOf](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IndexOf)
-   [LastIndexOf](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#LastIndexOf)
-   [Map](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Map)
-   [Merge](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Merge)
-   [Reverse](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Reverse)
-   [Reduce](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Reduce)
-   [Replace](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Replace)
-   [ReplaceAll](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ReplaceAll)
-   [Repeat](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Repeat)
-   [Shuffle](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Shuffle)
-   [Sort](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Sort)
-   [SortBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#SortBy)
-   [SortByField<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#SortByField)
-   [Some](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Some)
-   [StringSlice<sup>deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#StringSlice)
-   [SymmetricDifference](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#SymmetricDifference)
-   [ToSlice](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ToSlice)
-   [ToSlicePointer](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ToSlicePointer)
-   [Unique](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Unique)
-   [UniqueBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#UniqueBy)
-   [Union](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Union)
-   [UniqueBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#UniqueBy)
-   [UpdateAt](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#UpdateAt)
-   [Without](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Without)
-   [KeyBy](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#KeyBy)

### 17. strutil包含字符串处理的相关函数。

```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### 函数列表:
- **<big>After</big>** : 返回源字符串中指定字符串首次出现时的位置之后的子字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#After)] 
[[play](https://go.dev/play/p/RbCOQqCDA7m)]
- **<big>AfterLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之后的子字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#AfterLast)] 
[[play](https://go.dev/play/p/1TegARrb8Yn)]
- **<big>Before</big>** : 返回源字符串中指定字符串第一次出现时的位置之前的子字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Before)] 
[[play](https://go.dev/play/p/JAWTZDS4F5w)]
- **<big>BeforeLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之前的子字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#BeforeLast)] 
[[play](https://go.dev/play/p/pJfXXAoG_Te)]
- **<big>CamelCase</big>** : 将字符串转换为CamelCase驼峰式字符串, 非字母和数字会被忽略。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#CamelCase)]
[[play](https://go.dev/play/p/9eXP3tn2tUy)]
- **<big>Capitalize</big>** : 将字符串的第一个字符转换为大写。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Capitalize)]
[[play](https://go.dev/play/p/2OAjgbmAqHZ)]
- **<big>IsString</big>** : 判断传入参数的数据类型是否为字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#IsString)]
[[play](https://go.dev/play/p/IOgq7oF9ERm)]
- **<big>KebabCase</big>** : 将字符串转换为kebab-case形式字符串, 非字母和数字会被忽略。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#KebabCase)]
[[play](https://go.dev/play/p/dcZM9Oahw-Y)]
- **<big>UpperKebabCase</big>** : 将字符串转换为大写KEBAB-CASE形式字符串, 非字母和数字会被忽略。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#UpperKebabCase)]
[[play](https://go.dev/play/p/zDyKNneyQXk)]
- **<big>LowerFirst</big>** : 将字符串的第一个字符转换为小写形式。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#LowerFirst)]
[[play](https://go.dev/play/p/CbzAyZmtJwL)]
- **<big>UpperFirst</big>** : 将字符串的第一个字符转换为大写形式。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#UpperFirst)]
[[play](https://go.dev/play/p/sBbBxRbs8MM)]
- **<big>PadEnd</big>** : 如果字符串短于限制大小，则在右侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#PadEnd)]
[[play](https://go.dev/play/p/9xP8rN0vz--)]
- **<big>PadStart</big>** : 如果字符串短于限制大小，则在左侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#PadStart)]
[[play](https://go.dev/play/p/xpTfzArDfvT)]
- **<big>Reverse</big>** : 返回字符顺序与给定字符串相反的字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Reverse)]
[[play](https://go.dev/play/p/adfwalJiecD)]
- **<big>SnakeCase</big>** : 将字符串转换为snake_case形式, 非字母和数字会被忽略。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#SnakeCase)]
[[play](https://go.dev/play/p/tgzQG11qBuN)]
- **<big>UpperSnakeCase</big>** : 将字符串转换为大写SNAKE_CASE形式, 非字母和数字会被忽略。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#UpperSnakeCase)]
[[play](https://go.dev/play/p/4COPHpnLx38)]
- **<big>SplitEx</big>** : 拆分给定的字符串可以控制结果切片是否包含空字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#SplitEx)]
[[play](https://go.dev/play/p/Us-ySSbWh-3)]
- **<big>Substring</big>** : 根据指定的位置和长度截取子字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Substring)]
- **<big>Wrap</big>** : 用给定字符包裹传入的字符串
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Wrap)]
[[play](https://go.dev/play/p/KoZOlZDDt9y)]
- **<big>Unwrap</big>** : 从另一个字符串中解开一个给定的字符串。 将更改源字符串。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Unwrap)]
[[play](https://go.dev/play/p/Ec2q4BzCpG-)]

### 18. system包含os, runtime, shell command的相关函数。

```go
import "github.com/duke-git/lancet/v2/system"
```

#### 函数列表:
- **<big>IsWindows</big>** : 检查当前操作系统是否是windows。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#IsWindows)]
[[play](https://go.dev/play/p/XzJULbzmf9m)]
- **<big>IsLinux</big>** : 检查当前操作系统是否是linux。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#IsLinux)]
[[play](https://go.dev/play/p/zIflQgZNuxD)]
- **<big>IsMac</big>** : 检查当前操作系统是否是macos。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#IsMac)]
[[play](https://go.dev/play/p/Mg4Hjtyq7Zc)]
- **<big>GetOsEnv</big>** : 根据key获取对应的环境变量值
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#GetOsEnv)]
[[play](https://go.dev/play/p/D88OYVCyjO-)]
- **<big>SetOsEnv</big>** : 设置环境变量。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#SetOsEnv)]
[[play](https://go.dev/play/p/D88OYVCyjO-)]
- **<big>RemoveOsEnv</big>** : 删除环境变量。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#RemoveOsEnv)]
[[play](https://go.dev/play/p/fqyq4b3xUFQ)]
- **<big>CompareOsEnv</big>** : 换取环境变量并与传入值进行比较。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#CompareOsEnv)]
[[play](https://go.dev/play/p/BciHrKYOHbp)]
- **<big>ExecCommand</big>** : 执行shell命令。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#ExecCommand)]
[[play](https://go.dev/play/p/n-2fLyZef-4)]
- **<big>GetOsBits</big>** : 获取当前操作系统位数(32/64)。
[[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#GetOsBits)]
[[play](https://go.dev/play/p/ml-_XH3gJbW)]

### 19. validator验证器包，包含常用字符串格式验证函数。

```go
import "github.com/duke-git/lancet/v2/validator"
```

#### 函数列表:

-   [ContainChinese](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainChinese)
-   [ContainLetter](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainLetter)
-   [ContainLower](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainLower)
-   [ContainUpper](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainUpper)
-   [IsAlpha](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAlpha)
-   [IsAllUpper](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAllUpper)
-   [IsAllLower](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAllLower)
-   [IsBase64](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsBase64)
-   [IsChineseMobile](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChineseMobile)
-   [IsChineseIdNum](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChineseIdNum)
-   [IsChinesePhone](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChinesePhone)
-   [IsCreditCard](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsCreditCard)
-   [IsDns](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsDns)
-   [IsEmail](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsEmail)
-   [IsEmptyString](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsEmptyString)
-   [IsFloatStr](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsFloatStr)
-   [IsNumberStr](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsNumberStr)
-   [IsJSON](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsJSON)
-   [IsRegexMatch](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsRegexMatch)
-   [IsIntStr](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIntStr)
-   [IsIp](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIp)
-   [IsIpV4](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIpV4)
-   [IsIpV6](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIpV6)
-   [IsStrongPassword](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsStrongPassword)
-   [IsUrl](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsUrl)
-   [IsWeakPassword](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsWeakPassword)
-   [IsZeroValue](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsZeroValue)
-   [IsGBK](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsGBK)

### 20. xerror包实现一些错误处理函数

```go
import "github.com/duke-git/lancet/v2/xerror"
```

#### 函数列表:

-   [Unwrap](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#Unwrap)

## 如何贡献代码

非常感激任何的代码提交以使 lancet 的功能越来越强大。创建 pull request 时请遵守以下规则。

1. Fork lancet 仓库。
2. 创建自己的特性分支。
3. 提交变更。
4. Push 分支。
5. 创建新的 pull request。
