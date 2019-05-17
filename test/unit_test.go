package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	//"github.com/mozillazg/go-pinyin"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"log"
)
import "../util"

/**
 * created: 2019/5/16 17:28
 * By Will Fan
 */
func ExampleJson() {
	//fmt.Println(util.ReadConf("config/verbs.conf"))

	jsonData, err := json.Marshal(util.ActionParam{"查找", "人员"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	// Output:
}
func ExampleAP() {

	util.GetActionParam("切换车辆")
	// Output:
}
func ExamplePinYin() {

	//hans := "音乐"
	//a := pinyin.NewArgs()
	//fmt.Println(pinyin.LazyPinyin(hans, a))
	dict := pinyin.NewDict()
	s := dict.Convert(`音乐乐山`, "").None()
	fmt.Println(s)
	// Output:
}
func ExampleCarId() {

	carid := strings.Map(func(r rune) rune {
		if v, ok := map[rune]rune{
			'一': '1', '二': '2', '三': '3', '四': '4', '五': '5',
			'六': '6', '七': '7', '八': '8', '九': '9', '零': '0',
		}[r]; !ok {
			return r
		} else {
			return v
		}
	}, "京N二三四五六")
	fmt.Println(carid)
	// Output:
}

func ExampleT() {

	re := regexp.MustCompile(`((\d+)年)?(\d+)月(\d+)`)

	match := re.FindStringSubmatch("2019年5月17号")
	//fmt.Printf("%v\n%d\n", match, len(match))
	for i, v := range match {
		if v != "" {
			fmt.Printf("%d:%v\n", i, v)
		}
	}
	if match[1] != "" {
		str := strings.Join(match[2:], "-")
		fmt.Println(str)
	} else {
		tms := []string{strconv.Itoa(time.Now().Year())}
		tms = append(tms, match[3:]...)
		str := strings.Join(tms, "-")
		fmt.Println(str)
	}
	// Output:

	//fmt.Println(re.ReplaceAllStringFunc("2019年5月17日", repl))
	//fmt.Println(re.ReplaceAllStringFunc("2019年5月17号", repl))
	//fmt.Println(re.ReplaceAllStringFunc("2019年5月17", repl))
}
