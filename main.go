package main

import (
	t2v "baiduAi_voice/text2voice"
	"fmt"
)

func main() {
	text := "语音识别词库设置,上传识别自定义词条，训练专属识别模型。提交的词条越多、越全，识别效果提升越明显。 "
	err := t2v.Text2Voice(text)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success!!")
	}
}