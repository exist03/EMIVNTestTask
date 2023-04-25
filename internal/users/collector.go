package users

import "fmt"

type Collector struct {
	TelegramUsername string
	Nickname         string
}

func (c Collector) String() string {
	return fmt.Sprintf("TG Username: %s\nNickname: %s\n", c.TelegramUsername, c.Nickname)
}
