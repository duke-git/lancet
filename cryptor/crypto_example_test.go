package cryptor

import "fmt"

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
	err := GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		return
	}

	fmt.Println("foo")

	// Output:
	// foo
}

func ExampleRsaEncrypt() {
	// Create ras private and public pem file
	err := GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		return
	}

	data := []byte("hello")
	encrypted := RsaEncrypt(data, "rsa_public.pem")
	decrypted := RsaDecrypt(encrypted, "rsa_private.pem")

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleRsaDecrypt() {
	// Create ras private and public pem file
	err := GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		return
	}

	data := []byte("hello")
	encrypted := RsaEncrypt(data, "rsa_public.pem")
	decrypted := RsaDecrypt(encrypted, "rsa_private.pem")

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

func ExampleHmacSha1() {
	str := "hello"
	key := "12345"

	hms := HmacSha1(str, key)

	fmt.Println(hms)

	// Output:
	// 5c6a9db0cccb92e36ed0323fd09b7f936de9ace0
}

func ExampleHmacSha256() {
	str := "hello"
	key := "12345"

	hms := HmacSha256(str, key)

	fmt.Println(hms)

	// Output:
	// 315bb93c4e989862ba09cb62e05d73a5f376cb36f0d786edab0c320d059fde75
}

func ExampleHmacSha512() {
	str := "hello"
	key := "12345"

	hms := HmacSha512(str, key)

	fmt.Println(hms)

	// Output:
	// dd8f1290a9dd23d354e2526d9a2e9ce8cffffdd37cb320800d1c6c13d2efc363288376a196c5458daf53f8e1aa6b45a6d856303d5c0a2064bff9785861d48cfc
}

func ExampleMd5String() {
	str := "hello"

	md5Str := Md5String(str)

	fmt.Println(md5Str)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}

func ExampleSha1() {
	str := "hello"

	result := Sha1(str)

	fmt.Println(result)

	// Output:
	// aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}

func ExampleSha256() {
	str := "hello"

	result := Sha256(str)

	fmt.Println(result)

	// Output:
	// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}

func ExampleSha512() {
	str := "hello"

	result := Sha512(str)

	fmt.Println(result)

	// Output:
	// 9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
}
