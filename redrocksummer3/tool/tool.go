package tool

import (
	"fmt"
	"log"
	"strconv"
	"strings"

)

func HandleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Chuli(body string) (string, string, int, string) {
	fmt.Println("s++", body)
	str  := strings.Split(body,`"`)
	//数字处理
	num       := strings.Split(str[22],":")
	num_id    := strings.Split(num[1],",")
	num0, err := strconv.Atoi(num_id[0])
	if err != nil {
		fmt.Println("err:", err)
	}

	//时间处理
	//times, _ := time.ParseInLocation("2006-01-02", str[25], time.Local)

	return str[15], str[19], num0, str[25]
}

