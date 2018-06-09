package main

import (
	"github.com/sinxsoft/three-door/app"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	okTimes := 0
	//执行一万次，看看有多少次是ok的
	for t := 0; t < 10000; t++ {
		var d = [3]app.Door{}
		//fmt.Print(d)
		RandInitDoors(&d)
		//随机选一个门的序号：0，1，2中选一个
		rand := rand.Intn(time.Now().Nanosecond())
		//确定一个随机door编号
		rand = rand % 3
		//如果嘉宾本来就选中了
		if d[rand].GetType() == 1 {
			//主持人有两个可能，随便开一扇门
			//嘉宾更换选择必输，不用增加okTimes
		} else {
			//没选中的话，主持人则只有一扇门可以开，嘉宾更换后必赢
			okTimes++
		}
	}
	//最后打印10000个结果里面选中car的总和
	fmt.Print("数字：" +string(okTimes))
}

func  RandInitDoors( d *[3]app.Door) {
	rand := rand.Intn(time.Now().Nanosecond())
	rand = rand % 3

	if !d[rand].IsInit {
		//随机找一个door，设置为汽车
		d[rand].IsInit = true
		d[rand].SetType(1)
	}

	for x := 0; x < 3; x++ {
		//除了汽车之外，都设置位羊
		if !d[x].IsInit {
			d[x].IsInit = true
			d[x].SetType(2)
		}
	}
}
