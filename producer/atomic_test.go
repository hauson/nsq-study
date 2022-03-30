package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)


// The swap operation, implemented by the SwapT functions, is the atomic
// equivalent of:
//
//	old = *addr
//	*addr = new
//	return old
//
// The compare-and-swap operation, implemented by the CompareAndSwapT
// functions, is the atomic equivalent of:
//
//	if *addr == old {
//		*addr = new
//		return true
//	}
//	return false
//
// The add operation, implemented by the AddT functions, is the atomic
// equivalent of:
//
//	*addr += delta
//	return *addr
//
// The load and store operations, implemented by the LoadT and StoreT
// functions, are the atomic equivalents of "return *addr" and
// "*addr = val".
//

func Test_CompareAndSwapInt32(t *testing.T) {
	/*
	addr == old: 为真，替换为new值
	addr != old: 为假， 不替换
	 */

	var addr int32 = 0
	if atomic.CompareAndSwapInt32(&addr, 0, 1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	fmt.Println(addr)
}

func Test_Load(t *testing.T) {
	var addr int32 = 7
	value := atomic.LoadInt32(&addr)
	fmt.Println("value", value)
	fmt.Println("addr", addr)
}