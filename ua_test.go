package useragent

import (
	"fmt"
	"testing"
)

func TestUa(t *testing.T) {
	//var str = "Mozilla/5.0 (Linux; Android 12; M2102J2SC Build/SKQ1.211006.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/96.0.4664.104 Mobile Safari/537.36KSADSDK_V3.3.22.1_mobi.oneway.ks_s.1.0.14"
	//var str = "Mozilla/5.0 (Linux; Android 11; ONEPLUS A6010 Build/RKQ1.201217.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/90.0.4430.210 Mobile Safari/537.36KSADSDK_V3.3.38_com.zhs.android.fgzy_408.105"
	//var str = "Mozilla/5.0 (Linux; Android 8.1.0; OPPO R11t Build/OPM1.171019.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/62.0.3202.84 Mobile Safari/537.36 haokan/7.29.0.10 (Baidu; P1 8.1.0) nadcorevendor/5.12.0.23"
	//var str = "Mozilla/5.0 (Linux; Android 11; SM202210 Build/RP1A.201005.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/80.0.3987.132 Mobile Safari/537.36KSADSDK_V3.3.34.1_com.ucmobile.lite_13.9.9.1161"
	//var str = "Mozilla/5.0 (Linux; Android 12; RMX3610 Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/97.0.4692.98 Mobile Safari/537.36KSADSDK_V3.3.36_com.hainanys.kxssp_1.1.2"
	//var str = "Mozilla/5.0 (Linux; Android 11; 21091116AC Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/94.0.4606.85 Mobile Safari/537.36KSADSDK_V3.3.23_com.example.huigouzixun_3.1"
	var str = "Mozilla/5.0 (Linux; Android 11; 220233L2C Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/97.0.4692.98 Mobile Safari/537.36KSADSDK_V3.3.31_com.xiaoxiaoleyx.hlttx_1.0.4"
	ua := Parse(str)
	fmt.Println("OS:", ua.OS, "OSVersion:", ua.OSVersion, "DevType:", ua.DevType, "DevBrand:", ua.DevBrand, "DevModel:", ua.DevModel)
}
