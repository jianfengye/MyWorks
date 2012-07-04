package main

import (
    "github.com/lxn/walk"
    "strconv"
    "net"
    "fmt"
    "strings"
    "time"
)
const breakChars = "    "

type TalkWindow struct {
	*walk.MainWindow
	ShowText *walk.TextEdit
	SendText *walk.TextEdit
	ListenPort int
	SendPort int
}

func NewTalkWindow(mv *MainWindow, listenPort int, sendPort int) {
    walk.Initialize(walk.InitParams{PanicOnError: true})
    defer walk.Shutdown()
    
    talkWnd, err := walk.NewMainWindow()
    if err != nil {
        return
    }
    
    tw := &TalkWindow{MainWindow: talkWnd, ListenPort: listenPort, SendPort: sendPort}
    
	tw.SetTitle("I'm listing in" + strconv.Itoa(tw.ListenPort))
	
	tw.ShowText, _ = walk.NewTextEdit(tw)
	tw.ShowText.SetX(10)
	tw.ShowText.SetY(10)
	tw.ShowText.SetWidth(280)
	tw.ShowText.SetHeight(300)
	tw.ShowText.SetReadOnly(true)
	
	tw.SendText, _ = walk.NewTextEdit(tw)
	tw.SendText.SetX(10)
	tw.SendText.SetY(320)
	tw.SendText.SetWidth(200)
	tw.SendText.SetHeight(30)
	
	button1, _ := walk.NewPushButton(tw)
	button1.SetText("发送")
	button1.SetX(220)
	button1.SetY(320)
	button1.SetWidth(70)
	button1.SetHeight(30)
	button1.Clicked().Attach(func() {
	    tw.Send()
	})
    
	tw.SetSize(walk.Size{320, 400})
	tw.Show()

    go tw.Listen()
	tw.Run()
}

func (this *TalkWindow)Send() error {
    txt := this.SendText.Text()
    conn, err := net.Dial("tcp", "localhost:" + strconv.Itoa(this.SendPort))
    if err != nil {
    	return err
    }
    
    lenth := len([]byte(txt))
    pre := Int32ToStream(int32(lenth),BigEndian)
    
    fmt.Fprintf(conn, string(pre) + txt)
    this.SendText.SetText("")
    return nil
}

func (this *TalkWindow)Listen() error {
    ln, err := net.Listen("tcp", ":" + strconv.Itoa(this.ListenPort))
    if err != nil {
    	return err
    }
    for {
    	conn, err := ln.Accept()
    	if err != nil {
    		continue
    	}
    	go func(){
    	    buffer := make([]byte, 4)
    	    conn.Read(buffer)
    	    lenth := StreamToInt32(buffer, BigEndian)
    	    
    	    contentBuf := make([]byte, lenth)
    	    conn.Read(contentBuf)
    	    
    	    text := strings.TrimSpace(string(contentBuf))
    	    fmt.Println(text)
    	    this.ShowText.SetText(this.ShowText.Text() + time.Now().Format("2006-01-02 10:13:40") + breakChars + strconv.Itoa(this.SendPort) + ":" + text + "\r\n")
    	}()
    }
    return nil
}
