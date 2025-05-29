package cryptor

import (
	"crypto"
	"fmt"
)

func ExampleAesEcbEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesEcbEncrypt([]byte(data), []byte(key))

	decrypted := AesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesEcbDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesEcbEncrypt([]byte(data), []byte(key))

	decrypted := AesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCbcEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCbcEncrypt([]byte(data), []byte(key))

	decrypted := AesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCbcDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCbcEncrypt([]byte(data), []byte(key))

	decrypted := AesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCtrCrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCtrCrypt([]byte(data), []byte(key))
	decrypted := AesCtrCrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCtrEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	enCrypt := AesCtrEncrypt([]byte(data), []byte(key))
	deCrypt := AesCtrDecrypt(enCrypt, []byte(key))

	fmt.Println(string(deCrypt))

	// Output:
	// hello
}

func ExampleAesCtrDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	enCrypt := AesCtrEncrypt([]byte(data), []byte(key))
	deCrypt := AesCtrDecrypt(enCrypt, []byte(key))

	fmt.Println(string(deCrypt))

	// Output:
	// hello
}

func ExampleAesCfbEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCfbEncrypt([]byte(data), []byte(key))

	decrypted := AesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCfbDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCfbEncrypt([]byte(data), []byte(key))

	decrypted := AesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesOfbEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesOfbEncrypt([]byte(data), []byte(key))

	decrypted := AesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesOfbDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesOfbEncrypt([]byte(data), []byte(key))

	decrypted := AesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesGcmEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesGcmEncrypt([]byte(data), []byte(key))

	decrypted := AesGcmDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesGcmDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesGcmEncrypt([]byte(data), []byte(key))

	decrypted := AesGcmDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesEcbEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesEcbEncrypt([]byte(data), []byte(key))

	decrypted := DesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesEcbDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesEcbEncrypt([]byte(data), []byte(key))

	decrypted := DesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCbcEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCbcEncrypt([]byte(data), []byte(key))

	decrypted := DesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCbcDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCbcEncrypt([]byte(data), []byte(key))

	decrypted := DesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCtrCrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCtrCrypt([]byte(data), []byte(key))
	decrypted := DesCtrCrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCtrDecrypt() {
	data := "hello"
	key := "abcdefgh"

	enCrypt := DesCtrEncrypt([]byte(data), []byte(key))
	deCrypt := DesCtrDecrypt(enCrypt, []byte(key))

	fmt.Println(string(deCrypt))

	// Output:
	// hello
}

func ExampleDesCfbEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCfbEncrypt([]byte(data), []byte(key))

	decrypted := DesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCfbDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCfbEncrypt([]byte(data), []byte(key))

	decrypted := DesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesOfbEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesOfbEncrypt([]byte(data), []byte(key))

	decrypted := DesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesOfbDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesOfbEncrypt([]byte(data), []byte(key))

	decrypted := DesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleGenerateRsaKey() {
	// Create ras private and public pem file
	err := GenerateRsaKey(4096, "rsa_private_example.pem", "rsa_public_example.pem")
	if err != nil {
		return
	}

	fmt.Println("foo")

	// Output:
	// foo
}

