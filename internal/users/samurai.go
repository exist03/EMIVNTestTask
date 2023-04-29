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
	return fmt.Sprintf("Nickname: %s\nTG Username: %s\nOwner: %s\n", s.Nickname, s.TelegramUsername, s.Owner)
}
