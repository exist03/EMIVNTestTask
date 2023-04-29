package users

import "fmt"

type Samurai struct {
	TelegramUsername string
	Nickname         string
	Owner            string
}

func CreateSamurai(username, nickname string) Samurai {
	return Samurai{Nickname: nickname, TelegramUsername: username}
}

func (s *Samurai) SetOwner(owner string) {
	s.Owner = owner
}

func (s Samurai) String() string {
	return fmt.Sprintf("Owner: %s\nNickname: %s\nTG Username: %s\n\n", s.Owner, s.Nickname, s.TelegramUsername)
}