func ExampleRsaEncrypt() {
	// Create ras private and public pem file
	err := GenerateRsaKey(4096, "rsa_private_example.pem", "rsa_public_example.pem")
	if err != nil {
		return
	}

	data := []byte("hello")
	encrypted := RsaEncrypt(data, "rsa_public_example.pem")
	decrypted := RsaDecrypt(encrypted, "rsa_private_example.pem")

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleRsaDecrypt() {
	// Create ras private and public pem file
	err := GenerateRsaKey(4096, "rsa_private_example.pem", "rsa_public_example.pem")
	if err != nil {
		return
	}

	data := []byte("hello")
	encrypted := RsaEncrypt(data, "rsa_public_example.pem")
	decrypted := RsaDecrypt(encrypted, "rsa_private_example.pem")

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleBase64StdEncode() {
	base64Str := Base64StdEncode("hello")

	fmt.Println(base64Str)

	// Output:
	// aGVsbG8=
}

func ExampleBase64StdDecode() {
	str := Base64StdDecode("aGVsbG8=")

	fmt.Println(str)

	// Output:
	// hello
}

func ExampleHmacMd5() {
	str := "hello"
	key := "12345"

	hms := HmacMd5(str, key)
	fmt.Println(hms)

	// Output:
	// e834306eab892d872525d4918a7a639a
}

func ExampleHmacMd5WithBase64() {
	str := "hello"
	key := "12345"

	hms := HmacMd5WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// 6DQwbquJLYclJdSRinpjmg==
}

func ExampleHmacSha1() {
	str := "hello"
	key := "12345"

	hms := HmacSha1(str, key)
	fmt.Println(hms)

	// Output:
	// 5c6a9db0cccb92e36ed0323fd09b7f936de9ace0
}

func ExampleHmacSha1WithBase64() {
	str := "hello"
	key := "12345"

	hms := HmacSha1WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// XGqdsMzLkuNu0DI/0Jt/k23prOA=
}

func ExampleHmacSha256() {
	str := "hello"
	key := "12345"

	hms := HmacSha256(str, key)
	fmt.Println(hms)

	// Output:
	// 315bb93c4e989862ba09cb62e05d73a5f376cb36f0d786edab0c320d059fde75
}

func ExampleHmacSha256WithBase64() {
	str := "hello"
	key := "12345"

	hms := HmacSha256WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// MVu5PE6YmGK6Ccti4F1zpfN2yzbw14btqwwyDQWf3nU=
}

func ExampleHmacSha512() {
	str := "hello"
	key := "12345"

	hms := HmacSha512(str, key)
	fmt.Println(hms)

	// Output:
	// dd8f1290a9dd23d354e2526d9a2e9ce8cffffdd37cb320800d1c6c13d2efc363288376a196c5458daf53f8e1aa6b45a6d856303d5c0a2064bff9785861d48cfc
}

func ExampleHmacSha512WithBase64() {
	str := "hello"
	key := "12345"

	hms := HmacSha512WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// 3Y8SkKndI9NU4lJtmi6c6M///dN8syCADRxsE9Lvw2Mog3ahlsVFja9T+OGqa0Wm2FYwPVwKIGS/+XhYYdSM/A==
}

func ExampleMd5String() {
	md5Str := Md5String("hello")
	fmt.Println(md5Str)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}

func ExampleMd5StringWithBase64() {
	md5Str := Md5StringWithBase64("hello")
	fmt.Println(md5Str)

	// Output:
	// XUFAKrxLKna5cZ2REBfFkg==
}

func ExampleMd5Byte() {
	md5Str := Md5Byte([]byte{'a'})
	fmt.Println(md5Str)

	// Output:
	// 0cc175b9c0f1b6a831c399e269772661
}

func ExampleMd5ByteWithBase64() {
	md5Str := Md5ByteWithBase64([]byte("hello"))
	fmt.Println(md5Str)

	// Output:
	// XUFAKrxLKna5cZ2REBfFkg==
}

func ExampleSha1() {
	result := Sha1("hello")
	fmt.Println(result)

	// Output:
	// aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}

func ExampleSha1WithBase64() {
	result := Sha1WithBase64("hello")
	fmt.Println(result)

	// Output:
	// qvTGHdzF6KLavt4PO0gs2a6pQ00=
}

func ExampleSha256() {
	result := Sha256("hello")
	fmt.Println(result)

	// Output:
	// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}

func ExampleSha256WithBase64() {
	result := Sha256WithBase64("hello")
	fmt.Println(result)

	// Output:
	// LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=
}

func ExampleSha512() {
	result := Sha512("hello")
	fmt.Println(result)

	// Output:
	// 9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
}

func ExampleSha512WithBase64() {
	result := Sha512WithBase64("hello")
	fmt.Println(result)

	// Output:
	// m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw==
}

func ExampleRsaEncryptOAEP() {
	pri, pub := GenerateRsaKeyPair(1024)

	data := []byte("hello world")
	label := []byte("123456")

	encrypted, err := RsaEncryptOAEP(data, label, *pub)
	if err != nil {
		return
	}

	decrypted, err := RsaDecryptOAEP([]byte(encrypted), label, *pri)
	if err != nil {
		return
	}

	fmt.Println(string(decrypted))

	// Output:
	// hello world
}

func ExampleRsaSign() {
	data := []byte("This is a test data for RSA signing")
	hash := crypto.SHA256

	privateKey := "./rsa_private_example.pem"
	publicKey := "./rsa_public_example.pem"

	signature, err := RsaSign(hash, data, privateKey)
	if err != nil {
		return
	}

	err = RsaVerifySign(hash, data, signature, publicKey)
	if err != nil {
		return
	}

	fmt.Println("ok")

	// Output:
	// ok
}

func ExampleRsaVerifySign() {
	data := []byte("This is a test data for RSA signing")
	hash := crypto.SHA256

	privateKey := "./rsa_private_example.pem"
	publicKey := "./rsa_public_example.pem"

	signature, err := RsaSign(hash, data, privateKey)
	if err != nil {
		return
	}

	err = RsaVerifySign(hash, data, signature, publicKey)
	if err != nil {
		return
	}

	fmt.Println("ok")

	// Output:
	// ok
}
