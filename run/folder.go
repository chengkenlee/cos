package run

import (
	"context"
	"cos/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/tencentyun/cos-go-sdk-v5"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type SelfListener struct {
}

// 自定义进度回调需要实现 ProgressChangedCallback 方法
func (l *SelfListener) ProgressChangedCallback(event *cos.ProgressEvent) {
	c := color.New()
	switch event.EventType {
	case cos.ProgressDataEvent:
		fmt.Printf("\r[ConsumedBytes/TotalBytes: %s/%s, %d%%]",
			c.Add(color.FgHiYellow).Sprint(event.ConsumedBytes),
			c.Add(color.FgHiGreen).Sprint(event.TotalBytes),
			event.ConsumedBytes*100/event.TotalBytes,
		)
		//case cos.ProgressFailedEvent:
		//	fmt.Printf("\nTransfer Failed: %v", event.Err)
	}
}

// 多线程批量上传文件
func threadUploadFiles(cosfilename, localfilepath string) {
	/*Case 1 通过默认回调查看上传进度*/
	f, err := os.Open(localfilepath)
	if err != nil {
		util.Logger.Error(err.Error())
	}
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
			// 设置默认的进度回调函数
			Listener: &cos.DefaultProgressListener{},
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
			XCosACL: "private",
		},
	}
	_, err = util.Client.Object.Put(context.Background(), cosfilename, f, opt)
	/*Case 2 通过自定义方式查看上传进度*/
	opt.Listener = &SelfListener{}
	_, err = util.Client.Object.PutFromFile(context.Background(), cosfilename, localfilepath, opt)
}

// 创建文件夹
func createObjectDir(suffix string) {
	// 传递大小为0的输入流
	_, err := util.Client.Object.Put(context.Background(), suffix+"/", strings.NewReader(""), nil)
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
}

// 判断目录或文件
func FolderOrFile() {
	var wait sync.WaitGroup
	for _, s := range util.Config.GetStringSlice("scan.object") {
		wait.Add(1)
		go func(s string) {
			defer wait.Done()
			suffix := strings.Split(s, "\\")[len(strings.Split(s, "\\"))-1]

			b, err := os.Stat(s)
			if err != nil {
				util.Logger.Error(err.Error())
			}
			//判断是否是文件夹类型
			if b.IsDir() {
				//如果是一个文件夹，那么检测cos是否也存在对应文件夹，如不存在，则创建
				createObjectDir(suffix)

				err := filepath.Walk(s, walkFunc)
				if err != nil {
					util.Logger.Error(err.Error())
					return
				}
				return
			}
			c := color.New()
			if checkObjectBucketCosFileIsExist(suffix) {
				util.Logger.Info(fmt.Sprintf("%s:[%s]", c.Add(color.FgHiGreen).Sprint("cos已经存在"), suffix))
			} else {
				util.Logger.Info(fmt.Sprintf("%s:[%s]", c.Add(color.FgHiYellow).Sprint("准备上传文件"), suffix))
				threadUploadFiles(suffix, s)
			}
		}(s)
	}
	wait.Wait()
	util.Logger.Info("done.")
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		dir := strings.Split(path, "\\")[len(strings.Split(path, "\\"))-2]
		filename := strings.Split(path, "\\")[len(strings.Split(path, "\\"))-1]
		cosfile := fmt.Sprintf("/%s/%s", dir, filename)
		c := color.New()
		if checkObjectBucketCosFileIsExist(cosfile) {
			util.Logger.Info(fmt.Sprintf("%s:[%s] -> [%s]", c.Add(color.FgHiGreen).Sprint("cos已经存在"), path, cosfile))
			return nil
		} else {
			util.Logger.Info(fmt.Sprintf("%s:[%s] -> [%s]", c.Add(color.FgHiYellow).Sprint("准备上传文件"), path, cosfile))
			threadUploadFiles(cosfile, path)
		}
	}
	return nil
}
