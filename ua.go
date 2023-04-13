package useragent

import (
	"github.com/mileusna/useragent"
	"github.com/saxon134/go-utils/saData"
	"github.com/saxon134/go-utils/saData/saHit"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
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
			if resp == nil {
				return
			}

			//过滤只保存mob、pad类型
			buf := new(strings.Builder)
			_, _ = io.Copy(buf, resp.Body)
			var str = buf.String()
			var lines = strings.Split(str, "\n")
			var models = make([]string, 0, len(lines))
			for idx, line := range lines {
				re, _ := regexp.Compile(`"*"`)
				line = re.ReplaceAllString(line, ",")
				var items = strings.Split(line, ",")
				//格式校验，如果格式有变化则不处理
				if idx == 0 {
					if len(items) < 7 || items[0] != "model" || items[1] != "dtype" || items[2] != "brand" {
						return
					}
				}

				if len(items) >= 7 {
					if items[1] == "mob" || items[1] == "pad" {
						var brand = items[2]
						var ary = strings.Split(items[0], " ")
						var model = ary[len(ary)-1]
						var modelName = items[6]
						{
							ary = strings.Split(modelName, "（")
							modelName = ary[0]
							modelName = strings.Split(modelName, "(")[0]
							ary = strings.Split(modelName, " ")
							modelName = ary[0]
							if len(ary) >= 2 {
								modelName += " " + ary[1]
							}
						}
						models = append(models, saData.JoinStr(saHit.Str(items[1] == "mob", "m", "p"), ",", brand, ",", model, ",", modelName))
					}
				}
			}

			//保存手机型号数据
			f, err := os.Create("PhoneModels.txt")
			if err != nil {
				log.Println("Create PhoneModels file error :", err)
				return
			}
			_, _ = io.Copy(f, strings.NewReader(strings.Join(models, "\n")))
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

	//从文件读取手机型号数据
	var fileModels = []string{}
	buf, _ := os.ReadFile("PhoneModels.txt")
	if buf != nil {
		fileModels = strings.Split(string(buf), "\n")
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
		if fileModels != nil {
			userAgent = strings.ToLower(userAgent)
			for _, line := range fileModels {
				var ary = strings.Split(line, ",")
				if len(ary) >= 2 {
					if strings.Contains(userAgent, strings.ToLower(ary[0])) {
						if ary[0] == "m" {
							result.DevType = 1 //手机
						} else if ary[0] == "p" {
							result.DevType = 2 //平板
						}

						result.DevBrand = ary[1]
						result.DevBrand = strings.ToUpper(result.DevBrand[:1]) + result.DevBrand[1:]

						if len(ary) >= 4 {
							result.DevModel = ary[3]
						} else if len(ary) >= 3 {
							result.DevModel = ary[2]
						}
					}
				}
			}
		}
	} else {
		var brand = strings.ToUpper(result.DevBrand)
		result.DevModel = strings.TrimPrefix(result.DevModel, brand+" ")
		result.DevModel = strings.TrimPrefix(result.DevModel, brand)
		if fileModels != nil {
			for _, line := range fileModels {
				var ary = strings.Split(line, ",")
				if strings.Contains(result.DevModel, strings.ToLower(ary[0])) {
					if ary[0] == "m" {
						result.DevType = 1 //手机
					} else if ary[0] == "p" {
						result.DevType = 2 //平板
					}

					result.DevBrand = ary[1]
					result.DevBrand = strings.ToUpper(result.DevBrand[:1]) + result.DevBrand[1:]

					if len(ary) >= 4 {
						result.DevModel = ary[3]
					} else if len(ary) >= 3 {
						result.DevModel = ary[2]
					}
				}
			}
		}
	}

	return result
}
