package useragent

import (
	"github.com/saxon134/go-utils/saData"
	"strings"
)

var models = map[string][]string{
	"Huawei":   {"MHA-", "LON-", "ALP-", "BLA-", "NEO-", "HMA-", "LYA-", "EVR-", "LYA-", "TAS-", "LIO-", "OCE-", "NOH-", "NOP-", "TAH-", "TET-", "DAV-", "EVA-", "VIE-", "VTR-", "VKY-", "EML-", "CLT-", "ELE-", "VOG-", "ANA-", "ELS-", "JAD-", "WAS-", "PIC-", "BAC-", "HWI-", "ANE-", "PAR-", "INE-", "VCE-", "MAR-", "SEA-", "GLK-", "SPN-", "WLZ-", "JNY-", "JEF-", "CDY-", "CND-", "CDL-", "ANG-", "BRQ-", "JSC-", "CHL-", "NAM-", "RTE-", "SNE-", "POT-", "NCE-", "DIG-", "TRT-", "SLA-", "FIG-", "FLA-", "LDN-", "ATU-", "DRA-", "JKM-", "ARS-", "DUB-", "POT-", "MRD-", "STK-", "ART-", "AQM-", "MED-", "DVC-", "WKG-", "FRL-", "PPA-", "MLD-", "MNT", "TDT", "Hebe", "F810", "ALH"},
	"Xiaomi":   {"MI-", "MDT", "MDE", "MCE", "M1", "M2", "MCT", "MDG", "MDI"},
	"360":      {"1503-M02", "1503-A01", "1505-A01", "1505-A02", "1603-A03", "1605-A01", "1605-A02", "1607-A01", "1801-A01"},
	"Vivo":     {"V18", "V19", "V20", "V21"},
	"Oppo":     {"PA", "PD", "PE", "PC"},
	"Mpman":    {"MP"},
	"Motorola": {"XT1", "XT2"},
	"Nokia":    {"TA-"},
	"Nubia":    {"NX4", "NX5", "NX6"},
}

var brands = []string{"Huawei", "Xiaomi", "Vivo", "Oppo", "Samsung", "Apple", "Hera:Hinova", "LG", "Motorola", "Lenovo", "Amazon", "Nokia", "Lenovo", "Sony", "Xenium:Philips", "8848"}

func parseBrand(model string, userAgent string) string {
	if model == "" || model == "Unknown" {
		return "Unknown"
	}

	for k, ary := range models {
		for _, v := range ary {
			if strings.HasPrefix(model, v) {
				return k
			}
		}

		if len(model) >= 6 && saData.Int(model[:6]) > 0 {
			return "Redmi"
		}

		if strings.HasPrefix(model, "M") {
			if saData.Int(model[1:4]) > 0 {
				return "Meizu"
			}
		}

		if strings.HasPrefix(model, "meizu") {
			return "Meizu"
		}

		if strings.HasPrefix(model, "RMX") {
			return "Realme"
		}
	}

	var ary = strings.Split(model, " ")
	if len(ary) >= 2 {
		for _, b := range brands {
			var brandAry = strings.Split(b, ":")
			if strings.ToLower(brandAry[0]) == strings.ToLower(ary[0]) {
				if len(brandAry) == 2 {
					return brandAry[1]
				} else {
					return brandAry[0]
				}
			}
		}
	}

	var start = strings.Index(userAgent, "Build/")
	var end = strings.Index(userAgent[start:], model)
	if end > 6 {
		var brand = userAgent[start+6 : start+end]
		if strings.Contains(brand, " ") == false {
			return brand
		}
	}
	return "Unknown"
}
