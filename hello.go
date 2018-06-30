package hello

import (
	"fmt"
	"time"
)

var SupportedLangs = map[string]string{
	"en-US": "Hello, World!",
	"ru-RU": "Здравствуй, Мир!",
	"zh-CN": "你好，世界!",
	"fr-FR": "Bonjour le Monde!",
}

func World() error {
	return WorldIn("en-US")
}

func WorldIn(lang string) error {
	if greeting, found := SupportedLangs[lang]; found {
		_, err := fmt.Println(greeting)
		time.Sleep(1 * time.Second)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("failed to hello: lang '%s' is not supported", lang)
}

