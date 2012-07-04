package main

import (
    "github.com/lxn/walk"
)

type MainWindow struct {
	*walk.MainWindow
}


func main() {
    walk.Initialize(walk.InitParams{PanicOnError: true})
    defer walk.Shutdown()
    
    mainWnd, err := walk.NewMainWindow()
    if err != nil {
        return
    }
    
    mw := &MainWindow{MainWindow: mainWnd}
    
	mw.SetTitle("SocketIm Example")
	
	button1, _ := walk.NewPushButton(mw)
	button1.SetText("start port 8000")
	button1.SetX(10)
	button1.SetY(10)
	button1.SetWidth(100)
	button1.SetHeight(30)
		
	button1.Clicked().Attach(func() {
		go NewTalkWindow(mw, 8000, 8001)
		button1.SetEnabled(false)
	})
	
	button2, _ := walk.NewPushButton(mw)
	button2.SetText("start port 8001")
	button2.SetX(10)
	button2.SetY(60)
	button2.SetWidth(100)
	button2.SetHeight(30)
	
	button2.Clicked().Attach(func() {
		go NewTalkWindow(mw, 8001, 8000)
		button2.SetEnabled(false)
	})
    
	mw.SetSize(walk.Size{120, 150})
	mw.Show()

	mw.Run()
}
