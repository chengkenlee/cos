package run

import (
	"cos/util"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func init() {
	u, _ := url.Parse("https://chengkenlee-1256202734.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	util.Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  util.Config.GetString("cos.SecretId"),
			SecretKey: util.Config.GetString("cos.SecretKey"),
		},
	})
}

func Run() {
	queryBucketsList()
	queryBucketObjectDir()
	FolderOrFile()
	util.Logger.Info("初始化完成！")
	var wg sync.WaitGroup
	for _, s := range util.Config.GetStringSlice("scan.object") {
		wg.Add(1)
		go func(s string) {
			b, err := os.Stat(s)
			if err != nil {
				util.Logger.Error(err.Error())
			}
			//判断是否是文件夹类型
			if b.IsDir() {
				//开启文件监听器
				fsnotifymon(s)
			}
		}(s)
	}
	util.Logger.Info("monitoring...")
	wg.Wait()
}
