package useragent

import (
	"bufio"
	"github.com/mileusna/useragent"
	"github.com/saxon134/go-utils/saData/saHit"
	"io"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
)

type UserAgent struct {
	OS        string //Android、iOS、macOS、Windows、Windows Phone、Linux、ChromeOS
	OSVersion string //系统版本号
	DevType   int    //0-未知 1-手机 2-平板 3-PC
	DevBrand  string //手机品牌，如华为
	DevModel  string //手机型号，如：mate50
}

func init() {
	//文件不存在，或者下载时间超过30天
	var needRefresh = false
	finfo, _ := os.Stat("PhoneModels.txt")
	if finfo == nil {
		needRefresh = true
	} else {
		linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
		if linuxFileAttr.Size < 1024 || linuxFileAttr == nil || time.Now().Unix()-linuxFileAttr.Ctimespec.Sec > 24*60*60*30 {
			needRefresh = true
		}
	}

	if needRefresh {
		go func() {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", "https://raw.githubusercontent.com/KHwang9883/MobileModels/master/scripts/models.csv", nil)
			resp, _ := client.Do(req)
			f, _ := os.Create("PhoneModels.txt")
			_, _ = io.Copy(f, resp.Body)
		}()
	}
}

func Parse(userAgent string) UserAgent {
	ua := useragent.Parse(userAgent)
	var result = UserAgent{
		OS:        "",
		OSVersion: "",
		DevType:   0,
		DevBrand:  "",
		DevModel:  "",
	}

	result.OS = ua.OS
	result.OSVersion = ua.OSVersion

	if ua.Mobile {
		result.DevType = 1
	} else if ua.Tablet {
		result.DevType = 2
	} else if ua.Desktop {
		result.DevType = 3
	}

	result.DevBrand = ""
	result.DevModel = saHit.Str(ua.Device != "", ua.Device, "Unknown")
	if result.OS == "iOS" || result.OS == "macOS" {
		result.DevBrand = "Apple"
	} else if strings.Contains(userAgent, "HUAWEI") {
		result.DevBrand = "Huawei"
	} else if strings.Contains(userAgent, "HONOR") {
		result.DevBrand = "Honor"
	} else if strings.Contains(userAgent, "vivo") {
		result.DevBrand = "Vivo"
	} else if strings.Contains(userAgent, "OPPO") {
		result.DevBrand = "Oppo"
	} else if strings.Contains(userAgent, "Redmi") {
		result.DevBrand = "Redmi"
	} else if strings.HasPrefix(result.DevModel, "SM-") {
		result.DevBrand = "Samsung"
	} else if strings.HasPrefix(result.DevModel, "ONEPLUS") {
		result.DevBrand = "Oneplus"
	} else if strings.HasPrefix(result.DevModel, "ZTE") {
		result.DevBrand = "ZTE"
	} else if strings.Contains(userAgent, "Hinova") {
		result.DevBrand = "Hinova"
	} else {
		result.DevBrand = parseBrand(result.DevModel, userAgent)
	}

	if result.DevBrand == "Unknown" {
		f, _ := os.Open("PhoneModels.txt")
		if f != nil {
			userAgent = strings.ToLower(userAgent)
			buf := bufio.NewReader(f)
			for {
				line, err := buf.ReadString('\n')
				if err != nil {
					break
				}

				var ary = strings.Split(line, ",")
				if len(ary) >= 6 {
					if strings.Contains(userAgent, strings.ToLower(ary[0])) {
						if result.DevModel == "Unknown" {
							result.DevModel = ary[0]
						}
						result.DevBrand = ary[2]
						result.DevBrand = strings.ToUpper(result.DevBrand[:1]) + result.DevBrand[1:]
						if ary[1] == "mob" {
							result.DevType = 1
						} else if ary[1] == "pad" {
							result.DevType = 2
						}
					}
				}
			}
		}
	}

	var brand = strings.ToUpper(result.DevBrand)
	result.DevModel = strings.TrimPrefix(result.DevModel, brand+" ")
	result.DevModel = strings.TrimPrefix(result.DevModel, brand)

	return result
}
