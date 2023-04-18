package users

type Samurai struct {
	TurnOver         float64
	TelegramUsername string
	Nickname         string
	Owner            Daimyo
}

func CreateSamurai(username, nickname string) Samurai {
	return Samurai{Nickname: nickname, TelegramUsername: username}
}

func (s *Samurai) SetTurnover(num float64) {
	s.TurnOver = num
}

func (s *Samurai) SetOwner(daimyo *Daimyo) {
	s.Owner = *daimyo
}
