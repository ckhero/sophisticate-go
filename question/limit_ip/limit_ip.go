package limit_ip

import (
	"fmt"
	"sync"
	"time"
)

/**
在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success: 100。
 */

type LimitIp struct {
	m sync.Map
}
func(l *LimitIp) IsInvalid(user string) bool {
	v, ok := l.m.LoadOrStore(user, time.Now())
	if !ok {
		fmt.Printf("名单里没有，正常访问1 - [%s] \n", user)
		return true
	}
	if time.Now().Second() - v.(time.Time).Second() >= 2 {
		fmt.Printf("名单里没有，正常访问2 - [%s] \n", user)
		l.m.Store(user, time.Now())
		return true
	}
	fmt.Printf("time, %d \n", v.(time.Time).Second())
	fmt.Printf("名单里有，拒绝访问 - [%s] \n", user)
	return false
}