package users

import "fmt"

type Daimyo struct {
	CardList         []Card             //by nickname
	SamuraiList      map[string]Samurai //same
	Owner            string
	TelegramUsername interface{}
	Nickname         string
}

func (d Daimyo) String() string {
	return fmt.Sprintf("TG Username: %s\nNickname: %s\nOwner: %s\n\n", d.TelegramUsername, d.Nickname, d.Owner)
}
