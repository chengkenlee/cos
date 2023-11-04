package run

import (
	"context"
	"cos/util"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 列出目录下对象
func queryBucketObjectDir() {
	var marker string

	opt := &cos.BucketGetOptions{
		Delimiter: "/",  // deliter 表示分隔符, 设置为/表示列出当前目录下的 object, 设置为空表示列出所有的 object
		MaxKeys:   1000, // 设置最大遍历出多少个对象, 一次 listobject 最大支持1000
	}
	isTruncated := true
	for isTruncated {
		c := color.New()
		opt.Marker = marker
		v, _, err := util.Client.Bucket.Get(context.Background(), opt)
		if err != nil {
			fmt.Println(err)
			break
		}
		for _, content := range v.Contents {
			marshal, err := json.Marshal(&content)
			if err != nil {
				util.Logger.Info(err.Error())
				return
			}
			util.Logger.Info(fmt.Sprintf("文件:%s", c.Add(color.FgHiYellow).Sprint(string(marshal))))
		}
		// common prefix 表示表示被 delimiter 截断的路径, 如 delimter 设置为/, common prefix 则表示所有子目录的路径
		for _, commonPrefix := range v.CommonPrefixes {
			util.Logger.Info(fmt.Sprintf("文件夹:%s", c.Add(color.FgHiGreen).Sprint(commonPrefix)))
		}
		isTruncated = v.IsTruncated // 是否还有数据
		marker = v.NextMarker       // 设置下次请求的起始 key
	}

}

// 查询存储桶列表
func queryBucketsList() {
	s, _, err := util.Client.Service.Get(context.Background())
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	for _, b := range s.Buckets {
		util.Logger.Info(fmt.Sprintf("%v", b))
	}
}

// 检查对象是否存在
func checkObjectBucketCosFileIsExist(cosfile string) bool {
	ok, err := util.Client.Object.IsExist(context.Background(), cosfile)
	if err == nil && ok {
		return true
	} else if err != nil {
		util.Logger.Error(err.Error())
		return false
	} else {
		return false
	}
}
