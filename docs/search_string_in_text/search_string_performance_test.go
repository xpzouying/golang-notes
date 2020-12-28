package main

import (
	"regexp"
	"strings"
	"testing"
)

var text = `登陆月球是指人类控制无人太空船或者直接驾驶太空船降落在月球上。到目前为止，只有美国、前苏联和中国成功把探测器送到月球表面，只有美国成功派出宇航员登陆月球表面。当中，美国在阿波罗计划执行了六次载人登月任务。苏联月球2号于1959年9月撞击月球，是首个登陆月球的探测器。美国阿波罗11号于1969年7月成功登陆月球，太空人尼尔·阿姆斯特朗和巴兹·奥尔德林成为历史上最早登陆月球的人类。法国小说家儒勒·凡尔纳的1865年科幻小说《从地球到月球》则是人类出现最早有关登陆月球的概念之一。美国在1972年12月最后一次离开月球表面，是迄今为止，唯一一个成功进行月球任务的国家。2019年1月3日中国嫦娥四号飞船首次降落在月球的另一侧。在这之前，所有的软着陆都在月球的正面进行。`

var ss = []string{
	"日本",
	"月球",
	"美国",
	"苏联",
	"中国",
}

func containSubstr(text string, ss []string) bool {
	for _, s := range ss {

		if exists := strings.Contains(text, s); exists {
			return true
		}
	}

	return false
}

func BenchmarkStringContain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		containSubstr(text, ss)
	}
}

func BenchmarkRegularExpression(b *testing.B) {

	reg, err := regexp.Compile("(?:月球|美国|中国|苏联|日本)")
	if err != nil {
		b.Error("compile regexp error")
	}

	for i := 0; i < b.N; i++ {

		if exists := reg.MatchString(text); !exists {
			b.Error("not exists")
		}
	}
}
