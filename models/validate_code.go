package models

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//GenValidateCode 生成随机N位数验证码
func GenValidateCode(width int, username string) (string, error) {
	//随机验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}

	//存储redis
	if err := saveRedis(username, sb.String()); err != nil {
		return "", err
	}

	return sb.String(), nil
}

//存储验证码到redis
func saveRedis(username, code string) error {
	redisConfigStr := ""
	//redisConfigStr := golib.GetConfigSafeString("sessionProviderConfig", "")
	redisConfig := strings.Split(redisConfigStr, ",")
	if len(redisConfig) == 2 {
		redisConfig = append(redisConfig, "")
	}
	//redis, err := golib.NewRedisByAddress(redisConfig[0], redisConfig[2])
	//if err != nil {
	//	return err
	//}
	//
	//key := fmt.Sprintf("relay_code_user_%s", username)
	//return redis.SetWithExpire(key, code, 60*5)

	return nil
}

//VerifyValidateCode 校验验证码
func VerifyValidateCode(username, code string) (bool, error) {
	if len(code) == 0 {
		return false, fmt.Errorf("验证码为空!")
	}
	if len(username) == 0 {
		return false, fmt.Errorf("账户名为空!")
	}

	//redisConfigStr := golib.GetConfigSafeString("sessionProviderConfig", "")
	//redisConfig := strings.Split(redisConfigStr, ",")
	//if len(redisConfig) == 2 {
	//	redisConfig = append(redisConfig, "")
	//}
	//redis, err := golib.NewRedisByAddress(redisConfig[0], redisConfig[2])
	//if err != nil {
	//	return false, err
	//}
	//
	//key := fmt.Sprintf("relay_code_user_%s", username)
	//realCode, err := redis.Get(key)
	//if err != nil {
	//	return false, err
	//}
	//
	//if code == realCode {
	//	return true, nil
	//}

	return false, nil
}
