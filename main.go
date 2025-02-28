package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var rwMutex sync.RWMutex

func moveRight() {
	// AppleScript 命令
	appleScript := `tell application "System Events" to key code 124 using control down`

	// 执行命令
	cmd := exec.Command("osascript", "-e", appleScript)
	cmd.Run()
}
func moveLeft() {
	// AppleScript 命令
	appleScript := `tell application "System Events" to key code 123 using control down`
	// 执行命令
	cmd := exec.Command("osascript", "-e", appleScript)
	cmd.Run()
}
func moveUp() {
	// AppleScript 命令
	appleScript := `tell application "System Events" to key code 126 using control down`
	// 执行命令
	cmd := exec.Command("osascript", "-e", appleScript)
	cmd.Run()
}
func moveDown() {
	// AppleScript 命令
	appleScript := `tell application "System Events" to key code 125 using control down`
	// 执行命令
	cmd := exec.Command("osascript", "-e", appleScript)
	cmd.Run()
}
func main() {
	var (
		startX, startY int // 记录中键按下时的坐标
	)
	screenRect := robotgo.GetScreenRect()
	fmt.Println("screenRect", screenRect)
	horizontalMinDiff := screenRect.W / 20
	verticalMinDiff := screenRect.H / 20
	for {
		ok := hook.AddEvent("center")
		if ok {
			startX, startY = robotgo.Location()
			var x, y int
			var timeout = false
			s := hook.Start()
			var check = func() {
				rwMutex.RLock()
				if !timeout {
					if x-startX > horizontalMinDiff {
						fmt.Println("move right")
						moveRight()
						hook.End()
					} else if startX-x > horizontalMinDiff {
						fmt.Println("move left")
						moveLeft()
						hook.End()
					} else if y-startY > verticalMinDiff {
						fmt.Println("move down")
						moveDown()
						hook.End()
					} else if startY-y > verticalMinDiff {
						fmt.Println("move up")
						moveUp()
						hook.End()
					}
				} else {
					fmt.Println("timeout")
					hook.End()
				}
				rwMutex.RUnlock()
			}
			go func() {
				time.Sleep(time.Millisecond * 2000)
				rwMutex.Lock()
				timeout = true
				rwMutex.Unlock()
			}()
			hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
				x = int(e.X)
				y = int(e.Y)
				check()
			})
			<-hook.Process(s)
		}
	}
}
