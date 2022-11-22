# User Agent Parser

## 前言

这是一个对 Nodejs 的 `ua-parser-js` 包的 go版本

## Get

```shell
go get github.com/ggchivalrous/db-uaparser
```

## Example

```go
package main

import "github.com/ggchivalrous/db-uaparser"

func main() {
    ua := "Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 635) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537"
    uaInfo := dbuaparser.GetInfo(ua)

    fmt.Println(uaInfo.Browser) // iemobile
    fmt.Println(uaInfo.BrowserVer) // 11.0
    fmt.Println(uaInfo.Cpu) // arm
    fmt.Println(uaInfo.Device) // lumia 635
    fmt.Println(uaInfo.Vendor) // nokia
    fmt.Println(uaInfo.Type) // mobile
    fmt.Println(uaInfo.Engine) // trident
    fmt.Println(uaInfo.EngineVer) // 7.0
    fmt.Println(uaInfo.Os) // windows phone
    fmt.Println(uaInfo.OsVer) // 8.1
}
```