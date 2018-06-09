package main

import (
	"github.com/sinxsoft/three-door/app"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	okTimes := 0
	for t := 0; t < 10000; t++ {
		var d = [3]app.Door{}
		//fmt.Print(d)
		RandInitDoors(&d)
		//随机选一个门的序号：0，1，2中选一个
		rand := rand.Intn(time.Now().Nanosecond())
		//确定选中
		rand = rand % 3
		//判断是否选中

		//选中
		if d[rand].GetType() == 1 {
			//主持人有两个可能，随便开一扇门
			//嘉宾必输，不用增加okTimes
		} else {
			//没选中的话，主持人则只有一扇门可以开，嘉宾必赢
			okTimes++
			//结束
		}
	}
	fmt.Print("数字：" +string(okTimes))
}

func  RandInitDoors( d *[3]app.Door) {
	rand := rand.Intn(time.Now().Nanosecond())
	rand = rand % 3

	if !d[rand].IsInit {
		d[rand].IsInit = true
		d[rand].SetType(1)
	}

	for x := 0; x < 3; x++ {
		if !d[x].IsInit {
			d[x].IsInit = true
			d[x].SetType(2)
		}
	}
}
