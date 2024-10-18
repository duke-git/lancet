//go:build windows

package fileutil

import (
	"fmt"
	"syscall"
	"unsafe"
)

// tagVS_FIXEDFILEINFO 参考结构体https://learn.microsoft.com/zh-cn/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type tagVS_FIXEDFILEINFO struct {
	Signature        uint32
	StructVersion    uint32
	FileVersionMS    uint32
	FileVersionLS    uint32
	ProductVersionMS uint32
	ProductVersionLS uint32
	FileFlagsMask    uint32
	FileFlags        uint32
	FileOS           uint32
	FileType         uint32
	FileSubtype      uint32
	FileDateMS       uint32
	FileDateLS       uint32
}

// GetExeOrDllVersion get the version of exe or dll file on windows.
// Play: todo
func GetExeOrDllVersion(filePath string) (string, error) {
	// 加载系统dll
	versionDLL := syscall.NewLazyDLL("version.dll")
	getFileVersionInfoSize := versionDLL.NewProc("GetFileVersionInfoSizeW")
	getFileVersionInfo := versionDLL.NewProc("GetFileVersionInfoW")
	verQueryValue := versionDLL.NewProc("VerQueryValueW")

	// 转换路径为UTF-16
	filePathPtr, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to convert file path to UTF-16: %w", err)
	}

	// 获取version信息大小
	size, _, err := getFileVersionInfoSize.Call(
		uintptr(unsafe.Pointer(filePathPtr)),
		0,
	)
	if size == 0 {
		return "", fmt.Errorf("unable to obtain version information size: %v", err)
	}

	// 加载version信息
	data := make([]byte, size)
	ret, _, err := getFileVersionInfo.Call(uintptr(unsafe.Pointer(filePathPtr)), 0, size, uintptr(unsafe.Pointer(&data[0])))
	if ret == 0 {
		return "", fmt.Errorf("unable to obtain version information: %v", err)
	}

	// 查询version信息
	var fixedInfo *tagVS_FIXEDFILEINFO
	var fixedInfoLen uint32
	u16, err := syscall.UTF16PtrFromString(`\`)
	if err != nil {
		return "", fmt.Errorf("unable to convert file path to UTF-16: %w", err)
	}
	ret, _, err = verQueryValue.Call(
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(unsafe.Pointer(u16)),
		uintptr(unsafe.Pointer(&fixedInfo)),
		uintptr(unsafe.Pointer(&fixedInfoLen)),
	)
	if ret == 0 {
		return "", fmt.Errorf("unable to query version information: %v", err)
	}

	// 转换结构体
	major := fixedInfo.FileVersionMS >> 16
	minor := fixedInfo.FileVersionMS & 0xFFFF
	build := fixedInfo.FileVersionLS >> 16
	revision := fixedInfo.FileVersionLS & 0xFFFF

	return fmt.Sprintf("%d.%d.%d.%d", major, minor, build, revision), nil
}
