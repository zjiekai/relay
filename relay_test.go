package main_test

import (
	"bytes"
	"testing"

	"github.com/golang/glog"
)

// func TestGo(t *testing.T) {
// 	var c = make(chan struct{})
// 	var o sync.Once
//
// 	fn := func(a int) {
// 		var t int
// 		t = a
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(t)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 		o.Do(func() {
// 			close(c)
// 		})
// 	}
//
// 	go fn(3)
// 	go fn(5)
// 	<-c
// }

func TestGo2(t *testing.T) {
	output := &bytes.Buffer{}
	var _ = output
}

func TestCipher(t *testing.T) {

}

func TestGlog(t *testing.T) {
	//$ go test -v -logtostderr
	glog.Info("zjk")
}
