<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-v1.16-9cf)
[![Release](https://img.shields.io/badge/release-1.4.0-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com//duke-git/lancet?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet)
[![Go Report Card](https://goreportcard.com/badge/github.com/duke-git/lancet)](https://goreportcard.com/report/github.com/duke-git/lancet)
[![test](https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/duke-git/lancet/actions/workflows/codecov.yml)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/v1/LICENSE)

</div>

<div STYLE="page-break-after: always;"></div>
<p style="font-size: 18px">
    lancet（柳叶刀）是一个全面、高效、可复用的go语言工具函数库。 lancet受到了java apache common包和lodash.js的启发。
</p>

简体中文 | [English](./README.md)

## 特性

-   👏 全面、高效、可复用
-   💪 400+常用 go 工具函数，支持 string、slice、datetime、net、crypt...
-   💅 只依赖 go 标准库
-   🌍 所有导出函数单元测试覆盖率 100%

## 安装

```go
go get github.com/duke-git/lancet
```

## 用法

lancet 是以包的结构组织代码的，使用时需要导入相应的包名。例如：如果使用字符串相关函数，需要导入 strutil 包:

```go
import "github.com/duke-git/lancet/strutil"
```

## 例子

此处以字符串工具函数 Reverse（逆序字符串）为例，需要导入 strutil 包:

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

## API 文档


### 1. compare包提供几个轻量级的类型比较函数。
```go
import "github.com/duke-git/lancet/compare"
```

#### Function list:

-   [Equal](https://github.com/duke-git/lancet/blob/v1/docs/compare_zh-CN.md#Equal)
-   [EqualValue](https://github.com/duke-git/lancet/blob/v1/docs/compare_zh-CN.md#EqualValue)
-   [LessThan](https://github.com/duke-git/lancet/blob/v1/docs/compare_zh-CN.md#LessThan)
-   [GreaterThan](https://github.com/duke-git/lancet/blob/v1/docs/compare_zh-CN.md#GreaterThan)
-   [LessOrEqual](https://github.com/duke-git/lancet/blob/v1/docs/compare_zh-CN.md#LessOrEqual)
-   [GreaterOrEqual](https://github.com/duke-git/lancet/blob/v1/docs/compare_zh-CN.md#GreaterOrEqual)

### 2. convertor转换器包支持一些常见的数据类型转换。

```go
import "github.com/duke-git/lancet/convertor"
```

#### 函数列表:

-   [ColorHexToRGB](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ColorHexToRGB)
-   [ColorRGBToHex](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ColorRGBToHex)
-   [ToBool](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToBool)
-   [ToBytes](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToBytes)
-   [ToChar](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToChar)
-   [ToChannel](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToChannel)
-   [ToInt](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToInt)
-   [ToJson](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToJson)
-   [ToString](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToString)
-   [StructToMap](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#StructToMap)
-   [EncodeByte](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#EncodeByte)
-   [DecodeByte](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#DecodeByte)
-   [DeepClone](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#DeepClone)
-   [CopyProperties](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#CopyProperties)
-   [ToInterface](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#ToInterface)
-   [Utf8ToGbk](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#Utf8ToGbk)
-   [GbkToUtf8](https://github.com/duke-git/lancet/blob/v1/docs/convertor_zh-CN.md#GbkToUtf8)
### 3. cryptor 加密包支持数据加密和解密，获取 md5，hash 值。支持 base64, md5, hmac, aes, des, rsa。

```go
import "github.com/duke-git/lancet/cryptor"
```

#### 函数列表:

-   [AesEcbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesEcbEncrypt)
-   [AesEcbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesEcbDecrypt)
-   [AesCbcEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesCbcEncrypt)
-   [AesCbcDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesCbcDecrypt)
-   [AesCtrCrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesCtrCrypt)
-   [AesCfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesCfbEncrypt)
-   [AesCfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesCfbDecrypt)
-   [AesOfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesOfbEncrypt)
-   [AesOfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#AesOfbDecrypt)
-   [Base64StdEncode](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Base64StdEncode)
-   [Base64StdDecode](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Base64StdDecode)
-   [DesEcbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesEcbEncrypt)
-   [DesEcbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesEcbDecrypt)
-   [DesCbcEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesCbcEncrypt)
-   [DesCbcDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesCbcDecrypt)
-   [DesCtrCrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesCtrCrypt)
-   [DesCfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesCfbEncrypt)
-   [DesCfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesCfbDecrypt)
-   [DesOfbEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesOfbEncrypt)
-   [DesOfbDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#DesOfbDecrypt)
-   [HmacMd5](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacMd5)
-   [HmacMd5WithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacMd5WithBase64)
-   [HmacSha1WithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacSha1WithBase64)
-   [HmacSha1](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacSha1)
-   [HmacSha256WithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacSha256WithBase64)
-   [HmacSha1](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacSha1)
-   [HmacSha512WithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacSha512WithBase64)
-   [HmacSha1](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#HmacSha1)
-   [Md5String](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Md5String)
-   [Md5StringWithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Md5StringWithBase64)
-   [Md5Byte](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Md5Byte)
-   [Md5ByteWithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Md5ByteWithBase64)
-   [Md5File](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Md5File)
-   [Sha1](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Sha1)
-   [Sha1WithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Sha1WithBase64)
-   [Sha256](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Sha256)
-   [Sha256WithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Sha256WithBase64)
-   [Sha512](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Sha512)
-   [Sha512WithBase64](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#Sha512WithBase64)
-   [GenerateRsaKey](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#GenerateRsaKey)
-   [RsaEncrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#RsaEncrypt)
-   [RsaDecrypt](https://github.com/duke-git/lancet/blob/v1/docs/cryptor_zh-CN.md#RsaDecrypt)

### 4. datetime 日期时间处理包，格式化日期，比较日期。

```go
import "github.com/duke-git/lancet/datetime"
```

#### 函数列表:

-   [AddDay](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#AddDay)
-   [AddHour](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#AddHour)
-   [AddMinute](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#AddMinute)
-   [AddYear](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#AddYear)
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
-   [GetNowDate](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#GetNowDate)
-   [GetNowTime](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#GetNowTime)
-   [GetNowDateTime](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#GetNowDateTime)
-   [GetZeroHourTimestamp](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#GetZeroHourTimestamp)
-   [GetNightTimestamp](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#GetNightTimestamp)
-   [FormatTimeToStr](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#FormatTimeToStr)
-   [FormatStrToTime](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#FormatStrToTime)
-   [NewUnix](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#NewUnix)
-   [NewUnixNow](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#NewUnixNow)
-   [NewFormat](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#NewFormat)
-   [NewISO8601](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#NewISO8601)
-   [ToUnix](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#ToUnix)
-   [ToFormat](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#ToFormat)
-   [ToFormatForTpl](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#ToFormatForTpl)
-   [ToIso8601](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#ToIso8601)
-   [IsLeapYear](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#IsLeapYear)
-   [BetweenSeconds](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#BetweenSeconds)
-   [DayOfYear](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#DayOfYear)
-   [IsWeekend](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#IsWeekend)
-   [NowDateOrTime](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#NowDateOrTime)
-   [Timestamp](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#Timestamp)
-   [TimestampMilli](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#TimestampMilli)
-   [TimestampMicro](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#TimestampMicro)
-   [TimestampNano](https://github.com/duke-git/lancet/blob/v1/docs/datetime_zh-CN.md#TimestampNano)

### 5. fileutil 包支持文件基本操作。

```go
import "github.com/duke-git/lancet/fileutil"
```

#### 函数列表：

-   [ClearFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#ClearFile)
-   [CreateFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#CreateFile)
-   [CreateDir](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#CreateDir)
-   [CopyFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#CopyFile)
-   [FileMode](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#FileMode)
-   [MiMeType](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#MiMeType)
-   [IsExist](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#IsExist)
-   [IsLink](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#IsLink)
-   [IsDir](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#IsDir)
-   [ListFileNames](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#ListFileNames)
-   [RemoveFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#RemoveFile)
-   [ReadFileToString](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#ReadFileToString)
-   [ReadFileByLine](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#ReadFileByLine)
-   [Zip](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#Zip)
-   [ZipAppendEntry](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#ZipAppendEntry)
-   [UnZip](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#UnZip)
-   [CurrentPath](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#CurrentPath)
-   [IsZipFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#IsZipFile)
-   [FileSize](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#FileSize)
-   [MTime](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#MTime)
-   [Sha](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#Sha)
-   [ReadCsvFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#ReadCsvFile)
-   [WriteCsvFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#WriteCsvFile)
-   [WriteStringToFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#WriteStringToFile)
-   [WriteBytesToFile](https://github.com/duke-git/lancet/blob/v1/docs/fileutil_zh-CN.md#WriteBytesToFile)

### 6. formatter 格式化器包含一些数据格式化处理方法。

```go
import "github.com/duke-git/lancet/formatter"
```

#### 函数列表:

-   [Comma](https://github.com/duke-git/lancet/blob/v1/docs/formatter_zh-CN.md#Comma)
-   [Pretty](https://github.com/duke-git/lancet/blob/v1/docs/formatter_zh-CN.md#Pretty)
-   [PrettyToWriter](https://github.com/duke-git/lancet/blob/v1/docs/formatter_zh-CN.md#PrettyToWriter)
-   [DecimalBytes](https://github.com/duke-git/lancet/blob/v1/docs/formatter_zh-CN.md#DecimalBytes)
-   [BinaryBytes](https://github.com/duke-git/lancet/blob/v1/docs/formatter_zh-CN.md#BinaryBytes)
-   [ParseDecimalBytes](https://github.com/duke-git/lancet/blob/v1/docs/formatter_zh-CN.md#ParseDecimalBytes)
-   [ParseBinaryBytes](https://github.com/duke-git/lancet/blob/v1/docs/formatter_zh-CN.md#ParseBinaryBytes)

### 7. function 函数包控制函数执行流程，包含部分函数式编程。

```go
import "github.com/duke-git/lancet/function"
```

#### 函数列表:

-   [After](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#After)
-   [Before](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Before)
-   [Curry](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Curry)
-   [Compose](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Compose)
-   [Debounced](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Debounced)
-   [Delay](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Delay)
-   [Pipeline](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Pipeline)
-   [Schedule](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Schedule)
-   [Watcher](https://github.com/duke-git/lancet/blob/v1/docs/function_zh-CN.md#Watcher)

### 8. mathutil 包实现了一些数学计算的函数。

```go
import "github.com/duke-git/lancet/mathutil"
```

#### Function list:

-   [Exponent](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#Exponent)
-   [Fibonacci](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#Fibonacci)
-   [Factorial](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#Factorial)
-   [Percent](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#Percent)
-   [RoundToFloat](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#RoundToFloat)
-   [RoundToString](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#RoundToString)
-   [TruncRound](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#TruncRound)
-   [AngleToRadian](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#AngleToRadian)
-   [RadianToAngle](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#RadianToAngle)
-   [PointDistance](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#PointDistance)
-   [IsPrime](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#IsPrime)
-   [GCD](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#GCD)
-   [LCM](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#LCM)
-   [Cos](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#Cos)
-   [Sin](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#Sin)
-   [Log](https://github.com/duke-git/lancet/blob/v1/docs/mathutil_zh-CN.md#Log)
### 9. netutil 网络包支持获取 ip 地址，发送 http 请求。

```go
import "github.com/duke-git/lancet/netutil"
```

#### 函数列表:

-   [ConvertMapToQueryString](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#ConvertMapToQueryString)
-   [EncodeUrl](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#EncodeUrl)
-   [GetInternalIp](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#GetInternalIp)
-   [GetIps](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#GetIps)
-   [GetMacAddrs](https://github.com/duke-git/lancet/blob/v1/docs/netutil.md#GetMacAddrs)
-   [GetPublicIpInfo](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#GetPublicIpInfo)
-   [GetRequestPublicIp](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#GetRequestPublicIp)
-   [IsPublicIP](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#IsPublicIP)
-   [IsInternalIP](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#IsInternalIP)
-   [HttpGet](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#HttpGet)
-   [HttpDelete](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#HttpDelete)
-   [HttpPost](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#HttpPost)
-   [HttpPut](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#HttpPut)
-   [HttpPatch](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#HttpPatch)
-   [ParseHttpResponse](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#ParseHttpResponse)
-   [UploadFile](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#UploadFile)
-   [DownloadFile](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#DownloadFile)
-   [IsPingConnected](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#IsPingConnected)
-   [IsTelnetConnected](https://github.com/duke-git/lancet/blob/v1/docs/netutil_zh-CN.md#IsTelnetConnected)

### 10. random 随机数生成器包，可以生成随机[]bytes, int, string。

```go
import "github.com/duke-git/lancet/random"
```

#### 函数列表:

-   [RandBytes](https://github.com/duke-git/lancet/blob/v1/docs/random_zh-CN.md#RandBytes)
-   [RandInt](https://github.com/duke-git/lancet/blob/v1/docs/random_zh-CN.md#RandInt)
-   [RandString](https://github.com/duke-git/lancet/blob/v1/docs/random_zh-CN.md#RandString)
-   [RandUpper](https://github.com/duke-git/lancet/blob/v1/docs/random_zh-CN.md#RandUpper)
-   [RandLower](https://github.com/duke-git/lancet/blob/v1/docs/random_zh-CN.md#RandLower)
-   [RandNumeral](https://github.com/duke-git/lancet/blob/v1/docs/random_zh-CN.md#RandNumeral)
-   [RandNumeralOrLetter](https://github.com/duke-git/lancet/blob/v1/docs/random_zh-CN.md#RandNumeralOrLetter)
-   [UUIdV4](https://github.com/duke-git/lancet/blob/v1/docs/random.md#UUIdV4)
-   [RandUniqueIntSlice](https://github.com/duke-git/lancet/blob/v1/docs/random.md#RandUniqueIntSlice)

### 11. retry 重试执行函数直到函数运行成功或被 context cancel。

```go
import "github.com/duke-git/lancet/retry"
```

#### 函数列表:

-   [Context](https://github.com/duke-git/lancet/blob/v1/docs/retry_zh-CN.md#Context)
-   [Retry](https://github.com/duke-git/lancet/blob/v1/docs/retry_zh-CN.md#Retry)
-   [RetryFunc](https://github.com/duke-git/lancet/blob/v1/docs/retry_zh-CN.md#RetryFunc)
-   [RetryDuration](https://github.com/duke-git/lancet/blob/v1/docs/retry_zh-CN.md#RetryDuration)
-   [RetryTimes](https://github.com/duke-git/lancet/blob/v1/docs/retry_zh-CN.md#RetryTimes)

### 12. slice 包包含操作切片的方法集合。

```go
import "github.com/duke-git/lancet/slice"
```

#### 函数列表:

-   [AppendIfAbsent](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#AppendIfAbsent)
-   [Contain](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Contain)
-   [ContainSubSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#ContainSubSlice)
-   [Chunk](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Chunk)
-   [Compact](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Compact)
-   [Concat](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Concat)
-   [Count](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Count)
-   [Difference](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Difference)
-   [DifferenceBy](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#DifferenceBy)
-   [DeleteByIndex](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#DeleteByIndex)
-   [Drop](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Drop)
-   [Every](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Every)
-   [Equal](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Equal)
-   [EqualWith](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#EqualWith)
-   [Filter](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Filter)
-   [Find](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Find)
-   [FindLast](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#FindLast)
-   [FlattenDeep](https://github.com/duke-git/lancet/blob/v1/docs/slice.md#FlattenDeep)
-   [ForEach](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#ForEach)
-   [GroupBy](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#GroupBy)
-   [IntSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#IntSlice)
-   [IndexOf](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#IndexOf)
-   [LastIndexOf](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#LastIndexOf)
-   [InterfaceSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#InterfaceSlice)
-   [Intersection](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Intersection)
-   [InsertByIndex](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#InsertByIndex)
-   [Map](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Map)
-   [ReverseSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#ReverseSlice)
-   [Reduce](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Reduce)
-   [Shuffle](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Shuffle)
-   [SortByField](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#SortByField)
-   [Some](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Some)
-   [StringSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#StringSlice)
-   [ToSlice](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#ToSlice)
-   [ToSlicePointer](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#ToSlice)
-   [Unique](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Unique)
-   [UniqueBy](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#UniqueBy)
-   [Union](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Union)
-   [UpdateByIndex](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#UpdateByIndex)
-   [Without](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Without)
-   [Join](https://github.com/duke-git/lancet/blob/v1/docs/slice_zh-CN.md#Join)

### 13. strutil 包含处理字符串的相关函数。

```go
import "github.com/duke-git/lancet/strutil"
```

#### 函数列表:

-   [After](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#After)
-   [AfterLast](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#AfterLast)
-   [Before](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#Before)
-   [BeforeLast](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#BeforeLast)
-   [CamelCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#CamelCase)
-   [Capitalize](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#Capitalize)
-   [ContainsAll](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#ContainsAll)
-   [ContainsAny](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#ContainsAny)
-   [IsString](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#IsString)
-   [KebabCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#KebabCase)
-   [UpperKebabCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#UpperKebabCase)
-   [LowerFirst](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#LowerFirst)
-   [UpperFirst](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#UpperFirst)
-   [Pad](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#Pad)
-   [PadEnd](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#PadEnd)
-   [PadStart](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#PadStart)
-   [Reverse](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#Reverse)
-   [SnakeCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#SnakeCase)
-   [UpperSnakeCase](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#UpperSnakeCase)
-   [SplitEx](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#SplitEx)
-   [Wrap](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#Wrap)
-   [Unwrap](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#Unwrap)
-   [SplitWords](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#SplitWords)
-   [WordCount](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#WordCount)
-   [RemoveNonPrintable](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#RemoveNonPrintable)
-   [StringToBytes](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#StringToBytes)
-   [BytesToString](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#BytesToString)
-   [IsBlank](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#IsBlank)
-   [HasPrefixAny](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#HasPrefixAny)
-   [HasSuffixAny](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#HasSuffixAny)
-   [IndexOffset](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#IndexOffset)
-   [ReplaceWithMap](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#ReplaceWithMap)
-   [Trim](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#Trim)
-   [SplitAndTrim](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#SplitAndTrim)
-   [HideString](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#HideString)
-   [RemoveWhiteSpace](https://github.com/duke-git/lancet/blob/v1/docs/strutil_zh-CN.md#RemoveWhiteSpace)

### 14. system 包含 os, runtime, shell command 相关函数。

```go
import "github.com/duke-git/lancet/system"
```

#### 函数列表:

-   [IsWindows](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#IsWindows)
-   [IsLinux](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#IsLinux)
-   [IsMac](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#IsMac)
-   [GetOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#GetOsEnv)
-   [SetOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#SetOsEnv)
-   [RemoveOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#RemoveOsEnv)
-   [CompareOsEnv](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#CompareOsEnv)
-   [ExecCommand](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#ExecCommand)
-   [GetOsBits](https://github.com/duke-git/lancet/blob/v1/docs/system_zh-CN.md#GetOsBits)

### 15. validator 验证器包，包含常用字符串格式验证函数。

```go
import "github.com/duke-git/lancet/validator"
```

#### 函数列表:

-   [ContainChinese](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#ContainChinese)
-   [ContainLetter](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#ContainLetter)
-   [ContainLower](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#ContainLower)
-   [ContainUpper](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#ContainUpper)
-   [IsAlpha](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsAlpha)
-   [IsAllUpper](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsAllUpper)
-   [IsAllLower](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsAllLower)
-   [IsBase64](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsBase64)
-   [IsChineseMobile](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsChineseMobile)
-   [IsChineseIdNum](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsChineseIdNum)
-   [IsChinesePhone](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsChinesePhone)
-   [IsCreditCard](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsCreditCard)
-   [IsDns](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsDns)
-   [IsEmail](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsEmail)
-   [IsEmptyString](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsEmptyString)
-   [IsInt](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsInt)
-   [IsFloat](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsFloat)
-   [IsNumber](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsNumber)
-   [IsIntStr](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsIntStr)
-   [IsFloatStr](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsFloatStr)
-   [IsNumberStr](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsNumberStr)
-   [IsJSON](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsJSON)
-   [IsRegexMatch](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsRegexMatch)
-   [IsIp](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsIp)
-   [IsIpV4](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsIpV4)
-   [IsIpV6](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsIpV6)
-   [IsStrongPassword](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsStrongPassword)
-   [IsUrl](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsUrl)
-   [IsWeakPassword](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsWeakPassword)
-   [IsZeroValue](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsZeroValue)
-   [IsGBK](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsGBK)
-   [IsASCII](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsASCII)
-   [IsPrintable](https://github.com/duke-git/lancet/blob/v1/docs/validator_zh-CN.md#IsPrintable)

## 如何贡献代码

非常感激任何的代码提交以使 lancet 的功能越来越强大。创建 pull request 时请遵守以下规则。

1. Fork lancet 仓库。
2. 创建自己的特性分支。
3. 提交变更。
4. Push 分支。
5. 创建新的 pull request。
