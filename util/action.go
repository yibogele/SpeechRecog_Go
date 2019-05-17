package util

import (
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type ActionParam struct {
	Action string `json:"action"`
	Target string `json:"target"`
}

func GetActionParam(text string) ActionParam {
	for _, action := range ReadConf(GetCurrentDirectory() + ActionConfFile) {
		if strings.HasPrefix(text, action) {
			idx := utf8.RuneCountInString(action)
			target := string([]rune(text)[idx:])

			return getActionParamTrans(action, target)
		}
	}

	return ActionParam{}
}

func getActionParamTrans(action, target string) ActionParam {
	switch action {
	case "选择项目":
		return ActionParam{action, pinyin.NewDict().Convert(target, "").None()}
	case "选择日期":
		re := regexp.MustCompile(`((\d+)年)?(\d+)月(\d+)`)
		match := re.FindStringSubmatch(target)
		var str string

		if match[1] != "" {
			str = strings.Join(match[2:], "-")
		} else {
			tms := []string{strconv.Itoa(time.Now().Year())}
			tms = append(tms, match[3:]...)
			str = strings.Join(tms, "-")
		}
		return ActionParam{action, str}
	case "选择":
		switch {
		case strings.HasPrefix(target, "人员"):
			name := string([]rune(target)[len([]rune("人员")):])
			return ActionParam{action,
				pinyin.NewDict().Convert(name, "").None()}
		case strings.HasPrefix(target, "车牌号"):
			carId := string([]rune(target)[len([]rune("车牌号")):])
			return ActionParam{action,
				strings.Map(func(r rune) rune {
					if v, ok := map[rune]rune{
						'一': '1', '二': '2', '三': '3', '四': '4', '五': '5',
						'六': '6', '七': '7', '八': '8', '九': '9', '零': '0',
					}[r]; !ok {
						return r
					} else {
						return v
					}
				}, carId)}
		case strings.HasPrefix(target, "设施"):
			dev := string([]rune(target)[len([]rune("设施")):])
			return ActionParam{action, dev}
		}
	case "勾选":
		switch target {
		case "公测":
			return ActionParam{action, "公厕"}

		}
	}
	return ActionParam{action, target}
}
