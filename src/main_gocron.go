/**
 * @author zhaiyuanji
 * @date 2021年11月17日 2:31 下午
 */
package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func main()  {
	s := gocron.NewScheduler(time.UTC)

	counter := 0
	s.Every(1).Days().At("08:00").Do(func() {
		counter = counter + 1
		fmt.Println("time: ", counter)
	})
	s.StartBlocking()
}
