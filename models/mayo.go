package models

import (
	"errors"
)

type Mayo struct {
	Id       int    `json:"maYoId" orm:"column(mayo_id)"`
	Username string `json:"username" orm:"column(mayo_name)"`
	Password string `json:"password" orm:"column(mayo_pass)"`
	Phone    string `json:"phone" orm:"column(mayo_phone)"`
	Sign     string `json:"sign" orm:"column(mayo_sign)"`
	Sex      int    `json:"sex" orm:"column(mayo_sex)"`
	Emotion  int    `json:"emotion" orm:"column(mayo_emotion)"`
	Birthday string `json:"birthday" orm:"column(mayo_birthday)"`
	Location string `json:"location" orm:"column(mayo_location)"`
	School   string `json:"school" orm:"column(mayo_school)"`
	Trade    string `json:"trade" orm:"column(mayo_trade)"`
	Register string `json:"register" orm:"column(mayo_register)"`
	State    string `json:"state" orm:"column(mayo_state)"`
}

func AddUser(u Mayo) int {

	return u.Id
}

func GetUser(uid string) (u *Mayo, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("Mayo not exists")
}

func UpdateUser(uid string, uu *Mayo) (a *Mayo, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("Mayo Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
