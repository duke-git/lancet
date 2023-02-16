<div align=center>
<img src="./logo.png" width="200" height="200"/>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-2.1.15-green.svg)](https://github.com/duke-git/lancet/releases)
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

-   👏 Comprehensive, efficient and reusable.
-   💪 400+ go util functions, support string, slice, datetime, net, crypt...
-   💅 Only depend on the go standard library.
-   🌍 Unit test for every exported function.

## Installation

### Note:

1. <b>For users who use go1.18 and above, it is recommended to install lancet v2.x.x. Cause in v2.x.x all functions was rewriten with generics of go1.18.</b>

```go
go get github.com/duke-git/lancet/v2 // will install latest version of v2.x.x
```

2. <b>For users who use version below go1.18, you should install v1.x.x. The latest of v1.x.x is v1.3.6. </b>

```go
go get github.com/duke-git/lancet@v1.3.6 // below go1.18, install latest version of v1.x.x
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

### 1. Algorithm package implements some basic algorithm. eg. sort, search.

```go
import "github.com/duke-git/lancet/v2/algorithm"
```

#### Function list:

-   **<big>BubbleSort</big>** : sorts slice with bubble sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BubbleSort)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>CountSort</big>** : sorts slice with bubble sort algorithm, don't change original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#CountSort)]
    [[play](https://go.dev/play/p/tB-Umgm0DrP)]
-   **<big>HeapSort</big>** : sorts slice with heap sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#HeapSort)]
    [[play](https://go.dev/play/p/u6Iwa1VZS_f)]
-   **<big>InsertionSort</big>** : sorts slice with insertion sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#InsertionSort)]
    [[play](https://go.dev/play/p/G5LJiWgJJW6)]
-   **<big>MergeSort</big>** : sorts slice with merge sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#MergeSort)]
    [[play](https://go.dev/play/p/ydinn9YzUJn)]
-   **<big>QuickSort</big>** : sorts slice with quick sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#QuickSort)]
    [[play](https://go.dev/play/p/7Y7c1Elk3ax)]
-   **<big>SelectionSort</big>** : sorts slice with selection sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#SelectionSort)]
    [[play](https://go.dev/play/p/oXovbkekayS)]
-   **<big>ShellSort</big>** : sorts slice with shell sort algorithm, will change the original slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#ShellSort)]
    [[play](https://go.dev/play/p/3ibkszpJEu3)]
-   **<big>BinarySearch</big>** : returns the index of target within a sorted slice, use binary search (recursive call itself).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BinarySearch)]
    [[play](https://go.dev/play/p/t6MeGiUSN47)]
-   **<big>BinaryIterativeSearch</big>** : returns the index of target within a sorted slice, use binary search (no recursive).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#BinaryIterativeSearch)]
    [[play](https://go.dev/play/p/Anozfr8ZLH3)]
-   **<big>LinearSearch</big>** : returns the index of target in slice base on equal function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#LinearSearch)]
    [[play](https://go.dev/play/p/IsS7rgn5s3x)]
-   **<big>LRUCache</big>** : implements memory cache with lru algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/algorithm.md#LRUCache)]
    [[play](https://go.dev/play/p/-EZjgOURufP)]

### 2. Concurrency package contain some functions to support concurrent programming. eg, goroutine, channel, async.

```go
import "github.com/duke-git/lancet/v2/concurrency"
```

#### Function list:

-   **<big>NewChannel</big>** : create a Channel pointer instance.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#NewChannel)]
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Bridge</big>** : link multiply channels into one channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/Bridge.md#NewChannel)]
    [[play](https://go.dev/play/p/qmWSy1NVF-Y)]
-   **<big>FanIn</big>** : merge multiple channels into one channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#FanIn)]
    [[play](https://go.dev/play/p/2VYFMexEvTm)]
-   **<big>Generate</big>** : creates a channel, then put values into the channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Generate)]
    [[play](https://go.dev/play/p/7aB4KyMMp9A)]
-   **<big>Or</big>** : read one or more channels into one channel, will close when any readin channel is closed.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Or)]
    [[play](https://go.dev/play/p/Wqz9rwioPww)]
-   **<big>OrDone</big>** : read a channel into another channel, will close until cancel context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#OrDone)]
    [[play](https://go.dev/play/p/lm_GoS6aDjo)]
-   **<big>Repeat</big>** : create channel, put values into the channel repeatly until cancel the context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Repeat)]
    [[play](https://go.dev/play/p/k5N_ALVmYjE)]
-   **<big>RepeatFn</big>** : create a channel, excutes fn repeatly, and put the result into the channel, until close context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#RepeatFn)]
    [[play](https://go.dev/play/p/4J1zAWttP85)]
-   **<big>Take</big>** : create a channel whose values are taken from another channel with limit number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Take)]
    [[play](https://go.dev/play/p/9Utt-1pDr2J)]
-   **<big>Tee</big>** : split one chanel into two channels, until cancel the context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/concurrency.md#Tee)]
    [[play](https://go.dev/play/p/3TQPKnCirrP)]

### 3. Condition package contains some functions for conditional judgment. eg. And, Or, TernaryOperator...

```go
import "github.com/duke-git/lancet/v2/condition"
```

#### Function list:

-   **<big>Bool</big>** : returns the truthy value of anything.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Bool)]
    [[play](https://go.dev/play/p/ETzeDJRSvhm)]
-   **<big>And</big>** : returns true if both a and b are truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#And)]
    [[play](https://go.dev/play/p/W1SSUmt6pvr)]
-   **<big>Or</big>** : returns false if neither a nor b is truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Or)]
[[play](https://go.dev/play/p/UlQTxHaeEkq)]]
-   **<big>Xor</big>** : returns true if a or b but not both is truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Xor)]
    [[play](https://go.dev/play/p/gObZrW7ZbG8)]
-   **<big>Nor</big>** : returns true if neither a nor b is truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Nor)]
    [[play](https://go.dev/play/p/g2j08F_zZky)
-   **<big>Xnor</big>** : returns true if both a and b or neither a nor b are truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Xnor)]
    [[play](https://go.dev/play/p/OuDB9g51643)]
-   **<big>Nand</big>** : returns false if both a and b are truthy.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#Nand)]
    [[play](https://go.dev/play/p/vSRMLxLIbq8)]
-   **<big>TernaryOperator</big>** : ternary operator.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/condition.md#TernaryOperator)]
    [[play](https://go.dev/play/p/ElllPZY0guT)]

### 4. Convertor package contains some functions for data convertion.

```go
import "github.com/duke-git/lancet/v2/convertor"
```

#### Function list:

-   **<big>ColorHexToRGB</big>** : convert color hex to color rgb.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ColorHexToRGB)]
    [[play](https://go.dev/play/p/o7_ft-JCJBV)]
-   **<big>ColorRGBToHex</big>** : convert rgb color to hex color.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ColorRGBToHex)]
    [[play](https://go.dev/play/p/nzKS2Ro87J1)]
-   **<big>ToBool</big>** : convert string to bool.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToBool)]
    [[play](https://go.dev/play/p/ARht2WnGdIN)]
-   **<big>ToBytes</big>** : convert value to byte slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToBytes)]
    [[play](https://go.dev/play/p/fAMXYFDvOvr)]
-   **<big>ToChar</big>** : convert string to char slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToChar)]
    [[play](https://go.dev/play/p/JJ1SvbFkVdM)]
-   **<big>ToChannel</big>** : convert a collection of elements to a read-only channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToChannel)]
    [[play](https://go.dev/play/p/hOx_oYZbAnL)]
-   **<big>ToFloat</big>** : convert value to float64, if param is a invalid floatable, will return 0.0 and error.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToFloat)]
    [[play](https://go.dev/play/p/4YTmPCibqHJ)]
-   **<big>ToInt</big>** : convert value to int64 value, if input is not numerical, return 0 and error.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToInt)]
    [[play](https://go.dev/play/p/9_h9vIt-QZ_b)]
-   **<big>ToJson</big>** : convert value to a json string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToJson)]
    [[play](https://go.dev/play/p/2rLIkMmXWvR)]
-   **<big>ToMap</big>** : convert a slice of structs to a map based on iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToMap)]
    [[play](https://go.dev/play/p/tVFy7E-t24l)]
-   **<big>ToPointer</big>** : return a pointer of passed value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToPointer)]
    [[play](https://go.dev/play/p/ASf_etHNlw1)]
-   **<big>ToString</big>** : convert value to string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#ToString)]
    [[play](https://go.dev/play/p/nF1zOOslpQq)]
-   **<big>StructToMap</big>** : convert struct to map, only convert exported struct field.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#StructToMap)]
    [[play](https://go.dev/play/p/KYGYJqNUBOI)]
-   **<big>MapToSlice</big>** : convert map to slice based on iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#MapToSlice)]
    [[play](https://go.dev/play/p/dmX4Ix5V6Wl)]
-   **<big>EncodeByte</big>** : encode data to byte slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#EncodeByte)]
    [[play](https://go.dev/play/p/DVmM1G5JfuP)]
-   **<big>DecodeByte</big>** : decode byte slice data to target object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#DecodeByte)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>DeepClone</big>** : creates a deep copy of passed item, can't clone unexported field of struct.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/convertor.md#DeepClone)]
    [[play](https://go.dev/play/p/j4DP5dquxnk)]


### 5. Cryptor package is for data encryption and decryption.

```go
import "github.com/duke-git/lancet/v2/cryptor"
```

#### Function list:

-   **<big>AesEcbEncrypt</big>** : encrypt byte slice data with key use AES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesEcbEncrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesEcbDecrypt</big>** : decrypt byte slice data with key use AES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesEcbDecrypt)]
    [[play](https://go.dev/play/p/zI6xsmuQRbn)]
-   **<big>AesCbcEncrypt</big>** : encrypt byte slice data with key use AES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCbcEncrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCbcDecrypt</big>** : decrypt byte slice data with key use AES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCbcDecrypt)]
    [[play](https://go.dev/play/p/IOq_g8_lKZD)]
-   **<big>AesCtrCrypt</big>** : encrypt/ decrypt byte slice data with key use AES CRC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCtrCrypt)]
    [[play](https://go.dev/play/p/SpaZO0-5Nsp)]
-   **<big>AesCfbEncrypt</big>** : encrypt byte slice data with key use AES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCfbEncrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesCfbDecrypt</big>** : decrypt byte slice data with key use AES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesCfbDecrypt)]
    [[play](https://go.dev/play/p/tfkF10B13kH)]
-   **<big>AesOfbEncrypt</big>** : encrypt byte slice data with key use AES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesOfbEncrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>AesOfbDecrypt</big>** : decrypt byte slice data with key use AES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#AesOfbDecrypt)]
    [[play](https://go.dev/play/p/VtHxtkUj-3F)]
-   **<big>Base64StdEncode</big>** : encode string with base64 encoding.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Base64StdEncode)]
    [[play](https://go.dev/play/p/VOaUyQUreoK)]
-   **<big>Base64StdDecode</big>** : decode string with base64 encoding.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Base64StdDecode)]
    [[play](https://go.dev/play/p/RWQylnJVgIe)]
-   **<big>DesEcbEncrypt</big>** : encrypt byte slice data with key use DES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesEcbEncrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesEcbDecrypt</big>** : decrypt byte slice data with key use DES ECB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesEcbDecrypt)]
    [[play](https://go.dev/play/p/8qivmPeZy4P)]
-   **<big>DesCbcEncrypt</big>** : encrypt byte slice data with key use DES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCbcEncrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCbcDecrypt</big>** : decrypt byte slice data with key use DES CBC algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCbcDecrypt)]
    [[play](https://go.dev/play/p/4cC4QvWfe3_1)]
-   **<big>DesCtrCrypt</big>** : encrypt/decrypt byte slice data with key use DES CRY algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCtrCrypt)]
    [[play](https://go.dev/play/p/9-T6OjKpcdw)]
-   **<big>DesCfbEncrypt</big>** : encrypt byte slice data with key use DES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCfbEncrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesCfbDecrypt</big>** : decrypt byte slice data with key use DES CFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesCfbDecrypt)]
    [[play](https://go.dev/play/p/y-eNxcFBlxL)]
-   **<big>DesOfbEncrypt</big>** : encrypt byte slice data with key use DES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesOfbEncrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>DesOfbDecrypt</big>** : decrypt byte slice data with key use DES OFB algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#DesOfbDecrypt)]
    [[play](https://go.dev/play/p/74KmNadjN1J)]
-   **<big>HmacMd5</big>** : return the md5 hmac hash of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacMd5)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>HmacSha1</big>** : return the hmac hash of string use sha1.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha1)]
    [[play](https://go.dev/play/p/1UI4oQ4WXKM)]
-   **<big>HmacSha256</big>** : return the hmac hash of string use sha256.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha256)]
    [[play](https://go.dev/play/p/HhpwXxFhhC0)]
-   **<big>HmacSha512</big>** : return the hmac hash of string use sha512.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#HmacSha512)]
    [[play](https://go.dev/play/p/59Od6m4A0Ud)]
-   **<big>Md5String</big>** : return the md5 value of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Md5String)]
    [[play](https://go.dev/play/p/1bLcVetbTOI)]
-   **<big>Md5File</big>** : return the md5 value of file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Md5File)]
-   **<big>Sha1</big>** : return the sha1 value (SHA-1 hash algorithm) of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha1)]
    [[play](https://go.dev/play/p/_m_uoD1deMT)]
-   **<big>Sha256</big>** : return the sha256 value (SHA-256 hash algorithm) of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha256)]
    [[play](https://go.dev/play/p/tU9tfBMIAr1)]
-   **<big>Sha512</big>** : return the sha512 value (SHA-512 hash algorithm) of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#Sha512)]
    [[play](https://go.dev/play/p/3WsvLYZxsHa)]
-   **<big>GenerateRsaKey</big>** : create rsa private and public pemo file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#GenerateRsaKey)]
    [[play](https://go.dev/play/p/zutRHrDqs0X)]
-   **<big>RsaEncrypt</big>** : encrypt data with ras algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#RsaEncrypt)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]
-   **<big>RsaDecrypt</big>** : decrypt data with ras algorithm.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/cryptor.md#RsaDecrypt)]
    [[play](https://go.dev/play/p/uef0q1fz53I)]

### 6. Datetime package supports date and time format and compare.

```go
import "github.com/duke-git/lancet/v2/datetime"
```

#### Function list:

-   **<big>AddDay</big>** : add or sub day to the time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddDay)]
    [[play](https://go.dev/play/p/dIGbs_uTdFa)]
-   **<big>AddHour</big>** : add or sub day to the time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddHour)]
    [[play](https://go.dev/play/p/rcMjd7OCsi5)]
-   **<big>AddMinute</big>** : add or sub day to the time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#AddMinute)]
    [[play](https://go.dev/play/p/nT1heB1KUUK)]
-   **<big>BeginOfMinute</big>** : return the date time at the begin of minute of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMinute)]
    [[play](https://go.dev/play/p/ieOLVJ9CiFT)]
-   **<big>BeginOfHour</big>** : return the date time at the begin of hour of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfHour)]
    [[play](https://go.dev/play/p/GhdGFnDWpYs)]
-   **<big>BeginOfDay</big>** : return the date time at the begin of day of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfDay)]
    [[play](https://go.dev/play/p/94m_UT6cWs9)]
-   **<big>BeginOfWeek</big>** : return the date time at the begin of week of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfWeek)]
    [[play](https://go.dev/play/p/ynjoJPz7VNV)]
-   **<big>BeginOfMonth</big>** : return the date time at the begin of month of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfMonth)]
    [[play](https://go.dev/play/p/bWXVFsmmzwL)]
-   **<big>BeginOfYear</big>** : return the date time at the begin of year of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#BeginOfYear)]
    [[play](https://go.dev/play/p/i326DSwLnV8)]
-   **<big>EndOfMinute</big>** : return the date time at the end of minute of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMinute)]
    [[play](https://go.dev/play/p/yrL5wGzPj4z)]
-   **<big>EndOfHour</big>** : return the date time at the end of hour of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfHour)]
    [[play](https://go.dev/play/p/6ce3j_6cVqN)]
-   **<big>EndOfDay</big>** : return the date time at the end of day of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfDay)]
    [[play](https://go.dev/play/p/eMBOvmq5Ih1)]
-   **<big>EndOfWeek</big>** : return the date time at the end of week of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfWeek)]
    [[play](https://go.dev/play/p/i08qKXD9flf)]
-   **<big>EndOfMonth</big>** : return the date time at the end of month of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfMonth)]
    [[play](https://go.dev/play/p/_GWh10B3Nqi)]
-   **<big>EndOfYear</big>** : return the date time at the end of year of specific date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#EndOfYear)]
    [[play](https://go.dev/play/p/G01cKlMCvNm)]
-   **<big>GetNowDate</big>** : return format yyyy-mm-dd of current date.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowDate)]
    [[play](https://go.dev/play/p/PvfkPpcpBBf)]
-   **<big>GetNowTime</big>** : return format hh-mm-ss of current time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowTime)]
    [[play](https://go.dev/play/p/l7BNxCkTmJS)]
-   **<big>GetNowDateTime</big>** : return format yyyy-mm-dd hh-mm-ss of current datetime.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNowDateTime)]
    [[play](https://go.dev/play/p/pI4AqngD0al)]
-   **<big>GetZeroHourTimestamp</big>** : return timestamp of zero hour (timestamp of 00:00).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetZeroHourTimestamp)]
    [[play](https://go.dev/play/p/QmL2oIaGE3q)]
-   **<big>GetNightTimestamp</big>** : return timestamp of zero hour (timestamp of 23:59).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#GetNightTimestamp)]
    [[play](https://go.dev/play/p/UolysR3MYP1)]
-   **<big>FormatTimeToStr</big>** : convert time to string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#FormatTimeToStr)]
    [[play](https://go.dev/play/p/_Ia7M8H_OvE)]
-   **<big>FormatStrToTime</big>** : convert string to time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#FormatStrToTime)]
    [[play](https://go.dev/play/p/1h9FwdU8ql4)]
-   **<big>NewUnix</big>** : return unix timestamp of specific time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewUnix)]
    [[play](https://go.dev/play/p/psoSuh_kLRt)]
-   **<big>NewUnixNow</big>** : return unix timestamp of current time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewUnixNow)]
    [[play](https://go.dev/play/p/U4PPx-9D0oz)]
-   **<big>NewFormat</big>** : return unix timestamp of specific time string, t should be "yyyy-mm-dd hh:mm:ss".
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>NewISO8601</big>** : return unix timestamp of specific iso8601 time string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#NewISO8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]
-   **<big>ToUnix</big>** : return unix timestamp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToUnix)]
    [[play](https://go.dev/play/p/_LUiwAdocjy)]
-   **<big>ToFormat</big>** : return the time string 'yyyy-mm-dd hh:mm:ss' of unix time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToFormat)]
    [[play](https://go.dev/play/p/VkW08ZOaXPZ)]
-   **<big>ToFormatForTpl</big>** : return the time string which format is specific tpl.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToFormatForTpl)]
    [[play](https://go.dev/play/p/nyXxXcQJ8L5)]
-   **<big>ToIso8601</big>** : return iso8601 time string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datetime.md#ToIso8601)]
    [[play](https://go.dev/play/p/mkhOHQkdeA2)]




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

#### Structure list:

-   **<big>List</big>** : a linear table, implemented with slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/list.md)]
-   **<big>Link</big>** : link list structure, contains singly link and doubly link.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/link.md)]
-   **<big>Stack</big>** : stack structure(fifo), contains array stack and link stack.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/stack.md)]
-   **<big>Queue</big>** : queue structure(filo), contains array queue, circular queue, link queue and priority queue.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/queue.md)]
-   **<big>Set</big>** : a data container, like slice, but element of set is not duplicate.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/set.md)]
-   **<big>Tree</big>** : binary search tree structure.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/tree.md)]
-   **<big>Heap</big>** : a binary max heap.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/heap.md)]
-   **<big>Hashmap</big>** : hash map structure.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/datastructure/hashmap.md)]

### 8. Fileutil package implements some basic functions for file operations.

```go
import "github.com/duke-git/lancet/v2/fileutil"
```

#### Function list：

-   **<big>ClearFile</big>** : write empty string to target file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ClearFile)]
    [[play](https://go.dev/play/p/NRZ0ZT-G94H)]
-   **<big>CreateFile</big>** : create file in path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CreateFile)]
    [[play](https://go.dev/play/p/lDt8PEsTNKI)]
-   **<big>CreateDir</big>** : create directory in absolute path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CreateDir)]
    [[play](https://go.dev/play/p/qUuCe1OGQnM)]
-   **<big>CopyFile</big>** :copy src file to dest file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#CopyFile)]
    [[play](https://go.dev/play/p/Jg9AMJMLrJi)]
-   **<big>FileMode</big>** : return file's mode and permission.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#FileMode)]
    [[play](https://go.dev/play/p/2l2hI42fA3p)]
-   **<big>MiMeType</big>** : return file mime type.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#MiMeType)]
    [[play](https://go.dev/play/p/bd5sevSUZNu)]
-   **<big>IsExist</big>** : checks if a file or directory exists.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsExist)]
    [[play](https://go.dev/play/p/nKKXt8ZQbmh)]
-   **<big>IsLink</big>** : checks if a file is symbol link or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsLink)]
    [[play](https://go.dev/play/p/TL-b-Kzvf44)]
-   **<big>IsDir</big>** : checks if the path is directory or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#IsDir)]
    [[play](https://go.dev/play/p/WkVwEKqtOWk)]
-   **<big>ListFileNames</big>** : return all file names in the path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ListFileNames)]
    [[play](https://go.dev/play/p/Tjd7Y07rejl)]
-   **<big>RemoveFile</big>** : remove file, param should be file path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#RemoveFile)]
    [[play](https://go.dev/play/p/P2y0XW8a1SH)]
-   **<big>ReadFileToString</big>** : return string of file content.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ReadFileToString)]
    [[play](https://go.dev/play/p/cmfwp_5SQTp)]
-   **<big>ReadFileByLine</big>** : read file line by line, return string slice of file content.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#ReadFileByLine)]
    [[play](https://go.dev/play/p/svJP_7ZrBrD)]
-   **<big>Zip</big>** : create zip file.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#Zip)]
    [[play](https://go.dev/play/p/j-3sWBp8ik_P)]
-   **<big>UnZip</big>** : unzip the zip file and save it to dest path.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/fileutil.md#UnZip)]
    [[play](https://go.dev/play/p/g0w34kS7B8m)]

### 9. Formatter contains some functions for data formatting.

```go
import "github.com/duke-git/lancet/v2/formatter"
```

#### Function list:

-   **<big>Comma</big>** : add comma to a number value by every 3 numbers from right, ahead by symbol char.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/formatter.md#Comma)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]

### 10. Function package can control the flow of function execution and support part of functional programming

```go
import "github.com/duke-git/lancet/v2/function"
```

#### Function list:

-   **<big>After</big>** : return a function that invokes passed funcation once the returned function is called more than n times.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#After)]
    [[play](https://go.dev/play/p/eRD5k2vzUVX)]
-   **<big>Before</big>** : return a function that invokes passed funcation once the returned function is called less than n times
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#Before)]
    [[play](https://go.dev/play/p/0HqUDIFZ3IL)]
-   **<big>CurryFn</big>** : make a curry function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#CurryFn)]
    [[play](https://go.dev/play/p/5HopfDwANKX)]
-   **<big>Compose</big>** : compose the functions from right to left.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#Compose)]
    [[play](https://go.dev/play/p/KKfugD4PKYF)]
-   **<big>Delay</big>** : call the function after delayed time.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#Delay)]
    [[play](https://go.dev/play/p/Ivtc2ZE-Tye)]
-   **<big>Debounced</big>** : creates a debounced function that delays invoking fn until after wait duration have elapsed since the last time the debounced function was invoked.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#Debounced)]
    [[play](https://go.dev/play/p/absuEGB_GN7)]
-   **<big>Schedule</big>** : invoke function every duration time, util close the returned bool channel.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#Schedule)]
    [[play](https://go.dev/play/p/hbON-Xeyn5N)]
-   **<big>Pipeline</big>** : takes a list of functions and returns a function whose param will be passed into the functions one by one.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#Pipeline)]
    [[play](https://go.dev/play/p/mPdUVvj6HD6)]
-   **<big>Watcher</big>** : Watcher is used for record code excution time. can start/stop/reset the watch timer. get the elapsed time of function execution.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/function.md#Watcher)]
    [[play](https://go.dev/play/p/l2yrOpCLd1I)]


### 11. Maputil package includes some functions to manipulate map.

```go
import "github.com/duke-git/lancet/v2/maputil"
```

#### Function list:

-   **<big>ForEach</big>** : executes iteratee funcation for every key and value pair in map.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#ForEach)]
    [[play](https://go.dev/play/p/OaThj6iNVXK)]
-   **<big>Filter</big>** : iterates over map, return a new map contains all key and value pairs pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Filter)]
    [[play](https://go.dev/play/p/fSvF3wxuNG7)]
-   **<big>Intersect</big>** : iterates over maps, return a new map of key and value pairs in all given maps.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Intersect)]
    [[play](https://go.dev/play/p/Zld0oj3sjcC)]
-   **<big>Keys</big>** : returns a slice of the map's keys.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Keys)]
    [[play](https://go.dev/play/p/xNB5bTb97Wd)]
-   **<big>Merge</big>** : merge maps, next key will overwrite previous key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Merge)]
    [[play](https://go.dev/play/p/H95LENF1uB-)]
-   **<big>Minus</big>** : creates a map of whose key in mapA but not in mapB.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Minus)]
    [[play](https://go.dev/play/p/3u5U9K7YZ9m)]
-   **<big>Values</big>** : returns a slice of the map's values.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#Values)]
    [[play](https://go.dev/play/p/CBKdUc5FTW6)]
-   **<big>IsDisjoint</big>** : check two map are disjoint if they have no keys in common.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/maputil.md#IsDisjoint)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]

### 12. Mathutil package implements some functions for math calculation.

```go
import "github.com/duke-git/lancet/v2/mathutil"
```

#### Function list:

-   **<big>Average</big>** :return average value of numbers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Average)]
    [[play](https://go.dev/play/p/Vv7LBwER-pz)]
-   **<big>Exponent</big>** : calculate x^n for int64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Exponent)]
    [[play](https://go.dev/play/p/uF3HGNPk8wr)]
-   **<big>Fibonacci</big>** :calculate fibonacci number before n for int.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Fibonacci)]
    [[play](https://go.dev/play/p/IscseUNMuUc)]
-   **<big>Factorial</big>** : calculate x! for uint.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Factorial)]
    [[play](https://go.dev/play/p/tt6LdOK67Nx)]
-   **<big>Max</big>** : return maximum value of numbers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Max)]
    [[play](https://go.dev/play/p/cN8DHI0rTkH)]
-   **<big>MaxBy</big>** : return the maximum value of a slice using the given comparator function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#MaxBy)]
    [[play](https://go.dev/play/p/pbe2MT-7DV2)]
-   **<big>Min</big>** : return minimum value of numbers.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Min)]
    [[play](https://go.dev/play/p/21BER_mlGUj)]
-   **<big>MinBy</big>** : return the minimum value of a slice using the given comparator function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#MinBy)]
    [[play](https://go.dev/play/p/N9qgYg_Ho6f)]
-   **<big>Percent</big>** : calculate the percentage of value to total.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#Percent)]
    [[play](https://go.dev/play/p/QQM9B13coSP)]
-   **<big>RoundToFloat</big>** : round up to n decimal places for float64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#RoundToFloat)]
    [[play](https://go.dev/play/p/ghyb528JRJL)]
-   **<big>RoundToString</big>** : round up to n decimal places for float64, return string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#RoundToString)]
    [[play](https://go.dev/play/p/kZwpBRAcllO)]
-   **<big>TruncRound</big>** : round off n decimal places for int64.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/mathutil.md#TruncRound)]
    [[play](https://go.dev/play/p/aumarSHIGzP)]

### 13. Netutil package contains functions to get net information and send http request.

```go
import "github.com/duke-git/lancet/v2/netutil"
```

#### Function list:

-   **<big>ConvertMapToQueryString</big>** : convert map to sorted url query string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#ConvertMapToQueryString)]
    [[play](https://go.dev/play/p/jnNt_qoSnRi)]
-   **<big>EncodeUrl</big>** : encode url(?a=1&b=[2] -> ?a=1&b=%5B2%5D).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#EncodeUrl)]
    [[play](https://go.dev/play/p/bsZ6BRC4uKI)]
-   **<big>GetInternalIp</big>** : return internal ipv4.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetInternalIp)]
    [[play](https://go.dev/play/p/5mbu-gFp7ei)]
-   **<big>GetIps</big>** : return all ipv4 of current system.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetIps)]
    [[play](https://go.dev/play/p/NUFfcEmukx1)]
-   **<big>GetMacAddrs</big>** : return mac address of current system.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetMacAddrs)]
    [[play](https://go.dev/play/p/Rq9UUBS_Xp1)]
-   **<big>GetPublicIpInfo</big>** : return [public ip information](http://ip-api.com/json/).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetPublicIpInfo)]
    [[play](https://go.dev/play/p/YDxIfozsRHR)]
-   **<big>GetRequestPublicIp</big>** : return the http request public ip.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#GetRequestPublicIp)]
    [[play](https://go.dev/play/p/kxU-YDc_eBo)]
-   **<big>IsPublicIP</big>** : verify a ip is public or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#IsPublicIP)]
    [[play](https://go.dev/play/p/nmktSQpJZnn)]
-   **<big>IsInternalIP</big>** : verify an ip is intranet or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#IsInternalIP)]
    [[play](https://go.dev/play/p/sYGhXbgO4Cb)]
-   **<big>HttpRequest</big>** : a composed http request used for HttpClient send request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>HttpClient</big>** : a http client tool, used for sending http request
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpClient)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>SendRequest</big>** : send http request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#SendRequest)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>DecodeResponse</big>** : decode http response into target object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#DecodeResponse)]
    [[play](https://go.dev/play/p/jUSgynekH7G)]
-   **<big>StructToUrlValues</big>** : convert struct to url valuse.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#StructToUrlValues)]
    [[play](https://go.dev/play/p/pFqMkM40w9z)]
-   **<big>HttpGet<sup>deprecated</sup></big>** : send get http request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpGet)]
-   **<big>HttpDelete<sup>deprecated</sup></big>** : send delete http request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpDelete)]
-   **<big>HttpPost<sup>deprecated</sup></big>** : send post http request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPost)]
-   **<big>HttpPut<sup>deprecated</sup></big>** : send put http request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPut)]
-   **<big>HttpPatch<sup>deprecated</sup></big>** : send patch http request.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#HttpPatch)]
-   **<big>ParseHttpResponse</big>** : decode http response into target object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/netutil.md#ParseHttpResponse)]



### 14. Random package implements some basic functions to generate random int and string.

```go
import "github.com/duke-git/lancet/v2/random"
```

#### Function list:

-   **<big>RandBytes</big>** : generate random byte slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandBytes)]
    [[play](https://go.dev/play/p/EkiLESeXf8d)]
-   **<big>RandInt</big>** : generate random int number between min and max.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandInt)]
    [[play](https://go.dev/play/p/pXyyAAI5YxD)]
-   **<big>RandString</big>** : generate random string of specific length.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandString)]
    [[play](https://go.dev/play/p/W2xvRUXA7Mi)]
-   **<big>RandUpper</big>** : generate a random upper case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandUpper)]
    [[play](https://go.dev/play/p/29QfOh0DVuh)]
-   **<big>RandLower</big>** : generate a random lower case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandLower)]
    [[play](https://go.dev/play/p/XJtZ471cmtI)]
-   **<big>RandNumeral</big>** : generate a random numeral string of specific length.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandNumeral)]
    [[play](https://go.dev/play/p/g4JWVpHsJcf)]
-   **<big>RandNumeralOrLetter</big>** : generate a random numeral or letter string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#RandNumeralOrLetter)]
    [[play](https://go.dev/play/p/19CEQvpx2jD)]
-   **<big>UUIdV4</big>** : generate a random UUID of version 4 according to RFC 4122.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/random.md#UUIdV4)]
    [[play](https://go.dev/play/p/_Z9SFmr28ft)]

### 15. Retry package is for executing a function repeatedly until it was successful or canceled by the context.

```go
import "github.com/duke-git/lancet/v2/retry"
```

#### Function list:

-   **<big>Context</big>** : set retry context config option.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry.md#Context)]
    [[play](https://go.dev/play/p/xnAOOXv9GkS)]
-   **<big>Retry</big>** : executes the retryFunc repeatedly until it was successful or canceled by the context.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry.md#Retry)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryFunc</big>** : function that retry executes.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryFunc)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryDuration</big>** : set duration of retry
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryDuration)]
    [[play](https://go.dev/play/p/nk2XRmagfVF)]
-   **<big>RetryTimes</big>** : set times of retry.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/retry.md#RetryTimes)]
    [[play](https://go.dev/play/p/ssfVeU2SwLO)]

### 16. Slice contains some functions to manipulate slice.

```go
import "github.com/duke-git/lancet/v2/slice"
```

#### Function list:

-   **<big>AppendIfAbsent</big>** : if the item is absent,append it to the slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#AppendIfAbsent)]
    [[play](https://go.dev/play/p/GNdv7Jg2Taj)]
-   **<big>Contain</big>** : check if the value is in the slice or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Contain)]
    [[play](https://go.dev/play/p/_454yEHcNjf)]
-   **<big>ContainSubSlice</big>** : check if the slice contain a given subslice or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ContainSubSlice)]
    [[play](https://go.dev/play/p/bcuQ3UT6Sev)]
-   **<big>Chunk</big>** : creates a slice of elements split into groups the length of size.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Chunk)]
    [[play](https://go.dev/play/p/b4Pou5j2L_C)]
-   **<big>Compact</big>** : creates an slice with all falsey values removed. The values false, nil, 0, and "" are falsey.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Compact)]
    [[play](https://go.dev/play/p/pO5AnxEr3TK)]
-   **<big>Concat</big>** : creates a new slice concatenating slice with any additional slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Concat)]
    [[play](https://go.dev/play/p/gPt-q7zr5mk)]
-   **<big>Count</big>** : returns the number of occurrences of the given item in the slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Count)]
    [[play](https://go.dev/play/p/Mj4oiEnQvRJ)]
-   **<big>CountBy</big>** : iterates over elements of slice with predicate function, returns the number of all matched elements.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#CountBy)]
    [[play](https://go.dev/play/p/tHOccTMDZCC)]
-   **<big>Difference</big>** : creates an slice of whose element in slice but not in compared slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Difference)]
    [[play](https://go.dev/play/p/VXvadzLzhDa)]
-   **<big>DifferenceBy</big>** : accepts iteratee which is invoked for each element of slice and values to generate the criterion by which they're compared.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DifferenceBy)]
    [[play](https://go.dev/play/p/DiivgwM5OnC)]
-   **<big>DifferenceWith</big>** : accepts comparator which is invoked to compare elements of slice to values.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DifferenceWith)]
    [[play](https://go.dev/play/p/v2U2deugKuV)]
-   **<big>DeleteAt</big>** : delete the element of slice from specific start index to end index - 1.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DeleteAt)]
    [[play](https://go.dev/play/p/pJ-d6MUWcvK)]
-   **<big>Drop</big>** : drop n elements from the start of a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Drop)]
    [[play](https://go.dev/play/p/jnPO2yQsT8H)]
-   **<big>DropRight</big>** : drop n elements from the end of a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DropRight)]
    [[play](https://go.dev/play/p/8bcXvywZezG)]
-   **<big>DropWhile</big>** : drop n elements from the start of a slice while predicate function returns true.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DropWhile)]
    [[play](https://go.dev/play/p/4rt252UV_qs)]
-   **<big>DropRightWhile</big>** : drop n elements from the end of a slice while predicate function returns true.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#DropRightWhile)]
    [[play](https://go.dev/play/p/6wyK3zMY56e)]
-   **<big>Equal</big>** : checks if two slices are equal: the same length and all elements' order and value are equal.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Equal)]
    [[play](https://go.dev/play/p/WcRQJ37ifPa)]
-   **<big>EqualWith</big>** : checks if two slices are equal with comparator func.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#EqualWith)]
    [[play](https://go.dev/play/p/b9iygtgsHI1)]
-   **<big>Every</big>** : return true if all of the values in the slice pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Every)]
    [[play](https://go.dev/play/p/R8U6Sl-j8cD)]
-   **<big>Filter</big>** : iterates over elements of slice, returning an slice of all elements pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Filter)]
    [[play](https://go.dev/play/p/SdPna-7qK4T)]
-   **<big>Find</big>** : iterates over elements of slice, returning the first one that passes a truth test on predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Find)]
    [[play](https://go.dev/play/p/CBKeBoHVLgq)]
-   **<big>FindLast</big>** : return the last item that passes a truth test on predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#FindLast)]
    [[play](https://go.dev/play/p/FFDPV_j7URd)]
-   **<big>Flatten</big>** : flattens slice one level.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Flatten)]
    [[play](https://go.dev/play/p/hYa3cBEevtm)]
-   **<big>FlattenDeep</big>** : flattens slice recursive to one level.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#FlattenDeep)]
    [[play](https://go.dev/play/p/yjYNHPyCFaF)]
-   **<big>ForEach</big>** : iterates over elements of slice and invokes function for each element.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ForEach)]
    [[play](https://go.dev/play/p/DrPaa4YsHRF)]
-   **<big>GroupBy</big>** : iterate over elements of the slice, each element will be group by criteria, returns two slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#GroupBy)]
    [[play](https://go.dev/play/p/QVkPxzPR0iA)]
-   **<big>GroupWith</big>** : return a map composed of keys generated from the resultults of running each element of slice thru iteratee.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#GroupWith)]
    [[play](https://go.dev/play/p/ApCvMNTLO8a)]
-   **<big>IntSlice<sup>deprecated</sup></big>** : convert param to int slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IntSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>InterfaceSlice<sup>deprecated</sup></big>** : convert param to interface slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#InterfaceSlice)]
    [[play](https://go.dev/play/p/FdQXF0Vvqs-)]
-   **<big>Intersection</big>** : creates a slice of unique elements that included by all slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Intersection)]
    [[play](https://go.dev/play/p/anJXfB5wq_t)]
-   **<big>InsertAt</big>** : insert the value or other slice into slice at index.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#InsertAt)]
    [[play](https://go.dev/play/p/hMLNxPEGJVE)]
-   **<big>IndexOf</big>** : returns the index at which the first occurrence of an item is found in a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IndexOf)]
    [[play](https://go.dev/play/p/MRN1f0FpABb)]
-   **<big>LastIndexOf</big>** : returns the index at which the last occurrence of the item is found in a slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#LastIndexOf)]
    [[play](https://go.dev/play/p/DokM4cf1IKH)]
-   **<big>Map</big>** : creates an slice of values by running each element of slice thru iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Map)]
    [[play](https://go.dev/play/p/biaTefqPquw)]
-   **<big>Merge</big>** : merge all given slices into one slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Merge)]
    [[play](https://go.dev/play/p/lbjFp784r9N)]
-   **<big>Reverse</big>** : return slice of element order is reversed to the given slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Reverse)]
    [[play](https://go.dev/play/p/8uI8f1lwNrQ)]
-   **<big>Reduce</big>** : creates an slice of values by running each element of slice thru iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Reduce)]
    [[play](https://go.dev/play/p/_RfXJJWIsIm)]
-   **<big>Replace</big>** : returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Replace)]
    [[play](https://go.dev/play/p/P5mZp7IhOFo)]
-   **<big>ReplaceAll</big>** : returns a copy of the slice with all non-overlapping instances of old replaced by new.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ReplaceAll)]
    [[play](https://go.dev/play/p/CzqXMsuYUrx)]
-   **<big>Repeat</big>** : creates a slice with length n whose elements are passed item.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Repeat)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>Shuffle</big>** : shuffle the slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Shuffle)]
    [[play](https://go.dev/play/p/YHvhnWGU3Ge)]
-   **<big>IsAscending</big>** : Checks if a slice is ascending order.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IsAscending)]
    [[play](https://go.dev/play/p/9CtsFjet4SH)]
-   **<big>IsDescending</big>** : Checks if a slice is descending order.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IsDescending)]
    [[play](https://go.dev/play/p/U_LljFXma14)]
-   **<big>IsSorted</big>** : Checks if a slice is sorted (ascending or descending).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IsSorted)]
    [[play](https://go.dev/play/p/nCE8wPLwSA-)]
-   **<big>IsSortedByKey</big>** : Checks if a slice is sorted by iteratee function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#IsSortedByKey)]
    [[play](https://go.dev/play/p/tUoGB7DOHI4)]
-   **<big>Sort</big>** : sorts a slice of any ordered type(number or string).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Sort)]
    [[play](https://go.dev/play/p/V9AVjzf_4Fk)]
-   **<big>SortBy</big>** : sorts the slice in ascending order as determined by the less function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#SortBy)]
    [[play](https://go.dev/play/p/DAhLQSZEumm)]
-   **<big>SortByField<sup>deprecated</sup></big>** : return sorted slice by specific field.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#SortByField)]
    [[play](https://go.dev/play/p/fU1prOBP9p1)]
-   **<big>Some</big>** : return true if any of the values in the list pass the predicate function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Some)]
    [[play](https://go.dev/play/p/4pO9Xf9NDGS)]
-   **<big>StringSlice<sup>deprecated</sup></big>** : convert param to slice of string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#StringSlice)]
    [[play](https://go.dev/play/p/W0TZDWCPFcI)]
-   **<big>SymmetricDifference</big>** : the symmetric difference of two slice, also known as the disjunctive union.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#h42nJX5xMln)]
    [[play](https://go.dev/play/p/1CbOmtgILUU)]
-   **<big>ToSlice</big>** : returns a slices of a variable parameter transformation.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ToSlice)]
    [[play](https://go.dev/play/p/YzbzVq5kscN)]
-   **<big>ToSlicePointer</big>** : returns a pointer to the slices of a variable parameter transformation.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#ToSlicePointer)]
    [[play](https://go.dev/play/p/gx4tr6_VXSF)]
-   **<big>Unique</big>** : remove duplicate elements in slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Unique)]
    [[play](https://go.dev/play/p/AXw0R3ZTE6a)]
-   **<big>UniqueBy</big>** : call iteratee func with every item of slice, then remove duplicated.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#UniqueBy)]
    [[play](https://go.dev/play/p/UR323iZLDpv)]
-   **<big>Union</big>** : creates a slice of unique elements, in order, from all given slices.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Union)]
    [[play](https://go.dev/play/p/hfXV1iRIZOf)]
-   **<big>UnionBy</big>** : accepts iteratee which is invoked for each element of each slice, then union slice.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#UnionBy)]
    [[play](https://go.dev/play/p/HGKHfxKQsFi)]
-   **<big>UpdateAt</big>** : update the slice element at index.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#UpdateAt)]
    [[play](https://go.dev/play/p/f3mh2KloWVm)]
-   **<big>Without</big>** : creates a slice excluding all given items.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#Without)]
    [[play](https://go.dev/play/p/bwhEXEypThg)]
-   **<big>KeyBy</big>** : converts a slice to a map based on a callback function.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/slice.md#KeyBy)]
    [[play](https://go.dev/play/p/uXod2LWD1Kg)]

### 17. Strutil package contains some functions to manipulate string.

```go
import "github.com/duke-git/lancet/v2/strutil"
```

#### Function list:

-   **<big>After</big>** : returns the substring after the first occurrence of a specific string in the source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#After)]
    [[play](https://go.dev/play/p/RbCOQqCDA7m)]
-   **<big>AfterLast</big>** : returns the substring after the last occurrence of a specific string in the source string. [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#AfterLast)]
    [[play](https://go.dev/play/p/1TegARrb8Yn)]
-   **<big>Before</big>** : returns the substring before the first occurrence of a specific string in the source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Before)]
    [[play](https://go.dev/play/p/JAWTZDS4F5w)]
-   **<big>BeforeLast</big>** : returns the substring before the last occurrence of a specific string in the source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#BeforeLast)]
    [[play](https://go.dev/play/p/pJfXXAoG_Te)]
-   **<big>CamelCase</big>** : coverts source string to its camelCase string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#CamelCase)]
    [[play](https://go.dev/play/p/9eXP3tn2tUy)]
-   **<big>Capitalize</big>** : converts the first character of source string to upper case and the remaining to lower case.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Capitalize)]
    [[play](https://go.dev/play/p/2OAjgbmAqHZ)]
-   **<big>IsString</big>** : checks if the parameter value data type is string or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#IsString)]
    [[play](https://go.dev/play/p/IOgq7oF9ERm)]
-   **<big>KebabCase</big>** : coverts string to kebab-case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#KebabCase)]
    [[play](https://go.dev/play/p/dcZM9Oahw-Y)]
-   **<big>UpperKebabCase</big>** : coverts string to upper KEBAB-CASE string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#UpperKebabCase)]
    [[play](https://go.dev/play/p/zDyKNneyQXk)]
-   **<big>LowerFirst</big>** : converts the first character of string to lower case.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#LowerFirst)]
    [[play](https://go.dev/play/p/CbzAyZmtJwL)]
-   **<big>UpperFirst</big>** : converts the first character of string to upper case.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#UpperFirst)]
    [[play](https://go.dev/play/p/sBbBxRbs8MM)]
-   **<big>PadEnd</big>** : pads string with given characters on the right side if it's shorter than limit size. Padding characters are truncated if they exceed size.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#PadEnd)]
    [[play](https://go.dev/play/p/9xP8rN0vz--)]
-   **<big>PadStart</big>** : pads string with given characters on the left side if it's shorter than limit size. Padding characters are truncated if they exceed size.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#PadStart)]
    [[play](https://go.dev/play/p/xpTfzArDfvT)]
-   **<big>Reverse</big>** : returns string whose char order is reversed to the given string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Reverse)]
    [[play](https://go.dev/play/p/adfwalJiecD)]
-   **<big>SnakeCase</big>** : coverts string to snake_case string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#SnakeCase)]
    [[play](https://go.dev/play/p/tgzQG11qBuN)]
-   **<big>UpperSnakeCase</big>** : coverts string to upper SNAKE_CASE string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#UpperSnakeCase)]
    [[play](https://go.dev/play/p/4COPHpnLx38)]
-   **<big>SplitEx</big>** : split a given string which can control the result slice contains empty string or not.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#SplitEx)]
    [[play](https://go.dev/play/p/Us-ySSbWh-3)]
-   **<big>Substring</big>** : returns a substring of the specific length starting at the specific offset position.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Substring)]
    [[play](https://go.dev/play/p/q3sM6ehnPDp)]
-   **<big>Wrap</big>** : wrap a string with given string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Wrap)]
    [[play](https://go.dev/play/p/KoZOlZDDt9y)]
-   **<big>Unwrap</big>** : unwrap a given string from anther string. will change source string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/strutil.md#Unwrap)]
    [[play](https://go.dev/play/p/Ec2q4BzCpG-)]

### 19. System package contain some functions about os, runtime, shell command.

```go
import "github.com/duke-git/lancet/v2/system"
```

#### Function list:

-   **<big>IsWindows</big>** : check if current os is windows.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsWindows)]
    [[play](https://go.dev/play/p/XzJULbzmf9m)]
-   **<big>IsLinux</big>** : check if current os is linux.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsLinux)]
    [[play](https://go.dev/play/p/zIflQgZNuxD)]
-   **<big>IsMac</big>** : check if current os is macos.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#IsMac)]
    [[play](https://go.dev/play/p/Mg4Hjtyq7Zc)]
-   **<big>GetOsEnv</big>** : get the value of the environment variable named by the key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#GetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>SetOsEnv</big>** : set the value of the environment variable named by the key.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#SetOsEnv)]
    [[play](https://go.dev/play/p/D88OYVCyjO-)]
-   **<big>RemoveOsEnv</big>** : remove a single environment variable.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#RemoveOsEnv)]
    [[play](https://go.dev/play/p/fqyq4b3xUFQ)]
-   **<big>CompareOsEnv</big>** : get env named by the key and compare it with passed env.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#CompareOsEnv)]
    [[play](https://go.dev/play/p/BciHrKYOHbp)]
-   **<big>ExecCommand</big>** : execute command, return the stdout and stderr string of command, and error if error occurs.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#ExecCommand)]
    [[play](https://go.dev/play/p/n-2fLyZef-4)]
-   **<big>GetOsBits</big>** : return current os bits (32 or 64).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/system.md#GetOsBits)]
    [[play](https://go.dev/play/p/ml-_XH3gJbW)]

### 19. Validator package contains some functions for data validation.

```go
import "github.com/duke-git/lancet/v2/validator"
```

#### Function list:

-   **<big>ContainChinese</big>** : check if the string contain mandarin chinese.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainChinese)]
    [[play](https://go.dev/play/p/7DpU0uElYeM)]
-   **<big>ContainLetter</big>** : check if the string contain at least one letter.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainLetter)]
    [[play](https://go.dev/play/p/lqFD04Yyewp)]
-   **<big>ContainLower</big>** : check if the string contain at least one lower case letter a-z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainLower)]
    [[play](https://go.dev/play/p/Srqi1ItvnAA)]
-   **<big>ContainUpper</big>** : check if the string contain at least one upper case letter A-Z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#ContainUpper)]
    [[play](https://go.dev/play/p/CmWeBEk27-z)]
-   **<big>IsAlpha</big>** : checks if the string contains only letters (a-zA-Z).
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAlpha)]
    [[play](https://go.dev/play/p/7Q5sGOz2izQ)]
-   **<big>IsAllUpper</big>** : check if the string is all upper case letters A-Z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAllUpper)]
    [[play](https://go.dev/play/p/ZHctgeK1n4Z)]
-   **<big>IsAllLower</big>** : check if the string is all lower case letters a-z.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsAllLower)]
    [[play](https://go.dev/play/p/GjqCnOfV6cM)]
-   **<big>IsBase64</big>** : check if the string is base64 string.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsBase64)]
    [[play](https://go.dev/play/p/sWHEySAt6hl)]
-   **<big>IsChineseMobile</big>** : check if the string is chinese mobile number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChineseMobile)]
    [[play](https://go.dev/play/p/GPYUlGTOqe3)]
-   **<big>IsChineseIdNum</big>** : check if the string is chinese id card.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChineseIdNum)]
    [[play](https://go.dev/play/p/d8EWhl2UGDF)]
-   **<big>IsChinesePhone</big>** : check if the string is chinese phone number.(xxx-xxxxxxxx or xxxx-xxxxxxx.)
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsChinesePhone)]
    [[play](https://go.dev/play/p/RUD_-7YZJ3I)]
-   **<big>IsCreditCard</big>** : check if the string is credit card.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsCreditCard)]
    [[play](https://go.dev/play/p/sNwwL6B0-v4)]
-   **<big>IsDns</big>** : check if the string is dns.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsDns)]
    [[play](https://go.dev/play/p/jlYApVLLGTZ)]
-   **<big>IsEmail</big>** : check if the string is a email address.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsEmail)]
    [[play](https://go.dev/play/p/Os9VaFlT33G)]
-   **<big>IsEmptyString</big>** : check if the string is empty.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsEmptyString)]
    [[play](https://go.dev/play/p/dpzgUjFnBCX)]
-   **<big>IsFloatStr</big>** : check if the string can convert to a float.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#LOYwS_Oyl7U)]
    [[play](https://go.dev/play/p/LOYwS_Oyl7U)]
-   **<big>IsNumberStr</big>** : check if the string can convert to a number.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsNumberStr)]
    [[play](https://go.dev/play/p/LzaKocSV79u)]
-   **<big>IsJSON</big>** : check if the string is valid JSON.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsJSON)]
    [[play](https://go.dev/play/p/sRS6c4K8jGk)]
-   **<big>IsRegexMatch</big>** : check if the string match the regexp.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsRegexMatch)]
    [[play](https://go.dev/play/p/z_XeZo_litG)]
-   **<big>IsIntStr</big>** : check if the string can convert to a integer.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIntStr)]
    [[play](https://go.dev/play/p/jQRtFv-a0Rk)]
-   **<big>IsIp</big>** : check if the string is ip.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIp)]
    [[play](https://go.dev/play/p/FgcplDvmxoD)]
-   **<big>IsIpV4</big>** : check if the string is ipv4.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIpV4)]
    [[play](https://go.dev/play/p/zBGT99EjaIu)]
-   **<big>IsIpV6</big>** : check if the string is ipv6.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsIpV6)]
    [[play](https://go.dev/play/p/AHA0r0AzIdC)]
-   **<big>IsStrongPassword</big>** : check if the string is strong password.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsStrongPassword)]
    [[play](https://go.dev/play/p/QHdVcSQ3uDg)]
-   **<big>IsUrl</big>** : check if the string is url.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsUrl)]
    [[play](https://go.dev/play/p/pbJGa7F98Ka)]
-   **<big>IsWeakPassword</big>** : check if the string is weak password.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsWeakPassword)]
    [[play](https://go.dev/play/p/wqakscZH5gH)]
-   **<big>IsZeroValue</big>** : check if value is a zero value.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsZeroValue)]
    [[play](https://go.dev/play/p/UMrwaDCi_t4)]
-   **<big>IsGBK</big>** : check if data encoding is gbk.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/validator.md#IsGBK)]
    [[play](https://go.dev/play/p/E2nt3unlmzP)]

### 20. xerror package implements helpers for errors.

```go
import "github.com/duke-git/lancet/v2/xerror"
```

#### Function list:
-   **<big>New</big>** : creates a new XError pointer instance with message.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#New)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>Wrap</big>** : creates a new XError pointer instance based on error object, and add message.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#Wrap)]
    [[play](https://go.dev/play/p/5385qT2dCi4)]
-   **<big>Unwrap</big>** : returns unwrapped XError from err by errors.As. If no XError, returns nil.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#Unwrap)]
    [[play](https://go.dev/play/p/LKMLep723tu)]
-   **<big>XError_Wrap</big>** : creates a new XError and copy message and id to new one.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_Wrap)]
    [[play](https://go.dev/play/p/5385qT2dCi4)]
-   **<big>XError_Unwrap</big>** : Compatible with github.com/pkg/errors.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_Unwrap)]
    [[play](https://go.dev/play/p/VUXJ8BST4c6)]
-   **<big>XError_With</big>** : adds key and value related to the XError object.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_With)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_Id</big>** : sets XError object id to check equality in XError.Is.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_Id)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Is</big>** : checks if target error is XError and Error.id of two errors are matched.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_Is)]
    [[play](https://go.dev/play/p/X6HBlsy58U9)]
-   **<big>XError_Values</big>** : returns map of key and value that is set by XError.With function. 
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_Values)]
    [[play](https://go.dev/play/p/ow8UISXX_Dp)]
-   **<big>XError_StackTrace</big>** : returns stack trace which is compatible with pkg/errors.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_StackTrace)]
    [[play](https://go.dev/play/p/6FAvSQpa7pc)]
-   **<big>XError_Info</big>** : returns information of xerror, which can be printed.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_Info)]
    [[play](https://go.dev/play/p/1ZX0ME1F-Jb)]
-   **<big>XError_Error</big>** : implements standard error interface.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#XError_Error)]
    [[play](https://go.dev/play/p/w4oWZts7q7f)]
-   **<big>TryUnwrap</big>** : check if err is nil then it returns a valid value. If err is not nil, TryUnwrap panics with err.
    [[doc](https://github.com/duke-git/lancet/blob/main/docs/xerror.md#TryUnwrap)]
    [[play](https://go.dev/play/p/acyZVkNZEeW)]


## How to Contribute

I really appreciate any code commits which make lancet lib powerful. Please follow the rules below to create your pull request.

1. Fork the repository.
2. Create your feature branch.
3. Commit your changes.
4. Push to the branch
5. Create new pull request.
