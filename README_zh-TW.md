<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-2.4.0-green.svg)](https://github.com/duke-git/lancet/releases)
[![GoDoc](https://godoc.org/github.com/duke-git/lancet/v2?status.svg)](https://pkg.go.dev/github.com/duke-git/lancet/v2)
[![test](https://github.com/duke-git/lancet/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/duke-git/lancet/actions/workflows/codecov.yml)
[![codecov](https://codecov.io/gh/duke-git/lancet/branch/main/graph/badge.svg?token=FC48T1F078)](https://codecov.io/gh/duke-git/lancet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/main/LICENSE)
[![Gurubase](https://img.shields.io/badge/Gurubase-Ask%20Lancet%20Guru-006BFF)](https://gurubase.io/g/lancet)

</div>

<div STYLE="page-break-after: always;"></div>

<p style="font-size: 20px"> 
    Lancet 是一個全面、高效、可重複使用的 Go 工具函式庫，靈感來自 Java 的 Apache Commons 套件與 lodash.js。
</p>

## <a href="https://www.golancet.cn/en/" target="_blank"> 官方網站</a> | [English](./README.md) | [简体中文](./README_zh-CN.md)

## 特色

- 👏 全面、高效、可重複使用。
- 💪 700 多個 Go 工具函式，支援字串、切片、日期時間、網路、加密...
- 💅 只依賴兩類函式庫：Go 標準函式庫與 golang.org/x。
- 🌍 每個匯出的函式都有單元測試。

## 安裝

### 注意：

1. <b>使用 go1.18 以上版本的使用者，建議安裝 lancet v2.x.x。因為 v2.x.x 中所有函式都已使用 go1.18 的泛型重寫。</b>

```go
go get github.com/duke-git/lancet/v2 // will install latest version of v2.x.x
```

2. <b>使用 go1.18 以下版本的使用者，應安裝 v1.x.x。v1.x.x 的最新版本為 v1.4.6。 </b>

```go
go get github.com/duke-git/lancet // below go1.18, install latest version of v1.x.x
```

## 用法

Lancet 將程式碼組織為套件結構，使用時需要匯入對應的套件名稱。例如，若要使用字串相關的函式，請如下匯入 strutil 套件：

```go
import "github.com/duke-git/lancet/v2/strutil"
```

## 範例

這裡以字串函式 Reverse（反轉字串）為例，需要匯入 strutil 套件。

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

## 文件

### <span id="index">索引<span>

- [Algorithm](#user-content-algorithm)
- [Compare](#user-content-compare)
- [Concurrency](#user-content-concurrency)
- [Condition](#user-content-condition)
- [Convertor](#user-content-convertor)
- [Cryptor](#user-content-cryptor)
- [Datetime](#user-content-datetime)
- [Datastructure](#user-content-datastructure)
- [EventBus](#user-content-eventbus)
- [Enum](#user-content-enum)
- [Fileutil](#user-content-fileutil)
- [Formatter](#user-content-formatter)
- [Function](#user-content-function)
- [Maputil](#user-content-maputil)
- [Mathutil](#user-content-mathutil)
- [Netutil](#user-content-netutil)
- [Pointer](#user-content-pointer)
- [Random](#user-content-random)
- [Retry](#user-content-retry)
- [Slice](#user-content-slice)
- [Stream](#user-content-stream)
- [Structs](#user-content-structs)
- [Strutil](#user-content-strutil)
- [System](#user-content-system)
- [Tuple](#user-content-tuple)
- [Validator](#user-content-validator)
- [Xerror](#user-content-xerror)

<h3 id="algorithm">1. Algorithm 套件實作了一些基本演算法，例如排序、搜尋。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/algorithm"
```

#### 函式清單:

- **<big>BubbleSort</big>** : 使用氣泡排序演算法對切片排序，會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#BubbleSort)]
  [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
- **<big>CountSort</big>** : 使用氣泡排序演算法對切片排序，不會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#CountSort)]
  [[play](https://go.dev/play/p/tB-Umgm0DrP)]
- **<big>HeapSort</big>** : 使用堆積排序演算法對切片排序，會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#HeapSort)]
  [[play](https://go.dev/play/p/u6Iwa1VZS_f)]
- **<big>InsertionSort</big>** : 使用插入排序演算法對切片排序，會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#InsertionSort)]
  [[play](https://go.dev/play/p/G5LJiWgJJW6)]
- **<big>MergeSort</big>** : 使用合併排序演算法對切片排序，會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#MergeSort)]
  [[play](https://go.dev/play/p/ydinn9YzUJn)]
- **<big>QuickSort</big>** : 使用快速排序演算法對切片排序，會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#QuickSort)]
  [[play](https://go.dev/play/p/7Y7c1Elk3ax)]
- **<big>SelectionSort</big>** : 使用選擇排序演算法對切片排序，會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#SelectionSort)]
  [[play](https://go.dev/play/p/oXovbkekayS)]
- **<big>ShellSort</big>** : 使用希爾排序演算法對切片排序，會改變原始切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#ShellSort)]
  [[play](https://go.dev/play/p/3ibkszpJEu3)]
- **<big>BinarySearch</big>** : 在已排序的切片中使用二分搜尋（遞迴呼叫自身）回傳目標值的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#BinarySearch)]
  [[play](https://go.dev/play/p/t6MeGiUSN47)]
- **<big>BinaryIterativeSearch</big>** : 在已排序的切片中使用二分搜尋（非遞迴）回傳目標值的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#BinaryIterativeSearch)]
  [[play](https://go.dev/play/p/Anozfr8ZLH3)]
- **<big>LinearSearch</big>** : 依據相等判斷函式回傳目標值在切片中的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#LinearSearch)]
  [[play](https://go.dev/play/p/IsS7rgn5s3x)]
- **<big>LRUCache</big>** : 以 LRU 演算法實作的記憶體快取。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#LRUCache)]
  [[play](https://go.dev/play/p/-EZjgOURufP)]

<h3 id="compare"> 2. Compare 套件提供針對任意型別的輕量比較函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/compare"
```

#### 函式清單:

- **<big>Equal</big>** : 檢查兩個值是否相等（同時比對型別與值）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#Equal)]
  [[play](https://go.dev/play/p/wmVxR-to4lz)]
- **<big>EqualValue</big>** : 檢查兩個值是否相等（僅比對值）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#EqualValue)]
  [[play](https://go.dev/play/p/fxnna_LLD9u)]
- **<big>LessThan</big>** : 檢查值 `left` 是否小於值 `right`。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#LessThan)]
  [[play](https://go.dev/play/p/cYh7FQQj0ne)]
- **<big>GreaterThan</big>** : 檢查值 `left` 是否大於值 `right`。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#GreaterThan)]
  [[play](https://go.dev/play/p/9-NYDFZmIMp)]
- **<big>LessOrEqual</big>** : 檢查值 `left` 是否小於或等於值 `right`。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#LessOrEqual)]
  [[play](https://go.dev/play/p/e4T_scwoQzp)]
- **<big>GreaterOrEqual</big>** : 檢查值 `left` 是否大於或等於值 `right`。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#GreaterOrEqual)]
  [[play](https://go.dev/play/p/vx8mP0U8DFk)]
- **<big>InDelta</big>** : 檢查兩個值在容許誤差 delta 內是否相等。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#InDelta)]

<h3 id="concurrency"> 3. Concurrency 套件包含一些支援並行程式設計的函式，例如 goroutine、channel、非同步。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/concurrency"
```

#### 函式清單:

- **<big>NewChannel</big>** : 建立一個 Channel 指標實例。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#NewChannel)]
  [[play](https://go.dev/play/p/7aB4KyMMp9A)]
- **<big>Bridge</big>** : 將多個通道串接成一個通道。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/Bridge.md#NewChannel)]
  [[play](https://go.dev/play/p/qmWSy1NVF-Y)]
- **<big>FanIn</big>** : 將多個通道合併成一個通道。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#FanIn)]
  [[play](https://go.dev/play/p/2VYFMexEvTm)]
- **<big>Generate</big>** : 建立一個通道，並將數個值放入該通道。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Generate)]
  [[play](https://go.dev/play/p/7aB4KyMMp9A)]
- **<big>Or</big>** : 讀取一個或多個通道並匯入同一個通道，任一來源通道關閉時即會關閉。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Or)]
  [[play](https://go.dev/play/p/Wqz9rwioPww)]
- **<big>OrDone</big>** : 讀取一個通道並匯入另一個通道，直到取消 context 為止才會關閉。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#OrDone)]
  [[play](https://go.dev/play/p/lm_GoS6aDjo)]
- **<big>Repeat</big>** : 建立通道並重複放入值，直到取消 context 為止。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Repeat)]
  [[play](https://go.dev/play/p/k5N_ALVmYjE)]
- **<big>RepeatFn</big>** : 建立通道並重複執行 fn，將結果放入通道，直到 context 關閉為止。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#RepeatFn)]
  [[play](https://go.dev/play/p/4J1zAWttP85)]
- **<big>Take</big>** : 建立一個通道，其值取自另一個通道，並限制取出的數量。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Take)]
  [[play](https://go.dev/play/p/9Utt-1pDr2J)]
- **<big>Tee</big>** : 將一個通道拆分成兩個通道，直到取消 context 為止。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Tee)]
  [[play](https://go.dev/play/p/3TQPKnCirrP)]
- **<big>NewKeyedLocker</big>** : KeyedLocker 是依鍵值加鎖的簡易實作，支援非阻塞式取得鎖。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#NewKeyedLocker)]
  [[play](https://go.dev/play/p/GzeyC33T5rw)]
- **<big>Do</big>** :取得指定鍵值的鎖並執行所提供的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Do)]
  [[play](https://go.dev/play/p/GzeyC33T5rw)]
- **<big>NewRWKeyedLocker</big>** :RRWKeyedLocker 是 KeyedLocker 的讀寫鎖版本。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#NewRWKeyedLocker)]
  [[play](https://go.dev/play/p/ZrCr8sMo77T)]
- **<big>RLock</big>** : 取得指定鍵值的讀鎖並執行所提供的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#RLock)]
  [[play](https://go.dev/play/p/ZrCr8sMo77T)]
- **<big>Lock</big>** : 取得指定鍵值的寫鎖並執行所提供的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Lock)]
  [[play](https://go.dev/play/p/WgAcXbOPKGk)]
- **<big>NewTryKeyedLocker</big>** : TryKeyedLocker 是 KeyedLocker 的非阻塞版本。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#NewTryKeyedLocker)]
  [[play](https://go.dev/play/p/VG9qLvyetE2)]
- **<big>TryLock</big>** : TryLock 嘗試取得指定鍵值的鎖。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#TryLock)]
  [[play](https://go.dev/play/p/VG9qLvyetE2)]
- **<big>Unlock</big>** : Unlock 釋放指定鍵值的鎖。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Unlock)]
  [[play](https://go.dev/play/p/VG9qLvyetE2)]

<h3 id="condition"> 4. Condition 套件包含一些條件判斷函式，例如 And、Or、TernaryOperator... &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/condition"
```

#### 函式清單:

- **<big>Bool</big>** : 回傳任意值的真值判斷結果。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Bool)]
  [[play](https://go.dev/play/p/ETzeDJRSvhm)]
- **<big>And</big>** : 當 a 與 b 皆為真值時回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#And)]
  [[play](https://go.dev/play/p/W1SSUmt6pvr)]
- **<big>Or</big>** : 當 a 與 b 皆不為真值時回傳 false。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Or)]
  [[play](https://go.dev/play/p/UlQTxHaeEkq)]
- **<big>Xor</big>** : 當 a 或 b 其中之一（但非兩者）為真值時回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Xor)]
  [[play](https://go.dev/play/p/gObZrW7ZbG8)]
- **<big>Nor</big>** : 當 a 與 b 皆不為真值時回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Nor)]
  [[play](https://go.dev/play/p/g2j08F_zZky)]
- **<big>Xnor</big>** : 當 a 與 b 皆為真值，或皆不為真值時回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Xnor)]
  [[play](https://go.dev/play/p/OuDB9g51643)]
- **<big>Nand</big>** : 當 a 與 b 皆為真值時回傳 false。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Nand)]
  [[play](https://go.dev/play/p/vSRMLxLIbq8)]
- **<big>TernaryOperator</big>** : 三元運算子。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#TernaryOperator)]
  [[play](https://go.dev/play/p/ElllPZY0guT)]

<h3 id="convertor"> 5. Convertor 套件包含一些資料轉換函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/convertor"
```

#### 函式清單:

- **<big>ColorHexToRGB</big>** : 將十六進位色碼轉換為 RGB 色彩。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ColorHexToRGB)]
  [[play](https://go.dev/play/p/o7_ft-JCJBV)]
- **<big>ColorRGBToHex</big>** : 將 RGB 色彩轉換為十六進位色碼。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ColorRGBToHex)]
  [[play](https://go.dev/play/p/nzKS2Ro87J1)]
- **<big>ToBool</big>** : 將字串轉換為布林值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToBool)]
  [[play](https://go.dev/play/p/ARht2WnGdIN)]
- **<big>ToBytes</big>** : 將值轉換為位元組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToBytes)]
  [[play](https://go.dev/play/p/fAMXYFDvOvr)]
- **<big>ToChar</big>** : 將字串轉換為字元切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToChar)]
  [[play](https://go.dev/play/p/JJ1SvbFkVdM)]
- **<big>ToChannel</big>** : 將一組元素轉換為唯讀通道。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToChannel)]
  [[play](https://go.dev/play/p/hOx_oYZbAnL)]
- **<big>ToFloat</big>** : 將值轉換為 float64，若參數無法轉為浮點數，則回傳 0.0 與錯誤。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToFloat)]
  [[play](https://go.dev/play/p/4YTmPCibqHJ)]
- **<big>ToInt</big>** : 將值轉換為 int64，若輸入不是數值，則回傳 0 與錯誤。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToInt)]
  [[play](https://go.dev/play/p/9_h9vIt-QZ_b)]
- **<big>ToJson</big>** : 將值轉換為 JSON 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToJson)]
  [[play](https://go.dev/play/p/2rLIkMmXWvR)]
- **<big>ToMap</big>** : 依據 iteratee 函式將結構切片轉換為 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToMap)]
  [[play](https://go.dev/play/p/tVFy7E-t24l)]
- **<big>ToPointer</big>** : 回傳傳入值的指標。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToPointer)]
  [[play](https://go.dev/play/p/ASf_etHNlw1)]
- **<big>ToPointers</big>** : 將值的切片轉換為指標的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToPointers)]
  [[play](https://go.dev/play/p/ZUoXd2i5ZkV)]
- **<big>FromPointer</big>** : 回傳指標所指向的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#FromPointer)]
  [[play](https://go.dev/play/p/wAp90V7Zu6g)]
- **<big>FromPointers</big>** : 將指標的切片轉換為值的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#FromPointers)]
  [[play](https://go.dev/play/p/qIPsyYtNy3Q)]
- **<big>ToString</big>** : 將值轉換為字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToString)]
  [[play](https://go.dev/play/p/nF1zOOslpQq)]
- **<big>StructToMap</big>** : 將結構轉換為 map，僅轉換匯出的結構欄位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#StructToMap)]
  [[play](https://go.dev/play/p/KYGYJqNUBOI)]
- **<big>MapToSlice</big>** : 依據 iteratee 函式將 map 轉換為切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#MapToSlice)]
  [[play](https://go.dev/play/p/dmX4Ix5V6Wl)]
- **<big>EncodeByte</big>** : 將資料編碼為位元組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#EncodeByte)]
  [[play](https://go.dev/play/p/DVmM1G5JfuP)]
- **<big>DecodeByte</big>** : 將位元組切片資料解碼為目標物件。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#DecodeByte)]
  [[play](https://go.dev/play/p/zI6xsmuQRbn)]
- **<big>DeepClone</big>** : 建立傳入項目的深層複本，無法複製結構中未匯出的欄位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#DeepClone)]
  [[play](https://go.dev/play/p/j4DP5dquxnk)]
- **<big>CopyProperties</big>** : 將來源結構的每個欄位複製到目的結構。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#CopyProperties)]
  [[play](https://go.dev/play/p/oZujoB5Sgg5)]
- **<big>ToInterface</big>** : 將 reflect value 轉換為其介面型別。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToInterface)]
  [[play](https://go.dev/play/p/syqw0-WG7Xd)]
- **<big>Utf8ToGbk</big>** : 將 UTF-8 編碼的資料轉換為 GBK 編碼的資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#Utf8ToGbk)]
  [[play](https://go.dev/play/p/9FlIaFLArIL)]
- **<big>GbkToUtf8</big>** : 將 GBK 編碼的資料轉換為 UTF-8 編碼的資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#GbkToUtf8)]
  [[play](https://go.dev/play/p/OphmHCN_9u8)]
- **<big>ToStdBase64</big>** : 將值轉換為標準 Base64 編碼的字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToStdBase64)]
  [[play](https://go.dev/play/p/_fLJqJD3NMo)]
- **<big>ToUrlBase64</big>** : 將值轉換為 URL Base64 編碼的字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToUrlBase64)]
  [[play](https://go.dev/play/p/C_d0GlvEeUR)]
- **<big>ToRawStdBase64</big>** : 將值轉換為原始（無填充）標準 Base64 編碼的字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToRawStdBase64)]
  [[play](https://go.dev/play/p/wSAr3sfkDcv)]
- **<big>ToRawUrlBase64</big>** : 將值轉換為原始（無填充）URL Base64 編碼的字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToRawUrlBase64)]
  [[play](https://go.dev/play/p/HwdDPFcza1O)]
- **<big>ToBigInt</big>** : 將任何支援的整數型別（int、int64、uint64 等）轉換為 \*big.Int。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToBigInt)]
  [[play](https://go.dev/play/p/X3itkCxwB_x)]

<h3 id="cryptor"> 6. Cryptor 套件用於資料的加密與解密。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### 函式清單:

- **<big>AesEcbEncrypt</big>** : 使用 AES ECB 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesEcbEncrypt)]
  [[play](https://go.dev/play/p/zI6xsmuQRbn)]
- **<big>AesEcbDecrypt</big>** : 使用 AES ECB 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesEcbDecrypt)]
  [[play](https://go.dev/play/p/zI6xsmuQRbn)]
- **<big>AesCbcEncrypt</big>** : 使用 AES CBC 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCbcEncrypt)]
  [[play](https://go.dev/play/p/IOq_g8_lKZD)]
- **<big>AesCbcDecrypt</big>** : 使用 AES CBC 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCbcDecrypt)]
  [[play](https://go.dev/play/p/IOq_g8_lKZD)]
- **<big>AesCtrCrypt</big>** : 使用 AES CRC 演算法以金鑰加密／解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCtrCrypt)]
  [[play](https://go.dev/play/p/SpaZO0-5Nsp)]
- **<big>AesCfbEncrypt</big>** : 使用 AES CFB 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCfbEncrypt)]
  [[play](https://go.dev/play/p/tfkF10B13kH)]
- **<big>AesCfbDecrypt</big>** : 使用 AES CFB 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCfbDecrypt)]
  [[play](https://go.dev/play/p/tfkF10B13kH)]
- **<big>AesOfbEncrypt</big>** : 使用 AES OFB 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesOfbEncrypt)]
  [[play](https://go.dev/play/p/VtHxtkUj-3F)]
- **<big>AesOfbDecrypt</big>** : 使用 AES OFB 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesOfbDecrypt)]
  [[play](https://go.dev/play/p/VtHxtkUj-3F)]
- **<big>AesGcmEncrypt</big>** : 使用 AES GCM 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesGcmEncrypt)]
  [[play](https://go.dev/play/p/rUt0-DmsPCs)]
- **<big>AesGcmDecrypt</big>** : 使用 AES GCM 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesGcmDecrypt)]
  [[play](https://go.dev/play/p/rUt0-DmsPCs)]
- **<big>Base64StdEncode</big>** : 使用 base64 編碼對字串編碼。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Base64StdEncode)]
  [[play](https://go.dev/play/p/VOaUyQUreoK)]
- **<big>Base64StdDecode</big>** : 使用 base64 編碼對字串解碼。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Base64StdDecode)]
  [[play](https://go.dev/play/p/RWQylnJVgIe)]
- **<big>DesEcbEncrypt</big>** : 使用 DES ECB 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesEcbEncrypt)]
  [[play](https://go.dev/play/p/8qivmPeZy4P)]
- **<big>DesEcbDecrypt</big>** : 使用 DES ECB 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesEcbDecrypt)]
  [[play](https://go.dev/play/p/8qivmPeZy4P)]
- **<big>DesCbcEncrypt</big>** : 使用 DES CBC 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCbcEncrypt)]
  [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
- **<big>DesCbcDecrypt</big>** : 使用 DES CBC 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCbcDecrypt)]
  [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
- **<big>DesCtrCrypt</big>** : 使用 DES CRY 演算法以金鑰加密／解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCtrCrypt)]
  [[play](https://go.dev/play/p/9-T6OjKpcdw)]
- **<big>DesCfbEncrypt</big>** : 使用 DES CFB 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCfbEncrypt)]
  [[play](https://go.dev/play/p/y-eNxcFBlxL)]
- **<big>DesCfbDecrypt</big>** : 使用 DES CFB 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCfbDecrypt)]
  [[play](https://go.dev/play/p/y-eNxcFBlxL)]
- **<big>DesOfbEncrypt</big>** : 使用 DES OFB 演算法以金鑰加密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesOfbEncrypt)]
  [[play](https://go.dev/play/p/74KmNadjN1J)]
- **<big>DesOfbDecrypt</big>** : 使用 DES OFB 演算法以金鑰解密位元組切片資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesOfbDecrypt)]
  [[play](https://go.dev/play/p/74KmNadjN1J)]
- **<big>HmacMd5</big>** : 回傳字串的 md5 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacMd5)]
  [[play](https://go.dev/play/p/uef0q1fz53I)]
- **<big>HmacMd5WithBase64</big>** : 回傳 base64 字串的 md5 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacMd5WithBase64)]
- **<big>HmacSha1</big>** : 使用 sha1 回傳字串的 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha1)]
  [[play](https://go.dev/play/p/1UI4oQ4WXKM)]
- **<big>HmacSha1WithBase64</big>** : 使用 sha1 回傳字串經 base64 編碼的 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha1WithBase64)]
  [[play](https://go.dev/play/p/47JmmGrnF7B)]
- **<big>HmacSha256</big>** : 使用 sha256 回傳字串的 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha256)]
  [[play](https://go.dev/play/p/HhpwXxFhhC0)]
- **<big>HmacSha256WithBase64</big>** : 使用 sha256 回傳字串經 base64 編碼的 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha256WithBase64)]
  [[play](https://go.dev/play/p/EKbkUvPTLwO)]
- **<big>HmacSha512</big>** : 使用 sha512 回傳字串的 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha512)]
  [[play](https://go.dev/play/p/59Od6m4A0Ud)]
- **<big>HmacSha512WithBase64</big>** : 使用 sha512 回傳字串經 base64 編碼的 hmac 雜湊值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha512WithBase64)]
  [[play](https://go.dev/play/p/c6dSe3E2ydU)]
- **<big>Md5Byte</big>** : 回傳位元組切片的 md5 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5Byte)]
  [[play](https://go.dev/play/p/suraalH8lyC)]
- **<big>Md5ByteWithBase64</big>** : 回傳位元組切片經 base64 編碼的 md5 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5ByteWithBase64)]
  [[play](https://go.dev/play/p/Tcb-Z7LN2ax)]
- **<big>Md5String</big>** : 回傳字串的 md5 值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5String)]
  [[play](https://go.dev/play/p/1bLcVetbTOI)]
- **<big>Md5StringWithBase64</big>** : 回傳字串經 base64 編碼的 md5 值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5StringWithBase64)]
  [[play](https://go.dev/play/p/Lx4gH7Vdr5_y)]
- **<big>Md5File</big>** : 回傳檔案的 md5 值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5File)]
- **<big>Sha1</big>** : 回傳 base64 字串的 sha1 值（SHA-1 雜湊演算法）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha1)]
  [[play](https://go.dev/play/p/_m_uoD1deMT)]
- **<big>Sha1WithBase64</big>** : 回傳字串的 sha1 值（SHA-1 雜湊演算法）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha1WithBase64)]
  [[play](https://go.dev/play/p/fSyx-Gl2l2-)]
- **<big>Sha256</big>** : 回傳字串的 sha256 值（SHA-256 雜湊演算法）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha256)]
  [[play](https://go.dev/play/p/tU9tfBMIAr1)]
- **<big>Sha256WithBase64</big>** : 回傳 base64 字串的 sha256 值（SHA256 雜湊演算法）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha256WithBase64)]
  [[play](https://go.dev/play/p/85IXJHIal1k)]
- **<big>Sha512</big>** : 回傳字串的 sha512 值（SHA-512 雜湊演算法）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha512)]
  [[play](https://go.dev/play/p/3WsvLYZxsHa)]
- **<big>Sha512WithBase64</big>** : 回傳 base64 字串的 sha512 值（SHA-512 雜湊演算法）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha512WithBase64)]
  [[play](https://go.dev/play/p/q_fY2rA-k5I)]
- **<big>GenerateRsaKey</big>** : 建立 RSA 私鑰與公鑰的 pem 檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#GenerateRsaKey)]
  [[play](https://go.dev/play/p/zutRHrDqs0X)]
- **<big>RsaEncrypt</big>** : 使用 RSA 演算法加密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaEncrypt)]
  [[play](https://go.dev/play/p/7_zo6mrx-eX)]
- **<big>RsaDecrypt</big>** : 使用 RSA 演算法解密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaDecrypt)]
  [[play](https://go.dev/play/p/7_zo6mrx-eX)]
- **<big>GenerateRsaKeyPair</big>** : 建立 RSA 私鑰與公鑰。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#GenerateRsaKeyPair)]
  [[play](https://go.dev/play/p/sSVmkfENKMz)]
- **<big>RsaEncryptOAEP</big>** : 使用 RSA-OAEP 加密指定的資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaEncryptOAEP)]
  [[play](https://go.dev/play/p/sSVmkfENKMz)]
- **<big>RsaDecryptOAEP</big>** : 使用 RSA-OAEP 解密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaDecryptOAEP)]
  [[play](https://go.dev/play/p/sSVmkfENKMz)]
- **<big>RsaSign</big>** : 使用 RSA 對資料簽章。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaSign)]
  [[play](https://go.dev/play/p/qhsbf8BJ6Mf)]
- **<big>RsaVerifySign</big>** : 使用 RSA 驗證資料的簽章。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaVerifySign)]
  [[play](https://go.dev/play/p/qhsbf8BJ6Mf)]
- **<big>GenerateSm2Key</big>** : 產生 SM2 私鑰與公鑰。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#GenerateSm2Key)]
  [[play](https://go.dev/play/p/bKYMqRLvIx3)]
- **<big>Sm2Encrypt</big>** : 使用 SM2 公鑰加密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sm2Encrypt)]
  [[play](https://go.dev/play/p/bKYMqRLvIx3)]
- **<big>Sm2Decrypt</big>** : 使用 SM2 私鑰解密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sm2Decrypt)]
  [[play](https://go.dev/play/p/bKYMqRLvIx3)]
- **<big>Sm3</big>** : 回傳資料的 SM3 雜湊值（256 位元）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sm3)]
  [[play](https://go.dev/play/p/zDAQpteAiOc)]
- **<big>Sm4EcbEncrypt</big>** : 使用 SM4 ECB 模式加密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sm4EcbEncrypt)]
  [[play](https://go.dev/play/p/l5IQxYuuaED)]
- **<big>Sm4EcbDecrypt</big>** : 使用 SM4 ECB 模式解密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sm4EcbDecrypt)]
  [[play](https://go.dev/play/p/l5IQxYuuaED)]
- **<big>Sm4CbcEncrypt</big>** : 使用 SM4 CBC 模式加密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sm4CbcEncrypt)]
  [[play](https://go.dev/play/p/65Q6iYhLRTa)]
- **<big>Sm4CbcDecrypt</big>** : 使用 SM4 CBC 模式解密資料。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sm4CbcDecrypt)]
  [[play](https://go.dev/play/p/65Q6iYhLRTa)]

<h3 id="datetime"> 7. Datetime 套件支援日期與時間的格式化與比較。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/datetime"
```

#### 函式清單:

- **<big>AddDay</big>** : 對時間增加或減少天數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddDay)]
  [[play](https://go.dev/play/p/dIGbs_uTdFa)]
- **<big>AddHour</big>** : 對時間增加或減少天數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddHour)]
  [[play](https://go.dev/play/p/rcMjd7OCsi5)]
- **<big>AddMinute</big>** : 對時間增加或減少天數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddMinute)]
  [[play](https://go.dev/play/p/nT1heB1KUUK)]
- **<big>AddWeek</big>** : 對時間增加或減少週數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddWeek)]
  [[play](https://go.dev/play/p/M9TqdMiaA2p)]
- **<big>AddMonth</big>** : 對時間增加或減少月數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddMonth)]
  [[play](https://go.dev/play/p/DLoiOnpLvsN)]
- **<big>AddYear</big>** : 對時間增加或減少年數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddYear)]
  [[play](https://go.dev/play/p/MqW2ujnBx10)]
- **<big>AddDaySafe</big>** : 對時間增加或減少天數，並確保回傳的日期不會超出目標年月的有效日期。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddDaySafe)]
  [[play](https://go.dev/play/p/JTohZFpoDJ3)]
- **<big>AddMonthSafe</big>** : 對時間增加或減少月數，並確保回傳的日期不會超出目標年月的有效日期。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddMonthSafe)]
  [[play](https://go.dev/play/p/KLw0lo6mbVW)]
- **<big>AddYearSafe</big>** : 對時間增加或減少年數，並確保回傳的日期不會超出目標年月的有效日期。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddYearSafe)]
  [[play](https://go.dev/play/p/KVGXWZZ54ZH)]
- **<big>BeginOfMinute</big>** : 回傳指定日期該分鐘開始的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfMinute)]
  [[play](https://go.dev/play/p/ieOLVJ9CiFT)]
- **<big>BeginOfHour</big>** : 回傳指定日期該小時開始的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfHour)]
  [[play](https://go.dev/play/p/GhdGFnDWpYs)]
- **<big>BeginOfDay</big>** : 回傳指定日期當天開始的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfDay)]
  [[play](https://go.dev/play/p/94m_UT6cWs9)]
- **<big>BeginOfWeek</big>** : 回傳指定日期當週開始的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfWeek)]
  [[play](https://go.dev/play/p/DCHdcL6gnfV)]
- **<big>BeginOfMonth</big>** : 回傳指定日期當月開始的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfMonth)]
  [[play](https://go.dev/play/p/bWXVFsmmzwL)]
- **<big>BeginOfYear</big>** : 回傳指定日期當年開始的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfYear)]
  [[play](https://go.dev/play/p/i326DSwLnV8)]
- **<big>EndOfMinute</big>** : 回傳指定日期該分鐘結束的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfMinute)]
  [[play](https://go.dev/play/p/yrL5wGzPj4z)]
- **<big>EndOfHour</big>** : 回傳指定日期該小時結束的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfHour)]
  [[play](https://go.dev/play/p/6ce3j_6cVqN)]
- **<big>EndOfDay</big>** : 回傳指定日期當天結束的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfDay)]
  [[play](https://go.dev/play/p/eMBOvmq5Ih1)]
- **<big>EndOfWeek</big>** : 回傳指定日期當週結束的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfWeek)]
  [[play](https://go.dev/play/p/mGSA162YgX9)]
- **<big>EndOfMonth</big>** : 回傳指定日期當月結束的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfMonth)]
  [[play](https://go.dev/play/p/_GWh10B3Nqi)]
- **<big>EndOfYear</big>** : 回傳指定日期當年結束的日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfYear)]
  [[play](https://go.dev/play/p/G01cKlMCvNm)]
- **<big>GetNowDate</big>** : 回傳目前日期，格式為 yyyy-mm-dd。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNowDate)]
  [[play](https://go.dev/play/p/PvfkPpcpBBf)]
- **<big>GetNowTime</big>** : 回傳目前時間，格式為 hh-mm-ss。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNowTime)]
  [[play](https://go.dev/play/p/l7BNxCkTmJS)]
- **<big>GetNowDateTime</big>** : 回傳目前日期時間，格式為 yyyy-mm-dd hh-mm-ss。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNowDateTime)]
  [[play](https://go.dev/play/p/pI4AqngD0al)]
- **<big>GetTodayStartTime</big>** : 回傳今天的起始時間，格式：yyyy-mm-dd 00:00:00。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetTodayStartTime)]
  [[play](https://go.dev/play/p/84siyYF7t99)]
- **<big>GetTodayEndTime</big>** : 回傳今天的結束時間，格式：yyyy-mm-dd 23:59:59。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetTodayEndTime)]
  [[play](https://go.dev/play/p/jjrLnfoqgn3)]
- **<big>GetZeroHourTimestamp</big>** : 回傳零時的時間戳（00:00 的時間戳）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetZeroHourTimestamp)]
  [[play](https://go.dev/play/p/QmL2oIaGE3q)]
- **<big>GetNightTimestamp</big>** : 回傳當天最後一刻的時間戳（23:59 的時間戳）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNightTimestamp)]
  [[play](https://go.dev/play/p/UolysR3MYP1)]
- **<big>FormatTimeToStr</big>** : 將時間轉換為字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#FormatTimeToStr)]
  [[play](https://go.dev/play/p/_Ia7M8H_OvE)]
- **<big>FormatStrToTime</big>** : 將字串轉換為時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#FormatStrToTime)]
  [[play](https://go.dev/play/p/1h9FwdU8ql4)]
- **<big>NewUnix</big>** : 回傳指定時間的 unix 時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewUnix)]
  [[play](https://go.dev/play/p/psoSuh_kLRt)]
- **<big>NewUnixNow</big>** : 回傳目前時間的 unix 時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewUnixNow)]
  [[play](https://go.dev/play/p/U4PPx-9D0oz)]
- **<big>NewFormat</big>** : 回傳指定時間字串的 unix 時間戳，t 應為 "yyyy-mm-dd hh:mm:ss"。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewFormat)]
  [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
- **<big>NewISO8601</big>** : 回傳指定 iso8601 時間字串的 unix 時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewISO8601)]
  [[play](https://go.dev/play/p/mkhOHQkdeA2)]
- **<big>ToUnix</big>** : 回傳 unix 時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToUnix)]
  [[play](https://go.dev/play/p/_LUiwAdocjy)]
- **<big>ToFormat</big>** : 回傳 unix 時間對應的時間字串 'yyyy-mm-dd hh:mm:ss'。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToFormat)]
  [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
- **<big>ToFormatForTpl</big>** : 回傳格式為指定樣板 tpl 的時間字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToFormatForTpl)]
  [[play](https://go.dev/play/p/nyXxXcQJ8L5)]
- **<big>ToIso8601</big>** : 回傳 iso8601 時間字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToIso8601)]
  [[play](https://go.dev/play/p/mkhOHQkdeA2)]
- **<big>IsLeapYear</big>** : 檢查參數 `year` 是否為閏年。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#IsLeapYear)]
  [[play](https://go.dev/play/p/xS1eS2ejGew)]
- **<big>BetweenSeconds</big>** : 回傳兩個時間之間相差的秒數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BetweenSeconds)]
  [[play](https://go.dev/play/p/n3YDRyfyXJu)]
- **<big>DayOfYear</big>** : 回傳參數日期 `t` 是該年的第幾天。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#DayOfYear)]
  [[play](https://go.dev/play/p/0hjqhTwFNlH)]
- **<big>IsWeekend</big>** : 檢查傳入的時間是否為週末。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#IsWeekend)]
  [[play](https://go.dev/play/p/cupRM5aZOIY)]
- **<big>NowDateOrTime</big>** : 回傳以指定格式與時區表示的目前日期時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NowDateOrTime)]
  [[play](https://go.dev/play/p/EZ-begEjtT0)]
- **<big>Timestamp</big>** : 回傳目前的秒級時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#Timestamp)]
- **<big>TimestampMilli</big>** : 回傳目前的毫秒級時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#TimestampMilli)]
  [[play](https://go.dev/play/p/4gvEusOTu1T)]
- **<big>TimestampMicro</big>** : 回傳目前的微秒級時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#TimestampMicro)]
  [[play](https://go.dev/play/p/2maANglKHQE)]
- **<big>TimestampNano</big>** : 回傳目前的奈秒級時間戳。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#TimestampNano)]
  [[play](https://go.dev/play/p/A9Oq_COrcCF)]
- **<big>TrackFuncTime</big>** : 追蹤函式的執行時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#TrackFuncTime)]
  [[play](https://go.dev/play/p/QBSEdfXHPTp)]
- **<big>DaysBetween</big>** : 回傳兩個時間之間相差的天數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#DaysBetween)]
  [[play](https://go.dev/play/p/qD6qGb3TbOy)]
- **<big>GenerateDatetimesBetween</big>** : 回傳兩個時間之間的字串切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GenerateDatetimesBetween)]
  [[play](https://go.dev/play/p/6kHBpAxD9ZC)]
- **<big>Min</big>** : 回傳給定時間中最早的時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#Min)]
  [[play](https://go.dev/play/p/MCIDvHNOGGb)]
- **<big>Max</big>** : 回傳給定時間中最晚的時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#Max)]
  [[play](https://go.dev/play/p/9m6JMk1LB7-)]
- **<big>MaxMin</big>** : 回傳給定時間中最晚與最早的時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#MaxMin)]
  [[play](https://go.dev/play/p/rbW51cDtM_2)]

<h3 id="datastructure"> 8. Datastructure 套件包含一些常見的資料結構，例如 list、linklist、stack、queue、set、tree、graph。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

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
import optional "github.com/duke-git/lancet/v2/datastructure/optional"
```

#### 結構清單:

- **<big>List</big>** : 線性表，以切片實作。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/list.md)]
- **<big>CopyOnWriteList</big>** : 以 Go 切片為基礎、具備執行緒安全的串列實作。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/copyonwritelist.md)]
- **<big>Link</big>** : 鏈結串列結構，包含單向鏈結與雙向鏈結。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/link.md)]
- **<big>Stack</big>** : 堆疊結構（fifo），包含陣列堆疊與鏈結堆疊。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/stack.md)]
- **<big>Queue</big>** : 佇列結構（filo），包含陣列佇列、環狀佇列、鏈結佇列與優先佇列。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/queue.md)]
- **<big>Set</big>** : 資料容器，類似切片，但集合中的元素不會重複。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/set.md)]
- **<big>Tree</big>** : 二元搜尋樹結構。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/tree.md)]
- **<big>Heap</big>** : 二元最大堆積。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/heap.md)]
- **<big>Hashmap</big>** : 雜湊表結構。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/hashmap.md)]
- **<big>Optional</big>** : Optional 容器。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/optional.md)]

<h3 id="eventbus"> 9. EventBus 是用於處理應用程式內事件的事件匯流排。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">Index</a></h3>

```go
import "github.com/duke-git/lancet/v2/eventbus"
```

#### 函式清單:

- **<big>NewEventBus</big>** : 建立一個 EventBus 實例。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#NewEventBus)]
  [[play](https://go.dev/play/p/gHbOPV_NUOJ)]
- **<big>Subscribe</big>** : 以指定的事件主題與監聽函式訂閱事件。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#Subscribe)]
  [[play](https://go.dev/play/p/EYGf_8cHei-)]
- **<big>Unsubscribe</big>** : 以指定的事件主題與監聽函式取消訂閱事件。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#Unsubscribe)]
  [[play](https://go.dev/play/p/Tmh7Ttfvprf)]
- **<big>Publish</big>** : 以指定的事件主題與資料內容發布事件。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#Publish)]
  [[play](https://go.dev/play/p/gHTtVexFSH9)]
- **<big>ClearListeners</big>** : 清除所有監聽者。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#ClearListeners)]
  [[play](https://go.dev/play/p/KBfBYlKPgqD)]
- **<big>ClearListenersByTopic</big>** : 依主題清除所有監聽者。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#ClearListenersByTopic)]
  [[play](https://go.dev/play/p/gvMljmJOZmU)]
- **<big>GetListenersCount</big>** : 回傳指定事件主題的監聽者數量。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#GetListenersCount)]
  [[play](https://go.dev/play/p/8VPJsMQgStM)]
- **<big>GetAllListenersCount</big>** : 回傳所有監聽者的總數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#GetAllListenersCount)]
  [[play](https://go.dev/play/p/PUlr0xcpEOz)]
- **<big>GetEvents</big>** : 回傳所有事件主題。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#GetEvents)]
  [[play](https://go.dev/play/p/etgjjcOtAjX)]
- **<big>SetErrorHandler</big>** : 設定錯誤處理函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/eventbus.md#SetErrorHandler)]
  [[play](https://go.dev/play/p/gmB0gnFe5mc)]

<h3 id="enum"> 10. enum 套件提供簡單的列舉實作。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">Index</a></h3>

```go
import "github.com/duke-git/lancet/v2/enum"
```

#### 函式清單:

- **<big>NewItem</big>** : 建立新的列舉項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#NewEventBus)]
  [[play](https://go.dev/play/p/8qNsLw01HD5)]
- **<big>NewItemsFromPairs</big>** : 從 Pair 結構的切片建立列舉項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#NewItemsFromPairs)]
  [[play](https://go.dev/play/p/xKnoGa7gnev)]
- **<big>Value</big>** : 回傳列舉項目的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Value)]
  [[play](https://go.dev/play/p/xKnoGa7gnev)]
- **<big>Name</big>** : 回傳列舉項目的名稱。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Name)]
  [[play](https://go.dev/play/p/xKnoGa7gnev)]
- **<big>Valid</big>** : 檢查列舉項目是否有效。若提供自訂檢查函式，則使用該函式驗證值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Valid)]
  [[play](https://go.dev/play/p/pA3lYY2VSm3)]
- **<big>MarshalJSON</big>** : json.Marshaler 介面的實作。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#MarshalJSON)]
  [[play](https://go.dev/play/p/zIZEdAnneB5)]
- **<big>NewRegistry</big>** : 建立新的列舉註冊表。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#NewRegistry)]
  [[play](https://go.dev/play/p/ABEXsYfJKMo)]
- **<big>Add</big>** : 將列舉項目加入註冊表。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Add)]
  [[play](https://go.dev/play/p/ABEXsYfJKMo)]
- **<big>Remove</big>** : 依值從註冊表移除列舉項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Remove)]
  [[play](https://go.dev/play/p/dSG84wQ3TuC)]
- **<big>Update</big>** : 依值更新註冊表中列舉項目的名稱。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Update)]
  [[play](https://go.dev/play/p/Ol0moT1J9Xl)]
- **<big>GetByValue</big>** : 依值取得列舉項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#GetByValue)]
  [[play](https://go.dev/play/p/niJ1U2KlE_m)]
- **<big>GetByName</big>** : 依名稱取得列舉項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#GetByName)]
  [[play](https://go.dev/play/p/49ie_gpqH0m)]
- **<big>Items</big>** : 回傳註冊表中所有列舉項目的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Items)]
  [[play](https://go.dev/play/p/lAJFAradbvQ)]
- **<big>Contains</big>** : 檢查註冊表中是否存在指定值的列舉項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Contains)]
  [[play](https://go.dev/play/p/_T-lPYkZn2j)]
- **<big>Size</big>** : 回傳註冊表中列舉項目的數量。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Size)]
  [[play](https://go.dev/play/p/TeDArWhlQe2)]
- **<big>Range</big>** : 迭代註冊表中所有列舉項目並套用給定的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Range)]
  [[play](https://go.dev/play/p/GPsZbQbefWN)]
- **<big>SortedItems</big>** : 回傳依給定 less 函式排序後的所有列舉項目切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#SortedItems)]
  [[play](https://go.dev/play/p/tN9RE_m_WEI)]
- **<big>Filter</big>** : 回傳符合給定判斷函式的列舉項目切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/enum.md#Filter)]
  [[play](https://go.dev/play/p/uTUpTdcyoCU)]

<h3 id="fileutil"> 11. Fileutil 套件實作了一些檔案操作的基本函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/fileutil"
```

#### 函式清單：

- **<big>ClearFile</big>** : 將空字串寫入目標檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ClearFile)]
  [[play](https://go.dev/play/p/NRZ0ZT-G94H)]
- **<big>CreateFile</big>** : 在指定路徑建立檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CreateFile)]
  [[play](https://go.dev/play/p/lDt8PEsTNKI)]
- **<big>CreateDir</big>** : 在絕對路徑建立目錄。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CreateDir)]
  [[play](https://go.dev/play/p/qUuCe1OGQnM)]
- **<big>CopyFile</big>** : 將來源檔案複製到目的檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CopyFile)]
  [[play](https://go.dev/play/p/Jg9AMJMLrJi)]
- **<big>CopyDir</big>** : 將來源目錄複製到目的目錄。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CopyDir)]
  [[play](https://go.dev/play/p/YAyFTA_UuPb)]
- **<big>FileMode</big>** : 回傳檔案的模式與權限。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#FileMode)]
  [[play](https://go.dev/play/p/2l2hI42fA3p)]
- **<big>MiMeType</big>** : 回傳檔案的 MIME 類型。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#MiMeType)]
  [[play](https://go.dev/play/p/bd5sevSUZNu)]
- **<big>IsExist</big>** : 檢查檔案或目錄是否存在。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsExist)]
  [[play](https://go.dev/play/p/nKKXt8ZQbmh)]
- **<big>IsLink</big>** : 檢查檔案是否為符號連結。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsLink)]
  [[play](https://go.dev/play/p/TL-b-Kzvf44)]
- **<big>IsDir</big>** : 檢查路徑是否為目錄。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsDir)]
  [[play](https://go.dev/play/p/WkVwEKqtOWk)]
- **<big>ListFileNames</big>** : 回傳路徑下所有的檔案名稱。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ListFileNames)]
  [[play](https://go.dev/play/p/Tjd7Y07rejl)]
- **<big>RemoveFile</big>** : 移除檔案，參數應為檔案路徑。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#RemoveFile)]
  [[play](https://go.dev/play/p/P2y0XW8a1SH)]
- **<big>RemoveDir</big>** : 刪除目錄。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#RemoveDir)]
  [[play](https://go.dev/play/p/Oa6KnPek2uy)]
- **<big>ReadFileToString</big>** : 回傳檔案內容的字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadFileToString)]
  [[play](https://go.dev/play/p/cmfwp_5SQTp)]
- **<big>ReadFileByLine</big>** : 逐行讀取檔案，回傳檔案內容的字串切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadFileByLine)]
  [[play](https://go.dev/play/p/svJP_7ZrBrD)]
- **<big>Zip</big>** : 為 fpath 建立 zip 檔案，fpath 可以是檔案或目錄。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#Zip)]
  [[play](https://go.dev/play/p/j-3sWBp8ik_P)]
- **<big>ZipAppendEntry</big>** : 將 fpath 指定的單一檔案或目錄附加到既有的 zip 檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ZipAppendEntry)]
- **<big>UnZip</big>** : 解壓縮 zip 檔案並儲存到目的路徑。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#UnZip)]
  [[play](https://go.dev/play/p/g0w34kS7B8m)]
- **<big>CurrentPath</big>** : 回傳目前的絕對路徑。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CurrentPath)]
  [[play](https://go.dev/play/p/s74a9iBGcSw)]
- **<big>IsZipFile</big>** : 檢查檔案是否為 zip 檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsZipFile)]
  [[play](https://go.dev/play/p/9M0g2j_uF_e)]
- **<big>FileSize</big>** : 回傳檔案大小（以位元組為單位）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#FileSize)]
  [[play](https://go.dev/play/p/H9Z05uD-Jjc)]
- **<big>MTime</big>** : 回傳檔案的修改時間（unix 時間戳）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#MTime)]
  [[play](https://go.dev/play/p/s_Tl7lZoAaY)]
- **<big>Sha</big>** : 回傳檔案的 sha 值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#Sha)]
  [[play](https://go.dev/play/p/VfEEcO2MJYf)]
- **<big>ReadCsvFile</big>** : 將檔案內容讀入切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadCsvFile)]
  [[play](https://go.dev/play/p/OExTkhGEd3_u)]
- **<big>WriteCsvFile</big>** : 將內容寫入目標 csv 檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#WriteCsvFile)]
- **<big>WriteMapsToCsv</big>** : 將 map 的切片寫入 csv 檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#WriteMapsToCsv)]
  [[play](https://go.dev/play/p/umAIomZFV1c)]
- **<big>WriteBytesToFile</big>** : 將位元組寫入目標檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#WriteBytesToFile)]
  [[play](https://go.dev/play/p/s7QlDxMj3P8)]
- **<big>WriteStringToFile</big>** : 將字串寫入目標檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#WriteStringToFile)]
  [[play](https://go.dev/play/p/GhLS6d8lH_g)]
- **<big>ReadFile</big>** : 讀取檔案或 URL。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadFile)]
- **<big>ChunkRead</big>** : 從檔案的指定偏移量讀取一個區塊，並回傳該區塊內的所有行。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ChunkRead)]
  [[play](https://go.dev/play/p/r0hPmKWhsgf)]
- **<big>ParallelChunkRead</big>** : 並行讀取檔案，並將每一批行資料傳送到指定的通道。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ParallelChunkRead)]  
  [[play](https://go.dev/play/p/teMXnCsdSEw)]
- **<big>GetExeOrDllVersion</big>** : 取得 Windows 作業系統上 exe 或 dll 檔案的版本。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#GetExeOrDllVersion)]
  [[play](https://go.dev/play/p/iLRrDBhE38E)]

<h3 id="formatter"> 12. Formatter 包含一些資料格式化函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/formatter"
```

#### 函式清單:

- **<big>Comma</big>** : 從右側每三位數加上逗號，並在前面加上前綴符號字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#Comma)]
  [[play](https://go.dev/play/p/eRD5k2vzUVX)]
- **<big>Pretty</big>** : 將資料美化輸出為 JSON 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#Pretty)]
  [[play](https://go.dev/play/p/YsciGj3FH2x)]
- **<big>PrettyToWriter</big>** : 將資料美化編碼並寫入 writer。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#PrettyToWriter)]
  [[play](https://go.dev/play/p/LPLZ3lDi5ma)]
- **<big>DecimalBytes</big>** : 依十進位標準（基數 1000）回傳人類可讀的位元組大小。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#DecimalBytes)]
  [[play](https://go.dev/play/p/FPXs1suwRcs)]
- **<big>BinaryBytes</big>** : 依二進位標準（基數 1024）回傳人類可讀的位元組大小。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#BinaryBytes)]
  [[play](https://go.dev/play/p/G9oHHMCAZxP)]
- **<big>ParseDecimalBytes</big>** : 將人類可讀的位元組大小字串轉換為其所代表的數量（基數 1000）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#ParseDecimalBytes)]
  [[play](https://go.dev/play/p/Am98ybWjvjj)]
- **<big>ParseBinaryBytes</big>** : 將人類可讀的位元組大小字串轉換為其所代表的數量（基數 1024）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#ParseBinaryBytes)]
  [[play](https://go.dev/play/p/69v1tTT62x8)]
- **<big>ParseCNAddress</big>** : 智慧解析中文地址字串，並擷取結構化資訊（省、市、區、街道、姓名、電話等），支援包含縣級市在內的各種地址格式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#ParseCNAddress)]
  [[play](https://go.dev/play/p/o5l09hQopEV)]
- **<big>ParsePersonInfo</big>** : 從地址字串中擷取使用者資訊（姓名、電話、身分證號、郵遞區號），並與位置地址分開。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#ParsePersonInfo)]
  [[play](https://go.dev/play/p/JO-uTlJlTy7)]

<h3 id="function"> 13. Function 套件可控制函式的執行流程，並支援部分函式程式設計。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/function"
```

#### 函式清單:

- **<big>After</big>** : 回傳一個函式，當該函式被呼叫超過 n 次後才會呼叫傳入的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#After)]
  [[play](https://go.dev/play/p/eRD5k2vzUVX)]
- **<big>Before</big>** : 回傳一個函式，當該函式被呼叫少於 n 次時才會呼叫傳入的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Before)]
  [[play](https://go.dev/play/p/0HqUDIFZ3IL)]
- **<big>CurryFn</big>** : 建立柯里化（curry）函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#CurryFn)]
  [[play](https://go.dev/play/p/5HopfDwANKX)]
- **<big>Compose</big>** : 由右至左組合多個函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Compose)]
  [[play](https://go.dev/play/p/KKfugD4PKYF)]
- **<big>Delay</big>** : 延遲指定時間後呼叫函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Delay)]
  [[play](https://go.dev/play/p/Ivtc2ZE-Tye)]
- **<big>Debounced<sup>deprecated</sup></big>** : 建立防抖（debounce）函式，自上次呼叫該防抖函式起，經過 wait 時間後才會呼叫 fn。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Debounced)]
  [[play](https://go.dev/play/p/absuEGB_GN7)]
- **<big>Debounce</big>** : 建立所提供函式的防抖版本。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Debounce)]
  [[play](https://go.dev/play/p/-dGFrYn_1Zi)]
- **<big>Throttle</big>** : 建立所提供函式的節流（throttle）版本。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Throttle)]
  [[play](https://go.dev/play/p/HpoMov-tJSN)]
- **<big>Schedule</big>** : 每隔指定時間呼叫函式，直到關閉所回傳的布林通道為止。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Schedule)]
  [[play](https://go.dev/play/p/hbON-Xeyn5N)]
- **<big>Pipeline</big>** : 接收一組函式並回傳一個函式，其參數會依序傳入這些函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Pipeline)]
  [[play](https://go.dev/play/p/mPdUVvj6HD6)]
- **<big>AcceptIf</big>** : 回傳另一個與 apply 函式簽章相同的函式，但額外包含一個布林值以表示成功或失敗。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#AcceptIf)]
  [[play](https://go.dev/play/p/XlXHHtzCf7d)]
- **<big>And</big>** : 回傳由多個判斷函式組合而成、代表邏輯 AND 的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#And)]
  [[play](https://go.dev/play/p/dTBHJMQ0zD2)]
- **<big>Or</big>** : 回傳由多個判斷函式組合而成、代表邏輯 OR 的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Or)]
  [[play](https://go.dev/play/p/LitCIsDFNDA)]
- **<big>Negate</big>** : 回傳代表此判斷函式邏輯否定的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Negate)]
  [[play](https://go.dev/play/p/jbI8BtgFnVE)]
- **<big>Nor</big>** : 回傳由多個判斷函式組合而成、代表邏輯 NOR 的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Nor)]
  [[play](https://go.dev/play/p/2KdCoBEOq84)]
- **<big>Nand</big>** : 回傳由多個判斷函式組合而成、代表邏輯 NAND 的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Nand)]
  [[play](https://go.dev/play/p/Rb-FdNGpgSO)]
- **<big>Xnor</big>** : 回傳由多個判斷函式組合而成、代表邏輯 XNOR 的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Xnor)]
  [[play](https://go.dev/play/p/FJxko8SFbqc)]
- **<big>Watcher</big>** : Watcher 用於記錄程式碼執行時間，可啟動／停止／重設計時器，並取得函式執行所耗費的時間。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Watcher)]
  [[play](https://go.dev/play/p/l2yrOpCLd1I)]

<h3 id="maputil"> 14. Maputil 套件包含一些操作 map 的函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/maputil"
```

#### 函式清單:

- **<big>MapTo</big>** : 快速將任意值對應到結構或任何基本型別。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#MapTo)]
  [[play](https://go.dev/play/p/4K7KBEPgS5M)]
- **<big>ForEach</big>** : 對 map 中的每個鍵值對執行 iteratee 函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ForEach)]
  [[play](https://go.dev/play/p/OaThj6iNVXK)]
- **<big>Filter</big>** : 迭代 map，回傳包含所有通過判斷函式之鍵值對的新 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Filter)]
  [[play](https://go.dev/play/p/fSvF3wxuNG7)]
- **<big>FilterByKeys</big>** : 迭代 map，回傳鍵皆為給定鍵的新 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#FilterByKeys)]
  [[play](https://go.dev/play/p/7ov6BJHbVqh)]
- **<big>FilterByValues</big>** : 迭代 map，回傳值皆為給定值的新 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#FilterByValues)]
  [[play](https://go.dev/play/p/P3-9MdcXegR)]
- **<big>OmitBy</big>** : Filter 的相反，移除所有使判斷函式回傳 true 的 map 元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OmitBy)]
  [[play](https://go.dev/play/p/YJM4Hj5hNwm)]
- **<big>OmitByKeys</big>** : FilterByKeys 的相反，擷取鍵未被排除的所有 map 元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OmitByKeys)]
  [[play](https://go.dev/play/p/jXGrWDBfSRp)]
- **<big>OmitByValues</big>** : FilterByValues 的相反，移除所有值在給定切片中的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OmitByValues)]
  [[play](https://go.dev/play/p/XB7Y10uw20_U)]
- **<big>Intersect</big>** : 迭代多個 map，回傳包含所有給定 map 中鍵值對的新 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Intersect)]
  [[play](https://go.dev/play/p/Zld0oj3sjcC)]
- **<big>Keys</big>** : 回傳 map 的鍵的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Keys)]
  [[play](https://go.dev/play/p/xNB5bTb97Wd)]
- **<big>KeysBy</big>** : 建立一個切片，其元素為對 map 的每個鍵呼叫 mapper 函式的結果。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#KeysBy)]
  [[play](https://go.dev/play/p/hI371iB8Up8)]
- **<big>Merge</big>** : 合併 map，後面的鍵會覆蓋前面的鍵。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Merge)]
  [[play](https://go.dev/play/p/H95LENF1uB-)]
- **<big>Minus</big>** : 建立一個 map，其鍵存在於 mapA 但不在 mapB 中。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Minus)]
  [[play](https://go.dev/play/p/3u5U9K7YZ9m)]
- **<big>Values</big>** : 回傳 map 的值的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Values)]
  [[play](https://go.dev/play/p/CBKdUc5FTW6)]
- **<big>ValuesBy</big>** : 建立一個切片，其元素為對 map 的每個值呼叫 mapper 函式的結果。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ValuesBy)]
  [[play](https://go.dev/play/p/sg9-oRidh8f)]
- **<big>MapKeys</big>** : 透過處理鍵，將 map 轉換為另一種型別的 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#MapKeys)]
  [[play](https://go.dev/play/p/8scDxWeBDKd)]
- **<big>MapValues</big>** : 透過處理值，將 map 轉換為另一種型別的 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#MapValues)]
  [[play](https://go.dev/play/p/g92aY3fc7Iw)]
- **<big>Entries</big>** : 將 map 轉換為鍵值對的陣列。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Entries)]
  [[play](https://go.dev/play/p/Ltb11LNcElY)]
- **<big>FromEntries</big>** : 依據鍵值對的切片建立 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#FromEntries)]
  [[play](https://go.dev/play/p/fTdu4sCNjQO)]
- **<big>Transform</big>** : 將 map 轉換為另一種型別的 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Transform)]
  [[play](https://go.dev/play/p/P6ovfToM3zj)]
- **<big>IsDisjoint</big>** : 檢查兩個 map 是否不相交（沒有共同的鍵）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#IsDisjoint)]
  [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
- **<big>HasKey</big>** : 檢查 map 是否含有指定的鍵。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#HasKey)]
  [[play](https://go.dev/play/p/isZZHOsDhFc)]
- **<big>GetOrSet</big>** : 回傳給定鍵的值，若不存在則設定為給定的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#GetOrSet)]
  [[play](https://go.dev/play/p/IVQwO1OkEJC)]
- **<big>MapToStruct</big>** : 將 map 轉換為結構。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#MapToStruct)]
  [[play](https://go.dev/play/p/7wYyVfX38Dp)]
- **<big>ToSortedSlicesDefault</big>** : 將 map 轉換為依鍵排序的兩個切片：一個存放鍵，另一個存放值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ToSortedSlicesDefault)]
  [[play](https://go.dev/play/p/43gEM2po-qy)]
- **<big>ToSortedSlicesWithComparator</big>** : 將 map 轉換為使用自訂比較函式、依鍵排序的兩個切片：一個存放鍵，另一個存放值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ToSortedSlicesWithComparator)]
  [[play](https://go.dev/play/p/0nlPo6YLdt3)]
- **<big>NewOrderedMap</big>** : 建立新的 OrderedMap。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#NewOrderedMap)]
  [[play](https://go.dev/play/p/Y4ZJ_oOc1FU)]
- **<big>OrderedMap_Set</big>** : 為有序 map 設定給定的鍵值對。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Set)]
  [[play](https://go.dev/play/p/Y4ZJ_oOc1FU)]
- **<big>OrderedMap_Get</big>** : 回傳給定鍵的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Get)]
  [[play](https://go.dev/play/p/Y4ZJ_oOc1FU)]
- **<big>OrderedMap_Delete</big>** : 刪除給定鍵的鍵值對。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Delete)]
  [[play](ttps://go.dev/play/p/5bIi4yaZ3K-)]
- **<big>OrderedMap_Clear</big>** : 清空有序 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Clear)]
  [[play](https://go.dev/play/p/8LwoJyEfuFr)]
- **<big>OrderedMap_Front</big>** : 回傳第一個鍵值對。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Front)]
  [[play](https://go.dev/play/p/ty57XSimpoe)]
- **<big>OrderedMap_Back</big>** : 回傳最後一個鍵值對。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Back)]
  [[play](https://go.dev/play/p/rQMjp1yQmpa)]
- **<big>OrderedMap_Range</big>** : 對每個鍵值對呼叫給定的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Range)]
  [[play](https://go.dev/play/p/U-KpORhc7LZ)]
- **<big>OrderedMap_Keys</big>** : 依序回傳所有的鍵。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Keys)]
  [[play](https://go.dev/play/p/Vv_y9ExKclA)]
- **<big>OrderedMap_Values</big>** : 依序回傳所有的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Values)]
  [[play](https://go.dev/play/p/TWj5n1-PUfx)]
- **<big>OrderedMap_Elements</big>** : 依序回傳所有的鍵值對。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Elements)]
  [[play](https://go.dev/play/p/4BHG4kKz6bB)]
- **<big>OrderedMap_Len</big>** : 回傳鍵值對的數量。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Len)]
  [[play](https://go.dev/play/p/cLe6z2VX5N-)]
- **<big>OrderedMap_Contains</big>** : 若給定的鍵存在則回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Contains)]
  [[play](https://go.dev/play/p/QuwqqnzwDNX)]
- **<big>OrderedMap_Iter</big>** : 回傳一個依序產出鍵值對的通道。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_Iter)]
  [[play](https://go.dev/play/p/tlq2tdvicPt)]
- **<big>OrderedMap_ReverseIter</big>** : 回傳一個以相反順序產出鍵值對的通道。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_ReverseIter)]
  [[play](https://go.dev/play/p/8Q0ssg6hZzO)]
- **<big>OrderedMap_SortByKey</big>** : 依給定的 less 函式依鍵對 map 排序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_SortByKey)]
  [[play](https://go.dev/play/p/N7hjD_alZPq)]
- **<big>OrderedMap_MarshalJSON</big>** : 實作 json.Marshaler 介面。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_MarshalJSON)]
  [[play](https://go.dev/play/p/C-wAwydIAC7)]
- **<big>OrderedMap_UnmarshalJSON</big>** : 實作 json.Unmarshaler 介面。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OrderedMap_UnmarshalJSON)]
  [[play](https://go.dev/play/p/t_pkwerIRVx)]
- **<big>NewConcurrentMap</big>** : 以指定的分片數量建立 ConcurrentMap。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#NewConcurrentMap)]
  [[play](https://go.dev/play/p/3PenTPETJT0)]
- **<big>ConcurrentMap_Set</big>** : 設定鍵對應的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Set)]
  [[play](https://go.dev/play/p/3PenTPETJT0)]
- **<big>ConcurrentMap_Get</big>** : 取得 map 中某個鍵所儲存的值，若不存在則回傳 nil。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Get)]
  [[play](https://go.dev/play/p/3PenTPETJT0)]
- **<big>ConcurrentMap_GetOrSet</big>** : 若鍵存在，則回傳該鍵既有的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_GetOrSet)]
  [[play](https://go.dev/play/p/aDcDApOK01a)]
- **<big>ConcurrentMap_Delete</big>** : 刪除某個鍵對應的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Delete)]
  [[play](https://go.dev/play/p/uTIJZYhpVMS)]
- **<big>ConcurrentMap_GetAndDelete</big>** :若鍵存在，則回傳該鍵既有的值，然後刪除該鍵對應的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_GetAndDelete)]
  [[play](https://go.dev/play/p/ZyxeIXSZUiM)]
- **<big>ConcurrentMap_Has</big>** : 檢查 map 中是否有某個鍵對應的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Has)]
  [[play](https://go.dev/play/p/C8L4ul9TVwf)]
- **<big>ConcurrentMap_Range</big>** : 對 map 中每個分片內存在的每個鍵與值，依序呼叫 iterator。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Range)]
  [[play](https://go.dev/play/p/iqcy7P8P0Pr)]
- **<big>SortByKey</big>** : 依鍵對 map 排序，並回傳鍵已排序的新 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#SortByKey)]
  [[play](https://go.dev/play/p/PVdmBSnm6P_W)]
- **<big>GetOrDefault</big>** : 回傳給定鍵的值，若該鍵不存在則回傳預設值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#GetOrDefault)]
  [[play](https://go.dev/play/p/99QjSYSBdiM)]
- **<big>FindValuesBy</big>** : 回傳 map 中符合給定判斷函式之值的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#FindValuesBy)]
  [[play](https://go.dev/play/p/bvNwNBZDm6v)]
- **<big>ToMarkdownTable</big>** : 將 map 切片資料轉換為 Markdown 表格字串，支援自訂表頭顯示名稱與欄位顯示順序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ToMarkdownTable)]
  [[play](https://go.dev/play/p/w_pSLfeyEB5)]

<h3 id="mathutil"> 15. Mathutil 套件實作了一些數學計算函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### 函式清單:

- **<big>Average</big>** :回傳數值的平均值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Average)]
  [[play](https://go.dev/play/p/Vv7LBwER-pz)]
- **<big>Exponent</big>** : 計算 int64 的 x^n。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Exponent)]
  [[play](https://go.dev/play/p/uF3HGNPk8wr)]
- **<big>Fibonacci</big>** :計算 int 的 n 之前的費波那契數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Fibonacci)]
  [[play](https://go.dev/play/p/IscseUNMuUc)]
- **<big>Factorial</big>** : 計算 uint 的 x!。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Factorial)]
  [[play](https://go.dev/play/p/tt6LdOK67Nx)]
- **<big>Max</big>** : 回傳數值中的最大值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Max)]
  [[play](https://go.dev/play/p/cN8DHI0rTkH)]
- **<big>MaxBy</big>** : 使用給定的比較函式回傳切片中的最大值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#MaxBy)]
  [[play](https://go.dev/play/p/pbe2MT-7DV2)]
- **<big>Min</big>** : 回傳數值中的最小值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Min)]
  [[play](https://go.dev/play/p/21BER_mlGUj)]
- **<big>MinBy</big>** : 使用給定的比較函式回傳切片中的最小值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#MinBy)]
  [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
- **<big>Percent</big>** : 計算數值佔總數的百分比。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Percent)]
  [[play](https://go.dev/play/p/s0NdFCtwuyd)]
- **<big>RoundToFloat</big>** : 對 float64 四捨五入至小數點後 n 位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#RoundToFloat)]
  [[play](https://go.dev/play/p/ghyb528JRJL)]
- **<big>RoundToString</big>** : 對 float64 四捨五入至小數點後 n 位，並回傳字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#RoundToString)]
  [[play](https://go.dev/play/p/kZwpBRAcllO)]
- **<big>TruncRound</big>** : 對 int64 截去小數點後 n 位（不進行四捨五入）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#TruncRound)]
  [[play](https://go.dev/play/p/aumarSHIGzP)]
- **<big>CeilToFloat</big>** : 將浮點數無條件進位至小數點後 n 位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#CeilToFloat)]
  [[play](https://go.dev/play/p/8hOeSADZPCo)]
- **<big>CeilToString</big>** : 將浮點數無條件進位至小數點後 n 位，然後轉換為字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#CeilToString)]
  [[play](https://go.dev/play/p/wy5bYEyUKKG)]
- **<big>FloorToFloat</big>** : 將浮點數無條件捨去至小數點後 n 位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#FloorToFloat)]
  [[play](https://go.dev/play/p/vbCBrQHZEED)]
- **<big>FloorToString</big>** : 將浮點數無條件捨去至小數點後 n 位，然後轉換為字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#FloorToString)]
  [[play](https://go.dev/play/p/Qk9KPd2IdDb)]
- **<big>Range</big>** : 從 start 開始建立指定數量的數字切片，元素間距為 1。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Range)]
  [[play](https://go.dev/play/p/9ke2opxa8ZP)]
- **<big>RangeWithStep</big>** : 依指定的間距建立從 start 到 end 的數字切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Range)]
  [[play](https://go.dev/play/p/akLWz0EqOSM)]
- **<big>AngleToRadian</big>** : 將角度值轉換為弧度值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#AngleToRadian)]
  [[play](https://go.dev/play/p/CIvlICqrHql)]
- **<big>RadianToAngle</big>** : 將弧度值轉換為角度值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#RadianToAngle)]
  [[play](https://go.dev/play/p/dQtmOTUOMgi)]
- **<big>PointDistance</big>** : 取得兩點之間的距離。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#PointDistance)]
  [[play](https://go.dev/play/p/RrG4JIaziM8)]
- **<big>IsPrime</big>** : 檢查數字是否為質數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#IsPrime)]
  [[play](https://go.dev/play/p/Rdd8UTHZJ7u)]
- **<big>GCD</big>** : 回傳整數的最大公因數（GCD）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#GCD)]
  [[play](https://go.dev/play/p/CiEceLSoAKB)]
- **<big>LCM</big>** : 回傳整數的最小公倍數（LCM）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#LCM)]
  [[play](https://go.dev/play/p/EjcZxfY7G_g)]
- **<big>Cos</big>** : 回傳弧度參數的餘弦值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Cos)]
  [[play](https://go.dev/play/p/Sm89LoIfvFq)]
- **<big>Sin</big>** : 回傳弧度參數的正弦值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Sin)]
  [[play](https://go.dev/play/p/TWMQlMywDsP)]
- **<big>Log</big>** : 回傳以 n 為底的對數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Log)]
  [[play](https://go.dev/play/p/_d4bi8oyhat)]
- **<big>Sum</big>** : 回傳傳入數字的總和。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Sum)]
  [[play](https://go.dev/play/p/1To2ImAMJA7)]
- **<big>Abs</big>** : 回傳參數數字的絕對值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Sum)]
  [[play](https://go.dev/play/p/fsyBh1Os-1d)]
- **<big>Div</big>** : 回傳 x 除以 y 的結果。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Div)]
  [[play](https://go.dev/play/p/WLxDdGXXYat)]
- **<big>Variance</big>** : 回傳數值的變異數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Variance)]
  [[play](https://go.dev/play/p/uHuV4YgXf8F)]
- **<big>StdDev</big>** : 回傳數值的標準差。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#StdDev)]
  [[play](https://go.dev/play/p/FkNZDXvHD2l)]
- **<big>Permutation</big>** : 計算 P(n, k)。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Permutation)]
  [[play](https://go.dev/play/p/MgobwH_FOxj)]
- **<big>Combination</big>** : 計算 C(n, k)。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Combination)]
  [[play](https://go.dev/play/p/ENFQRDQUFi9)]

<h3 id="netutil"> 16. Netutil 套件包含取得網路資訊與發送 HTTP 請求的函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/netutil"
```

#### 函式清單:

- **<big>ConvertMapToQueryString</big>** : 將 map 轉換為已排序的 URL 查詢字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#ConvertMapToQueryString)]
  [[play](https://go.dev/play/p/jnNt_qoSnRi)]
- **<big>EncodeUrl</big>** : 對 URL 編碼(?a=1&b=[2] -> ?a=1&b=%5B2%5D)。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#EncodeUrl)]
  [[play](https://go.dev/play/p/bsZ6BRC4uKI)]
- **<big>GetInternalIp</big>** : 回傳內部 IPv4。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetInternalIp)]
  [[play](https://go.dev/play/p/5mbu-gFp7ei)]
- **<big>GetIps</big>** : 回傳目前系統的所有 IPv4。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetIps)]
  [[play](https://go.dev/play/p/NUFfcEmukx1)]
- **<big>GetMacAddrs</big>** : 回傳目前系統的 MAC 位址。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetMacAddrs)]
  [[play](https://go.dev/play/p/Rq9UUBS_Xp1)]
- **<big>GetPublicIpInfo</big>** : 回傳[公開 IP 資訊](http://ip-api.com/json/)。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetPublicIpInfo)]
  [[play](https://go.dev/play/p/YDxIfozsRHR)]
- **<big>GetRequestPublicIp</big>** : 回傳 HTTP 請求的公開 IP。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetRequestPublicIp)]
  [[play](https://go.dev/play/p/kxU-YDc_eBo)]
- **<big>IsPublicIP</big>** : 驗證 IP 是否為公開 IP。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsPublicIP)]
  [[play](https://go.dev/play/p/nmktSQpJZnn)]
- **<big>IsInternalIP</big>** : 驗證 IP 是否為內網 IP。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsInternalIP)]
  [[play](https://go.dev/play/p/sYGhXbgO4Cb)]
- **<big>HttpRequest</big>** : 供 HttpClient 發送請求使用的組合式 HTTP 請求。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpRequest)]
  [[play](https://go.dev/play/p/jUSgynekH7G)]
- **<big>HttpClient</big>** : 用於發送 HTTP 請求的 HTTP 用戶端工具。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpClient)]
  [[play](https://go.dev/play/p/jUSgynekH7G)]
- **<big>SendRequest</big>** : 發送 HTTP 請求。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#SendRequest)]
  [[play](https://go.dev/play/p/jUSgynekH7G)]
- **<big>DecodeResponse</big>** : 將 HTTP 回應解碼為目標物件。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#DecodeResponse)]
  [[play](https://go.dev/play/p/jUSgynekH7G)]
- **<big>StructToUrlValues</big>** : 將結構轉換為 URL values。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#StructToUrlValues)]
  [[play](https://go.dev/play/p/pFqMkM40w9z)]
- **<big>HttpGet<sup>deprecated</sup></big>** : 發送 HTTP GET 請求。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpGet)]
- **<big>HttpDelete<sup>deprecated</sup></big>** : 發送 HTTP DELETE 請求。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpDelete)]
- **<big>HttpPost<sup>deprecated</sup></big>** : 發送 HTTP POST 請求。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpPost)]
- **<big>HttpPut<sup>deprecated</sup></big>** : 發送 HTTP PUT 請求。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpPut)]
- **<big>HttpPatch<sup>deprecated</sup></big>** : 發送 HTTP PATCH 請求。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpPatch)]
- **<big>ParseHttpResponse</big>** : 將 HTTP 回應解碼為目標物件。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#ParseHttpResponse)]
- **<big>DownloadFile</big>** : 將 URL 上的檔案下載到本機檔案。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#DownloadFile)]
- **<big>UploadFile</big>** : 將檔案上傳到伺服器。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#UploadFile)]
- **<big>IsPingConnected</big>** : 檢查是否能 ping 通指定的主機。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsPingConnected)]
  [[play](https://go.dev/play/p/q8OzTijsA87)]
- **<big>IsTelnetConnected</big>** : 檢查是否能 telnet 連上指定的主機。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsTelnetConnected)]
  [[play](https://go.dev/play/p/yiLCGtQv_ZG)]
- **<big>BuildUrl</big>** : 依給定的參數建構 URL。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#BuildUrl)]
  [[play](https://go.dev/play/p/JLXl1hZK7l4)]
- **<big>AddQueryParams</big>** : 將查詢參數加入給定的 URL。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#AddQueryParams)]
  [[play](https://go.dev/play/p/JLXl1hZK7l4)]

<h3 id="pointer"> 17. Pointer 套件包含一些操作 Go 指標的工具函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/pointer"
```

#### 函式清單:

- **<big>ExtractPointer</big>** : 依給定的介面型別回傳底層的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#ExtractPointer)]
  [[play](https://go.dev/play/p/D7HFjeWU2ZP)]
- **<big>Of</big>** : 回傳指向值 `v` 的指標。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#Of)]
  [[play](https://go.dev/play/p/HFd70x4DrMj)]
- **<big>Unwrap</big>** : 回傳指標所指向的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#Unwrap)]
  [[play](https://go.dev/play/p/cgeu3g7cjWb)]
- **<big>UnwrapOr</big>** : UnwrapOr 回傳指標所指向的值，若指標為 nil 則回傳備用值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#UnwrapOr)]
  [[play](https://go.dev/play/p/mmNaLC38W8C)]
- **<big>UnwarpOrDefault</big>** : UnwarpOrDefault 回傳指標所指向的值，若指標為 nil 則回傳預設值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#UnwrapOrDefault)]
  [[play](https://go.dev/play/p/ZnGIHf8_o4E)]

<h3 id="random"> 18. Random 套件實作了產生隨機整數與字串的基本函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/random"
```

#### 函式清單:

- **<big>RandBytes</big>** : 產生隨機位元組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandBytes)]
  [[play](https://go.dev/play/p/EkiLESeXf8d)]
- **<big>RandInt</big>** : 產生介於 min 與 max 之間的隨機整數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandInt)]
  [[play](https://go.dev/play/p/pXyyAAI5YxD)]
- **<big>RandString</big>** : 產生指定長度的隨機字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandString)]
  [[play](https://go.dev/play/p/W2xvRUXA7Mi)]
- **<big>RandUpper</big>** : 產生隨機的大寫字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandUpper)]
  [[play](https://go.dev/play/p/29QfOh0DVuh)]
- **<big>RandLower</big>** : 產生隨機的小寫字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandLower)]
  [[play](https://go.dev/play/p/XJtZ471cmtI)]
- **<big>RandNumeral</big>** : 產生指定長度的隨機數字字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandNumeral)]
  [[play](https://go.dev/play/p/g4JWVpHsJcf)]
- **<big>RandNumeralOrLetter</big>** : 產生隨機的數字或字母字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandNumeralOrLetter)]
  [[play](https://go.dev/play/p/19CEQvpx2jD)]
- **<big>UUIdV4</big>** : 依據 RFC 4122 產生版本 4 的隨機 UUID。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#UUIdV4)]
  [[play](https://go.dev/play/p/_Z9SFmr28ft)]
- **<big>RandUniqueIntSlice</big>** : 產生不重複的隨機整數切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandUniqueIntSlice)]
  [[play](https://go.dev/play/p/uBkRSOz73Ec)]
- **<big>RandSymbolChar</big>** : 產生指定長度的隨機符號字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandSymbolChar)]
  [[play](https://go.dev/play/p/Im6ZJxAykOm)]
- **<big>RandFloat</big>** : 依指定精度產生介於 [min, max) 之間的隨機 float64 數字。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandFloat)]
  [[play](https://go.dev/play/p/zbD_tuobJtr)]
- **<big>RandFloats</big>** : 產生不重複的隨機 float64 數字切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandFloats)]
  [[play](https://go.dev/play/p/I3yndUQ-rhh)]
- **<big>RandStringSlice</big>** : 依據字元集產生長度為 strLen 的隨機字串切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandStringSlice)]
  [[play](https://go.dev/play/p/2_-PiDv3tGn)]
- **<big>RandBool</big>** : 產生隨機布林值（true 或 false）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandBool)]
  [[play](https://go.dev/play/p/to6BLc26wBv)]
- **<big>RandBoolSlice</big>** : 產生指定長度的隨機布林切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandBoolSlice)]
  [[play](https://go.dev/play/p/o-VSjPjnILI)]
- **<big>RandIntSlice</big>** : 產生隨機整數切片，數字範圍為 [min, max)。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandIntSlice)]
  [[play](https://go.dev/play/p/GATTQ5xTEG8)]
- **<big>RandFromGivenSlice</big>** : 從給定切片中產生一個隨機元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandFromGivenSlice)]
  [[play](https://go.dev/play/p/UrkWueF6yYo)]
- **<big>RandSliceFromGivenSlice</big>** : 從給定切片中產生長度為 num 的隨機切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandSliceFromGivenSlice)]
  [[play](https://go.dev/play/p/68UikN9d6VT)]
- **<big>RandNumberOfLength</big>** : 產生指定長度的隨機整數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandNumberOfLength)]
  [[play](https://go.dev/play/p/oyZbuV7bu7b)]

<h3 id="retry"> 19. Retry 套件用於重複執行函式，直到成功或被 context 取消為止。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### 函式清單:

- **<big>Context</big>** : 設定重試的 context 組態選項。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#Context)]
  [[play](https://go.dev/play/p/xnAOOXv9GkS)]
- **<big>Retry</big>** : 重複執行 retryFunc，直到成功或被 context 取消為止。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#Retry)]
  [[play](https://go.dev/play/p/nk2XRmagfVF)]
- **<big>RetryFunc</big>** : 重試所要執行的函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryFunc)]
  [[play](https://go.dev/play/p/nk2XRmagfVF)]
- **<big>RetryDuration</big>** : 設定重試的時間間隔。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryDuration)]
  [[play](https://go.dev/play/p/nk2XRmagfVF)]
- **<big>RetryTimes</big>** : 設定重試的次數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryTimes)]
  [[play](https://go.dev/play/p/ssfVeU2SwLO)]
- **<big>BackoffStrategy</big>** : 定義計算退避間隔方法的介面。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#BackoffStrategy)]
- **<big>RetryWithCustomBackoff</big>** : 設定任意自訂的退避策略。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryWithCustomBackoff)]
  [[play](https://go.dev/play/p/jIm_o2vb5Y4)]
- **<big>RetryWithLinearBackoff</big>** : 設定線性退避策略。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryWithLinearBackoff)]
  [[play](https://go.dev/play/p/PDet2ZQZwcB)]
- **<big>RetryWithExponentialWithJitterBackoff</big>** : 設定指數退避策略。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryWithExponentialWithJitterBackoff)]
  [[play](https://go.dev/play/p/xp1avQmn16X)]

<h3 id="slice"> 20. Slice 套件包含一些操作切片的函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### 函式清單:

- **<big>AppendIfAbsent</big>** : 若項目不存在，則將其附加到切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#AppendIfAbsent)]
  [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
- **<big>Contain</big>** : 檢查值是否在切片中。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Contain)]
  [[play](https://go.dev/play/p/_454yEHcNjf)]
- **<big>ContainBy</big>** : 若判斷函式回傳 true 則回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ContainBy)]
  [[play](https://go.dev/play/p/49tkHfX4GNc)]
- **<big>ContainSubSlice</big>** : 檢查切片是否包含給定的子切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ContainSubSlice)]
  [[play](https://go.dev/play/p/bcuQ3UT6Sev)]
- **<big>ContainAny</big>** : 檢查切片是否包含目標切片中的任一元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ContainAny)]
  [[play](https://go.dev/play/p/4xoxhc9XSSw)]
- **<big>Chunk</big>** : 建立依 size 長度分組的元素切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Chunk)]
  [[play](https://go.dev/play/p/b4Pou5j2L_C)]
- **<big>Compact</big>** : 建立移除所有假值的切片。false、nil、0 與 "" 皆為假值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Compact)]
  [[play](https://go.dev/play/p/pO5AnxEr3TK)]
- **<big>Concat</big>** : 建立一個新切片，將切片與任何額外的切片串接起來。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Concat)]
  [[play](https://go.dev/play/p/gPt-q7zr5mk)]
- **<big>Count</big>** : 回傳給定項目在切片中出現的次數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Count)]
  [[play](https://go.dev/play/p/Mj4oiEnQvRJ)]
- **<big>CountBy</big>** : 使用判斷函式迭代切片的元素，回傳所有符合元素的數量。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#CountBy)]
  [[play](https://go.dev/play/p/tHOccTMDZCC)]
- **<big>Difference</big>** : 建立一個切片，其元素存在於切片中但不在比較用的切片中。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Difference)]
  [[play](https://go.dev/play/p/VXvadzLzhDa)]
- **<big>DifferenceBy</big>** : 接收 iteratee，對切片的每個元素與值呼叫它以產生比較的依據。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DifferenceBy)]
  [[play](https://go.dev/play/p/DiivgwM5OnC)]
- **<big>DifferenceWith</big>** : 接收 comparator，呼叫它來比較切片的元素與值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DifferenceWith)]
  [[play](https://go.dev/play/p/v2U2deugKuV)]
- **<big>DeleteAt</big>** : 刪除切片中指定索引的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DeleteAt)]
  [[play](https://go.dev/play/p/800B1dPBYyd)]
- **<big>DeleteRange</big>** : 刪除切片中從起始索引到結束索引（不含）的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DeleteRange)]
  [[play](https://go.dev/play/p/945HwiNrnle)]
- **<big>Drop</big>** : 從切片開頭丟棄 n 個元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Drop)]
  [[play](https://go.dev/play/p/jnPO2yQsT8H)]
- **<big>DropRight</big>** : 從切片結尾丟棄 n 個元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DropRight)]
  [[play](https://go.dev/play/p/8bcXvywZezG)]
- **<big>DropWhile</big>** : 當判斷函式回傳 true 時，從切片開頭丟棄 n 個元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DropWhile)]
  [[play](https://go.dev/play/p/4rt252UV_qs)]
- **<big>DropRightWhile</big>** : 當判斷函式回傳 true 時，從切片結尾丟棄 n 個元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DropRightWhile)]
  [[play](https://go.dev/play/p/6wyK3zMY56e)]
- **<big>Equal</big>** : 檢查兩個切片是否相等：長度相同，且所有元素的順序與值皆相等。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Equal)]
  [[play](https://go.dev/play/p/WcRQJ37ifPa)]
- **<big>EqualWith</big>** : 使用比較函式檢查兩個切片是否相等。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#EqualWith)]
  [[play](https://go.dev/play/p/b9iygtgsHI1)]
- **<big>EqualUnordered</big>** : 檢查兩個切片是否相等：長度相同，且所有元素的值皆相等（不考慮順序）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#EqualUnordered)]
  [[play](https://go.dev/play/p/n8fSc2w8ZgX)]
- **<big>Every</big>** : 若切片中所有的值都通過判斷函式則回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Every)]
  [[play](https://go.dev/play/p/R8U6Sl-j8cD)]
- **<big>Filter</big>** : 迭代切片的元素，回傳所有通過判斷函式之元素的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Filter)]
  [[play](https://go.dev/play/p/SdPna-7qK4T)]
- **<big>FilterMap</big>** : 回傳對給定切片同時套用過濾與映射的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FilterMap)]
  [[play](https://go.dev/play/p/J94SZ_9MiIe)]
- **<big>Find<sup>deprecated</sup></big>** : 迭代切片的元素，回傳第一個通過判斷函式真值測試的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Find)]
  [[play](https://go.dev/play/p/CBKeBoHVLgq)]
- **<big>FindBy</big>** : 迭代切片的元素，回傳第一個通過判斷函式真值測試的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FindBy)]
  [[play](https://go.dev/play/p/n1lysBYl-GB)]
- **<big>FindLast<sup>deprecated</sup></big>** : 回傳最後一個通過判斷函式真值測試的項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FindLast)]
  [[play](https://go.dev/play/p/FFDPV_j7URd)]
- **<big>FindLastBy</big>** : 迭代切片的元素，回傳最後一個通過判斷函式真值測試的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FindLastBy)]
  [[play](https://go.dev/play/p/8iqomzyCl_s)]
- **<big>Flatten</big>** : 將切片攤平一層。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Flatten)]
  [[play](https://go.dev/play/p/hYa3cBEevtm)]
- **<big>FlattenDeep</big>** : 將切片遞迴攤平為一層。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FlattenDeep)]
  [[play](https://go.dev/play/p/yjYNHPyCFaF)]
- **<big>FlatMap</big>** : 處理切片，將其轉換並攤平為另一種型別的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FlatMap)]
  [[play](https://go.dev/play/p/_QARWlWs1N_F)]
- **<big>ForEach</big>** : 迭代切片的元素，並對每個元素呼叫函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ForEach)]
  [[play](https://go.dev/play/p/DrPaa4YsHRF)]
- **<big>ForEachWithBreak</big>** : 迭代切片的元素並對每個元素呼叫函式，當 iteratee 回傳 false 時，會中斷 for each 迴圈。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ForEachWithBreak)]
  [[play](https://go.dev/play/p/qScs39f3D9W)]
- **<big>ForEachConcurrent</big>** : 並行地對切片中的每個項目套用 iteratee 函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ForEachConcurrent)]
  [[play](https://go.dev/play/p/kT4XW7DKVoV)]
- **<big>GroupBy</big>** : 迭代切片的元素，依條件將每個元素分組，回傳兩個切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#GroupBy)]
  [[play](https://go.dev/play/p/QVkPxzPR0iA)]
- **<big>GroupWith</big>** : 回傳一個 map，其鍵由切片的每個元素經 iteratee 處理的結果產生。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#GroupWith)]
  [[play](https://go.dev/play/p/ApCvMNTLO8a)]
- **<big>IntSlice<sup>deprecated</sup></big>** : 將參數轉換為 int 切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IntSlice)]
  [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
- **<big>InterfaceSlice<sup>deprecated</sup></big>** : 將參數轉換為 interface 切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#InterfaceSlice)]
  [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
- **<big>Intersection</big>** : 建立由所有切片共同包含的不重複元素切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Intersection)]
  [[play](https://go.dev/play/p/anJXfB5wq_t)]
- **<big>InsertAt</big>** : 在指定索引處將值或其他切片插入切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#InsertAt)]
  [[play](https://go.dev/play/p/hMLNxPEGJVE)]
- **<big>IndexOf</big>** : 回傳項目在切片中第一次出現位置的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IndexOf)]
  [[play](https://go.dev/play/p/MRN1f0FpABb)]
- **<big>LastIndexOf</big>** : 回傳項目在切片中最後一次出現位置的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#LastIndexOf)]
  [[play](https://go.dev/play/p/DokM4cf1IKH)]
- **<big>Map</big>** : 透過讓切片的每個元素經過 iteratee 函式處理，建立值的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Map)]
  [[play](https://go.dev/play/p/biaTefqPquw)]
- **<big>MapConcurrent</big>** : 並行地對切片中的每個項目套用 iteratee 函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#MapConcurrent)]
  [[play](https://go.dev/play/p/H1ehfPkPen0)]
- **<big>Merge</big>** : 將所有給定的切片合併為一個切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Merge)]
  [[play](https://go.dev/play/p/lbjFp784r9N)]
- **<big>Reverse</big>** : 回傳元素順序與給定切片相反的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Reverse)]
  [[play](https://go.dev/play/p/8uI8f1lwNrQ)]
- **<big>Reduce<sup>deprecated</sup></big>** : 透過讓切片的每個元素經過 iteratee 函式處理，建立值的切片。（已棄用：請使用 ReduceBy）
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Reduce)]
  [[play](https://go.dev/play/p/_RfXJJWIsIm)]
- **<big>ReduceBy</big>** : 透過 reducer 函式累積切片中每個元素的處理結果，產生一個值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ReduceBy)]
  [[play](https://go.dev/play/p/YKDpLi7gtee)]
- **<big>ReduceRight</big>** : ReduceRight 與 ReduceBy 類似，但從右至左迭代切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ReduceRight)]
  [[play](https://go.dev/play/p/qT9dZC03A1K)]
- **<big>ReduceConcurrent</big>** : 並行地對切片中的每個項目套用 reducer 函式，將切片歸約為單一值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ReduceConcurrent)]
  [[play](https://go.dev/play/p/Tjwe6OtaG07)]
- **<big>Replace</big>** : 回傳切片的副本，其中前 n 個不重疊的 old 被取代為 new。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Replace)]
  [[play](https://go.dev/play/p/P5mZp7IhOFo)]
- **<big>ReplaceAll</big>** : 回傳切片的副本，其中所有不重疊的 old 被取代為 new。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ReplaceAll)]
  [[play](https://go.dev/play/p/CzqXMsuYUrx)]
- **<big>Repeat</big>** : 建立長度為 n、元素皆為傳入項目的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Repeat)]
  [[play](https://go.dev/play/p/1CbOmtgILUU)]
- **<big>Shuffle</big>** : 將切片隨機打亂。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Shuffle)]
  [[play](https://go.dev/play/p/YHvhnWGU3Ge)]
- **<big>ShuffleCopy</big>** : 回傳元素已隨機打亂的新切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ShuffleCopy)]
  [[play](https://go.dev/play/p/vqDa-Gs1vT0)]
- **<big>IsAscending</big>** : 檢查切片是否為遞增順序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsAscending)]
  [[play](https://go.dev/play/p/9CtsFjet4SH)]
- **<big>IsDescending</big>** : 檢查切片是否為遞減順序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsDescending)]
  [[play](https://go.dev/play/p/U_LljFXma14)]
- **<big>IsSorted</big>** : 檢查切片是否已排序（遞增或遞減）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsSorted)]
  [[play](https://go.dev/play/p/nCE8wPLwSA-)]
- **<big>IsSortedByKey</big>** : 檢查切片是否已依 iteratee 函式排序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsSortedByKey)]
  [[play](https://go.dev/play/p/tUoGB7DOHI4)]
- **<big>Sort</big>** : 對任何可排序型別（數字或字串）的切片排序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Sort)]
  [[play](https://go.dev/play/p/V9AVjzf_4Fk)]
- **<big>SortBy</big>** : 依 less 函式決定的順序將切片遞增排序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#SortBy)]
  [[play](https://go.dev/play/p/DAhLQSZEumm)]
- **<big>SortByField<sup>deprecated</sup></big>** : 回傳依指定欄位排序的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#SortByField)]
  [[play](https://go.dev/play/p/fU1prOBP9p1)]
- **<big>Some</big>** : 若清單中任一值通過判斷函式則回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Some)]
  [[play](https://go.dev/play/p/4pO9Xf9NDGS)]
- **<big>StringSlice<sup>deprecated</sup></big>** : 將參數轉換為字串切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#StringSlice)]
  [[play](https://go.dev/play/p/W0TZDWCPFcI)]
- **<big>SymmetricDifference</big>** : 兩個切片的對稱差集，也稱為互斥聯集。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#h42nJX5xMln)]
  [[play](https://go.dev/play/p/1CbOmtgILUU)]
- **<big>ToSlice</big>** : 回傳由可變參數轉換而成的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ToSlice)]
  [[play](https://go.dev/play/p/YzbzVq5kscN)]
- **<big>ToSlicePointer</big>** : 回傳指向由可變參數轉換而成之切片的指標。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ToSlicePointer)]
  [[play](https://go.dev/play/p/gx4tr6_VXSF)]
- **<big>Unique</big>** : 移除切片中重複的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Unique)]
  [[play](https://go.dev/play/p/AXw0R3ZTE6a)]
- **<big>UniqueBy</big>** : 依據 iteratee 函式回傳的值，移除輸入切片中重複的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UniqueBy)]
  [[play](https://go.dev/play/p/GY7JE4yikrl)]
- **<big>UniqueByComparator</big>** : 使用提供的比較函式，移除輸入切片中重複的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UniqueByComparator)]
  [[play](https://go.dev/play/p/rwSacr-ZHsR)]
- **<big>UniqueByField</big>** : 依結構欄位移除結構切片中重複的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UniqueByField)]
  [[play](https://go.dev/play/p/6cifcZSPIGu)]
- **<big>UniqueByConcurrent</big>** : 並行地移除切片中重複的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UniqueByConcurrent)]
  [[play](https://go.dev/play/p/wXZ7LcYRMGL)]
- **<big>Union</big>** : 依序建立由所有給定切片中的不重複元素組成的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Union)]
  [[play](https://go.dev/play/p/hfXV1iRIZOf)]
- **<big>UnionBy</big>** : 接收 iteratee，對每個切片的每個元素呼叫它，然後取聯集切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UnionBy)]
  [[play](https://go.dev/play/p/HGKHfxKQsFi)]
- **<big>UpdateAt</big>** : 更新切片中指定索引的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UpdateAt)]
  [[play](https://go.dev/play/p/f3mh2KloWVm)]
- **<big>Without</big>** : 建立排除所有給定項目的切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Without)]
  [[play](https://go.dev/play/p/bwhEXEypThg)]
- **<big>KeyBy</big>** : 依據回呼函式將切片轉換為 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#KeyBy)]
  [[play](https://go.dev/play/p/uXod2LWD1Kg)]
- **<big>Join</big>** : 以指定的分隔符號連接切片的項目。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Join)]
  [[play](https://go.dev/play/p/huKzqwNDD7V)]
- **<big>Partition</big>** : 依據給定判斷函式的評估結果，對切片的所有元素進行分區。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Partition)]
  [[play](https://go.dev/play/p/lkQ3Ri2NQhV)]
- **<big>Random</big>** : 取得切片的隨機項目並回傳其索引，當切片為空時回傳 -1。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Random)]
  [[play](https://go.dev/play/p/UzpGQptWppw)]
- **<big>SetToDefaultIf</big>** : 若元素符合給定的判斷函式，則將其設為預設值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#SetToDefaultIf)]
  [[play](https://go.dev/play/p/9AXGlPRC0-A)]
- **<big>Break</big>** : 在判斷函式第一次為 true 的位置，將清單拆成兩部分。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Break)]
- **<big>RightPadding</big>** : 在切片的右端加入填充。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#RightPadding)]
  [[play](https://go.dev/play/p/0_2rlLEMBXL)]
- **<big>LeftPadding</big>** : 在切片的左端加入填充。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#LeftPadding)]
  [[play](https://go.dev/play/p/jlQVoelLl2k)]
- **<big>Frequency</big>** : 計算切片中每個元素出現的頻率。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Frequency)]
  [[play](https://go.dev/play/p/CW3UVNdUZOq)]
- **<big>JoinFunc</big>** : 以給定的分隔符號將切片元素連接成單一字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#JoinFunc)]
  [[play](https://go.dev/play/p/55ib3SB5fM2)]
- **<big>ConcatBy</big>** : 使用提供的分隔符號與連接函式，將切片的元素串接成單一值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ConcatBy)]
  [[play](https://go.dev/play/p/6QcUpcY4UMW)]

<h3 id="stream"> 21. Stream 套件實作了支援循序與並行操作的元素序列。此套件是一項實驗，探索 Go 中的串流能否如同 Java 那樣運作，其功能非常有限。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/stream"
```

#### 函式清單:

- **<big>Of</big>** : 建立元素為指定值的串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Of)]
  [[play](https://go.dev/play/p/jI6_iZZuVFE)]
- **<big>FromSlice</big>** : 從切片建立串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FromSlice)]
  [[play](https://go.dev/play/p/wywTO0XZtI4)]
- **<big>FromChannel</big>** : 從通道建立串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FromChannel)]
  [[play](https://go.dev/play/p/9TZYugGMhXZ)]
- **<big>FromRange</big>** : 建立從 start 到 end 的數字串流，start 與 end 皆包含在內：[start, end]。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FromRange)]
  [[play](https://go.dev/play/p/9Ex1-zcg-B-)]
- **<big>Generate</big>** : 建立串流，其每個元素由提供的產生函式產生。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Generate)]
  [[play](https://go.dev/play/p/rkOWL1yA3j9)]
- **<big>Concat</big>** : 建立惰性串接的串流，其元素為第一個串流的所有元素，後接第二個串流的所有元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Concat)]
  [[play](https://go.dev/play/p/HM4OlYk_OUC)]
- **<big>Distinct</big>** : 建立並回傳移除重複項目的串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Distinct)]
  [[play](https://go.dev/play/p/eGkOSrm64cB)]
- **<big>Filter</big>** : 回傳由此串流中符合給定判斷函式的元素所組成的串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Filter)]
  [[play](https://go.dev/play/p/MFlSANo-buc)]
- **<big>FilterConcurrent</big>** : 並行地對輸入切片的每個元素套用提供的過濾函式 `predicate`。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FilterConcurrent)]
  [[play](https://go.dev/play/p/t_pkwerIRVx)]
- **<big>Map</big>** : 回傳由此串流的元素經給定函式處理後所組成的串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Map)]
  [[play](https://go.dev/play/p/OtNQUImdYko)]
- **<big>Peek</big>** : 回傳由此串流的元素所組成的串流，並在從結果串流消耗元素時，額外對每個元素執行提供的動作。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Peek)]
  [[play](https://go.dev/play/p/u1VNzHs6cb2)]
- **<big>Skip</big>** : 回傳在丟棄此串流前 n 個元素後，由剩餘元素所組成的串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Skip)]
  [[play](https://go.dev/play/p/fNdHbqjahum)]
- **<big>Limit</big>** : 回傳由此串流的元素所組成的串流，並截斷使其長度不超過 maxSize。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Limit)]
  [[play](https://go.dev/play/p/qsO4aniDcGf)]
- **<big>Reverse</big>** : 回傳元素順序與給定串流相反的串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Reverse)]
  [[play](https://go.dev/play/p/A8_zkJnLHm4)]
- **<big>ReverseCopy</big>** : 回傳元素順序與給定切片相反的新切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#ReverseCopy)]
  [[play](https://go.dev/play/p/c9arEaP7Cg-)]
- **<big>Range</big>** : 回傳元素範圍從原串流的 start（含）到 end（不含）的串流。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Range)]
  [[play](https://go.dev/play/p/indZY5V2f4j)]
- **<big>Sorted</big>** : 回傳由此串流的元素所組成的串流，並依提供的 less 函式排序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Sorted)]
  [[play](https://go.dev/play/p/XXtng5uonFj)]
- **<big>ForEach</big>** : 對此串流的每個元素執行一個動作。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#ForEach)]
  [[play](https://go.dev/play/p/Dsm0fPqcidk)]
- **<big>Reduce</big>** : 使用具結合律的累積函式對此串流的元素進行歸約，並回傳描述歸約值（若有）的 Optional。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Reduce)]
  [[play](https://go.dev/play/p/6uzZjq_DJLU)]
- **<big>FindFirst</big>** : 回傳此串流的第一個元素與 true；若串流為空，則回傳零值與 false。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FindFirst)]
  [[play](https://go.dev/play/p/9xEf0-6C1e3)]
- **<big>FindLast</big>** : 回傳此串流的最後一個元素與 true；若串流為空，則回傳零值與 false。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FindLast)]
  [[play](https://go.dev/play/p/WZD2rDAW-2h)]
- **<big>Max</big>** : 依提供的 less 函式回傳此串流的最大元素。less 函式：a > b
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Max)]
  [[play](https://go.dev/play/p/fm-1KOPtGzn)]
- **<big>Min</big>** : 依提供的 less 函式回傳此串流的最小元素。less 函式：a < b
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Min)]
  [[play](https://go.dev/play/p/vZfIDgGNRe_0)]
- **<big>AllMatch</big>** : 回傳此串流的所有元素是否皆符合提供的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#AllMatch)]
  [[play](https://go.dev/play/p/V5TBpVRs-Cx)]
- **<big>AnyMatch</big>** : 回傳此串流中是否有任何元素符合提供的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#AnyMatch)]
  [[play](https://go.dev/play/p/PTCnWn4OxSn)]
- **<big>NoneMatch</big>** : 回傳此串流中是否沒有任何元素符合提供的判斷函式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#NoneMatch)]
  [[play](https://go.dev/play/p/iWS64pL1oo3)]
- **<big>Count</big>** : 回傳串流中的元素數量。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Count)]
  [[play](https://go.dev/play/p/r3koY6y_Xo-)]
- **<big>ToSlice</big>** : 回傳串流中的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#ToSlice)]
  [[play](https://go.dev/play/p/jI6_iZZuVFE)]
- **<big>IndexOf</big>** : 回傳指定元素在此串流中第一次出現位置的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#IndexOf)]
  [[play](https://go.dev/play/p/tBV5Nc-XDX2)]
- **<big>LastIndexOf</big>** : 回傳指定元素在此串流中最後一次出現位置的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#LastIndexOf)]
  [[play](https://go.dev/play/p/CjeoNw2eac_G)]

<h3 id="structs"> 22. Structs 套件提供數個操作結構、標籤與欄位的高階函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/structs"
```

#### 函式清單:

- **<big>New</big>** : 建立 `Struct` 實例。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#New)]
  [[play](https://go.dev/play/p/O29l8kk-Z17)]
- **<big>ToMap</big>** : 將有效的結構轉換為 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#ToMap)]
  [[play](https://go.dev/play/p/qQbLySBgerZ)]
- **<big>Fields</big>** : 取得給定結構的所有欄位，這些欄位為抽象的結構欄位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#Fields)]
  [[play](https://go.dev/play/p/w3Kk_CyVY7D)]
- **<big>Field</big>** : 依給定的欄位名稱取得結構的抽象欄位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/api/packages/struct.md#Field)]
  [[play](https://go.dev/play/p/KocZMSYarza)]
- **<big>IsStruct</big>** : 檢查結構是否有效。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#IsStruct)]
  [[play](https://go.dev/play/p/bU2FSdkbK1C)]
- **<big>Tag</big>** : 取得 `Field` 的 `Tag`，`Tag` 為抽象的結構欄位標籤。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#Tag)]
  [[play](https://go.dev/play/p/DVrx5HvvUJr)]
- **<big>Name</big>** : 取得欄位名稱。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#Name)]
  [[play](https://go.dev/play/p/zfIGlqsatee)]
- **<big>Value</big>** : 取得 `Field` 的底層值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#Value)]
  [[play](https://go.dev/play/p/qufYEU2o4Oi)]
- **<big>Kind</big>** : 取得欄位的種類（kind）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#Kind)]
  [[play](https://go.dev/play/p/wg4NlcUNG5o)]
- **<big>IsEmbedded</big>** : 檢查欄位是否為嵌入欄位。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#IsEmbedded)]
  [[play](https://go.dev/play/p/wV2PrbYm3Ec)]
- **<big>IsExported</big>** : 檢查欄位是否已匯出。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#IsExported)]
  [[play](https://go.dev/play/p/csK4AXYaNbJ)]
- **<big>IsZero</big>** : 檢查欄位是否為零值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#IsZero)]
  [[play](https://go.dev/play/p/RzqpGISf87r)]
- **<big>IsSlice</big>** : 檢查欄位是否為切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#IsSlice)]
  [[play](https://go.dev/play/p/MKz4CgBIUrU)]
- **<big>IsTargetType</big>** : 檢查欄位是否為目標型別。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#IsTargetType)]
  [[play](https://go.dev/play/p/Ig75P-agN39)]
- **<big>TypeName</big>** : 回傳結構的型別名稱。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/struct.md#TypeName)]
  [[play](https://go.dev/play/p/SWLWd0XBaBb)]

<h3 id="strutil"> 23. Strutil 套件包含一些操作字串的函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### 函式清單:

- **<big>After</big>** : 回傳來源字串中，特定字串第一次出現之後的子字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#After)]
  [[play](https://go.dev/play/p/RbCOQqCDA7m)]
- **<big>AfterLast</big>** : 回傳來源字串中，特定字串最後一次出現之後的子字串。 [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#AfterLast)]
  [[play](https://go.dev/play/p/1TegARrb8Yn)]
- **<big>Before</big>** : 回傳來源字串中，特定字串第一次出現之前的子字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Before)]
  [[play](https://go.dev/play/p/JAWTZDS4F5w)]
- **<big>BeforeLast</big>** : 回傳來源字串中，特定字串最後一次出現之前的子字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#BeforeLast)]
  [[play](https://go.dev/play/p/pJfXXAoG_Te)]
- **<big>CamelCase</big>** : 將來源字串轉換為 camelCase 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#CamelCase)]
  [[play](https://go.dev/play/p/9eXP3tn2tUy)]
- **<big>UpperCamelCase</big>** : 將來源字串轉換為 UpperCamelCase 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#UpperCamelCase)]
  [[play](https://go.dev/play/p/vPHszjTXSFp)]
- **<big>Capitalize</big>** : 將來源字串的第一個字元轉為大寫，其餘轉為小寫。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Capitalize)]
  [[play](https://go.dev/play/p/2OAjgbmAqHZ)]
- **<big>ContainsAll</big>** : 若目標字串包含所有的子字串則回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#ContainsAll)]
  [[play](https://go.dev/play/p/KECtK2Os4zq)]
- **<big>ContainsAny</big>** : 若目標字串包含任一子字串則回傳 true。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#ContainsAny)]
  [[play](https://go.dev/play/p/dZGSSMB3LXE)]
- **<big>IsString</big>** : 檢查參數值的資料型別是否為字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IsString)]
  [[play](https://go.dev/play/p/IOgq7oF9ERm)]
- **<big>KebabCase</big>** : 將字串轉換為 kebab-case 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#KebabCase)]
  [[play](https://go.dev/play/p/dcZM9Oahw-Y)]
- **<big>UpperKebabCase</big>** : 將字串轉換為大寫的 KEBAB-CASE 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#UpperKebabCase)]
  [[play](https://go.dev/play/p/zDyKNneyQXk)]
- **<big>LowerFirst</big>** : 將字串的第一個字元轉為小寫。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#LowerFirst)]
  [[play](https://go.dev/play/p/CbzAyZmtJwL)]
- **<big>UpperFirst</big>** : 將字串的第一個字元轉為大寫。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#UpperFirst)]
  [[play](https://go.dev/play/p/sBbBxRbs8MM)]
- **<big>Pad</big>** : 若字串短於 size，則在其左右兩側填充。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Pad)]
  [[play](https://go.dev/play/p/NzImQq-VF8q)]
- **<big>PadEnd</big>** : 若字串短於限制長度，則在右側以給定字元填充，填充字元若超出長度則會被截斷。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#PadEnd)]
  [[play](https://go.dev/play/p/9xP8rN0vz--)]
- **<big>PadStart</big>** : 若字串短於限制長度，則在左側以給定字元填充，填充字元若超出長度則會被截斷。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#PadStart)]
  [[play](https://go.dev/play/p/xpTfzArDfvT)]
- **<big>Reverse</big>** : 回傳字元順序與給定字串相反的字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Reverse)]
  [[play](https://go.dev/play/p/adfwalJiecD)]
- **<big>SnakeCase</big>** : 將字串轉換為 snake_case 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SnakeCase)]
  [[play](https://go.dev/play/p/tgzQG11qBuN)]
- **<big>UpperSnakeCase</big>** : 將字串轉換為大寫的 SNAKE_CASE 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#UpperSnakeCase)]
  [[play](https://go.dev/play/p/4COPHpnLx38)]
- **<big>SplitEx</big>** : 分割給定的字串，可控制結果切片是否包含空字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SplitEx)]
  [[play](https://go.dev/play/p/Us-ySSbWh-3)]
- **<big>Substring</big>** : 回傳從指定偏移位置開始、具指定長度的子字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Substring)]
  [[play](https://go.dev/play/p/q3sM6ehnPDp)]
- **<big>Wrap</big>** : 以給定的字串包覆字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Wrap)]
  [[play](https://go.dev/play/p/KoZOlZDDt9y)]
- **<big>Unwrap</big>** : 從另一個字串中解除給定字串的包覆，會改變來源字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Unwrap)]
  [[play](https://go.dev/play/p/Ec2q4BzCpG-)]
- **<big>SplitWords</big>** : 將字串分割為單字，單字僅包含英文字母字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SplitWords)]
  [[play](https://go.dev/play/p/KLiX4WiysMM)]
- **<big>WordCount</big>** : 回傳字串中有意義單字的數量，單字僅包含英文字母字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#WordCount)]
  [[play](https://go.dev/play/p/bj7_odx3vRf)]
- **<big>RemoveNonPrintable</big>** : 移除字串中的不可列印字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#RemoveNonPrintable)]
  [[play](https://go.dev/play/p/og47F5x_jTZ)]
- **<big>StringToBytes</big>** : 不經記憶體配置，將字串轉換為位元組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#StringToBytes)]
  [[play](https://go.dev/play/p/7OyFBrf9AxA)]
- **<big>BytesToString</big>** : 不經記憶體配置，將位元組切片轉換為字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#BytesToString)]
  [[play](https://go.dev/play/p/6c68HRvJecH)]
- **<big>IsBlank</big>** : 檢查字串是否為空白或空字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IsBlank)]
  [[play](https://go.dev/play/p/6zXRH_c0Qd3)]
- **<big>IsNotBlank</big>** : 檢查字串是否不是空白且不是空字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IsNotBlank)]
  [[play](https://go.dev/play/p/e_oJW0RAquA)]
- **<big>HasPrefixAny</big>** : 檢查字串是否以指定字串陣列中的任一字串開頭。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#HasPrefixAny)]
  [[play](https://go.dev/play/p/8UUTl2C5slo)]
- **<big>HasSuffixAny</big>** : 檢查字串是否以指定字串陣列中的任一字串結尾。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#HasSuffixAny)]
  [[play](https://go.dev/play/p/sKWpCQdOVkx)]
- **<big>IndexOffset</big>** : 將字串偏移 `idxFrom` 之後，回傳 substr 第一次出現位置的索引。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IndexOffset)]
  [[play](https://go.dev/play/p/qZo4lV2fomB)]
- **<big>ReplaceWithMap</big>** : 回傳 `str` 的副本，其內容依 map 以無序、區分大小寫的方式取代。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#ReplaceWithMap)]
  [[play](https://go.dev/play/p/h3t7CNj2Vvu)]
- **<big>Trim</big>** : 去除字串開頭與結尾的空白（或其他字元）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Trim)]
  [[play](https://go.dev/play/p/Y0ilP0NRV3j)]
- **<big>SplitAndTrim</big>** : 以字串 `delimiter` 分割字串 `str` 為切片，並對切片的每個元素呼叫 Trim。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SplitAndTrim)]
  [[play](https://go.dev/play/p/ZNL6o4SkYQ7)]
- **<big>HideString</big>** : 以參數 `replaceChar` 隱藏來源字串中的部分字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#HideString)]
  [[play](https://go.dev/play/p/pzbaIVCTreZ)]
- **<big>RemoveWhiteSpace</big>** : 移除字串中的空白字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#RemoveWhiteSpace)]
  [[play](https://go.dev/play/p/HzLC9vsTwkf)]
- **<big>SubInBetween</big>** : 回傳來源字串中從起始位置到結束位置（不含）之間的子字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SubInBetween)]
  [[play](https://go.dev/play/p/EDbaRvjeNsv)]
- **<big>HammingDistance</big>** : 計算兩個字串之間的漢明距離。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#HammingDistance)]
  [[play](https://go.dev/play/p/glNdQEA9HUi)]
- **<big>Concat</big>** : 串接字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Concat)]
  [[play](https://go.dev/play/p/gD52SZHr4Kp)]
- **<big>Ellipsis</big>** : 將字串截斷至指定長度並附加省略符號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Ellipsis)]
  [[play](https://go.dev/play/p/i1vbdQiQVRR)]
- **<big>Shuffle</big>** : 打亂給定字串中字元的順序。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Shuffle)]
  [[play](https://go.dev/play/p/iStFwBwyGY7)]
- **<big>Rotate</big>** : 將字串依指定的字元數旋轉。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Rotate)]
  [[play](https://go.dev/play/p/Kf03iOeT5bd)]
- **<big>TemplateReplace</big>** : 以資料 map 中對應的值取代樣板字串中的佔位符。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#TemplateReplace)]
  [[play](https://go.dev/play/p/cXSuFvyZqv9)]
- **<big>RegexMatchAllGroups</big>** : 使用正規表示式比對字串中所有的子群組並回傳結果。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#RegexMatchAllGroups)]
  [[play](https://go.dev/play/p/JZiu0RXpgN-)]
- **<big>ExtractContent</big>** : 擷取來源字串中起始字串與結束字串之間的內容。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#ExtractContent)]
  [[play](https://go.dev/play/p/Ay9UIk7Rum9)]
- **<big>FindAllOccurrences</big>** : 回傳子字串在字串中所有出現位置。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#FindAllOccurrences)]
  [[play](https://go.dev/play/p/uvyA6azGLB1)]

<h3 id="system"> 24. System 套件包含一些與作業系統、執行環境、shell 指令相關的函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/system"
```

#### 函式清單:

- **<big>IsWindows</big>** : 檢查目前作業系統是否為 Windows。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#IsWindows)]
  [[play](https://go.dev/play/p/XzJULbzmf9m)]
- **<big>IsLinux</big>** : 檢查目前作業系統是否為 Linux。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#IsLinux)]
  [[play](https://go.dev/play/p/zIflQgZNuxD)]
- **<big>IsMac</big>** : 檢查目前作業系統是否為 macOS。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#IsMac)]
  [[play](https://go.dev/play/p/Mg4Hjtyq7Zc)]
- **<big>GetOsEnv</big>** : 取得以 key 命名的環境變數的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#GetOsEnv)]
  [[play](https://go.dev/play/p/D88OYVCyjO-)]
- **<big>SetOsEnv</big>** : 設定以 key 命名的環境變數的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#SetOsEnv)]
  [[play](https://go.dev/play/p/D88OYVCyjO-)]
- **<big>RemoveOsEnv</big>** : 移除單一環境變數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#RemoveOsEnv)]
  [[play](https://go.dev/play/p/fqyq4b3xUFQ)]
- **<big>CompareOsEnv</big>** : 取得以 key 命名的環境變數，並與傳入的環境值比對。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#CompareOsEnv)]
  [[play](https://go.dev/play/p/BciHrKYOHbp)]
- **<big>ExecCommand</big>** : 執行指令，回傳指令的標準輸出與標準錯誤字串，發生錯誤時一併回傳錯誤。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#ExecCommand)]
  [[play](https://go.dev/play/p/n-2fLyZef-4)]
- **<big>GetOsBits</big>** : 回傳目前作業系統的位元數（32 或 64）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#GetOsBits)]
  [[play](https://go.dev/play/p/ml-_XH3gJbW)]
- **<big>StartProcess</big>** : 以指定的名稱與參數啟動新的行程。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#StartProcess)]
  [[play](https://go.dev/play/p/5GVol6ryS_X)]
- **<big>StopProcess</big>** : 依 pid 停止行程。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#StopProcess)]
  [[play](https://go.dev/play/p/jJZhRYGGcmD)]
- **<big>KillProcess</big>** : 以指定的名稱與參數終止行程。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#KillProcess)]
  [[play](https://go.dev/play/p/XKmvV-ExBWa)]
- **<big>GetProcessInfo</big>** : 依 pid 取得行程的詳細資訊。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#GetProcessInfo)]
  [[play](https://go.dev/play/p/NQDVywEYYx7)]

<h3 id="tuple"> 25. Tuple 套件實作了 tuple 資料型別及其部分操作。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/tuple"
```

#### 函式清單:

- **<big>Tuple2</big>** : 代表含 2 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple2)]
  [[play](https://go.dev/play/p/3sHVqBQpLYN)]
- **<big>Tuple2_Unbox</big>** : 回傳 Tuple2 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple2_Unbox)]
  [[play](https://go.dev/play/p/0fD1qfCVwjm)]
- **<big>Zip2</big>** : 建立 Tuple2 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip2)]
  [[play](https://go.dev/play/p/4ncWJJ77Xio)]
- **<big>Unzip2</big>** : 從 Tuple2 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip2)]
  [[play](https://go.dev/play/p/KBecr60feXb)]
- **<big>Tuple3</big>** : 代表含 3 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple3)]
  [[play](https://go.dev/play/p/FtH2sdCLlCf)]
- **<big>Tuple3_Unbox</big>** : 回傳 Tuple3 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple3_Unbox)]
  [[play](https://go.dev/play/p/YojLy-id1BS)]
- **<big>Zip3</big>** : 建立 Tuple3 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip3)]
  [[play](https://go.dev/play/p/97NgmsTILfu)]
- **<big>Unzip3</big>** : 從 Tuple3 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip3)]
  [[play](https://go.dev/play/p/bba4cpAa7KO)]
- **<big>Tuple4</big>** : 代表含 4 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple4)]
  [[play](https://go.dev/play/p/D2EqDz096tk)]
- **<big>Tuple4_Unbox</big>** : 回傳 Tuple4 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple4_Unbox)]
  [[play](https://go.dev/play/p/ACj9YuACGgW)]
- **<big>Zip4</big>** : 建立 Tuple4 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip4)]
  [[play](https://go.dev/play/p/PEmTYVK5hL4)]
- **<big>Unzip4</big>** : 從 Tuple4 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip4)]
  [[play](https://go.dev/play/p/rb8z4gyYSRN)]
- **<big>Tuple5</big>** : 代表含 5 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple5)]
  [[play](https://go.dev/play/p/2WndmVxPg-r)]
- **<big>Tuple5_Unbox</big>** : 回傳 Tuple4 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple5_Unbox)]
  [[play](https://go.dev/play/p/GyIyZHjCvoS)]
- **<big>Zip5</big>** : 建立 Tuple5 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip5)]
  [[play](https://go.dev/play/p/fCAAJLMfBIP)]
- **<big>Unzip5</big>** : 從 Tuple5 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip5)]
  [[play](https://go.dev/play/p/gyl6vKfhqPb)]
- **<big>Tuple6</big>** : 代表含 6 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple6)]
  [[play](https://go.dev/play/p/VjqcCwEJZbs)]
- **<big>Tuple6_Unbox</big>** : 回傳 Tuple6 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple6_Unbox)]
  [[play](https://go.dev/play/p/FjIHV7lpxmW)]
- **<big>Zip6</big>** : 建立 Tuple6 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip6)]
  [[play](https://go.dev/play/p/oWPrnUYuFHo)]
- **<big>Unzip6</big>** : 從 Tuple6 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip6)]
  [[play](https://go.dev/play/p/l41XFqCyh5E)]
- **<big>Tuple7</big>** : 代表含 7 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple7)]
  [[play](https://go.dev/play/p/dzAgv_Ezub9)]
- **<big>Tuple7_Unbox</big>** : 回傳 Tuple7 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple7_Unbox)]
  [[play](https://go.dev/play/p/R9I8qeDk0zs)]
- **<big>Zip7</big>** : 建立 Tuple7 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip7)]
  [[play](https://go.dev/play/p/WUJuo897Egf)]
- **<big>Unzip7</big>** : 從 Tuple7 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip7)]
  [[play](https://go.dev/play/p/hws_P1Fr2j3)]
- **<big>Tuple8</big>** : 代表含 8 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple8)]
  [[play](https://go.dev/play/p/YA9S0rz3dRz)]
- **<big>Tuple8_Unbox</big>** : 回傳 Tuple8 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple8_Unbox)]
  [[play](https://go.dev/play/p/PRxLBBb4SMl)]
- **<big>Zip8</big>** : 建立 Tuple8 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip8)]
  [[play](https://go.dev/play/p/8V9jWkuJfaQ)]
- **<big>Unzip8</big>** : 從 Tuple8 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip8)]
  [[play](https://go.dev/play/p/1SndOwGsZB4)]
- **<big>Tuple9</big>** : 代表含 9 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple9)]
  [[play](https://go.dev/play/p/yS2NGGtZpQr)]
- **<big>Tuple9_Unbox</big>** : 回傳 Tuple9 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple9_Unbox)]
  [[play](https://go.dev/play/p/oFJFGTAuOa8)]
- **<big>Zip9</big>** : 建立 Tuple9 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip9)]
  [[play](https://go.dev/play/p/cgsL15QYnfz)]
- **<big>Unzip9</big>** : 從 Tuple9 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip9)]
  [[play](https://go.dev/play/p/91-BU_KURSA)]
- **<big>Tuple10</big>** : 代表含 10 個元素的 tuple。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple10)]
  [[play](https://go.dev/play/p/799qqZg0hUv)]
- **<big>Tuple10_Unbox</big>** : 回傳 Tuple10 中的值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple10_Unbox)]
  [[play](https://go.dev/play/p/qfyx3x_X0Cu)]
- **<big>Zip10</big>** : 建立 Tuple10 的切片，其元素對應於給定切片的元素。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip10)]
  [[play](https://go.dev/play/p/YSR-2cXnrY4)]
- **<big>Unzip10</big>** : 從 Tuple10 的切片建立一組切片。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip10)]
  [[play](https://go.dev/play/p/-taQB6Wfre_z)]

<h3 id="validator"> 26. Validator 套件包含一些資料驗證函式。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/validator"
```

#### 函式清單:

- **<big>ContainChinese</big>** : 檢查字串是否包含中文字（漢字）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainChinese)]
  [[play](https://go.dev/play/p/7DpU0uElYeM)]
- **<big>ContainLetter</big>** : 檢查字串是否至少包含一個字母。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainLetter)]
  [[play](https://go.dev/play/p/lqFD04Yyewp)]
- **<big>ContainLower</big>** : 檢查字串是否至少包含一個小寫字母 a-z。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainLower)]
  [[play](https://go.dev/play/p/Srqi1ItvnAA)]
- **<big>ContainUpper</big>** : 檢查字串是否至少包含一個大寫字母 A-Z。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainUpper)]
  [[play](https://go.dev/play/p/CmWeBEk27-z)]
- **<big>IsAlpha</big>** : 檢查字串是否僅包含字母（a-zA-Z）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAlpha)]
  [[play](https://go.dev/play/p/7Q5sGOz2izQ)]
- **<big>IsAllUpper</big>** : 檢查字串是否全為大寫字母 A-Z。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAllUpper)]
  [[play](https://go.dev/play/p/ZHctgeK1n4Z)]
- **<big>IsAllLower</big>** : 檢查字串是否全為小寫字母 a-z。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAllLower)]
  [[play](https://go.dev/play/p/GjqCnOfV6cM)]
- **<big>IsBase64</big>** : 檢查字串是否為 base64 字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsBase64)]
  [[play](https://go.dev/play/p/sWHEySAt6hl)]
- **<big>IsChineseMobile</big>** : 檢查字串是否為中國大陸手機號碼。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChineseMobile)]
  [[play](https://go.dev/play/p/GPYUlGTOqe3)]
- **<big>IsChineseIdNum</big>** : 檢查字串是否為中國身分證號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChineseIdNum)]
  [[play](https://go.dev/play/p/d8EWhl2UGDF)]
- **<big>IsChinesePhone</big>** : 檢查字串是否為中國電話號碼。(xxx-xxxxxxxx 或 xxxx-xxxxxxx。)
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChinesePhone)]
  [[play](https://go.dev/play/p/RUD_-7YZJ3I)]
- **<big>IsCreditCard</big>** : 檢查字串是否為信用卡號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsCreditCard)]
  [[play](https://go.dev/play/p/sNwwL6B0-v4)]
- **<big>IsDns</big>** : 檢查字串是否為 DNS。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsDns)]
  [[play](https://go.dev/play/p/jlYApVLLGTZ)]
- **<big>IsEmail</big>** : 檢查字串是否為電子郵件地址。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsEmail)]
  [[play](https://go.dev/play/p/HVQ5LAe-vFz)]
- **<big>IsEmptyString</big>** : 檢查字串是否為空。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsEmptyString)]
  [[play](https://go.dev/play/p/dpzgUjFnBCX)]
- **<big>IsFloat</big>** : 檢查值是否為浮點數（float32、float34）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsFloat)]
  [[play](https://go.dev/play/p/vsyG-sxr99_Z)]
- **<big>IsFloatStr</big>** : 檢查字串是否可轉換為浮點數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsFloatStr)]
  [[play](https://go.dev/play/p/LOYwS_Oyl7U)]
- **<big>IsNumber</big>** : 檢查值是否為數字（整數、浮點數）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsNumber)]
  [[play](https://go.dev/play/p/mdJHOAvtsvF)]
- **<big>IsNumberStr</big>** : 檢查字串是否可轉換為數字。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsNumberStr)]
  [[play](https://go.dev/play/p/LzaKocSV79u)]
- **<big>IsAlphaNumeric</big>** : 檢查字串是否為英數字。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAlphaNumeric)]
  [[play](https://go.dev/play/p/RHeESLrLg9c)]
- **<big>IsJSON</big>** : 檢查字串是否為有效的 JSON。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsJSON)]
  [[play](https://go.dev/play/p/8Kip1Itjiil)]
- **<big>IsRegexMatch</big>** : 檢查字串是否符合正規表示式。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsRegexMatch)]
  [[play](https://go.dev/play/p/z_XeZo_litG)]
- **<big>IsInt</big>** : 檢查字串是否可轉換為數字。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsInt)]
  [[play](https://go.dev/play/p/eFoIHbgzl-z)]
- **<big>IsIntStr</big>** : 檢查字串是否可轉換為整數。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIntStr)]
  [[play](https://go.dev/play/p/jQRtFv-a0Rk)]
- **<big>IsIp</big>** : 檢查字串是否為 IP。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIp)]
  [[play](https://go.dev/play/p/FgcplDvmxoD)]
- **<big>IsIpV4</big>** : 檢查字串是否為 IPv4。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIpV4)]
  [[play](https://go.dev/play/p/zBGT99EjaIu)]
- **<big>IsIpV6</big>** : 檢查字串是否為 IPv6。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIpV6)]
  [[play](https://go.dev/play/p/AHA0r0AzIdC)]
- **<big>IsIpPort</big>** : 檢查字串是否為 ip:port。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIpPort)]
  [[play](https://go.dev/play/p/xUmls_b9L29)]
- **<big>IsStrongPassword</big>** : 檢查字串是否為強密碼。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsStrongPassword)]
  [[play](https://go.dev/play/p/QHdVcSQ3uDg)]
- **<big>IsUrl</big>** : 檢查字串是否為 URL。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsUrl)]
  [[play](https://go.dev/play/p/pbJGa7F98Ka)]
- **<big>IsWeakPassword</big>** : 檢查字串是否為弱密碼。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsWeakPassword)]
  [[play](https://go.dev/play/p/wqakscZH5gH)]
- **<big>IsZeroValue</big>** : 檢查值是否為零值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsZeroValue)]
  [[play](https://go.dev/play/p/UMrwaDCi_t4)]
- **<big>IsGBK</big>** : 檢查資料編碼是否為 GBK。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsGBK)]
  [[play](https://go.dev/play/p/E2nt3unlmzP)]
- **<big>IsASCII</big>** : 檢查字串是否全為 ASCII 字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsASCII)]
  [[play](https://go.dev/play/p/hfQNPLX0jNa)]
- **<big>IsPrintable</big>** : 檢查字串是否全為可列印字元。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsPrintable)]
  [[play](https://go.dev/play/p/Pe1FE2gdtTP)]
- **<big>IsBin</big>** : 檢查給定的字串是否為有效的二進位值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsBin)]
  [[play](https://go.dev/play/p/ogPeg2XJH4P)]
- **<big>IsHex</big>** : 檢查給定的字串是否為有效的十六進位值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsHex)]
  [[play](https://go.dev/play/p/M2qpHbEwmm7)]
- **<big>IsBase64URL</big>** : 檢查給定的字串是否為有效的 URL 安全 Base64 編碼字串。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsBase64URL)]
  [[play](https://go.dev/play/p/vhl4mr8GZ6S)]
- **<big>IsJWT</big>** : 檢查給定的字串是否為有效的 JSON Web Token (JWT)。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsJWT)]
  [[play](https://go.dev/play/p/R6Op7heJbKI)]
- **<big>IsVisa</big>** : 檢查給定的字串是否為有效的 Visa 卡號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsVisa)]
  [[play](https://go.dev/play/p/SdS2keOyJsl)]
- **<big>IsMasterCard</big>** : 檢查給定的字串是否為有效的 MasterCard 卡號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsMasterCard)]
  [[play](https://go.dev/play/p/CwWBFRrG27b)]
- **<big>IsAmericanExpress</big>** : 檢查給定的字串是否為有效的美國運通卡號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAmericanExpress)]
  [[play](https://go.dev/play/p/HIDFpcOdpkd)]
- **<big>IsUnionPay</big>** : 檢查給定的字串是否為有效的銀聯卡號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsUnionPay)]
  [[play](https://go.dev/play/p/CUHPEwEITDf)]
- **<big>IsChinaUnionPay</big>** : 檢查給定的字串是否為有效的中國銀聯卡號。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChinaUnionPay)]
  [[play](https://go.dev/play/p/yafpdxLiymu)]
- **<big>IsPassport</big>** : 護照驗證（使用正規表示式）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsPassport)]
  [[play](https://go.dev/play/p/dvOiV2BW7Aw)]
- **<big>IsChineseHMPassport</big>** : 港澳居民來往內地通行證驗證（使用正規表示式）。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChineseHMPassport)]
  [[play](https://go.dev/play/p/xKG6spQTcY0)]

<h3 id="xerror"> 27. Xerror 套件實作了錯誤處理的輔助工具。 &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/xerror"
```

#### 函式清單:

- **<big>New</big>** : 以訊息建立新的 XError 指標實例。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#New)]
  [[play](https://go.dev/play/p/w4oWZts7q7f)]
- **<big>Wrap</big>** : 依據 error 物件建立新的 XError 指標實例，並加入訊息。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#Wrap)]
  [[play](https://go.dev/play/p/5385qT2dCi4)]
- **<big>Unwrap</big>** : 透過 errors.As 從 err 取得解開包裝的 XError，若沒有 XError 則回傳 nil。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#Unwrap)]
  [[play](https://go.dev/play/p/LKMLep723tu)]
- **<big>XError_Wrap</big>** : 建立新的 XError，並將訊息與 id 複製到新的 XError。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Wrap)]
  [[play](https://go.dev/play/p/RpjJ5u5sc97)]
- **<big>XError_Unwrap</big>** : 與 github.com/pkg/errors 相容。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Unwrap)]
  [[play](https://go.dev/play/p/VUXJ8BST4c6)]
- **<big>XError_With</big>** : 加入與 XError 物件相關的鍵與值。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_With)]
  [[play](https://go.dev/play/p/ow8UISXX_Dp)]
- **<big>XError_Id</big>** : 設定 XError 物件的 id，以便在 XError.Is 中檢查是否相等。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Id)]
  [[play](https://go.dev/play/p/X6HBlsy58U9)]
- **<big>XError_Is</big>** : 檢查目標錯誤是否為 XError，且兩個錯誤的 Error.id 是否相符。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Is)]
  [[play](https://go.dev/play/p/X6HBlsy58U9)]
- **<big>XError_Values</big>** : 回傳由 XError.With 函式所設定的鍵與值的 map。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Values)]
  [[play](https://go.dev/play/p/ow8UISXX_Dp)]
- **<big>XError_StackTrace</big>** : 回傳與 pkg/errors 相容的堆疊追蹤。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_StackTrace)]
  [[play](https://go.dev/play/p/6FAvSQpa7pc)]
- **<big>XError_Info</big>** : 回傳 xerror 的資訊，可用於列印。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Info)]
  [[play](https://go.dev/play/p/1ZX0ME1F-Jb)]
- **<big>XError_Error</big>** : 實作標準的 error 介面。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Error)]
  [[play](https://go.dev/play/p/w4oWZts7q7f)]
- **<big>TryUnwrap</big>** : 檢查 err 是否為 nil，若是則回傳有效的值；若 err 不為 nil，TryUnwrap 會以 err 觸發 panic。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#TryUnwrap)]
  [[play](https://go.dev/play/p/acyZVkNZEeW)]
- **<big>TryCatch</big>** : 簡易模擬 Java 風格的 try-catch。
  [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#TryCatch)]
  [[play](https://go.dev/play/p/D5Mdb0mRj0P)]

## 如何貢獻

#### [貢獻指南](./CONTRIBUTION.md)

## 貢獻者

感謝所有為 lancet 做出貢獻的人！

<a href="https://github.com/duke-git/lancet/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=duke-git/lancet" />
</a>

## 星標歷史

[![Star History Chart](https://api.star-history.com/svg?repos=duke-git/lancet&type=Date)](https://star-history.com/#duke-git/lancet&Date)
