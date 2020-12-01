/**
 *@Description
 *@ClassName log
 *@Date 2020/11/2 11:02 下午
 *@Author ckhero
 */

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"

	//"github.com/json-iterator/go/extra"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type StdStruct struct {
	Age int `json:"age"`
}

type StdStruct2 struct {
	Age string `json:"age"`
}
func main() {
	fmt.Println(hashPassword("18801613198"))
	// Using the default options
	salt, encodedPwd := password.Encode("generic password", nil)
	fmt.Println(encodedPwd)
	check := password.Verify("generic password", salt, encodedPwd, nil)
	fmt.Println(check) // true

	// Using custom options
	options := &password.Options{10, 10000, 50, md5.New}
	fmt.Println(encodedPwd)

	salt, encodedPwd = password.Encode("generic password", options)
	check = password.Verify("generic password", salt, encodedPwd, options)
	fmt.Println(check) // true
}

func hashPassword(passwordString string) (string, string) {

	return password.Encode(passwordString, getPasswordOptions())
}
const saltLen = 10
const iterations = 10000
const keyLen = 50
const DefaultExpireTime = 10 // 十分钟

func getPasswordOptions() *password.Options {
	return &password.Options{
		SaltLen:      saltLen,
		Iterations:   iterations,
		KeyLen:       keyLen,
		HashFunction: sha1.New,
	}
}
func DeepCopyPHP(src, dst interface{}) error {
	//data ,_ := json.Marshal(src)
	//return json.Unmarshal(data, dst)

	extra.RegisterFuzzyDecoders()
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(b, dst)
}