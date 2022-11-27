<div align=center>
<a href="https://uvdream.github.io/lancet-docs/en/"><img src="./logo.png" width="200" height="200"/></a>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-2.1.10-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com/duke-git/lancet/v2?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/duke-git/lancet/v2)](https://goreportcard.com/report/github.com/duke-git/lancet/v2)
[![test](https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/duke-git/lancet/actions/workflows/codecov.yml)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/main/LICENSE)

</div>

<div STYLE="page-break-after: always;"></div>

<p style="font-size: 20px"> 
    Lancet is a comprehensive, efficient, and reusable util function library of go. Inspired by the java apache common package and lodash.js.
</p>

English | [简体中文](./README_zh-CN.md) | [Website](https://uvdream.github.io/lancet-docs/en/)

## Feature

-   👏 Comprehensive, efficient and reusable.
-   💪 300+ go util functions, support string, slice, datetime, net, crypt...
-   💅 Only depend on the go standard library.
-   🌍 Unit test for every exported function.

## Installation

### Note:

1. <b>For users who use go1.18 and above, it is recommended to install lancet v2.x.x. Cause v2.x.x rewrite all functions with generics of go1.18.</b>

```go
go get github.com/duke-git/lancet/v2 // will install latest version of v2.x.x
```

2. <b>For users who use version below go1.18, you should install v1.x.x. now latest v1 is v1.3.4. </b>

```go
go get github.com/duke-git/lancet@v1.3.3 // below go1.18, install latest version of v1.x.x
```

## Usage

Lancet organizes the code into package structure, and you need to import the corresponding package name when use it. For example, if you use string-related functions,import the strutil package like below:

```go
import "github.com/duke-git/lancet/v2/strutil"
```

## Example

Here takes the string function Reverse (reverse order string) as an example, and the strutil package needs to be imported.

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

## API Documentation

## [lancet API doc](https://uvdream.github.io/lancet-docs/) Thanks [@UvDream](https://github.com/UvDream) for contributing.

### 1. Algorithm package implements some basic algorithm. eg. sort, search.

```go
import "github.com/duke-git/lancet/v2/algorithm"
```

#### Function list:

-   [BubbleSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BubbleSort)
-   [CountSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#CountSort)
-   [HeapSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#HeapSort)
-   [InsertionSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#InsertionSort)
-   [MergeSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#MergeSort)
-   [QuickSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#QuickSort)
-   [SelectionSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#SelectionSort)
-   [ShellSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#ShellSort)
-   [BinarySearch](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BinarySearch)
-   [BinaryIterativeSearch](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BinaryIterativeSearch)
-   [LinearSearch](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#LinearSearch)
-   [LRUCache](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#LRUCache)

### 2. Concurrency package contain some functions to support concurrent programming. eg, goroutine, channel, async.

```go
import "github.com/duke-git/lancet/v2/concurrency"
```

#### Function list:

-   [NewChannel](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#NewChannel)
-   [Bridge](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Bridge)
-   [FanIn](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#FanIn)
-   [Generate](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Generate)
-   [Or](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Or)
-   [OrDone](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#OrDone)
-   [Repeat](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Repeat)
-   [RepeatFn](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#RepeatFn)
-   [Take](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Take)
-   [Tee](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Tee)

### 3. Condition package contains some functions for conditional judgment. eg. And, Or, TernaryOperator...

```go
import "github.com/duke-git/lancet/v2/condition"
```

#### Function list:

-   [Bool](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Bool)
-   [And](https://github.com/duke-git/lancet/blob/main/docs/condition.md#And)
-   [Or](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Or)
-   [Xor](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Xor)
-   [Nor](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Nor)
-   [Nand](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Nand)
-   [TernaryOperator](https://github.com/duke-git/lancet/blob/main/docs/condition.md#TernaryOperator)

### 4. Convertor package contains some functions for data convertion.

```go
import "github.com/duke-git/lancet/v2/convertor"
```

#### Function list:

-   [ColorHexToRGB](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ColorHexToRGB)
-   [ColorRGBToHex](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ColorRGBToHex)
-   [ToBool](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToBool)
-   [ToBytes](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToBytes)
-   [ToChar](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToChar)
-   [ToChannel](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToChannel)
-   [ToFloat](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToFloat)
-   [ToInt](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToInt)
-   [ToJson](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToJson)
-   [ToMap](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToMap)
-   [ToPointer](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToPointer)
-   [ToString](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToString)
-   [StructToMap](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#StructToMap)
-   [MapToSlice](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#MapToSlice)
-   [EncodeByte](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#EncodeByte)
-   [DecodeByte](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#DecodeByte)

### 5. Cryptor package is for data encryption and decryption.

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### Function list:

-   [AesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesEcbEncrypt)
-   [AesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesEcbDecrypt)
-   [AesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCbcEncrypt)
-   [AesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCbcDecrypt)
-   [AesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCtrCrypt)
-   [AesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCfbEncrypt)
-   [AesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCfbDecrypt)
-   [AesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesOfbEncrypt)
-   [AesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesOfbDecrypt)
-   [Base64StdEncode](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Base64StdEncode)
-   [Base64StdDecode](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Base64StdDecode)
-   [DesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesEcbEncrypt)
-   [DesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesEcbDecrypt)
-   [DesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCbcEncrypt)
-   [DesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCbcDecrypt)
-   [DesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCtrCrypt)
-   [DesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCfbEncrypt)
-   [DesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCfbDecrypt)
-   [DesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesOfbEncrypt)
-   [DesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesOfbDecrypt)
-   [HmacMd5](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacMd5)
-   [HmacSha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha1)
-   [HmacSha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha256)
-   [HmacSha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha512)
-   [Md5String](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Md5String)
-   [Md5File](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Md5File)
-   [Sha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha1)
-   [Sha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha256)
-   [Sha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha512)
-   [GenerateRsaKey](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#GenerateRsaKey)
-   [RsaEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#RsaEncrypt)
-   [RsaDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#RsaDecrypt)

### 6. Datetime package supports date and time format and compare.

```go
import "github.com/duke-git/lancet/v2/datetime"
```

#### Function list:

-   [AddDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddDay)
-   [AddHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddHour)
-   [AddMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddMinute)
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
-   [GetNowDate](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowDate)
-   [GetNowTime](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowTime)
-   [GetNowDateTime](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowDateTime)
-   [GetZeroHourTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetZeroHourTimestamp)
-   [GetNightTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNightTimestamp)
-   [FormatTimeToStr](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#FormatTimeToStr)
-   [FormatStrToTime](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#FormatStrToTime)
-   [NewUnix](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewUnix)
-   [NewUnixNow](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewUnixNow)
-   [NewFormat](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewFormat)
-   [NewISO8601](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewISO8601)
-   [ToUnix](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToUnix)
-   [ToFormat](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToFormat)
-   [ToFormatForTpl](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToFormatForTpl)
-   [ToIso8601](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToIso8601)

### 7. Datastructure package constains some common data structure. eg. list, linklist, stack, queue, set, tree, graph.

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

-   [List](https://github.com/duke-git/lancet/blob/main/docs/datastructure/list.md)
-   [Linklist](https://github.com/duke-git/lancet/blob/main/docs/datastructure/linklist.md)
-   [Stack](https://github.com/duke-git/lancet/blob/main/docs/datastructure/stack.md)
-   [Queue](https://github.com/duke-git/lancet/blob/main/docs/datastructure/queue.md)
-   [Set](https://github.com/duke-git/lancet/blob/main/docs/datastructure/set.md)
-   [Tree](https://github.com/duke-git/lancet/blob/main/docs/datastructure/tree.md)
-   [Heap](https://github.com/duke-git/lancet/blob/main/docs/datastructure/heap.md)
-   [HashMap](https://github.com/duke-git/lancet/blob/main/docs/datastructure/hashmap.md)

### 8. Fileutil package implements some basic functions for file operations.

```go
import "github.com/duke-git/lancet/v2/fileutil"
```

#### Function list：

-   [ClearFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ClearFile)
-   [CreateFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CreateFile)
-   [CreateDir](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CreateDir)
-   [CopyFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CopyFile)
-   [FileMode](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#FileMode)
-   [MiMeType](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#MiMeType)
-   [IsExist](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsExist)
-   [IsLink](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsLink)
-   [IsDir](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsDir)
-   [ListFileNames](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ListFileNames)
-   [RemoveFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#RemoveFile)
-   [ReadFileToString](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ReadFileToString)
-   [ReadFileByLine](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ReadFileByLine)
-   [Zip](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#Zip)
-   [UnZip](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#UnZip)

### 9. Formatter contains some functions for data formatting.

```go
import "github.com/duke-git/lancet/v2/formatter"
```

#### Function list:

-   [Comma](https://github.com/duke-git/lancet/blob/main/docs/formatter.md#Comma)

### 10. Function package can control the flow of function execution and support part of functional programming

```go
import "github.com/duke-git/lancet/v2/function"
```

#### Function list:

-   [After](https://github.com/duke-git/lancet/blob/main/docs/function.md#After)
-   [Before](https://github.com/duke-git/lancet/blob/main/docs/function.md#Before)
-   [Curry](https://github.com/duke-git/lancet/blob/main/docs/function.md#Curry)
-   [Compose](https://github.com/duke-git/lancet/blob/main/docs/function.md#Compose)
-   [Debounced](https://github.com/duke-git/lancet/blob/main/docs/function.md#Debounced)
-   [Delay](https://github.com/duke-git/lancet/blob/main/docs/function.md#Delay)
-   [Pipeline](https://github.com/duke-git/lancet/blob/main/docs/function.md#Pipeline)
-   [Watcher](https://github.com/duke-git/lancet/blob/main/docs/function.md#Watcher)

### 11. Maputil package includes some functions to manipulate map.

```go
import "github.com/duke-git/lancet/v2/maputil"
```

#### Function list:

-   [ForEach](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#ForEach)
-   [Filter](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Filter)
-   [Intersect](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Intersect)
-   [Keys](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Keys)
-   [Merge](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Merge)
-   [Minus](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Minus)
-   [Values](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Values)
-   [IsDisjoint](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#IsDisjoint)

### 12. Mathutil package implements some functions for math calculation.

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### Function list:

-   [Average](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Average)
-   [Exponent](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Exponent)
-   [Fibonacci](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Fibonacci)
-   [Factorial](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Factorial)
-   [Max](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Max)
-   [MaxBy](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#MaxBy)
-   [Min](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Min)
-   [MinBy](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#MinBy)
-   [Percent](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Percent)
-   [RoundToFloat](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#RoundToFloat)
-   [RoundToString](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#RoundToString)
-   [TruncRound](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#TruncRound)

### 13. Netutil package contains functions to get net information and send http request.

```go
import "github.com/duke-git/lancet/v2/netutil"
```

#### Function list:

-   [ConvertMapToQueryString](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#ConvertMapToQueryString)
-   [EncodeUrl](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#EncodeUrl)
-   [GetInternalIp](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetInternalIp)
-   [GetIps](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetIps)
-   [GetMacAddrs](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetMacAddrs)
-   [GetPublicIpInfo](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetPublicIpInfo)
-   [GetRequestPublicIp](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetRequestPublicIp)
-   [IsPublicIP](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#IsPublicIP)
-   [IsInternalIP](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#IsInternalIP)
-   [HttpRequest](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpRequest)
-   [HttpClient](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpClient)
-   [SendRequest](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#SendRequest)
-   [DecodeResponse](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#DecodeResponse)
-   [StructToUrlValues](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#StructToUrlValues)

-   [HttpGet<sup>Deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpGet)
-   [HttpDelete<sup>Deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpDelete)
-   [HttpPost<sup>Deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPost)
-   [HttpPut<sup>Deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPut)
-   [HttpPatch<sup>Deprecated</sup>](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPatch)
-   [ParseHttpResponse](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#ParseHttpResponse)

### 14. Random package implements some basic functions to generate random int and string.

```go
import "github.com/duke-git/lancet/v2/random"
```

#### Function list:

-   [RandBytes](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandBytes)
-   [RandInt](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandInt)
-   [RandString](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandString)
-   [RandUpper](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandUpper)
-   [RandLower](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandLower)
-   [RandNumeral](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandNumeral)
-   [RandNumeralOrLetter](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandNumeralOrLetter)
-   [UUIdV4](https://github.com/duke-git/lancet/blob/main/docs/random.md#UUIdV4)

### 15. Retry package is for executing a function repeatedly until it was successful or canceled by the context.

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### Function list:

-   [Context](https://github.com/duke-git/lancet/blob/main/docs/retry.md#Context)
-   [Retry](https://github.com/duke-git/lancet/blob/main/docs/retry.md#Retry)
-   [RetryFunc](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryFunc)
-   [RetryDuration](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryDuration)
-   [RetryTimes](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryTimes)

### 16. Slice contains some functions to manipulate slice.

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### Function list:

-   [AppendIfAbsent](https://github.com/duke-git/lancet/blob/main/docs/slice.md#AppendIfAbsent)
-   [Contain](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Contain)
-   [ContainSubSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ContainSubSlice)
-   [Chunk](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Chunk)
-   [Compact](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Compact)
-   [Concat](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Concat)
-   [Count](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Count)
-   [Difference](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Difference)
-   [DifferenceBy](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DifferenceBy)
-   [DifferenceWith](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DifferenceWith)
-   [DeleteAt](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DeleteAt)
-   [Drop](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Drop)
-   [Equal](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Equal)
-   [EqualWith](https://github.com/duke-git/lancet/blob/main/docs/slice.md#EqualWith)
-   [Every](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Every)
-   [Filter](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Filter)
-   [Find](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Find)
-   [FindLast](https://github.com/duke-git/lancet/blob/main/docs/slice.md#FindLast)
-   [Flatten](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Flatten)
-   [FlattenDeep](https://github.com/duke-git/lancet/blob/main/docs/slice.md#FlattenDeep)
-   [ForEach](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ForEach)
-   [GroupBy](https://github.com/duke-git/lancet/blob/main/docs/slice.md#GroupBy)
-   [GroupWith](https://github.com/duke-git/lancet/blob/main/docs/slice.md#GroupWith)
-   [IntSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IntSlice)
-   [InterfaceSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#InterfaceSlice)
-   [Intersection](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Intersection)
-   [InsertAt](https://github.com/duke-git/lancet/blob/main/docs/slice.md#InsertAt)
-   [IndexOf](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IndexOf)
-   [LastIndexOf](https://github.com/duke-git/lancet/blob/main/docs/slice.md#LastIndexOf)
-   [Map](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Map)
-   [Reverse](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Reverse)
-   [Reduce](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Reduce)
-   [Replace](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Replace)
-   [ReplaceAll](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ReplaceAll)
-   [Shuffle](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Shuffle)
-   [SortByField](https://github.com/duke-git/lancet/blob/main/docs/slice.md#SortByField)
-   [Some](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Some)
-   [StringSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#StringSlice)
-   [SymmetricDifference](https://github.com/duke-git/lancet/blob/main/docs/slice.md#SymmetricDifference)
-   [ToSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ToSlice)
-   [ToSlicePointer](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ToSlicePointer)
-   [Unique](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Unique)
-   [UniqueBy](https://github.com/duke-git/lancet/blob/main/docs/slice.md#UniqueBy)
-   [Union](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Union)
-   [UnionBy](https://github.com/duke-git/lancet/blob/main/docs/slice.md#UnionBy)
-   [UpdateAt](https://github.com/duke-git/lancet/blob/main/docs/slice.md#UpdateAt)
-   [Without](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Without)
-   [KeyBy](https://github.com/duke-git/lancet/blob/main/docs/slice.md#KeyBy)

### 17. Strutil package contains some functions to manipulate string.

```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### Function list:

-   [After](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#After)
-   [AfterLast](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#AfterLast)
-   [Before](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Before)
-   [BeforeLast](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#BeforeLast)
-   [CamelCase](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#CamelCase)
-   [Capitalize](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Capitalize)
-   [IsString](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#IsString)
-   [KebabCase](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#KebabCase)
-   [LowerFirst](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#LowerFirst)
-   [UpperFirst](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#UpperFirst)
-   [PadEnd](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#PadEnd)
-   [PadStart](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#PadStart)
-   [Reverse](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Reverse)
-   [SnakeCase](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#SnakeCase)
-   [SplitEx](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#SplitEx)
-   [Wrap](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Wrap)
-   [Unwrap](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Unwrap)

### 19. System package contain some functions about os, runtime, shell command.

```go
import "github.com/duke-git/lancet/v2/system"
```

#### Function list:

-   [IsWindows](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsWindows)
-   [IsLinux](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsLinux)
-   [IsMac](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsMac)
-   [GetOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#GetOsEnv)
-   [SetOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#SetOsEnv)
-   [RemoveOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#RemoveOsEnv)
-   [CompareOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#CompareOsEnv)
-   [ExecCommand](https://github.com/duke-git/lancet/blob/main/docs/system.md#ExecCommand)
-   [GetOsBits](https://github.com/duke-git/lancet/blob/main/docs/system.md#GetOsBits)

### 19. Validator package contains some functions for data validation.

```go
import "github.com/duke-git/lancet/v2/validator"
```

#### Function list:

-   [ContainChinese](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainChinese)
-   [ContainLetter](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainLetter)
-   [ContainLower](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainLower)
-   [ContainUpper](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainUpper)
-   [IsAlpha](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAlpha)
-   [IsAllUpper](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAllUpper)
-   [IsAllLower](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAllLower)
-   [IsBase64](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsBase64)
-   [IsChineseMobile](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChineseMobile)
-   [IsChineseIdNum](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChineseIdNum)
-   [IsChinesePhone](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChinesePhone)
-   [IsCreditCard](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsCreditCard)
-   [IsDns](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsDns)
-   [IsEmail](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsEmail)
-   [IsEmptyString](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsEmptyString)
-   [IsFloatStr](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsFloatStr)
-   [IsNumberStr](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsNumberStr)
-   [IsJSON](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsJSON)
-   [IsRegexMatch](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsRegexMatch)
-   [IsIntStr](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIntStr)
-   [IsIp](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIp)
-   [IsIpV4](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIpV4)
-   [IsIpV6](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIpV6)
-   [IsStrongPassword](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsStrongPassword)
-   [IsUrl](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsUrl)
-   [IsWeakPassword](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsWeakPassword)
-   [IsZeroValue](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsZeroValue)

### 20. xerror package implements helpers for errors.

```go
import "github.com/duke-git/lancet/v2/xerror"
```

#### Function list:

-   [Unwrap](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#Unwrap)

## How to Contribute

I really appreciate any code commits which make lancet lib powerful. Please follow the rules below to create your pull request.

1. Fork the repository.
2. Create your feature branch.
3. Commit your changes.
4. Push to the branch
5. Create new pull request.
