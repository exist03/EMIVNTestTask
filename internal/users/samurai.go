package users

import "fmt"

type Samurai struct {
	TurnOver         float64
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
	return fmt.Sprintf("Nickname: %s\nTG Username: %s\nTurnover: %.2f\nOwner: %s\n", s.Nickname, s.TelegramUsername, s.TurnOver, s.Owner)
}
