package bizid

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
生成形如: 180102981234567812,
0-5:  	YYMMDD
6-7:  	CountrySN  国家编码
8-9:    BizSN
10-15:	device seq id
16-17:  毫秒数最后2位
*/

// 国家
type CountrySN int

// 业务
type BizSN int

const (
	EPAYSN CountrySN = 100
)

const (
	AppSN    BizSN = 1 // app项目编号
	ManageSN BizSN = 2 // 后台项目编号
)

func BizID(bizSN BizSN) int64 {
	countrySN := EPAYSN

	t := time.Now()
	nanos := t.UnixNano()
	millis := nanos / 1000000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Intn(10000000)

	bizIDStr := fmt.Sprintf("%d%02d%02d%02d%02d%07d%01d", t.Year()%100, t.Month(), t.Day(), countrySN, bizSN, id%10000000, millis%10)
	bizID, _ := strconv.ParseInt(bizIDStr, 10, 64)
	return bizID
}
