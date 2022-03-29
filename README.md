<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3D1.16<recommend>-9cf)
[![Release](https://img.shields.io/badge/release-2.0.2-green.svg)](https://github.com/duke-git/lancet/releases)
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

English | [简体中文](./README_zh-CN.md)

## Feature

- 👏 Comprehensive, efficient and reusable.
- 💪 250+ go util functions, support string, slice, datetime, net, crypt...
- 💅 Only depend on the go standard library.
- 🌍 Unit test for every exported function.

## Installation
### Note:

1. <b>For users who use go1.18 and above, it is recommended to install lancet v2.x.x. Cause v2.x.x rewrite all functions with generics of go1.18.</b>
```go
go get github.com/duke-git/lancet/v2 // will install latest version of v2.x.x
```

2. <b>For users who use version below go1.18, you should install v1.x.x. now latest v1 is v1.2.6. </b>
```go
go get github.com/duke-git/lancet@v1.2.6 // below go1.18, install latest version of v1.x.x
```

## Usage

Lancet organizes the code into package structure, and you need to import the corresponding package name when use it. For example, if you use string-related functions,import the strutil package like below:

```go
import "github.com/duke-git/lancet/v2/strutil"
```

## Example

Here takes the string function ReverseStr (reverse order string) as an example, and the strutil package needs to be imported.

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    s := "hello"
    rs := strutil.ReverseStr(s)
    fmt.Println(rs) //olleh
}
```

## API Documentation

### Algorithm package implements some basic algorithm. eg. sort, search.

```go
import "github.com/duke-git/lancet/v2/algorithm"
```
#### Function list:
- [BubbleSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BubbleSort)
- [CountSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#CountSort)
- [HeapSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#HeapSort)
- [InsertionSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#InsertionSort)
- [MergeSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#MergeSort)
- [QuickSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#QuickSort)
- [SelectionSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#SelectionSort)
- [ShellSort](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#ShellSort)
- [BinarySearch](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BinarySearch)
- [BinaryIterativeSearch](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BinaryIterativeSearch)
- [LinearSearch](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#LinearSearch)
- [LRUCache](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#LRUCache)


### Convertor package contains some functions for data convertion.

```go
import "github.com/duke-git/lancet/v2/convertor"
```
#### Function list:
- [ColorHexToRGB](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ColorHexToRGB)
- [ColorRGBToHex](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ColorRGBToHex)
- [ToBool](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToBool)
- [ToBytes](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToBytes)
- [ToChar](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToChar)
- [ToInt](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToInt)
- [ToJson](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToJson)
- [ToString](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToString)
- [StructToMap](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#StructToMap)
  
### Cryptor package is for data encryption and decryption.

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### Function list:
- [AesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesEcbEncrypt)
- [AesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesEcbDecrypt)
- [AesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCbcEncrypt)
- [AesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCbcDecrypt)
- [AesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCtrCrypt)
- [AesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCfbEncrypt)
- [AesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCfbDecrypt)
- [AesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesOfbEncrypt)
- [AesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesOfbDecrypt)
- [Base64StdEncode](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Base64StdEncode)
- [Base64StdDecode](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Base64StdDecode)
- [DesEcbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesEcbEncrypt)
- [DesEcbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesEcbDecrypt)
- [DesCbcEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCbcEncrypt)
- [DesCbcDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCbcDecrypt)
- [DesCtrCrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCtrCrypt)
- [DesCfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCfbEncrypt)
- [DesCfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCfbDecrypt)
- [DesOfbEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesOfbEncrypt)
- [DesOfbDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesOfbDecrypt)
- [HmacMd5](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacMd5)
- [HmacSha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha1)
- [HmacSha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha256)
- [HmacSha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha512)
- [Md5String](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Md5String)
- [Md5File](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Md5File)
- [Sha1](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha1)
- [Sha256](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha256)
- [Sha512](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha512)
- [GenerateRsaKey](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#GenerateRsaKey)
- [RsaEncrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#RsaEncrypt)
- [RsaDecrypt](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#RsaDecrypt)

### Datetime package supports date and time format and compare.


```go
import "github.com/duke-git/lancet/v2/datetime"
```
#### Function list:
- [AddDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddDay)
- [AddHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddHour)
- [AddMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddMinute)
- [BeginOfMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMinute)
- [BeginOfHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfHour)
- [BeginOfDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfDay)
- [BeginOfWeek](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfWeek)
- [BeginOfMonth](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMonth)
- [BeginOfYear](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfYear)
- [EndOfMinute](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMinute)
- [EndOfHour](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfHour)
- [EndOfDay](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfDay)
- [EndOfWeek](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfWeek)
- [EndOfMonth](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMonth)
- [EndOfYear](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfYear)
- [GetNowDate](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowDate)
- [GetNowTime](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowTime)
- [GetNowDateTime](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowDateTime)
- [GetZeroHourTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetZeroHourTimestamp)
- [GetNightTimestamp](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNightTimestamp)
- [FormatTimeToStr](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#FormatTimeToStr)
- [FormatStrToTime](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#FormatStrToTime)
- [NewUnix](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewUnix)
- [NewUnixNow](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewUnixNow)
- [NewFormat](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewFormat)
- [NewISO8601](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewISO8601)
- [ToUnix](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToUnix)
- [ToFormat](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToFormat)
- [ToFormatForTpl](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToFormatForTpl)
- [ToIso8601](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToIso8601)

### Fileutil package implements some basic functions for file operations.

```go
import "github.com/duke-git/lancet/v2/fileutil"
```

#### Function list：

- [ClearFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ClearFile)
- [CreateFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CreateFile)
- [CopyFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CopyFile)
- [FileMode](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#FileMode)
- [MiMeType](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#MiMeType)
- [IsExist](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsExist)
- [IsLink](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsLink)
- [IsDir](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsDir)
- [ListFileNames](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ListFileNames)
- [RemoveFile](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#RemoveFile)
- [ReadFileToString](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ReadFileToString)
- [ReadFileByLine](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ReadFileByLine)
- [Zip](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#Zip)
- [UnZip](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#UnZip)

### Formatter contains some functions for data formatting.

```go
import "github.com/duke-git/lancet/v2/formatter"
```
#### Function list:
- [Comma](https://github.com/duke-git/lancet/blob/main/docs/formatter.md#Comma)

### Function package can control the flow of function execution and support part of functional programming

```go
import "github.com/duke-git/lancet/v2/function"
```

#### Function list:
- [After](https://github.com/duke-git/lancet/blob/main/docs/function.md#After)
- [Before](https://github.com/duke-git/lancet/blob/main/docs/function.md#Before)
- [Curry](https://github.com/duke-git/lancet/blob/main/docs/function.md#Curry)
- [Compose](https://github.com/duke-git/lancet/blob/main/docs/function.md#Compose)
- [Debounced](https://github.com/duke-git/lancet/blob/main/docs/function.md#Debounced)
- [Delay](https://github.com/duke-git/lancet/blob/main/docs/function.md#Delay)
- [Watcher](https://github.com/duke-git/lancet/blob/main/docs/function.md#Watcher)


### Mathutil package implements some functions for math calculation.

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### Function list:
- [Average](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Average)
- [Exponent](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Exponent)
- [Fibonacci](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Fibonacci)
- [Factorial](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Factorial)
- [Max](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Max)
- [Min](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Min)
- [Percent](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Percent)
- [RoundToFloat](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#RoundToFloat)
- [RoundToString](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#RoundToString)
- [TruncRound](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#TruncRound)


### Netutil package contains functions to get net information and send http request.

```go
import "github.com/duke-git/lancet/v2/netutil"
```

#### Function list:
- [ConvertMapToQueryString](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#ConvertMapToQueryString)
- [GetInternalIp](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetInternalIp)
- [GetIps](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetIps)
- [GetMacAddrs](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetMacAddrs)
- [GetPublicIpInfo](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetPublicIpInfo)
- [IsPublicIP](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#IsPublicIP)
- [HttpGet](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpGet)
- [HttpDelete](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpDelete)
- [HttpPost](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPost)
- [HttpPut](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPut)
- [HttpPatch](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPatch)
- [ParseHttpResponse](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#ParseHttpResponse)

### Random package implements some basic functions to generate random int and string.

```go
import "github.com/duke-git/lancet/v2/random"
```

#### Function list:
- [RandBytes](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandBytes)
- [RandInt](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandInt)
- [RandString](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandString)
- [UUIdV4](https://github.com/duke-git/lancet/blob/main/docs/random.md#UUIdV4)

### Retry package is for executing a function repeatedly until it was successful or canceled by the context.

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### Function list:
- [Context](https://github.com/duke-git/lancet/blob/main/docs/retry.md#Context)
- [Retry](https://github.com/duke-git/lancet/blob/main/docs/retry.md#Retry)
- [RetryFunc](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryFunc)
- [RetryDuration](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryDuration)
- [RetryTimes](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryTimes)

### Slice contains some functions to manipulate slice.

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### Function list:
- [Contain](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Contain)
- [ContainSubSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ContainSubSlice)
- [Chunk](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Chunk)
- [Compact](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Compact)
- [Concat](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Concat)
- [Count](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Count)
- [Difference](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Difference)
- [DifferenceBy](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DifferenceBy)
- [DifferenceWith](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DifferenceWith)
- [DeleteAt](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DeleteAt)
- [Drop](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Drop)
- [Every](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Every)
- [Filter](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Filter)
- [Find](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Find)
- [FindLast](https://github.com/duke-git/lancet/blob/main/docs/slice.md#FindLast)
- [FlattenDeep](#FlattenDeep)
- [ForEach](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ForEach)
- [GroupBy](https://github.com/duke-git/lancet/blob/main/docs/slice.md#GroupBy)
- [GroupWith](https://github.com/duke-git/lancet/blob/main/docs/slice.md#GroupWith)
- [IntSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IntSlice)
- [InterfaceSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#InterfaceSlice)
- [Intersection](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Intersection)
- [InsertAt](https://github.com/duke-git/lancet/blob/main/docs/slice.md#InsertAt)
- [Map](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Map)
- [Reverse](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Reverse)
- [Reduce](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Reduce)
- [Shuffle](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Shuffle)
- [SortByField](https://github.com/duke-git/lancet/blob/main/docs/slice.md#SortByField)
- [Some](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Some)
- [StringSlice](https://github.com/duke-git/lancet/blob/main/docs/slice.md#StringSlice)
- [Unique](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Unique)
- [Union](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Union)
- [UpdateAt](https://github.com/duke-git/lancet/blob/main/docs/slice.md#UpdateAt)
- [Without](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Without)

### Strutil package contains some functions to manipulate string.
```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### Function list:

- [After](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#After)
- [AfterLast](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#AfterLast)
- [Before](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Before)
- [BeforeLast](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#BeforeLast)
- [CamelCase](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#CamelCase)
- [Capitalize](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Capitalize)
- [IsString](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#IsString)
- [KebabCase](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#KebabCase)
- [LowerFirst](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#LowerFirst)
- [UpperFirst](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#UpperFirst)
- [PadEnd](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#PadEnd)
- [PadStart](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#PadStart)
- [ReverseStr](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#ReverseStr)
- [SnakeCase](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#SnakeCase)
- [Wrap](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Wrap)
- [Unwrap](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Unwrap)
  
### System package contain some functions about os, runtime, shell command.

```go
import "github.com/duke-git/lancet/v2/system"
```
#### Function list:
- [IsWindows](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsWindows)
- [IsLinux](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsLinux)
- [IsMac](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsMac)
- [GetOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#GetOsEnv)
- [SetOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#SetOsEnv)
- [RemoveOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#RemoveOsEnv)
- [CompareOsEnv](https://github.com/duke-git/lancet/blob/main/docs/system.md#CompareOsEnv)
- [ExecCommand](https://github.com/duke-git/lancet/blob/main/docs/system.md#ExecCommand)

### Validator package contains some functions for data validation.

```go
import "github.com/duke-git/lancet/v2/validator"
```

#### Function list:

- [ContainChinese](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainChinese)
- [ContainLetter](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainLetter)
- [ContainLower](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainLower)
- [ContainUpper](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainUpper)
- [IsAlpha](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAlpha)
- [IsAllUpper](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAllUpper)
- [IsAllLower](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAllLower)
- [IsBase64](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsBase64)
- [IsChineseMobile](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChineseMobile)
- [IsChineseIdNum](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChineseIdNum)
- [IsChinesePhone](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChinesePhone)
- [IsCreditCard](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsCreditCard)
- [IsDns](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsDns)
- [IsEmail](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsEmail)
- [IsEmptyString](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsEmptyString)
- [IsFloatStr](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsFloatStr)
- [IsNumberStr](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsNumberStr)
- [IsJSON](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsJSON)
- [IsRegexMatch](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsRegexMatch)
- [IsIntStr](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIntStr)
- [IsIp](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIp)
- [IsIpV4](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIpV4)
- [IsIpV6](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIpV6)
- [IsStrongPassword](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsStrongPassword)
- [IsUrl](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsUrl)
- [IsWeakPassword](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsWeakPassword)
### xerror package implements helpers for errors.

```go
import "github.com/duke-git/lancet/v2/xerror"
```
#### Function list:
- [Unwrap](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#Unwrap)
  

## How to Contribute

I really appreciate any code commits which make lancet lib powerful. Please follow the rules below to create your pull request.

1. Fork the repository.
2. Create your feature branch.
3. Commit your changes.
4. Push to the branch
5. Create new pull request.
