package tray

import (
	"fmt"
	"testing"
)

//go:generate rsrc -arch amd64 -manifest app.manifest -o rsrc.syso -ico logo.ico
func TestNew(t *testing.T) {
	// 创建托盘
	tray, err := New(&Options{
		// 提示信息
		Tip: "test",
		// 托盘logo
		Icon: "./logo.ico",
		// 托盘点击事件
		Click: func(x, y int) {
			fmt.Println("x:", x, "y:", y)
		},
		// 托盘右键菜单
		Menus: Menus{
			{
				Name: "test",
				Action: func(tray *Tray) {
					fmt.Println("click test")
				},
			},
			{
				Name: "exit",
				Action: func(tray *Tray) {
					_ = tray.Stop()
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// 托盘信息
	tray.Message("running", "is running")

	// 启动托盘
	tray.Run()
}
