//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"github.com/csxuejin/sntp/netapp"
	"github.com/csxuejin/sntp/netevent"
)

func main() {
	var handler = netapp.GetHandler()
	netevent.Reactor.ListenUdp(123, handler)
	netevent.Reactor.Run()
}
