package main

import(
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"strings"
)

func main() {
	// 监听websocket端口
	http.Handle("/", websocket.Handler(ActionRoute))
	go http.ListenAndServe(":3201", nil)
	select {}
}

func ActionRoute(ws *websocket.Conn) {
	log.Println("connect a new conn")
	var err error
	for {
		var receive string
		// 收到消息：{"action":"startItuns", "params" : ""}
		if err = websocket.Message.Receive(ws, &receive); err != nil {
			log.Println("receive err:", err)
			continue
		}

		type Command struct {
			Action string `json:"action"`
			Params string `json:"params"`
		}

		command := new(Command)
		if err = json.Unmarshal([]byte(receive), command); err != nil {
			log.Println("receive message parse err, receive:", receive)
			continue
		}

		// 更多osascript可参考 http://hints.macworld.com/article.php?story=20011108211802830
		log.Println(command.Action)
		switch command.Action {
		case "play":
			// 执行osascript启动
			//osascript -e 'tell application "iTunes" to play';
			Run("", "osascript", "-e", `tell application "iTunes" to play`)
		case "volup":
			// 开大声音
			// 获取当前音量：osascript -e 'tell application "iTunes" to sound volume as integer'
			vol, err := Run("", "osascript", "-e", `tell application "iTunes" to sound volume as integer`)
			if err != nil {
				log.Println("get volumn err:", err)
				continue
			}
			volint, err := strconv.Atoi(strings.Trim(string(vol),"\n"))
			if err != nil {
				log.Println("get volumn err:", err)
				continue
			}
			volint = volint + 10
			if volint > 100 {
				volint = 100
			}
			Run("", "osascript", "-e", `tell application "iTunes" to set sound volume to ` + strconv.Itoa(volint))
			// osascript -e "tell application \"iTunes\" to set sound volume to $newvol";
		case "voldown":
			// 降低声音
			// 获取当前音量：osascript -e 'tell application "iTunes" to sound volume as integer'
			// osascript -e "tell application \"iTunes\" to set sound volume to $newvol";
			vol, err := Run("", "osascript", "-e", `tell application "iTunes" to sound volume as integer`)
			if err != nil {
				log.Println("get volumn err:", err)
				continue
			}
			volint, err := strconv.Atoi(strings.Trim(string(vol),"\n"))
			if err != nil {
				log.Println("get volumn err:", err)
				continue
			}
			volint = volint - 10
			if volint < 0 {
				volint = 0
			}
			Run("", "osascript", "-e", `tell application "iTunes" to set sound volume to ` + strconv.Itoa(volint))
		case "stop":
			// 关闭ituns
			// osascript -e 'tell application "iTunes" to stop';
			Run("", "osascript", "-e", `tell application "iTunes" to stop`)
		case "prev":
			// 上一首
			// osascript -e 'tell application "iTunes" to previous track';
			Run("", "osascript", "-e", `tell application "iTunes" to previous track`)
		case "next":
			// 下一首
			// osascript -e 'tell application "iTunes" to next track';
			Run("", "osascript", "-e", `tell application "iTunes" to next track`)
		}
	}
}