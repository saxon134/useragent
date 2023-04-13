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
	//var str = "Mozilla/5.0 (Linux; Android 11; 220233L2C Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/97.0.4692.98 Mobile Safari/537.36KSADSDK_V3.3.31_com.xiaoxiaoleyx.hlttx_1.0.4"
	//var str = "Mozilla/5.0 (Linux; Android 11; Hera-BD00 Build/HinovaHera-BD00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36KSADSDK_V3.3.32_com.kuaiyin.player_5.26.00"
	//var str = "Mozilla/5.0 (Linux; Android 11; F810 Build/JRDF810; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36KSADSDK_V3.3.32_com.ushaqi.zhuishushenqi.adfree_3.44.48"
	//var str = "Mozilla/5.0 (Linux; Android 11; ALH-BD00 Build/HinovaALH-BD00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36KSADSDK_V3.3.32_cn.xiaoniangao.dkapp_1.1.4"
	//var str = "Mozilla/5.0 (Linux; Android 9; Xenium S701 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/71.0.3578.99 Mobile Safari/537.36 haokan/7.29.0.10 (Baidu; P1 9) nadcorevendor/5.12.0.23"
	//var str = "Mozilla/5.0 (Linux; Android 7.0; S3 Prow Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/97.0.4692.98 Safari/537.36 T7/13.19 light/1.0 SP-engine/2.57.0 baiduboxapp/13.19.5.10 (Baidu; P1 7.0)"
	var str = "Apache-HttpClient/4.5.13 (Java/1.8.0_201)"
	//var str = ""
	ua := Parse(str)
	fmt.Println("OS:", ua.OS, "OSVersion:", ua.OSVersion, "DevType:", ua.DevType, "DevBrand:", ua.DevBrand, "DevModel:", ua.DevModel)
}
