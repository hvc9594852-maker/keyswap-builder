package main

import (
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const (
	en = "1234567890-=qwertyuiop[]\\asdfghjkl;'zxcvbnm,./!@#$%^&*()_+QWERTYUIOP{}|ASDFGHJKL:\"ZXCVBNM<>?"
	th = "ๅ/_ภถุึคตจขชๆไำพะัีรนยบลฃฟหกดเ้่าสวงผปแอิืทมใฝ+\"!#ฯ%&฿()_+-๑๒๓๔ู฿๕๖๗๘๙ฤฆฏโฌพ็โรศษศซ.()ฉฮฺ์?ฒฬฦ"
)

func convert(text string) string {
	var sb strings.Builder
	for _, r := range text {
		char := string(r)
		idxEn := strings.Index(en, char)
		idxTh := strings.Index(th, char)

		if idxEn != -1 {
			sb.WriteString(string([]rune(th)[idxEn]))
		} else if idxTh != -1 {
			sb.WriteString(string([]rune(en)[idxTh]))
		} else {
			sb.WriteString(char)
		}
	}
	return sb.String()
}

func main() {
	// ตั้งคีย์ลัด Ctrl + Shift + T
	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "t"}, func(e hook.Event) {
		oldClip, _ := clipboard.ReadAll()
		robotgo.KeyTap("c", "control")
		time.Sleep(100 * time.Millisecond)

		text, err := clipboard.ReadAll()
		if err == nil && text != "" {
			converted := convert(text)
			clipboard.WriteAll(converted)
			robotgo.KeyTap("v", "control")
			time.Sleep(100 * time.Millisecond)
		}
		clipboard.WriteAll(oldClip)
	})

	s := hook.Start()
	<-hook.Process(s)
}
