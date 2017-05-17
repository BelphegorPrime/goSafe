package main

type Website struct {
	ID		int
	Url     	string
	UserName     	string
	Password 	string
}

func (w Website) ToString() string {
	return w.Url+" "+w.UserName+" "+w.Password
}
