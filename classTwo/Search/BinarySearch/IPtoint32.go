/*
将ip地址转为int32类型整数
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ans int32 = 0
	ip := "127.0.0.1"
	ips := strings.Split(ip, ".")
	fmt.Println(ips)
	for i := len(ips) - 1; i >= 0; i-- {
		temp, err := strconv.ParseInt(ips[i], 10, 32) //将string转化为int32
		if temp > 255 || temp < 0 {
			fmt.Println("ip is err")
			return
		}
		if err != nil {
			fmt.Println("strconv.ParseInt err")
			return
		}
		ans = ans | (int32(temp) << (8 * i)) // 把每个8位放在合适位置，二进制形式组合起来
	}
	fmt.Println(ip, "ip to int32:", ans)
}
