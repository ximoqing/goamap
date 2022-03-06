# goamap
golang 高德地图api

## 安装
### golang version => 1.13
`
go get github.com/ximoqing/goamap
`

## 使用
```go
package main

import (
    "github.com/ximoqing/goamap"
)

func main() {
    ip := goamap.IpJson("你的高德key"，"要定位的ip")
    fmt.Println(ip)
}

```

## api功能
- [X] 地理/逆地理编码
- [X] 路径规划/路径规划 v2
- [X] 行政区域查询
- [X] 搜索POI/搜索POI v2
- [X] 交通事件
- [X] IP定位
- [X] 静态地图
- [X] 坐标转换
- [X] 天气查询
- [X] 输入提示
- [X] 轨迹纠偏
