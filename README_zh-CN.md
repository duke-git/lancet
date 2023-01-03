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

1. <b>使用 go1.18 及以上版本的用户，建议安装 v2.x.x。 因为 v2.x.x 应用 go1.18 的泛型重写了大部分函数。</b>

```go
go get github.com/duke-git/lancet/v2 //安装v2最新版本v2.x.x
```

2. <b>使用 go1.18 以下版本的用户，必须安装 v1.x.x。目前最新的 v1 版本是 v1.3.5。</b>

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

### 1. algorithm 包实现一些基本查找和排序算法。

```go
import "github.com/duke-git/lancet/v2/algorithm"
```

#### Function list:

-   **<big>BubbleSort</big>** : 使用冒泡排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#BubbleSort)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>CountSort</big>** : 使用计数排序算法对切片进行排序。不改变原数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#CountSort)]
    [[play](https://go.dev/play/p/tB-Umgm0DrP)]
-   **<big>HeapSort</big>** : 使用堆排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#HeapSort)]
    [[play](https://go.dev/play/p/u6Iwa1VZS_f)]
-   **<big>InsertionSort</big>** : 使用插入排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#InsertionSort)]
    [[play](https://go.dev/play/p/G5LJiWgJJW6)]
-   **<big>MergeSort</big>** : 使用合并排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#MergeSort)]
    [[play](https://go.dev/play/p/ydinn9YzUJn)]
-   **<big>QuickSort</big>** : 使用快速排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#QuickSort)]
    [[play](https://go.dev/play/p/7Y7c1Elk3ax)]
-   **<big>SelectionSort</big>** : 使用选择排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#SelectionSort)]
    [[play](https://go.dev/play/p/oXovbkekayS)]
-   **<big>ShellSort</big>** : 使用希尔排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#ShellSort)]
    [[play](https://go.dev/play/p/3ibkszpJEu3)]
-   **<big>BinarySearch</big>** : 返回排序切片中目标值的索引，使用二分搜索（递归调用）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#BinarySearch)]
    [[play](https://go.dev/play/p/t6MeGiUSN47)]
-   **<big>BinaryIterativeSearch</big>** :返回排序切片中目标值的索引，使用二分搜索（非递归）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#BinaryIterativeSearch)]
    [[play](https://go.dev/play/p/Anozfr8ZLH3)]
-   **<big>LinearSearch</big>** : 基于传入的相等函数返回切片中目标值的索引。（线性查找）
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#LinearSearch)]
-   **<big>LRUCache</big>** : 应用 lru 算法实现内存缓存.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md#LRUCache)]
    [[play](https://go.dev/play/p/-EZjgOURufP)]

### 2. concurrency 包含一些支持并发编程的功能。例如：goroutine, channel, async 等。

```go
import "github.com/duke-git/lancet/v2/concurrency"
```

#### Function list:

-   **<big>NewChannel</big>** : 返回一个 Channel 指针实例。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#NewChannel)]
-   **<big>Bridge</big>** : 将多个 channel 链接到一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/Bridge.md#NewChannel)]
-   **<big>FanIn</big>** : 将多个 channel 合并为一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#FanIn)]
-   **<big>Generate</big>** : 根据传入的值，生成 channel。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Generate)]
-   **<big>Or</big>** : 将一个或多个 channel 读取到一个 channel 中，当任何读取 channel 关闭时将结束读取。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Or)]
-   **<big>OrDone</big>** : 将一个 channel 读入另一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#OrDone)]
-   **<big>Repeat</big>** : 返回一个 channel，将参数`values`重复放入 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Repeat)]
-   **<big>RepeatFn</big>** : 返回一个 channel，重复执行函数 fn，并将结果放入返回的 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#RepeatFn)]
-   **<big>Take</big>** : 返回一个 channel，其值从另一个 channel 获取，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Take)]
-   **<big>Tee</big>** : 将一个 channel 分成两个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Tee)]

### 3. condition 包含一些用于条件判断的函数。

```go
import "github.com/duke-git/lancet/v2/condition"
```

#### Function list:

-   **<big>Bool</big>** : 返回传入参数的 bool 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Bool)]
    [[play](https://go.dev/play/p/ETzeDJRSvhm)]
-   **<big>And</big>** : 逻辑且操作，当切仅当 a 和 b 都为 true 时返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#And)]
    [[play](https://go.dev/play/p/W1SSUmt6pvr)]
-   **<big>Or</big>** : 逻辑或操作，当切仅当 a 和 b 都为 false 时返回 false。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Or)]
    [[play](https://go.dev/play/p/UlQTxHaeEkq)]
-   **<big>Xor</big>** : 逻辑异或操作，a 和 b 相同返回 false，a 和 b 不相同返回 true
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Xor)]
    [[play](https://go.dev/play/p/gObZrW7ZbG8)]
-   **<big>Nor</big>** : 异或的取反操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Nor)]
    [[play](https://go.dev/play/p/g2j08F_zZky)
-   **<big>Xnor</big>** : 如果 a 和 b 都是真的或 a 和 b 均是假的，则返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Xnor)]
    [[play](https://go.dev/play/p/OuDB9g51643)]
-   **<big>Nand</big>** : 如果 a 和 b 都为真，返回 false，否则返回 true
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#Nand)]
    [[play](https://go.dev/play/p/vSRMLxLIbq8)]
-   **<big>TernaryOperator</big>** : 三元运算符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition_zh-CN.md#TernaryOperator)]
    [[play](https://go.dev/play/p/ElllPZY0guT)]

### 4. convertor 转换器包支持一些常见的数据类型转换。

```go
import "github.com/duke-git/lancet/v2/convertor"
```

#### 函数列表:

-   **<big>ColorHexToRGB</big>** : 颜色值十六进制转 rgb。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ColorHexToRGB)]
    [[play](https://go.dev/play/p/o7_ft-JCJBV)]
-   **<big>ColorRGBToHex</big>** : 颜色值 rgb 转十六进制。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ColorRGBToHex)]
    [[play](https://go.dev/play/p/nzKS2Ro87J1)]
-   **<big>ToBool</big>** : 字符串转布尔类型，使用 strconv.ParseBool。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToBool)]
    [[play](https://go.dev/play/p/ARht2WnGdIN)]
-   **<big>ToBytes</big>** : interface 转字节切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToBytes)]
    [[play](https://go.dev/play/p/fAMXYFDvOvr)]
-   **<big>ToChar</big>** : 字符串转字符切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToChar)]
    [[play](https://go.dev/play/p/JJ1SvbFkVdM)]
-   **<big>ToChannel</big>** : 将切片转为只读 channel。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToChannel)]
    [[play](https://go.dev/play/p/hOx_oYZbAnL)]
-   **<big>ToFloat</big>** : 将 interface 转成 float64 类型，如果参数无法转换，会返回 0.0 和 error。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToFloat)]
    [[play](https://go.dev/play/p/4YTmPCibqHJ)]
-   **<big>ToInt</big>** : 将 interface 转成 int64 类型，如果参数无法转换，会返回 0 和 error。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToInt)]
    [[play](https://go.dev/play/p/9_h9vIt-QZ_b)]
-   **<big>ToJson</big>** : 将 interface 转成 json 字符串，如果参数无法转换，会返回""和 error。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToJson)]
    [[play](https://go.dev/play/p/2rLIkMmXWvR)]
-   **<big>ToMap</big>** : 将切片转为 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToMap)]
    [[play](https://go.dev/play/p/tVFy7E-t24l)]
-   **<big>ToPointer</big>** : 返回传入值的指针。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToPointer)]
    [[play](https://go.dev/play/p/ASf_etHNlw1)]
-   **<big>ToString</big>** : 将值转换为字符串，对于数字、字符串、[]byte，将转换为字符串。 对于其他类型（切片、映射、数组、结构）将调用 json.Marshal。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#ToString)]
    [[play](https://go.dev/play/p/nF1zOOslpQq)]
-   **<big>StructToMap</big>** : 将 struct 转成 map，只会转换 struct 中可导出的字段。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#StructToMap)]
    [[play](https://go.dev/play/p/KYGYJqNUBOI)]
-   **<big>MapToSlice</big>** : map 中 key 和 value 执行函数 iteratee 后，转为切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#MapToSlice)]
    [[play](https://go.dev/play/p/dmX4Ix5V6Wl)]
-   **<big>EncodeByte</big>** : 将传入的 data 编码成字节切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#EncodeByte)]
    [[play](https://go.dev/play/p/DVmM1G5JfuP)]
-   **<big>DecodeByte</big>** : 解码字节切片到目标对象，目标对象需要传入一个指针实例。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#DecodeByte)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]

### 5. cryptor 加密包支持数据加密和解密，获取 md5，hash 值。支持 base64, md5, hmac, aes, des, rsa。

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

-   **<big>ClearFile</big>** : 清空文件内容。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ClearFile)]
    [[play](https://go.dev/play/p/NRZ0ZT-G94H)]
-   **<big>CreateFile</big>** : 创建文件，创建成功返回 true, 否则返回 false。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CreateFile)]
    [[play](https://go.dev/play/p/lDt8PEsTNKI)]
-   **<big>CreateDir</big>** : 创建嵌套目录，例如/a/, /a/b/。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CreateDir)]
    [[play](https://go.dev/play/p/qUuCe1OGQnM)]
-   **<big>CopyFile</big>** :拷贝文件，会覆盖原有的文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#CopyFile)]
    [[play](https://go.dev/play/p/Jg9AMJMLrJi)]
-   **<big>FileMode</big>** : 获取文件 mode 信息。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#FileMode)]
    [[play](https://go.dev/play/p/2l2hI42fA3p)]
-   **<big>MiMeType</big>** : 获取文件 mime 类型, 参数的类型必须是 string 或者\*os.File。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#MiMeType)]
    [[play](https://go.dev/play/p/bd5sevSUZNu)]
-   **<big>IsExist</big>** : 判断文件或目录是否存在。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsExist)]
    [[play](https://go.dev/play/p/nKKXt8ZQbmh)]
-   **<big>IsLink</big>** : 判断文件是否是符号链接。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsLink)]
    [[play](https://go.dev/play/p/TL-b-Kzvf44)]
-   **<big>IsDir</big>** : 判断参数是否是目录。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#IsDir)]
    [[play](https://go.dev/play/p/WkVwEKqtOWk)]
-   **<big>ListFileNames</big>** : 返回目录下所有文件名。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ListFileNames)]
    [[play](https://go.dev/play/p/Tjd7Y07rejl)]
-   **<big>RemoveFile</big>** : 删除文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#RemoveFile)]
    [[play](https://go.dev/play/p/P2y0XW8a1SH)]
-   **<big>ReadFileToString</big>** : 读取文件内容并返回字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ReadFileToString)]
    [[play](https://go.dev/play/p/cmfwp_5SQTp)]
-   **<big>ReadFileByLine</big>** : 按行读取文件内容，返回字符串切片包含每一行。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#ReadFileByLine)]
    [[play](https://go.dev/play/p/svJP_7ZrBrD)]
-   **<big>Zip</big>** : zip 压缩文件, 参数可以是文件或目录。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#Zip)]
    [[play](https://go.dev/play/p/j-3sWBp8ik_P)]
-   **<big>UnZip</big>** : zip 解压缩文件并保存在目录中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil_zh-CN.md#UnZip)]
    [[play](https://go.dev/play/p/g0w34kS7B8m)]

### 9. formatter 格式化器包含一些数据格式化处理方法。

```go
import "github.com/duke-git/lancet/v2/formatter"
```

#### 函数列表:

-   **<big>Comma</big>** : 用逗号每隔 3 位分割数字/字符串，支持前缀添加符号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/formatter_zh-CN.md#Comma)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]

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

### 11. maputil 包括一些操作 map 的函数.

```go
import "github.com/duke-git/lancet/v2/maputil"
```

#### 函数列表:

-   **<big>ForEach</big>** : 对 map 中的每对 key 和 value 执行 iteratee 函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#ForEach)]
    [[play](https://go.dev/play/p/OaThj6iNVXK)]
-   **<big>Filter</big>** : 迭代 map 中的每对 key 和 value，返回 map，其中的 key 和 value 符合 predicate 函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Filter)]
    [[play](https://go.dev/play/p/fSvF3wxuNG7)]
-   **<big>Intersect</big>** : 多个 map 的交集操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Intersect)]
    [[play](https://go.dev/play/p/Zld0oj3sjcC)]
-   **<big>Keys</big>** : 返回 map 中所有 key 组成的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Keys)]
    [[play](https://go.dev/play/p/xNB5bTb97Wd)]
-   **<big>Merge</big>** : 合并多个 map, 相同的 key 会被之后的 key 覆盖。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Merge)]
    [[play](https://go.dev/play/p/H95LENF1uB-)]
-   **<big>Minus</big>** : 返回一个 map，其中的 key 存在于 mapA，不存在于 mapB。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Minus)]
    [[play](https://go.dev/play/p/3u5U9K7YZ9m)]
-   **<big>Values</big>** : 返回 map 中所有 values 组成的切片
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Values)]
    [[play](https://go.dev/play/p/CBKdUc5FTW6)]
-   **<big>IsDisjoint</big>** : 验证两个 map 是否具有不同的 key。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#IsDisjoint)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]

### 12. mathutil 包实现了一些数学计算的函数。

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### Function list:

-   **<big>Average</big>** :计算平均数，可能需要对结果调用 RoundToFloat 方法四舍五入。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Average)]
    [[play](https://go.dev/play/p/Vv7LBwER-pz)]
-   **<big>Exponent</big>** : 指数计算（x 的 n 次方）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Exponent)]
    [[play](https://go.dev/play/p/uF3HGNPk8wr)]
-   **<big>Fibonacci</big>** :计算斐波那契数列的第 n 个数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Fibonacci)]
    [[play](https://go.dev/play/p/IscseUNMuUc)]
-   **<big>Factorial</big>** : 计算阶乘。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Factorial)]
    [[play](https://go.dev/play/p/tt6LdOK67Nx)]
-   **<big>Max</big>** : 返回参数中的最大数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Max)]
    [[play](https://go.dev/play/p/cN8DHI0rTkH)]
-   **<big>MaxBy</big>** : 使用给定的比较器函数返回切片的最大值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#MaxBy)]
    [[play](https://go.dev/play/p/pbe2MT-7DV2)]
-   **<big>Min</big>** : 返回参数中的最小数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Min)]
    [[play](https://go.dev/play/p/21BER_mlGUj)]
-   **<big>MinBy</big>** : 使用给定的比较器函数返回切片的最小值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#MinBy)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
-   **<big>Percent</big>** : 计算百分比，可以指定保留 n 位小数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#Percent)]
-   **<big>RoundToFloat</big>** : 四舍五入，保留 n 位小数，返回 float64。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#RoundToFloat)]
    [[play](https://go.dev/play/p/ghyb528JRJL)]
-   **<big>RoundToString</big>** : 四舍五入，保留 n 位小数，返回 string。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#RoundToString)]
    [[play](https://go.dev/play/p/kZwpBRAcllO)]
-   **<big>TruncRound</big>** : 截短 n 位小数（不进行四舍五入）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil_zh-CN.md#TruncRound)]
    [[play](https://go.dev/play/p/aumarSHIGzP)]

### 13. netutil 网络包支持获取 ip 地址，发送 http 请求。

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

### 14. random 随机数生成器包，可以生成随机[]bytes, int, string。

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

### 15. retry 重试执行函数直到函数运行成功或被 context cancel。

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### 函数列表:

-   [Context](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Context)
-   [Retry](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Retry)
-   [RetryFunc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryFunc)
-   [RetryDuration](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryDuration)
-   [RetryTimes](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryTimes)

### 16. slice 包含操作切片的方法集合。

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

### 17. strutil 包含字符串处理的相关函数。

```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### 函数列表:

-   **<big>After</big>** : 返回源字符串中指定字符串首次出现时的位置之后的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#After)]
    [[play](https://go.dev/play/p/RbCOQqCDA7m)]
-   **<big>AfterLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之后的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#AfterLast)]
    [[play](https://go.dev/play/p/1TegARrb8Yn)]
-   **<big>Before</big>** : 返回源字符串中指定字符串第一次出现时的位置之前的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Before)]
    [[play](https://go.dev/play/p/JAWTZDS4F5w)]
-   **<big>BeforeLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之前的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#BeforeLast)]
    [[play](https://go.dev/play/p/pJfXXAoG_Te)]
-   **<big>CamelCase</big>** : 将字符串转换为 CamelCase 驼峰式字符串, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#CamelCase)]
    [[play](https://go.dev/play/p/9eXP3tn2tUy)]
-   **<big>Capitalize</big>** : 将字符串的第一个字符转换为大写。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Capitalize)]
    [[play](https://go.dev/play/p/2OAjgbmAqHZ)]
-   **<big>IsString</big>** : 判断传入参数的数据类型是否为字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#IsString)]
    [[play](https://go.dev/play/p/IOgq7oF9ERm)]
-   **<big>KebabCase</big>** : 将字符串转换为 kebab-case 形式字符串, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#KebabCase)]
    [[play](https://go.dev/play/p/dcZM9Oahw-Y)]
-   **<big>UpperKebabCase</big>** : 将字符串转换为大写 KEBAB-CASE 形式字符串, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#UpperKebabCase)]
    [[play](https://go.dev/play/p/zDyKNneyQXk)]
-   **<big>LowerFirst</big>** : 将字符串的第一个字符转换为小写形式。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#LowerFirst)]
    [[play](https://go.dev/play/p/CbzAyZmtJwL)]
-   **<big>UpperFirst</big>** : 将字符串的第一个字符转换为大写形式。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#UpperFirst)]
    [[play](https://go.dev/play/p/sBbBxRbs8MM)]
-   **<big>PadEnd</big>** : 如果字符串短于限制大小，则在右侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#PadEnd)]
    [[play](https://go.dev/play/p/9xP8rN0vz--)]
-   **<big>PadStart</big>** : 如果字符串短于限制大小，则在左侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#PadStart)]
    [[play](https://go.dev/play/p/xpTfzArDfvT)]
-   **<big>Reverse</big>** : 返回字符顺序与给定字符串相反的字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Reverse)]
    [[play](https://go.dev/play/p/adfwalJiecD)]
-   **<big>SnakeCase</big>** : 将字符串转换为 snake_case 形式, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#SnakeCase)]
    [[play](https://go.dev/play/p/tgzQG11qBuN)]
-   **<big>UpperSnakeCase</big>** : 将字符串转换为大写 SNAKE_CASE 形式, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#UpperSnakeCase)]
    [[play](https://go.dev/play/p/4COPHpnLx38)]
-   **<big>SplitEx</big>** : 拆分给定的字符串可以控制结果切片是否包含空字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#SplitEx)]
    [[play](https://go.dev/play/p/Us-ySSbWh-3)]
-   **<big>Substring</big>** : 根据指定的位置和长度截取子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Substring)]
-   **<big>Wrap</big>** : 用给定字符包裹传入的字符串
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Wrap)]
    [[play](https://go.dev/play/p/KoZOlZDDt9y)]
-   **<big>Unwrap</big>** : 从另一个字符串中解开一个给定的字符串。 将更改源字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Unwrap)]
    [[play](https://go.dev/play/p/Ec2q4BzCpG-)]

### 18. system 包含 os, runtime, shell command 的相关函数。

```go
import "github.com/duke-git/lancet/v2/system"
```

#### 函数列表:

-   **<big>IsWindows</big>** : 检查当前操作系统是否是 windows。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#IsWindows)]
    [[play](https://go.dev/play/p/XzJULbzmf9m)]
-   **<big>IsLinux</big>** : 检查当前操作系统是否是 linux。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#IsLinux)]
    [[play](https://go.dev/play/p/zIflQgZNuxD)]
-   **<big>IsMac</big>** : 检查当前操作系统是否是 macos。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#IsMac)]
    [[play](https://go.dev/play/p/Mg4Hjtyq7Zc)]
-   **<big>GetOsEnv</big>** : 根据 key 获取对应的环境变量值
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#GetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>SetOsEnv</big>** : 设置环境变量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#SetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>RemoveOsEnv</big>** : 删除环境变量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#RemoveOsEnv)]
    [[play](https://go.dev/play/p/fqyq4b3xUFQ)]
-   **<big>CompareOsEnv</big>** : 换取环境变量并与传入值进行比较。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#CompareOsEnv)]
    [[play](https://go.dev/play/p/BciHrKYOHbp)]
-   **<big>ExecCommand</big>** : 执行 shell 命令。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#ExecCommand)]
    [[play](https://go.dev/play/p/n-2fLyZef-4)]
-   **<big>GetOsBits</big>** : 获取当前操作系统位数(32/64)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system_zh-CN#GetOsBits)]
    [[play](https://go.dev/play/p/ml-_XH3gJbW)]

### 19. validator 验证器包，包含常用字符串格式验证函数。

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

### 20. xerror 包实现一些错误处理函数

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
