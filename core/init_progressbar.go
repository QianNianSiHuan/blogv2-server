package core

import (
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var ProgressbarMsg = make(chan string)
var ProgressbarEndMsg = make(chan bool)

type ChanWriter struct {
	ch chan string
}

func (w ChanWriter) Write(p []byte) (n int, err error) {
	msg := string(p)
	w.ch <- msg
	return len(p), nil
}
func InitProgressbar(maxNumber int) {
	var ProgressbarLogMsg = make(chan string)
	var logOutput = ChanWriter{ProgressbarLogMsg}
	logrus.SetOutput(logOutput)
	bar := progressbar.NewOptions(15,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()), //you should install "github.com/k0kubun/go-ansi"
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(25),
		progressbar.OptionSetDescription("[red]开始读取配置...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	count := 0
	go func() {
		for {
			select {
			case str := <-ProgressbarMsg:
				bar.Add(1)
				count += 1
				bar.Describe(fmt.Sprintf("[red]%-10s[reset]", str))
				time.Sleep(300 * time.Millisecond)
				continue
			}
		}
	}()

	go func() {
		for {
			select {
			case str := <-ProgressbarLogMsg:
				switch str {
				case "":
				default:
					bar.Add(1)
					count += 1
					progressbar.Bprintf(bar, str)
					continue
				}
			}
		}
	}()

	defer func() {
		close(ProgressbarMsg)
		close(ProgressbarLogMsg)
		close(ProgressbarEndMsg)
	}()

	select {
	case <-ProgressbarEndMsg:
		bar.Describe(fmt.Sprintf("%-10s", "[red]系统加载成功！！![reset]"))
		logrus.SetOutput(os.Stdout)
		bar.Close()
		return
	}
}
