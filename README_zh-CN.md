<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-2.3.8-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com/duke-git/lancet/v2?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/duke-git/lancet/v2)](https://goreportcard.com/report/github.com/duke-git/lancet/v2)
[![test](https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/duke-git/lancet/actions/workflows/codecov.yml)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/main/LICENSE)
[![Gurubase](https://img.shields.io/badge/Gurubase-Ask%20Lancet%20Guru-006BFF)](https://gurubase.io/g/lancet)

</div>

<div STYLE="page-break-after: always;"></div>
<p style="font-size: 18px">
    lancet（柳叶刀）是一个全面、高效、可复用的go语言工具函数库。 lancet受到了java apache common包和lodash.js的启发。
</p>

## <a href="https://www.golancet.cn/" target="_blank">官网</a> | [English](./README.md)

## 特性

-   👏 全面、高效、可复用。
-   💪 700+常用 go 工具函数，支持 string、slice、datetime、net、crypt...
-   💅 只依赖 go 标准库和 golang.org/x。
-   🌍 所有导出函数单元测试覆盖率 100%。

## 安装

### Note:

1. <b>使用 go1.18 及以上版本的用户，建议安装 v2.x.x。 因为 v2.x.x 应用 go1.18 的泛型重写了大部分函数。</b>

```go
go get github.com/duke-git/lancet/v2 //安装v2最新版本v2.x.x
```

2. <b>使用 go1.18 以下版本的用户，必须安装 v1.x.x。目前最新的 v1 版本是 v1.4.6。</b>

```go
go get github.com/duke-git/lancet// 使用go1.18以下版本, 必须安装v1.x.x版本
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

### <span id="index">目录<span>

-   [Algorithm](#user-content-algorithm)
-   [Compare](#user-content-compare)
-   [Concurrency](#user-content-concurrency)
-   [Condition](#user-content-condition)
-   [Convertor](#user-content-convertor)
-   [Cryptor](#user-content-cryptor)
-   [Datetime](#user-content-datetime)
-   [Datastructure](#user-content-datastructure)
-   [EventBus](#user-content-eventbus)
-   [Enum](#user-content-enum)
-   [Fileutil](#user-content-fileutil)
-   [Formatter](#user-content-formatter)
-   [Function](#user-content-function)
-   [Maputil](#user-content-maputil)
-   [Mathutil](#user-content-mathutil)
-   [Netutil](#user-content-netutil)
-   [Pointer](#user-content-pointer)
-   [Random](#user-content-random)
-   [Retry](#user-content-retry)
-   [Slice](#user-content-slice)
-   [Stream](#user-content-stream)
-   [Structs](#user-content-structs)
-   [Strutil](#user-content-strutil)
-   [System](#user-content-system)
-   [Tuple](#user-content-tuple)
-   [Validator](#user-content-validator)
-   [Xerror](#user-content-xerror)

<h3 id="algorithm"> 1. algorithm 包实现一些基本查找和排序算法。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/algorithm"
```

#### 函数列表:

-   **<big>BubbleSort</big>** : 使用冒泡排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#BubbleSort)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>CountSort</big>** : 使用计数排序算法对切片进行排序。不改变原数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#CountSort)]
    [[play](https://go.dev/play/p/tB-Umgm0DrP)]
-   **<big>HeapSort</big>** : 使用堆排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#HeapSort)]
    [[play](https://go.dev/play/p/u6Iwa1VZS_f)]
-   **<big>InsertionSort</big>** : 使用插入排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#InsertionSort)]
    [[play](https://go.dev/play/p/G5LJiWgJJW6)]
-   **<big>MergeSort</big>** : 使用合并排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#MergeSort)]
    [[play](https://go.dev/play/p/ydinn9YzUJn)]
-   **<big>QuickSort</big>** : 使用快速排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#QuickSort)]
    [[play](https://go.dev/play/p/7Y7c1Elk3ax)]
-   **<big>SelectionSort</big>** : 使用选择排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#SelectionSort)]
    [[play](https://go.dev/play/p/oXovbkekayS)]
-   **<big>ShellSort</big>** : 使用希尔排序算法对切片进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#ShellSort)]
    [[play](https://go.dev/play/p/3ibkszpJEu3)]
-   **<big>BinarySearch</big>** : 返回排序切片中目标值的索引，使用二分搜索（递归调用）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#BinarySearch)]
    [[play](https://go.dev/play/p/t6MeGiUSN47)]
-   **<big>BinaryIterativeSearch</big>** :返回排序切片中目标值的索引，使用二分搜索（非递归）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#BinaryIterativeSearch)]
    [[play](https://go.dev/play/p/Anozfr8ZLH3)]
-   **<big>LinearSearch</big>** : 基于传入的相等函数返回切片中目标值的索引。（线性查找）
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#LinearSearch)]
    [[play](https://go.dev/play/p/IsS7rgn5s3x)]
-   **<big>LRUCache</big>** : 应用 lru 算法实现内存缓存.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/algorithm.md#LRUCache)]
    [[play](https://go.dev/play/p/-EZjgOURufP)]

<h3 id="compare"> 2. compare 包提供几个轻量级的类型比较函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/compare"
```

#### 函数列表:

-   **<big>Equal</big>** : 检查两个值是否相等(检查类型和值)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/compare.md#Equal)]
    [[play](https://go.dev/play/p/wmVxR-to4lz)]
-   **<big>EqualValue</big>** : 检查两个值是否相等(只检查值)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/compare.md#EqualValue)]
    [[play](https://go.dev/play/p/fxnna_LLD9u)]
-   **<big>LessThan</big>** : 验证参数`left`的值是否小于参数`right`的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/compare.md#LessThan)]
    [[play](https://go.dev/play/p/cYh7FQQj0ne)]
-   **<big>GreaterThan</big>** : 验证参数`left`的值是否大于参数`right`的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/compare.md#GreaterThan)]
    [[play](https://go.dev/play/p/9-NYDFZmIMp)]
-   **<big>LessOrEqual</big>** : 验证参数`left`的值是否小于或等于参数`right`的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/compare.md#LessOrEqual)]
    [[play](https://go.dev/play/p/e4T_scwoQzp)]
-   **<big>GreaterOrEqual</big>** : 验证参数`left`的值是否大于或等于参数`right`的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/compare.md#GreaterOrEqual)]
    [[play](https://go.dev/play/p/vx8mP0U8DFk)]
-   **<big>InDelta</big>** : 检查增量内两个值是否相等。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/compare.md#InDelta)]

<h3 id="concurrency"> 3. concurrency 包含一些支持并发编程的功能。例如：goroutine, channel, async 等。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/concurrency"
```

#### 函数列表:

-   **<big>NewChannel</big>** : 返回一个 Channel 指针实例。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#NewChannel)]
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Bridge</big>** : 将多个 channel 链接到一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/Bridge.md#NewChannel)]
    [[play](https://go.dev/play/p/qmWSy1NVF-Y)]
-   **<big>FanIn</big>** : 将多个 channel 合并为一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#FanIn)]
    [[play](https://go.dev/play/p/2VYFMexEvTm)]
-   **<big>Generate</big>** : 根据传入的值，生成 channel。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Generate)]
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Or</big>** : 将一个或多个 channel 读取到一个 channel 中，当任何读取 channel 关闭时将结束读取。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Or)]
    [[play](https://go.dev/play/p/Wqz9rwioPww)]
-   **<big>OrDone</big>** : 将一个 channel 读入另一个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#OrDone)]
    [[play](https://go.dev/play/p/lm_GoS6aDjo)]
-   **<big>Repeat</big>** : 返回一个 channel，将参数`values`重复放入 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Repeat)]
    [[play](https://go.dev/play/p/k5N_ALVmYjE)]
-   **<big>RepeatFn</big>** : 返回一个 channel，重复执行函数 fn，并将结果放入返回的 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#RepeatFn)]
    [[play](https://go.dev/play/p/4J1zAWttP85)]
-   **<big>Take</big>** : 返回一个 channel，其值从另一个 channel 获取，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Take)]
    [[play](https://go.dev/play/p/9Utt-1pDr2J)]
-   **<big>Tee</big>** : 将一个 channel 分成两个 channel，直到取消上下文。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Tee)]
    [[play](https://go.dev/play/p/3TQPKnCirrP)]
-   **<big>NewKeyedLocker</big>** : NewKeyedLocker 创建一个新的 KeyedLocker，并为锁的过期设置指定的 TTL。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#NewKeyedLocker)]
    [[play](https://go.dev/play/p/GzeyC33T5rw)]
-   **<big>Do</big>** :为指定的键获取锁并执行提供的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Do)]
    [[play](https://go.dev/play/p/GzeyC33T5rw)]
-   **<big>NewRWKeyedLocker</big>** :RWKeyedLocker 是一个简单的键值读写锁实现，允许非阻塞的锁获取。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#NewRWKeyedLocker)]
    [[play](https://go.dev/play/p/ZrCr8sMo77T)]
-   **<big>RLock</big>** : 为指定的键获取读锁并执行提供的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#RLock)]
    [[play](https://go.dev/play/p/ZrCr8sMo77T)]
-   **<big>Lock</big>** : 为指定的键获取锁并执行提供的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Lock)]
    [[play](https://go.dev/play/p/WgAcXbOPKGk)]
-   **<big>NewTryKeyedLocker</big>** : 创建一个 TryKeyedLocker 实例，TryKeyedLocker 是 KeyedLocker 的非阻塞版本。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#NewTryKeyedLocker)]
    [[play](https://go.dev/play/p/VG9qLvyetE2)]
-   **<big>TryLock</big>** : TryLock 尝试获取指定键的锁。如果锁成功获取，则返回 true，否则返回 false。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#TryLock)]
    [[play](https://go.dev/play/p/VG9qLvyetE2)]
-   **<big>Unlock</big>** : 释放指定键的锁。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/concurrency.md#Unlock)]
    [[play](https://go.dev/play/p/VG9qLvyetE2)]

<h3 id="condition"> 4. condition 包含一些用于条件判断的函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/condition"
```

#### 函数列表:

-   **<big>Bool</big>** : 返回传入参数的 bool 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#Bool)]
    [[play](https://go.dev/play/p/ETzeDJRSvhm)]
-   **<big>And</big>** : 逻辑且操作，当切仅当 a 和 b 都为 true 时返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#And)]
    [[play](https://go.dev/play/p/W1SSUmt6pvr)]
-   **<big>Or</big>** : 逻辑或操作，当切仅当 a 和 b 都为 false 时返回 false。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#Or)]
    [[play](https://go.dev/play/p/UlQTxHaeEkq)]
-   **<big>Xor</big>** : 逻辑异或操作，a 和 b 相同返回 false，a 和 b 不相同返回 true
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#Xor)]
    [[play](https://go.dev/play/p/gObZrW7ZbG8)]
-   **<big>Nor</big>** : 异或的取反操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#Nor)]
    [[play](https://go.dev/play/p/g2j08F_zZky)]
-   **<big>Xnor</big>** : 如果 a 和 b 都是真的或 a 和 b 均是假的，则返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#Xnor)]
    [[play](https://go.dev/play/p/OuDB9g51643)]
-   **<big>Nand</big>** : 如果 a 和 b 都为真，返回 false，否则返回 true
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#Nand)]
    [[play](https://go.dev/play/p/vSRMLxLIbq8)]
-   **<big>TernaryOperator</big>** : 三元运算符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/condition.md#TernaryOperator)]
    [[play](https://go.dev/play/p/ElllPZY0guT)]

<h3 id="convertor"> 5. convertor 转换器包支持一些常见的数据类型转换。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/convertor"
```

#### 函数列表:

-   **<big>ColorHexToRGB</big>** : 颜色值十六进制转 rgb。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ColorHexToRGB)]
    [[play](https://go.dev/play/p/o7_ft-JCJBV)]
-   **<big>ColorRGBToHex</big>** : 颜色值 rgb 转十六进制。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ColorRGBToHex)]
    [[play](https://go.dev/play/p/nzKS2Ro87J1)]
-   **<big>ToBool</big>** : 字符串转布尔类型，使用 strconv.ParseBool。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToBool)]
    [[play](https://go.dev/play/p/ARht2WnGdIN)]
-   **<big>ToBytes</big>** : interface 转字节切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToBytes)]
    [[play](https://go.dev/play/p/fAMXYFDvOvr)]
-   **<big>ToChar</big>** : 字符串转字符切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToChar)]
    [[play](https://go.dev/play/p/JJ1SvbFkVdM)]
-   **<big>ToChannel</big>** : 将切片转为只读 channel。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToChannel)]
    [[play](https://go.dev/play/p/hOx_oYZbAnL)]
-   **<big>ToFloat</big>** : 将 interface 转成 float64 类型，如果参数无法转换，会返回 0.0 和 error。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToFloat)]
    [[play](https://go.dev/play/p/4YTmPCibqHJ)]
-   **<big>ToInt</big>** : 将 interface 转成 int64 类型，如果参数无法转换，会返回 0 和 error。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToInt)]
    [[play](https://go.dev/play/p/9_h9vIt-QZ_b)]
-   **<big>ToJson</big>** : 将 interface 转成 json 字符串，如果参数无法转换，会返回""和 error。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToJson)]
    [[play](https://go.dev/play/p/2rLIkMmXWvR)]
-   **<big>ToMap</big>** : 将切片转为 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToMap)]
    [[play](https://go.dev/play/p/tVFy7E-t24l)]
-   **<big>ToPointer</big>** : 返回传入值的指针。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToPointer)]
    [[play](https://go.dev/play/p/ASf_etHNlw1)]
-   **<big>ToPointers</big>** : 将值的切片转换为指针的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToPointers)]
    [[play](https://go.dev/play/p/ZUoXd2i5ZkV)]
-   **<big>FromPointer</big>** : 返回指针所指向的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#FromPointer)]
    [[play](https://go.dev/play/p/wAp90V7Zu6g)]
-   **<big>FromPointers</big>** : 将指针的切片转换为值的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#FromPointers)]
    [[play](https://go.dev/play/p/qIPsyYtNy3Q)]
-   **<big>ToString</big>** : 将值转换为字符串，对于数字、字符串、[]byte，将转换为字符串。 对于其他类型（切片、映射、数组、结构）将调用 json.Marshal。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToString)]
    [[play](https://go.dev/play/p/nF1zOOslpQq)]
-   **<big>StructToMap</big>** : 将 struct 转成 map，只会转换 struct 中可导出的字段。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#StructToMap)]
    [[play](https://go.dev/play/p/KYGYJqNUBOI)]
-   **<big>MapToSlice</big>** : map 中 key 和 value 执行函数 iteratee 后，转为切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#MapToSlice)]
    [[play](https://go.dev/play/p/dmX4Ix5V6Wl)]
-   **<big>EncodeByte</big>** : 将传入的 data 编码成字节切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#EncodeByte)]
    [[play](https://go.dev/play/p/DVmM1G5JfuP)]
-   **<big>DecodeByte</big>** : 解码字节切片到目标对象，目标对象需要传入一个指针实例。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#DecodeByte)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>DeepClone</big>** : 创建一个传入值的深拷贝, 无法克隆结构体的非导出字段。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#DeepClone)]
    [[play](https://go.dev/play/p/j4DP5dquxnk)]
-   **<big>CopyProperties</big>** : 拷贝不同结构体之间的同名字段。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#CopyProperties)]
    [[play](https://go.dev/play/p/oZujoB5Sgg5)]
-   **<big>ToInterface</big>** : 将反射值转换成对应的 interface 类型。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToInterface)]
    [[play](https://go.dev/play/p/syqw0-WG7Xd)]
-   **<big>Utf8ToGbk</big>** : utf8 编码转 GBK 编码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#Utf8ToGbk)]
    [[play](https://go.dev/play/p/9FlIaFLArIL)]
-   **<big>GbkToUtf8</big>** : GBK 编码转 utf8 编码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#GbkToUtf8)]
    [[play](https://go.dev/play/p/OphmHCN_9u8)]
-   **<big>ToStdBase64</big>** : 将值转换为 StdBase64 编码的字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToStdBase64)]
    [[play](https://go.dev/play/p/_fLJqJD3NMo)]
-   **<big>ToUrlBase64</big>** : 将值转换为 url Base64 编码的字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToUrlBase64)]
    [[play](https://go.dev/play/p/C_d0GlvEeUR)]
-   **<big>ToRawStdBase64</big>** : 将值转换为 RawStdBase64 编码的字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToRawStdBase64)]
    [[play](https://go.dev/play/p/wSAr3sfkDcv)]
-   **<big>ToRawUrlBase64</big>** : 将值转换为 RawUrlBase64 编码的字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToRawUrlBase64)]
    [[play](https://go.dev/play/p/HwdDPFcza1O)]
-   **<big>ToBigInt</big>** : 将整数转为\*big.Int。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/convertor.md#ToBigInt)]
    [[play](https://go.dev/play/p/X3itkCxwB_x)]

<h3 id="cryptor"> 6. cryptor 加密包支持数据加密和解密，获取 md5，hash 值。支持 base64, md5, hmac, aes, des, rsa。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### 函数列表:

-   **<big>AesEcbEncrypt</big>** : 使用 AES ECB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesEcbEncrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesEcbDecrypt</big>** : 使用 AES ECB 算法模解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesEcbDecrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesCbcEncrypt</big>** : 使用 AES CBC 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesCbcEncrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCbcDecrypt</big>** : 使用 AES CBC 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesCbcDecrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCtrCrypt<sup>deprecated</sup></big>** : 使用 AES CTR 算法模式加密/解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesCtrCrypt)]
    [[play](https://go.dev/play/p/SpaZO0-5Nsp)]
-   **<big>AesCtrEncrypt</big>** : 使用 AES CTR 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesCtrCrypt)]
    [[play](https://go.dev/play/p/x6pjPAvThRz)]
-   **<big>AesCtrDecrypt</big>** : 使用 AES CTR 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesCtrCrypt)]
    [[play](https://go.dev/play/p/x6pjPAvThRz)]
-   **<big>AesCfbEncrypt</big>** : 使用 AES CFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesCfbEncrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesCfbDecrypt</big>** : 使用 AES CFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesCfbDecrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesOfbEncrypt</big>** : 使用 AES OFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesOfbEncrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>AesOfbDecrypt</big>** : 使用 AES OFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesOfbDecrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>AesGcmEncrypt</big>** : 使用 AES GCM 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesGcmEncrypt)]
    [[play](https://go.dev/play/p/rUt0-DmsPCs)]
-   **<big>AesGcmDecrypt</big>** : 使用 AES GCM 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#AesGcmDecrypt)]
    [[play](https://go.dev/play/p/rUt0-DmsPCs)]
-   **<big>Base64StdEncode</big>** : 将字符串 base64 编码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Base64StdEncode)]
    [[play](https://go.dev/play/p/VOaUyQUreoK)]
-   **<big>Base64StdDecode</big>** : 解码 base64 字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Base64StdDecode)]
    [[play](https://go.dev/play/p/RWQylnJVgIe)]
-   **<big>DesEcbEncrypt</big>** : 使用 DES ECB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesEcbEncrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesEcbDecrypt</big>** : 使用 DES ECB 算法模解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesEcbDecrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesCbcEncrypt</big>** : 使用 DES CBC 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesCbcEncrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCbcDecrypt</big>** : 使用 DES CBC 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesCbcDecrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCtrCrypt<sup>deprecated</sup></big>** : 使用 DES CTR 算法模式加密/解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesCtrCrypt)]
    [[play](https://go.dev/play/p/9-T6OjKpcdw)]
-   **<big>DesCtrEncrypt</big>** : 使用 DES CTR 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesCtrEncrypt)]
    [[play](https://go.dev/play/p/S6p_WHCgH1d)]
-   **<big>DesCtrDecrypt</big>** : 使用 DES CTR 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesCtrDecrypt)]
    [[play](https://go.dev/play/p/S6p_WHCgH1d)]
-   **<big>DesCfbEncrypt</big>** : 使用 DES CFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesCfbEncrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesCfbDecrypt</big>** : 使用 DES CFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesCfbDecrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesOfbEncrypt</big>** : 使用 DES OFB 算法模式加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesOfbEncrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>DesOfbDecrypt</big>** : 使用 DES OFB 算法模式解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#DesOfbDecrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>HmacMd5</big>** : 返回字符串 md5 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacMd5)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>HmacMd5WithBase64</big>** : 获取字符串 md5 hmac base64 字符串值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacMd5WithBase64)]
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacMd5WithBase64)]
-   **<big>HmacSha1</big>** : 返回字符串 sha1 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacSha1)]
    [[play](https://go.dev/play/p/1UI4oQ4WXKM)]
-   **<big>HmacSha1WithBase64</big>** : 获取字符串的 sha1 base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacSha1WithBase64)]
    [[play](https://go.dev/play/p/47JmmGrnF7B)]
-   **<big>HmacSha256</big>** : 返回字符串 sha256 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacSha256)]
    [[play](https://go.dev/play/p/HhpwXxFhhC0)]
-   **<big>HmacSha256WithBase64</big>** : 获取字符串 sha256 hmac base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacSha256WithBase64)]
    [[play](https://go.dev/play/p/EKbkUvPTLwO)]
-   **<big>HmacSha512</big>** : 返回字符串 sha256 hmac 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacSha512)]
    [[play](https://go.dev/play/p/59Od6m4A0Ud)]
-   **<big>HmacSha512WithBase64</big>** : 获取字符串 sha512 hmac base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#HmacSha512WithBase64)]
    [[play](https://go.dev/play/p/c6dSe3E2ydU)]
-   **<big>Md5Byte</big>** : 返回 byte slice 的 md5 值.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Md5Byte)]
    [[play](https://go.dev/play/p/suraalH8lyC)]
-   **<big>Md5ByteWithBase64</big>** : 获取 byte slice 的 md5 base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Md5ByteWithBase64)]
    [[play](https://go.dev/play/p/Tcb-Z7LN2ax)]
-   **<big>Md5String</big>** : 返回字符串 md5 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Md5String)]
    [[play](https://go.dev/play/p/1bLcVetbTOI)]
-   **<big>Md5StringWithBase64</big>** : 获取字符串 md5 base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Md5StringWithBase64)]
    [[play](https://go.dev/play/p/Lx4gH7Vdr5_y)]
-   **<big>Md5File</big>** : 返回文件 md5 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Md5File)]
-   **<big>Sha1</big>** : 返回字符串 sha1 哈希值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Sha1)]
    [[play](https://go.dev/play/p/_m_uoD1deMT)]
-   **<big>Sha1WithBase64</big>** : 获取字符串 sha1 base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Sha1WithBase64)]
    [[play](https://go.dev/play/p/fSyx-Gl2l2-)]
-   **<big>Sha256</big>** :返回字符串 sha256 哈希值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Sha256)]
    [[play](https://go.dev/play/p/tU9tfBMIAr1)]
-   **<big>Sha256WithBase64</big>** : 获取字符串 sha256 base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Sha256WithBase64)]
    [[play](https://go.dev/play/p/85IXJHIal1k)]
-   **<big>Sha512</big>** : 返回字符串 sha512 哈希值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Sha512)]
    [[play](https://go.dev/play/p/3WsvLYZxsHa)]
-   **<big>Sha512WithBase64</big>** : 获取字符串 sha512 base64 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#Sha512WithBase64)]
    [[play](https://go.dev/play/p/q_fY2rA-k5I)]
-   **<big>GenerateRsaKey</big>** : 在当前目录下创建 rsa 私钥文件和公钥文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#GenerateRsaKey)]
    [[play](https://go.dev/play/p/zutRHrDqs0X)]
-   **<big>RsaEncrypt</big>** : 用公钥文件 ras 加密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#RsaEncrypt)]
    [[play](https://go.dev/play/p/7_zo6mrx-eX)]
-   **<big>RsaDecrypt</big>** : 用私钥文件 rsa 解密数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#RsaDecrypt)]
    [[play](https://go.dev/play/p/7_zo6mrx-eX)]
-   **<big>GenerateRsaKeyPair</big>** : 创建 rsa 公钥私钥和 key。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#GenerateRsaKeyPair)]
    [[play](https://go.dev/play/p/sSVmkfENKMz)]
-   **<big>RsaEncryptOAEP</big>** : rsa OAEP 加密。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#RsaEncryptOAEP)]
    [[play](https://go.dev/play/p/sSVmkfENKMz)]
-   **<big>RsaDecryptOAEP</big>** : rsa OAEP 解密。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#RsaDecryptOAEP)]
    [[play](https://go.dev/play/p/sSVmkfENKMz)]
-   **<big>RsaSign</big>** : 应用 RSA 算法签名数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#RsaSign)]
    [[play](https://go.dev/play/p/qhsbf8BJ6Mf)]
-   **<big>RsaVerifySign</big>** : 验证数据的签名是否符合 RSA 算法。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/cryptor.md#RsaVerifySign)]
    [[play](https://go.dev/play/p/qhsbf8BJ6Mf)]

<h3 id="datetime"> 7. datetime日期时间处理包，格式化日期，比较日期。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/datetime"
```

#### 函数列表:

-   **<big>AddDay</big>** : 将日期加/减天数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddDay)]
    [[play](https://go.dev/play/p/dIGbs_uTdFa)]
-   **<big>AddHour</big>** : 将日期加/减小时数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddHour)]
    [[play](https://go.dev/play/p/rcMjd7OCsi5)]
-   **<big>AddMinute</big>** : 将日期加/减分钟数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddMinute)]
    [[play](https://go.dev/play/p/nT1heB1KUUK)]
-   **<big>AddWeek</big>** : 将日期加/减星期数.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddWeek)]
    [[play](https://go.dev/play/p/M9TqdMiaA2p)]
-   **<big>AddMonth</big>** : 将日期加/减月数.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddMonth)]
    [[play](https://go.dev/play/p/DLoiOnpLvsN)]
-   **<big>AddYear</big>** : 将日期加/减分年数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddYear)]
    [[play](https://go.dev/play/p/MqW2ujnBx10)]
-   **<big>AddDaySafe</big>** : 增加/减少指定的天数，并确保日期是有效日期。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddDaySafe)]
    [[play](https://go.dev/play/p/JTohZFpoDJ3)]
-   **<big>AddMonthSafe</big>** : 增加/减少指定的月份，并确保日期是有效日期。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddMonthSafe)]
    [[play](https://go.dev/play/p/KLw0lo6mbVW)]
-   **<big>AddYearSafe</big>** : 增加/减少指定的年份，并确保日期是有效日期。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#AddYearSafe)]
    [[play](https://go.dev/play/p/KVGXWZZ54ZH)]
-   **<big>BeginOfMinute</big>** : 返回指定时间的分钟开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BeginOfMinute)]
    [[play](https://go.dev/play/p/ieOLVJ9CiFT)]
-   **<big>BeginOfHour</big>** : 返回指定时间的小时开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BeginOfHour)]
    [[play](https://go.dev/play/p/GhdGFnDWpYs)]
-   **<big>BeginOfDay</big>** : 返回指定时间的当天开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BeginOfDay)]
    [[play](https://go.dev/play/p/94m_UT6cWs9)]
-   **<big>BeginOfWeek</big>** : 返回指定时间的每周开始时间,默认开始时间星期日。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BeginOfWeek)]
    [[play](https://go.dev/play/p/DCHdcL6gnfV)]
-   **<big>BeginOfMonth</big>** : 返回指定时间的当月开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BeginOfMonth)]
    [[play](https://go.dev/play/p/bWXVFsmmzwL)]
-   **<big>BeginOfYear</big>** : 返回指定时间的当年开始时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BeginOfYear)]
    [[play](https://go.dev/play/p/i326DSwLnV8)]
-   **<big>EndOfMinute</big>** : 返回指定时间的分钟结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#EndOfMinute)]
    [[play](https://go.dev/play/p/yrL5wGzPj4z)]
-   **<big>EndOfHour</big>** : 返回指定时间的小时结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#EndOfHour)]
    [[play](https://go.dev/play/p/6ce3j_6cVqN)]
-   **<big>EndOfDay</big>** : 返回指定时间的当天结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#EndOfDay)]
    [[play](https://go.dev/play/p/eMBOvmq5Ih1)]
-   **<big>EndOfWeek</big>** : 返回指定时间的星期结束时间,默认结束时间星期六。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#EndOfWeek)]
    [[play](https://go.dev/play/p/mGSA162YgX9)]
-   **<big>EndOfMonth</big>** : 返回指定时间的月份结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#EndOfMonth)]
    [[play](https://go.dev/play/p/_GWh10B3Nqi)]
-   **<big>EndOfYear</big>** : 返回指定时间的年份结束时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#EndOfYear)]
    [[play](https://go.dev/play/p/G01cKlMCvNm)]
-   **<big>GetNowDate</big>** : 获取当天日期，返回格式：yyyy-mm-dd。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GetNowDate)]
    [[play](https://go.dev/play/p/PvfkPpcpBBf)]
-   **<big>GetNowTime</big>** : 获取当时时间，返回格式：hh:mm:ss。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GetNowTime)]
    [[play](https://go.dev/play/p/l7BNxCkTmJS)]
-   **<big>GetNowDateTime</big>** : 获取当时日期和时间，返回格式：yyyy-mm-dd hh:mm:ss。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GetNowDateTime)]
    [[play](https://go.dev/play/p/pI4AqngD0al)]
-   **<big>GetTodayStartTime</big>** : 返回当天开始时间， 格式: yyyy-mm-dd 00:00:00。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GetTodayStartTime)]
    [[play](https://go.dev/play/p/84siyYF7t99)]
-   **<big>GetTodayEndTime</big>** : 返回当天结束时间，格式: yyyy-mm-dd 23:59:59。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GetTodayEndTime)]
    [[play](https://go.dev/play/p/jjrLnfoqgn3)]
-   **<big>GetZeroHourTimestamp</big>** : 获取零时时间戳(timestamp of 00:00)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GetZeroHourTimestamp)]
    [[play](https://go.dev/play/p/QmL2oIaGE3q)]
-   **<big>GetNightTimestamp</big>** : 获取午夜时间戳(timestamp of 23:59)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GetNightTimestamp)]
    [[play](https://go.dev/play/p/UolysR3MYP1)]
-   **<big>FormatTimeToStr</big>** : 将日期格式化成字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#FormatTimeToStr)]
    [[play](https://go.dev/play/p/_Ia7M8H_OvE)]
-   **<big>FormatStrToTime</big>** : 将字符串格式化成时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#FormatStrToTime)]
    [[play](https://go.dev/play/p/1h9FwdU8ql4)]
-   **<big>NewUnix</big>** : 创建一个 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#NewUnix)]
    [[play](https://go.dev/play/p/psoSuh_kLRt)]
-   **<big>NewUnixNow</big>** : 创建一个当前时间的 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#NewUnixNow)]
    [[play](https://go.dev/play/p/U4PPx-9D0oz)]
-   **<big>NewFormat</big>** : 创建一个 yyyy-mm-dd hh:mm:ss 格式时间字符串的 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#NewFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>NewISO8601</big>** : 创建一个 iso8601 格式时间字符串的 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#NewISO8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]
-   **<big>ToUnix</big>** : 返回 unix 时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#ToUnix)]
    [[play](https://go.dev/play/p/_LUiwAdocjy)]
-   **<big>ToFormat</big>** : 返回格式'yyyy-mm-dd hh:mm:ss'的日期字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#ToFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>ToFormatForTpl</big>** : 返回 tpl 格式指定的日期字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#ToFormatForTpl)]
    [[play](https://go.dev/play/p/nyXxXcQJ8L5)]
-   **<big>ToIso8601</big>** : 返回 iso8601 日期字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#ToIso8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]
-   **<big>IsLeapYear</big>** :验证是否是闰年。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#IsLeapYear)]
    [[play](https://go.dev/play/p/xS1eS2ejGew)]
-   **<big>BetweenSeconds</big>** : 返回两个时间的间隔秒数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BetweenSeconds)]
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#BetweenSeconds)]
-   **<big>DayOfYear</big>** : 返回参数日期是一年中的第几天。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#DayOfYear)]
    [[play](https://go.dev/play/p/0hjqhTwFNlH)]
-   **<big>IsWeekend</big>** : 判断日期是否是周末。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#IsWeekend)]
    [[play](https://go.dev/play/p/cupRM5aZOIY)]
-   **<big>NowDateOrTime</big>** : 根据指定的格式和时区返回当前时间字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#NowDateOrTime)]
    [[play](https://go.dev/play/p/EZ-begEjtT0)]
-   **<big>Timestamp</big>** : 返回当前秒级时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#Timestamp)]
    [[play](https://go.dev/play/p/iU5b7Vvjx6x)]
-   **<big>TimestampMilli</big>** : 返回当前毫秒级时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#TimestampMilli)]
    [[play](https://go.dev/play/p/4gvEusOTu1T)]
-   **<big>TimestampMicro</big>** : 返回当前微秒级时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#TimestampMicro)]
    [[play](https://go.dev/play/p/2maANglKHQE)]
-   **<big>TimestampNano</big>** : 返回当前纳秒级时间戳。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#TimestampNano)]
    [[play](https://go.dev/play/p/A9Oq_COrcCF)]
-   **<big>TrackFuncTime</big>** : 测试函数执行时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#TrackFuncTime)]
    [[play](https://go.dev/play/p/QBSEdfXHPTp)]
-   **<big>DaysBetween</big>** : 返回两个日期之间的天数差。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#DaysBetween)]
    [[play](https://go.dev/play/p/qD6qGb3TbOy)]
-   **<big>GenerateDatetimesBetween</big>** : 生成从 start 到 end 的所有日期时间的字符串列表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#GenerateDatetimesBetween)]
    [[play](https://go.dev/play/p/6kHBpAxD9ZC)]
-   **<big>Min</big>** : 返回最早时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#Min)]
    [[play](https://go.dev/play/p/MCIDvHNOGGb)]
-   **<big>Max</big>** : 返回最晚时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#Max)]
    [[play](https://go.dev/play/p/9m6JMk1LB7-)]
-   **<big>MaxMin</big>** : 返回最早和最晚时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datetime.md#MaxMin)]
    [[play](https://go.dev/play/p/rbW51cDtM_2)]

<h3 id="datastructure"> 8. datastructure 包含一些普通的数据结构实现。例如：list, linklist, stack, queue, set, tree, graph。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import list "github.com/duke-git/lancet/v2/datastructure/list"
import copyonwritelist "github.com/duke-git/lancet/v2/datastructure/copyonwritelist"
import link "github.com/duke-git/lancet/v2/datastructure/link"
import stack "github.com/duke-git/lancet/v2/datastructure/stack"
import queue "github.com/duke-git/lancet/v2/datastructure/queue"
import set "github.com/duke-git/lancet/v2/datastructure/set"
import tree "github.com/duke-git/lancet/v2/datastructure/tree"
import heap "github.com/duke-git/lancet/v2/datastructure/heap"
import hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
```

#### 函数列表:

-   **<big>List</big>** : 线性表结构, 用切片实现。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/list.md)]
-   **<big>CopyOnWriteList</big>** : 是一个线程安全的 List 实现，底层使用 go 切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/copyonwritelist.md)]
-   **<big>Link</big>** : 链表解构, 包括单链表和双向链表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/link.md)]
-   **<big>Stack</big>** : 栈结构(fifo), 包括数组栈和链表栈。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/stack.md)]
-   **<big>Queue</big>** : 队列结构(filo), 包括数组队列，链表队列，循环队列，优先级队列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/queue.md)]
-   **<big>Set</big>** : 集合（set）结构。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/set.md)]
-   **<big>Tree</big>** : 二叉搜索树。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/tree.md)]
-   **<big>Heap</big>** : 二叉 max 堆。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/heap.md)]
-   **<big>Hashmap</big>** : 哈希映射。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/datastructure/hashmap.md)]

<h3 id="eventbus"> 9. EventbBus是一个事件总线，用于在应用程序中处理事件。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/eventbus"
```

#### 函数列表:

-   **<big>NewEventBus</big>** : 创建 EventBus 实例。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#NewEventBus)]
    [[play](https://go.dev/play/p/gHbOPV_NUOJ)]
-   **<big>Subscribe</big>** : 订阅具有特定事件主题和监听函数的事件。支持异步，事件优先级，事件过滤器。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#Subscribe)]
    [[play](https://go.dev/play/p/EYGf_8cHei-)]
-   **<big>Unsubscribe</big>** : 取消订阅具有特定事件主题的事件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#Unsubscribe)]
    [[play](https://go.dev/play/p/Tmh7Ttfvprf)]
-   **<big>Publish</big>** : 发布一个带有特定事件主题和数据负载的事件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#Publish)]
    [[play](https://go.dev/play/p/gHTtVexFSH9)]
-   **<big>ClearListeners</big>** : 清空所有事件监听器。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#ClearListeners)]
    [[play](https://go.dev/play/p/KBfBYlKPgqD)]
-   **<big>ClearListenersByTopic</big>** : 清空特定事件监听器。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#ClearListenersByTopic)]
    [[play](https://go.dev/play/p/gvMljmJOZmU)]
-   **<big>GetListenersCount</big>** : 获取特定事件的监听器数量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#GetListenersCount)]
    [[play](https://go.dev/play/p/8VPJsMQgStM)]
-   **<big>GetAllListenersCount</big>** : 获取所有事件的监听器数量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#GetAllListenersCount)]
    [[play](https://go.dev/play/p/PUlr0xcpEOz)]
-   **<big>GetEvents</big>** : 获取所有事件的 topic。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#GetEvents)]
    [[play](https://go.dev/play/p/etgjjcOtAjX)]
-   **<big>SetErrorHandler</big>** : 设置事件的错误处理函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/eventbus.md#SetErrorHandler)]
    [[play](https://go.dev/play/p/gmB0gnFe5mc)]

<h3 id="enum"> 10. Enum实现一个简单枚举工具包。. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">Index</a></h3>

```go
import "github.com/duke-git/lancet/v2/enum"
```

#### Function list:

-   **<big>NewItem</big>** : 创建枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#NewItem)]
    [[play](https://go.dev/play/p/8qNsLw01HD5)]
-   **<big>NewItemsFromPairs</big>** : 从 Pair 结构体的切片创建枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#NewItemsFromPairs)]
    [[play](https://go.dev/play/p/xKnoGa7gnev)]
-   **<big>Value</big>** : 返回枚举项的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Value)]
    [[play](https://go.dev/play/p/xKnoGa7gnev)]
-   **<big>Name</big>** : 返回枚举项的名称。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Name)]
    [[play](https://go.dev/play/p/xKnoGa7gnev)]
-   **<big>Valid</big>** : 检查枚举项是否有效。如果提供了自定义检查函数，将使用该函数验证值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Valid)]
    [[play](https://go.dev/play/p/pA3lYY2VSm3)]
-   **<big>MarshalJSON</big>** : 枚举项实现 json.Marshaler 接口。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#MarshalJSON)]
    [[play](https://go.dev/play/p/zIZEdAnneB5)]
-   **<big>NewRegistry</big>** : Registry 定义了一个通用的枚举注册表结构体。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#NewRegistry)]
    [[play](https://go.dev/play/p/ABEXsYfJKMo)]
-   **<big>Add</big>** : 向枚举注册表添加枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Add)]
    [[play](https://go.dev/play/p/ABEXsYfJKMo)]
-   **<big>Remove</big>** : 在枚举注册表中删除枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Remove)]
    [[play](https://go.dev/play/p/dSG84wQ3TuC)]
-   **<big>Update</big>** : 在枚举注册表中更新枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Update)]
    [[play](https://go.dev/play/p/Ol0moT1J9Xl)]
-   **<big>GetByValue</big>** : 在枚举注册表中通过值获取枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#GetByValue)]
    [[play](https://go.dev/play/p/niJ1U2KlE_m)]
-   **<big>GetByName</big>** : 在枚举注册表中通过名称获取枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#GetByName)]
    [[play](https://go.dev/play/p/49ie_gpqH0m)]
-   **<big>Items</big>** : 返回枚举注册表中的枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Items)]
    [[play](https://go.dev/play/p/lAJFAradbvQ)]
-   **<big>Contains</big>** : 检查注册表中是否存在具有给定值的枚举项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Contains)]
    [[play](https://go.dev/play/p/_T-lPYkZn2j)]
-   **<big>Size</big>** : 返回注册表中枚举项的数目。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Size)]
    [[play](https://go.dev/play/p/TeDArWhlQe2)]
-   **<big>Range</big>** : 遍历注册表中的所有枚举项，并应用给定的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Range)]
    [[play](https://go.dev/play/p/GPsZbQbefWN)]
-   **<big>SortedItems</big>** : 返回按给定比较函数排序的所有枚举项的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#SortedItems)]
    [[play](https://go.dev/play/p/tN9RE_m_WEI)]
-   **<big>Filter</big>** : 返回满足给定谓词函数的枚举项切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/enum.md#Filter)]
    [[play](https://go.dev/play/p/uTUpTdcyoCU)]

<h3 id="fileutil"> 11. fileutil 包含文件基本操作。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/fileutil"
```

#### 函数列表：

-   **<big>ClearFile</big>** : 清空文件内容。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ClearFile)]
    [[play](https://go.dev/play/p/NRZ0ZT-G94H)]
-   **<big>CreateFile</big>** : 创建文件，创建成功返回 true, 否则返回 false。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#CreateFile)]
    [[play](https://go.dev/play/p/lDt8PEsTNKI)]
-   **<big>CreateDir</big>** : 创建嵌套目录，例如/a/, /a/b/。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#CreateDir)]
    [[play](https://go.dev/play/p/qUuCe1OGQnM)]
-   **<big>CopyFile</big>** : 拷贝文件，会覆盖原有的文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#CopyFile)]
    [[play](https://go.dev/play/p/Jg9AMJMLrJi)]
-   **<big>CopyDir</big>** : 拷贝目录。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#CopyDir)]
    [[play](https://go.dev/play/p/YAyFTA_UuPb)]
-   **<big>FileMode</big>** : 获取文件 mode 信息。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#FileMode)]
    [[play](https://go.dev/play/p/2l2hI42fA3p)]
-   **<big>MiMeType</big>** : 获取文件 mime 类型, 参数的类型必须是 string 或者\*os.File。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#MiMeType)]
    [[play](https://go.dev/play/p/bd5sevSUZNu)]
-   **<big>IsExist</big>** : 判断文件或目录是否存在。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#IsExist)]
    [[play](https://go.dev/play/p/nKKXt8ZQbmh)]
-   **<big>IsLink</big>** : 判断文件是否是符号链接。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#IsLink)]
    [[play](https://go.dev/play/p/TL-b-Kzvf44)]
-   **<big>IsDir</big>** : 判断参数是否是目录。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#IsDir)]
    [[play](https://go.dev/play/p/WkVwEKqtOWk)]
-   **<big>ListFileNames</big>** : 返回目录下所有文件名。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ListFileNames)]
    [[play](https://go.dev/play/p/Tjd7Y07rejl)]
-   **<big>RemoveFile</big>** : 删除文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#RemoveFile)]
    [[play](https://go.dev/play/p/P2y0XW8a1SH)]
-   **<big>RemoveDir</big>** : 删除目录，支持传入删除前的回调函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#RemoveDir)]
    [[play](https://go.dev/play/p/Oa6KnPek2uy)]
-   **<big>ReadFileToString</big>** : 读取文件内容并返回字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ReadFileToString)]
    [[play](https://go.dev/play/p/cmfwp_5SQTp)]
-   **<big>ReadFileByLine</big>** : 按行读取文件内容，返回字符串切片包含每一行。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ReadFileByLine)]
    [[play](https://go.dev/play/p/svJP_7ZrBrD)]
-   **<big>Zip</big>** : zip 压缩文件, 参数可以是文件或目录。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#Zip)]
    [[play](https://go.dev/play/p/j-3sWBp8ik_P)]
-   **<big>ZipAppendEntry</big>** : 通过将单个文件或目录追加到现有的 zip 文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ZipAppendEntry)]
-   **<big>UnZip</big>** : zip 解压缩文件并保存在目录中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#UnZip)]
    [[play](https://go.dev/play/p/g0w34kS7B8m)]
-   **<big>CurrentPath</big>** : 返回当前位置的绝对路径。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#CurrentPath)]
    [[play](https://go.dev/play/p/s74a9iBGcSw)]
-   **<big>IsZipFile</big>** : 判断文件是否是 zip 压缩文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#IsZipFile)]
    [[play](https://go.dev/play/p/9M0g2j_uF_e)]
-   **<big>FileSize</big>** : 返回文件字节大小。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#FileSize)]
    [[play](https://go.dev/play/p/H9Z05uD-Jjc)]
-   **<big>MTime</big>** : 返回文件修改时间(unix timestamp)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#MTime)]
    [[play](https://go.dev/play/p/s_Tl7lZoAaY)]
-   **<big>Sha</big>** : 返回文件 sha 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#Sha)]
    [[play](https://go.dev/play/p/VfEEcO2MJYf)]
-   **<big>ReadCsvFile</big>** : 读取 csv 文件内容到切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ReadCsvFile)]
    [[play](https://go.dev/play/p/OExTkhGEd3_u)]
-   **<big>WriteCsvFile</big>** : 向 csv 文件写入切片数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#WriteCsvFile)]
-   **<big>WriteMapsToCsv</big>** : 将 map 切片写入 csv 文件中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#WriteMapsToCsv)]
    [[play](https://go.dev/play/p/umAIomZFV1c)]
-   **<big>WriteBytesToFile</big>** : 将 bytes 写入文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#WriteBytesToFile)]
    [[play](https://go.dev/play/p/s7QlDxMj3P8)]
-   **<big>WriteStringToFile</big>** : 将字符串写入文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#WriteStringToFile)]
    [[play](https://go.dev/play/p/GhLS6d8lH_g)]
-   **<big>ReadFile</big>** : 读取文件或者 URL。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ReadFile)]
-   **<big>ChunkRead</big>** : 从文件的指定偏移读取块并返回块内所有行。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ChunkRead)]
    [[play](https://go.dev/play/p/r0hPmKWhsgf)]
-   **<big>ParallelChunkRead</big>** : 并行读取文件并将每个块的行发送到指定通道。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#ParallelChunkRead)]
    [[play](https://go.dev/play/p/teMXnCsdSEw)]
-   **<big>GetExeOrDllVersion</big>** : 返回 exe,dll 文件版本号(仅 Window 平台)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/fileutil.md#GetExeOrDllVersion)]
    [[play](https://go.dev/play/p/iLRrDBhE38E)]

<h3 id="formatter"> 12. formatter 格式化器包含一些数据格式化处理方法。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/formatter"
```

#### 函数列表:

-   **<big>Comma</big>** : 用逗号每隔 3 位分割数字/字符串，支持添加前缀符号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#Comma)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]
-   **<big>Pretty</big>** : 返回 pretty JSON 字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#Pretty)]
    [[play](https://go.dev/play/p/YsciGj3FH2x)]
-   **<big>PrettyToWriter</big>** : Pretty encode 数据到 writer。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#PrettyToWriter)]
    [[play](https://go.dev/play/p/LPLZ3lDi5ma)]
-   **<big>DecimalBytes</big>** : 返回十进制标准（以 1000 为基数）下的可读字节单位字符串。precision 参数指定小数点后的位数，默认为 4。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#DecimalBytes)]
    [[play](https://go.dev/play/p/FPXs1suwRcs)]
-   **<big>BinaryBytes</big>** : 返回 binary 标准（以 1024 为基数）下的可读字节单位字符串。precision 参数指定小数点后的位数，默认为 4。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#BinaryBytes)]
    [[play](https://go.dev/play/p/G9oHHMCAZxP)]
-   **<big>ParseDecimalBytes</big>** : 将字节单位字符串转换成其所表示的字节数（以 1000 为基数）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#ParseDecimalBytes)]
    [[play](https://go.dev/play/p/Am98ybWjvjj)]
-   **<big>ParseBinaryBytes</big>** : 将字节单位字符串转换成其所表示的字节数（以 1024 为基数）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#ParseBinaryBytes)]
    [[play](https://go.dev/play/p/69v1tTT62x8)]
-   **<big>ParseCNAddress</big>** : 智能解析中国地址字符串并提取结构化信息（省、市、区、街道、姓名、电话等）。支持多种地址格式，包括县级市。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#ParseCNAddress)]
    [[play](https://go.dev/play/p/o5l09hQopEV)]
-   **<big>ParsePersonInfo</big>** : 从地址字符串中提取用户信息（姓名、电话、身份证、邮编）并将其与位置地址分离。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/formatter.md#ParsePersonInfo)]
    [[play](https://go.dev/play/p/JO-uTlJlTy7)]

<h3 id="function"> 13. function 函数包控制函数执行流程，包含部分函数式编程。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/function"
```

#### 函数列表:

-   **<big>After</big>** : 创建一个函数，当该函数被调用 n 或更多次之后将执行传入的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#After)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]
-   **<big>Before</big>** : 创建一个函数，当该函数被调用不超过 n 次时，将执行执行传入的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Before)]
    [[play](https://go.dev/play/p/0HqUDIFZ3IL)]
-   **<big>CurryFn</big>** : 创建柯里化函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#CurryFn)]
    [[play](https://go.dev/play/p/5HopfDwANKX)]
-   **<big>Compose</big>** : 从右至左组合函数列表 fnList，返回组合后的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Compose)]
    [[play](https://go.dev/play/p/KKfugD4PKYF)]
-   **<big>Delay</big>** : 延迟 delay 时间后调用函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Delay)]
    [[play](https://go.dev/play/p/Ivtc2ZE-Tye)]
-   **<big>Debounced<sup>deprecated</sup></big>** : 创建一个 debounced 函数，该函数延迟调用 fn 直到自上次调用 debounced 函数后等待持续时间过去。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Debounced)]
    [[play](https://go.dev/play/p/absuEGB_GN7)]
-   **<big>Debounce</big>** : 创建一个函数的去抖动版本。该去抖动函数仅在上次调用后的指定延迟时间过去之后才会调用原始函数。支持取消去抖动。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Debounce)]
    [[play](https://go.dev/play/p/-dGFrYn_1Zi)]
-   **<big>Throttle</big>** : 创建一个函数的节流版本。返回的函数保证在每个时间间隔内最多只会被调用一次。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Throttle)]
    [[play](https://go.dev/play/p/HpoMov-tJSN)]
-   **<big>Schedule</big>** : 每次持续时间调用函数，直到关闭返回的 channel。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Schedule)]
    [[play](https://go.dev/play/p/hbON-Xeyn5N)]
-   **<big>Pipeline</big>** : 从右至左执行函数列表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Pipeline)]
    [[play](https://go.dev/play/p/mPdUVvj6HD6)]
-   **<big>AcceptIf</big>** : AcceptIf 函数会返回另一个函数，该函数的签名与 apply 函数相同，但同时还会包含一个布尔值来表示成功或失败。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#AcceptIf)]
    [[play](https://go.dev/play/p/XlXHHtzCf7d)]
-   **<big>And</big>** : 返回一个复合谓词判断函数，该判断函数表示一组谓词的逻辑 and 操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#And)]
    [[play](https://go.dev/play/p/dTBHJMQ0zD2)]
-   **<big>Or</big>** : 返回一个复合谓词判断函数，该判断函数表示一组谓词的逻辑 or 操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Or)]
    [[play](https://go.dev/play/p/LitCIsDFNDA)]
-   **<big>Negate</big>** : 返回一个谓词函数，该谓词函数表示当前谓词的逻辑否定。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Negate)]
    [[play](https://go.dev/play/p/jbI8BtgFnVE)]
-   **<big>Nor</big>** : 返回一个复合谓词判断函数，该判断函数表示一组谓词的逻辑非或 nor 的操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Nor)]
    [[play](https://go.dev/play/p/2KdCoBEOq84)]
-   **<big>Nand</big>** : 返回一个复合谓词判断函数，该判断函数表示一组谓词的逻辑非与 nand 的操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Nand)]
    [[play](https://go.dev/play/p/Rb-FdNGpgSO)]
-   **<big>Xnor</big>** : 返回一个复合谓词判断函数，该判断函数表示一组谓词的逻辑异或 xnor 的操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Xnor)]
    [[play](https://go.dev/play/p/FJxko8SFbqc)]
-   **<big>Watcher</big>** : Watcher 用于记录代码执行时间。可以启动/停止/重置手表定时器。获取函数执行的时间。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/function.md#Watcher)]
    [[play](https://go.dev/play/p/l2yrOpCLd1I)]

<h3 id="maputil"> 14. maputil 包括一些操作 map 的函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/maputil"
```

#### 函数列表:

-   **<big>MapTo</big>** : 快速将 map 或者其他类型映射到结构体或者指定类型。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#MapTo)]
    [[play](https://go.dev/play/p/4K7KBEPgS5M)]
-   **<big>ForEach</big>** : 对 map 中的每对 key 和 value 执行 iteratee 函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ForEach)]
    [[play](https://go.dev/play/p/OaThj6iNVXK)]
-   **<big>Filter</big>** : 迭代 map 中的每对 key 和 value，返回 map，其中的 key 和 value 符合 predicate 函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Filter)]
    [[play](https://go.dev/play/p/fSvF3wxuNG7)]
-   **<big>FilterByKeys</big>** : 迭代 map, 返回一个新 map，其 key 都是给定的 key 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil#FilterByKeys)]
    [[play](https://go.dev/play/p/7ov6BJHbVqh)]
-   **<big>FilterByValues</big>** : 迭代 map, 返回一个新 map，其 value 都是给定的 value 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil#FilterByValues)]
    [[play](https://go.dev/play/p/P3-9MdcXegR)]
-   **<big>OmitBy</big>** : Filter 的反向操作, 迭代 map 中的每对 key 和 value, 删除符合 predicate 函数的 key, value, 返回新 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil#OmitBy)]
    [[play](https://go.dev/play/p/YJM4Hj5hNwm)]
-   **<big>OmitByKeys</big>** : FilterByKeys 的反向操作, 迭代 map, 返回一个新 map，其 key 不包括给定的 key 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil#OmitByKeys)]
    [[play](https://go.dev/play/p/jXGrWDBfSRp)]
-   **<big>OmitByValues</big>** : FilterByValues 的反向操作, 迭代 map, 返回一个新 map，其 value 不包括给定的 value 值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil#OmitByValues)]
    [[play](https://go.dev/play/p/XB7Y10uw20_U)]
-   **<big>Intersect</big>** : 多个 map 的交集操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Intersect)]
    [[play](https://go.dev/play/p/Zld0oj3sjcC)]
-   **<big>Keys</big>** : 返回 map 中所有 key 组成的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Keys)]
    [[play](https://go.dev/play/p/xNB5bTb97Wd)]
-   **<big>KeysBy</big>** : 创建一个切片，其元素是每个 map 的 key 调用 mapper 函数的结果。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil#KeysBy)]
    [[play](https://go.dev/play/p/hI371iB8Up8)]
-   **<big>Merge</big>** : 合并多个 map, 相同的 key 会被之后的 key 覆盖。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Merge)]
    [[play](https://go.dev/play/p/H95LENF1uB-)]
-   **<big>Minus</big>** : 返回一个 map，其中的 key 存在于 mapA，不存在于 mapB。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Minus)]
    [[play](https://go.dev/play/p/3u5U9K7YZ9m)]
-   **<big>Values</big>** : 返回 map 中所有 values 组成的切片
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Values)]
    [[play](https://go.dev/play/p/CBKdUc5FTW6)]
-   **<big>ValuesBy</big>** : 创建一个切片，其元素是每个 map 的 value 调用 mapper 函数的结果。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ValuesBy)]
    [[play](https://go.dev/play/p/sg9-oRidh8f)]
-   **<big>MapKeys</big>** : 操作 map 的每个 key，然后转为新的 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#MapKeys)]
    [[play](https://go.dev/play/p/8scDxWeBDKd)]
-   **<big>MapValues</big>** : 操作 map 的每个 value，然后转为新的 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#MapValues)]
    [[play](https://go.dev/play/p/g92aY3fc7Iw)]
-   **<big>Entries</big>** : 将 map 转换为键/值对切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Entries)]
    [[play](https://go.dev/play/p/Ltb11LNcElY)]
-   **<big>FromEntries</big>** : 基于键/值对的切片创建 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#FromEntries)]
    [[play](https://go.dev/play/p/fTdu4sCNjQO)]
-   **<big>Transform</big>** : 将 map 转换为其他类型的 map。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#Transform)]
    [[play](https://go.dev/play/p/P6ovfToM3zj)]
-   **<big>IsDisjoint</big>** : 验证两个 map 是否具有不同的 key。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#IsDisjoint)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
-   **<big>HasKey</big>** : 检查 map 是否包含某个 key。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#HasKey)]
    [[play](https://go.dev/play/p/isZZHOsDhFc)]
-   **<big>GetOrSet</big>** : 返回给定键的值，如果不存在则设置该值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#GetOrSet)]
    [[play](https://go.dev/play/p/IVQwO1OkEJC)]
-   **<big>MapToStruct</big>** : 将 map 转成 struct。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#MapToStruct)]
    [[play](https://go.dev/play/p/7wYyVfX38Dp)]
-   **<big>ToSortedSlicesDefault</big>** : 将 map 的 key 和 value 转化成两个根据 key 的值从小到大排序的切片，value 切片中元素的位置与 key 对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ToSortedSlicesDefault)]
    [[play](https://go.dev/play/p/43gEM2po-qy)]
-   **<big>ToSortedSlicesWithComparator</big>** : 将 map 的 key 和 value 转化成两个使用比较器函数根据 key 的值自定义排序规则的切片，value 切片中元素的位置与 key 对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ToSortedSlicesWithComparator)]
    [[play](https://go.dev/play/p/0nlPo6YLdt3)]
-   **<big>NewOrderedMap</big>** : 创建有序映射。有序映射是键值对的集合，其中键是唯一的，并且保留键插入的顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#NewOrderedMap)]
    [[play](https://go.dev/play/p/Y4ZJ_oOc1FU)]
-   **<big>OrderedMap_Set</big>** : 设置给定的键值对。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Set)]
    [[play](https://go.dev/play/p/Y4ZJ_oOc1FU)]
-   **<big>OrderedMap_Get</big>** : 返回给定键的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Get)]
    [[play](https://go.dev/play/p/Y4ZJ_oOc1FU)]
-   **<big>OrderedMap_Delete</big>** : 删除给定键的键值对。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Delete)]
    [[play](ttps://go.dev/play/p/5bIi4yaZ3K-)]
-   **<big>OrderedMap_Clear</big>** : 清空 map 数据。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Clear)]
    [[play](https://go.dev/play/p/8LwoJyEfuFr)]
-   **<big>OrderedMap_Front</big>** : 返回第一个键值对。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Front)]
    [[play](https://go.dev/play/p/ty57XSimpoe)]
-   **<big>OrderedMap_Back</big>** : 返回最后一个键值对。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Back)]
    [[play](https://go.dev/play/p/rQMjp1yQmpa)]
-   **<big>OrderedMap_Range</big>** : 为每个键值对调用给定的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Range)]
    [[play](https://go.dev/play/p/U-KpORhc7LZ)]
-   **<big>OrderedMap_Keys</big>** : 按顺序返回键的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Keys)]
    [[play](https://go.dev/play/p/Vv_y9ExKclA)]
-   **<big>OrderedMap_Values</big>** : 按顺序返回值的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Values)]
    [[play](https://go.dev/play/p/TWj5n1-PUfx)]
-   **<big>OrderedMap_Elements</big>** : 按顺序返回键值对。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Elements)]
    [[play](https://go.dev/play/p/4BHG4kKz6bB)]
-   **<big>OrderedMap_Len</big>** : 返回键值对的数量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Len)]
    [[play](https://go.dev/play/p/cLe6z2VX5N-)]
-   **<big>OrderedMap_Contains</big>** : 如果给定的键存在则返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Contains)]
    [[play](https://go.dev/play/p/QuwqqnzwDNX)]
-   **<big>OrderedMap_Iter</big>** : 返回按顺序产生键值对的通道。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_Iter)]
    [[play](https://go.dev/play/p/tlq2tdvicPt)]
-   **<big>OrderedMap_ReverseIter</big>** : 返回以相反顺序产生键值对的通道。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_ReverseIter)]
    [[play](https://go.dev/play/p/8Q0ssg6hZzO)]
-   **<big>OrderedMap_SortByKey</big>** : 使用传入的比较函数排序 map key。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_SortByKey)]
    [[play](https://go.dev/play/p/N7hjD_alZPq)]
-   **<big>OrderedMap_MarshalJSON</big>** : 实现 json.Marshaler 接口。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_MarshalJSON)]
    [[play](https://go.dev/play/p/C-wAwydIAC7)]
-   **<big>OrderedMap_UnmarshalJSON</big>** : 实现 json.Unmarshaler 接口。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#OrderedMap_UnmarshalJSON)]
    [[play](https://go.dev/play/p/8C3MvJ3-mut)]
-   **<big>NewConcurrentMap</big>** : ConcurrentMap 协程安全的 map 结构。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#NewConcurrentMap)]
    [[play](https://go.dev/play/p/3PenTPETJT0)]
-   **<big>ConcurrentMap_Set</big>** : 在 map 中设置 key 和 value。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ConcurrentMap_Set)]
    [[play](https://go.dev/play/p/3PenTPETJT0)]
-   **<big>ConcurrentMap_Get</big>** : 根据 key 获取 value, 如果不存在 key,返回零值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ConcurrentMap_Get)]
    [[play](https://go.dev/play/p/3PenTPETJT0)]
-   **<big>ConcurrentMap_GetOrSet</big>** : 返回键的现有值（如果存在），否则，设置 key 并返回给定值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ConcurrentMap_GetOrSet)]
    [[play](https://go.dev/play/p/aDcDApOK01a)]
-   **<big>ConcurrentMap_Delete</big>** : 删除 key。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ConcurrentMap_Delete)]
    [[play](https://go.dev/play/p/uTIJZYhpVMS)]
-   **<big>ConcurrentMap_GetAndDelete</big>** :获取 key，然后删除。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ConcurrentMap_GetAndDelete)]
    [[play](https://go.dev/play/p/ZyxeIXSZUiM)]
-   **<big>ConcurrentMap_Has</big>** : 验证是否包含 key。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ConcurrentMap_Has)]
    [[play](https://go.dev/play/p/C8L4ul9TVwf)]
-   **<big>ConcurrentMap_Range</big>** : 为 map 中每个键和值顺序调用迭代器。 如果 iterator 返回 false，则停止迭代。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ConcurrentMap_Range)]
    [[play](https://go.dev/play/p/iqcy7P8P0Pr)]
-   **<big>SortByKey</big>** : 对传入的 map 根据 key 进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#SortByKey)]
    [[play](https://go.dev/play/p/PVdmBSnm6P_W)]
-   **<big>GetOrDefault</big>** : 返回给定键的值，如果键不存在，则返回默认值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#GetOrDefault)]
    [[play](https://go.dev/play/p/99QjSYSBdiM)]
-   **<big>FindValuesBy</big>** : 返回一个切片，包含满足给定谓词判断函数的 map 中的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#FindValuesBy)]
    [[play](https://go.dev/play/p/bvNwNBZDm6v)]
-   **<big>ToMarkdownTable</big>** : 将一个 map 切片数据转换为 Markdown 表格字符串。支持自定义表头显示名称和列的显示顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/maputil.md#ToMarkdownTable)]
    [[play](https://go.dev/play/p/w_pSLfeyEB5)]

<h3 id="mathutil"> 15. mathutil 包实现了一些数学计算的函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### 函数列表:

-   **<big>Average</big>** :计算平均数，可能需要对结果调用 RoundToFloat 方法四舍五入。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Average)]
    [[play](https://go.dev/play/p/Vv7LBwER-pz)]
-   **<big>Exponent</big>** : 指数计算（x 的 n 次方）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Exponent)]
    [[play](https://go.dev/play/p/uF3HGNPk8wr)]
-   **<big>Fibonacci</big>** :计算斐波那契数列的第 n 个数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Fibonacci)]
    [[play](https://go.dev/play/p/IscseUNMuUc)]
-   **<big>Factorial</big>** : 计算阶乘。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Factorial)]
    [[play](https://go.dev/play/p/tt6LdOK67Nx)]
-   **<big>Max</big>** : 返回参数中的最大数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Max)]
    [[play](https://go.dev/play/p/cN8DHI0rTkH)]
-   **<big>MaxBy</big>** : 使用给定的比较器函数返回切片的最大值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#MaxBy)]
    [[play](https://go.dev/play/p/pbe2MT-7DV2)]
-   **<big>Min</big>** : 返回参数中的最小数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Min)]
    [[play](https://go.dev/play/p/21BER_mlGUj)]
-   **<big>MinBy</big>** : 使用给定的比较器函数返回切片的最小值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#MinBy)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
-   **<big>Percent</big>** : 计算百分比，可以指定保留 n 位小数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Percent)]
    [[play](https://go.dev/play/p/s0NdFCtwuyd)]
-   **<big>RoundToFloat</big>** : 四舍五入，保留 n 位小数，返回 float64。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#RoundToFloat)]
    [[play](https://go.dev/play/p/ghyb528JRJL)]
-   **<big>RoundToString</big>** : 四舍五入，保留 n 位小数，返回 string。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#RoundToString)]
    [[play](https://go.dev/play/p/kZwpBRAcllO)]
-   **<big>TruncRound</big>** : 截短 n 位小数（不进行四舍五入）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#TruncRound)]
    [[play](https://go.dev/play/p/aumarSHIGzP)]
-   **<big>CeilToFloat</big>** : 向上舍入（进一法），保留 n 位小数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#CeilToFloat)]
    [[play](https://go.dev/play/p/8hOeSADZPCo)]
-   **<big>CeilToString</big>** : 向上舍入（进一法），保留 n 位小数，返回字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#CeilToString)]
    [[play](https://go.dev/play/p/wy5bYEyUKKG)]
-   **<big>FloorToFloat</big>** : 向下舍入（去尾法），保留 n 位小数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#FloorToFloat)]
    [[play](https://go.dev/play/p/vbCBrQHZEED)]
-   **<big>FloorToString</big>** : 向下舍入（去尾法），保留 n 位小数，返回字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#FloorToString)]
    [[play](https://go.dev/play/p/Qk9KPd2IdDb)]
-   **<big>Range</big>** : 根据指定的起始值和数量，创建一个数字切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Range)]
    [[play](https://go.dev/play/p/9ke2opxa8ZP)]
-   **<big>RangeWithStep</big>** : 根据指定的起始值，结束值，步长，创建一个数字切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#RangeWithStep)]
    [[play](https://go.dev/play/p/akLWz0EqOSM)]
-   **<big>AngleToRadian</big>** : 将角度值转为弧度值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#AngleToRadian)]
    [[play](https://go.dev/play/p/CIvlICqrHql)]
-   **<big>RadianToAngle</big>** : 将弧度值转为角度值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#RadianToAngle)]
    [[play](https://go.dev/play/p/dQtmOTUOMgi)]
-   **<big>PointDistance</big>** : 计算两个坐标点的距离。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#PointDistance)]
    [[play](https://go.dev/play/p/RrG4JIaziM8)]
-   **<big>IsPrime</big>** : 判断质数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#IsPrime)]
    [[play](https://go.dev/play/p/Rdd8UTHZJ7u)]
-   **<big>GCD</big>** : 求最大公约数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#GCD)]
    [[play](https://go.dev/play/p/CiEceLSoAKB)]
-   **<big>LCM</big>** : 求最小公倍数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#LCM)]
    [[play](https://go.dev/play/p/EjcZxfY7G_g)]
-   **<big>Cos</big>** : 计算弧度的余弦值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Cos)]
    [[play](https://go.dev/play/p/Sm89LoIfvFq)]
-   **<big>Sin</big>** : 计算弧度的正弦值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Sin)]
    [[play](https://go.dev/play/p/TWMQlMywDsP)]
-   **<big>Log</big>** : 计算以 base 为底 n 的对数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Log)]
    [[play](https://go.dev/play/p/_d4bi8oyhat)]
-   **<big>Sum</big>** : 求传入参数之和。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Sum)]
    [[play](https://go.dev/play/p/1To2ImAMJA7)]
-   **<big>Abs</big>** : 求绝对值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Sum)]
    [[play](https://go.dev/play/p/fsyBh1Os-1d)]
-   **<big>Div</big>** : 除法运算。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Div)]
    [[play](https://go.dev/play/p/WLxDdGXXYat)]
-   **<big>Variance</big>** : 计算方差。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Variance)]
    [[play](https://go.dev/play/p/uHuV4YgXf8F)]
-   **<big>StdDev</big>** : 计算标准差。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#StdDev)]
    [[play](https://go.dev/play/p/FkNZDXvHD2l)]
-   **<big>Permutation</big>** : 计算排列数 P(n, k)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Permutation)]
    [[play](https://go.dev/play/p/MgobwH_FOxj)]
-   **<big>Combination</big>** : 计算组合数 C(n, k)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/mathutil.md#Combination)]
    [[play](https://go.dev/play/p/ENFQRDQUFi9)]

<h3 id="netutil"> 16. netutil 网络包支持获取 ip 地址，发送 http 请求。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/netutil"
```

#### 函数列表:

-   **<big>ConvertMapToQueryString</big>** : 将 map 转换成 http 查询字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#ConvertMapToQueryString)]
    [[play](https://go.dev/play/p/jnNt_qoSnRi)]
-   **<big>EncodeUrl</big>** : 编码 url query string 的值(?a=1&b=[2] -> ?a=1&b=%5B2%5D)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#EncodeUrl)]
    [[play](https://go.dev/play/p/bsZ6BRC4uKI)]
-   **<big>GetInternalIp</big>** : 获取内部 ipv4。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#GetInternalIp)]
    [[play](https://go.dev/play/p/5mbu-gFp7ei)]
-   **<big>GetIps</big>** : 获取系统 ipv4 地址列表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#GetIps)]
    [[play](https://go.dev/play/p/NUFfcEmukx1)]
-   **<big>GetMacAddrs</big>** : 获取系统 mac 地址列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#GetMacAddrs)]
    [[play](https://go.dev/play/p/Rq9UUBS_Xp1)]
-   **<big>GetPublicIpInfo</big>** : 获取[公网 ip 信息](http://ip-api.com/json/).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#GetPublicIpInfo)]
    [[play](https://go.dev/play/p/YDxIfozsRHR)]
-   **<big>GetRequestPublicIp</big>** : 获取 http 请求 ip。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#GetRequestPublicIp)]
    [[play](https://go.dev/play/p/kxU-YDc_eBo)]
-   **<big>IsPublicIP</big>** : 判断 ip 是否是公共 ip。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#IsPublicIP)]
    [[play](https://go.dev/play/p/nmktSQpJZnn)]
-   **<big>IsInternalIP</big>** : 判断 ip 是否是局域网 ip。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#IsInternalIP)]
    [[play](https://go.dev/play/p/sYGhXbgO4Cb)]
-   **<big>HttpRequest</big>** : 用于抽象 HTTP 请求实体的结构。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#HttpRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>HttpClient</big>** : 用于发送 HTTP 请求。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#HttpClient)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>SendRequest</big>** : 发送 http 请求。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#SendRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>DecodeResponse</big>** : 解析 http 响应体到目标结构体。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#DecodeResponse)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>StructToUrlValues</big>** : 将结构体转为 url values, 仅转化结构体导出字段并且包含`json` tag。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#StructToUrlValues)]
    [[play](https://go.dev/play/p/pFqMkM40w9z)]
-   **<big>HttpGet<sup>deprecated</sup></big>** : 发送 http get 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#HttpGet)]
-   **<big>HttpDelete<sup>deprecated</sup></big>** : 发送 http delete 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#HttpDelete)]
-   **<big>HttpPost<sup>deprecated</sup></big>** : 发送 http post 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#HttpPost)]
-   **<big>HttpPut<sup>deprecated</sup></big>** : 发送 http put 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#HttpPut)]
-   **<big>HttpPatch<sup>deprecated</sup></big>** : 发送 http patch 请求（已弃用：SendRequest 代替）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#HttpPatch)]
-   **<big>ParseHttpResponse</big>** : 解析 http 响应体到目标结构体。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#ParseHttpResponse)]
-   **<big>DownloadFile</big>** : 从指定的 server 地址下载文件。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#DownloadFile)]
-   **<big>UploadFile</big>** : 将文件上传指定的 server 地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#UploadFile)]
-   **<big>IsPingConnected</big>** : 检查能否 ping 通主机。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#IsPingConnected)]
    [[play](https://go.dev/play/p/q8OzTijsA87)]
-   **<big>IsTelnetConnected</big>** : 检查能否 telnet 到主机。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#IsTelnetConnected)]
    [[play](https://go.dev/play/p/yiLCGtQv_ZG)]
-   **<big>BuildUrl</big>** : 创建 url 字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#BuildUrl)]
    [[play](https://go.dev/play/p/JLXl1hZK7l4)]
-   **<big>AddQueryParams</big>** : 向 url 添加查询参数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/netutil.md#AddQueryParams)]
    [[play](https://go.dev/play/p/JLXl1hZK7l4)]

<h3 id="pointer"> 17. pointer 包支持一些指针类型的操作。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/pointer"
```

#### 函数列表:

-   **<big>ExtractPointer</big>** : 返回传入 interface 的底层值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/pointer.md#ExtractPointer)]
    [[play](https://go.dev/play/p/D7HFjeWU2ZP)]
-   **<big>Of</big>** : 返回传入参数的指针值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/pointer.md#Of)]
    [[play](https://go.dev/play/p/HFd70x4DrMj)]
-   **<big>Unwrap</big>** : 返回传入指针指向的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/pointer.md#Unwrap)]
    [[play](https://go.dev/play/p/cgeu3g7cjWb)
-   **<big>UnwarpOr</big>** : 返回指针的值，如果指针为零值，则返回 fallback。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/pointer.md#UnwrapOr)]
    [[play](https://go.dev/play/p/mmNaLC38W8C)]
-   **<big>UnwarpOrDefault</big>** : 返回指针的值，如果指针为零值，则返回相应零值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/pointer.md#UnwrapOrDefault)]
    [[play](https://go.dev/play/p/ZnGIHf8_o4E)]

<h3 id="random"> 18. random 随机数生成器包，可以生成随机[]bytes, int, string。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/random"
```

#### 函数列表:

-   **<big>RandBytes</big>** : 生成随机字节切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandBytes)]
    [[play](https://go.dev/play/p/EkiLESeXf8d)]
-   **<big>RandInt</big>** : 生成随机 int, 范围[min, max)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandInt)]
    [[play](https://go.dev/play/p/pXyyAAI5YxD)]
-   **<big>RandString</big>** : 生成给定长度的随机字符串，只包含字母(a-zA-Z)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandString)]
    [[play](https://go.dev/play/p/W2xvRUXA7Mi)]
-   **<big>RandUpper</big>** : 生成给定长度的随机大写字母字符串(A-Z)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandUpper)]
    [[play](https://go.dev/play/p/29QfOh0DVuh)]
-   **<big>RandLower</big>** : 生成给定长度的随机小写字母字符串(a-z)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandLower)]
    [[play](https://go.dev/play/p/XJtZ471cmtI)]
-   **<big>RandNumeral</big>** : 生成给定长度的随机数字字符串(0-9)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandNumeral)]
    [[play](https://go.dev/play/p/g4JWVpHsJcf)]
-   **<big>RandNumeralOrLetter</big>** : 生成给定长度的随机字符串（数字+字母)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandNumeralOrLetter)]
    [[play](https://go.dev/play/p/19CEQvpx2jD)]
-   **<big>UUIdV4</big>** : 生成 UUID v4 字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#UUIdV4)]
    [[play](https://go.dev/play/p/_Z9SFmr28ft)]
-   **<big>RandUniqueIntSlice</big>** : 生成一个不重复的随机 int 切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandUniqueIntSlice)]
    [[play](https://go.dev/play/p/uBkRSOz73Ec)]
-   **<big>RandSymbolChar</big>** : 生成给定长度的随机符号字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandSymbolChar)]
    [[play](https://go.dev/play/p/Im6ZJxAykOm)]
-   **<big>RandFloat</big>** : 生成随机 float64 数字，可以指定范围和精度。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandFloat)]
    [[play](https://go.dev/play/p/zbD_tuobJtr)]
-   **<big>RandFloats</big>** : 生成随机 float64 数字切片，可以指定长度，范围和精度.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandFloats)]
    [[play](https://go.dev/play/p/uBkRSOz73Ec)]
-   **<big>RandStringSlice</big>** : 生成随机字符串 slice。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandStringSlice)]
    [[play](https://go.dev/play/p/2_-PiDv3tGn)]
-   **<big>RandBool</big>** : 生成随机 bool 值(true or false)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandBool)]
    [[play](https://go.dev/play/p/to6BLc26wBv)]
-   **<big>RandBoolSlice</big>** : 生成特定长度的随机 bool slice。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandBoolSlice)]
    [[play](https://go.dev/play/p/o-VSjPjnILI)]
-   **<big>RandIntSlice</big>** : 生成一个特定长度的随机 int 切片，数值范围[min, max)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandIntSlice)]
    [[play](https://go.dev/play/p/GATTQ5xTEG8)]
-   **<big>RandFromGivenSlice</big>** : 从给定切片中随机生成元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandFromGivenSlice)]
    [[play](https://go.dev/play/p/UrkWueF6yYo)]
-   **<big>RandSliceFromGivenSlice</big>** : 从给定切片中生成长度为 num 的随机切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandSliceFromGivenSlice)]
    [[play](https://go.dev/play/p/68UikN9d6VT)]
-   **<big>RandNumberOfLength</big>** : 生成指定长度的随机数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/random.md#RandNumberOfLength)]
    [[play](https://go.dev/play/p/oyZbuV7bu7b)]

<h3 id="retry"> 19. retry 重试执行函数直到函数运行成功或被 context cancel。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### 函数列表:

-   **<big>Context</big>** : 设置重试 context 参数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#Context)]
    [[play](https://go.dev/play/p/xnAOOXv9GkS)]
-   **<big>Retry</big>** : 重试执行函数 retryFunc，直到函数运行成功，或被 context 取消。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#Retry)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryFunc</big>** : 重试执行的函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#RetryFunc)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryTimes</big>** : 设置重试次数，默认 5。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#RetryTimes)]
    [[play](https://go.dev/play/p/ssfVeU2SwLO)]
-   **<big>BackoffStrategy</big>** : 定义计算退避间隔的方法的接口。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#BackoffStrategy)]
-   **<big>RetryWithCustomBackoff</big>** : 设置自定义退避策略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#RetryWithCustomBackoff)]
    [[play](https://go.dev/play/p/jIm_o2vb5Y4)]
-   **<big>RetryWithLinearBackoff</big>** : 设置线性策略退避。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#RetryWithLinearBackoff)]
    [[play](https://go.dev/play/p/PDet2ZQZwcB)]
-   **<big>RetryWithExponentialWithJitterBackoff</big>** : 设置指数策略退避。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/retry.md#RetryWithExponentialWithJitterBackoff)]
    [[play](https://go.dev/play/p/xp1avQmn16X)]

<h3 id="slice"> 20. slice 包含操作切片的方法集合。&nbsp; &nbsp; &nbsp; &nbsp; <a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### 函数列表:

-   **<big>AppendIfAbsent</big>** : 当前切片中不包含值时，将该值追加到切片中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#AppendIfAbsent)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>Contain</big>** : 判断 slice 是否包含 value。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Contain)]
    [[play](https://go.dev/play/p/_454yEHcNjf)]
-   **<big>ContainBy</big>** : 根据 predicate 函数判断切片是否包含某个值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ContainBy)]
    [[play](https://go.dev/play/p/49tkHfX4GNc)]
-   **<big>ContainSubSlice</big>** : 判断 slice 是否包含 subslice。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ContainSubSlice)]
    [[play](https://go.dev/play/p/bcuQ3UT6Sev)]
-   **<big>ContainAny</big>** : 判断 slice 是否包含 targets 切片中的任意一个元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ContainAny)]
    [[play](https://go.dev/play/p/4xoxhc9XSSw)]
-   **<big>Chunk</big>** : 按照 size 参数均分 slice。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Chunk)]
    [[play](https://go.dev/play/p/b4Pou5j2L_C)]
-   **<big>Compact</big>** : 去除 slice 中的假值（false values are false, nil, 0, ""）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Compact)]
    [[play](https://go.dev/play/p/pO5AnxEr3TK)]
-   **<big>Concat</big>** : 合并多个 slices 到一个 slice 中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Concat)]
    [[play](https://go.dev/play/p/gPt-q7zr5mk)]
-   **<big>Count</big>** : 返回切片中指定元素的个数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Count)]
    [[play](https://go.dev/play/p/Mj4oiEnQvRJ)]
-   **<big>CountBy</big>** : 遍历切片，对每个元素执行函数 predicate. 返回符合函数返回值为 true 的元素的个数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#CountBy)]
    [[play](https://go.dev/play/p/tHOccTMDZCC)]
-   **<big>Difference</big>** : 创建一个切片，其元素不包含在另一个给定切片中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Difference)]
    [[play](https://go.dev/play/p/VXvadzLzhDa)]
-   **<big>DifferenceBy</big>** : 将两个 slice 中的每个元素调用 iteratee 函数，并比较它们的返回值，如果不相等返回在 slice 中对应的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#DifferenceBy)]
    [[play](https://go.dev/play/p/DiivgwM5OnC)]
-   **<big>DifferenceWith</big>** : 接受比较器函数，该比较器被调用以将切片的元素与值进行比较。 结果值的顺序和引用由第一个切片确定。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#DifferenceWith)]
    [[play](https://go.dev/play/p/v2U2deugKuV)]
-   **<big>DeleteAt</big>** : 删除切片中指定索引到的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#DeleteAt)]
    [[play](https://go.dev/play/p/800B1dPBYyd)]
-   **<big>DeleteRange</big>** : 删除切片中指定开始索引到结束索引的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#DeleteRange)]
    [[play](https://go.dev/play/p/945HwiNrnle)]
-   **<big>Drop</big>** : 从切片头部删除 n 个元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Drop)]
    [[play](https://go.dev/play/p/jnPO2yQsT8H)]
-   **<big>DropRight</big>** : 从切片尾部删除 n 个元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#DropRight)]
    [[play](https://go.dev/play/p/8bcXvywZezG)]
-   **<big>DropWhile</big>** : 从切片的头部删除 n 个元素，这个 n 个元素满足 predicate 函数返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#DropWhile)]
    [[play](https://go.dev/play/p/4rt252UV_qs)]
-   **<big>DropRightWhile</big>** : 从切片的尾部删除 n 个元素，这个 n 个元素满足 predicate 函数返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#DropRightWhile)]
    [[play](https://go.dev/play/p/6wyK3zMY56e)]
-   **<big>Equal</big>** : 检查两个切片是否相等，相等条件：切片长度相同，元素顺序和值都相同。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Equal)]
    [[play](https://go.dev/play/p/WcRQJ37ifPa)]
-   **<big>EqualWith</big>** : 检查两个切片是否相等，相等条件：对两个切片的元素调用比较函数 comparator，返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#EqualWith)]
    [[play](https://go.dev/play/p/b9iygtgsHI1)]
-   **<big>EqualUnordered</big>** : 检查两个切片是否相等，元素数量相同，值相等，不考虑元素顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#EqualUnordered)]
    [[play](https://go.dev/play/p/n8fSc2w8ZgX)]
-   **<big>Every</big>** : 如果切片中的所有值都通过谓词函数，则返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Every)]
    [[play](https://go.dev/play/p/R8U6Sl-j8cD)]
-   **<big>Filter</big>** : 返回切片中通过 predicate 函数真值测试的所有元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Filter)]
    [[play](https://go.dev/play/p/SdPna-7qK4T)]
-   **<big>FilterMap</big>** : 返回一个将 filter 和 map 操作应用于给定切片的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#FilterMap)]
    [[play](https://go.dev/play/p/J94SZ_9MiIe)]
-   **<big>Find<sup>deprecated</sup></big>** : 遍历切片的元素，返回第一个通过 predicate 函数真值测试的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Find)]
    [[play](https://go.dev/play/p/CBKeBoHVLgq)]
-   **<big>FindBy</big>** : 遍历切片的元素，返回第一个通过 predicate 函数真值测试的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#FindBy)]
    [[play](https://go.dev/play/p/n1lysBYl-GB)]
-   **<big>FindLast<sup>deprecated</sup></big>** : 从头到尾遍历 slice 的元素，返回最后一个通过 predicate 函数真值测试的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#FindLast)]
    [[play](https://go.dev/play/p/FFDPV_j7URd)]
-   **<big>FindLastBy</big>** : 从遍历 slice 的元素，返回最后一个通过 predicate 函数真值测试的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#FindLastBy)]
    [[play](https://go.dev/play/p/8iqomzyCl_s)]
-   **<big>Flatten</big>** : 将多维切片展平一层。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Flatten)]
    [[play](https://go.dev/play/p/hYa3cBEevtm)]
-   **<big>FlattenDeep</big>** : 将多维切片递归展平到一层。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#FlattenDeep)]
    [[play](https://go.dev/play/p/yjYNHPyCFaF)]
-   **<big>FlatMap</big>** : 将切片转换为其它类型切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#FlatMap)]
    [[play](https://go.dev/play/p/_QARWlWs1N_F)]
-   **<big>ForEach</big>** : 遍历切片的元素并为每个元素调用 iteratee 函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ForEach)]
    [[play](https://go.dev/play/p/DrPaa4YsHRF)]
-   **<big>ForEachConcurrent</big>** : 对 slice 并发执行 foreach 操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ForEachConcurrent)]
    [[play](https://go.dev/play/p/kT4XW7DKVoV)]
-   **<big>ForEachWithBreak</big>** : 遍历切片的元素并为每个元素调用 iteratee 函数，当 iteratee 函数返回 false 时，终止遍历。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ForEachWithBreak)]
    [[play](https://go.dev/play/p/qScs39f3D9W)]
-   **<big>GroupBy</big>** : 迭代切片的元素，每个元素将按条件分组，返回两个切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#GroupBy)]
    [[play](https://go.dev/play/p/QVkPxzPR0iA)]
-   **<big>GroupWith</big>** : 创建一个 map，key 是 iteratee 遍历 slice 中的每个元素返回的结果。值是切片元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#GroupWith)]
    [[play](https://go.dev/play/p/ApCvMNTLO8a)]
-   **<big>IntSlice<sup>deprecated</sup></big>** : 将接口切片转换为 int 切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#IntSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>InterfaceSlice<sup>deprecated</sup></big>** : 将值转换为 interface 切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#InterfaceSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>Intersection</big>** : 返回多个切片的交集。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Intersection)]
    [[play](https://go.dev/play/p/anJXfB5wq_t)]
-   **<big>InsertAt</big>** : 将元素插入到索引处的切片中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#InsertAt)]
    [[play](https://go.dev/play/p/hMLNxPEGJVE)]
-   **<big>IndexOf</big>** : 返回在切片中找到值的第一个匹配项的索引，如果找不到值，则返回-1。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#IndexOf)]
    [[play](https://go.dev/play/p/MRN1f0FpABb)]
-   **<big>LastIndexOf</big>** : 返回在切片中找到最后一个值的索引，如果找不到该值，则返回-1。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#LastIndexOf)]
    [[play](https://go.dev/play/p/DokM4cf1IKH)]
-   **<big>Map</big>** : 对 slice 中的每个元素执行 map 函数以创建一个新切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Map)]
    [[play](https://go.dev/play/p/biaTefqPquw)]
-   **<big>MapConcurrent</big>** : 对 slice 并发执行 map 操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#MapConcurrent)]
    [[play](https://go.dev/play/p/H1ehfPkPen0)]
-   **<big>Merge</big>** : 合并多个切片（不会消除重复元素)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Merge)]
    [[play](https://go.dev/play/p/lbjFp784r9N)]
-   **<big>Reverse</big>** : 反转切片中的元素顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Reverse)]
    [[play](https://go.dev/play/p/8uI8f1lwNrQ)]
-   **<big>ReverseCopy</big>** : 反转切片中的元素顺序, 不改变原 slice。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#ReverseCopy)]
    [[play](https://go.dev/play/p/c9arEaP7Cg-)]
-   **<big>Reduce<sup>deprecated</sup></big>** : 将切片中的元素依次运行 iteratee 函数，返回运行结果。(废弃：建议使用 ReduceBy)
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Reduce)]
    [[play](https://go.dev/play/p/_RfXJJWIsIm)]
-   **<big>ReduceBy</big>** : 对切片元素执行 reduce 操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ReduceBy)]
    [[play](https://go.dev/play/p/YKDpLi7gtee)]
-   **<big>ReduceRight</big>** : 类似 ReduceBy 操作，迭代切片元素顺序从右至左。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ReduceRight)]
    [[play](https://go.dev/play/p/qT9dZC03A1K)]
-   **<big>ReduceConcurrent</big>** : 对切片元素执行并发 reduce 操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ReduceConcurrent)]
    [[play](https://go.dev/play/p/Tjwe6OtaG07)]
-   **<big>Replace</big>** : 返回切片的副本，其中前 n 个不重叠的 old 替换为 new。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Replace)]
    [[play](https://go.dev/play/p/P5mZp7IhOFo)]
-   **<big>ReplaceAll</big>** : 返回切片的副本，将其中 old 全部替换为 new。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ReplaceAll)]
    [[play](https://go.dev/play/p/CzqXMsuYUrx)]
-   **<big>Repeat</big>** : 创建一个切片，包含 n 个传入的 item。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Repeat)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>Shuffle</big>** : 随机打乱切片中的元素顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Shuffle)]
    [[play](https://go.dev/play/p/YHvhnWGU3Ge)]
-   **<big>ShuffleCopy</big>** : 随机打乱切片中的元素顺序, 不改变原切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ShuffleCopy)]
    [[play](https://go.dev/play/p/vqDa-Gs1vT0)]
-   **<big>IsAscending</big>** : 检查切片元素是否按升序排列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#IsAscending)]
    [[play](https://go.dev/play/p/9CtsFjet4SH)]
-   **<big>IsDescending</big>** : 检查切片元素是否按降序排列。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#IsDescending)]
    [[play](https://go.dev/play/p/U_LljFXma14)]
-   **<big>IsSorted</big>** : 检查切片元素是否是有序的（升序或降序）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#IsSorted)]
    [[play](https://go.dev/play/p/nCE8wPLwSA-)]
-   **<big>IsSortedByKey</big>** : 通过 iteratee 函数，检查切片元素是否是有序的。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#IsSortedByKey)]
    [[play](https://go.dev/play/p/tUoGB7DOHI4)]
-   **<big>Sort</big>** : 对任何有序类型（数字或字符串）的切片进行排序，使用快速排序算法。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Sort)]
    [[play](https://go.dev/play/p/V9AVjzf_4Fk)]
-   **<big>SortBy</big>** : 按照 less 函数确定的升序规则对切片进行排序。排序不保证稳定性。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#SortBy)]
    [[play](https://go.dev/play/p/DAhLQSZEumm)]
-   **<big>SortByField<sup>deprecated</sup></big>** : 按字段对结构切片进行排序。slice 元素应为 struct，字段类型应为 int、uint、string 或 bool。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#SortByField)]
    [[play](https://go.dev/play/p/fU1prOBP9p1)]
-   **<big>Some</big>** : 如果列表中的任何值通过谓词函数，则返回 true。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Some)]
    [[play](https://go.dev/play/p/4pO9Xf9NDGS)]
-   **<big>StringSlice<sup>deprecated</sup></big>** : 将接口切片转换为字符串切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#StringSlice)]
    [[play](https://go.dev/play/p/W0TZDWCPFcI)]
-   **<big>SymmetricDifference</big>** : 返回一个切片，其中的元素存在于参数切片中，但不同时存储在于参数切片中（交集取反）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#h42nJX5xMln)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>ToSlice</big>** : 将可变参数转为切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ToSlice)]
    [[play](https://go.dev/play/p/YzbzVq5kscN)]
-   **<big>ToSlicePointer</big>** : 将可变参数转为指针切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ToSlicePointer)]
    [[play](https://go.dev/play/p/gx4tr6_VXSF)]
-   **<big>Unique</big>** : 删除切片中的重复元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Unique)]
    [[play](https://go.dev/play/p/AXw0R3ZTE6a)]
-   **<big>UniqueBy</big>** : 根据迭代函数返回的值，从输入切片中移除重复元素。此函数保持元素的顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#UniqueBy)]
    [[play](https://go.dev/play/p/GY7JE4yikrl)]
-   **<big>UniqueByComparator</big>** : 使用提供的比较器函数从输入切片中移除重复元素。此函数保持元素的顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#UniqueByComparator)]
    [[play](https://go.dev/play/p/rwSacr-ZHsR)]
-   **<big>UniqueByField</big>** : 根据 struct 字段对 struct 切片去重复。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#UniqueByField)]
    [[play](https://go.dev/play/p/6cifcZSPIGu)]
-   **<big>UniqueByConcurrent</big>** : 并发的从输入切片中移除重复元素，结果保持元素的顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#UniqueByConcurrent)]
    [[play](https://go.dev/play/p/wXZ7LcYRMGL)]
-   **<big>Union</big>** : 合并多个切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Union)]
    [[play](https://go.dev/play/p/hfXV1iRIZOf)]
-   **<big>UnionBy</big>** : 对切片的每个元素调用函数后，合并多个切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#UnionBy)]
    [[play](https://go.dev/play/p/HGKHfxKQsFi)]
-   **<big>UpdateAt</big>** : 更新索引处的切片元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#UpdateAt)]
    [[play](https://go.dev/play/p/f3mh2KloWVm)]
-   **<big>Without</big>** : 创建一个不包括所有给定值的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Without)]
    [[play](https://go.dev/play/p/bwhEXEypThg)]
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#KeyBy)]
    [[play](https://go.dev/play/p/uXod2LWD1Kg)]
-   **<big>Join</big>** : 用指定的分隔符链接切片元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Join)]
    [[play](https://go.dev/play/p/huKzqwNDD7V)]
-   **<big>Partition</big>** : 根据给定的 predicate 判断函数分组切片元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Partition)]
    [[play](https://go.dev/play/p/lkQ3Ri2NQhV)]
-   **<big>Random</big>** : 随机返回切片中元素以及下标, 当切片长度为 0 时返回下标-1。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Random)]
    [[play](https://go.dev/play/p/UzpGQptWppw)]
-   **<big>SetToDefaultIf</big>** : 根据给定给定的 predicate 判定函数来修改切片中的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#SetToDefaultIf)]
    [[play](https://go.dev/play/p/9AXGlPRC0-A)]
-   **<big>Break</big>** : 根据判断函数将切片分成两部分。它开始附加到与函数匹配的第一个元素之后的第二个切片。第一个匹配之后的所有元素都包含在第二个切片中，无论它们是否与函数匹配。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Break)]
-   **<big>RightPadding</big>** : 在切片的右部添加元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#RightPadding)]
    [[play](https://go.dev/play/p/0_2rlLEMBXL)]
-   **<big>LeftPadding</big>** : 在切片的左部添加元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#LeftPadding)]
    [[play](https://go.dev/play/p/jlQVoelLl2k)]
-   **<big>Frequency</big>** : 计算切片中每个元素出现的频率。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#Frequency)]
    [[play](https://go.dev/play/p/CW3UVNdUZOq)]
-   **<big>JoinFunc</big>** : 将切片元素用给定的分隔符连接成一个单一的字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#JoinFunc)]
    [[play](https://go.dev/play/p/55ib3SB5fM2)]
-   **<big>ConcatBy</big>** : 将切片中的元素连接成一个值，使用指定的分隔符和连接器函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/slice.md#ConcatBy)]
    [[play](https://go.dev/play/p/6QcUpcY4UMW)]

<h3 id="stream"> 21. stream 流，该包仅验证简单的 stream 实现，功能有限。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/stream"
```

#### 函数列表:

-   **<big>Of</big>** : 创建元素为指定值的 stream。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Of)]
    [[play](https://go.dev/play/p/jI6_iZZuVFE)]
-   **<big>FromSlice</big>** : 从切片创建 stream。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#FromSlice)]
    [[play](https://go.dev/play/p/wywTO0XZtI4)]
-   **<big>FromChannel</big>** : 从通道创建 stream。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#FromChannel)]
    [[play](https://go.dev/play/p/9TZYugGMhXZ)]
-   **<big>FromRange</big>** : 指定一个数字范围创建 stream, 范围两端点值都包括在内。. [start, end]
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#FromRange)]
    [[play](https://go.dev/play/p/9Ex1-zcg-B-)]
-   **<big>Generate</big>** : 创建一个 stream，其中每个元素都由提供的生成器函数生成。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Generate)]
    [[play](https://go.dev/play/p/rkOWL1yA3j9)]
-   **<big>Concat</big>** : 创建一个延迟连接 stream，其元素是第一个 stream 的所有元素，后跟第二个 stream 的全部元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Concat)]
    [[play](https://go.dev/play/p/HM4OlYk_OUC)]
-   **<big>Distinct</big>** : 创建并返回一个 stream，用于删除重复的项。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Distinct)]
    [[play](https://go.dev/play/p/eGkOSrm64cB)]
-   **<big>Filter</big>** : 返回一个通过判定函数的 stream。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Filter)]
    [[play](https://go.dev/play/p/MFlSANo-buc)]
-   **<big>FilterConcurrent</big>** : 对 slice 并发执行 filter 操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#FilterConcurrent)]
    [[play](https://go.dev/play/p/t_pkwerIRVx)]
-   **<big>Map</big>** : 返回一个 stream，该 stream 由将给定函数应用于源 stream 元素的元素组成。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Map)]
    [[play](https://go.dev/play/p/OtNQUImdYko)]
-   **<big>Peek</big>** : 返回一个由源 stream 的元素组成的 stream，并在从生成的 stream 中消耗元素时对每个元素执行所提供的操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Peek)]
    [[play](https://go.dev/play/p/u1VNzHs6cb2)]
-   **<big>Skip</big>** : 在丢弃 stream 的前 n 个元素后，返回由源 stream 的其余元素组成的 stream。如果此 stream 包含的元素少于 n 个，则将返回一个空 stream。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Skip)]
    [[play](https://go.dev/play/p/fNdHbqjahum)]
-   **<big>Limit</big>** : 返回由源 stream 的元素组成的 stream，该 stream 被截断为长度不超过 maxSize。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Limit)]
    [[play](https://go.dev/play/p/qsO4aniDcGf)]
-   **<big>Reverse</big>** : 返回元素与源 stream 的顺序相反的 stream。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Reverse)]
    [[play](https://go.dev/play/p/A8_zkJnLHm4)]
-   **<big>Range</big>** : 返回一个 stream，该 stream 的元素在从源 stream 的开始（包含）到结束（排除）的范围内。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Range)]
    [[play](https://go.dev/play/p/indZY5V2f4j)]
-   **<big>Sorted</big>** : 返回一个 stream，该 stream 由源 stream 的元素组成，并根据提供的 less 函数进行排序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Sorted)]
    [[play](https://go.dev/play/p/XXtng5uonFj)]
-   **<big>ForEach</big>** : 对 stream 的每个元素执行一个操作。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#ForEach)]
    [[play](https://go.dev/play/p/Dsm0fPqcidk)]
-   **<big>Reduce</big>** : 使用关联累加函数对 stream 的元素执行 reduce 操作，并 reduce 操作结果（如果有）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Reduce)]
    [[play](https://go.dev/play/p/6uzZjq_DJLU)]
-   **<big>FindFirst</big>** : 返回此 stream 的第一个元素，如果 stream 为空，则返回零值和 false。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#FindFirst)]
    [[play](https://go.dev/play/p/9xEf0-6C1e3)]
-   **<big>FindLast</big>** : 返回此 stream 的最后一个元素，如果 stream 为空，则返回零值和 false。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#FindLast)]
    [[play](https://go.dev/play/p/WZD2rDAW-2h)]
-   **<big>Max</big>** : 根据提供的 less 函数返回 stream 的最大元素。less 函数: a > b
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Max)]
    [[play](https://go.dev/play/p/fm-1KOPtGzn)]
-   **<big>Min</big>** : 根据提供的 less 函数返回 stream 的最小元素。less 函数: a < b
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Min)]
    [[play](https://go.dev/play/p/vZfIDgGNRe_0)]
-   **<big>AllMatch</big>** : 判断 stream 的所有元素是否全部匹配指定判定函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#AllMatch)]
    [[play](https://go.dev/play/p/V5TBpVRs-Cx)]
-   **<big>AnyMatch</big>** : 判断 stream 是否包含匹配指定判定函数的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#AnyMatch)]
    [[play](https://go.dev/play/p/PTCnWn4OxSn)]
-   **<big>NoneMatch</big>** : 判断 stream 的元素是否全部不匹配指定的判定函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#NoneMatch)]
    [[play](https://go.dev/play/p/iWS64pL1oo3)]
-   **<big>Count</big>** : 返回 stream 中元素的数量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#Count)]
    [[play](https://go.dev/play/p/r3koY6y_Xo-)]
-   **<big>ToSlice</big>** : 返回 stream 中的元素切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#ToSlice)]
    [[play](https://go.dev/play/p/jI6_iZZuVFE)]
-   **<big>IndexOf</big>** : 返回在 stream 中找到值的第一个匹配项的索引，如果找不到值，则返回-1。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#IndexOf)]
    [[play](https://go.dev/play/p/tBV5Nc-XDX2)]
-   **<big>LastIndexOf</big>** : 返回在 stream 中找到值的最后一个匹配项的索引，如果找不到值，则返回-1。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/stream.md#LastIndexOf)]
    [[play](https://go.dev/play/p/CjeoNw2eac_G)]

<h3 id="structs"> 22. structs 提供操作 struct, tag, field 的相关函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/structs"
```

#### 函数列表:

-   **<big>New</big>** : `Struct`结构体的构造函数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#New)]
    [[play](https://go.dev/play/p/O29l8kk-Z17)]
-   **<big>ToMap</big>** : 将一个合法的 struct 对象转换为 map[string]any。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#ToMap)]
    [[play](https://go.dev/play/p/qQbLySBgerZ)]
-   **<big>Fields</big>** : 获取一个 struct 对象的属性列表。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#Fields)]
    [[play](https://go.dev/play/p/w3Kk_CyVY7D)]
-   **<big>Field</big>** : 根据属性名获取一个 struct 对象的属性。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#Fields)]
    [[play](https://go.dev/play/p/KocZMSYarza)]
-   **<big>IsStruct</big>** : 判断是否为一个合法的 struct 对象。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#IsStruct)]
    [[play](https://go.dev/play/p/bU2FSdkbK1C)]
-   **<big>Tag</big>** : 获取`Field`的`Tag`，默认的 tag key 是 json。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#Tag)]
    [[play](https://go.dev/play/p/DVrx5HvvUJr)]
-   **<big>Name</big>** : 获取属性名。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#Name)]
    [[play](https://go.dev/play/p/zfIGlqsatee)]
-   **<big>Value</big>** : 获取`Field`属性的值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#Value)]
    [[play](https://go.dev/play/p/qufYEU2o4Oi)]
-   **<big>Kind</big>** : 获取属性 Kind。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#Kind)]
    [[play](https://go.dev/play/p/wg4NlcUNG5o)]
-   **<big>IsEmbedded</big>** : 判断属性是否为嵌入。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#IsEmbedded)]
    [[play](https://go.dev/play/p/wV2PrbYm3Ec)]
-   **<big>IsExported</big>** : 判断属性是否导出。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#IsExported)]
    [[play](https://go.dev/play/p/csK4AXYaNbJ)]
-   **<big>IsZero</big>** : 判断属性是否为零值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#IsZero)]
    [[play](https://go.dev/play/p/RzqpGISf87r)]
-   **<big>IsSlice</big>** : 判断属性是否是切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#IsSlice)]
    [[play](https://go.dev/play/p/MKz4CgBIUrU)]
-   **<big>IsTargetType</big>** : 判断属性是否是目标类型。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#IsTargetType)]
    [[play](https://go.dev/play/p/Ig75P-agN39)]
-   **<big>TypeName</big>** : 获取结构体类型名。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#TypeName)]
    [[play](https://go.dev/play/p/SWLWd0XBaBb)]

<h3 id="strutil"> 23. strutil 包含字符串处理的相关函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### 函数列表:

-   **<big>After</big>** : 返回源字符串中指定字符串首次出现时的位置之后的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#After)]
    [[play](https://go.dev/play/p/RbCOQqCDA7m)]
-   **<big>AfterLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之后的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#AfterLast)]
    [[play](https://go.dev/play/p/1TegARrb8Yn)]
-   **<big>Before</big>** : 返回源字符串中指定字符串第一次出现时的位置之前的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Before)]
    [[play](https://go.dev/play/p/JAWTZDS4F5w)]
-   **<big>BeforeLast</big>** : 返回源字符串中指定字符串最后一次出现时的位置之前的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#BeforeLast)]
    [[play](https://go.dev/play/p/pJfXXAoG_Te)]
-   **<big>CamelCase</big>** : 将字符串转换为 CamelCase 驼峰式字符串, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#CamelCase)]
    [[play](https://go.dev/play/p/9eXP3tn2tUy)]
-   **<big>Capitalize</big>** : 将字符串的第一个字符转换为大写。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Capitalize)]
    [[play](https://go.dev/play/p/2OAjgbmAqHZ)]
-   **<big>ContainsAll</big>** : 判断字符串是否包括全部给定的子字符串切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#ContainsAll)]
    [[play](https://go.dev/play/p/KECtK2Os4zq)]
-   **<big>ContainsAny</big>** : 判断字符串是否包括给定的子字符串切片中任意一个子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#ContainsAny)]
    [[play](https://go.dev/play/p/dZGSSMB3LXE)]
-   **<big>IsString</big>** : 判断传入参数的数据类型是否为字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#IsString)]
    [[play](https://go.dev/play/p/IOgq7oF9ERm)]
-   **<big>KebabCase</big>** : 将字符串转换为 kebab-case 形式字符串, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#KebabCase)]
    [[play](https://go.dev/play/p/dcZM9Oahw-Y)]
-   **<big>UpperKebabCase</big>** : 将字符串转换为大写 KEBAB-CASE 形式字符串, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#UpperKebabCase)]
    [[play](https://go.dev/play/p/zDyKNneyQXk)]
-   **<big>LowerFirst</big>** : 将字符串的第一个字符转换为小写形式。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#LowerFirst)]
    [[play](https://go.dev/play/p/CbzAyZmtJwL)]
-   **<big>UpperFirst</big>** : 将字符串的第一个字符转换为大写形式。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#UpperFirst)]
    [[play](https://go.dev/play/p/sBbBxRbs8MM)]
-   **<big>Pad</big>** : 如果字符串长度短于 size，则在左右两侧填充字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Pad)]
    [[play](https://go.dev/play/p/NzImQq-VF8q)]
-   **<big>PadEnd</big>** : 如果字符串短于限制大小，则在右侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#PadEnd)]
    [[play](https://go.dev/play/p/9xP8rN0vz--)]
-   **<big>PadStart</big>** : 如果字符串短于限制大小，则在左侧用给定字符填充字符串。 如果填充字符超出大小，它们将被截断。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#PadStart)]
    [[play](https://go.dev/play/p/xpTfzArDfvT)]
-   **<big>Reverse</big>** : 返回字符顺序与给定字符串相反的字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Reverse)]
    [[play](https://go.dev/play/p/adfwalJiecD)]
-   **<big>SnakeCase</big>** : 将字符串转换为 snake_case 形式, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#SnakeCase)]
    [[play](https://go.dev/play/p/tgzQG11qBuN)]
-   **<big>UpperSnakeCase</big>** : 将字符串转换为大写 SNAKE_CASE 形式, 非字母和数字会被忽略。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#UpperSnakeCase)]
    [[play](https://go.dev/play/p/4COPHpnLx38)]
-   **<big>SplitEx</big>** : 拆分给定的字符串可以控制结果切片是否包含空字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#SplitEx)]
    [[play](https://go.dev/play/p/Us-ySSbWh-3)]
-   **<big>Substring</big>** : 根据指定的位置和长度截取子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Substring)]
    [[play](https://go.dev/play/p/q3sM6ehnPDp)]
-   **<big>Wrap</big>** : 用给定字符包裹传入的字符串
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Wrap)]
    [[play](https://go.dev/play/p/KoZOlZDDt9y)]
-   **<big>Unwrap</big>** : 从另一个字符串中解开一个给定的字符串。 将更改源字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Unwrap)]
    [[play](https://go.dev/play/p/Ec2q4BzCpG-)]
-   **<big>SplitWords</big>** : 将字符串拆分为单词，只支持字母字符单词。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#SplitWords)]
    [[play](https://go.dev/play/p/KLiX4WiysMM)]
-   **<big>WordCount</big>** : 返回有意义单词的数量，只支持字母字符单词。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#WordCount)]
    [[play](https://go.dev/play/p/bj7_odx3vRf)]
-   **<big>RemoveNonPrintable</big>** : 删除字符串中不可打印的字符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#RemoveNonPrintable)]
    [[play](https://go.dev/play/p/og47F5x_jTZ)]
-   **<big>StringToBytes</big>** : 在不分配内存的情况下将字符串转换为字节片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#StringToBytes)]
    [[play](https://go.dev/play/p/7OyFBrf9AxA)]
-   **<big>BytesToString</big>** : 在不分配内存的情况下将字节切片转换为字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#BytesToString)]
    [[play](https://go.dev/play/p/6c68HRvJecH)]
-   **<big>IsBlank</big>** : 检查字符串是否为空格或空。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#IsBlank)]
    [[play](https://go.dev/play/p/6zXRH_c0Qd3)]
-   **<big>IsNotBlank</big>** : 检查字符串是否不为空。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#IsNotBlank)]
    [[play](https://go.dev/play/p/e_oJW0RAquA)]
-   **<big>HasPrefixAny</big>** : 检查字符串是否以指定字符串数组中的任何一个开头。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#HasPrefixAny)]
    [[play](https://go.dev/play/p/8UUTl2C5slo)]
-   **<big>HasSuffixAny</big>** : 检查字符串是否以指定字符串数组中的任何一个结尾。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#HasSuffixAny)]
    [[play](https://go.dev/play/p/sKWpCQdOVkx)]
-   **<big>IndexOffset</big>** : 将字符串偏移 idxFrom 后，返回字符串中第一个 substr 实例的索引。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#IndexOffset)]
    [[play](https://go.dev/play/p/qZo4lV2fomB)]
-   **<big>ReplaceWithMap</big>** : 返回 string 的副本，以无序的方式被 map 替换，区分大小写。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#ReplaceWithMap)]
    [[play](https://go.dev/play/p/h3t7CNj2Vvu)]
-   **<big>Trim</big>** : 从字符串的开头和结尾去除空格（或其他字符）。 可选参数 characterMask 指定额外的剥离字符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Trim)]
    [[play](https://go.dev/play/p/Y0ilP0NRV3j)]
-   **<big>SplitAndTrim</big>** : 将字符串 str 按字符串 delimiter 拆分为一个切片，并对该数组的每个元素调用 Trim。忽略 Trim 后为空的元素。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#SplitAndTrim)]
    [[play](https://go.dev/play/p/ZNL6o4SkYQ7)]
-   **<big>HideString</big>** : 隐藏源字符串中的一些字符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#HideString)]
    [[play](https://go.dev/play/p/pzbaIVCTreZ)]
-   **<big>RemoveWhiteSpace</big>** : 删除字符串中的空格。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#RemoveWhiteSpace)]
    [[play](https://go.dev/play/p/HzLC9vsTwkf)]
-   **<big>SubInBetween</big>** : 获取字符串中指定的起始字符串 start 和终止字符串 end 直接的子字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#SubInBetween)]
    [[play](https://go.dev/play/p/EDbaRvjeNsv)]
-   **<big>HammingDistance</big>** : 计算两个字符串之间的汉明距离。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#HammingDistance)]
    [[play](https://go.dev/play/p/glNdQEA9HUi)]
-   **<big>Concat</big>** : 拼接字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Concat)]
    [[play](https://go.dev/play/p/gD52SZHr4Kp)]
-   **<big>Ellipsis</big>** : 将字符串截断到指定长度，并在末尾添加省略号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Ellipsis)]
    [[play](https://go.dev/play/p/i1vbdQiQVRR)]
-   **<big>Shuffle</big>** : 打乱给定字符串中的字符顺序。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Shuffle)]
    [[play](https://go.dev/play/p/iStFwBwyGY7)]
-   **<big>Rotate</big>** : 按指定的字符数旋转字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#Rotate)]
    [[play](https://go.dev/play/p/Kf03iOeT5bd)]
-   **<big>TemplateReplace</big>** : 将模板字符串中的占位符替换为 map 中的相应值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#TemplateReplace)]
    [[play](https://go.dev/play/p/cXSuFvyZqv9)]
-   **<big>RegexMatchAllGroups</big>** : 使用正则表达式匹配字符串中的所有子组并返回结果。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#RegexMatchAllGroups)]
    [[play](https://go.dev/play/p/JZiu0RXpgN-)]
-   **<big>ExtractContent</big>** : 提取两个标记之间的内容。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#ExtractContent)]
    [[play](https://go.dev/play/p/Ay9UIk7Rum9)]
-   **<big>FindAllOccurrences</big>** : 返回子字符串在字符串中所有出现的位置。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/strutil.md#FindAllOccurrences)]
    [[play](https://go.dev/play/p/uvyA6azGLB1)]

<h3 id="system"> 24. system 包含 os, runtime, shell command 的相关函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/system"
```

#### 函数列表:

-   **<big>IsWindows</big>** : 检查当前操作系统是否是 windows。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#IsWindows)]
    [[play](https://go.dev/play/p/XzJULbzmf9m)]
-   **<big>IsLinux</big>** : 检查当前操作系统是否是 linux。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#IsLinux)]
    [[play](https://go.dev/play/p/zIflQgZNuxD)]
-   **<big>IsMac</big>** : 检查当前操作系统是否是 macos。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#IsMac)]
    [[play](https://go.dev/play/p/Mg4Hjtyq7Zc)]
-   **<big>GetOsEnv</big>** : 根据 key 获取对应的环境变量值
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#GetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>SetOsEnv</big>** : 设置环境变量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#SetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>RemoveOsEnv</big>** : 删除环境变量。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#RemoveOsEnv)]
    [[play](https://go.dev/play/p/fqyq4b3xUFQ)]
-   **<big>CompareOsEnv</big>** : 换取环境变量并与传入值进行比较。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#CompareOsEnv)]
    [[play](https://go.dev/play/p/BciHrKYOHbp)]
-   **<big>ExecCommand</big>** : 执行 shell 命令。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#ExecCommand)]
    [[play](https://go.dev/play/p/n-2fLyZef-4)]
-   **<big>GetOsBits</big>** : 获取当前操作系统位数(32/64)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system#GetOsBits)]
    [[play](https://go.dev/play/p/ml-_XH3gJbW)]
-   **<big>StartProcess</big>** :创建进程。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system.md#StartProcess)]
    [[play](https://go.dev/play/p/5GVol6ryS_X)]
-   **<big>StopProcess</big>** : 停止进程。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system.md#StopProcess)]
    [[play](https://go.dev/play/p/jJZhRYGGcmD)]
-   **<big>KillProcess</big>** : 杀掉进程。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system.md#KillProcess)]
    [[play](https://go.dev/play/p/XKmvV-ExBWa)]
-   **<big>GetProcessInfo</big>** : 根据进程 id 获取进程信息。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/system.md#GetProcessInfo)]
    [[play](https://go.dev/play/p/NQDVywEYYx7)]

<h3 id="tuple"> 25. Tuple 包实现一个元组数据类型。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/tuple"
```

#### 函数列表:

-   **<big>Tuple2</big>** : 2 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple2)]
    [[play](https://go.dev/play/p/3sHVqBQpLYN)]
-   **<big>Tuple2_Unbox</big>** : 返回 2 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple2_Unbox)]
    [[play](https://go.dev/play/p/0fD1qfCVwjm)]
-   **<big>Zip2</big>** : 创建一个 Tuple2 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip2)]
    [[play](https://go.dev/play/p/4ncWJJ77Xio)]
-   **<big>Unzip2</big>** : 根据传入的 Tuple2 切片，创建一组和 Tuple2 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip2)]
    [[play](https://go.dev/play/p/KBecr60feXb)]
-   **<big>Tuple3</big>** : 3 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple3)]
    [[play](https://go.dev/play/p/FtH2sdCLlCf)]
-   **<big>Tuple3_Unbox</big>** : 返回 3 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple3_Unbox)]
    [[play](https://go.dev/play/p/YojLy-id1BS)]
-   **<big>Zip3</big>** : 创建一个 Tuple3 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip3)]
    [[play](https://go.dev/play/p/97NgmsTILfu)]
-   **<big>Unzip3</big>** : 根据传入的 Tuple3 切片，创建一组和 Tuple3 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip3)]
    [[play](https://go.dev/play/p/bba4cpAa7KO)]
-   **<big>Tuple4</big>** : 4 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple4)]
    [[play](https://go.dev/play/p/D2EqDz096tk)]
-   **<big>Tuple4_Unbox</big>** : 返回 4 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple4_Unbox)]
    [[play](https://go.dev/play/p/ACj9YuACGgW)]
-   **<big>Zip4</big>** : 创建一个 Tuple4 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip4)]
    [[play](https://go.dev/play/p/PEmTYVK5hL4)]
-   **<big>Unzip4</big>** : 根据传入的 Tuple4 切片，创建一组和 Tuple4 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip4)]
    [[play](https://go.dev/play/p/rb8z4gyYSRN)]
-   **<big>Tuple5</big>** : 5 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple5)]
    [[play](https://go.dev/play/p/2WndmVxPg-r)]
-   **<big>Tuple5_Unbox</big>** : 返回 5 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple5_Unbox)]
    [[play](https://go.dev/play/p/GyIyZHjCvoS)]
-   **<big>Zip5</big>** : 创建一个 Tuple5 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip5)]
    [[play](https://go.dev/play/p/fCAAJLMfBIP)]
-   **<big>Unzip5</big>** : 根据传入的 Tuple5 切片，创建一组和 Tuple5 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip5)]
    [[play](https://go.dev/play/p/gyl6vKfhqPb)]
-   **<big>Tuple6</big>** : 6 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple6)]
    [[play](https://go.dev/play/p/VjqcCwEJZbs)]
-   **<big>Tuple6_Unbox</big>** : 返回 6 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple6_Unbox)]
    [[play](https://go.dev/play/p/FjIHV7lpxmW)]
-   **<big>Zip6</big>** : 创建一个 Tuple6 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip6)]
    [[play](https://go.dev/play/p/oWPrnUYuFHo)]
-   **<big>Unzip6</big>** : 根据传入的 Tuple6 切片，创建一组和 Tuple6 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip6)]
    [[play](https://go.dev/play/p/l41XFqCyh5E)]
-   **<big>Tuple7</big>** : 7 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple7)]
    [[play](https://go.dev/play/p/dzAgv_Ezub9)]
-   **<big>Tuple7_Unbox</big>** : 返回 7 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple7_Unbox)]
    [[play](https://go.dev/play/p/R9I8qeDk0zs)]
-   **<big>Zip7</big>** : 创建一个 Tuple7 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip7)]
    [[play](https://go.dev/play/p/WUJuo897Egf)]
-   **<big>Unzip7</big>** : 根据传入的 Tuple7 切片，创建一组和 Tuple7 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip7)]
    [[play](https://go.dev/play/p/hws_P1Fr2j3)]
-   **<big>Tuple8</big>** : 8 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple8)]
    [[play](https://go.dev/play/p/YA9S0rz3dRz)]
-   **<big>Tuple8_Unbox</big>** : 返回 8 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple8_Unbox)]
    [[play](https://go.dev/play/p/PRxLBBb4SMl)]
-   **<big>Zip8</big>** : 创建一个 Tuple8 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip8)]
    [[play](https://go.dev/play/p/8V9jWkuJfaQ)]
-   **<big>Unzip8</big>** : 根据传入的 Tuple8 切片，创建一组和 Tuple8 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip8)]
    [[play](https://go.dev/play/p/1SndOwGsZB4)]
-   **<big>Tuple9</big>** : 9 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple9)]
    [[play](https://go.dev/play/p/yS2NGGtZpQr)]
-   **<big>Tuple9_Unbox</big>** : 返回 9 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple9_Unbox)]
    [[play](https://go.dev/play/p/oFJFGTAuOa8)]
-   **<big>Zip9</big>** : 创建一个 Tuple9 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip9)]
    [[play](https://go.dev/play/p/cgsL15QYnfz)]
-   **<big>Unzip9</big>** : 根据传入的 Tuple9 切片，创建一组和 Tuple9 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip9)]
    [[play](https://go.dev/play/p/91-BU_KURSA)]
-   **<big>Tuple10</big>** : 10 元元组
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple10)]
    [[play](https://go.dev/play/p/799qqZg0hUv)]
-   **<big>Tuple10_Unbox</big>** : 返回 10 元元组的字段值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Tuple10_Unbox)]
    [[play](https://go.dev/play/p/qfyx3x_X0Cu)]
-   **<big>Zip10</big>** : 创建一个 Tuple10 元组切片, 其中元组的元素和传入切片元素相对应。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Zip10)]
    [[play](https://go.dev/play/p/YSR-2cXnrY4)]
-   **<big>Unzip10</big>** : 根据传入的 Tuple10 切片，创建一组和 Tuple10 元素相对应的切片。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/tuple.md#Unzip10)]
    [[play](https://go.dev/play/p/-taQB6Wfre_z)]

<h3 id="validator"> 26. validator 验证器包，包含常用字符串格式验证函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/validator"
```

#### 函数列表:

-   **<big>ContainChinese</big>** : 验证字符串是否包含中文字符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#ContainChinese)]
    [[play](https://go.dev/play/p/7DpU0uElYeM)]
-   **<big>ContainLetter</big>** : 验证字符串是否包含至少一个英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#ContainLetter)]
    [[play](https://go.dev/play/p/lqFD04Yyewp)]
-   **<big>ContainLower</big>** : 验证字符串是否包含至少一个英文小写字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#ContainLower)]
    [[play](https://go.dev/play/p/Srqi1ItvnAA)]
-   **<big>ContainUpper</big>** : 验证字符串是否包含至少一个英文大写字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#ContainUpper)]
    [[play](https://go.dev/play/p/CmWeBEk27-z)]
-   **<big>IsAlpha</big>** : 验证字符串是否只包含英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsAlpha)]
    [[play](https://go.dev/play/p/7Q5sGOz2izQ)]
-   **<big>IsAllUpper</big>** : 验证字符串是否全是大写英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsAllUpper)]
    [[play](https://go.dev/play/p/ZHctgeK1n4Z)]
-   **<big>IsAllLower</big>** : 验证字符串是否全是小写英文字母。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsAllLower)]
    [[play](https://go.dev/play/p/GjqCnOfV6cM)]
-   **<big>IsBase64</big>** : 验证字符串是否是 base64 编码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsBase64)]
    [[play](https://go.dev/play/p/sWHEySAt6hl)]
-   **<big>IsChineseMobile</big>** : 验证字符串是否是中国手机号码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsChineseMobile)]
    [[play](https://go.dev/play/p/GPYUlGTOqe3)]
-   **<big>IsChineseIdNum</big>** : 验证字符串是否是中国身份证号码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsChineseIdNum)]
    [[play](https://go.dev/play/p/d8EWhl2UGDF)]
-   **<big>IsChinesePhone</big>** : 验证字符串是否是中国电话座机号码(xxx-xxxxxxxx or xxxx-xxxxxxx.)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsChinesePhone)]
    [[play](https://go.dev/play/p/RUD_-7YZJ3I)]
-   **<big>IsCreditCard</big>** : 验证字符串是否是信用卡号码。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsCreditCard)]
    [[play](https://go.dev/play/p/sNwwL6B0-v4)]
-   **<big>IsDns</big>** : 验证字符串是否是有效 dns。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsDns)]
    [[play](https://go.dev/play/p/jlYApVLLGTZ)]
-   **<big>IsEmail</big>** : 验证字符串是否是有效电子邮件地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsEmail)]
    [[play](https://go.dev/play/p/Os9VaFlT33G)]
-   **<big>IsEmptyString</big>** : 验证字符串是否是空字符串。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsEmptyString)]
    [[play](https://go.dev/play/p/dpzgUjFnBCX)]
-   **<big>IsFloat</big>** : 验证参数是否是浮点数((float32，float34)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsFloat)]
    [[play](https://go.dev/play/p/vsyG-sxr99_Z)]
-   **<big>IsFloatStr</big>** : 验证字符串是否是可以转换为浮点数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsFloatStr)]
    [[play](https://go.dev/play/p/LOYwS_Oyl7U)]
-   **<big>IsNumber</big>** : 验证参数是否是数字(integer，float)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsNumber)]
    [[play](https://go.dev/play/p/mdJHOAvtsvF)]
-   **<big>IsNumberStr</big>** : 验证字符串是否是可以转换为数字。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsNumberStr)]
    [[play](https://go.dev/play/p/LzaKocSV79u)]
-   **<big>IsAlphaNumeric</big>** : 验证字符串是字母或数字。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsAlphaNumeric)]
    [[play](https://go.dev/play/p/RHeESLrLg9c)]
-   **<big>IsJSON</big>** : 验证字符串是否是有效 json。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsJSON)]
    [[play](https://go.dev/play/p/8Kip1Itjiil)]
-   **<big>IsRegexMatch</big>** : 验证字符串是否可以匹配正则表达式。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsRegexMatch)]
    [[play](https://go.dev/play/p/z_XeZo_litG)]
-   **<big>IsInt</big>** : 验证参数是否是整数(int, unit)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsInt)]
    [[play](https://go.dev/play/p/eFoIHbgzl-z)]
-   **<big>IsIntStr</big>** : 验证字符串是否是可以转换为整数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsIntStr)]
    [[play](https://go.dev/play/p/jQRtFv-a0Rk)]
-   **<big>IsIp</big>** : 验证字符串是否是 ip 地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsIp)]
    [[play](https://go.dev/play/p/FgcplDvmxoD)]
-   **<big>IsIpV4</big>** : 验证字符串是否是 ipv4 地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsIpV4)]
    [[play](https://go.dev/play/p/zBGT99EjaIu)]
-   **<big>IsIpV6</big>** : 验证字符串是否是 ipv6 地址。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsIpV6)]
    [[play](https://go.dev/play/p/AHA0r0AzIdC)]
-   **<big>IsIpPort</big>** : 检查字符串是否是 ip:port 格式。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsIpPort)]
    [[play](https://go.dev/play/p/xUmls_b9L29)]
-   **<big>IsStrongPassword</big>** : 验证字符串是否是强密码：（字母+数字+特殊字符)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsStrongPassword)]
    [[play](https://go.dev/play/p/QHdVcSQ3uDg)]
-   **<big>IsUrl</big>** : 验证字符串是否是 url。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsUrl)]
    [[play](https://go.dev/play/p/pbJGa7F98Ka)]
-   **<big>IsWeakPassword</big>** : 验证字符串是否是弱密码（只包含字母+数字）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsWeakPassword)]
    [[play](https://go.dev/play/p/wqakscZH5gH)]
-   **<big>IsZeroValue</big>** : 判断传入的参数值是否为零值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsZeroValue)]
    [[play](https://go.dev/play/p/UMrwaDCi_t4)]
-   **<big>IsGBK</big>** : 检查数据编码是否为 gbk（汉字内部代码扩展规范）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsGBK)]
    [[play](https://go.dev/play/p/E2nt3unlmzP)]
-   **<big>IsASCII</big>** : 验证字符串全部为 ASCII 字符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsASCII)]
    [[play](https://go.dev/play/p/hfQNPLX0jNa)]
-   **<big>IsPrintable</big>** : 检查字符串是否全部为可打印字符。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsPrintable)]
    [[play](https://go.dev/play/p/Pe1FE2gdtTP)]
-   **<big>IsBin</big>** : 检查字符串是否是有效的二进制数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsBin)]
    [[play](https://go.dev/play/p/ogPeg2XJH4P)]
-   **<big>IsHex</big>** : 检查字符串是否是有效的十六进制数。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsHex)]
    [[play](https://go.dev/play/p/M2qpHbEwmm7)]
-   **<big>IsBase64URL</big>** : 检查字符串是否是有效的 base64 url。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsBase64URL)]
    [[play](https://go.dev/play/p/vhl4mr8GZ6S)]
-   **<big>IsJWT</big>** : 检查字符串是否是有效的 JSON Web Token (JWT)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsJWT)]
    [[play](https://go.dev/play/p/R6Op7heJbKI)]
-   **<big>IsVisa</big>** : 检查字符串是否是有效的 visa 卡号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsVisa)]
    [[play](https://go.dev/play/p/SdS2keOyJsl)]
-   **<big>IsMasterCard</big>** : 检查字符串是否是有效的 MasterCard 卡号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsMasterCard)]
    [[play](https://go.dev/play/p/CwWBFRrG27b)]
-   **<big>IsAmericanExpress</big>** : 检查字符串是否是有效的 American Express 卡号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsAmericanExpress)]
    [[play](https://go.dev/play/p/HIDFpcOdpkd)]
-   **<big>IsUnionPay</big>** : 检查字符串是否是有效的美国银联卡号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsUnionPay)]
    [[play](https://go.dev/play/p/CUHPEwEITDf)]
-   **<big>IsChinaUnionPay</big>** : 检查字符串是否是有效的中国银联卡号。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsChinaUnionPay)]
    [[play](https://go.dev/play/p/yafpdxLiymu)]
-   **<big>IsPassport</big>** : 判断护照(正则判断)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsPassport)]
    [[play](https://go.dev/play/p/dvOiV2BW7Aw)]
-   **<big>IsChineseHMPassport</big>** : 判断港澳台通行证(正则判断)。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/validator.md#IsChineseHMPassport)]
    [[play](https://go.dev/play/p/xKG6spQTcY0)]

<h3 id="xerror"> 27. xerror 包实现一些错误处理函数。&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">回到目录</a></h3>

```go
import "github.com/duke-git/lancet/v2/xerror"
```

#### 函数列表:

-   **<big>New</big>** : 创建 XError 对象实例。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#New)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>Wrap</big>** : 根据 error 对象创建 XError 对象实例，可添加 message。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#Wrap)]
    [[play](https://go.dev/play/p/5385qT2dCi4)]
-   **<big>Unwrap</big>** : 从 error 对象中解构出 XError。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#Unwrap)]
    [[play](https://go.dev/play/p/LKMLep723tu)]
-   **<big>XError_Wrap</big>** : 创建新的 XError 对象并将消息和 id 复制到新的对象中。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_Wrap)]
    [[play](https://go.dev/play/p/RpjJ5u5sc97)]
-   **<big>XError_Unwrap</big>** : 解构 XEerror 为 error 对象。适配 github.com/pkg/errors。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_Unwrap)
    [[play](https://go.dev/play/p/VUXJ8BST4c6)]
-   **<big>XError_With</big>** : 添加与 XError 对象的键和值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_With)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_Id</big>** : 设置 XError 对象的 id。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_Id)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Is</big>** : 检查目标 error 是否为 XError，两个错误中的 error.id 是否匹配。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_Is)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Values</big>** : 返回由 With 设置的键和值的映射。将合并所有 XError 键和值。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_Values)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_StackTrace</big>** : 返回与 pkg/error 兼容的堆栈信息。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_StackTrace)]
    [[play](https://go.dev/play/p/6FAvSQpa7pc)]
-   **<big>XError_Info</big>** : 返回可打印的 XError 对象信息。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_Info)]
    [[play](https://go.dev/play/p/1ZX0ME1F-Jb)]
-   **<big>XError_Error</big>** : 实现标准库的 error 接口。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#XError_Error)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>TryUnwrap</big>** : 检查 error, 如果 err 为 nil 则展开，则它返回一个有效值，如果 err 不是 nil 则 Unwrap 使用 err 发生 panic。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#TryUnwrap)]
    [[play](https://go.dev/play/p/acyZVkNZEeW)]
-   **<big>TryCatch</big>** : 简单实现的 java 风格异常处理（try-catch-finally）。
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/xerror.md#TryCatch)]
    [[play](https://go.dev/play/p/D5Mdb0mRj0P)]

## 如何贡献代码

#### [代码贡献指南](./CONTRIBUTION.zh-CN.md)

## 贡献者

感谢所有为 lancet 贡献过代码的人！

<a href="https://github.com/duke-git/lancet/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=duke-git/lancet" />
</a>

## GitHub Stars

[![Star History Chart](https://api.star-history.com/svg?repos=duke-git/lancet&type=Date)](https://star-history.com/#duke-git/lancet&Date)
