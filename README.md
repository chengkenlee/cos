# cos sync

## 定时监控本地文件夹变化， 并同步文件到腾讯云cos对象存储

```
定时监控本地文件夹变化， 并同步文件到腾讯云cos对象存储
2023-11-04T23:12:39.197+0800    info    log format is success.
2023-11-04T23:12:39.454+0800    info    {chengkenlee-1256202734 ap-guangzhou 2023-11-04T08:55:16Z}
2023-11-04T23:12:39.865+0800    info    准备上传文件:[D:\Picture\图片20231104\IMG_2197.JPG] -> [/图片20231104/IMG_2197.JPG]
Transfer Start    [ConsumedBytes/TotalBytes: 0/5581253]
Transfer Data     [ConsumedBytes/TotalBytes: 5581253/5581253, 100%]
Transfer Complete [ConsumedBytes/TotalBytes: 5581253/5581253]
[ConsumedBytes/TotalBytes: 5581253/5581253, 100%]2023-11-04T23:12:42.162+0800  info
准备上传文件:[D:\Picture\图片20231104\IMG_2198.JPG] -> [/图片20231104/IMG_2198.JPG]
Transfer Start    [ConsumedBytes/TotalBytes: 0/5940951]
Transfer Data     [ConsumedBytes/TotalBytes: 5940951/5940951, 100%]
Transfer Complete [ConsumedBytes/TotalBytes: 5940951/5940951]
[ConsumedBytes/TotalBytes: 5940951/5940951, 100%]2023-11-04T23:12:45.561+0800  info
准备上传文件:[D:\Picture\图片20231104\IMG_2199.JPG] -> [/图片20231104/IMG_2199.JPG]
Transfer Start    [ConsumedBytes/TotalBytes: 0/5333541]
Transfer Complete [ConsumedBytes/TotalBytes: 5333541/5333541]
[ConsumedBytes/TotalBytes: 5333541/5333541, 100%]2023-11-04T23:12:49.331+0800  info
准备上传文件:[D:\Picture\图片20231104\IMG_2200.JPG] -> [/图片20231104/IMG_2200.JPG]
Transfer Start    [ConsumedBytes/TotalBytes: 0/5385300]
Transfer Data     [ConsumedBytes/TotalBytes: 5385300/5385300, 100%]

[ConsumedBytes/TotalBytes: 5385300/5385300, 100%]2023-11-04T23:12:52.780+0800  info
准备上传文件:[D:\Picture\图片20231104\IMG_2201.JPG] -> [/图片20231104/IMG_2201.JPG]
Transfer Start    [ConsumedBytes/TotalBytes: 0/4457319]
Transfer Data     [ConsumedBytes/TotalBytes: 4457319/4457319, 100%]
Transfer Complete [ConsumedBytes/TotalBytes: 4457319/4457319]
[ConsumedBytes/TotalBytes: 4457319/4457319, 100%]2023-11-04T23:12:54.623+0800  info                                                                                                                                                                                              d
Transfer Data     [ConsumedBytes/TotalBytes: 5385300/5385300, 100%]
Transfer Complete [ConsumedBytes/TotalBytes: 5385300/5385300]
[ConsumedBytes/TotalBytes: 5385300/5385300, 100%]2023-11-04T23:12:52.780+0800  info
准备上传文件:[D:\Picture\图片20231104\IMG_2201.JPG] -> [/图片20231104/IMG_2201.JPG]
Transfer Start    [ConsumedBytes/TotalBytes: 0/4457319]
Transfer Data     [ConsumedBytes/TotalBytes: 4457319/4457319, 100%]
Transfer Complete [ConsumedBytes/TotalBytes: 4457319/4457319]
[ConsumedBytes/TotalBytes: 4457319/4457319, 100%]2023-11-04T23:12:54.623+0800  info                                                                                                                                                                                              d
one.
2023-11-04T23:12:54.623+0800    info    初始化完成！
one.
2023-11-04T23:12:54.623+0800    info    初始化完成！
2023-11-04T23:12:54.624+0800    info    monitoring...
2023-11-04T23:13:14.598+0800    info    检测到目录发生动态变化~ {"Name":"D:\\Picture\\图片20231104\\2023_04_08_11_54_IMG_5634.JPG","Op":1}
2023-11-04T23:13:14.742+0800    info    准备上传文件:[D:\Picture\图片20231104\2023_04_08_11_54_IMG_5634.JPG] -> [/图片20231104/2023_04_08_11_54_IMG_5634.JPG]
2023-11-04T23:13:14.598+0800    info    检测到目录发生动态变化~ {"Name":"D:\\Picture\\图片20231104\\2023_04_08_11_54_IMG_5634.JPG","Op":1}
2023-11-04T23:13:14.742+0800    info    准备上传文件:[D:\Picture\图片20231104\2023_04_08_11_54_IMG_5634.JPG] -> [/图片20231104/2023_04_08_11_54_IMG_5634.JPG]
Transfer Start    [ConsumedBytes/TotalBytes: 0/2442239]
Transfer Data     [ConsumedBytes/TotalBytes: 2442239/2442239, 100%]
Transfer Complete [ConsumedBytes/TotalBytes: 2442239/2442239]
[ConsumedBytes/TotalBytes: 2442239/2442239, 100%]2023-11-04T23:13:16.328+0800  info                                                                                                                                                                                              c
os已经存在:[D:\Picture\图片20231104\IMG_2197.JPG] -> [/图片20231104/IMG_2197.JPG]
2023-11-04T23:13:16.401+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2198.JPG] -> [/图片20231104/IMG_2198.JPG]
2023-11-04T23:13:16.470+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2199.JPG] -> [/图片20231104/IMG_2199.JPG]
2023-11-04T23:13:16.529+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2200.JPG] -> [/图片20231104/IMG_2200.JPG]
2023-11-04T23:13:16.591+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2201.JPG] -> [/图片20231104/IMG_2201.JPG]
2023-11-04T23:13:16.591+0800    info    检测到目录发生动态变化~ {"Name":"D:\\Picture\\图片20231104\\2023_04_08_11_54_IMG_5634.JPG","Op":2}
2023-11-04T23:13:16.724+0800    info    cos已经存在:[D:\Picture\图片20231104\2023_04_08_11_54_IMG_5634.JPG] -> [/图片20231104/2023_04_08_11_54_IMG_5634.JPG]
2023-11-04T23:13:16.777+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2197.JPG] -> [/图片20231104/IMG_2197.JPG]
2023-11-04T23:13:16.840+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2198.JPG] -> [/图片20231104/IMG_2198.JPG]
2023-11-04T23:13:16.900+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2199.JPG] -> [/图片20231104/IMG_2199.JPG]
2023-11-04T23:13:16.959+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2200.JPG] -> [/图片20231104/IMG_2200.JPG]
2023-11-04T23:13:17.021+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2201.JPG] -> [/图片20231104/IMG_2201.JPG]
2023-11-04T23:13:17.022+0800    info    检测到目录发生动态变化~ {"Name":"D:\\Picture\\图片20231104\\2023_04_08_11_54_IMG_5634.JPG","Op":2}
2023-11-04T23:13:17.155+0800    info    cos已经存在:[D:\Picture\图片20231104\2023_04_08_11_54_IMG_5634.JPG] -> [/图片20231104/2023_04_08_11_54_IMG_5634.JPG]
2023-11-04T23:13:17.219+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2197.JPG] -> [/图片20231104/IMG_2197.JPG]
2023-11-04T23:13:17.285+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2198.JPG] -> [/图片20231104/IMG_2198.JPG]
2023-11-04T23:13:17.344+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2199.JPG] -> [/图片20231104/IMG_2199.JPG]
2023-11-04T23:13:17.418+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2200.JPG] -> [/图片20231104/IMG_2200.JPG]
2023-11-04T23:13:17.485+0800    info    cos已经存在:[D:\Picture\图片20231104\IMG_2201.JPG] -> [/图片20231104/IMG_2201.JPG]

```

## Author

* **Author**  - **_ChengKen_**

  ###### **The wind blows away the thoughts, rolled up unruly time**


## Best Regards

* Hat tip to anyone whose code was used
* Inspiration
* etc
