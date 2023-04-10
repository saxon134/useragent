package useragent

import (
	"github.com/saxon134/go-utils/saData"
	"strings"
)

var brands = map[string][]string{
	"Xiaomi": {"MI-", "MDT", "MDE", "MCE", "M1", "M2", "MCT", "MDG", "MDI"},
	"360":    {"1503-M02", "1503-A01", "1505-A01", "1505-A02", "1603-A03", "1605-A01", "1605-A02", "1607-A01", "1801-A01"},
	"Vivo":   {"V18", "V19", "V20", "V21"},
	"Oppo":   {"PA", "PD", "PE", "PC"},
}

func parseBrand(model string) string {
	if model == "" || model == "Unknown" {
		return "Unknown"
	}

	for k, ary := range brands {
		for _, v := range ary {
			if strings.HasPrefix(model, v) {
				return k
			}
		}

		if saData.Int(model[:6]) > 0 {
			return "Redmi"
		}

		if strings.HasPrefix(model, "M") {
			if saData.Int(model[1:4]) > 0 {
				return "Meizu"
			}
		}

		if strings.HasPrefix(model, "RMX") {
			return "Realme"
		}
	}

	return "Unknown"
}
