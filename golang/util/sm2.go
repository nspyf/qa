
package util

import (
	"errors"
	"github.com/tjfoc/gmsm/sm2"
	"strings"
)

type SM2obj struct {
	// 加密钥匙和解密钥匙均可作为公、私钥，根据用途决定
	encryptKey string // 用于加密
	decryptKey string // 用于解密

	pair *sm2.PrivateKey // 初始化后才能使用
}

 // 产生一对密钥，并初始化
func (obj *SM2obj) Generate() (err error) {
	obj.pair, err = sm2.GenerateKey()
	if err != nil {
		return err
	}

	//fmt.Println(obj.pair.D.String())

	obj.encryptKey = StrBase64Encode(obj.pair.PublicKey.X.String()+"."+obj.pair.PublicKey.Y.String())
	obj.decryptKey = StrBase64Encode(obj.pair.D.String())

	return nil
}

 // 初始化EncryptKey
func (obj *SM2obj) WriteEncryptKey(encryptKey string) error {
	encryptKeyStr, err := StrBase64Decode(encryptKey)
	if err != nil {
		return err
	}

	encryptKeyStrArr := strings.Split(encryptKeyStr,".")
	if len(encryptKeyStrArr) != 2 {
		return errors.New("encryptKey Error")
	}

	obj.pair.PublicKey.X.SetString(encryptKeyStrArr[0], 10)
	obj.pair.PublicKey.Y.SetString(encryptKeyStrArr[1], 10)

	return nil
}

// 初始化DecryptKey
func (obj *SM2obj) WriteDecryptKey(decryptKey string) error {
	decryptKeyStr, err := StrBase64Decode(decryptKey)
	if err != nil {
		return err
	}

	obj.pair.PublicKey.Y.SetString(decryptKeyStr, 10)

	return nil
}

func (obj *SM2obj) ReadEncryptKey() string {
	return obj.encryptKey
}

func (obj *SM2obj) ReadDecryptKey() string {
	return obj.decryptKey
}

// 使用前需WriteEncryptKey
func (obj *SM2obj) Encrypt(data string) (string,error) {
	s,err := obj.pair.Encrypt([]byte(data))
	if err != nil {
		return "", err
	}

	return ByteBase64Encode(s), nil
}

// 使用前需WriteDecryptKey
func (obj *SM2obj) Decrypt(data string) (string,error) {
	s,err := obj.pair.Decrypt([]byte(data))
	if err != nil {
		return "", err
	}

	return ByteBase64Encode(s), nil
}

// 使用前需WriteEncryptKey
func (obj *SM2obj) Sign(data string) (string,error) {
	hashStr := SHA256(data)
	s,err := obj.pair.PublicKey.Encrypt([]byte(hashStr))
	if err != nil {
		return "", err
	}

	return ByteBase64Encode(s), nil
}

// 使用前需WriteDecryptKey
func (obj *SM2obj) Verify(data string,signature string) (bool,error) {
	hashStr := SHA256(data)

	str,err := StrBase64Decode(signature)
	if err != nil {
		return false, err
	}

	deByte,err := obj.pair.Decrypt([]byte(str))
	if err != nil {
		return false, err
	}

	return hashStr == string(deByte),nil
}
