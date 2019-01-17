package helper

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"
)

const (
	OperateTypeCreate = 1
	OperateTypeEdit   = 2
)

func GetRandomStrBy(strLen uint) string {
	return GetRandomStr(strLen, "")
}

func GetRandomStr(strLen uint, str string) string {
	if str == "" {
		str = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	bytes := []byte(str)
	var result []byte
	var i uint
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(str)
	for i = 0; i < strLen; i++ {
		result = append(result, bytes[r.Intn(length)])
	}
	return string(result)
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func NowHour() string {
	return time.Now().Format("15")
}

func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func CreateNewError(msg string) error {
	return errors.New(msg)
}
