package pkg

import (
	"syscall"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

//Load Windows DLL File, user32.dll
var libuser32 *syscall.LazyDLL = syscall.NewLazyDLL("user32.dll")

//Load function from winuser.h
//hWnd indicates layered window, created by specifying WS_EX_LAYERED
//When creating the window with the CreateWindowEx function or by setting WS_EX_LAYERED via SetWindowLong after the window has been created.
//I called the function SetWindowLong to set window WS_EX_LAYERED
//A crColor, COLORREF structure, that specifies the transparency color key to be used when composing the layered window. All pixels painted by the window in this color will be transparent.
//The type of crColor in Go Walk is walk.Color.
//bAlpha value used to describe the opacity of the layered window.
//If bAlpha is 0, the window will be completly transparent. So I have not used this now.
//dwFlags is the action to be taken.
//In 0x0000001, Use crColor as the transparency color.
func SetLayeredWindowAttributes(hWnd win.HWND, crColor walk.Color, bAlpha byte, dwFlags int32) bool {
	setLayeredWindowAttributes := libuser32.NewProc("SetLayeredWindowAttributes")
	r, _, _ := syscall.Syscall6(setLayeredWindowAttributes.Addr(), 4, uintptr(hWnd), uintptr(crColor), uintptr(bAlpha), uintptr(dwFlags), 0, 0)
	if r == 0 {
		return false
	}
	return true
}
