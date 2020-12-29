package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func DoErr(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func WriteJSON(path string,v interface{}) error {
	outData,err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, outData, 0777)
	if err != nil {
		return err
	}
	return nil
}

func ReadJSON(path string,v interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf,v)
	if err != nil {
		return err
	}
	return nil
}

func DecodeReader(r io.ReadCloser,v interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf,v)
	if err != nil {
		return err
	}
	return nil
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func StrBase64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func ByteBase64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func StrBase64Decode(str string) (string,error) {
	resBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(resBytes),nil
}

