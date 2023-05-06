package users

import "fmt"

type Samurai struct {
	TelegramUsername interface{}
	Nickname         string
	Owner            string
}

func (s Samurai) String() string {
	return fmt.Sprintf("Owner: %s\nNickname: %s\nTG Username: %s\n\n", s.Owner, s.Nickname, s.TelegramUsername)
}
