package main

import (
	"embed"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/quavious/go-windows-gadget/pkg"
)

//go:embed assets/*
var assets embed.FS

var window *walk.MainWindow
var windowIcon *walk.Icon

func main() {
	bgColor := walk.RGB(180, 180, 180)

	imageBuf, err := assets.ReadFile("assets/sample.png")
	image := pkg.LoadImage(imageBuf, err)

	iconBuf, err := assets.ReadFile("assets/sample_icon.png")
	icon := pkg.LoadImage(iconBuf, err)

	size := image.Size()
	pictureWidth, pictureHeight := size.Width, size.Height

	err = (MainWindow{
		AssignTo: &window,
		Title:    "Cat",
		Size:     Size{Width: pictureWidth, Height: pictureHeight},
		Layout:   VBox{MarginsZero: true, SpacingZero: true},
		Background: SolidColorBrush{
			Color: bgColor,
		},
		Children: []Widget{
			ImageView{
				Image:  image,
				Mode:   ImageViewModeIdeal,
				Margin: 0,
			},
		},
	}).Create()
	pkg.HandleError(err)

	ic, err := walk.NewIconFromBitmap(icon)
	pkg.HandleError(err)

	err = window.SetIcon(ic)
	pkg.HandleError(err)

	if windowIcon != nil {
		windowIcon.Dispose()
	} else {
		windowIcon = ic
		defer windowIcon.Dispose()
	}

	handle := window.Handle()

	flag := win.GetWindowLong(handle, win.GWL_STYLE)
	flag &= ^win.WS_THICKFRAME
	win.SetWindowLong(handle, win.GWL_STYLE, flag)
	win.SetBkMode(win.HDC(handle), win.TRANSPARENT)

	win.SetWindowLong(handle, win.GWL_EXSTYLE, win.GetWindowLong(handle, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
	//Load function from winuser.h
	//I hand over one crColor and dwFlags value to this method.
	pkg.SetLayeredWindowAttributes(handle, bgColor, 255, 0x00000001)

	win.ShowWindow(handle, win.SW_SHOW)
	window.Run()
}
