package dbuaparser_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/dlclark/regexp2"
)

func TestSingle(t *testing.T) {
	ua := "Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 635) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537"
	regStr := `((?:avr32|ia64(?=;))|68k(?=\))|\barm(?=v(?:[1-7]|[5-7]1)l?|;|eabi)|(?=atmel )avr|(?:irix|mips|sparc)(?:64)?\b|pa-risc)`
	exec := reg(regStr)
	strList := exec.FindStringMatch(ua)

	fmt.Println(len(strList), strList[1])
}

func reg(v string) IRegexp {
	res := IRegexp{}
	reg, err := regexp.Compile(v)

	if err != nil {
		reg2, err := regexp2.Compile(v, regexp2.IgnoreCase)

		if err != nil {
			panic(err)
		}

		fmt.Println("使用: reg2")
		res.reg2 = reg2
	} else {
		fmt.Println("使用: reg")
		res.reg = reg
	}
	return res
}

type IRegexp struct {
	reg2 *regexp2.Regexp
	reg  *regexp.Regexp
}

func (i *IRegexp) FindStringMatch(s string) []string {
	if i.reg != nil {
		return i.reg.FindStringSubmatch(s)
	} else {
		match, _ := i.reg2.FindStringMatch(s)
		if match == nil {
			return []string{}
		}
		matchList := match.Groups()
		matchLen := match.GroupCount()
		strList := make([]string, matchLen)
		for i := 0; i < matchLen; i++ {
			strList[i] = matchList[i].String()
		}
		return strList
	}
}
