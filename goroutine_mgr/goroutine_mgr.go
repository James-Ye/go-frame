package goroutine_mgr

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func GetCurrentGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

type goChannel struct {
	Control chan bool
	Params  uintptr
}

type Goroutinemanager struct {
	mutex      sync.Mutex
	grchannels map[int64]*goChannel
}

func (gm *Goroutinemanager) GetChannels() *map[int64]*goChannel {
	return &gm.grchannels
}

func (gm *Goroutinemanager) Regist(goid int64) error {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	ch := new(goChannel)

	ch.Control = make(chan bool)
	if gm.grchannels == nil {
		gm.grchannels = make(map[int64]*goChannel)
	} else if _, ok := gm.grchannels[goid]; ok {
		return fmt.Errorf("goroutine channel already defined")
	}

	gm.grchannels[goid] = ch
	return nil
}

func (gm *Goroutinemanager) IsNeedCleanup(goid int64) bool {
	needcleanup := false
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	select {
	case ch := <-gm.grchannels[goid].Control:
		needcleanup = ch
		delete(gm.grchannels, goid)
	default:
		fmt.Println("no signal")
	}

	return needcleanup
}

func (gm *Goroutinemanager) Size() int {
	return len(gm.grchannels)
}

func (gm *Goroutinemanager) Stop(goid *int64) {
	if goid != nil {
		if channel, ok := gm.grchannels[*goid]; ok {
			channel.Control <- true
		}
	} else {
		for _, v := range gm.grchannels {
			v.Control <- true
		}
	}
}
