package entity

import "time"

type User struct {
	ID        uint
	FullName  string
	Handle    string
	Email     string
	Password  string
	BirthDate time.Time
	AvatarURL string
	//add followers, followings, posts
}
