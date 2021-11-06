package util

import (
	"reflect"
	"unsafe"
)

type SelectType struct {
	Chan    chan unsafe.Pointer                                 //用于协程间通信的channel
	Para    unsafe.Pointer                                      //启动外部协程传递的参数,可为nil
	Routine func(ch chan<- unsafe.Pointer, para unsafe.Pointer) //启动协程的函数,如果为nil,需要自己手动外部启动协程
	OnValid func(id int, value unsafe.Pointer)                  //收到channel回复的处理函数
}

func MuiltpleSelect(eles []SelectType) {
	cases := make([]reflect.SelectCase, len(eles))
	for i, ele := range eles {
		if ele.Routine != nil {
			go ele.Routine(ele.Chan, ele.Para)
		}
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ele.Chan)}
	}

	remaining := len(cases)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			//channel已经关闭,置空此channel
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining -= 1
			continue
		}
		cb := eles[chosen].OnValid
		if cb != nil {
			cb(chosen, unsafe.Pointer(value.Pointer()))
		}
	}
}

type ISelectType interface {
	GetChannel() chan interface{}        //协程间通信的channel
	Routine()                            //启动协程的函数
	OnValid(id int, value reflect.Value) //收到channel回复的处理函数
}

func MuiltpleISelect(eles []ISelectType) {
	cases := make([]reflect.SelectCase, len(eles))
	for i, ele := range eles {
		go ele.Routine()
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ele.GetChannel())}
	}

	remaining := len(cases)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			//channel已经关闭,置空此channel
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining -= 1
			continue
		}
		eles[chosen].OnValid(chosen, value)
	}
}
