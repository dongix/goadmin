package service

import (
	"connection"
	"fmt"
	"log"
)

// GenSMSCode 生成验证码
func GenSMSCode(mobile string) error {
	client, error := connection.GetRedisClient()
	defer client.Close()
	if error != nil {
		log.Fatal(error.Error())
	}
	code := fmt.Sprintf("sms:%s", mobile)
	err := client.Set(code, "123456", -0).Err()
	return err
}
