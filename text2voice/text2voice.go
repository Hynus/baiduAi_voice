package text2voice

import (
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"os"
	"time"
	"strconv"
)

var (
	BaiduAiApiKey = "grF7SveGj2LIwIYobfNBPO9n"
	BaiduAiSecretKey = "iUDZHYH8qhWWWuUc9m9A7oPcmzdzkRGH"
)

func Text2Voice(text string) error {
	timestamp := time.Now().Unix()
	filename := "E:/ACA/Learning/MyGoLang/MyGo/src/baiduAi_voice/genVoice/" + strconv.FormatInt(timestamp,10) + ".mp3"
	client := voice.NewVoiceClient(BaiduAiApiKey, BaiduAiSecretKey)
	voiceFile, err := client.TextToSpeech(text)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(voiceFile); err != nil {
		return err
	}
	return nil
}