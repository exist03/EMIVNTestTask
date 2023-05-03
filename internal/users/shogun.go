package users

import (
	"fmt"
)

type Shogun struct {
	DaimyoList       map[string]Daimyo
	TelegramUsername string
	Nickname         string
}

func (s Shogun) String() string {
	return fmt.Sprintf("TG Username: %s\nNickname: %s\n", s.TelegramUsername, s.Nickname)
}
