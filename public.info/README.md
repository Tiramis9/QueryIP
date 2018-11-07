#### overview

依赖第三方geoip2免费IP库 github.com/oschwald/geoip2-golang

可以根据ip地址的归属地，显示本地的语言，目前仅支持以下语言
          		China     = "CN"  = "zh-CN"
			America    ="US" = "en"
			Japanese   = "JP"   = "ja"
			Korea     = "KR"  = "ru"
			Filipino  = "PH"  = "fr"
			Brazil    = "BR"   = "pt-BR"
			Spain     = "ES"  = "es"
			Deutschland = "DE" = "de"

#### Installation 

```
go get github.com/oschwald/maxminddb-golang
```

生成可执行文件.exe,在public目录下执行以下命令

```
go build 
```



#### comments

命令行测试方法//curl -X POST "http://172.16.101.8:8080/ip"  -d '{99.9.0.1}' 

开启服务，8080端口可以在./config/cfg.go里更改端口，172.16.101是本机的ip地址，可以通过ipcofig查看本机的地址信息

./cmd/config/cfg.go  is configuration file

./cmd/task/info.go  is Task distribution handle

./server/DealHander.go  is Processing service 



