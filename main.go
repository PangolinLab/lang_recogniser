package lang_recogniser

import "github.com/abadojack/whatlanggo"
import "github.com/liuzl/gocc"

var t2s, _ = gocc.New("t2s")

func DetectChineseScript(text string) string {

	simplified, _ := t2s.Convert(text)

	if simplified == text {
		return "zh-Hans"
	}

	return "zh-Hant"
}

func DetectLangTag(text string) string {

	info := whatlanggo.Detect(text)

	switch info.Lang {

	case whatlanggo.Cmn:
		return DetectChineseScript(text)

	case whatlanggo.Eng:
		return "en"

	case whatlanggo.Jpn:
		return "ja"

	case whatlanggo.Kor:
		return "ko"

	case whatlanggo.Pes:
		return "fa-IR"

	case whatlanggo.Rus:
		return "ru-RU"

	case whatlanggo.Tur:
		return "tr-Latn"

	case whatlanggo.Arb:
		return "ar-Arab"

	case whatlanggo.Vie:
		return "vi-Latn"

	case whatlanggo.Amh:
		return "am-Ethi"

	case whatlanggo.Uzb:
		return "uz-Latn"

	default:
		return "unk"
	}
}

func Recogniser(s string) []string {
	var langSet []string

	for _, obj := range s {
		langSet = append(langSet, DetectLangTag(string(obj)))
	}

	return langSet
}
