package useragent

import (
	"github.com/mileusna/useragent"
	"github.com/saxon134/go-utils/saData/saHit"
	"strings"
)

type UserAgent struct {
	OS        string //Android、iOS、macOS、Windows、Windows Phone、Linux、ChromeOS
	OSVersion string //系统版本号
	DevType   int    //0-未知 1-手机 2-平板 3-PC
	DevBrand  string //手机品牌，如华为
	DevModel  string //手机型号，如：mate50
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

	var brand = strings.ToUpper(result.DevBrand)
	result.DevModel = strings.TrimPrefix(result.DevModel, brand+" ")
	result.DevModel = strings.TrimPrefix(result.DevModel, brand)

	return result
}
