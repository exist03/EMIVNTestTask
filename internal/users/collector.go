package users

import "fmt"

type Collector struct {
	TelegramUsername interface{}
	Nickname         string
}

func (c Collector) String() string {
	return fmt.Sprintf("TG Username: %s\nNickname: %s\n\n", c.TelegramUsername, c.Nickname)
}
