package dbuaparser_test

import (
	"fmt"
	"testing"

	dbuaparser "github.com/ggchivalrous/db-uaparser"
)

func TestRegDevice(t *testing.T) {
	deviceLen := len(deviceList)
	errCount := 0
	for _, v := range deviceList {
		res := dbuaparser.GetDevice(v.Ua)
		state := res.Device == v.Expect.Name && res.Vendor == v.Expect.Vendor && res.Type == v.Expect.Type
		if !state {
			errCount++
			fmt.Println("匹配: ", v.Desc, " 失败，结果：name: ", res.Device, " vendor:", res.Vendor, " type: ", res.Type)
		}
	}

	fmt.Println("匹配总数: ", deviceLen, " 成功总数:", deviceLen-errCount, "失败总数: ", errCount)

	if errCount > 0 {
		t.Error("匹配失败数量: ", errCount)
	}
}

type TestDevice struct {
	Desc   string
	Ua     string
	Expect TestDeviceRes
}

type TestDeviceRes struct {
	Name   string
	Vendor string
	Type   string
}

var deviceList = []TestDevice{
	{
		Desc: "ASUS Nexus 7",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 7 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "nexus 7",
			Type:   "tablet",
		},
	},
	{
		Desc: "ASUS Padfone",
		Ua:   "Mozilla/5.0 (Linux; Android 4.1.1; PadFone 2 Build/JRO03L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.117 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "padfone",
			Type:   "tablet",
		},
	},
	{
		Desc: "ASUS ZenPad 10",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; P00C Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "p00c",
			Type:   "tablet",
		},
	},
	{
		Desc: "ASUS ZenPad Z8s",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; ASUS_P00J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.111 Safari/537.36\n",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "p00j",
			Type:   "tablet",
		},
	},
	{
		Desc: "ASUS ROG",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; ZS600KL Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.126 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "zs600kl",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS ROG II",
		Ua:   "Mozilla/5.0 (Linux; Android 9; ASUS_I001DA Build/PKQ1.190414.001; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "i001da",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone 2",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0; ASUS ZenFone 2 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/37.0.0.0 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "zenfone 2",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone 3 Deluxe",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; ASUS_Z016D Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "z016d",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone 5",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0; ZE620KL Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.158 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "ze620kl",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone 7",
		Ua:   "Mozilla/5.0 (Linux; Android 10; ASUS_I002D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.81 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "i002d",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone 7 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 10; ZS671KS) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "zs671ks",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone Max Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; ZB602KL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "zb602kl",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone Max Pro (M1)",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; ASUS_X00TD Build/OPM1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.137 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "x00td",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone Max M2",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; ASUS_X01AD) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.99 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "x01ad",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone Max Pro M2",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; ASUS_X01BDA) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.99 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "x01bda",
			Type:   "mobile",
		},
	},
	{
		Desc: "ASUS Zenfone Go",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; ASUS_X009DA Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "asus",
			Name:   "x009da",
			Type:   "mobile",
		},
	},
	{
		Desc: "Acer Iconia A1-810",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.2; A1-810 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.109 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "acer",
			Name:   "a1-810",
			Type:   "tablet",
		},
	},
	{
		Desc: "BlackBerry Priv",
		Ua:   "User-Agent: Mozilla/5.0 (Linux; Android 5.1.1; STV100-1 Build/LMY47V; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/46.0.2490.76 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "blackberry",
			Name:   "stv100-1",
			Type:   "mobile",
		},
	},
	{
		Desc: "BlackBerry Keyone",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; BBB100-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.111 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "blackberry",
			Name:   "bbb100-1",
			Type:   "mobile",
		},
	},
	{
		Desc: "BlackBerry Key2",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; BBF100-1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "blackberry",
			Name:   "bbf100-1",
			Type:   "mobile",
		},
	},
	{
		Desc: "BlackBerry Key2 LE",
		Ua:   "User-Agent: Mozilla/5.0 (Linux; Android 8.1.0; BBE100-1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "blackberry",
			Name:   "bbe100-1",
			Type:   "mobile",
		},
	},
	{
		Desc: "Essential PH-1",
		Ua:   "Mozilla/5.0 (Linux; Android 9; PH-1 Build/PPR1.180905.036) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "essential",
			Name:   "ph-1",
			Type:   "mobile",
		},
	},
	{
		Desc: "Fairphone 1U",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.2.2; FP1U Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "fairphone",
			Name:   "fp1u",
			Type:   "mobile",
		},
	},
	{
		Desc: "Fairphone 2",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.2; FP2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "fairphone",
			Name:   "fp2",
			Type:   "mobile",
		},
	},
	{
		Desc: "Fairphone 3",
		Ua:   "Mozilla/5.0 (Linux; Android 9; FP3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.93 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "fairphone",
			Name:   "fp3",
			Type:   "mobile",
		},
	},
	{
		Desc: "HTC Desire 820",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 820 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "htc",
			Name:   "desire 820",
			Type:   "mobile",
		},
	},
	{
		Desc: "HTC Evo Shift 4G",
		Ua:   "Mozilla/5.0 (Linux; U; Android 2.3.4; en-us; Sprint APA7373KT Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0",
		Expect: TestDeviceRes{
			Vendor: "sprint",
			Name:   "apa7373kt",
			Type:   "mobile",
		},
	},
	{
		Desc: "HTC Nexus 9",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0; Nexus 9 Build/LRX21R) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Mobile Crosswalk/7.36.154.13 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "htc",
			Name:   "nexus 9",
			Type:   "tablet",
		},
	},
	{
		Desc: "Huawei Honor",
		Ua:   "Mozilla/5.0 (Linux; U; Android 2.3; xx-xx; U8860 Build/HuaweiU8860) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "u8860",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Honor 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YAL-L41) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yal-l41",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Honor 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YAL-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yal-al10",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nexus 6P",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 6P Build/MTC19V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.81 Mobile Safari/537",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "nexus 6p",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P10",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; VTR-L09 Build/HUAWEIVTR-L09; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "vtr-l09",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Y3II",
		Ua:   "Mozilla/5.0 (Linux; U; Android 5.1; xx-xx; HUAWEI LUA-L03 Build/HUAWEILUA-L03) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/39.0.0.0 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "lua-l03",
			Type:   "mobile",
		},
	},
	{
		Desc: "HUAWEI MediaPad M3 Lite 10",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; BAH-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "bah-l09",
			Type:   "tablet",
		},
	},
	{
		Desc: "HUAWEI MediaPad M5 Lite",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; BAH2-W19 Build/HUAWEIBAH2-W19; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "bah2-w19",
			Type:   "tablet",
		},
	},
	{
		Desc: "HUAWEI MediaPad M5",
		Ua:   "Mozilla/5.0 (Linux; Android 9; SHT-AL09 Build/HUAWEISHT-AL09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "sht-al09",
			Type:   "tablet",
		},
	},
	{
		Desc: "HUAWEI MediaPad T5",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; AGS2-L09 Build/HUAWEIAGS2-L09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/84.0.4147.125 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "ags2-l09",
			Type:   "tablet",
		},
	},
	{
		Desc: "HUAWEI MediaPad T10",
		Ua:   "Mozilla/5.0 (Linux; Android 10; AGR-W09 Build/HUAWEIAGR-W09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "agr-w09",
			Type:   "tablet",
		},
	},
	{
		Desc: "HUAWEI MediaPad T10s",
		Ua:   "Mozilla/5.0 (Linux; Android 10; AGS3-W09 Build/HUAWEIAGS3-W09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "ags3-w09",
			Type:   "tablet",
		},
	},
	{
		Desc: "Huawei MatePad T 10",
		Ua:   "Mozilla/5.0 (Linux; Android 10; AGR-L09; HMSCore 5.0.4.301) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 HuaweiBrowser/11.0.3.304 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "agr-l09",
			Type:   "tablet",
		},
	},
	{
		Desc: "Huawei M3",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; BTV-W09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "btv-w09",
			Type:   "tablet",
		},
	},
	{
		Desc: "Huawei Mate 10 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0; BLA-L29 Build/HUAWEIBLA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3236.6 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "bla-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate X",
		Ua:   "Mozilla/5.0 (Linux; Android 9; TAH-AN00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.111 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "tah-an00",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate X2",
		Ua:   "Mozilla/5.0 (Linux; Android 10; TET-AN00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "tet-an00",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 X",
		Ua:   "Mozilla/5.0 (Linux; Android 9; EVR-L29 Build/HUAWEIEVR-L29; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "evr-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LYA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "lya-l09",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LYA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "lya-al00",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LYA-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "lya-al10",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LYA-L0C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "lya-l0c",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LYA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "lya-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LYA-TL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "lya-tl00",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P20 Lite",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; ANE-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.143 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "ane-lx1",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P20",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; EML-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "eml-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P20 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; CLT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "clt-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P30",
		Ua:   "Mozilla/5.0 (Linux; Android 9; ELE-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "ele-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P30 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.143 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "vog-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P40",
		Ua:   "Mozilla/5.0 (Linux; Android 10; ANA-AN00 Build/HUAWEIANA-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.26 SP-engine/2.22.0 baiduboxapp/11.26.0.10 (Baidu; P1 10) NABar/1.0",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "ana-an00",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P40 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 10; ELS-AN00 Build/HUAWEIELS-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Mobile Safari/537.36 mailapp/6.0.0",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "els-an00",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei 30 Pro+",
		Ua:   "Mozilla/5.0 (Linux; Android 10; EBG-AN10 Build/HUAWEIEBG-AN10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36 EdgA/42.0.0.2741",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "ebg-an10",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei 30S",
		Ua:   "Mozilla/5.0 (Linux; Android 10; CDY-AN90 Build/HUAWEICDY-AN90; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Mobile Safari/537.36 mailapp/5.8.0",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "cdy-an90",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nova 5T",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YAL-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yal-l21",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nova 5T",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YAL-L61) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yal-l61",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nova 5T",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YAL-L71) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yal-l71",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nova 5T",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YAL-L61D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yal-l61d",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nova 5T",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YALE-L61A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yale-l61a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nova 5T",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YALE-L61D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yale-l61d",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Nova 5T",
		Ua:   "Mozilla/5.0 (Linux; Android 10; YALE-L71A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "yale-l71a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Enjoy10e",
		Ua:   "Dalvik/2.1.0 (Linux; U; Android 10; MED-AL00 Build/HUAWEIMED-AL00)",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "med-al00",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Honor 6A",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; DLI-L22 Build/HONORDLI-L22; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.116 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/252.0.0.22.355;]",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "dli-l22",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Honor 7",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; PLK-L01 Build/HONORPLK-L01; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "plk-l01",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei 10 Lite",
		Ua:   "Mozilla/5.0 (Linux; Android 9; HRY-LX1 Build/HONORHRY-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "hry-lx1",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Y7 2018",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; LDN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.62 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "ldn-l01",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Honor 8X",
		Ua:   "Mozilla/5.0 (Linux; Android 9; JSN-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "jsn-l21",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Y6 2019",
		Ua:   "Mozilla/5.0 (Linux; Android 9; MRD-LX1N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "mrd-lx1n",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Y9 2019",
		Ua:   "Mozilla/5.0 (Linux; Android 9; JKM-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "jkm-lx2",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Y5",
		Ua:   "Mozilla/5.0 (Linux; Android 9; AMN-LX3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "amn-lx3",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Y7p",
		Ua:   "Mozilla/5.0 (Linux; Android 9; ART-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.92 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "art-l29",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Mate 20 Lite",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; SNE-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "sne-lx1",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei P10 Lite",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "was-lx1a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Y5 Lite 2018",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; DRA-LX5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "dra-lx5",
			Type:   "mobile",
		},
	},
	{
		Desc: "Huawei Honor 8C",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; BKK-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "huawei",
			Name:   "bkk-lx2",
			Type:   "mobile",
		},
	},
	{
		Desc: "iPad using UCBrowser",
		Ua:   "Mozilla/5.0 (iPad; U; CPU OS 11_2 like Mac OS X; zh-CN; iPad5,3) AppleWebKit/534.46 (KHTML, like Gecko) UCBrowser/3.0.1.776 U3/ Mobile/10A403 Safari/7543.48.3",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "ipad",
			Type:   "tablet",
		},
	},
	{
		Desc: "iPad Air",
		Ua:   "Mozilla/5.0 (iPad; CPU OS 12_4_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPad4,1;FBMD/iPad;FBSN/iOS;FBSV/12.4.5;FBSS/2;FBID/tablet;FBLC/en_US;FBOP/5;FBCR/]",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "ipad",
			Type:   "tablet",
		},
	},
	{
		Desc: "iPad using Facebook Browser",
		Ua:   "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) WebKit/8610 (KHTML, like Gecko) Mobile/18D70 [FBAN/FBIOS;FBDV/iPad7,11;FBMD/iPad;FBSN/iOS;FBSV/14.4.2;FBSS/2;FBID/tablet;FBLC/en_US;FBOP/5]",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "ipad",
			Type:   "tablet",
		},
	},
	{
		Desc: "iPod",
		Ua:   "Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B554a Safari/9537.53",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "ipod touch",
			Type:   "mobile",
		},
	},
	{
		Desc: "Lenovo Tab 2",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0.1; Lenovo TAB 2 A7-30HC Build/LRX21M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.157 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lenovo",
			Name:   "tab 2 a7",
			Type:   "tablet",
		},
	},
	{
		Desc: "Lenovo Phone",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; Lenovo PB2-650M Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.105 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/311.0.0.44.117;]",
		Expect: TestDeviceRes{
			Vendor: "lenovo",
			Name:   "pb2-650m",
			Type:   "mobile",
		},
	},
	{
		Desc: "Lenovo Tab 3 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; Lenovo YT3-X90F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.99 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lenovo",
			Name:   "yt3-x90f",
			Type:   "tablet",
		},
	},
	{
		Desc: "Lenovo Tab 4",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.1; Lenovo TB-X304F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.99 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lenovo",
			Name:   "tb-x304f",
			Type:   "tablet",
		},
	},
	{
		Desc: "Lenovo Tab 4",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; Lenovo TAB 2 A7-30HC) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lenovo",
			Name:   "tab 2 a7",
			Type:   "tablet",
		},
	},
	{
		Desc: "Lenovo Tab 7 Essential",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; Lenovo TB-7304X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lenovo",
			Name:   "tb-7304x",
			Type:   "tablet",
		},
	},
	{
		Desc: "Lenovo Tab M10",
		Ua:   "Mozilla/5.0 (Linux; arm_64; Android 9; Lenovo TB-X606F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 YaBrowser/20.9.4.99.01 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lenovo",
			Name:   "tb-x606f",
			Type:   "tablet",
		},
	},
	{
		Desc: "LG V40 ThinQ",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LM-V405) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "lm-v405",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG K30",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; LM-X410.F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "lm-x410.f",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG K30",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LM-X410.FGN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.93 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "lm-x410.fgn",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG Stylo 5",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "lm-q720",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG G7 ThinQ",
		Ua:   "Mozilla/5.0 (Linux; Android 9; LM-G710VM Build/PKQ1.181105.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "lm-g710vm",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG K500",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; LG-K500 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "k500",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG Nexus 4",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.1; Nexus 4 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "nexus 4",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG Nexus 4",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.3; en-us; Google Nexus 4 - 4.3 - API 18 - 768x1280 Build/JLS36G) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "nexus 4",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG Nexus 5",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 5 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "nexus 5",
			Type:   "mobile",
		},
	},
	{
		Desc: "LG Wing",
		Ua:   "Mozilla/5.0 (Linux; Android 10; LM-F100N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "lm-f100n",
			Type:   "mobile",
		},
	},
	{
		Desc: "Meizu M5 Note",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; M5 Note Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043024 Safari/537.36 MicroMessenger/6.5.7.1040 NetType/WIFI Language/zh_CN",
		Expect: TestDeviceRes{
			Vendor: "meizu",
			Name:   "m5 note",
			Type:   "mobile",
		},
	},
	{
		Desc: "Microsoft Lumia 950",
		Ua:   "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586",
		Expect: TestDeviceRes{
			Vendor: "microsoft",
			Name:   "lumia 950",
			Type:   "mobile",
		},
	},
	{
		Desc: "Microsoft Surface Duo",
		Ua:   "Dalvik/2.1.0 (Linux; U; Android 10; Surface Duo Build/2020.1014.61)",
		Expect: TestDeviceRes{
			Vendor: "microsoft",
			Name:   "surface duo",
			Type:   "tablet",
		},
	},
	{
		Desc: "Motorola Moto X",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.4; XT1097 Build/KXE21.187-38) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.109 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "xt1097",
			Type:   "mobile",
		},
	},
	{
		Desc: "Meizu M3S",
		Ua:   "Mozilla/5.0 (X11; Linux; Android 5.1; MZ-M3s Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrom/45.0.2454.94 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "meizu",
			Name:   "m3s",
			Type:   "mobile",
		},
	},
	{
		Desc: "Microsoft Lumia 950",
		Ua:   "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586",
		Expect: TestDeviceRes{
			Vendor: "microsoft",
			Name:   "lumia 950",
			Type:   "mobile",
		},
	},
	{
		Desc: "Motorola Nexus 6",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.20 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "nexus 6",
			Type:   "mobile",
		},
	},
	{
		Desc: "Motorola Droid RAZR 4G",
		Ua:   "Mozilla/5.0 (Linux; U; Android 2.3; xx-xx; DROID RAZR 4G Build/6.5.1-73_DHD-11_M1-29) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "droid razr 4g",
			Type:   "mobile",
		},
	},
	{
		Desc: "Motorola RAZR 2019",
		Ua:   "Mozilla/5.0 (Linux; Android 9; motorola razr) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/11.1 Chrome/75.0.3770.143 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "razr",
			Type:   "mobile",
		},
	},
	{
		Desc: "iPhone",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "iphone",
			Type:   "mobile",
		},
	},
	{
		Desc: "iPhone SE",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone8,4;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/2;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "iphone",
			Type:   "mobile",
		},
	},
	{
		Desc: "iPhone SE using Facebook App",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone8,4;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/2;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "iphone",
			Type:   "mobile",
		},
	},
	{
		Desc: "iPhone 11 Pro Max",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone12,5;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "iphone",
			Type:   "mobile",
		},
	},
	{
		Desc: "iPhone XS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone11,2;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "iphone",
			Type:   "mobile",
		},
	},
	{
		Desc: "iPod touch",
		Ua:   "Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0_2 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A501 Safari/9537.53",
		Expect: TestDeviceRes{
			Vendor: "apple",
			Name:   "ipod touch",
			Type:   "mobile",
		},
	},
	{
		Desc: "Moto X",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.2; xx-xx; XT1058 Build/13.9.0Q2.X-70-GHOST-ATT_LE-2) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "xt1058",
			Type:   "mobile",
		},
	},
	{
		Desc: "Motorola Moto g(6) Play",
		Ua:   "Mozilla/5.0 (Linux; Android 9; moto g(6) play) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "moto g(6) play",
			Type:   "mobile",
		},
	},
	{
		Desc: "Motorola Moto g(7) Supra",
		Ua:   "Mozilla/5.0 (Linux; Android 9; moto g(7) supra Build/PCOS29.114-134-2; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/73.0.3683.90 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "moto g(7) supra",
			Type:   "mobile",
		},
	},
	{
		Desc: "Motorola Moto E",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NDQS26.69-64-11-7; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "motorola",
			Name:   "moto e (4)",
			Type:   "mobile",
		},
	},
	{
		Desc: "Nokia3xx",
		Ua:   "Nokia303/14.87 CLDC-1.1",
		Expect: TestDeviceRes{
			Vendor: "nokia",
			Name:   "303",
			Type:   "mobile",
		},
	},
	{
		Desc: "Nokia 3.2",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Nokia 3.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "nokia",
			Name:   "3.2",
			Type:   "mobile",
		},
	},
	{
		Desc: "Nokia N9",
		Ua:   "Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13",
		Expect: TestDeviceRes{
			Vendor: "nokia",
			Name:   "n9",
			Type:   "mobile",
		},
	},
	{
		Desc: "Nokia 2720 Flip",
		Ua:   "Mozilla/5.0 (Mobile; Nokia_2720_Flip; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5.2",
		Expect: TestDeviceRes{
			Vendor: "nokia",
			Name:   "2720 flip",
			Type:   "mobile",
		},
	},
	{
		Desc: "Oculus Quest",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Quest) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/15.0.0.0.22.280317669 SamsungBrowser/4.0 Chrome/89.0.4389.90 VR Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "facebook",
			Name:   "quest",
			Type:   "wearable",
		},
	},
	{
		Desc: "Oculus Quest 2",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Quest 2) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/15.0.0.0.22.280317669 SamsungBrowser/4.0 Chrome/89.0.4389.90 VR Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "facebook",
			Name:   "quest 2",
			Type:   "wearable",
		},
	},
	{
		Desc: "OnePlus One",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.4; A0001 Build/KTU84Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "a0001",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus One",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; OnePlus One A0001 Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.117 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "a0001",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus 2",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; ONE A2003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.93 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "a2003",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus 3",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.1; ONEPLUS A3000 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.98 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "a3000",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus 6",
		Ua:   "Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.89 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "a6003",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus 6T",
		Ua:   "Mozilla/5.0 (Linux; Android 9; ONEPLUS A6010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "a6010",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus 8T",
		Ua:   "Mozilla/5.0 (Linux; Android 11; KB2005) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "kb2005",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus 8 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 10; IN2025) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.119 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "in2025",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus Nord N100",
		Ua:   "Mozilla/5.0 (Linux; Android 10; BE2015 Build/QKQ1.200719.002; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "be2015",
			Type:   "mobile",
		},
	},
	{
		Desc: "OnePlus Nord N10 5G",
		Ua:   "Mozilla/5.0 (Linux; Android 10; BE2029) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.185 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oneplus",
			Name:   "be2029",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Neo",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.2.2; zh-cn; R831T Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 OppoBrowser/3.3.2 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "r831t",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO R7s",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; OPPO R7s Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 Chrome/37.0.0.0 MQQBrowser/7.1 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "r7s",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO A3s",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; CPH1803 Build/OPM1.171019.026; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "cph1803",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO A12",
		Ua:   "Mozilla/5.0 (Linux; Android 9; CPH2083) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "cph2083",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Reno",
		Ua:   "Mozilla/5.0 (Linux; Android 9; PCAT00 Build/PKQ1.190101.001; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "pcat00",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Reno3 Pro 5G",
		Ua:   "Mozilla/5.0 (Linux; Android 10; PCLM50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.117 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "pclm50",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Reno4 SE",
		Ua:   "Mozilla/5.0 (Linux; U; Android 10; xx-xx; PEAM00 Build/QP1A.190711.020) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "peam00",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Reno4 5G",
		Ua:   "Mozilla/5.0 (Linux; Android 10; PDPM00 Build/QKQ1.200216.002; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "pdpm00",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Reno4 Pro 5G",
		Ua:   "Mozilla/5.0 (Linux; U; Android 10; xx-xx; PDNT00 Build/QKQ1.200216.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "pdnt00",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Reno5 A",
		Ua:   "Mozilla/5.0 (Linux; Android 11; A101OP) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "a101op",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Find X",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; PAFM00 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "pafm00",
			Type:   "mobile",
		},
	},
	{
		Desc: "OPPO Find 7a",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.3; xx-xx; X9007 Build/JLS36C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "oppo",
			Name:   "x9007",
			Type:   "mobile",
		},
	},
	{
		Desc: "Realme C1",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; RMX1811 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.126 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "realme",
			Name:   "rmx1811",
			Type:   "mobile",
		},
	},
	{
		Desc: "Realme C2",
		Ua:   "Mozilla/5.0 (Linux; Android 9; RMX1941) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "realme",
			Name:   "rmx1941",
			Type:   "mobile",
		},
	},
	{
		Desc: "Realme Narzo 20",
		Ua:   "Mozilla/5.0 (Linux; U; Android 10; xx-xx; RMX2193 Build/QP1A.190711.020) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "realme",
			Name:   "rmx2193",
			Type:   "mobile",
		},
	},
	{
		Desc: "Realme 2 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 9; RMX1801) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "realme",
			Name:   "rmx1801",
			Type:   "mobile",
		},
	},
	{
		Desc: "Roku",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36 Roku/DVP-8.10 (468.10E04145A)",
		Expect: TestDeviceRes{
			Vendor: "roku",
			Name:   "dvp-8.10",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Roku",
		Ua:   "Roku4640X/DVP-7.70 (297.70E04154A)",
		Expect: TestDeviceRes{
			Vendor: "roku",
			Name:   "dvp-7.70",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Xiaomi TV",
		Ua:   "Mozilla/5.0 (Linux; Android 10; MiTV-MOOQ0 Build/QTG3.200305.006; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/94.0.4606.61 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mitv-mooq0",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Kindle Fire HD",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; KFTT Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.4 Mobile Safari/535.19 Silk-Accelerated=true",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "kftt",
			Type:   "tablet",
		},
	},
	{
		Desc: "Kindle Fire HD",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; KFTT) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.4 Mobile Safari/535.19 Silk-Accelerated=true",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "kftt",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy A21s",
		Ua:   "Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-A217F) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/11.0 Chrome/75.0.3770.143 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-a217f",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy A31",
		Ua:   "Mozilla/5.0 (Linux; Android 10; SM-A315G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-a315g",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy A50",
		Ua:   "Mozilla/5.0 (Linux; Android 9; SM-A505F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.105 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-a505f",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy A80",
		Ua:   "Mozilla/5.0 (Linux; Android 9; SM-A805F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.112 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-a805f",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy Fold",
		Ua:   "Mozilla/5.0 (Linux; Android 9; SAMSUNG SM-F900U Build/PPR1.180610.011) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/9.2 Chrome/67.0.3396.87 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-f900u",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy Z Flip",
		Ua:   "Mozilla/5.0 (Linux; Android 10; SM-F700N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-f700n",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy Z Fold2",
		Ua:   "Mozilla/5.0 (Linux; Android 10; SM-F916B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-f916b",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy S10E",
		Ua:   "Mozilla/5.0 (Linux; Android 9; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-g970f",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy Note 10+",
		Ua:   "Mozilla/5.0 (Linux; Android 9; SM-N976V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.89 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-n976v",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung SM-C5000",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; SM-C5000 Build/MMB29M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 wkbrowser 4.1.35 3065",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-c5000",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy Note 8",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.2; GT-N5100 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "gt-n5100",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung SM-T231",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; SM-T231 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-t231",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy Tab 6 Lite",
		Ua:   "Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-P610) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/12.0 Chrome/79.0.3945.136 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-p610",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy Tab A 9.7",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.1; SM-P550 Build/NMF26X; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.90 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-p550",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy Tab A 10.1",
		Ua:   " Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-T515) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/13.0 Chrome/83.0.4103.106 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-t515",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy Tab S7",
		Ua:   "Mozilla/5.0 (Linux; Android 10; SM-T870) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-t870",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy Tab S8",
		Ua:   "Mozilla/5.0 (Linux; Android 12; SM-X706B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-x706b",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy Tab S",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; SM-T700 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-t700",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Galaxy Tab Pro 10.1",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; SM-T520 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-t520",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung Note 10.1",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; SM-P605) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-p605",
			Type:   "tablet",
		},
	},
	{
		Desc: "Samsung SmartTV2011",
		Ua:   "HbbTV/1.1.1 (;;;;;) Maple;2011",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "smarttv2011",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Samsung SmartTV2012",
		Ua:   "HbbTV/1.1.1 (;Samsung;SmartTV2012;;;) WebKit",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "smarttv2012",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Samsung SmartTV2014",
		Ua:   "HbbTV/1.1.1 (;Samsung;SmartTV2014;T-NT14UDEUC-1060.4;;) WebKit",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "smarttv2014",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Sharp AQUOS-TVX19B",
		Ua:   "Mozilla/5.0 (Linux; Android 9; AQUOS-TVX19B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sharp",
			Name:   "aquos-tvx19b",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Sharp Aquos B10",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; SH-A01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sharp",
			Name:   "sh-a01",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sharp Aquos L2",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; SH-L02 Build/S4045) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sharp",
			Name:   "sh-l02",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sharp Aquos R2",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0; SHV42) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.92 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sharp",
			Name:   "shv42",
			Type:   "mobile",
		},
	},
	{
		Desc: "SONY Xperia 1 III",
		Ua:   "Mozilla/5.0 (Linux; Android 11; A101SO) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "a101so",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony G8141 (Xperia XZ Premium)",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; G8141) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "g8141",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony C5303 (Xperia SP)",
		Ua:   "Mozilla/5.0 (Linux; Android 4.3; C5303 Build/12.1.A.1.205) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "c5303",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony SO-02F (Xperia Z1 F)",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.2; SO-02F Build/14.1.H.2.119) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "so-02f",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony D6653 (Xperia Z3)",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4; D6653 Build/23.0.A.0.376) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "d6653",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony Xperia SOL25 (ZL2)",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.4; SOL25 Build/17.1.1.C.1.64) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "sol25",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony Xperia SP",
		Ua:   "Mozilla/5.0 (Linux; Android 4.3; C5302 Build/12.1.A.1.201) AppleWebkit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "c5302",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony Xperia L4",
		Ua:   "Mozilla/5.0 (Linux; Android 9; XQ-AD51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.83 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xq-ad51",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony Xperia 1ii",
		Ua:   "Mozilla/5.0 (Linux; Android 10; XQ-AT51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xq-at51",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony Xperia 1ii",
		Ua:   "Mozilla/5.0 (Linux; Android 10; SOG01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "sog01",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony Xperia 10ii",
		Ua:   "Mozilla/5.0 (Linux; Android 10; XQ-AU52) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xq-au52",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony Xperia Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 10; XQ-AQ52) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.185 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xq-aq52",
			Type:   "mobile",
		},
	},
	{
		Desc: "Sony SGP521 (Xperia Z2 Tablet)",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4; SGP521 Build/17.1.A.0.432) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xperia tablet",
			Type:   "tablet",
		},
	},
	{
		Desc: "Sony Xperia Z2 Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; SGP561) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.99 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xperia tablet",
			Type:   "tablet",
		},
	},
	{
		Desc: "Sony Tablet S",
		Ua:   "Mozilla/5.0 (Linux; U; Android 3.1; Sony Tablet S Build/THMAS10000) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xperia tablet",
			Type:   "tablet",
		},
	},
	{
		Desc: "Sony Tablet Z LTE",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.1; SonySGP321 Build/10.2.C.0.143) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "xperia tablet",
			Type:   "tablet",
		},
	},
	{
		Desc: "Sony BRAVIA 4K GB ATV3",
		Ua:   "Mozilla/5.0 (Linux; Andr0id 9; BRAVIA 4K GB ATV3 Build/PTT1.190515.001.S38) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36 OPR/46.0.2207.0 OMI/4.13.0.180.DIA5.104 Model/Sony-BRAVIA-4K-GB-ATV3",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "bravia 4k gb atv3",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Sony BRAVIA 4K GB ATV3",
		Ua:   "Mozilla/5.0 (Linux; Android 9; BRAVIA 4K GB ATV3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "bravia 4k gb atv3",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Sony Bravia 4k UR2",
		Ua:   "Mozilla/5.0 (Linux: Andr0id 9: BRAVIA 4K UR2 Build/PTT1.190515.001.S104) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36 OPR/46.0.2207.0 OMI/4.13.5.431.DIA5HBBTV.250 Model/Sony-BRAVIA-4K-UR2",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "bravia 4k ur2",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Xiaomi 2013023",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.2.2; en-US; 2013023 Build/HM2013023) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/10.0.1.512 U3/0.8.0 Mobile Safari/533.1",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "2013023",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Hongmi Note 1W",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.2.2; zh-CN; HM NOTE 1W Build/JDQ39) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.7.9.439 U3/0.8.0 Mobile Safari/533.1",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "hm note 1w",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi 3C",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.3; zh-CN; MI 3C Build/JLS36C) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.7.9.439 U3/0.8.0 Mobile Safari/533.1",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi 3c",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi 5",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; MI 5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.83 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi 5",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi 6",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1; MI 6 Build/NMF26X; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi 6",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi 5s Plus",
		Ua:   "Mozilla/5.0 (Linux; U; Android 6.0.1; zh-cn; MI 5s Plus Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.146 Mobile Safari/537.36 XiaoMi/MiuiBrowser/8.7.1",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi 5s plus",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi A1",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi a1",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi Note",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.4; MI NOTE LTE Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi note lte",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi One Plus",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; MI-ONE Plus Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi-one plus",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi Max 3",
		Ua:   "Mozilla/5.0 (Linux; Android 9; MI MAX 3 Build/PKQ1.181007.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi max 3",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi A1",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Mi A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.101 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi a1",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi A2 Lite",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Mi A2 Lite) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.62 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi a2 lite",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi 9 SE",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Mi 9 SE) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi 9 se",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi A2",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Mi A2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi a2",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Mi CC9",
		Ua:   "Mozilla/5.0 (Linux; U; Android 11; zh-cn; MI CC 9 Build/RKQ1.200826.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.116 Mobile Safari/537.36 XiaoMi/MiuiBrowser/15.5.18",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi cc 9",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi MI PAD 2",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1; MI PAD 2 Build/LMY47I; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/60.0.3112.107 Safari/537.36 [FB_IAB/FB4A;FBAV/137.0.0.24.91;]",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi pad 2",
			Type:   "tablet",
		},
	},
	{
		Desc: "Xiaomi MI PAD 4 PLUS",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1; MI PAD 4 PLUS) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "mi pad 4 plus",
			Type:   "tablet",
		},
	},
	{
		Desc: "Xiaomi POCO X2",
		Ua:   "Mozilla/5.0 (Linux; Android 10; POCO X2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "poco x2",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Redmi 4A",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; Redmi 4A Build/MMB29M; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "redmi 4a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Redmi K30 5G",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Redmi K30 5G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "redmi k30 5g",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Redmi K30 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Redmi K30 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "redmi k30 pro",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Redmi Note 3",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; Redmi Note 3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "redmi note 3",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xiaomi Redmi Note 9 Pro Max",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Redmi Note 9 Pro Max) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.99 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "xiaomi",
			Name:   "redmi note 9 pro max",
			Type:   "mobile",
		},
	},
	{
		Desc: "PlayStation 4",
		Ua:   "Mozilla/5.0 (PlayStation 4 3.00) AppleWebKit/537.73 (KHTML, like Gecko)",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "playstation 4",
			Type:   "console",
		},
	},
	{
		Desc: "PlayStation 5",
		Ua:   "Mozilla/5.0 (Playstation; Playstation 5/1.05) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Safari/605.1.15",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "playstation 5",
			Type:   "console",
		},
	},
	{
		Desc: "PlayStation Vita",
		Ua:   "Mozilla/5.0 (PlayStation Vita 3.52) AppleWebKit/537.73 (KHTML, like Gecko) Silk/3.2",
		Expect: TestDeviceRes{
			Vendor: "sony",
			Name:   "playstation vita",
			Type:   "console",
		},
	},
	{
		Desc: "Nintendo Switch",
		Ua:   "Mozilla/5.0 (Nintendo Switch; WifiWebAuthApplet) AppleWebKit/606.4 (KHTML, like Gecko) NF/6.0.1.15.4 NintendoBrowser/5.1.0.20393",
		Expect: TestDeviceRes{
			Vendor: "nintendo",
			Name:   "switch",
			Type:   "console",
		},
	},
	{
		Desc: "Nintendo WiiU",
		Ua:   "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.9 NintendoBrowser/4.2.0.11146.EU",
		Expect: TestDeviceRes{
			Vendor: "nintendo",
			Name:   "wiiu",
			Type:   "console",
		},
	},
	{
		Desc: "Nintendo Wii",
		Ua:   "Opera/9.10 (Nintendo Wii; U; ; 1621; en)",
		Expect: TestDeviceRes{
			Vendor: "nintendo",
			Name:   "wii",
			Type:   "console",
		},
	},
	{
		Desc: "Nintendo 3DS",
		Ua:   "Mozilla/5.0 (Nintendo 3DS; U; ; en) Version/1.7610.EU",
		Expect: TestDeviceRes{
			Vendor: "nintendo",
			Name:   "3ds",
			Type:   "console",
		},
	},
	{
		Desc: "Nintendo 3DS",
		Ua:   "Mozilla/5.0 (New Nintendo 3DS like iPhone) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.0.5.15 Mobile NintendoBrowser/1.3.10126.EU",
		Expect: TestDeviceRes{
			Vendor: "nintendo",
			Name:   "3ds",
			Type:   "console",
		},
	},
	{
		Desc: "Galaxy Nexus",
		Ua:   "Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "galaxy nexus",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy C9 Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; SAMSUNG SM-C900F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.2 Chrome/44.0.2403.133 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-c900f",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy S5",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0; SM-G900F Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.78 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-g900f",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy S6",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; SM-G920I Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-g920i",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy S6 Edge",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; SM-G925I Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-g925i",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy Note 5 Chrome",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; SM-N920C Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.91 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-n920c",
			Type:   "mobile",
		},
	},
	{
		Desc: "Samsung Galaxy Note 5 Samsung Browser",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; SAMSUNG SM-N920C Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "samsung",
			Name:   "sm-n920c",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Chromecast",
		Ua:   "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.84 Safari/537.36 CrKey/1.22.79313",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "chromecast",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Google Pixel C",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; Pixel C Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel c",
			Type:   "tablet",
		},
	},
	{
		Desc: "Google Pixel C",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; Pixel C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel c",
			Type:   "tablet",
		},
	},
	{
		Desc: "Google Pixel",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1; Pixel Build/NDE63V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.85 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel XL",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1; Pixel XL Build/NDE63X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.85 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel XL",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Pixel XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 2",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; Pixel 2 Build/OPM1.171019.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 2",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 2 XL",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; Pixel 2 XL Build/OPM1.171019.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 2 xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 2 XL",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Pixel 2 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 2 xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 3",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Pixel 3 Build/PD1A.180720.030) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 3",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 3 XL",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Pixel 3 XL Build/PD1A.180720.030) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 3 xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 3 XL",
		Ua:   "Mozilla/5.0 (Linux; Android 9; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 3 xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 3a",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Pixel 3a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 3a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 3a XL",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Pixel 3a XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 3a xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 4",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Pixel 4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 4",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 4a",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Pixel 4a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.83 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 4a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 4 XL",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Pixel 4 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 4 xl",
			Type:   "mobile",
		},
	},
	{
		Desc: "Google Pixel 5",
		Ua:   "Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "google",
			Name:   "pixel 5",
			Type:   "mobile",
		},
	},
	{
		Desc: "Generic Android Device",
		Ua:   "Mozilla/5.0 (Linux; U; Android 6.0.1; i980 Build/MRA58K)",
		Expect: TestDeviceRes{
			Vendor: "generic",
			Name:   "android 6.0.1",
		},
	},
	{
		Desc: "Android Phone Unidentified Vendor (docomo F-04K)",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; F-04K Build/V15R060P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.137 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Name: "f-04k",
			Type: "mobile",
		},
	},
	{
		Desc: "docomo SH-02M",
		Ua:   "Mozilla/5.0 (Linux; Android 9; SH-02M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "sharp",
			Name:   "sh-02m",
			Type:   "mobile",
		},
	},
	{
		Desc: "Android Tablet Unidentified Vendor (docomo F-02K)",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; F-02K Build/V44R059G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.109 Safari/537.36",
		Expect: TestDeviceRes{
			Name: "f-02k",
			Type: "tablet",
		},
	},
	{
		Desc: "Android Tablet Unidentified Vendor (docomo d-02K)",
		Ua:   "Mozilla/5.0 (Linux; Android 9; d-02K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Safari/537.36",
		Expect: TestDeviceRes{
			Name: "d-02k",
			Type: "tablet",
		},
	},
	{
		Desc: "LG VK Series Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0.2; VK700 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.84 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "vk700",
			Type:   "tablet",
		},
	},
	{
		Desc: "LG LK Series Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0.1; LGLK430 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/38.0.2125.102 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lg",
			Name:   "lk430",
			Type:   "tablet",
		},
	},
	{
		Desc: "RCA Voyager III Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; RCT6973W43 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "rca",
			Name:   "rct6973w43",
			Type:   "tablet",
		},
	},
	{
		Desc: "RCA Voyager II Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0; RCT6773W22B Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "rca",
			Name:   "rct6773w22b",
			Type:   "tablet",
		},
	},
	{
		Desc: "Verizon Quanta Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; QMV7B Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "verizon",
			Name:   "qmv7b",
			Type:   "tablet",
		},
	},
	{
		Desc: "Verizon Ellipsis 8 Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; QTAQZ3 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "verizon",
			Name:   "qtaqz3",
			Type:   "tablet",
		},
	},
	{
		Desc: "Verizon Ellipsis 8HD Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; QTASUN1 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.81 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "verizon",
			Name:   "qtasun1",
			Type:   "tablet",
		},
	},
	{
		Desc: "Dell Venue 8 Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; Venue 8 3830 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "dell",
			Name:   "venue 8 3830",
			Type:   "tablet",
		},
	},
	{
		Desc: "Dell Venue 7 Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; Venue 7 3730 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "dell",
			Name:   "venue 7 3730",
			Type:   "tablet",
		},
	},
	{
		Desc: "Barnes & Noble Nook HD+ Tablet",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; Barnes & Noble Nook HD+ Build/JZO54K; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "barnes & noble",
			Name:   "nook hd+",
			Type:   "tablet",
		},
	},
	{
		Desc: "Barnes & Noble V400 Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 4.0.4; BNTV400 Build/IMM76L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.111 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "barnes & noble",
			Name:   "v400",
			Type:   "tablet",
		},
	},
	{
		Desc: "NuVision TM101A540N Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1; TM101A540N Build/LMY47I; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/50.0.2661.86 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "nuvision",
			Name:   "tm101a540n",
			Type:   "tablet",
		},
	},
	{
		Desc: "ZTE-Z431",
		Ua:   "ZTE-Z431/1.4.0 NetFront/4.2 QTV5.1 Profile/MIDP-2.1 Configuration/CLDC-1.1",
		Expect: TestDeviceRes{
			Vendor: "zte",
			Name:   "z431",
			Type:   "mobile",
		},
	},
	{
		Desc: "ZTE",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; ZTE-Z740G Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "zte",
			Name:   "z740g",
			Type:   "mobile",
		},
	},
	{
		Desc: "ZTE K Series Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; K88 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "zte",
			Name:   "k88",
			Type:   "tablet",
		},
	},
	{
		Desc: "ZTE Nubia Red Magic 3",
		Ua:   "Mozilla/5.0 (Linux; Android 9; NX629J Build/PKQ1.190321.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/45016 Mobile Safari/537.36 MMWEBID/4064 MicroMessenger/7.0.10.1580(0x27000A34) Process/tools NetType/WIFI Language/zh_CN ABI/arm64",
		Expect: TestDeviceRes{
			Vendor: "zte",
			Name:   "nx629j",
			Type:   "mobile",
		},
	},
	{
		Desc: "ZTE Blade A5",
		Ua:   "Mozilla/5.0 (Linux; Android 9; ZTE Blade A5 2019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "zte",
			Name:   "blade a5 2019",
			Type:   "mobile",
		},
	},
	{
		Desc: "ZTE BLADE V0730",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; ZTE BLADE V0730) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "zte",
			Name:   "blade v0730",
			Type:   "mobile",
		},
	},
	{
		Desc: "ZTE B2017G",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.1; ZTE B2017G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.93 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "zte",
			Name:   "b2017g",
			Type:   "mobile",
		},
	},
	{
		Desc: "Swizz GEN610",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; GEN610 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "swiss",
			Name:   "gen610",
			Type:   "mobile",
		},
	},
	{
		Desc: "Swizz ZUR700",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; ZUR700 Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "swiss",
			Name:   "zur700",
			Type:   "tablet",
		},
	},
	{
		Desc: "Zeki TB782b Tablet",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.0.4; en-US; TB782B Build/IMM76D) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.0.2.299 U3/0.8.0 Mobile Safari/534.31",
		Expect: TestDeviceRes{
			Vendor: "zeki",
			Name:   "tb782b",
			Type:   "tablet",
		},
	},
	{
		Desc: "Dragon Touch Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 4.0.4; DT9138B Build/IMM76D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "dragon touch",
			Name:   "9138b",
			Type:   "tablet",
		},
	},
	{
		Desc: "Insignia Tablet",
		Ua:   "Mozilla/5.0 (Linux; U; Android 6.0.1; NS-P08A7100 Build/MMB29M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "insignia",
			Name:   "ns-p08a7100",
			Type:   "tablet",
		},
	},
	{
		Desc: "Voice Xtreme V75",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.2.1; en-us; V75 Build/JOP40D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "voice",
			Name:   "v75",
			Type:   "mobile",
		},
	},
	{
		Desc: "LvTel V11",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; V11 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/39.0.0.0 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lvtel",
			Name:   "v11",
			Type:   "mobile",
		},
	},
	{
		Desc: "Envizen Tablet V100MD",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; V100MD Build/V100MD.20130816) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "envizen",
			Name:   "v100md",
			Type:   "tablet",
		},
	},
	{
		Desc: "Rotor Tablet",
		Ua:   "mozilla/5.0 (linux; android 5.0.1; tu_1491 build/lrx22c) applewebkit/537.36 (khtml, like gecko) chrome/43.0.2357.93 safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "rotor",
			Name:   "1491",
			Type:   "tablet",
		},
	},
	{
		Desc: "MachSpeed Tablets",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; Trio 7.85 vQ Build/Trio_7.85_vQ) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "machspeed",
			Name:   "trio 7.85 vq",
			Type:   "tablet",
		},
	},
	{
		Desc: "Trinity Tablets",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0.1; Trinity T101 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "trinity",
			Name:   "t101",
			Type:   "tablet",
		},
	},
	{
		Desc: "NextBook Next7",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; Next7P12 Build/IMM76I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		Expect: TestDeviceRes{
			Vendor: "nextbook",
			Name:   "next7p12",
			Type:   "tablet",
		},
	},
	{
		Desc: "NextBook Tablets",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0; NXA8QC116 Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "nextbook",
			Name:   "nxa8qc116",
			Type:   "tablet",
		},
	},
	{
		Desc: "Le Pan Tablets",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.2; Le Pan TC802A Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lepan",
			Name:   "tc802a",
			Type:   "tablet",
		},
	},
	{
		Desc: "Le Pan Tablets",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.2; Le Pan TC802A Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "lepan",
			Name:   "tc802a",
			Type:   "tablet",
		},
	},
	{
		Desc: "Amazon Alexa Echo Show",
		Ua:   "AlexaWebMediaPlayer/1.0.200641.0 (Linux;Android 5.1.1)",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "alexa",
			Type:   "tablet",
		},
	},
	{
		Desc: "Amazon Kindle Fire Tablet",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.4.3; en-us; KFSAWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.66 like Chrome/39.0.2171.93 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "kfsawi",
			Type:   "tablet",
		},
	},
	{
		Desc: "Amazon Kindle Fire Tablet",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.4.3; en-us; KFSAWI) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.66 like Chrome/39.0.2171.93 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "kfsawi",
			Type:   "tablet",
		},
	},
	{
		Desc: "Amazon Kindle Fire Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 9; KFMAWI Build/PS7312; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "kfmawi",
			Type:   "tablet",
		},
	},
	{
		Desc: "Amazon Fire TV",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.2; AFTB Build/JDQ39) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.173 Mobile Safari/537.22",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "b",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Amazon Fire TV",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; AFTT) AppleWebKit/537.36 (KHTML, like Gecko) Silk/86.3.20 like Chrome/86.0.4240.198 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "t",
			Type:   "smarttv",
		},
	},
	{
		Desc: "Gigaset Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 4.2.2; Gigaset QV830 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "gigaset",
			Name:   "qv830",
			Type:   "tablet",
		},
	},
	{
		Desc: "Amazon Fire 7",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; KFAUWI) AppleWebKit/537.36 (KHTML, like Gecko) Silk/80.5.3 like Chrome/80.0.3987.162 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "amazon",
			Name:   "kfauwi",
			Type:   "tablet",
		},
	},
	{
		Desc: "AT&T Radiant Core U304AA",
		Ua:   "Dalvik/2.1.0 (Linux; U; Android 9; U304AA Build/P00610)",
		Expect: TestDeviceRes{
			Vendor: "at&t",
			Name:   "u304aa",
			Type:   "mobile",
		},
	},
	{
		Desc: "Vodafone Smart Tab 4G",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.4; Vodafone Smart Tab 4G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "vodafone",
			Name:   "smart tab 4g",
			Type:   "tablet",
		},
	},
	{
		Desc: "Vodafone Smart ultra 6",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0.2; Vodafone Smart ultra 6 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.133 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "vodafone",
			Name:   "smart ultra 6",
			Type:   "tablet",
		},
	},
	{
		Desc: "Alcatel",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.2; ALCATEL A564C Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.133 Mobile Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "alcatel",
			Name:   "a564c",
			Type:   "mobile",
		},
	},
	{
		Desc: "Alcatel Go Flip",
		Ua:   "Mozilla/5.0 (Mobile; ALCATEL4044T; rv:37.0) Gecko/37.0 Firefox/37.0 KaiOS/1.0",
		Expect: TestDeviceRes{
			Vendor: "alcatel",
			Name:   "4044t",
			Type:   "mobile",
		},
	},
	{
		Desc: "Xbox One",
		Ua:   "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; Xbox; Xbox One)",
		Expect: TestDeviceRes{
			Vendor: "microsoft",
			Name:   "xbox one",
			Type:   "console",
		},
	},
	{
		Desc: "Xbox",
		Ua:   "Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Xbox)",
		Expect: TestDeviceRes{
			Vendor: "microsoft",
			Name:   "xbox",
			Type:   "console",
		},
	},
	{
		Desc: "Nvidia Shield Tablet",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; SHIELD Tablet Build/LVY48E; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/45.0.2454.19 Safari/537.36",
		Expect: TestDeviceRes{
			Vendor: "nvidia",
			Name:   "shield tablet",
			Type:   "tablet",
		},
	},
	{
		Desc: "Vivo Y52s",
		Ua:   "Mozilla/5.0 (Linux; Android 10; V2057A Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/12.10 SP-engine/2.28.0 baiduboxapp/12.10.0.10 (Baidu; P1 10) NABar/1.0",
		Expect: TestDeviceRes{
			Vendor: "vivo",
			Name:   "v2057a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Vivo X60",
		Ua:   "Mozilla/5.0 (Linux; Android 11; V2046A; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/8.8.71.0",
		Expect: TestDeviceRes{
			Vendor: "vivo",
			Name:   "v2046a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Vivo Y79A",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.2; vivo Y79A Build/N2G47H; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/9.0.14.0",
		Expect: TestDeviceRes{
			Vendor: "vivo",
			Name:   "y79a",
			Type:   "mobile",
		},
	},
	{
		Desc: "Vivo Y97",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; V1813T Build/O11019; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/9.0.14.0",
		Expect: TestDeviceRes{
			Vendor: "vivo",
			Name:   "v1813t",
			Type:   "mobile",
		},
	},
	{
		Desc: "Vivo iQOO Pro",
		Ua:   "Mozilla/5.0 (Linux; Android 11; V1916A; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/9.1.10.6",
		Expect: TestDeviceRes{
			Vendor: "vivo",
			Name:   "v1916a",
			Type:   "mobile",
		},
	},
}
