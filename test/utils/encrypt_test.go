package utils

import (
	"fmt"
	"gorm.io/gorm/utils"
	"testing"
)

func TestAesBase64Encrypt(t *testing.T) {
	str := "0x0000000002c4702b4fe19b9b89650f11af77d33e67a544efec90bd8ed3a7cf16"
	base64Encrypt, err := utils.AesEcp.AesBase64Encrypt(str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("====== base64Encrypt: %v", base64Encrypt)
}

func TestAesBase64Decrypt(t *testing.T) {
	str := "1s+K3y8Iz4HoT2a96nKc/jyMxwHXp7qMaFS4YzHyTcshf4L/c8MF6qQ1uEV8pyUq90z3Tm/GIF9EHpte41eetLmZmBSTW+/SKlYUbfgSWhU="
	base64Decrypt, err := utils.AesEcp.AesBase64Decrypt(str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("====== base64Decrypt: %v", base64Decrypt)
}

func TestMd4(t *testing.T) {
	fmt.Println(utils.Md4("123456"))
}

func TestMd5(t *testing.T) {
	fmt.Println(utils.Md5("123456"))
}

func TestSha256(t *testing.T) {
	fmt.Println(utils.Sha256("123456"))
}

func TestSha512(t *testing.T) {
	fmt.Println(utils.Sha512("123456"))
}

func TestFileMd5(t *testing.T) {
	fileMd5, err := utils.FileMd5("./a.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fileMd5)
}
