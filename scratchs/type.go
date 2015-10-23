package main

import (
	"fmt"
)

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func NewPlayer(id int, name string, location string, gameid int) *Player {
	return &Player{
		User:   &User{id, name, location},
		GameId: gameid,
	}
}

func main() {
	user := &Player{
		User:   &User{12, "11", "10"},
		GameId: 12,
	}
	fmt.Println(user.Greetings())
}
