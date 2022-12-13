package scribeencryptdecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"

	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)


func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func deriveeKey(passphrase []byte, salt []byte) []byte {

	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha1.New)
}
func AESDecrypt(encryptedString string, sPassphrase string, saltString string) string {
	crypt, _ := base64.StdEncoding.DecodeString(encryptedString[24:])
	initialVector, _ := base64.StdEncoding.DecodeString(encryptedString[:24])
	//fmt.Println(encryptedString[24:])
	//fmt.Println(encryptedString[:24])

	passphrase := []byte(sPassphrase)
	salt := []byte(saltString)
	key := deriveeKey(passphrase, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, initialVector)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)
	return string(PKCS5Trimming(decrypted))
}

func init() {
	_ = function.Register(&aesdecrypt{})
}

type aesdecrypt struct {
}

func (s *aesdecrypt) Name() string {
	return "aesdecrypt"
}

func (aesdecrypt) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString,data.TypeString,data.TypeString}, false
}

func (aesdecrypt) Eval(params ...interface{}) (interface{}, error) {
	
	encryptedText := params[0].(string)
	passphrase := params[1].(string)
	saltstring:=params[2].(string)	

	decryptedString := AESDecrypt(encryptedText, passphrase, saltstring)
	
	
	return decryptedString,nil;


}