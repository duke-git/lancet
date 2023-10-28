<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-2.2.7-green.svg)](https://github.com/duke-git/lancet/releases)
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

## <a href="https://www.golancet.cn/en/" target="_blank"> Website</a> | [简体中文](./README_zh-CN.md)

## Features

-   👏 Comprehensive, efficient and reusable.
-   💪 600+ go util functions, support string, slice, datetime, net, crypt...
-   💅 Only depends on two kinds of libraries: go standard library and golang.org/x.
-   🌍 Unit test for every exported function.

## Installation

### Note:

1. <b>For users who use go1.18 and above, it is recommended to install lancet v2.x.x. Cause in v2.x.x all functions was rewriten with generics of go1.18.</b>

```go
go get github.com/duke-git/lancet/v2 // will install latest version of v2.x.x
```

2. <b>For users who use version below go1.18, you should install v1.x.x. The latest of v1.x.x is v1.4.2. </b>

```go
go get github.com/duke-git/lancet // below go1.18, install latest version of v1.x.x
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

## Documentation

### <span id="index">Index<span>

-   [Algorithm](#user-content-algorithm)
-   [Compare](#user-content-compare)
-   [Concurrency](#user-content-concurrency)
-   [Condition](#user-content-condition)
-   [Convertor](#user-content-convertor)
-   [Cryptor](#user-content-cryptor)
-   [Datetime](#user-content-datetime)
-   [Datastructure](#user-content-datastructure)
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

<h3 id="algorithm">1. Algorithm package implements some basic algorithm. eg. sort, search. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/algorithm"
```

#### Function list:

-   **<big>BubbleSort</big>** : sorts slice with bubble sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#BubbleSort)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>CountSort</big>** : sorts slice with bubble sort algorithm, don't change original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#CountSort)]
    [[play](https://go.dev/play/p/tB-Umgm0DrP)]
-   **<big>HeapSort</big>** : sorts slice with heap sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#HeapSort)]
    [[play](https://go.dev/play/p/u6Iwa1VZS_f)]
-   **<big>InsertionSort</big>** : sorts slice with insertion sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#InsertionSort)]
    [[play](https://go.dev/play/p/G5LJiWgJJW6)]
-   **<big>MergeSort</big>** : sorts slice with merge sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#MergeSort)]
    [[play](https://go.dev/play/p/ydinn9YzUJn)]
-   **<big>QuickSort</big>** : sorts slice with quick sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#QuickSort)]
    [[play](https://go.dev/play/p/7Y7c1Elk3ax)]
-   **<big>SelectionSort</big>** : sorts slice with selection sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#SelectionSort)]
    [[play](https://go.dev/play/p/oXovbkekayS)]
-   **<big>ShellSort</big>** : sorts slice with shell sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#ShellSort)]
    [[play](https://go.dev/play/p/3ibkszpJEu3)]
-   **<big>BinarySearch</big>** : returns the index of target within a sorted slice, use binary search (recursive call itself).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#BinarySearch)]
    [[play](https://go.dev/play/p/t6MeGiUSN47)]
-   **<big>BinaryIterativeSearch</big>** : returns the index of target within a sorted slice, use binary search (no recursive).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#BinaryIterativeSearch)]
    [[play](https://go.dev/play/p/Anozfr8ZLH3)]
-   **<big>LinearSearch</big>** : returns the index of target in slice base on equal function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#LinearSearch)]
    [[play](https://go.dev/play/p/IsS7rgn5s3x)]
-   **<big>LRUCache</big>** : implements memory cache with lru algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/algorithm.md#LRUCache)]
    [[play](https://go.dev/play/p/-EZjgOURufP)]

<h3 id="compare"> 2. Compare package provides a lightweight comparison function on any type. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/compare"
```

#### Function list:

-   **<big>Equal</big>** : Checks if two values are equal or not. (check both type and value)
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#Equal)]
    [[play](https://go.dev/play/p/wmVxR-to4lz)]
-   **<big>EqualValue</big>** : Checks if two values are equal or not. (check value only)
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#EqualValue)]
    [[play](https://go.dev/play/p/fxnna_LLD9u)]
-   **<big>LessThan</big>** : Checks if value `left` less than value `right`.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#LessThan)]
    [[play](https://go.dev/play/p/cYh7FQQj0ne)]
-   **<big>GreaterThan</big>** : Checks if value `left` greater than value `right`.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#GreaterThan)]
    [[play](https://go.dev/play/p/9-NYDFZmIMp)]
-   **<big>LessOrEqual</big>** : Checks if value `left` less than or equal than value `right`.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#LessOrEqual)]
    [[play](https://go.dev/play/p/e4T_scwoQzp)]
-   **<big>GreaterOrEqual</big>** : Checks if value `left` less greater or equal than value `right`.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#GreaterOrEqual)]
    [[play](https://go.dev/play/p/vx8mP0U8DFk)]
-   **<big>InDelta</big>** : Checks if two values are equal or not within a delta.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/compare.md#InDelta)]

<h3 id="concurrency"> 3. Concurrency package contain some functions to support concurrent programming. eg, goroutine, channel, async. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/concurrency"
```

#### Function list:

-   **<big>NewChannel</big>** : create a Channel pointer instance.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#NewChannel)]
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Bridge</big>** : link multiply channels into one channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/Bridge.md#NewChannel)]
    [[play](https://go.dev/play/p/qmWSy1NVF-Y)]
-   **<big>FanIn</big>** : merge multiple channels into one channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#FanIn)]
    [[play](https://go.dev/play/p/2VYFMexEvTm)]
-   **<big>Generate</big>** : creates a channel, then put values into the channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Generate)]
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Or</big>** : read one or more channels into one channel, will close when any readin channel is closed.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Or)]
    [[play](https://go.dev/play/p/Wqz9rwioPww)]
-   **<big>OrDone</big>** : read a channel into another channel, will close until cancel context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#OrDone)]
    [[play](https://go.dev/play/p/lm_GoS6aDjo)]
-   **<big>Repeat</big>** : create channel, put values into the channel repeatly until cancel the context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Repeat)]
    [[play](https://go.dev/play/p/k5N_ALVmYjE)]
-   **<big>RepeatFn</big>** : create a channel, excutes fn repeatly, and put the result into the channel, until close context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#RepeatFn)]
    [[play](https://go.dev/play/p/4J1zAWttP85)]
-   **<big>Take</big>** : create a channel whose values are taken from another channel with limit number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Take)]
    [[play](https://go.dev/play/p/9Utt-1pDr2J)]
-   **<big>Tee</big>** : split one chanel into two channels, until cancel the context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/concurrency.md#Tee)]
    [[play](https://go.dev/play/p/3TQPKnCirrP)]

<h3 id="condition"> 4. Condition package contains some functions for conditional judgment. eg. And, Or, TernaryOperator...&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/condition"
```

#### Function list:

-   **<big>Bool</big>** : returns the truthy value of anything.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Bool)]
    [[play](https://go.dev/play/p/ETzeDJRSvhm)]
-   **<big>And</big>** : returns true if both a and b are truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#And)]
    [[play](https://go.dev/play/p/W1SSUmt6pvr)]
-   **<big>Or</big>** : returns false if neither a nor b is truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Or)]
    [[play](https://go.dev/play/p/UlQTxHaeEkq)]
-   **<big>Xor</big>** : returns true if a or b but not both is truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Xor)]
    [[play](https://go.dev/play/p/gObZrW7ZbG8)]
-   **<big>Nor</big>** : returns true if neither a nor b is truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Nor)]
    [[play](https://go.dev/play/p/g2j08F_zZky)]
-   **<big>Xnor</big>** : returns true if both a and b or neither a nor b are truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Xnor)]
    [[play](https://go.dev/play/p/OuDB9g51643)]
-   **<big>Nand</big>** : returns false if both a and b are truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#Nand)]
    [[play](https://go.dev/play/p/vSRMLxLIbq8)]
-   **<big>TernaryOperator</big>** : ternary operator.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/condition.md#TernaryOperator)]
    [[play](https://go.dev/play/p/ElllPZY0guT)]

<h3 id="convertor"> 5. Convertor package contains some functions for data convertion. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a> </h3>

```go
import "github.com/duke-git/lancet/v2/convertor"
```

#### Function list:

-   **<big>ColorHexToRGB</big>** : convert color hex to color rgb.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ColorHexToRGB)]
    [[play](https://go.dev/play/p/o7_ft-JCJBV)]
-   **<big>ColorRGBToHex</big>** : convert rgb color to hex color.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ColorRGBToHex)]
    [[play](https://go.dev/play/p/nzKS2Ro87J1)]
-   **<big>ToBool</big>** : convert string to bool.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToBool)]
    [[play](https://go.dev/play/p/ARht2WnGdIN)]
-   **<big>ToBytes</big>** : convert value to byte slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToBytes)]
    [[play](https://go.dev/play/p/fAMXYFDvOvr)]
-   **<big>ToChar</big>** : convert string to char slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToChar)]
    [[play](https://go.dev/play/p/JJ1SvbFkVdM)]
-   **<big>ToChannel</big>** : convert a collection of elements to a read-only channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToChannel)]
    [[play](https://go.dev/play/p/hOx_oYZbAnL)]
-   **<big>ToFloat</big>** : convert value to float64, if param is a invalid floatable, will return 0.0 and error.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToFloat)]
    [[play](https://go.dev/play/p/4YTmPCibqHJ)]
-   **<big>ToInt</big>** : convert value to int64 value, if input is not numerical, return 0 and error.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToInt)]
    [[play](https://go.dev/play/p/9_h9vIt-QZ_b)]
-   **<big>ToJson</big>** : convert value to a json string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToJson)]
    [[play](https://go.dev/play/p/2rLIkMmXWvR)]
-   **<big>ToMap</big>** : convert a slice of structs to a map based on iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToMap)]
    [[play](https://go.dev/play/p/tVFy7E-t24l)]
-   **<big>ToPointer</big>** : return a pointer of passed value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToPointer)]
    [[play](https://go.dev/play/p/ASf_etHNlw1)]
-   **<big>ToString</big>** : convert value to string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToString)]
    [[play](https://go.dev/play/p/nF1zOOslpQq)]
-   **<big>StructToMap</big>** : convert struct to map, only convert exported struct field.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#StructToMap)]
    [[play](https://go.dev/play/p/KYGYJqNUBOI)]
-   **<big>MapToSlice</big>** : convert map to slice based on iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#MapToSlice)]
    [[play](https://go.dev/play/p/dmX4Ix5V6Wl)]
-   **<big>EncodeByte</big>** : encode data to byte slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#EncodeByte)]
    [[play](https://go.dev/play/p/DVmM1G5JfuP)]
-   **<big>DecodeByte</big>** : decode byte slice data to target object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#DecodeByte)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>DeepClone</big>** : creates a deep copy of passed item, can't clone unexported field of struct.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#DeepClone)]
    [[play](https://go.dev/play/p/j4DP5dquxnk)]
-   **<big>CopyProperties</big>** : copies each field from the source struct into the destination struct.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#CopyProperties)]
    [[play](https://go.dev/play/p/oZujoB5Sgg5)]
-   **<big>ToInterface</big>** : converts reflect value to its interface type.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#ToInterface)]
    [[play](https://go.dev/play/p/syqw0-WG7Xd)]
-   **<big>Utf8ToGbk</big>** : converts utf8 encoding data to GBK encoding data
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#Utf8ToGbk)]
    [[play](https://go.dev/play/p/9FlIaFLArIL)]
-   **<big>GbkToUtf8</big>** : converts GBK encoding data to utf8 encoding data.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/convertor.md#GbkToUtf8)]
    [[play](https://go.dev/play/p/OphmHCN_9u8)]

<h3 id="cryptor"> 6. Cryptor package is for data encryption and decryption.&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### Function list:

-   **<big>AesEcbEncrypt</big>** : encrypt byte slice data with key use AES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesEcbEncrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesEcbDecrypt</big>** : decrypt byte slice data with key use AES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesEcbDecrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesCbcEncrypt</big>** : encrypt byte slice data with key use AES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCbcEncrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCbcDecrypt</big>** : decrypt byte slice data with key use AES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCbcDecrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCtrCrypt</big>** : encrypt/ decrypt byte slice data with key use AES CRC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCtrCrypt)]
    [[play](https://go.dev/play/p/SpaZO0-5Nsp)]
-   **<big>AesCfbEncrypt</big>** : encrypt byte slice data with key use AES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCfbEncrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesCfbDecrypt</big>** : decrypt byte slice data with key use AES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesCfbDecrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesOfbEncrypt</big>** : encrypt byte slice data with key use AES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesOfbEncrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>AesOfbDecrypt</big>** : decrypt byte slice data with key use AES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#AesOfbDecrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>Base64StdEncode</big>** : encode string with base64 encoding.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Base64StdEncode)]
    [[play](https://go.dev/play/p/VOaUyQUreoK)]
-   **<big>Base64StdDecode</big>** : decode string with base64 encoding.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Base64StdDecode)]
    [[play](https://go.dev/play/p/RWQylnJVgIe)]
-   **<big>DesEcbEncrypt</big>** : encrypt byte slice data with key use DES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesEcbEncrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesEcbDecrypt</big>** : decrypt byte slice data with key use DES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesEcbDecrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesCbcEncrypt</big>** : encrypt byte slice data with key use DES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCbcEncrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCbcDecrypt</big>** : decrypt byte slice data with key use DES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCbcDecrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCtrCrypt</big>** : encrypt/decrypt byte slice data with key use DES CRY algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCtrCrypt)]
    [[play](https://go.dev/play/p/9-T6OjKpcdw)]
-   **<big>DesCfbEncrypt</big>** : encrypt byte slice data with key use DES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCfbEncrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesCfbDecrypt</big>** : decrypt byte slice data with key use DES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesCfbDecrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesOfbEncrypt</big>** : encrypt byte slice data with key use DES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesOfbEncrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>DesOfbDecrypt</big>** : decrypt byte slice data with key use DES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#DesOfbDecrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>HmacMd5</big>** : return the md5 hmac hash of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacMd5)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>HmacMd5WithBase64</big>** : return the md5 hmac hash of base64 string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacMd5WithBase64)]
-   **<big>HmacSha1</big>** : return the hmac hash of string use sha1.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha1)]
    [[play](https://go.dev/play/p/1UI4oQ4WXKM)]
-   **<big>HmacSha1WithBase64</big>** : return the hmac hash of string use sha1 with base64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha1WithBase64)]
    [[play](https://go.dev/play/p/47JmmGrnF7B)]
-   **<big>HmacSha256</big>** : return the hmac hash of string use sha256.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha256)]
    [[play](https://go.dev/play/p/HhpwXxFhhC0)]
-   **<big>HmacSha256WithBase64</big>** : return the hmac hash of string use sha256 with base64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha256WithBase64)]
    [[play](https://go.dev/play/p/EKbkUvPTLwO)]
-   **<big>HmacSha512</big>** : return the hmac hash of string use sha512.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha512)]
    [[play](https://go.dev/play/p/59Od6m4A0Ud)]
-   **<big>HmacSha512WithBase64</big>** : return the hmac hash of string use sha512 with base64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#HmacSha512WithBase64)]
    [[play](https://go.dev/play/p/c6dSe3E2ydU)]
-   **<big>Md5Byte</big>** : return the md5 string of byte slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5Byte)]
    [[play](https://go.dev/play/p/suraalH8lyC)]
-   **<big>Md5ByteWithBase64</big>** : return the md5 string of byte slice with base64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5ByteWithBase64)]
    [[play](https://go.dev/play/p/Tcb-Z7LN2ax)]
-   **<big>Md5String</big>** : return the md5 value of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5String)]
    [[play](https://go.dev/play/p/1bLcVetbTOI)]
-   **<big>Md5StringWithBase64</big>** : return the md5 value of string with base64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5StringWithBase64)]
    [[play](https://go.dev/play/p/Lx4gH7Vdr5_y)]
-   **<big>Md5File</big>** : return the md5 value of file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Md5File)]
-   **<big>Sha1</big>** : return the sha1 value (SHA-1 hash algorithm) of base64 string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha1)]
    [[play](https://go.dev/play/p/_m_uoD1deMT)]
-   **<big>Sha1WithBase64</big>** : return the sha1 value (SHA-1 hash algorithm) of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha1WithBase64)]
    [[play](https://go.dev/play/p/fSyx-Gl2l2-)]
-   **<big>Sha256</big>** : return the sha256 value (SHA-256 hash algorithm) of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha256)]
    [[play](https://go.dev/play/p/tU9tfBMIAr1)]
-   **<big>Sha256WithBase64</big>** : return the sha256 value (SHA256 hash algorithm) of base64 string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha256WithBase64)]
    [[play](https://go.dev/play/p/85IXJHIal1k)]
-   **<big>Sha512</big>** : return the sha512 value (SHA-512 hash algorithm) of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha512)]
    [[play](https://go.dev/play/p/3WsvLYZxsHa)]
-   **<big>Sha512WithBase64</big>** : return the sha512 value (SHA-512 hash algorithm) of base64 string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#Sha512WithBase64)]
    [[play](https://go.dev/play/p/q_fY2rA-k5I)]
-   **<big>GenerateRsaKey</big>** : create rsa private and public pemo file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#GenerateRsaKey)]
    [[play](https://go.dev/play/p/zutRHrDqs0X)]
-   **<big>RsaEncrypt</big>** : encrypt data with ras algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaEncrypt)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>RsaDecrypt</big>** : decrypt data with ras algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaDecrypt)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>GenerateRsaKeyPair</big>** : creates rsa private and public key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#GenerateRsaKeyPair)]
    [[play](https://go.dev/play/p/sSVmkfENKMz)]
-   **<big>RsaEncryptOAEP</big>** : encrypts the given data with RSA-OAEP.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaEncryptOAEP)]
    [[play](https://go.dev/play/p/sSVmkfENKMz)]
-   **<big>RsaDecryptOAEP</big>** : decrypts the data with RSA-OAEP
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/cryptor.md#RsaDecryptOAEP)]
    [[play](https://go.dev/play/p/sSVmkfENKMz)]


<h3 id="datetime"> 7. Datetime package supports date and time format and compare. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/datetime"
```

#### Function list:

-   **<big>AddDay</big>** : add or sub day to the time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddDay)]
    [[play](https://go.dev/play/p/dIGbs_uTdFa)]
-   **<big>AddHour</big>** : add or sub day to the time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddHour)]
    [[play](https://go.dev/play/p/rcMjd7OCsi5)]
-   **<big>AddMinute</big>** : add or sub day to the time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddMinute)]
    [[play](https://go.dev/play/p/nT1heB1KUUK)]
-   **<big>AddYear</big>** : add or sub year to the time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#AddYear)]
    [[play](https://go.dev/play/p/MqW2ujnBx10)]
-   **<big>BeginOfMinute</big>** : return the date time at the begin of minute of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfMinute)]
    [[play](https://go.dev/play/p/ieOLVJ9CiFT)]
-   **<big>BeginOfHour</big>** : return the date time at the begin of hour of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfHour)]
    [[play](https://go.dev/play/p/GhdGFnDWpYs)]
-   **<big>BeginOfDay</big>** : return the date time at the begin of day of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfDay)]
    [[play](https://go.dev/play/p/94m_UT6cWs9)]
-   **<big>BeginOfWeek</big>** : return the date time at the begin of week of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfWeek)]
    [[play](https://go.dev/play/p/ynjoJPz7VNV)]
-   **<big>BeginOfMonth</big>** : return the date time at the begin of month of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfMonth)]
    [[play](https://go.dev/play/p/bWXVFsmmzwL)]
-   **<big>BeginOfYear</big>** : return the date time at the begin of year of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BeginOfYear)]
    [[play](https://go.dev/play/p/i326DSwLnV8)]
-   **<big>EndOfMinute</big>** : return the date time at the end of minute of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfMinute)]
    [[play](https://go.dev/play/p/yrL5wGzPj4z)]
-   **<big>EndOfHour</big>** : return the date time at the end of hour of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfHour)]
    [[play](https://go.dev/play/p/6ce3j_6cVqN)]
-   **<big>EndOfDay</big>** : return the date time at the end of day of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfDay)]
    [[play](https://go.dev/play/p/eMBOvmq5Ih1)]
-   **<big>EndOfWeek</big>** : return the date time at the end of week of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfWeek)]
    [[play](https://go.dev/play/p/i08qKXD9flf)]
-   **<big>EndOfMonth</big>** : return the date time at the end of month of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfMonth)]
    [[play](https://go.dev/play/p/_GWh10B3Nqi)]
-   **<big>EndOfYear</big>** : return the date time at the end of year of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#EndOfYear)]
    [[play](https://go.dev/play/p/G01cKlMCvNm)]
-   **<big>GetNowDate</big>** : return format yyyy-mm-dd of current date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNowDate)]
    [[play](https://go.dev/play/p/PvfkPpcpBBf)]
-   **<big>GetNowTime</big>** : return format hh-mm-ss of current time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNowTime)]
    [[play](https://go.dev/play/p/l7BNxCkTmJS)]
-   **<big>GetNowDateTime</big>** : return format yyyy-mm-dd hh-mm-ss of current datetime.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNowDateTime)]
    [[play](https://go.dev/play/p/pI4AqngD0al)]
-   **<big>GetTodayStartTime</big>** : return the start time of today, format: yyyy-mm-dd 00:00:00.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetTodayStartTime)]
    [[play](https://go.dev/play/p/84siyYF7t99)]
-   **<big>GetTodayEndTime</big>** : return the end time of today, format: yyyy-mm-dd 23:59:59.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetTodayEndTime)]
    [[play](https://go.dev/play/p/jjrLnfoqgn3)]
-   **<big>GetZeroHourTimestamp</big>** : return timestamp of zero hour (timestamp of 00:00).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetZeroHourTimestamp)]
    [[play](https://go.dev/play/p/QmL2oIaGE3q)]
-   **<big>GetNightTimestamp</big>** : return timestamp of zero hour (timestamp of 23:59).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#GetNightTimestamp)]
    [[play](https://go.dev/play/p/UolysR3MYP1)]
-   **<big>FormatTimeToStr</big>** : convert time to string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#FormatTimeToStr)]
    [[play](https://go.dev/play/p/_Ia7M8H_OvE)]
-   **<big>FormatStrToTime</big>** : convert string to time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#FormatStrToTime)]
    [[play](https://go.dev/play/p/1h9FwdU8ql4)]
-   **<big>NewUnix</big>** : return unix timestamp of specific time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewUnix)]
    [[play](https://go.dev/play/p/psoSuh_kLRt)]
-   **<big>NewUnixNow</big>** : return unix timestamp of current time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewUnixNow)]
    [[play](https://go.dev/play/p/U4PPx-9D0oz)]
-   **<big>NewFormat</big>** : return unix timestamp of specific time string, t should be "yyyy-mm-dd hh:mm:ss".
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>NewISO8601</big>** : return unix timestamp of specific iso8601 time string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NewISO8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]
-   **<big>ToUnix</big>** : return unix timestamp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToUnix)]
    [[play](https://go.dev/play/p/_LUiwAdocjy)]
-   **<big>ToFormat</big>** : return the time string 'yyyy-mm-dd hh:mm:ss' of unix time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>ToFormatForTpl</big>** : return the time string which format is specific tpl.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToFormatForTpl)]
    [[play](https://go.dev/play/p/nyXxXcQJ8L5)]
-   **<big>ToIso8601</big>** : return iso8601 time string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#ToIso8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]
-   **<big>IsLeapYear</big>** : check if param `year` is leap year or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#IsLeapYear)]
    [[play](https://go.dev/play/p/xS1eS2ejGew)]
-   **<big>BetweenSeconds</big>** : returns the number of seconds between two times.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#BetweenSeconds)]
    [[play](https://go.dev/play/p/n3YDRyfyXJu)]
-   **<big>DayOfYear</big>** : returns which day of the year the parameter date `t` is.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#DayOfYear)]
    [[play](https://go.dev/play/p/0hjqhTwFNlH)]
-   **<big>IsWeekend</big>** : checks if passed time is weekend or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#IsWeekend)]
    [[play](https://go.dev/play/p/cupRM5aZOIY)]
-   **<big>NowDateOrTime</big>** : returns current datetime with specific format and timezone.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#NowDateOrTime)]
    [[play](https://go.dev/play/p/EZ-begEjtT0)]
-   **<big>Timestamp</big>** : returns current second timestamp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#Timestamp)]
-   **<big>TimestampMilli</big>** : returns current mill second timestamp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#TimestampMilli)]
    [[play](https://go.dev/play/p/4gvEusOTu1T)]
-   **<big>TimestampMicro</big>** : returns current micro second timestamp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#TimestampMicro)]
    [[play](https://go.dev/play/p/2maANglKHQE)]
-   **<big>TimestampNano</big>** : returns current nano second timestamp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datetime.md#TimestampNano)]
    [[play](https://go.dev/play/p/A9Oq_COrcCF)]

<h3 id="datastructure"> 8. Datastructure package constains some common data structure. eg. list, linklist, stack, queue, set, tree, graph. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

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

#### Structure list:

-   **<big>List</big>** : a linear table, implemented with slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/list.md)]
-   **<big>CopyOnWriteList</big>** : a thread-safe list implementation that uses go slicing as its base.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/copyonwritelist.md)]
-   **<big>Link</big>** : link list structure, contains singly link and doubly link.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/link.md)]
-   **<big>Stack</big>** : stack structure(fifo), contains array stack and link stack.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/stack.md)]
-   **<big>Queue</big>** : queue structure(filo), contains array queue, circular queue, link queue and priority queue.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/queue.md)]
-   **<big>Set</big>** : a data container, like slice, but element of set is not duplicate.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/set.md)]
-   **<big>Tree</big>** : binary search tree structure.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/tree.md)]
-   **<big>Heap</big>** : a binary max heap.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/heap.md)]
-   **<big>Hashmap</big>** : hash map structure.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/datastructure/hashmap.md)]

<h3 id="fileutil"> 9. Fileutil package implements some basic functions for file operations. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/fileutil"
```

#### Function list：

-   **<big>ClearFile</big>** : write empty string to target file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ClearFile)]
    [[play](https://go.dev/play/p/NRZ0ZT-G94H)]
-   **<big>CreateFile</big>** : create file in path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CreateFile)]
    [[play](https://go.dev/play/p/lDt8PEsTNKI)]
-   **<big>CreateDir</big>** : create directory in absolute path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CreateDir)]
    [[play](https://go.dev/play/p/qUuCe1OGQnM)]
-   **<big>CopyFile</big>** :copy src file to dest file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CopyFile)]
    [[play](https://go.dev/play/p/Jg9AMJMLrJi)]
-   **<big>FileMode</big>** : return file's mode and permission.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#FileMode)]
    [[play](https://go.dev/play/p/2l2hI42fA3p)]
-   **<big>MiMeType</big>** : return file mime type.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#MiMeType)]
    [[play](https://go.dev/play/p/bd5sevSUZNu)]
-   **<big>IsExist</big>** : checks if a file or directory exists.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsExist)]
    [[play](https://go.dev/play/p/nKKXt8ZQbmh)]
-   **<big>IsLink</big>** : checks if a file is symbol link or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsLink)]
    [[play](https://go.dev/play/p/TL-b-Kzvf44)]
-   **<big>IsDir</big>** : checks if the path is directory or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsDir)]
    [[play](https://go.dev/play/p/WkVwEKqtOWk)]
-   **<big>ListFileNames</big>** : return all file names in the path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ListFileNames)]
    [[play](https://go.dev/play/p/Tjd7Y07rejl)]
-   **<big>RemoveFile</big>** : remove file, param should be file path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#RemoveFile)]
    [[play](https://go.dev/play/p/P2y0XW8a1SH)]
-   **<big>ReadFileToString</big>** : return string of file content.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadFileToString)]
    [[play](https://go.dev/play/p/cmfwp_5SQTp)]
-   **<big>ReadFileByLine</big>** : read file line by line, return string slice of file content.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadFileByLine)]
    [[play](https://go.dev/play/p/svJP_7ZrBrD)]
-   **<big>Zip</big>** : create a zip file of fpath, fpath could be a file or a directory.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#Zip)]
    [[play](https://go.dev/play/p/j-3sWBp8ik_P)]
-   **<big>ZipAppendEntry</big>** : append a single file or directory by fpath to an existing zip file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ZipAppendEntry)]
-   **<big>UnZip</big>** : unzip the zip file and save it to dest path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#UnZip)]
    [[play](https://go.dev/play/p/g0w34kS7B8m)]
-   **<big>CurrentPath</big>** : return current absolute path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#CurrentPath)]
    [[play](https://go.dev/play/p/s74a9iBGcSw)]
-   **<big>IsZipFile</big>** : checks if file is zip file or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#IsZipFile)]
    [[play](https://go.dev/play/p/9M0g2j_uF_e)]
-   **<big>FileSize</big>** : return file size in bytes.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#FileSize)]
    [[play](https://go.dev/play/p/H9Z05uD-Jjc)]
-   **<big>MTime</big>** : return file modified time(unix timestamp).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#MTime)]
    [[play](https://go.dev/play/p/s_Tl7lZoAaY)]
-   **<big>Sha</big>** : return file sha value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#Sha)]
    [[play](https://go.dev/play/p/VfEEcO2MJYf)]
-   **<big>ReadCsvFile</big>** : read file content into slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadCsvFile)]
    [[play](https://go.dev/play/p/OExTkhGEd3_u)]
-   **<big>WriteCsvFile</big>** : write content to target csv file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#WriteCsvFile)]
    [[play](https://go.dev/play/p/dAXm58Q5U1o)]
-   **<big>WriteBytesToFile</big>** : write bytes to target file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#WriteBytesToFile)]
    [[play](https://go.dev/play/p/s7QlDxMj3P8)]
-   **<big>WriteStringToFile</big>** : write string to target file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#WriteStringToFile)]
    [[play](https://go.dev/play/p/GhLS6d8lH_g)]
-   **<big>ReadFile</big>** : read file or url.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/fileutil.md#ReadFile)]

<h3 id="formatter"> 10. Formatter contains some functions for data formatting. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/formatter"
```

#### Function list:

-   **<big>Comma</big>** : add comma to a number value by every 3 numbers from right, ahead by symbol char.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#Comma)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]
-   **<big>Pretty</big>** : pretty print data to JSON string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#Pretty)]
    [[play](https://go.dev/play/p/YsciGj3FH2x)]
-   **<big>PrettyToWriter</big>** : pretty encode data to writer.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#PrettyToWriter)]
    [[play](https://go.dev/play/p/LPLZ3lDi5ma)]
-   **<big>DecimalBytes</big>** : returns a human readable byte size under decimal standard (base 1000).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#DecimalBytes)]
    [[play](https://go.dev/play/p/FPXs1suwRcs)]
-   **<big>BinaryBytes</big>** : returns a human-readable byte size under binary standard (base 1024).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#BinaryBytes)]
    [[play](https://go.dev/play/p/G9oHHMCAZxP)]
-   **<big>ParseDecimalBytes</big>** : return the human readable bytes size string into the amount it represents(base 1000).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#ParseDecimalBytes)]
    [[play](https://go.dev/play/p/Am98ybWjvjj)]
-   **<big>ParseBinaryBytes</big>** : return the human readable bytes size string into the amount it represents(base 1024).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/formatter.md#ParseBinaryBytes)]
    [[play](https://go.dev/play/p/69v1tTT62x8)]

<h3 id="function"> 11. Function package can control the flow of function execution and support part of functional programming.&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/function"
```

#### Function list:

-   **<big>After</big>** : return a function that invokes passed funcation once the returned function is called more than n times.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#After)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]
-   **<big>Before</big>** : return a function that invokes passed funcation once the returned function is called less than n times
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Before)]
    [[play](https://go.dev/play/p/0HqUDIFZ3IL)]
-   **<big>CurryFn</big>** : make a curry function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#CurryFn)]
    [[play](https://go.dev/play/p/5HopfDwANKX)]
-   **<big>Compose</big>** : compose the functions from right to left.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Compose)]
    [[play](https://go.dev/play/p/KKfugD4PKYF)]
-   **<big>Delay</big>** : call the function after delayed time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Delay)]
    [[play](https://go.dev/play/p/Ivtc2ZE-Tye)]
-   **<big>Debounced</big>** : creates a debounced function that delays invoking fn until after wait duration have elapsed since the last time the debounced function was invoked.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Debounced)]
    [[play](https://go.dev/play/p/absuEGB_GN7)]
-   **<big>Schedule</big>** : invoke function every duration time, util close the returned bool channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Schedule)]
    [[play](https://go.dev/play/p/hbON-Xeyn5N)]
-   **<big>Pipeline</big>** : takes a list of functions and returns a function whose param will be passed into the functions one by one.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Pipeline)]
    [[play](https://go.dev/play/p/mPdUVvj6HD6)]
-   **<big>Watcher</big>** : Watcher is used for record code excution time. can start/stop/reset the watch timer. get the elapsed time of function execution.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/function.md#Watcher)]
    [[play](https://go.dev/play/p/l2yrOpCLd1I)]

<h3 id="maputil"> 12. Maputil package includes some functions to manipulate map.&nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/maputil"
```

#### Function list:

-   **<big>MapTo</big>** : quick map any value to struct or any base type.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#MapTo)]
    [[play](https://go.dev/play/p/4K7KBEPgS5M)]
-   **<big>ForEach</big>** : executes iteratee funcation for every key and value pair in map.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ForEach)]
    [[play](https://go.dev/play/p/OaThj6iNVXK)]
-   **<big>Filter</big>** : iterates over map, return a new map contains all key and value pairs pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Filter)]
    [[play](https://go.dev/play/p/fSvF3wxuNG7)]
-   **<big>FilterByKeys</big>** : iterates over map, return a new map whose keys are all given keys
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#FilterByKeys)]
    [[play](https://go.dev/play/p/7ov6BJHbVqh)]
-   **<big>FilterByValues</big>** : iterates over map, return a new map whose values are all given values.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#FilterByValues)]
    [[play](https://go.dev/play/p/P3-9MdcXegR)]
-   **<big>OmitBy</big>** : the opposite of Filter, removes all the map elements for which the predicate function returns true.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OmitBy)]
    [[play](https://go.dev/play/p/YJM4Hj5hNwm)]
-   **<big>OmitByKeys</big>** : the opposite of FilterByKeys, extracts all the map elements which keys are not omitted.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OmitByKeys)]
    [[play](https://go.dev/play/p/jXGrWDBfSRp)]
-   **<big>OmitByValues</big>** : the opposite of FilterByValues. remov all elements whose value are in the give slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#OmitByValues)]
    [[play](https://go.dev/play/p/XB7Y10uw20_U)]
-   **<big>Intersect</big>** : iterates over maps, return a new map of key and value pairs in all given maps.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Intersect)]
    [[play](https://go.dev/play/p/Zld0oj3sjcC)]
-   **<big>Keys</big>** : returns a slice of the map's keys.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Keys)]
    [[play](https://go.dev/play/p/xNB5bTb97Wd)]
-   **<big>KeysBy</big>** : creates a slice whose element is the result of function mapper invoked by every map's key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#KeysBy)]
    [[play](https://go.dev/play/p/hI371iB8Up8)]
-   **<big>Merge</big>** : merge maps, next key will overwrite previous key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Merge)]
    [[play](https://go.dev/play/p/H95LENF1uB-)]
-   **<big>Minus</big>** : creates a map of whose key in mapA but not in mapB.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Minus)]
    [[play](https://go.dev/play/p/3u5U9K7YZ9m)]
-   **<big>Values</big>** : returns a slice of the map's values.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Values)]
    [[play](https://go.dev/play/p/CBKdUc5FTW6)]
-   **<big>ValuesBy</big>** : creates a slice whose element is the result of function mapper invoked by every map's value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ValuesBy)]
    [[play](https://go.dev/play/p/sg9-oRidh8f)]
-   **<big>MapKeys</big>** : transforms a map to other type map by manipulating it's keys.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#MapKeys)]
    [[play](https://go.dev/play/p/8scDxWeBDKd)]
-   **<big>MapValues</big>** : transforms a map to other type map by manipulating it's values.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#MapValues)]
    [[play](https://go.dev/play/p/g92aY3fc7Iw)]
-   **<big>Entries</big>** : transforms a map into array of key/value pairs.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Entries)]
    [[play](https://go.dev/play/p/Ltb11LNcElY)]
-   **<big>FromEntries</big>** : creates a map based on a slice of key/value pairs.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#FromEntries)]
    [[play](https://go.dev/play/p/fTdu4sCNjQO)]
-   **<big>Transform</big>** : transform a map to another type map.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#Transform)]
    [[play](https://go.dev/play/p/P6ovfToM3zj)]
-   **<big>IsDisjoint</big>** : check two map are disjoint if they have no keys in common.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#IsDisjoint)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
-   **<big>HasKey</big>** : checks if map has key or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#HasKey)]
    [[play](https://go.dev/play/p/isZZHOsDhFc)]
-   **<big>NewConcurrentMap</big>** : creates a ConcurrentMap with specific shard count.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#NewConcurrentMap)]
    [[play](https://go.dev/play/p/3PenTPETJT0)]
-   **<big>ConcurrentMap_Set</big>** : set the value for a key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Set)]
    [[play](https://go.dev/play/p/3PenTPETJT0)]
-   **<big>ConcurrentMap_Get</big>** : get the value stored in the map for a key, or nil if no.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Get)]
    [[play](https://go.dev/play/p/3PenTPETJT0)]
-   **<big>ConcurrentMap_GetOrSet</big>** : returns the existing value for the key if present.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_GetOrSet)]
    [[play](https://go.dev/play/p/aDcDApOK01a)]
-   **<big>ConcurrentMap_Delete</big>** : delete the value for a key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Delete)]
    [[play](https://go.dev/play/p/uTIJZYhpVMS)]
-   **<big>ConcurrentMap_GetAndDelete</big>** :returns the existing value for the key if present and then delete the value for the key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_GetAndDelete)]
    [[play](https://go.dev/play/p/ZyxeIXSZUiM)]
-   **<big>ConcurrentMap_Has</big>** : checks if map has the value for a key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Has)]
    [[play](https://go.dev/play/p/C8L4ul9TVwf)]
-   **<big>ConcurrentMap_Range</big>** : calls iterator sequentially for each key and value present in each of the shards in the map.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/maputil.md#ConcurrentMap_Range)]
    [[play](https://go.dev/play/p/iqcy7P8P0Pr)]

<h3 id="mathutil"> 13. Mathutil package implements some functions for math calculation. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### Function list:

-   **<big>Average</big>** :return average value of numbers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Average)]
    [[play](https://go.dev/play/p/Vv7LBwER-pz)]
-   **<big>Exponent</big>** : calculate x^n for int64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Exponent)]
    [[play](https://go.dev/play/p/uF3HGNPk8wr)]
-   **<big>Fibonacci</big>** :calculate fibonacci number before n for int.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Fibonacci)]
    [[play](https://go.dev/play/p/IscseUNMuUc)]
-   **<big>Factorial</big>** : calculate x! for uint.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Factorial)]
    [[play](https://go.dev/play/p/tt6LdOK67Nx)]
-   **<big>Max</big>** : return maximum value of numbers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Max)]
    [[play](https://go.dev/play/p/cN8DHI0rTkH)]
-   **<big>MaxBy</big>** : return the maximum value of a slice using the given comparator function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#MaxBy)]
    [[play](https://go.dev/play/p/pbe2MT-7DV2)]
-   **<big>Min</big>** : return minimum value of numbers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Min)]
    [[play](https://go.dev/play/p/21BER_mlGUj)]
-   **<big>MinBy</big>** : return the minimum value of a slice using the given comparator function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#MinBy)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
-   **<big>Percent</big>** : calculate the percentage of value to total.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Percent)]
    [[play](https://go.dev/play/p/s0NdFCtwuyd)]
-   **<big>RoundToFloat</big>** : round up to n decimal places for float64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#RoundToFloat)]
    [[play](https://go.dev/play/p/ghyb528JRJL)]
-   **<big>RoundToString</big>** : round up to n decimal places for float64, return string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#RoundToString)]
    [[play](https://go.dev/play/p/kZwpBRAcllO)]
-   **<big>TruncRound</big>** : round off n decimal places for int64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#TruncRound)]
    [[play](https://go.dev/play/p/aumarSHIGzP)]
-   **<big>Range</big>** : Creates a slice of numbers from start with specified count, element step is 1.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Range)]
    [[play](https://go.dev/play/p/9ke2opxa8ZP)]
-   **<big>RangeWithStep</big>** : Creates a slice of numbers from start to end with specified step.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Range)]
    [[play](https://go.dev/play/p/akLWz0EqOSM)]
-   **<big>AngleToRadian</big>** : converts angle value to radian value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#AngleToRadian)]
    [[play](https://go.dev/play/p/CIvlICqrHql)]
-   **<big>RadianToAngle</big>** : converts radian value to angle value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#RadianToAngle)]
    [[play](https://go.dev/play/p/dQtmOTUOMgi)]
-   **<big>PointDistance</big>** : get two points distance.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#PointDistance)]
    [[play](https://go.dev/play/p/RrG4JIaziM8)]
-   **<big>IsPrime</big>** : checks if number is prime number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#IsPrime)]
    [[play](https://go.dev/play/p/Rdd8UTHZJ7u)]
-   **<big>GCD</big>** : return greatest common divisor (GCD) of integers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#GCD)]
    [[play](https://go.dev/play/p/CiEceLSoAKB)]
-   **<big>LCM</big>** : return Least Common Multiple (LCM) of integers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#LCM)]
    [[play](https://go.dev/play/p/EjcZxfY7G_g)]
-   **<big>Cos</big>** : return the cosine of the radian argument.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Cos)]
    [[play](https://go.dev/play/p/Sm89LoIfvFq)]
-   **<big>Sin</big>** : return the sine of the radian argument.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Sin)]
    [[play](https://go.dev/play/p/TWMQlMywDsP)]
-   **<big>Log</big>** : returns the logarithm of base n.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Log)]
    [[play](https://go.dev/play/p/_d4bi8oyhat)]
-   **<big>Sum</big>** : return sum of passed numbers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Sum)]
    [[play](https://go.dev/play/p/1To2ImAMJA7)]
-   **<big>Abs</big>** : returns the absolute value of param nubmer.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/mathutil.md#Sum)]
    [[play](https://go.dev/play/p/fsyBh1Os-1d)]

<h3 id="netutil"> 14. Netutil package contains functions to get net information and send http request. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/netutil"
```

#### Function list:

-   **<big>ConvertMapToQueryString</big>** : convert map to sorted url query string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#ConvertMapToQueryString)]
    [[play](https://go.dev/play/p/jnNt_qoSnRi)]
-   **<big>EncodeUrl</big>** : encode url(?a=1&b=[2] -> ?a=1&b=%5B2%5D).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#EncodeUrl)]
    [[play](https://go.dev/play/p/bsZ6BRC4uKI)]
-   **<big>GetInternalIp</big>** : return internal ipv4.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetInternalIp)]
    [[play](https://go.dev/play/p/5mbu-gFp7ei)]
-   **<big>GetIps</big>** : return all ipv4 of current system.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetIps)]
    [[play](https://go.dev/play/p/NUFfcEmukx1)]
-   **<big>GetMacAddrs</big>** : return mac address of current system.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetMacAddrs)]
    [[play](https://go.dev/play/p/Rq9UUBS_Xp1)]
-   **<big>GetPublicIpInfo</big>** : return [public ip information](http://ip-api.com/json/).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetPublicIpInfo)]
    [[play](https://go.dev/play/p/YDxIfozsRHR)]
-   **<big>GetRequestPublicIp</big>** : return the http request public ip.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#GetRequestPublicIp)]
    [[play](https://go.dev/play/p/kxU-YDc_eBo)]
-   **<big>IsPublicIP</big>** : verify a ip is public or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsPublicIP)]
    [[play](https://go.dev/play/p/nmktSQpJZnn)]
-   **<big>IsInternalIP</big>** : verify an ip is intranet or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsInternalIP)]
    [[play](https://go.dev/play/p/sYGhXbgO4Cb)]
-   **<big>HttpRequest</big>** : a composed http request used for HttpClient send request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>HttpClient</big>** : a http client tool, used for sending http request
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpClient)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>SendRequest</big>** : send http request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#SendRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>DecodeResponse</big>** : decode http response into target object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#DecodeResponse)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>StructToUrlValues</big>** : convert struct to url valuse.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#StructToUrlValues)]
    [[play](https://go.dev/play/p/pFqMkM40w9z)]
-   **<big>HttpGet<sup>deprecated</sup></big>** : send http get request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpGet)]
-   **<big>HttpDelete<sup>deprecated</sup></big>** : send http delete request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpDelete)]
-   **<big>HttpPost<sup>deprecated</sup></big>** : send http post request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpPost)]
-   **<big>HttpPut<sup>deprecated</sup></big>** : send http put request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpPut)]
-   **<big>HttpPatch<sup>deprecated</sup></big>** : send http patch request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#HttpPatch)]
-   **<big>ParseHttpResponse</big>** : decode http response into target object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#ParseHttpResponse)]
-   **<big>DownloadFile</big>** : download the file exist in url to a local file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#DownloadFile)]
-   **<big>UploadFile</big>** : upload the file to a server.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#UploadFile)]
-   **<big>IsPingConnected</big>** : checks if can ping the specified host or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsPingConnected)]
    [[play](https://go.dev/play/p/q8OzTijsA87)]
-   **<big>IsTelnetConnected</big>** : checks if can if can telnet the specified host or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/netutil.md#IsTelnetConnected)]
    [[play](https://go.dev/play/p/yiLCGtQv_ZG)]

<h3 id="pointer"> 15. Pointer package contains some util functions to operate go pointer. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/pointer"
```

#### Function list:

-   **<big>ExtractPointer</big>** : return the underlying value by the given interface type.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#ExtractPointer)]
    [[play](https://go.dev/play/p/D7HFjeWU2ZP)]
-   **<big>Of</big>** : return a pointer to the value `v`.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#Of)]
    [[play](https://go.dev/play/p/HFd70x4DrMj)]
-   **<big>Unwrap</big>** : return the value from the pointer.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#Unwrap)]
    [[play](https://go.dev/play/p/cgeu3g7cjWb)]
-   **<big>UnwarpOr</big>** : UnwarpOr returns the value from the pointer or fallback if the pointer is nil.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#UnwrapOr)]
    [[play](https://go.dev/play/p/mmNaLC38W8C)]
-   **<big>UnwarpOrDefault</big>** : UnwarpOrDefault returns the value from the pointer or the default value if the pointer is nil.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/pointer.md#UnwrapOrDefault)]
    [[play](https://go.dev/play/p/ZnGIHf8_o4E)]

<h3 id="random"> 16. Random package implements some basic functions to generate random int and string. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/random"
```

#### Function list:

-   **<big>RandBytes</big>** : generate random byte slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandBytes)]
    [[play](https://go.dev/play/p/EkiLESeXf8d)]
-   **<big>RandInt</big>** : generate random int number between min and max.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandInt)]
    [[play](https://go.dev/play/p/pXyyAAI5YxD)]
-   **<big>RandString</big>** : generate random string of specific length.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandString)]
    [[play](https://go.dev/play/p/W2xvRUXA7Mi)]
-   **<big>RandUpper</big>** : generate a random upper case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandUpper)]
    [[play](https://go.dev/play/p/29QfOh0DVuh)]
-   **<big>RandLower</big>** : generate a random lower case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandLower)]
    [[play](https://go.dev/play/p/XJtZ471cmtI)]
-   **<big>RandNumeral</big>** : generate a random numeral string of specific length.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandNumeral)]
    [[play](https://go.dev/play/p/g4JWVpHsJcf)]
-   **<big>RandNumeralOrLetter</big>** : generate a random numeral or letter string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandNumeralOrLetter)]
    [[play](https://go.dev/play/p/19CEQvpx2jD)]
-   **<big>UUIdV4</big>** : generate a random UUID of version 4 according to RFC 4122.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#UUIdV4)]
    [[play](https://go.dev/play/p/_Z9SFmr28ft)]
-   **<big>RandUniqueIntSlice</big>** : generate a slice of random int of length n that do not repeat.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/random.md#RandUniqueIntSlice)]
    [[play](https://go.dev/play/p/uBkRSOz73Ec)]

<h3 id="retry"> 17. Retry package is for executing a function repeatedly until it was successful or canceled by the context. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### Function list:

-   **<big>Context</big>** : set retry context config option.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#Context)]
    [[play](https://go.dev/play/p/xnAOOXv9GkS)]
-   **<big>Retry</big>** : executes the retryFunc repeatedly until it was successful or canceled by the context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#Retry)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryFunc</big>** : function that retry executes.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryFunc)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryDuration</big>** : set duration of retry
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryDuration)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryTimes</big>** : set times of retry.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/retry.md#RetryTimes)]
    [[play](https://go.dev/play/p/ssfVeU2SwLO)]

<h3 id="slice"> 18. Slice contains some functions to manipulate slice. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### Function list:

-   **<big>AppendIfAbsent</big>** : if the item is absent,append it to the slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#AppendIfAbsent)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>Contain</big>** : check if the value is in the slice or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Contain)]
    [[play](https://go.dev/play/p/_454yEHcNjf)]
-   **<big>ContainBy</big>** : returns true if predicate function return true.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ContainBy)]
    [[play](https://go.dev/play/p/49tkHfX4GNc)]
-   **<big>ContainSubSlice</big>** : check if the slice contain a given subslice or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ContainSubSlice)]
    [[play](https://go.dev/play/p/bcuQ3UT6Sev)]
-   **<big>Chunk</big>** : creates a slice of elements split into groups the length of size.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Chunk)]
    [[play](https://go.dev/play/p/b4Pou5j2L_C)]
-   **<big>Compact</big>** : creates an slice with all falsey values removed. The values false, nil, 0, and "" are falsey.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Compact)]
    [[play](https://go.dev/play/p/pO5AnxEr3TK)]
-   **<big>Concat</big>** : creates a new slice concatenating slice with any additional slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Concat)]
    [[play](https://go.dev/play/p/gPt-q7zr5mk)]
-   **<big>Count</big>** : returns the number of occurrences of the given item in the slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Count)]
    [[play](https://go.dev/play/p/Mj4oiEnQvRJ)]
-   **<big>CountBy</big>** : iterates over elements of slice with predicate function, returns the number of all matched elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#CountBy)]
    [[play](https://go.dev/play/p/tHOccTMDZCC)]
-   **<big>Difference</big>** : creates an slice of whose element in slice but not in compared slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Difference)]
    [[play](https://go.dev/play/p/VXvadzLzhDa)]
-   **<big>DifferenceBy</big>** : accepts iteratee which is invoked for each element of slice and values to generate the criterion by which they're compared.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DifferenceBy)]
    [[play](https://go.dev/play/p/DiivgwM5OnC)]
-   **<big>DifferenceWith</big>** : accepts comparator which is invoked to compare elements of slice to values.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DifferenceWith)]
    [[play](https://go.dev/play/p/v2U2deugKuV)]
-   **<big>DeleteAt</big>** : delete the element of slice from specific start index to end index - 1.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DeleteAt)]
    [[play](https://go.dev/play/p/pJ-d6MUWcvK)]
-   **<big>Drop</big>** : drop n elements from the start of a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Drop)]
    [[play](https://go.dev/play/p/jnPO2yQsT8H)]
-   **<big>DropRight</big>** : drop n elements from the end of a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DropRight)]
    [[play](https://go.dev/play/p/8bcXvywZezG)]
-   **<big>DropWhile</big>** : drop n elements from the start of a slice while predicate function returns true.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DropWhile)]
    [[play](https://go.dev/play/p/4rt252UV_qs)]
-   **<big>DropRightWhile</big>** : drop n elements from the end of a slice while predicate function returns true.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#DropRightWhile)]
    [[play](https://go.dev/play/p/6wyK3zMY56e)]
-   **<big>Equal</big>** : checks if two slices are equal: the same length and all elements' order and value are equal.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Equal)]
    [[play](https://go.dev/play/p/WcRQJ37ifPa)]
-   **<big>EqualWith</big>** : checks if two slices are equal with comparator func.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#EqualWith)]
    [[play](https://go.dev/play/p/b9iygtgsHI1)]
-   **<big>Every</big>** : return true if all of the values in the slice pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Every)]
    [[play](https://go.dev/play/p/R8U6Sl-j8cD)]
-   **<big>Filter</big>** : iterates over elements of slice, returning an slice of all elements pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Filter)]
    [[play](https://go.dev/play/p/SdPna-7qK4T)]
-   **<big>FilterMap</big>** : returns a slice which apply both filtering and mapping to the given slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FilterMap)]
    [[play](https://go.dev/play/p/J94SZ_9MiIe)]
-   **<big>Find<sup>deprecated</sup></big>** : iterates over elements of slice, returning the first one that passes a truth test on predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Find)]
    [[play](https://go.dev/play/p/CBKeBoHVLgq)]
-   **<big>FindBy</big>** : iterates over elements of slice, returning the first one that passes a truth test on predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FindBy)]
    [[play](https://go.dev/play/p/n1lysBYl-GB)]
-   **<big>FindLast<sup>deprecated</sup></big>** : return the last item that passes a truth test on predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FindLast)]
    [[play](https://go.dev/play/p/FFDPV_j7URd)]
-   **<big>FindLastBy</big>** : iterates over elements of slice, returning the last one that passes a truth test on predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FindLastBy)]
    [[play](https://go.dev/play/p/8iqomzyCl_s)]
-   **<big>Flatten</big>** : flattens slice one level.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Flatten)]
    [[play](https://go.dev/play/p/hYa3cBEevtm)]
-   **<big>FlattenDeep</big>** : flattens slice recursive to one level.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FlattenDeep)]
    [[play](https://go.dev/play/p/yjYNHPyCFaF)]
-   **<big>FlatMap</big>** : manipulates a slice and transforms and flattens it to a slice of another type.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#FlatMap)]
    [[play](https://go.dev/play/p/_QARWlWs1N_F)]
-   **<big>ForEach</big>** : iterates over elements of slice and invokes function for each element.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ForEach)]
    [[play](https://go.dev/play/p/DrPaa4YsHRF)]
-   **<big>ForEachWithBreak</big>** : iterates over elements of slice and invokes function for each element, when iteratee return false, will break the for each loop.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ForEachWithBreak)]
    [[play](https://go.dev/play/p/qScs39f3D9W)]
-   **<big>GroupBy</big>** : iterate over elements of the slice, each element will be group by criteria, returns two slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#GroupBy)]
    [[play](https://go.dev/play/p/QVkPxzPR0iA)]
-   **<big>GroupWith</big>** : return a map composed of keys generated from the resultults of running each element of slice thru iteratee.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#GroupWith)]
    [[play](https://go.dev/play/p/ApCvMNTLO8a)]
-   **<big>IntSlice<sup>deprecated</sup></big>** : convert param to int slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IntSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>InterfaceSlice<sup>deprecated</sup></big>** : convert param to interface slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#InterfaceSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>Intersection</big>** : creates a slice of unique elements that included by all slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Intersection)]
    [[play](https://go.dev/play/p/anJXfB5wq_t)]
-   **<big>InsertAt</big>** : insert the value or other slice into slice at index.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#InsertAt)]
    [[play](https://go.dev/play/p/hMLNxPEGJVE)]
-   **<big>IndexOf</big>** : returns the index at which the first occurrence of an item is found in a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IndexOf)]
    [[play](https://go.dev/play/p/MRN1f0FpABb)]
-   **<big>LastIndexOf</big>** : returns the index at which the last occurrence of the item is found in a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#LastIndexOf)]
    [[play](https://go.dev/play/p/DokM4cf1IKH)]
-   **<big>Map</big>** : creates an slice of values by running each element of slice thru iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Map)]
    [[play](https://go.dev/play/p/biaTefqPquw)]
-   **<big>Merge</big>** : merge all given slices into one slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Merge)]
    [[play](https://go.dev/play/p/lbjFp784r9N)]
-   **<big>Reverse</big>** : return slice of element order is reversed to the given slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Reverse)]
    [[play](https://go.dev/play/p/8uI8f1lwNrQ)]
-   **<big>Reduce<sup>deprecated</sup></big>** : creates an slice of values by running each element of slice thru iteratee function.(Deprecated: use ReduceBy)
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Reduce)]
    [[play](https://go.dev/play/p/_RfXJJWIsIm)]
-   **<big>ReduceBy</big>** : produces a value from slice by accumulating the result of each element as passed through the reducer function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ReduceBy)]
    [[play](https://go.dev/play/p/YKDpLi7gtee)]
-   **<big>ReduceRight</big>** : ReduceRight is like ReduceBy, but it iterates over elements of slice from right to left.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ReduceRight)]
    [[play](https://go.dev/play/p/qT9dZC03A1K)]
-   **<big>Replace</big>** : returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Replace)]
    [[play](https://go.dev/play/p/P5mZp7IhOFo)]
-   **<big>ReplaceAll</big>** : returns a copy of the slice with all non-overlapping instances of old replaced by new.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ReplaceAll)]
    [[play](https://go.dev/play/p/CzqXMsuYUrx)]
-   **<big>Repeat</big>** : creates a slice with length n whose elements are passed item.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Repeat)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>Shuffle</big>** : shuffle the slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Shuffle)]
    [[play](https://go.dev/play/p/YHvhnWGU3Ge)]
-   **<big>IsAscending</big>** : Checks if a slice is ascending order.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsAscending)]
    [[play](https://go.dev/play/p/9CtsFjet4SH)]
-   **<big>IsDescending</big>** : Checks if a slice is descending order.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsDescending)]
    [[play](https://go.dev/play/p/U_LljFXma14)]
-   **<big>IsSorted</big>** : Checks if a slice is sorted (ascending or descending).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsSorted)]
    [[play](https://go.dev/play/p/nCE8wPLwSA-)]
-   **<big>IsSortedByKey</big>** : Checks if a slice is sorted by iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#IsSortedByKey)]
    [[play](https://go.dev/play/p/tUoGB7DOHI4)]
-   **<big>Sort</big>** : sorts a slice of any ordered type(number or string).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Sort)]
    [[play](https://go.dev/play/p/V9AVjzf_4Fk)]
-   **<big>SortBy</big>** : sorts the slice in ascending order as determined by the less function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#SortBy)]
    [[play](https://go.dev/play/p/DAhLQSZEumm)]
-   **<big>SortByField<sup>deprecated</sup></big>** : return sorted slice by specific field.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#SortByField)]
    [[play](https://go.dev/play/p/fU1prOBP9p1)]
-   **<big>Some</big>** : return true if any of the values in the list pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Some)]
    [[play](https://go.dev/play/p/4pO9Xf9NDGS)]
-   **<big>StringSlice<sup>deprecated</sup></big>** : convert param to slice of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#StringSlice)]
    [[play](https://go.dev/play/p/W0TZDWCPFcI)]
-   **<big>SymmetricDifference</big>** : the symmetric difference of two slice, also known as the disjunctive union.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#h42nJX5xMln)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>ToSlice</big>** : returns a slices of a variable parameter transformation.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ToSlice)]
    [[play](https://go.dev/play/p/YzbzVq5kscN)]
-   **<big>ToSlicePointer</big>** : returns a pointer to the slices of a variable parameter transformation.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#ToSlicePointer)]
    [[play](https://go.dev/play/p/gx4tr6_VXSF)]
-   **<big>Unique</big>** : remove duplicate elements in slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Unique)]
    [[play](https://go.dev/play/p/AXw0R3ZTE6a)]
-   **<big>UniqueBy</big>** : call iteratee func with every item of slice, then remove duplicated.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UniqueBy)]
    [[play](https://go.dev/play/p/UR323iZLDpv)]
-   **<big>Union</big>** : creates a slice of unique elements, in order, from all given slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Union)]
    [[play](https://go.dev/play/p/hfXV1iRIZOf)]
-   **<big>UnionBy</big>** : accepts iteratee which is invoked for each element of each slice, then union slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UnionBy)]
    [[play](https://go.dev/play/p/HGKHfxKQsFi)]
-   **<big>UpdateAt</big>** : update the slice element at index.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#UpdateAt)]
    [[play](https://go.dev/play/p/f3mh2KloWVm)]
-   **<big>Without</big>** : creates a slice excluding all given items.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Without)]
    [[play](https://go.dev/play/p/bwhEXEypThg)]
-   **<big>KeyBy</big>** : converts a slice to a map based on a callback function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#KeyBy)]
    [[play](https://go.dev/play/p/uXod2LWD1Kg)]
-   **<big>Join</big>** : join the slice item with specify separator.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Join)]
    [[play](https://go.dev/play/p/huKzqwNDD7V)]
-   **<big>Partition</big>** : partition all slice elements with the evaluation of the given predicate functions.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/slice.md#Partition)]
    [[play](https://go.dev/play/p/lkQ3Ri2NQhV)]

<h3 id="stream"> 19. Stream package implements a sequence of elements supporting sequential and operations. this package is an experiment to explore if stream in go can work as the way java does. its function is very limited. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/stream"
```

#### Function list:

-   **<big>Of</big>** : creates a stream whose elements are the specified values.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Of)]
    [[play](https://go.dev/play/p/jI6_iZZuVFE)]
-   **<big>FromSlice</big>** : creates a stream from slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FromSlice)]
    [[play](https://go.dev/play/p/wywTO0XZtI4)]
-   **<big>FromChannel</big>** : creates a stream from channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FromChannel)]
    [[play](https://go.dev/play/p/9TZYugGMhXZ)]
-   **<big>FromRange</big>** : creates a number stream from start to end. both start and end are included. [start, end]
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FromRange)]
    [[play](https://go.dev/play/p/9Ex1-zcg-B-)]
-   **<big>Generate</big>** : creates a stream where each element is generated by the provided generater function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Generate)]
    [[play](https://go.dev/play/p/rkOWL1yA3j9)]
-   **<big>Concat</big>** : creates a lazily concatenated stream whose elements are all the elements of the first stream followed by all the elements of the second stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Concat)]
    [[play](https://go.dev/play/p/HM4OlYk_OUC)]
-   **<big>Distinct</big>** : creates returns a stream that removes the duplicated items.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Distinct)]
    [[play](https://go.dev/play/p/eGkOSrm64cB)]
-   **<big>Filter</big>** : returns a stream consisting of the elements of this stream that match the given predicate.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Filter)]
    [[play](https://go.dev/play/p/MFlSANo-buc)]
-   **<big>Map</big>** : returns a stream consisting of the elements of this stream that apply the given function to elements of stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Map)]
    [[play](https://go.dev/play/p/OtNQUImdYko)]
-   **<big>Peek</big>** : returns a stream consisting of the elements of this stream, additionally performing the provided action on each element as elements are consumed from the resulting stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Peek)]
    [[play](https://go.dev/play/p/u1VNzHs6cb2)]
-   **<big>Skip</big>** : returns a stream consisting of the remaining elements of this stream after discarding the first n elements of the stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Skip)]
    [[play](https://go.dev/play/p/fNdHbqjahum)]
-   **<big>Limit</big>** : returns a stream consisting of the elements of this stream, truncated to be no longer than maxSize in length.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Limit)]
    [[play](https://go.dev/play/p/qsO4aniDcGf)]
-   **<big>Reverse</big>** : returns a stream whose elements are reverse order of given stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Reverse)]
    [[play](https://go.dev/play/p/A8_zkJnLHm4)]
-   **<big>Range</big>** : returns a stream whose elements are in the range from start(included) to end(excluded) original stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Range)]
    [[play](https://go.dev/play/p/indZY5V2f4j)]
-   **<big>Sorted</big>** : returns a stream consisting of the elements of this stream, sorted according to the provided less function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Sorted)]
    [[play](https://go.dev/play/p/XXtng5uonFj)]
-   **<big>ForEach</big>** : performs an action for each element of this stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#ForEach)]
    [[play](https://go.dev/play/p/Dsm0fPqcidk)]
-   **<big>Reduce</big>** : performs a reduction on the elements of this stream, using an associative accumulation function, and returns an Optional describing the reduced value, if any.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Reduce)]
    [[play](https://go.dev/play/p/6uzZjq_DJLU)]
-   **<big>FindFirst</big>** : returns the first element of this stream and true, or zero value and false if the stream is empty.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FindFirst)]
    [[play](https://go.dev/play/p/9xEf0-6C1e3)]
-   **<big>FindLast</big>** : returns the last element of this stream and true, or zero value and false if the stream is empty.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#FindLast)]
    [[play](https://go.dev/play/p/WZD2rDAW-2h)]
-   **<big>Max</big>** : returns the maximum element of this stream according to the provided less function. less fuction: a > b
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Max)]
    [[play](https://go.dev/play/p/fm-1KOPtGzn)]
-   **<big>Min</big>** : returns the minimum element of this stream according to the provided less function. less fuction: a < b
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Min)]
    [[play](https://go.dev/play/p/vZfIDgGNRe_0)]
-   **<big>AllMatch</big>** : returns whether all elements of this stream match the provided predicate.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#AllMatch)]
    [[play](https://go.dev/play/p/V5TBpVRs-Cx)]
-   **<big>AnyMatch</big>** : returns whether any elements of this stream match the provided predicate.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#AnyMatch)]
    [[play](https://go.dev/play/p/PTCnWn4OxSn)]
-   **<big>NoneMatch</big>** : returns whether no elements of this stream match the provided predicate.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#NoneMatch)]
    [[play](https://go.dev/play/p/iWS64pL1oo3)]
-   **<big>Count</big>** : returns the count of elements in the stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#Count)]
    [[play](https://go.dev/play/p/r3koY6y_Xo-)]
-   **<big>ToSlice</big>** : returns the elements in the stream.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/stream.md#ToSlice)]
    [[play](https://go.dev/play/p/jI6_iZZuVFE)]

<h3 id="structs"> 20. Structs package provides several high level functions to manipulate struct, tag, and field. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/structs"
```

#### Function list:

-   **<big>New</big>** : creates a `Struct` instance.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/struct.md#New)]
-   **<big>ToMap</big>** : converts a valid struct to a map.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/struct.md#ToMap)]
-   **<big>Fields</big>** : get all fields of a given struct, that the fields are abstract struct field.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/struct.md#Fields)]
-   **<big>IsStruct</big>** : check if the struct is valid.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/struct.md#IsStruct)]
-   **<big>Tag</big>** : get a `Tag` of the `Field`, `Tag` is a abstract struct field tag
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#Tag)]
-   **<big>Name</big>** : get the field name.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#Name)]
-   **<big>Value</big>** : get the `Field` underlying value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#Value)]
-   **<big>Kind</big>** : get the field's kind
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#Kind)]
-   **<big>IsEmbedded</big>** : check if the field is an embedded field.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#IsEmbedded)]
-   **<big>IsExported</big>** : check if the field is exporte
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#IsExported)]
-   **<big>IsZero</big>** : check if the field is zero value
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#IsZero)]
-   **<big>IsSlice</big>** : check if the field is a slice
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/structs/field.md#IsSlice)]

<h3 id="strutil"> 21. Strutil package contains some functions to manipulate string. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### Function list:

-   **<big>After</big>** : returns the substring after the first occurrence of a specific string in the source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#After)]
    [[play](https://go.dev/play/p/RbCOQqCDA7m)]
-   **<big>AfterLast</big>** : returns the substring after the last occurrence of a specific string in the source string. [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#AfterLast)]
    [[play](https://go.dev/play/p/1TegARrb8Yn)]
-   **<big>Before</big>** : returns the substring before the first occurrence of a specific string in the source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Before)]
    [[play](https://go.dev/play/p/JAWTZDS4F5w)]
-   **<big>BeforeLast</big>** : returns the substring before the last occurrence of a specific string in the source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#BeforeLast)]
    [[play](https://go.dev/play/p/pJfXXAoG_Te)]
-   **<big>CamelCase</big>** : coverts source string to its camelCase string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#CamelCase)]
    [[play](https://go.dev/play/p/9eXP3tn2tUy)]
-   **<big>Capitalize</big>** : converts the first character of source string to upper case and the remaining to lower case.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Capitalize)]
    [[play](https://go.dev/play/p/2OAjgbmAqHZ)]
-   **<big>ContainsAll</big>** : return true if target string contains all the substrings.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#ContainsAll)]
    [[play](https://go.dev/play/p/KECtK2Os4zq)]
-   **<big>ContainsAny</big>** : return true if target string contains any one of the substrings.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#ContainsAny)]
    [[play](https://go.dev/play/p/dZGSSMB3LXE)]
-   **<big>IsString</big>** : checks if the parameter value data type is string or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IsString)]
    [[play](https://go.dev/play/p/IOgq7oF9ERm)]
-   **<big>KebabCase</big>** : coverts string to kebab-case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#KebabCase)]
    [[play](https://go.dev/play/p/dcZM9Oahw-Y)]
-   **<big>UpperKebabCase</big>** : coverts string to upper KEBAB-CASE string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#UpperKebabCase)]
    [[play](https://go.dev/play/p/zDyKNneyQXk)]
-   **<big>LowerFirst</big>** : converts the first character of string to lower case.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#LowerFirst)]
    [[play](https://go.dev/play/p/CbzAyZmtJwL)]
-   **<big>UpperFirst</big>** : converts the first character of string to upper case.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#UpperFirst)]
    [[play](https://go.dev/play/p/sBbBxRbs8MM)]
-   **<big>Pad</big>** : pads string on the left and right side if it's shorter than size.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Pad)]
    [[play](https://go.dev/play/p/NzImQq-VF8q)]
-   **<big>PadEnd</big>** : pads string with given characters on the right side if it's shorter than limit size. Padding characters are truncated if they exceed size.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#PadEnd)]
    [[play](https://go.dev/play/p/9xP8rN0vz--)]
-   **<big>PadStart</big>** : pads string with given characters on the left side if it's shorter than limit size. Padding characters are truncated if they exceed size.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#PadStart)]
    [[play](https://go.dev/play/p/xpTfzArDfvT)]
-   **<big>Reverse</big>** : returns string whose char order is reversed to the given string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Reverse)]
    [[play](https://go.dev/play/p/adfwalJiecD)]
-   **<big>SnakeCase</big>** : coverts string to snake_case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SnakeCase)]
    [[play](https://go.dev/play/p/tgzQG11qBuN)]
-   **<big>UpperSnakeCase</big>** : coverts string to upper SNAKE_CASE string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#UpperSnakeCase)]
    [[play](https://go.dev/play/p/4COPHpnLx38)]
-   **<big>SplitEx</big>** : split a given string which can control the result slice contains empty string or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SplitEx)]
    [[play](https://go.dev/play/p/Us-ySSbWh-3)]
-   **<big>Substring</big>** : returns a substring of the specific length starting at the specific offset position.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Substring)]
    [[play](https://go.dev/play/p/q3sM6ehnPDp)]
-   **<big>Wrap</big>** : wrap a string with given string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Wrap)]
    [[play](https://go.dev/play/p/KoZOlZDDt9y)]
-   **<big>Unwrap</big>** : unwrap a given string from anther string. will change source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Unwrap)]
    [[play](https://go.dev/play/p/Ec2q4BzCpG-)]
-   **<big>SplitWords</big>** : splits a string into words, word only contains alphabetic characters.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SplitWords)]
    [[play](https://go.dev/play/p/KLiX4WiysMM)]
-   **<big>WordCount</big>** : return the number of meaningful word of a string, word only contains alphabetic characters.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#WordCount)]
    [[play](https://go.dev/play/p/bj7_odx3vRf)]
-   **<big>RemoveNonPrintable</big>** : remove non-printable characters from a string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#RemoveNonPrintable)]
    [[play](https://go.dev/play/p/og47F5x_jTZ)]
-   **<big>StringToBytes</big>** : converts a string to byte slice without a memory allocation.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#StringToBytes)]
    [[play](https://go.dev/play/p/7OyFBrf9AxA)]
-   **<big>BytesToString</big>** : converts a byte slice to string without a memory allocation.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#BytesToString)]
    [[play](https://go.dev/play/p/6c68HRvJecH)]
-   **<big>IsBlank</big>** : checks if a string is whitespace or empty.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IsBlank)]
    [[play](https://go.dev/play/p/6zXRH_c0Qd3)]
-   **<big>IsNotBlank</big>** : checks if a string is not whitespace or not empty.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IsNotBlank)]
    [[play](https://go.dev/play/p/e_oJW0RAquA)]
-   **<big>HasPrefixAny</big>** : checks if a string starts with any of an array of specified strings.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#HasPrefixAny)]
    [[play](https://go.dev/play/p/8UUTl2C5slo)]
-   **<big>HasSuffixAny</big>** : checks if a string ends with any of an array of specified strings.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#HasSuffixAny)]
    [[play](https://go.dev/play/p/sKWpCQdOVkx)]
-   **<big>IndexOffset</big>** : returns the index of the first instance of substr in string after offsetting the string by `idxFrom`.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#IndexOffset)]
    [[play](https://go.dev/play/p/qZo4lV2fomB)]
-   **<big>ReplaceWithMap</big>** : returns a copy of `str`, which is replaced by a map in unordered way, case-sensitively.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#ReplaceWithMap)]
    [[play](https://go.dev/play/p/h3t7CNj2Vvu)]
-   **<big>Trim</big>** : strips whitespace (or other characters) from the beginning and end of a string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#Trim)]
    [[play](https://go.dev/play/p/Y0ilP0NRV3j)]
-   **<big>SplitAndTrim</big>** : splits string `str` by a string `delimiter` to a slice, and calls Trim to every element of slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#SplitAndTrim)]
    [[play](https://go.dev/play/p/ZNL6o4SkYQ7)]
-   **<big>HideString</big>** : Hide some chars in source string with param `replaceChar`.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#HideString)]
    [[play](https://go.dev/play/p/pzbaIVCTreZ)]
-   **<big>RemoveWhiteSpace</big>** : remove whitespace characters from a string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/strutil.md#RemoveWhiteSpace)]
    [[play](https://go.dev/play/p/HzLC9vsTwkf)]


<h3 id="system"> 22. System package contain some functions about os, runtime, shell command. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/system"
```

#### Function list:

-   **<big>IsWindows</big>** : check if current os is windows.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#IsWindows)]
    [[play](https://go.dev/play/p/XzJULbzmf9m)]
-   **<big>IsLinux</big>** : check if current os is linux.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#IsLinux)]
    [[play](https://go.dev/play/p/zIflQgZNuxD)]
-   **<big>IsMac</big>** : check if current os is macos.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#IsMac)]
    [[play](https://go.dev/play/p/Mg4Hjtyq7Zc)]
-   **<big>GetOsEnv</big>** : get the value of the environment variable named by the key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#GetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>SetOsEnv</big>** : set the value of the environment variable named by the key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#SetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>RemoveOsEnv</big>** : remove a single environment variable.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#RemoveOsEnv)]
    [[play](https://go.dev/play/p/fqyq4b3xUFQ)]
-   **<big>CompareOsEnv</big>** : get env named by the key and compare it with passed env.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#CompareOsEnv)]
    [[play](https://go.dev/play/p/BciHrKYOHbp)]
-   **<big>ExecCommand</big>** : execute command, return the stdout and stderr string of command, and error if error occurs.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#ExecCommand)]
    [[play](https://go.dev/play/p/n-2fLyZef-4)]
-   **<big>GetOsBits</big>** : return current os bits (32 or 64).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/system.md#GetOsBits)]
    [[play](https://go.dev/play/p/ml-_XH3gJbW)]

<h3 id="tuple"> 23. Tuple package implements tuple data type and some operations on it. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/tuple"
```

#### Function list:

-   **<big>Tuple2</big>** : represents a 2 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple2)]
    [[play](https://go.dev/play/p/3sHVqBQpLYN)]
-   **<big>Tuple2_Unbox</big>** : returns values in Tuple2.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple2_Unbox)]
    [[play](https://go.dev/play/p/0fD1qfCVwjm)]
-   **<big>Zip2</big>** : create a slice of Tuple2, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip2)]
    [[play](https://go.dev/play/p/4ncWJJ77Xio)]
-   **<big>Unzip2</big>** : create a group of slice from a slice of Tuple2.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip2)]
    [[play](https://go.dev/play/p/KBecr60feXb)]
-   **<big>Tuple3</big>** : represents a 3 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple3)]
    [[play](https://go.dev/play/p/FtH2sdCLlCf)]
-   **<big>Tuple3_Unbox</big>** : returns values in Tuple3.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple3_Unbox)]
    [[play](https://go.dev/play/p/YojLy-id1BS)]
-   **<big>Zip3</big>** : create a slice of Tuple3, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip3)]
    [[play](https://go.dev/play/p/97NgmsTILfu)]
-   **<big>Unzip3</big>** : create a group of slice from a slice of Tuple3.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip3)]
    [[play](https://go.dev/play/p/bba4cpAa7KO)]
-   **<big>Tuple4</big>** : represents a 4 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple4)]
    [[play](https://go.dev/play/p/D2EqDz096tk)]
-   **<big>Tuple4_Unbox</big>** : returns values in Tuple4.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple4_Unbox)]
    [[play](https://go.dev/play/p/ACj9YuACGgW)]
-   **<big>Zip4</big>** : create a slice of Tuple4, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip4)]
    [[play](https://go.dev/play/p/PEmTYVK5hL4)]
-   **<big>Unzip4</big>** : create a group of slice from a slice of Tuple4.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip4)]
    [[play](https://go.dev/play/p/rb8z4gyYSRN)]
-   **<big>Tuple5</big>** : represents a 5 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple5)]
    [[play](https://go.dev/play/p/2WndmVxPg-r)]
-   **<big>Tuple5_Unbox</big>** : returns values in Tuple4.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple5_Unbox)]
    [[play](https://go.dev/play/p/GyIyZHjCvoS)]
-   **<big>Zip5</big>** : create a slice of Tuple5, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip5)]
    [[play](https://go.dev/play/p/fCAAJLMfBIP)]
-   **<big>Unzip5</big>** : create a group of slice from a slice of Tuple5.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip5)]
    [[play](https://go.dev/play/p/gyl6vKfhqPb)]
-   **<big>Tuple6</big>** : represents a 6 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple6)]
    [[play](https://go.dev/play/p/VjqcCwEJZbs)]
-   **<big>Tuple6_Unbox</big>** : returns values in Tuple6.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple6_Unbox)]
    [[play](https://go.dev/play/p/FjIHV7lpxmW)]
-   **<big>Zip6</big>** : create a slice of Tuple6, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip6)]
    [[play](https://go.dev/play/p/oWPrnUYuFHo)]
-   **<big>Unzip6</big>** : create a group of slice from a slice of Tuple6.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip6)]
    [[play](https://go.dev/play/p/l41XFqCyh5E)]
-   **<big>Tuple7</big>** : represents a 7 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple7)]
    [[play](https://go.dev/play/p/dzAgv_Ezub9)]
-   **<big>Tuple7_Unbox</big>** : returns values in Tuple7.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple7_Unbox)]
    [[play](https://go.dev/play/p/R9I8qeDk0zs)]
-   **<big>Zip7</big>** : create a slice of Tuple7, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip7)]
    [[play](https://go.dev/play/p/WUJuo897Egf)]
-   **<big>Unzip7</big>** : create a group of slice from a slice of Tuple7.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip7)]
    [[play](https://go.dev/play/p/hws_P1Fr2j3)]
-   **<big>Tuple8</big>** : represents a 8 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple8)]
    [[play](https://go.dev/play/p/YA9S0rz3dRz)]
-   **<big>Tuple8_Unbox</big>** : returns values in Tuple8.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple8_Unbox)]
    [[play](https://go.dev/play/p/PRxLBBb4SMl)]
-   **<big>Zip8</big>** : create a slice of Tuple8, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip8)]
    [[play](https://go.dev/play/p/8V9jWkuJfaQ)]
-   **<big>Unzip8</big>** : create a group of slice from a slice of Tuple8.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip8)]
    [[play](https://go.dev/play/p/1SndOwGsZB4)]
-   **<big>Tuple9</big>** : represents a 9 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple9)]
    [[play](https://go.dev/play/p/yS2NGGtZpQr)]
-   **<big>Tuple9_Unbox</big>** : returns values in Tuple9.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple9_Unbox)]
    [[play](https://go.dev/play/p/oFJFGTAuOa8)]
-   **<big>Zip9</big>** : create a slice of Tuple9, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip9)]
    [[play](https://go.dev/play/p/cgsL15QYnfz)]
-   **<big>Unzip9</big>** : create a group of slice from a slice of Tuple9.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip9)]
    [[play](https://go.dev/play/p/91-BU_KURSA)]
-   **<big>Tuple10</big>** : represents a 10 elemnets tuple.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple10)]
    [[play](https://go.dev/play/p/799qqZg0hUv)]
-   **<big>Tuple10_Unbox</big>** : returns values in Tuple10.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Tuple10_Unbox)]
    [[play](https://go.dev/play/p/qfyx3x_X0Cu)]
-   **<big>Zip10</big>** : create a slice of Tuple10, whose elements are correspond to the given slice elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Zip10)]
    [[play](https://go.dev/play/p/YSR-2cXnrY4)]
-   **<big>Unzip10</big>** : create a group of slice from a slice of Tuple10.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/tuple.md#Unzip10)]
    [[play](https://go.dev/play/p/-taQB6Wfre_z)]

<h3 id="validator"> 24. Validator package contains some functions for data validation. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/validator"
```

#### Function list:

-   **<big>ContainChinese</big>** : check if the string contain mandarin chinese.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainChinese)]
    [[play](https://go.dev/play/p/7DpU0uElYeM)]
-   **<big>ContainLetter</big>** : check if the string contain at least one letter.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainLetter)]
    [[play](https://go.dev/play/p/lqFD04Yyewp)]
-   **<big>ContainLower</big>** : check if the string contain at least one lower case letter a-z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainLower)]
    [[play](https://go.dev/play/p/Srqi1ItvnAA)]
-   **<big>ContainUpper</big>** : check if the string contain at least one upper case letter A-Z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#ContainUpper)]
    [[play](https://go.dev/play/p/CmWeBEk27-z)]
-   **<big>IsAlpha</big>** : checks if the string contains only letters (a-zA-Z).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAlpha)]
    [[play](https://go.dev/play/p/7Q5sGOz2izQ)]
-   **<big>IsAllUpper</big>** : check if the string is all upper case letters A-Z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAllUpper)]
    [[play](https://go.dev/play/p/ZHctgeK1n4Z)]
-   **<big>IsAllLower</big>** : check if the string is all lower case letters a-z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAllLower)]
    [[play](https://go.dev/play/p/GjqCnOfV6cM)]
-   **<big>IsBase64</big>** : check if the string is base64 string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsBase64)]
    [[play](https://go.dev/play/p/sWHEySAt6hl)]
-   **<big>IsChineseMobile</big>** : check if the string is chinese mobile number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChineseMobile)]
    [[play](https://go.dev/play/p/GPYUlGTOqe3)]
-   **<big>IsChineseIdNum</big>** : check if the string is chinese id card.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChineseIdNum)]
    [[play](https://go.dev/play/p/d8EWhl2UGDF)]
-   **<big>IsChinesePhone</big>** : check if the string is chinese phone number.(xxx-xxxxxxxx or xxxx-xxxxxxx.)
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChinesePhone)]
    [[play](https://go.dev/play/p/RUD_-7YZJ3I)]
-   **<big>IsCreditCard</big>** : check if the string is credit card.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsCreditCard)]
    [[play](https://go.dev/play/p/sNwwL6B0-v4)]
-   **<big>IsDns</big>** : check if the string is dns.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsDns)]
    [[play](https://go.dev/play/p/jlYApVLLGTZ)]
-   **<big>IsEmail</big>** : check if the string is a email address.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsEmail)]
    [[play](https://go.dev/play/p/Os9VaFlT33G)]
-   **<big>IsEmptyString</big>** : check if the string is empty.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsEmptyString)]
    [[play](https://go.dev/play/p/dpzgUjFnBCX)]
-   **<big>IsFloat</big>** : check if the value is float(float32, float34) or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsFloat)]
    [[play](https://go.dev/play/p/vsyG-sxr99_Z)]
-   **<big>IsFloatStr</big>** : check if the string can convert to a float.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsFloatStr)]
    [[play](https://go.dev/play/p/LOYwS_Oyl7U)]
-   **<big>IsNumber</big>** : check if the value is number(integer, float) or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsNumber)]
    [[play](https://go.dev/play/p/mdJHOAvtsvF)]
-   **<big>IsNumberStr</big>** : check if the string can convert to a number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsNumberStr)]
    [[play](https://go.dev/play/p/LzaKocSV79u)]
-   **<big>IsJSON</big>** : check if the string is valid JSON.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsJSON)]
    [[play](https://go.dev/play/p/8Kip1Itjiil)]
-   **<big>IsRegexMatch</big>** : check if the string match the regexp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsRegexMatch)]
    [[play](https://go.dev/play/p/z_XeZo_litG)]
-   **<big>IsInt</big>** : check if the string can convert to a number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsInt)]
    [[play](https://go.dev/play/p/eFoIHbgzl-z)]
-   **<big>IsIntStr</big>** : check if the string can convert to a integer.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIntStr)]
    [[play](https://go.dev/play/p/jQRtFv-a0Rk)]
-   **<big>IsIp</big>** : check if the string is ip.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIp)]
    [[play](https://go.dev/play/p/FgcplDvmxoD)]
-   **<big>IsIpV4</big>** : check if the string is ipv4.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIpV4)]
    [[play](https://go.dev/play/p/zBGT99EjaIu)]
-   **<big>IsIpV6</big>** : check if the string is ipv6.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsIpV6)]
    [[play](https://go.dev/play/p/AHA0r0AzIdC)]
-   **<big>IsStrongPassword</big>** : check if the string is strong password.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsStrongPassword)]
    [[play](https://go.dev/play/p/QHdVcSQ3uDg)]
-   **<big>IsUrl</big>** : check if the string is url.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsUrl)]
    [[play](https://go.dev/play/p/pbJGa7F98Ka)]
-   **<big>IsWeakPassword</big>** : check if the string is weak password.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsWeakPassword)]
    [[play](https://go.dev/play/p/wqakscZH5gH)]
-   **<big>IsZeroValue</big>** : check if value is a zero value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsZeroValue)]
    [[play](https://go.dev/play/p/UMrwaDCi_t4)]
-   **<big>IsGBK</big>** : check if data encoding is gbk.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsGBK)]
    [[play](https://go.dev/play/p/E2nt3unlmzP)]
-   **<big>IsASCII</big>** : checks if string is all ASCII char.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsASCII)]
    [[play](https://go.dev/play/p/hfQNPLX0jNa)]
-   **<big>IsPrintable</big>** : checks if string is all printable chars.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsPrintable)]
    [[play](https://go.dev/play/p/Pe1FE2gdtTP)]
-   **<big>IsBin</big>** : check if a give string is a valid binary value or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsBin)]
    [[play](https://go.dev/play/p/ogPeg2XJH4P)]
-   **<big>IsHex</big>** : check if a give string is a valid hexadecimal value or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsHex)]
    [[play](https://go.dev/play/p/M2qpHbEwmm7)]
-   **<big>IsBase64URL</big>** : check if a give string is a valid URL-safe Base64 encoded string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsBase64URL)]
    [[play](https://go.dev/play/p/vhl4mr8GZ6S)]
-   **<big>IsJWT</big>** : check if a give string is a valid JSON Web Token (JWT).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsJWT)]
    [[play](https://go.dev/play/p/R6Op7heJbKI)]
-   **<big>IsVisa</big>** : check if a give string is a valid visa card nubmer or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsVisa)]
    [[play](https://go.dev/play/p/SdS2keOyJsl)]
-   **<big>IsMasterCard</big>** : check if a give string is a valid master card nubmer or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsMasterCard)]
    [[play](https://go.dev/play/p/CwWBFRrG27b)]
-   **<big>IsAmericanExpress</big>** : check if a give string is a valid american expression card nubmer or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsAmericanExpress)]
    [[play](https://go.dev/play/p/HIDFpcOdpkd)]
-   **<big>IsUnionPay</big>** : check if a give string is a valid union pay nubmer or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsUnionPay)]
    [[play](https://go.dev/play/p/CUHPEwEITDf)]
-   **<big>IsChinaUnionPay</big>** : check if a give string is a valid china union pay nubmer or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/validator.md#IsChinaUnionPay)]
    [[play](https://go.dev/play/p/yafpdxLiymu)]

<h3 id="xerror"> 25. Xerror package implements helpers for errors. &nbsp; &nbsp; &nbsp; &nbsp;<a href="#index">index</a></h3>

```go
import "github.com/duke-git/lancet/v2/xerror"
```

#### Function list:

-   **<big>New</big>** : creates a new XError pointer instance with message.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#New)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>Wrap</big>** : creates a new XError pointer instance based on error object, and add message.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#Wrap)]
    [[play](https://go.dev/play/p/5385qT2dCi4)]
-   **<big>Unwrap</big>** : returns unwrapped XError from err by errors.As. If no XError, returns nil.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#Unwrap)]
    [[play](https://go.dev/play/p/LKMLep723tu)]
-   **<big>XError_Wrap</big>** : creates a new XError and copy message and id to new one.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Wrap)]
    [[play](https://go.dev/play/p/RpjJ5u5sc97)]
-   **<big>XError_Unwrap</big>** : Compatible with github.com/pkg/errors.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Unwrap)]
    [[play](https://go.dev/play/p/VUXJ8BST4c6)]
-   **<big>XError_With</big>** : adds key and value related to the XError object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_With)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_Id</big>** : sets XError object id to check equality in XError.Is.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Id)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Is</big>** : checks if target error is XError and Error.id of two errors are matched.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Is)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Values</big>** : returns map of key and value that is set by XError.With function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Values)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_StackTrace</big>** : returns stack trace which is compatible with pkg/errors.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_StackTrace)]
    [[play](https://go.dev/play/p/6FAvSQpa7pc)]
-   **<big>XError_Info</big>** : returns information of xerror, which can be printed.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Info)]
    [[play](https://go.dev/play/p/1ZX0ME1F-Jb)]
-   **<big>XError_Error</big>** : implements standard error interface.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#XError_Error)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>TryUnwrap</big>** : check if err is nil then it returns a valid value. If err is not nil, TryUnwrap panics with err.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/en/api/packages/xerror.md#TryUnwrap)]
    [[play](https://go.dev/play/p/acyZVkNZEeW)]

## How to Contribute

#### [Contributing Guide](./CONTRIBUTING.en-US.md)


## Sponsor

Hello, I am a software developer and have been engaged in development work for 13 years. Love open source software. And be willing to put in the energy for it. I am the creator of project lancet. Since Lancet was released as open source two years ago, it has been used by more than 1,000 internal and external projects. lancet will always be free for all users. Your support is a powerful encouragement for me to continue my struggle. Thanks! You can use WeChat to scans the following QR code or clicks the following sponsor button to initiate sponsorship.

<div style="position: relative;display: inline-block;">
    <img src="./docs/public/wechat_pay.png" width="260" height="260"/>
    <a href="https://en.liberapay.com/Duke_Du/donate" target="\_blank"><img src="./docs/public/sponsor_btn.png" width="220" height="60"/></a>
</div>

<br/>
<br/>

*Donated funds will be used to maintain [lancet](https://www.golancet.cn/en/) website and pay for cloud server costs. Or just buy me a cup of ☕️ when I'm sleepy writing code.*


## Contributors
Thank you to all the people who contributed to lancet!

<a href="https://github.com/duke-git/lancet/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=duke-git/lancet" />
</a>