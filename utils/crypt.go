package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetRandKey() string {
	rand.Seed(time.Now().UnixNano())
	s := rand.Int31n(10000)
	return fmt.Sprintf("%04d", s)
}

func GetUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7
	return hex.EncodeToString(uuid), nil
}

func GetLiveRandom() string {
	//rand.Seed(time.Now().UnixNano())
	//s := rand.Int31n(100000000)
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
	//return fmt.Sprintf("%08v", s)
}
func HmacSha256(key, data string) string {
	fmt.Println("HmacSha256 key")
	fmt.Println(key)
	fmt.Println("HmacSha256 data")
	fmt.Println(data)

	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data))
	fmt.Println(hex.EncodeToString(hash.Sum([]byte(""))))
	return hex.EncodeToString(hash.Sum([]byte("")))
}
