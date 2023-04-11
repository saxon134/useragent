# useragent

### 介绍

通过userAgent获取手机品牌、型号，系统版本等信息

小米、华为等品牌有不少机型UA不标准，需要特殊处理

手机品牌收集参考：https://github.com/matiji66/MobileModels


### 安装
```
go get -u github.com/saxon134/useragent
```

### 使用

```
var str = "Mozilla/5.0 (Linux; Android 12; M2102J2SC Build/SKQ1.211006.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/96.0.4664.104 Mobile Safari/537.36KSADSDK_V3.3.22.1_mobi.oneway.ks_s.1.0.14"
ua := Parse(str)
fmt.Println("OS:", ua.OS, "OSVersion:", ua.OSVersion, "DevType:", ua.DevType, "DevBrand:", ua.DevBrand, "DevModel:", ua.DevModel)

```
