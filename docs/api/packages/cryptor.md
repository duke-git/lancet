# Cryptor

cryptor 包包含数据加密和解密功能。支持 base64, md5, hmac, hash, aes, des, rsa。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/cryptor/basic.go](https://github.com/duke-git/lancet/blob/main/cryptor/basic.go)
-   [https://github.com/duke-git/lancet/blob/main/cryptor/crypto.go](https://github.com/duke-git/lancet/blob/main/cryptor/crypto.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/cryptor"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [AesEcbEncrypt](#AesEcbEncrypt)
-   [AesEcbDecrypt](#AesEcbDecrypt)
-   [AesCbcEncrypt](#AesCbcEncrypt)
-   [AesCbcDecrypt](#AesCbcDecrypt)
-   [AesCtrCrypt](#AesCtrCrypt)
-   [AesCfbEncrypt](#AesCfbEncrypt)
-   [AesCfbDecrypt](#AesCfbDecrypt)
-   [AesOfbEncrypt](#AesOfbEncrypt)
-   [AesOfbDecrypt](#AesOfbDecrypt)
-   [Base64StdEncode](#Base64StdEncode)
-   [Base64StdDecode](#Base64StdDecode)
-   [DesEcbEncrypt](#DesEcbEncrypt)
-   [DesEcbDecrypt](#DesEcbDecrypt)
-   [DesCbcEncrypt](#DesCbcEncrypt)
-   [DesCbcDecrypt](#DesCbcDecrypt)
-   [DesCtrCrypt](#DesCtrCrypt)
-   [DesCfbEncrypt](#DesCfbEncrypt)
-   [DesCfbDecrypt](#DesCfbDecrypt)
-   [DesOfbEncrypt](#DesOfbEncrypt)
-   [DesOfbDecrypt](#DesOfbDecrypt)
-   [HmacMd5](#HmacMd5)
-   [HmacMd5WithBase64](#HmacMd5WithBase64)
-   [HmacSha1](#HmacSha1)
-   [HmacSha1WithBase64](#HmacSha1WithBase64)
-   [HmacSha256](#HmacSha256)
-   [HmacSha256WithBase64](#HmacSha256WithBase64)
-   [HmacSha512](#HmacSha512)
-   [HmacSha512WithBase64](#HmacSha512WithBase64)
-   [Md5String](#Md5String)
-   [Md5StringWithBase64](#Md5StringWithBase64)
-   [Md5Byte](#Md5Byte)
-   [Md5ByteWithBase64](#Md5ByteWithBase64)
-   [Md5File](#Md5File)
-   [Sha1](#Sha1)
-   [Sha1WithBase64](#Sha1WithBase64)
-   [Sha256](#Sha256)
-   [Sha256WithBase64](#Sha256WithBase64)
-   [Sha512](#Sha512)
-   [Sha512WithBase64](#Sha512WithBase64)
-   [GenerateRsaKey](#GenerateRsaKey)
-   [RsaEncrypt](#RsaEncrypt)
-   [RsaDecrypt](#RsaDecrypt)
-   [GenerateRsaKeyPair](#GenerateRsaKeyPair)
-   [RsaEncryptOAEP](#RsaEncryptOAEP)
-   [RsaDecryptOAEP](#RsaDecryptOAEP)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="AesEcbEncrypt">AesEcbEncrypt</span>

<p>使用AES ECB算法模式加密数据. 参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesEcbEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/zI6xsmuQRbn)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesEcbDecrypt">AesEcbDecrypt</span>

<p>使用AES ECB算法模式解密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesEcbDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/zI6xsmuQRbn)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCbcEncrypt">AesCbcEncrypt</span>

<p>使用AES CBC算法模式加密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCbcEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/IOq_g8_lKZD)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCbcDecrypt">AesCbcDecrypt</span>

<p>使用AES CBC算法模式解密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCbcDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/IOq_g8_lKZD)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCtrCrypt">AesCtrCrypt</span>

<p>使用AES CTR算法模式加密/解密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCtrCrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/SpaZO0-5Nsp)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCtrCrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCtrCrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCfbEncrypt">AesCfbEncrypt</span>

<p>使用AES CFB算法模式加密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCfbEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/tfkF10B13kH)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCfbDecrypt">AesCfbDecrypt</span>

<p>使用AES CFB算法模式解密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCfbDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/tfkF10B13kH)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesOfbEncrypt">AesOfbEncrypt</span>

<p>使用AES OFB算法模式加密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesOfbEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VtHxtkUj-3F)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="AesCfbDecrypt">AesOfbDecrypt</span>

<p>使用AES OFB算法模式解密数据，参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesOfbDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VtHxtkUj-3F)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefghijklmnop"

    encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="Base64StdEncode">Base64StdEncode</span>

<p>将字符串base64编码。</p>

<b>函数签名:</b>

```go
func Base64StdEncode(s string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VOaUyQUreoK)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    base64Str := cryptor.Base64StdEncode("hello")
    fmt.Println(base64Str)

    // Output:
    // aGVsbG8=
}
```

### <span id="Base64StdDecode">Base64StdDecode</span>

<p>解码base64字符串。</p>

<b>函数签名:</b>

```go
func Base64StdDecode(s string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/RWQylnJVgIe)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := cryptor.Base64StdDecode("aGVsbG8=")
    fmt.Println(str)

    // Output:
    // hello
}
```

### <span id="DesEcbEncrypt">DesEcbEncrypt</span>

<p>使用DES ECB算法模式加密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesEcbEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/8qivmPeZy4P)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))

    decrypted := cryptor.DesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesEcbDecrypt">DesEcbDecrypt</span>

<p>使用DES ECB算法模式解决密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesEcbDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/8qivmPeZy4P)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))

    decrypted := cryptor.DesEcbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCbcEncrypt">DesCbcEncrypt</span>

<p>使用DES CBC算法模式加密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesCbcEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/4cC4QvWfe3_1)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCbcDecrypt">DesCbcDecrypt</span>

<p>使用DES CBC算法模式解密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesCbcDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/4cC4QvWfe3_1)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCbcDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCtrCrypt">DesCtrCrypt</span>

<p>使用DES CTR算法模式加密/解密数据，参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesCtrCrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9-T6OjKpcdw)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCtrCrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCtrCrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCfbEncrypt">DesCfbEncrypt</span>

<p>使用DES CFB算法模式加密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesCfbEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/y-eNxcFBlxL)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesCfbDecrypt">DesCfbDecrypt</span>

<p>使用DES CFB算法模式解决密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesCfbDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/y-eNxcFBlxL)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesOfbEncrypt">DesOfbEncrypt</span>

<p>使用DES OFB算法模式加密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesOfbEncrypt(data, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/74KmNadjN1J)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesOfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="DesOfbDecrypt">DesOfbDecrypt</span>

<p>使用DES OFB算法模式解密数据，参数`key`的长度是8。</p>

<b>函数签名:</b>

```go
func DesOfbDecrypt(encrypted, key []byte) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/74KmNadjN1J)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello"
    key := "abcdefgh"

    encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesOfbDecrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="HmacMd5">HmacMd5</span>

<p>获取字符串md5 hmac值。</p>

<b>函数签名:</b>

```go
func HmacMd5(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/uef0q1fz53I)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacMd5(str, key)
    fmt.Println(hms)

    // Output:
    // e834306eab892d872525d4918a7a639a
}
```

### <span id="HmacMd5WithBase64">HmacMd5WithBase64</span>

<p>获取字符串md5 hmac base64字符串值。</p>

<b>函数签名:</b>

```go
func HmacMd5WithBase64(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/UY0ng2AefFC)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacMd5WithBase64(str, key)
    fmt.Println(hms)

    // Output:
    // 6DQwbquJLYclJdSRinpjmg==
}
```

### <span id="HmacSha1">HmacSha1</span>

<p>获取字符串的sha1 hmac值。</p>

<b>函数签名:</b>

```go
func HmacSha1(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/1UI4oQ4WXKM)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha1(str, key)
    fmt.Println(hms)

    // Output:
    // 5c6a9db0cccb92e36ed0323fd09b7f936de9ace0
}
```

### <span id="HmacSha1WithBase64">HmacSha1WithBase64</span>

<p>获取字符串的sha1 base64值。</p>

<b>函数签名:</b>

```go
func HmacSha1WithBase64(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/47JmmGrnF7B)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha1WithBase64(str, key)
    fmt.Println(hms)

    // Output:
    // XGqdsMzLkuNu0DI/0Jt/k23prOA=
}
```

### <span id="HmacSha256">HmacSha256</span>

<p>获取字符串sha256 hmac值。</p>

<b>函数签名:</b>

```go
func HmacSha256(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/HhpwXxFhhC0)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha256(str, key)
    fmt.Println(hms)

    // Output:
    // 315bb93c4e989862ba09cb62e05d73a5f376cb36f0d786edab0c320d059fde75
}
```

### <span id="HmacSha256WithBase64">HmacSha256WithBase64</span>

<p>获取字符串sha256 hmac base64值。</p>

<b>函数签名:</b>

```go
func HmacSha256WithBase64(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/EKbkUvPTLwO)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha256WithBase64(str, key)
    fmt.Println(hms)

    // Output:
    // MVu5PE6YmGK6Ccti4F1zpfN2yzbw14btqwwyDQWf3nU=
}
```

### <span id="HmacSha512">HmacSha512</span>

<p>获取字符串sha512 hmac值。</p>

<b>函数签名:</b>

```go
func HmacSha512(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/59Od6m4A0Ud)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha512(str, key)
    fmt.Println(hms)

    // Output:
    // dd8f1290a9dd23d354e2526d9a2e9ce8cffffdd37cb320800d1c6c13d2efc363288376a196c5458daf53f8e1aa6b45a6d856303d5c0a2064bff9785861d48cfc
}
```

### <span id="HmacSha512WithBase64">HmacSha512WithBase64</span>

<p>获取字符串sha512 hmac base64值。</p>

<b>函数签名:</b>

```go
func HmacSha512WithBase64(str, key string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/c6dSe3E2ydU)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"
    key := "12345"

    hms := cryptor.HmacSha512WithBase64(str, key)
    fmt.Println(hms)

    // Output:
    // 3Y8SkKndI9NU4lJtmi6c6M///dN8syCADRxsE9Lvw2Mog3ahlsVFja9T+OGqa0Wm2FYwPVwKIGS/+XhYYdSM/A==
}
```

### <span id="Md5String">Md5String</span>

<p>获取字符串md5值。</p>

<b>函数签名:</b>

```go
func Md5String(str string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/1bLcVetbTOI)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"

    md5Str := cryptor.Md5String(str)
    fmt.Println(md5Str)

    // Output:
    // 5d41402abc4b2a76b9719d911017c592
}
```

### <span id="Md5StringWithBase64">Md5StringWithBase64</span>

<p>获取字符串md5 base64值。</p>

<b>函数签名:</b>

```go
func Md5StringWithBase64(s string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Tcb-Z7LN2ax)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    md5Str := cryptor.Md5StringWithBase64("hello")
    fmt.Println(md5Str)

    // Output:
    // XUFAKrxLKna5cZ2REBfFkg==
}
```

### <span id="Md5Byte">Md5Byte</span>

<p>获取byte slice的md5值。</p>

<b>函数签名:</b>

```go
func Md5Byte(data []byte) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/suraalH8lyC)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    md5Str := cryptor.Md5Byte([]byte{'a'})
    fmt.Println(md5Str)

    // Output:
    // 0cc175b9c0f1b6a831c399e269772661
}
```

### <span id="Md5ByteWithBase64">Md5ByteWithBase64</span>

<p>获取byte slice的md5 base64值。</p>

<b>函数签名:</b>

```go
func Md5ByteWithBase64(data []byte) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Lx4gH7Vdr5_y)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    md5Str := cryptor.Md5ByteWithBase64([]byte("hello"))
    fmt.Println(md5Str)

    // Output:
    // XUFAKrxLKna5cZ2REBfFkg==
}
```

### <span id="Md5File">Md5File</span>

<p>获取文件md5值。</p>

<b>函数签名:</b>

```go
func Md5File(filepath string) (string, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.Md5File("./main.go"))
    fmt.Println(s)
}
```

### <span id="Sha1">Sha1</span>

<p>获取字符串sha1值。</p>

<b>函数签名:</b>

```go
func Sha1(str string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_m_uoD1deMT)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"

    result := cryptor.Sha1(str)
    fmt.Println(result)

    // Output:
    // aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}
```

### <span id="Sha1WithBase64">Sha1WithBase64</span>

<p>获取字符串sha1 base64值。</p>

<b>函数签名:</b>

```go
func Sha1WithBase64(str string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/fSyx-Gl2l2-)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    result := cryptor.Sha1WithBase64("hello")
    fmt.Println(result)

    // Output:
    // qvTGHdzF6KLavt4PO0gs2a6pQ00=
}
```

### <span id="Sha256">Sha256</span>

<p>获取字符串sha256值。</p>

<b>函数签名:</b>

```go
func Sha256(str string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/tU9tfBMIAr1)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"

    result := cryptor.Sha256(str)
    fmt.Println(result)

    // Output:
    // 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}
```

### <span id="Sha256WithBase64">Sha256WithBase64</span>

<p>获取字符串sha256 base64值。</p>

<b>函数签名:</b>

```go
func Sha256WithBase64(str string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/85IXJHIal1k)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    result := cryptor.Sha256WithBase64("hello")
    fmt.Println(result)

    // Output:
    // LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=
}
```

### <span id="Sha512">Sha512</span>

<p>获取字符串sha512值。</p>

<b>函数签名:</b>

```go
func Sha512(str string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/3WsvLYZxsHa)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := "hello"

    result := cryptor.Sha512(str)
    fmt.Println(result)

    // Output:
    // 9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
}
```

### <span id="Sha512WithBase64">Sha512WithBase64</span>

<p>获取字符串sha512 base64值。</p>

<b>函数签名:</b>

```go
func Sha512WithBase64(str string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/q_fY2rA-k5I)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    result := cryptor.Sha512WithBase64("hello")
    fmt.Println(result)

    // Output:
    // m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw==
}
```

### <span id="GenerateRsaKey">GenerateRsaKey</span>

<p>在当前目录下创建rsa私钥文件和公钥文件。</p>

<b>函数签名:</b>

```go
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/zutRHrDqs0X)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="RsaEncrypt">RsaEncrypt</span>

<p>用公钥文件ras加密数据。</p>

<b>函数签名:</b>

```go
func RsaEncrypt(data []byte, pubKeyFileName string) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/uef0q1fz53I)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
    if err != nil {
        return
    }

    data := []byte("hello")
    encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
    decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="RsaDecrypt">RsaDecrypt</span>

<p>用私钥文件rsa解密数据。</p>

<b>函数签名:</b>

```go
func RsaDecrypt(data []byte, privateKeyFileName string) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/uef0q1fz53I)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
    if err != nil {
        return
    }

    data := []byte("hello")
    encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
    decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")

    fmt.Println(string(decrypted))

    // Output:
    // hello
}
```

### <span id="GenerateRsaKeyPair">GenerateRsaKeyPair</span>

<p>创建rsa公钥私钥和key。</p>

<b>函数签名:</b>

```go
func GenerateRsaKeyPair(keySize int) (*rsa.PrivateKey, *rsa.PublicKey)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/sSVmkfENKMz)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    pri, pub := cryptor.GenerateRsaKeyPair(1024)
}
```

### <span id="RsaEncryptOAEP">RsaEncryptOAEP</span>

<p>rsa OAEP加密。</p>

<b>函数签名:</b>

```go
func RsaEncryptOAEP(data []byte, label []byte, key rsa.PublicKey) ([]byte, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/sSVmkfENKMz)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    pri, pub := cryptor.GenerateRsaKeyPair(1024)

    data := []byte("hello world")
    label := []byte("123456")

    encrypted, err := cryptor.RsaEncryptOAEP(data, label, *pub)
    if err != nil {
        return
    }

    decrypted, err := cryptor.RsaDecryptOAEP([]byte(encrypted), label, *pri)
    if err != nil {
        return
    }

    fmt.Println(string(decrypted))

    // Output:
    // hello world
}
```

### <span id="RsaDecryptOAEP">RsaDecryptOAEP</span>

<p>rsa OAEP解密。</p>

<b>函数签名:</b>

```go
func RsaDecryptOAEP(ciphertext []byte, label []byte, key rsa.PrivateKey) ([]byte, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/sSVmkfENKMz)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    pri, pub := cryptor.GenerateRsaKeyPair(1024)

    data := []byte("hello world")
    label := []byte("123456")

    encrypted, err := cryptor.RsaEncryptOAEP(data, label, *pub)
    if err != nil {
        return
    }

    decrypted, err := cryptor.RsaDecryptOAEP([]byte(encrypted), label, *pri)
    if err != nil {
        return
    }

    fmt.Println(string(decrypted))

    // Output:
    // hello world
}
```
