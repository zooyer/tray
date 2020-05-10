# Go tray for windows

Go tray for windows.



#### Download and Install

```shell
go get github.com/zooyer/tray
```



#### Generate rsrc.syso(depend app.manifest)

```shell
go get github.com/akavel/rsrc
rsrc -arch amd64 -manifest app.manifest -o rsrc.syso
```



#### Example

```go
//go:generate rsrc -arch amd64 -manifest app.manifest -o rsrc.syso
func main() {
	tray, err := New(&Options{
		Tip: "test",
		Icon: "./logo.ico",
		Click: func(x, y int) {
			fmt.Println("x:", x, "y:", y)
		},
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

	tray.Message("running", "is running")

	tray.Run()
}
```

