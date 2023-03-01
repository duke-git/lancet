<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-2.1.16-green.svg)](https://github.com/duke-git/lancet/releases)
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
-   💪 400+常用 go 工具函数，支持 string、slice、datetime、net、crypt...
-   💅 只依赖 go 标准库
-   🌍 所有导出函数单元测试覆盖率 100%

## 安装

### Note:

1. <b>使用 go1.18 及以上版本的用户，建议安装 v2.x.x。 因为 v2.x.x 应用 go1.18 的泛型重写了大部分函数。</b>

```go
go get github.com/duke-git/lancet/v2 //安装v2最新版本v2.x.x
```

2. <b>使用 go1.18 以下版本的用户，必须安装 v1.x.x。目前最新的 v1 版本是 v1.3.7。</b>

```go
go get github.com/duke-git/lancet@v1.3.7 // 使用go1.18以下版本, 必须安装v1.x.x版本
```

## 用法

lancet 是以包的结构组织代码的，使用时需要导入相应的包名。例如：如果使用字符串相关函数，需要导入 strutil 包:

```go
import "github.com/duke-git/lancet/v2/strutil"
```

## 示例

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
    [[play](https://go.dev/play/p/IsS7rgn5s3x)]
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
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Bridge</big>** : 将多个 channel 链接到一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/Bridge.md#NewChannel)]
    [[play](https://go.dev/play/p/qmWSy1NVF-Y)]
-   **<big>FanIn</big>** : 将多个 channel 合并为一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#FanIn)]
    [[play](https://go.dev/play/p/2VYFMexEvTm)]
-   **<big>Generate</big>** : 根据传入的值，生成 channel。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Generate)]
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Or</big>** : 将一个或多个 channel 读取到一个 channel 中，当任何读取 channel 关闭时将结束读取。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Or)]
    [[play](https://go.dev/play/p/Wqz9rwioPww)]
-   **<big>OrDone</big>** : 将一个 channel 读入另一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#OrDone)]
    [[play](https://go.dev/play/p/lm_GoS6aDjo)]
-   **<big>Repeat</big>** : 返回一个 channel，将参数`values`重复放入 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Repeat)]
    [[play](https://go.dev/play/p/k5N_ALVmYjE)]
-   **<big>RepeatFn</big>** : 返回一个 channel，重复执行函数 fn，并将结果放入返回的 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#RepeatFn)]
    [[play](https://go.dev/play/p/4J1zAWttP85)]
-   **<big>Take</big>** : 返回一个 channel，其值从另一个 channel 获取，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Take)]
    [[play](https://go.dev/play/p/9Utt-1pDr2J)]
-   **<big>Tee</big>** : 将一个 channel 分成两个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency_zh-CN.md#Tee)]
    [[play](https://go.dev/play/p/3TQPKnCirrP)]

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
[[play](https://go.dev/play/p/OuDB9g51643)]]
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
-   **<big>DeepClone</big>** : 创建一个传入值的深拷贝, 无法克隆结构体的非导出字段。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#DeepClone)]
    [[play](https://go.dev/play/p/j4DP5dquxnk)]
-   **<big>CopyProperties</big>** : 拷贝不同结构体之间的同名字段。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor_zh-CN.md#CopyProperties)]
    [[play](https://go.dev/play/p/FOVY3XJL-6B)]

### 5. cryptor 加密包支持数据加密和解密，获取 md5，hash 值。支持 base64, md5, hmac, aes, des, rsa。

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### 函数列表:

-   **<big>AesEcbEncrypt</big>** : 使用 AES ECB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesEcbEncrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesEcbDecrypt</big>** : 使用 AES ECB 算法模解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesEcbDecrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesCbcEncrypt</big>** : 使用 AES CBC 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCbcEncrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCbcDecrypt</big>** : 使用 AES CBC 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCbcDecrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCtrCrypt</big>** : 使用 AES CTR 算法模式加密/解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCtrCrypt)]
    [[play](https://go.dev/play/p/SpaZO0-5Nsp)]
-   **<big>AesCfbEncrypt</big>** : 使用 AES CFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCfbEncrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesCfbDecrypt</big>** : 使用 AES CFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesCfbDecrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesOfbEncrypt</big>** : 使用 AES OFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesOfbEncrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>AesOfbDecrypt</big>** : 使用 AES OFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#AesOfbDecrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>Base64StdEncode</big>** : 将字符串 base64 编码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Base64StdEncode)]
    [[play](https://go.dev/play/p/VOaUyQUreoK)]
-   **<big>Base64StdDecode</big>** : 解码 base64 字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Base64StdDecode)]
    [[play](https://go.dev/play/p/RWQylnJVgIe)]
-   **<big>DesEcbEncrypt</big>** : 使用 DES ECB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesEcbEncrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesEcbDecrypt</big>** : 使用 DES ECB 算法模解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesEcbDecrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesCbcEncrypt</big>** : 使用 DES CBC 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCbcEncrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCbcDecrypt</big>** : 使用 DES CBC 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCbcDecrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCtrCrypt</big>** : 使用 DES CTR 算法模式加密/解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCtrCrypt)]
    [[play](https://go.dev/play/p/9-T6OjKpcdw)]
-   **<big>DesCfbEncrypt</big>** : 使用 DES CFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCfbEncrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesCfbDecrypt</big>** : 使用 DES CFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesCfbDecrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesOfbEncrypt</big>** : 使用 DES OFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesOfbEncrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>DesOfbDecrypt</big>** : 使用 DES OFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#DesOfbDecrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>HmacMd5</big>** : 返回字符串 md5 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacMd5)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>HmacSha1</big>** : 返回字符串 sha1 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha1)]
    [[play](https://go.dev/play/p/1UI4oQ4WXKM)]
-   **<big>HmacSha256</big>** : 返回字符串 sha256 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha256)]
    [[play](https://go.dev/play/p/HhpwXxFhhC0)]
-   **<big>HmacSha512</big>** : 返回字符串 sha256 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#HmacSha512)]
    [[play](https://go.dev/play/p/59Od6m4A0Ud)]
-   **<big>Md5String</big>** : 返回字符串 md5 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Md5String)]
    [[play](https://go.dev/play/p/1bLcVetbTOI)]
-   **<big>Md5File</big>** : 返回文件 md5 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Md5File)]
-   **<big>Sha1</big>** : 返回字符串 sha1 哈希值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha1)]
    [[play](https://go.dev/play/p/_m_uoD1deMT)]
-   **<big>Sha256</big>** :返回字符串 sha256 哈希值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha256)]
    [[play](https://go.dev/play/p/tU9tfBMIAr1)]
-   **<big>Sha512</big>** : 返回字符串 sha512 哈希值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#Sha512)]
    [[play](https://go.dev/play/p/3WsvLYZxsHa)]
-   **<big>GenerateRsaKey</big>** : 在当前目录下创建 rsa 私钥文件和公钥文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#GenerateRsaKey)]
    [[play](https://go.dev/play/p/zutRHrDqs0X)]
-   **<big>RsaEncrypt</big>** : 用公钥文件 ras 加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#RsaEncrypt)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>RsaDecrypt</big>** : 用私钥文件 rsa 解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor_zh-CN.md#RsaDecrypt)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]

### 6. datetime 日期时间处理包，格式化日期，比较日期。

```go
import "github.com/duke-git/lancet/v2/datetime"
```

#### 函数列表:

-   **<big>AddDay</big>** : 将日期加/减天数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddDay)]
    [[play](https://go.dev/play/p/dIGbs_uTdFa)]
-   **<big>AddHour</big>** : 将日期加/减小时数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddHour)]
    [[play](https://go.dev/play/p/rcMjd7OCsi5)]
-   **<big>AddMinute</big>** : 将日期加/减分钟数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#AddMinute)]
    [[play](https://go.dev/play/p/nT1heB1KUUK)]
-   **<big>BeginOfMinute</big>** : 返回指定时间的分钟开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#BeginOfMinute)]
    [[play](https://go.dev/play/p/ieOLVJ9CiFT)]
-   **<big>BeginOfHour</big>** : 返回指定时间的小时开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#BeginOfHour)]
    [[play](https://go.dev/play/p/GhdGFnDWpYs)]
-   **<big>BeginOfDay</big>** : 返回指定时间的当天开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#BeginOfDay)]
    [[play](https://go.dev/play/p/94m_UT6cWs9)]
-   **<big>BeginOfWeek</big>** : 返回指定时间的每周开始时间,默认开始时间星期日。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#BeginOfWeek)]
    [[play](https://go.dev/play/p/ynjoJPz7VNV)]
-   **<big>BeginOfMonth</big>** : 返回指定时间的当月开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#BeginOfMonth)]
    [[play](https://go.dev/play/p/bWXVFsmmzwL)]
-   **<big>BeginOfYear</big>** : 返回指定时间的当年开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#BeginOfYear)]
    [[play](https://go.dev/play/p/i326DSwLnV8)]
-   **<big>EndOfMinute</big>** : 返回指定时间的分钟结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#EndOfMinute)]
    [[play](https://go.dev/play/p/yrL5wGzPj4z)]
-   **<big>EndOfHour</big>** : 返回指定时间的小时结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#EndOfHour)]
    [[play](https://go.dev/play/p/6ce3j_6cVqN)]
-   **<big>EndOfDay</big>** : 返回指定时间的当天结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#EndOfDay)]
    [[play](https://go.dev/play/p/eMBOvmq5Ih1)]
-   **<big>EndOfWeek</big>** : 返回指定时间的星期结束时间,默认结束时间星期六。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#EndOfWeek)]
    [[play](https://go.dev/play/p/i08qKXD9flf)]
-   **<big>EndOfMonth</big>** : 返回指定时间的月份结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#EndOfMonth)]
    [[play](https://go.dev/play/p/_GWh10B3Nqi)]
-   **<big>EndOfYear</big>** : 返回指定时间的年份结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#EndOfYear)]
    [[play](https://go.dev/play/p/G01cKlMCvNm)]
-   **<big>GetNowDate</big>** : 获取当天日期，返回格式：yyyy-mm-dd。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowDate)]
    [[play](https://go.dev/play/p/PvfkPpcpBBf)]
-   **<big>GetNowTime</big>** : 获取当时时间，返回格式：hh:mm:ss。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowTime)]
    [[play](https://go.dev/play/p/l7BNxCkTmJS)]
-   **<big>GetNowDateTime</big>** : 获取当时日期和时间，返回格式：yyyy-mm-dd hh:mm:ss。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNowDateTime)]
    [[play](https://go.dev/play/p/pI4AqngD0al)]
-   **<big>GetZeroHourTimestamp</big>** : 获取零时时间戳(timestamp of 00:00)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetZeroHourTimestamp)]
    [[play](https://go.dev/play/p/QmL2oIaGE3q)]
-   **<big>GetNightTimestamp</big>** : 获取午夜时间戳(timestamp of 23:59)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#GetNightTimestamp)]
    [[play](https://go.dev/play/p/UolysR3MYP1)]
-   **<big>FormatTimeToStr</big>** : 将日期格式化成字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#FormatTimeToStr)]
    [[play](https://go.dev/play/p/_Ia7M8H_OvE)]
-   **<big>FormatStrToTime</big>** : 将字符串格式化成时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#FormatStrToTime)]
    [[play](https://go.dev/play/p/1h9FwdU8ql4)]
-   **<big>NewUnix</big>** : 创建一个 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewUnix)]
    [[play](https://go.dev/play/p/psoSuh_kLRt)]
-   **<big>NewUnixNow</big>** : 创建一个当前时间的 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewUnixNow)]
    [[play](https://go.dev/play/p/U4PPx-9D0oz)]
-   **<big>NewFormat</big>** : 创建一个 yyyy-mm-dd hh:mm:ss 格式时间字符串的 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>NewISO8601</big>** : 创建一个 iso8601 格式时间字符串的 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#NewISO8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]
-   **<big>ToUnix</big>** : 返回 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToUnix)]
    [[play](https://go.dev/play/p/_LUiwAdocjy)]
-   **<big>ToFormat</big>** : 返回格式'yyyy-mm-dd hh:mm:ss'的日期字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>ToFormatForTpl</big>** : 返回 tpl 格式指定的日期字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToFormatForTpl)]
    [[play](https://go.dev/play/p/nyXxXcQJ8L5)]
-   **<big>ToIso8601</big>** : 返回 iso8601 日期字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime_zh-CN.md#ToIso8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]

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

-   **<big>List</big>** : 线性表结构, 用切片实现。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/list_zh-CN.md)]
-   **<big>Link</big>** : 链表解构, 包括单链表和双向链表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/link_zh-CN.md)]
-   **<big>Stack</big>** : 栈结构(fifo), 包括数组栈和链表栈。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/stack_zh-CN.md)]
-   **<big>Queue</big>** : 队列结构(filo), 包括数组队列，链表队列，循环队列，优先级队列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/queue_zh-CN.md)]
-   **<big>Set</big>** : 集合（set）结构。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/set_zh-CN.md)]
-   **<big>Tree</big>** : 二叉搜索树。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/tree_zh-CN.md)]
-   **<big>Heap</big>** : 二叉 max 堆。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/heap_zh-CN.md)]
-   **<big>Hashmap</big>** : 哈希映射。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/hashmap_zh-CN.md)]

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

-   **<big>After</big>** : 创建一个函数，当该函数被调用 n 或更多次之后将执行传入的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#After)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]
-   **<big>Before</big>** : 创建一个函数，当该函数被调用不超过 n 次时，将执行执行传入的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Before)]
    [[play](https://go.dev/play/p/0HqUDIFZ3IL)]
-   **<big>CurryFn</big>** : 创建柯里化函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#CurryFn)]
    [[play](https://go.dev/play/p/5HopfDwANKX)]
-   **<big>Compose</big>** : 从右至左组合函数列表 fnList，返回组合后的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Compose)]
    [[play](https://go.dev/play/p/KKfugD4PKYF)]
-   **<big>Delay</big>** : 延迟 delay 时间后调用函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Delay)]
    [[play](https://go.dev/play/p/Ivtc2ZE-Tye)]
-   **<big>Debounced</big>** : 创建一个 debounced 函数，该函数延迟调用 fn 直到自上次调用 debounced 函数后等待持续时间过去。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Debounced)]
    [[play](https://go.dev/play/p/absuEGB_GN7)]
-   **<big>Schedule</big>** : 每次持续时间调用函数，直到关闭返回的 channel。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Schedule)]
    [[play](https://go.dev/play/p/hbON-Xeyn5N)]
-   **<big>Pipeline</big>** : 从右至左执行函数列表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Pipeline)]
    [[play](https://go.dev/play/p/mPdUVvj6HD6)]
-   **<big>Watcher</big>** : Watcher 用于记录代码执行时间。可以启动/停止/重置手表定时器。获取函数执行的时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function_zh-CN.md#Watcher)]
    [[play](https://go.dev/play/p/l2yrOpCLd1I)]

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
-   **<big>FilterByKeys</big>** : 迭代 map, 返回一个新 map，其 key 都是给定的 key 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#FilterByKeys)]
    [[play](https://go.dev/play/p/7ov6BJHbVqh)]
-   **<big>FilterByValues</big>** : 迭代 map, 返回一个新 map，其 value 都是给定的 value 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#FilterByValues)]
    [[play](https://go.dev/play/p/P3-9MdcXegR)]
-   **<big>OmitBy</big>** : Filter 的反向操作, 迭代 map 中的每对 key 和 value, 删除符合 predicate 函数的 key, value, 返回新 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#OmitBy)]
    [[play](https://go.dev/play/p/YJM4Hj5hNwm)]
-   **<big>OmitByKeys</big>** : FilterByKeys 的反向操作, 迭代 map, 返回一个新 map，其 key 不包括给定的 key 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#OmitByKeys)]
    [[play](https://go.dev/play/p/jXGrWDBfSRp)]
-   **<big>OmitByValues</big>** : FilterByValues 的反向操作, 迭代 map, 返回一个新 map，其 value 不包括给定的 value 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#OmitByValues)]
    [[play](https://go.dev/play/p/XB7Y10uw20_U)]
-   **<big>Intersect</big>** : 多个 map 的交集操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Intersect)]
    [[play](https://go.dev/play/p/Zld0oj3sjcC)]
-   **<big>Keys</big>** : 返回 map 中所有 key 组成的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Keys)]
    [[play](https://go.dev/play/p/xNB5bTb97Wd)]
-   **<big>KeysBy</big>** : 创建一个切片，其元素是每个 map 的 key 调用 mapper 函数的结果。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#KeysBy)]
    [[play](https://go.dev/play/p/hI371iB8Up8)]
-   **<big>Merge</big>** : 合并多个 map, 相同的 key 会被之后的 key 覆盖。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Merge)]
    [[play](https://go.dev/play/p/H95LENF1uB-)]
-   **<big>Minus</big>** : 返回一个 map，其中的 key 存在于 mapA，不存在于 mapB。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Minus)]
    [[play](https://go.dev/play/p/3u5U9K7YZ9m)]
-   **<big>Values</big>** : 返回 map 中所有 values 组成的切片
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN.md#Values)]
    [[play](https://go.dev/play/p/CBKdUc5FTW6)]
-   **<big>ValuesBy</big>** : 创建一个切片，其元素是每个 map 的 value 调用 mapper 函数的结果。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#ValuesBy)]
    [[play](https://go.dev/play/p/sg9-oRidh8f)]
-   **<big>MapKeys</big>** : 操作 map 的每个 key，然后转为新的 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#MapKeys)]
    [[play](https://go.dev/play/p/8scDxWeBDKd)]
-   **<big>MapValues</big>** : 操作 map 的每个 value，然后转为新的 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#MapValues)]
    [[play](https://go.dev/play/p/g92aY3fc7Iw)]
-   **<big>Entries</big>** : 将 map 转换为键/值对切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#Entries)]
    [[play](https://go.dev/play/p/Ltb11LNcElY)]
-   **<big>FromEntries</big>** : 基于键/值对的切片创建 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#FromEntries)]
    [[play](https://go.dev/play/p/fTdu4sCNjQO)]
-   **<big>Transform</big>** : 将 map 转换为其他类型的 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil_zh-CN#Transform)]
    [[play](https://go.dev/play/p/P6ovfToM3zj)]
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
    [[play](https://go.dev/play/p/QQM9B13coSP)]
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

-   **<big>ConvertMapToQueryString</big>** : 将 map 转换成 http 查询字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#ConvertMapToQueryString)]
    [[play](https://go.dev/play/p/jnNt_qoSnRi)]
-   **<big>EncodeUrl</big>** : 编码 url query string 的值(?a=1&b=[2] -> ?a=1&b=%5B2%5D)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#EncodeUrl)]
    [[play](https://go.dev/play/p/bsZ6BRC4uKI)]
-   **<big>GetInternalIp</big>** : 获取内部 ipv4。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetInternalIp)]
    [[play](https://go.dev/play/p/5mbu-gFp7ei)]
-   **<big>GetIps</big>** : 获取系统 ipv4 地址列表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetIps)]
    [[play](https://go.dev/play/p/NUFfcEmukx1)]
-   **<big>GetMacAddrs</big>** : 获取系统 mac 地址列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetMacAddrs)]
    [[play](https://go.dev/play/p/Rq9UUBS_Xp1)]
-   **<big>GetPublicIpInfo</big>** : 获取[公网 ip 信息](http://ip-api.com/json/).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetPublicIpInfo)]
    [[play](https://go.dev/play/p/YDxIfozsRHR)]
-   **<big>GetRequestPublicIp</big>** : 获取 http 请求 ip。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#GetRequestPublicIp)]
    [[play](https://go.dev/play/p/kxU-YDc_eBo)]
-   **<big>IsPublicIP</big>** : 判断 ip 是否是公共 ip。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#IsPublicIP)]
    [[play](https://go.dev/play/p/nmktSQpJZnn)]
-   **<big>IsInternalIP</big>** : 判断 ip 是否是局域网 ip。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#IsInternalIP)]
    [[play](https://go.dev/play/p/sYGhXbgO4Cb)]
-   **<big>HttpRequest</big>** : 用于抽象 HTTP 请求实体的结构。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>HttpClient</big>** : 用于发送 HTTP 请求。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpClient)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>SendRequest</big>** : 发送 http 请求。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#SendRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>DecodeResponse</big>** : 解析 http 响应体到目标结构体。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#DecodeResponse)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>StructToUrlValues</big>** : 将结构体转为 url values, 仅转化结构体导出字段并且包含`json` tag。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#StructToUrlValues)]
    [[play](https://go.dev/play/p/pFqMkM40w9z)]
-   **<big>HttpGet<sup>deprecated</sup></big>** : 发送 http get 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpGet)]
-   **<big>HttpDelete<sup>deprecated</sup></big>** : 发送 http delete 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpDelete)]
-   **<big>HttpPost<sup>deprecated</sup></big>** : 发送 http post 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPost)]
-   **<big>HttpPut<sup>deprecated</sup></big>** : 发送 http put 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPut)]
-   **<big>HttpPatch<sup>deprecated</sup></big>** : 发送 http patch 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#HttpPatch)]
-   **<big>ParseHttpResponse</big>** : 解析 http 响应体到目标结构体。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil_zh-CN.md#ParseHttpResponse)]

### 14. random 随机数生成器包，可以生成随机[]bytes, int, string。

```go
import "github.com/duke-git/lancet/v2/random"
```

#### 函数列表:

-   **<big>RandBytes</big>** : 生成随机字节切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandBytes)]
    [[play](https://go.dev/play/p/EkiLESeXf8d)]
-   **<big>RandInt</big>** : 生成随机 int, 范围[min, max)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandInt)]
    [[play](https://go.dev/play/p/pXyyAAI5YxD)]
-   **<big>RandString</big>** : 生成给定长度的随机字符串，只包含字母(a-zA-Z)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandString)]
    [[play](https://go.dev/play/p/W2xvRUXA7Mi)]
-   **<big>RandUpper</big>** : 生成给定长度的随机大写字母字符串(A-Z)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandUpper)]
    [[play](https://go.dev/play/p/29QfOh0DVuh)]
-   **<big>RandLower</big>** : 生成给定长度的随机小写字母字符串(a-z)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandLower)]
    [[play](https://go.dev/play/p/XJtZ471cmtI)]
-   **<big>RandNumeral</big>** : 生成给定长度的随机数字字符串(0-9)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandNumeral)]
    [[play](https://go.dev/play/p/g4JWVpHsJcf)]
-   **<big>RandNumeralOrLetter</big>** : 生成给定长度的随机字符串（数字+字母)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#RandNumeralOrLetter)]
    [[play](https://go.dev/play/p/19CEQvpx2jD)]
-   **<big>UUIdV4</big>** : 生成 UUID v4 字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random_zh-CN.md#UUIdV4)]
    [[play](https://go.dev/play/p/_Z9SFmr28ft)]

### 15. retry 重试执行函数直到函数运行成功或被 context cancel。

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### 函数列表:

-   **<big>Context</big>** : 设置重试 context 参数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Context)]
    [[play](https://go.dev/play/p/xnAOOXv9GkS)]
-   **<big>Retry</big>** : 重试执行函数 retryFunc，直到函数运行成功，或被 context 取消。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#Retry)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryFunc</big>** : 重试执行的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryFunc)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryDuration</big>** : 设置重试间隔时间，默认 3 秒。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryDuration)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryTimes</big>** : 设置重试次数，默认 5。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry_zh-CN.md#RetryTimes)]
    [[play](https://go.dev/play/p/ssfVeU2SwLO)]

### 16. slice 包含操作切片的方法集合。

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### 函数列表:

-   **<big>AppendIfAbsent</big>** : 当前切片中不包含值时，将该值追加到切片中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#AppendIfAbsent)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>Contain</big>** : 判断 slice 是否包含 value。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Contain)]
    [[play](https://go.dev/play/p/_454yEHcNjf)]
-   **<big>ContainBy</big>** : 根据 predicate 函数判断切片是否包含某个值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ContainBy)]
    [[play](https://go.dev/play/p/49tkHfX4GNc)]
-   **<big>ContainSubSlice</big>** : 判断 slice 是否包含 subslice。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ContainSubSlice)]
    [[play](https://go.dev/play/p/bcuQ3UT6Sev)]
-   **<big>Chunk</big>** : 按照 size 参数均分 slice。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Chunk)]
    [[play](https://go.dev/play/p/b4Pou5j2L_C)]
-   **<big>Compact</big>** : 去除 slice 中的假值（false values are false, nil, 0, ""）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Compact)]
    [[play](https://go.dev/play/p/pO5AnxEr3TK)]
-   **<big>Concat</big>** : 合并多个 slices 到一个 slice 中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Concat)]
    [[play](https://go.dev/play/p/gPt-q7zr5mk)]
-   **<big>Count</big>** : 返回切片中指定元素的个数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Count)]
    [[play](https://go.dev/play/p/Mj4oiEnQvRJ)]
-   **<big>CountBy</big>** : 遍历切片，对每个元素执行函数 predicate. 返回符合函数返回值为 true 的元素的个数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#CountBy)]
    [[play](https://go.dev/play/p/tHOccTMDZCC)]
-   **<big>Difference</big>** : 创建一个切片，其元素不包含在另一个给定切片中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Difference)]
    [[play](https://go.dev/play/p/VXvadzLzhDa)]
-   **<big>DifferenceBy</big>** : 将两个 slice 中的每个元素调用 iteratee 函数，并比较它们的返回值，如果不相等返回在 slice 中对应的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DifferenceBy)]
    [[play](https://go.dev/play/p/DiivgwM5OnC)]
-   **<big>DifferenceWith</big>** : 接受比较器函数，该比较器被调用以将切片的元素与值进行比较。 结果值的顺序和引用由第一个切片确定。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DifferenceWith)]
    [[play](https://go.dev/play/p/v2U2deugKuV)]
-   **<big>DeleteAt</big>** : 删除切片中指定开始索引到结束索引的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DeleteAt)]
    [[play](https://go.dev/play/p/pJ-d6MUWcvK)]
-   **<big>Drop</big>** : 从切片头部删除 n 个元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Drop)]
    [[play](https://go.dev/play/p/jnPO2yQsT8H)]
-   **<big>DropRight</big>** : 从切片尾部删除 n 个元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DropRight)]
    [[play](https://go.dev/play/p/8bcXvywZezG)]
-   **<big>DropWhile</big>** : 从切片的头部删除 n 个元素，这个 n 个元素满足 predicate 函数返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DropWhile)]
    [[play](https://go.dev/play/p/4rt252UV_qs)]
-   **<big>DropRightWhile</big>** : 从切片的尾部删除 n 个元素，这个 n 个元素满足 predicate 函数返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#DropRightWhile)]
    [[play](https://go.dev/play/p/6wyK3zMY56e)]
-   **<big>Equal</big>** : 检查两个切片是否相等，相等条件：切片长度相同，元素顺序和值都相同。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Equal)]
    [[play](https://go.dev/play/p/WcRQJ37ifPa)]
-   **<big>EqualWith</big>** : 检查两个切片是否相等，相等条件：对两个切片的元素调用比较函数 comparator，返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#EqualWith)]
    [[play](https://go.dev/play/p/b9iygtgsHI1)]
-   **<big>Every</big>** : 如果切片中的所有值都通过谓词函数，则返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Every)]
    [[play](https://go.dev/play/p/R8U6Sl-j8cD)]
-   **<big>Filter</big>** : 返回切片中通过 predicate 函数真值测试的所有元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Filter)]
    [[play](https://go.dev/play/p/SdPna-7qK4T)]
-   **<big>FilterMap</big>** : 返回一个将 filter 和 map 操作应用于给定切片的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#FilterMap)]
    [[play](https://go.dev/play/p/J94SZ_9MiIe)]
-   **<big>Find</big>** : 遍历切片的元素，返回第一个通过 predicate 函数真值测试的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Find)]
    [[play](https://go.dev/play/p/CBKeBoHVLgq)]
-   **<big>FindLast</big>** : 从头到尾遍历 slice 的元素，返回最后一个通过 predicate 函数真值测试的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#FindLast)]
    [[play](https://go.dev/play/p/FFDPV_j7URd)]
-   **<big>Flatten</big>** : 将多维切片展平一层。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Flatten)]
    [[play](https://go.dev/play/p/hYa3cBEevtm)]
-   **<big>FlattenDeep</big>** : 将多维切片递归展平到一层。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#FlattenDeep)]
    [[play](https://go.dev/play/p/yjYNHPyCFaF)]
-   **<big>FlatMap</big>** : 将切片转换为其它类型切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#FlatMap)]
    [[play](https://go.dev/play/p/_QARWlWs1N_F)]
-   **<big>ForEach</big>** : 遍历切片的元素并为每个元素调用 iteratee 函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ForEach)]
    [[play](https://go.dev/play/p/DrPaa4YsHRF)]
-   **<big>GroupBy</big>** : 迭代切片的元素，每个元素将按条件分组，返回两个切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#GroupBy)]
    [[play](https://go.dev/play/p/QVkPxzPR0iA)]
-   **<big>GroupWith</big>** : 创建一个 map，key 是 iteratee 遍历 slice 中的每个元素返回的结果。值是切片元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#GroupWith)]
    [[play](https://go.dev/play/p/ApCvMNTLO8a)]
-   **<big>IntSlice<sup>deprecated</sup></big>** : 将接口切片转换为 int 切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IntSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>InterfaceSlice<sup>deprecated</sup></big>** : 将值转换为 interface 切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#InterfaceSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>Intersection</big>** : 返回多个切片的交集。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Intersection)]
    [[play](https://go.dev/play/p/anJXfB5wq_t)]
-   **<big>InsertAt</big>** : 将元素插入到索引处的切片中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#InsertAt)]
    [[play](https://go.dev/play/p/hMLNxPEGJVE)]
-   **<big>IndexOf</big>** : 返回在切片中找到值的第一个匹配项的索引，如果找不到值，则返回-1。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IndexOf)]
    [[play](https://go.dev/play/p/MRN1f0FpABb)]
-   **<big>LastIndexOf</big>** : 返回在切片中找到最后一个值的索引，如果找不到该值，则返回-1。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#LastIndexOf)]
    [[play](https://go.dev/play/p/DokM4cf1IKH)]
-   **<big>Map</big>** : 对 slice 中的每个元素执行 map 函数以创建一个新切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Map)]
    [[play](https://go.dev/play/p/biaTefqPquw)]
-   **<big>Merge</big>** : 合并多个切片（不会消除重复元素)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Merge)]
    [[play](https://go.dev/play/p/lbjFp784r9N)]
-   **<big>Reverse</big>** : 反转切片中的元素顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Reverse)]
    [[play](https://go.dev/play/p/8uI8f1lwNrQ)]
-   **<big>Reduce</big>** : 将切片中的元素依次运行 iteratee 函数，返回运行结果。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Reduce)]
    [[play](https://go.dev/play/p/_RfXJJWIsIm)]
-   **<big>Replace</big>** : 返回切片的副本，其中前 n 个不重叠的 old 替换为 new。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Replace)]
    [[play](https://go.dev/play/p/P5mZp7IhOFo)]
-   **<big>ReplaceAll</big>** : 返回切片的副本，将其中 old 全部替换为 new。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ReplaceAll)]
    [[play](https://go.dev/play/p/CzqXMsuYUrx)]
-   **<big>Repeat</big>** : 创建一个切片，包含 n 个传入的 item。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Repeat)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>Shuffle</big>** : 随机打乱切片中的元素顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Shuffle)]
    [[play](https://go.dev/play/p/YHvhnWGU3Ge)]
-   **<big>IsAscending</big>** : 检查切片元素是否按升序排列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IsAscending)]
    [[play](https://go.dev/play/p/9CtsFjet4SH)]
-   **<big>IsDescending</big>** : 检查切片元素是否按降序排列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IsDescending)]
    [[play](https://go.dev/play/p/U_LljFXma14)]
-   **<big>IsSorted</big>** : 检查切片元素是否是有序的（升序或降序）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IsSorted)]
    [[play](https://go.dev/play/p/nCE8wPLwSA-)]
-   **<big>IsSortedByKey</big>** : 通过 iteratee 函数，检查切片元素是否是有序的。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#IsSortedByKey)]
    [[play](https://go.dev/play/p/tUoGB7DOHI4)]
-   **<big>Sort</big>** : 对任何有序类型（数字或字符串）的切片进行排序，使用快速排序算法。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Sort)]
    [[play](https://go.dev/play/p/V9AVjzf_4Fk)]
-   **<big>SortBy</big>** : 按照 less 函数确定的升序规则对切片进行排序。排序不保证稳定性。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#SortBy)]
    [[play](https://go.dev/play/p/DAhLQSZEumm)]
-   **<big>SortByField<sup>deprecated</sup></big>** : 按字段对结构切片进行排序。slice 元素应为 struct，字段类型应为 int、uint、string 或 bool。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#SortByField)]
    [[play](https://go.dev/play/p/fU1prOBP9p1)]
-   **<big>Some</big>** : 如果列表中的任何值通过谓词函数，则返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Some)]
    [[play](https://go.dev/play/p/4pO9Xf9NDGS)]
-   **<big>StringSlice<sup>deprecated</sup></big>** : 将接口切片转换为字符串切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#StringSlice)]
    [[play](https://go.dev/play/p/W0TZDWCPFcI)]
-   **<big>SymmetricDifference</big>** : 返回一个切片，其中的元素存在于参数切片中，但不同时存储在于参数切片中（交集取反）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#h42nJX5xMln)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>ToSlice</big>** : 将可变参数转为切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ToSlice)]
    [[play](https://go.dev/play/p/YzbzVq5kscN)]
-   **<big>ToSlicePointer</big>** : 将可变参数转为指针切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#ToSlicePointer)]
    [[play](https://go.dev/play/p/gx4tr6_VXSF)]
-   **<big>Unique</big>** : 删除切片中的重复元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Unique)]
    [[play](https://go.dev/play/p/AXw0R3ZTE6a)]
-   **<big>UniqueBy</big>** : 对切片的每个元素调用 iteratee 函数，然后删除重复元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#UniqueBy)]
    [[play](https://go.dev/play/p/UR323iZLDpv)]
-   **<big>Union</big>** : 合并多个切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Union)]
    [[play](https://go.dev/play/p/hfXV1iRIZOf)]
-   **<big>UnionBy</big>** : 对切片的每个元素调用函数后，合并多个切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#UnionBy)]
    [[play](https://go.dev/play/p/HGKHfxKQsFi)]
-   **<big>UpdateAt</big>** : 更新索引处的切片元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#UpdateAt)]
    [[play](https://go.dev/play/p/f3mh2KloWVm)]
-   **<big>Without</big>** : 创建一个不包括所有给定值的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#Without)]
    [[play](https://go.dev/play/p/bwhEXEypThg)]
-   **<big>KeyBy</big>** :将切片每个元素调用函数后转为 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice_zh-CN.md#KeyBy)]
    [[play](https://go.dev/play/p/uXod2LWD1Kg)]

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
-   **<big>Pad</big>** : 如果字符串长度短于 size，则在左右两侧填充字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Pad)]
    [[play](https://go.dev/play/p/NzImQq-VF8q)]
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
    [[play](https://go.dev/play/p/q3sM6ehnPDp)]
-   **<big>Wrap</big>** : 用给定字符包裹传入的字符串
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Wrap)]
    [[play](https://go.dev/play/p/KoZOlZDDt9y)]
-   **<big>Unwrap</big>** : 从另一个字符串中解开一个给定的字符串。 将更改源字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#Unwrap)]
    [[play](https://go.dev/play/p/Ec2q4BzCpG-)]
-   **<big>SplitWords</big>** : 将字符串拆分为单词，只支持字母字符单词。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#SplitWords)]
    [[play](https://go.dev/play/p/KLiX4WiysMM)]
-   **<big>WordCount</big>** : 返回有意义单词的数量，只支持字母字符单词。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil_zh-CN.md#WordCount)]
    [[play](https://go.dev/play/p/bj7_odx3vRf)]


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

-   **<big>ContainChinese</big>** : 验证字符串是否包含中文字符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainChinese)]
    [[play](https://go.dev/play/p/7DpU0uElYeM)]
-   **<big>ContainLetter</big>** : 验证字符串是否包含至少一个英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainLetter)]
    [[play](https://go.dev/play/p/lqFD04Yyewp)]
-   **<big>ContainLower</big>** : 验证字符串是否包含至少一个英文小写字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainLower)]
    [[play](https://go.dev/play/p/Srqi1ItvnAA)]
-   **<big>ContainUpper</big>** : 验证字符串是否包含至少一个英文大写字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#ContainUpper)]
    [[play](https://go.dev/play/p/CmWeBEk27-z)]
-   **<big>IsAlpha</big>** : 验证字符串是否只包含英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAlpha)]
    [[play](https://go.dev/play/p/7Q5sGOz2izQ)]
-   **<big>IsAllUpper</big>** : 验证字符串是否全是大写英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAllUpper)]
    [[play](https://go.dev/play/p/ZHctgeK1n4Z)]
-   **<big>IsAllLower</big>** : 验证字符串是否全是小写英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsAllLower)]
    [[play](https://go.dev/play/p/GjqCnOfV6cM)]
-   **<big>IsBase64</big>** : 验证字符串是否是 base64 编码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsBase64)]
    [[play](https://go.dev/play/p/sWHEySAt6hl)]
-   **<big>IsChineseMobile</big>** : 验证字符串是否是中国手机号码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChineseMobile)]
    [[play](https://go.dev/play/p/GPYUlGTOqe3)]
-   **<big>IsChineseIdNum</big>** : 验证字符串是否是中国身份证号码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChineseIdNum)]
    [[play](https://go.dev/play/p/d8EWhl2UGDF)]
-   **<big>IsChinesePhone</big>** : 验证字符串是否是中国电话座机号码(xxx-xxxxxxxx or xxxx-xxxxxxx.)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsChinesePhone)]
    [[play](https://go.dev/play/p/RUD_-7YZJ3I)]
-   **<big>IsCreditCard</big>** : 验证字符串是否是信用卡号码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsCreditCard)]
    [[play](https://go.dev/play/p/sNwwL6B0-v4)]
-   **<big>IsDns</big>** : 验证字符串是否是有效 dns。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsDns)]
    [[play](https://go.dev/play/p/jlYApVLLGTZ)]
-   **<big>IsEmail</big>** : 验证字符串是否是有效电子邮件地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsEmail)]
    [[play](https://go.dev/play/p/Os9VaFlT33G)]
-   **<big>IsEmptyString</big>** : 验证字符串是否是空字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsEmptyString)]
    [[play](https://go.dev/play/p/dpzgUjFnBCX)]
-   **<big>IsFloatStr</big>** : 验证字符串是否是可以转换为浮点数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#LOYwS_Oyl7U)]
    [[play](https://go.dev/play/p/LOYwS_Oyl7U)]
-   **<big>IsNumberStr</big>** : 验证字符串是否是可以转换为数字。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsNumberStr)]
    [[play](https://go.dev/play/p/LzaKocSV79u)]
-   **<big>IsJSON</big>** : 验证字符串是否是有效 json。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsJSON)]
    [[play](https://go.dev/play/p/sRS6c4K8jGk)]
-   **<big>IsRegexMatch</big>** : 验证字符串是否可以匹配正则表达式。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsRegexMatch)]
    [[play](https://go.dev/play/p/z_XeZo_litG)]
-   **<big>IsIntStr</big>** : 验证字符串是否是可以转换为整数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIntStr)]
    [[play](https://go.dev/play/p/jQRtFv-a0Rk)]
-   **<big>IsIp</big>** : 验证字符串是否是 ip 地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIp)]
    [[play](https://go.dev/play/p/FgcplDvmxoD)]
-   **<big>IsIpV4</big>** : 验证字符串是否是 ipv4 地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIpV4)]
    [[play](https://go.dev/play/p/zBGT99EjaIu)]
-   **<big>IsIpV6</big>** : 验证字符串是否是 ipv6 地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsIpV6)]
    [[play](https://go.dev/play/p/AHA0r0AzIdC)]
-   **<big>IsStrongPassword</big>** : 验证字符串是否是强密码：（字母+数字+特殊字符)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsStrongPassword)]
    [[play](https://go.dev/play/p/QHdVcSQ3uDg)]
-   **<big>IsUrl</big>** : 验证字符串是否是 url。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsUrl)]
    [[play](https://go.dev/play/p/pbJGa7F98Ka)]
-   **<big>IsWeakPassword</big>** : 验证字符串是否是弱密码（只包含字母+数字）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsWeakPassword)]
    [[play](https://go.dev/play/p/wqakscZH5gH)]
-   **<big>IsZeroValue</big>** : 判断传入的参数值是否为零值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsZeroValue)]
    [[play](https://go.dev/play/p/UMrwaDCi_t4)]
-   **<big>IsGBK</big>** : 检查数据编码是否为 gbk（汉字内部代码扩展规范）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator_zh-CN.md#IsGBK)]
    [[play](https://go.dev/play/p/E2nt3unlmzP)]

### 20. xerror 包实现一些错误处理函数

```go
import "github.com/duke-git/lancet/v2/xerror"
```

#### 函数列表:

-   **<big>New</big>** : 创建 XError 对象实例。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#New)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>Wrap</big>** : 根据 error 对象创建 XError 对象实例，可添加 message。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#Wrap)]
    [[play](https://go.dev/play/p/5385qT2dCi4)]
-   **<big>Unwrap</big>** : 从 error 对象中解构出 XError。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#Unwrap)]
    [[play](https://go.dev/play/p/LKMLep723tu)]
-   **<big>XError_Wrap</big>** : 创建新的 XError 对象并将消息和 id 复制到新的对象中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_Wrap)]
    [[play](https://go.dev/play/p/5385qT2dCi4)]
-   **<big>XError_Unwrap</big>** : 解构 XEerror 为 error 对象。适配 github.com/pkg/errors。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_Unwrap)]
    [[play](https://go.dev/play/p/VUXJ8BST4c6)]
-   **<big>XError_With</big>** : 添加与 XError 对象的键和值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_With)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_Id</big>** : 设置 XError 对象的 id。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_Id)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Is</big>** : 检查目标 error 是否为 XError，两个错误中的 error.id 是否匹配。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_Is)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Values</big>** : 返回由 With 设置的键和值的映射。将合并所有 XError 键和值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_Values)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_StackTrace</big>** : 返回与 pkg/error 兼容的堆栈信息。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_StackTrace)]
    [[play](https://go.dev/play/p/6FAvSQpa7pc)]
-   **<big>XError_Info</big>** : 返回可打印的 XError 对象信息。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_Info)]
    [[play](https://go.dev/play/p/1ZX0ME1F-Jb)]
-   **<big>XError_Error</big>** : 实现标准库的 error 接口。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#XError_Error)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>TryUnwrap</big>** : 检查 error, 如果 err 为 nil 则展开，则它返回一个有效值，如果 err 不是 nil 则 Unwrap 使用 err 发生 panic。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror_zh-CN.md#TryUnwrap)]
    [[play](https://go.dev/play/p/acyZVkNZEeW)]

## 如何贡献代码

非常感激任何的代码提交以使 lancet 的功能越来越强大。创建 pull request 时请遵守以下规则。

1. Fork lancet 仓库。
2. 创建自己的特性分支。
3. 提交变更。
4. Push 分支。
5. 创建新的 pull request。
