package monitor

import "github.com/golang/glog"

var openPipe chan struct{}
var closePipe chan struct{}

func Init() {
  openPipe = make(chan struct{})
  closePipe = make(chan struct{})
  go monitor()
}

func monitor() {
  pipeCount := 0
  for {
    select {
    case <-openPipe:
      pipeCount++
      glog.Infof("a new pipe opened. %v pipe(s) open", pipeCount)
    case <-closePipe:
      pipeCount--
      glog.Infof("an old pipe closed. %v pipe(s) open", pipeCount)
    }
  }
}

func PipeOpen() {
  openPipe <- struct{}{}
}

func PipeClose() {
  closePipe <- struct{}{}
}
