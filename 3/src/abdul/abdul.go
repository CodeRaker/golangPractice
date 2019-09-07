package abdul

import "time"

func Abdulize(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- ("Abdul\n")
		time.Sleep(1 * time.Second)
	}
}