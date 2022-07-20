# Cryptor
cryptor加密包支持数据加密和解密，获取md5，hash值。支持base64, md5, hmac, aes, des, rsa。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/cryptor/aes.go](https://github.com/duke-git/lancet/blob/main/cryptor/aes.go)
- [https://github.com/duke-git/lancet/blob/main/cryptor/des.go](https://github.com/duke-git/lancet/blob/main/cryptor/des.go)
- [https://github.com/duke-git/lancet/blob/main/cryptor/basic.go](https://github.com/duke-git/lancet/blob/main/cryptor/basic.go)
- [https://github.com/duke-git/lancet/blob/main/cryptor/rsa.go](https://github.com/duke-git/lancet/blob/main/cryptor/rsa.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/v2/cryptor"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [AesEcbEncrypt](#AesEcbEncrypt)
- [AesEcbDecrypt](#AesEcbDecrypt)
- [AesCbcEncrypt](#AesCbcEncrypt)
- [AesCbcDecrypt](#AesCbcDecrypt)
- [AesCtrCrypt](#AesCtrCrypt)
- [AesCfbEncrypt](#AesCfbEncrypt)
- [AesCfbDecrypt](#AesCfbDecrypt)
- [AesOfbEncrypt](#AesOfbEncrypt)
- [AesOfbDecrypt](#AesOfbDecrypt)
- [Base64StdEncode](#Base64StdEncode)
- [Base64StdDecode](#Base64StdDecode)

- [DesEcbEncrypt](#DesEcbEncrypt)
- [DesEcbDecrypt](#DesEcbDecrypt)
- [DesCbcEncrypt](#DesCbcEncrypt)
- [DesCbcDecrypt](#DesCbcDecrypt)
- [DesCtrCrypt](#DesCtrCrypt)
- [DesCfbEncrypt](#DesCfbEncrypt)
- [DesCfbDecrypt](#DesCfbDecrypt)
- [DesOfbEncrypt](#DesOfbEncrypt)
- [DesOfbDecrypt](#DesOfbDecrypt)

- [HmacMd5](#HmacMd5)
- [HmacSha1](#HmacSha1)
- [HmacSha256](#HmacSha256)
- [HmacSha512](#HmacSha512)
- [Md5String](#Md5String)
- [Md5File](#Md5File)
- [Sha1](#Sha1)
- [Sha256](#Sha256)
- [Sha512](#Sha512)
- [GenerateRsaKey](#GenerateRsaKey)
- [RsaEncrypt](#RsaEncrypt)
- [RsaDecrypt](#RsaDecrypt)


<div STYLE="page-break-after: always;"></div>

## 文档



### <span id="AesEcbEncrypt">AesEcbEncrypt</span>

<p>使用AES ECB算法模式加密数据. 参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesEcbEncrypt(data, key []byte) []byte
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))

    fmt.Println(string(encrypted))
}
```



### <span id="AesEcbDecrypt">AesEcbDecrypt</span>

<p>使用AES ECB算法模式解密数据. 参数`key`的长度是16, 24 or 32。

<b>函数签名:</b>

```go
func AesEcbDecrypt(encrypted, key []byte) []byte
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesEcbDecrypt(encrypted, []byte(key))
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="AesCbcEncrypt">AesCbcEncrypt</span>

<p>使用AES CBC算法模式加密数据. 参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCbcEncrypt(data, key []byte) []byte
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))

    fmt.Println(string(encrypted))
}
```



### <span id="AesCbcDecrypt">AesCbcDecrypt</span>

<p>使用AES CBC算法模式解密数据. 参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCbcDecrypt(encrypted, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))
    
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="AesCtrCrypt">AesCtrCrypt</span>

<p>使用AES CTR算法模式加密/解密数据. 参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCtrCrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesCtrCrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCtrCrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="AesCfbEncrypt">AesCfbEncrypt</span>

<p>使用AES CFB算法模式加密数据. 参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCfbEncrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))
    fmt.Println(string(encrypted))
}
```



### <span id="AesCfbDecrypt">AesCfbDecrypt</span>

<p>使用AES CFB算法模式解密数据. 参数`key`的长度是16, 24 or 32。</p>

<b>函数签名:</b>

```go
func AesCfbDecrypt(encrypted, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="AesOfbEncrypt">AesOfbEncrypt</span>

<p>使用AES OFB算法模式加密数据. 参数`key`的长度是16, 24 or 32</p>

<b>函数签名:</b>

```go
func AesOfbEncrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))
    fmt.Println(string(encrypted))
}
```



### <span id="AesCfbDecrypt">AesOfbDecrypt</span>

<p>使用AES OFB算法模式解密数据. 参数`key`的长度是16, 24 or 32</p>

<b>函数签名:</b>

```go
func AesOfbDecrypt(encrypted, key []byte) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefghijklmnop"
    encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.AesOfbDecrypt(encrypted, []byte(key))
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="Base64StdEncode">Base64StdEncode</span>

<p>将字符串base64编码</p>

<b>函数签名:</b>

```go
func Base64StdEncode(s string) string
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    base64Str := cryptor.Base64StdEncode("hello world")
    fmt.Println(base64Str) //aGVsbG8gd29ybGQ=
}
```



### <span id="Base64StdDecode">Base64StdDecode</span>

<p>解码base64字符串</p>

<b>函数签名:</b>

```go
func Base64StdDecode(s string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    str := cryptor.Base64StdDecode("aGVsbG8gd29ybGQ=")
    fmt.Println(str) //hello world
}
```



### <span id="DesEcbEncrypt">DesEcbEncrypt</span>

<p>使用DES ECB算法模式加密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesEcbEncrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))

    fmt.Println(string(encrypted))
}
```



### <span id="DesEcbDecrypt">DesEcbDecrypt</span>

<p>使用DES ECB算法模式解密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesEcbDecrypt(encrypted, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesEcbDecrypt(encrypted, []byte(key))
  
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="DesCbcEncrypt">DesCbcEncrypt</span>

<p>使用DES CBC算法模式加密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesCbcEncrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))

    fmt.Println(string(encrypted))
}
```



### <span id="DesCbcDecrypt">DesCbcDecrypt</span>

<p>使用DES CBC算法模式解密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesCbcDecrypt(encrypted, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCbcDecrypt(encrypted, []byte(key))
    
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="DesCtrCrypt">DesCtrCrypt</span>

<p>使用DES CTR算法模式加密/解密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesCtrCrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesCtrCrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCtrCrypt(encrypted, []byte(key))

    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="DesCfbEncrypt">DesCfbEncrypt</span>

<p>使用DES CFB算法模式加密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesCfbEncrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))
    fmt.Println(string(encrypted))
}
```



### <span id="DesCfbDecrypt">DesCfbDecrypt</span>

<p>使用DES CFB算法模式解密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesCfbDecrypt(encrypted, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesCfbDecrypt(encrypted, []byte(key))
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="DesOfbEncrypt">DesOfbEncrypt</span>

<p>使用DES OFB算法模式加密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesOfbEncrypt(data, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))
    fmt.Println(string(encrypted))
}
```



### <span id="DesOfbDecrypt">DesOfbDecrypt</span>

<p>使用DES OFB算法模式解密数据. 参数`key`的长度是8</p>

<b>函数签名:</b>

```go
func DesOfbDecrypt(encrypted, key []byte) []byte
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    data := "hello world"
    key := "abcdefgh"
    encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))
    decrypted := cryptor.DesOfbDecrypt(encrypted, []byte(key))
    fmt.Println(string(decrypted)) //hello world
}
```



### <span id="HmacMd5">HmacMd5</span>

<p>获取字符串md5 hmac值</p>

<b>函数签名:</b>

```go
func HmacMd5(data, key string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.HmacMd5("hello world", "12345"))
    fmt.Println(s) //5f4c9faaff0a1ad3007d9ddc06abe36d
}
```



### <span id="HmacSha1">HmacSha1</span>

<p>获取字符串sha1 hmac值</p>

<b>函数签名:</b>

```go
func HmacSha1(data, key string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.HmacSha1("hello world", "12345"))
    fmt.Println(s) //3826f812255d8683f051ee97346d1359234d5dbd
}
```



### <span id="HmacSha256">HmacSha256</span>

<p>获取字符串sha256 hmac值</p>

<b>函数签名:</b>

```go
func HmacSha256(data, key string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.HmacSha256("hello world", "12345"))
    fmt.Println(s) //9dce2609f2d67d41f74c7f9efc8ccd44370d41ad2de52982627588dfe7289ab8
}
```



### <span id="HmacSha512">HmacSha512</span>

<p>获取字符串sha512 hmac值</p>

<b>函数签名:</b>

```go
func HmacSha512(data, key string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.HmacSha512("hello world", "12345"))
    fmt.Println(s) 
 //5b1563ac4e9b49c9ada8ccb232588fc4f0c30fd12f756b3a0b95af4985c236ca60925253bae10ce2c6bf9af1c1679b51e5395ff3d2826c0a2c7c0d72225d4175
}
```



### <span id="Md5String">Md5String</span>

<p>获取字符串md5值</p>

<b>函数签名:</b>

```go
func Md5String(s string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.Md5String("hello"))
    fmt.Println(s) //5d41402abc4b2a76b9719d911017c592
}
```



### <span id="Md5File">Md5File</span>

<p>获取文件md5值.</p>

<b>函数签名:</b>

```go
func Md5File(filepath string) (string, error)
```

<b>列子:</b>

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

<p>获取字符串sha1值</p>

<b>函数签名:</b>

```go
func Sha1(data string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.Sha1("hello world"))
    fmt.Println(s) //2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
}
```



### <span id="Sha256">Sha256</span>

<p>获取字符串sha256值</p>

<b>函数签名:</b>

```go
func Sha256(data string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.Sha256("hello world"))
    fmt.Println(s) //b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}
```



### <span id="Sha512">Sha512</span>

<p>获取字符串sha512值</p>

<b>函数签名:</b>

```go
func Sha512(data string) string
```

<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/cryptor"
)

func main() {
    s := cryptor.Sha512("hello world"))
    fmt.Println(s) //309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
}
```



### <span id="GenerateRsaKey">GenerateRsaKey</span>

<p>在当前目录下创建rsa私钥文件和公钥文件</p>

<b>函数签名:</b>

```go
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error
```

<b>列子:</b>

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

<p>用公钥文件ras加密数据</p>

<b>函数签名:</b>

```go
func RsaEncrypt(data []byte, pubKeyFileName string) []byte
```

<b>列子:</b>

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
    data := []byte("hello world")
    encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
    decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")
  	fmt.Println(string(decrypted)) //hello world
}
```



### <span id="RsaDecrypt">RsaDecrypt</span>

<p>用私钥文件rsa解密数据</p>

<b>函数签名:</b>

```go
func RsaDecrypt(data []byte, privateKeyFileName string) []byte
```

<b>列子:</b>

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
    data := []byte("hello world")
    encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
    decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")
  	fmt.Println(string(decrypted)) //hello world
}
```



