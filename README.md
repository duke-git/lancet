<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-v1.16-9cf)
[![Release](https://img.shields.io/badge/release-1.3.4-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com//duke-git/lancet?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet)
[![Go Report Card](https://goreportcard.com/badge/github.com/duke-git/lancet)](https://goreportcard.com/report/github.com/duke-git/lancet)
[![test](https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/duke-git/lancet/actions/workflows/codecov.yml)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/v1/LICENSE)

</div>

<div STYLE="page-break-after: always;"></div>

<p style="font-size: 20px"> 
    Lancet is a comprehensive, efficient, and reusable util function library of go. Inspired by the java apache common package and lodash.js.
</p>

English | [简体中文](./README_zh-CN.md)

## Feature

-   👏 Comprehensive, efficient and reusable.
-   💪 200+ go util functions, support string, slice, datetime, net, crypt...
-   💅 Only depend on the go standard library.
-   🌍 Unit test for every exported function.

## Installation

```go
go get github.com/duke-git/lancet
```

## Usage

Lancet organizes the code into package structure, and you need to import the corresponding package name when use it. For example, if you use string-related functions,import the strutil package like below:

```go
import "github.com/duke-git/lancet/strutil"
```

## Example

Here takes the string function Reverse (reverse order string) as an example, and the strutil package needs to be imported.

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    s := "hello"
    rs := strutil.Reverse(s)
    fmt.Println(rs) //olleh
}
```

## API Documentation

### 1. Convertor package contains some functions for data convertion.

```go
import "github.com/duke-git/lancet/convertor"
```

#### Function list:

-   [ColorHexToRGB](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ColorHexToRGB)
-   [ColorRGBToHex](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ColorRGBToHex)
-   [ToBool](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ToBool)
-   [ToBytes](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ToBytes)
-   [ToChar](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ToChar)
-   [ToChannel](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ToChannel)
-   [ToInt](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ToInt)
-   [ToJson](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ToJson)
-   [ToString](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#ToString)
-   [StructToMap](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#StructToMap)
-   [EncodeByte](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#EncodeByte)
-   [DecodeByte](https://github.com/duke-git/lancet/blob/v1/docs/convertor.md#DecodeByte)

### 2. Cryptor package is for data encryption and decryption.

```go
import "github.com/duke-git/lancet/cryptor"
```

#### Function list:

-   [AesEcbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesEcbEncrypt)
-   [AesEcbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesEcbDecrypt)
-   [AesCbcEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesCbcEncrypt)
-   [AesCbcDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesCbcDecrypt)
-   [AesCtrCrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesCtrCrypt)
-   [AesCfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesCfbEncrypt)
-   [AesCfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesCfbDecrypt)
-   [AesOfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesOfbEncrypt)
-   [AesOfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#AesOfbDecrypt)
-   [Base64StdEncode](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#Base64StdEncode)
-   [Base64StdDecode](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#Base64StdDecode)
-   [DesEcbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesEcbEncrypt)
-   [DesEcbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesEcbDecrypt)
-   [DesCbcEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesCbcEncrypt)
-   [DesCbcDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesCbcDecrypt)
-   [DesCtrCrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesCtrCrypt)
-   [DesCfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesCfbEncrypt)
-   [DesCfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesCfbDecrypt)
-   [DesOfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesOfbEncrypt)
-   [DesOfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#DesOfbDecrypt)
-   [HmacMd5](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#HmacMd5)
-   [HmacSha1](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#HmacSha1)
-   [HmacSha256](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#HmacSha256)
-   [HmacSha512](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#HmacSha512)
-   [Md5String](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#Md5String)
-   [Md5File](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#Md5File)
-   [Sha1](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#Sha1)
-   [Sha256](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#Sha256)
-   [Sha512](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#Sha512)
-   [GenerateRsaKey](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#GenerateRsaKey)
-   [RsaEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#RsaEncrypt)
-   [RsaDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor.md#RsaDecrypt)

### 3. Datetime package supports date and time format and compare.

```go
import "github.com/duke-git/lancet/datetime"
```

#### Function list:

-   [AddDay](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#AddDay)
-   [AddHour](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#AddHour)
-   [AddMinute](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#AddMinute)
-   [BeginOfMinute](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#BeginOfMinute)
-   [BeginOfHour](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#BeginOfHour)
-   [BeginOfDay](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#BeginOfDay)
-   [BeginOfWeek](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#BeginOfWeek)
-   [BeginOfMonth](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#BeginOfMonth)
-   [BeginOfYear](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#BeginOfYear)
-   [EndOfMinute](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#EndOfMinute)
-   [EndOfHour](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#EndOfHour)
-   [EndOfDay](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#EndOfDay)
-   [EndOfWeek](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#EndOfWeek)
-   [EndOfMonth](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#EndOfMonth)
-   [EndOfYear](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#EndOfYear)
-   [GetNowDate](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#GetNowDate)
-   [GetNowTime](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#GetNowTime)
-   [GetNowDateTime](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#GetNowDateTime)
-   [GetZeroHourTimestamp](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#GetZeroHourTimestamp)
-   [GetNightTimestamp](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#GetNightTimestamp)
-   [FormatTimeToStr](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#FormatTimeToStr)
-   [FormatStrToTime](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#FormatStrToTime)
-   [NewUnix](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#NewUnix)
-   [NewUnixNow](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#NewUnixNow)
-   [NewFormat](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#NewFormat)
-   [NewISO8601](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#NewISO8601)
-   [ToUnix](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#ToUnix)
-   [ToFormat](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#ToFormat)
-   [ToFormatForTpl](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#ToFormatForTpl)
-   [ToIso8601](https://github.com/duke-git/lancet/blob/v1/docs/datetime.md#ToIso8601)

### 4. Fileutil package implements some basic functions for file operations.

```go
import "github.com/duke-git/lancet/fileutil"
```

#### Function list：

-   [ClearFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#ClearFile)
-   [CreateFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#CreateFile)
-   [CreateDir](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#CreateDir)
-   [CopyFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#CopyFile)
-   [FileMode](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#FileMode)
-   [MiMeType](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#MiMeType)
-   [IsExist](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#IsExist)
-   [IsLink](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#IsLink)
-   [IsDir](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#IsDir)
-   [ListFileNames](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#ListFileNames)
-   [RemoveFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#RemoveFile)
-   [ReadFileToString](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#ReadFileToString)
-   [ReadFileByLine](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#ReadFileByLine)
-   [Zip](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#Zip)
-   [UnZip](https://github.com/duke-git/lancet/blob/v1/docs/fileutil.md#UnZip)

### 5. Formatter contains some functions for data formatting.

```go
import "github.com/duke-git/lancet/formatter"
```

#### Function list:

-   [Comma](https://github.com/duke-git/lancet/blob/v1/docs/formatter.md#Comma)

### Function package can control the flow of function execution and support part of functional programming

```go
import "github.com/duke-git/lancet/function"
```

#### Function list:

-   [After](https://github.com/duke-git/lancet/blob/v1/docs/function.md#After)
-   [Before](https://github.com/duke-git/lancet/blob/v1/docs/function.md#Before)
-   [Curry](https://github.com/duke-git/lancet/blob/v1/docs/function.md#Curry)
-   [Compose](https://github.com/duke-git/lancet/blob/v1/docs/function.md#Compose)
-   [Debounced](https://github.com/duke-git/lancet/blob/v1/docs/function.md#Debounced)
-   [Delay](https://github.com/duke-git/lancet/blob/v1/docs/function.md#Delay)
-   [Watcher](https://github.com/duke-git/lancet/blob/v1/docs/function.md#Watcher)

### 6. Mathutil package implements some functions for math calculation.

```go
import "github.com/duke-git/lancet/mathutil"
```

#### Function list:

-   [Exponent](https://github.com/duke-git/lancet/blob/v1/docs/mathutil.md#Exponent)
-   [Fibonacci](https://github.com/duke-git/lancet/blob/v1/docs/mathutil.md#Fibonacci)
-   [Factorial](https://github.com/duke-git/lancet/blob/v1/docs/mathutil.md#Factorial)
-   [Percent](https://github.com/duke-git/lancet/blob/v1/docs/mathutil.md#Percent)
-   [RoundToFloat](https://github.com/duke-git/lancet/blob/v1/docs/mathutil.md#RoundToFloat)
-   [RoundToString](https://github.com/duke-git/lancet/blob/v1/docs/mathutil.md#RoundToString)
-   [TruncRound](https://github.com/duke-git/lancet/blob/v1/docs/mathutil.md#TruncRound)

### 7. Netutil package contains functions to get net information and send http request.

```go
import "github.com/duke-git/lancet/netutil"
```

#### Function list:

-   [ConvertMapToQueryString](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#ConvertMapToQueryString)
-   [EncodeUrl](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#EncodeUrl)
-   [GetInternalIp](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#GetInternalIp)
-   [GetIps](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#GetIps)
-   [GetMacAddrs](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#GetMacAddrs)
-   [GetPublicIpInfo](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#GetPublicIpInfo)
-   [GetRequestPublicIp](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#GetRequestPublicIp)
-   [IsPublicIP](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#IsPublicIP)
-   [IsInternalIP](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#IsInternalIP)
-   [HttpGet](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#HttpGet)
-   [HttpDelete](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#HttpDelete)
-   [HttpPost](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#HttpPost)
-   [HttpPut](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#HttpPut)
-   [HttpPatch](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#HttpPatch)
-   [ParseHttpResponse](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#ParseHttpResponse)

### 8. Random package implements some basic functions to generate random int and string.

```go
import "github.com/duke-git/lancet/random"
```

#### Function list:

-   [RandBytes](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandBytes)
-   [RandInt](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandInt)
-   [RandString](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandString)
-   [RandUpper](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandUpper)
-   [RandLower](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandLower)
-   [RandNumeral](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandNumeral)
-   [RandNumeralOrLetter](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandNumeralOrLetter)
-   [UUIdV4](https://github.com/duke-git/lancet/blob/v1/docs/random.md#UUIdV4)

### 9. Retry package is for executing a function repeatedly until it was successful or canceled by the context.

```go
import "github.com/duke-git/lancet/retry"
```

#### Function list:

-   [Context](https://github.com/duke-git/lancet/blob/v1/docs/retry.md#Context)
-   [Retry](https://github.com/duke-git/lancet/blob/v1/docs/retry.md#Retry)
-   [RetryFunc](https://github.com/duke-git/lancet/blob/v1/docs/retry.md#RetryFunc)
-   [RetryDuration](https://github.com/duke-git/lancet/blob/v1/docs/retry.md#RetryDuration)
-   [RetryTimes](https://github.com/duke-git/lancet/blob/v1/docs/retry.md#RetryTimes)

### 10. Slice contains some functions to manipulate slice.

```go
import "github.com/duke-git/lancet/slice"
```

#### Function list:

-   [AppendIfAbsent](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#AppendIfAbsent)
-   [Contain](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Contain)
-   [ContainSubSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#ContainSubSlice)
-   [Chunk](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Chunk)
-   [Compact](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Compact)
-   [Concat](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Concat)
-   [Count](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Count)
-   [Difference](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Difference)
-   [DifferenceBy](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#DifferenceBy)
-   [DeleteByIndex](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#DeleteByIndex)
-   [Drop](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Drop)
-   [Every](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Every)
-   [Equal](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Equal)
-   [EqualWith](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#EqualWith)
-   [Filter](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Filter)
-   [Find](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Find)
-   [FindLast](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#FindLast)
-   [FlattenDeep](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#FlattenDeep)
-   [ForEach](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#ForEach)
-   [GroupBy](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#GroupBy)
-   [IntSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#IntSlice)
-   [IndexOf](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#IndexOf)
-   [LastIndexOf](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#LastIndexOf)
-   [InterfaceSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#InterfaceSlice)
-   [Intersection](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Intersection)
-   [InsertByIndex](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#InsertByIndex)
-   [Map](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Map)
-   [ReverseSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#ReverseSlice)
-   [Reduce](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Reduce)
-   [Shuffle](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Shuffle)
-   [SortByField](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#SortByField)
-   [Some](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Some)
-   [StringSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#StringSlice)
-   [ToSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#ToSlice)
-   [ToSlicePointer](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#ToSlice)
-   [Unique](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Unique)
-   [UniqueBy](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#UniqueBy)
-   [Union](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Union)
-   [UpdateByIndex](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#UpdateByIndex)
-   [Without](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#Without)

### 11. Strutil package contains some functions to manipulate string.

```go
import "github.com/duke-git/lancet/strutil"
```

#### Function list:

-   [After](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#After)
-   [AfterLast](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#AfterLast)
-   [Before](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#Before)
-   [BeforeLast](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#BeforeLast)
-   [CamelCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#CamelCase)
-   [Capitalize](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#Capitalize)
-   [IsString](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#IsString)
-   [KebabCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#KebabCase)
-   [LowerFirst](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#LowerFirst)
-   [UpperFirst](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#UpperFirst)
-   [PadEnd](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#PadEnd)
-   [PadStart](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#PadStart)
-   [Reverse](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#Reverse)
-   [SnakeCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#SnakeCase)
-   [SplitEx](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#SplitEx)
-   [Wrap](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#Wrap)
-   [Unwrap](https://github.com/duke-git/lancet/blob/v1/docs/strutil.md#Unwrap)

### 12. System package contain some functions about os, runtime, shell command.

```go
import "github.com/duke-git/lancet/system"
```

#### Function list:

-   [IsWindows](https://github.com/duke-git/lancet/blob/v1/docs/system.md#IsWindows)
-   [IsLinux](https://github.com/duke-git/lancet/blob/v1/docs/system.md#IsLinux)
-   [IsMac](https://github.com/duke-git/lancet/blob/v1/docs/system.md#IsMac)
-   [GetOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system.md#GetOsEnv)
-   [SetOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system.md#SetOsEnv)
-   [RemoveOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system.md#RemoveOsEnv)
-   [CompareOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system.md#CompareOsEnv)
-   [ExecCommand](https://github.com/duke-git/lancet/blob/v1/docs/system.md#ExecCommand)
-   [GetOsBits](https://github.com/duke-git/lancet/blob/v1/docs/system.md#GetOsBits)

### 13. Validator package contains some functions for data validation.

```go
import "github.com/duke-git/lancet/validator"
```

#### Function list:

-   [ContainChinese](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#ContainChinese)
-   [ContainLetter](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#ContainLetter)
-   [ContainLower](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#ContainLower)
-   [ContainUpper](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#ContainUpper)
-   [IsAlpha](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsAlpha)
-   [IsAllUpper](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsAllUpper)
-   [IsAllLower](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsAllLower)
-   [IsBase64](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsBase64)
-   [IsChineseMobile](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsChineseMobile)
-   [IsChineseIdNum](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsChineseIdNum)
-   [IsChinesePhone](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsChinesePhone)
-   [IsCreditCard](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsCreditCard)
-   [IsDns](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsDns)
-   [IsEmail](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsEmail)
-   [IsEmptyString](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsEmptyString)
-   [IsFloatStr](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsFloatStr)
-   [IsNumberStr](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsNumberStr)
-   [IsJSON](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsJSON)
-   [IsRegexMatch](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsRegexMatch)
-   [IsIntStr](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsIntStr)
-   [IsIp](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsIp)
-   [IsIpV4](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsIpV4)
-   [IsIpV6](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsIpV6)
-   [IsStrongPassword](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsStrongPassword)
-   [IsUrl](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsUrl)
-   [IsWeakPassword](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsWeakPassword)
-   [IsZeroValue](https://github.com/duke-git/lancet/blob/v1/docs/validator.md#IsZeroValue)

## How to Contribute

I really appreciate any code commits which make lancet lib powerful. Please follow the rules below to create your pull request.

1. Fork the repository.
2. Create your feature branch.
3. Commit your changes.
4. Push to the branch
5. Create new pull request.
