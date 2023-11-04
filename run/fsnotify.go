package run

import (
	"cos/util"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"path/filepath"
	"strings"
)

func fsnotifymon(path string) {
	// 创建文件系统监控器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		util.Logger.Error(err.Error())
	}
	defer watcher.Close()
	// 添加需要监控的目录
	err = watcher.Add(path)
	if err != nil {
		util.Logger.Error(err.Error())
	}
	// 开始监听文件更改事件
	for {
		select {
		case event := <-watcher.Events:
			marshal, err := json.Marshal(&event)
			if err != nil {
				util.Logger.Error(err.Error())
				return
			}
			c := color.New()
			util.Logger.Info(fmt.Sprintf("%s %s", c.Add(color.FgHiYellow).Sprint("检测到目录发生动态变化~"), c.Add(color.FgHiBlue).Sprint(string(marshal))))
			/*操作监控上传*/
			suffix := strings.Split(path, "\\")[len(strings.Split(path, "\\"))-1]
			createObjectDir(suffix)
			err = filepath.Walk(path, walkFunc)
			if err != nil {
				util.Logger.Error(err.Error())
				return
			}
			/*end*/
		case err := <-watcher.Errors:
			util.Logger.Error(err.Error())
		}
	}
}
