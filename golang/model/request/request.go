package req_mod

type Register struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type Login struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type Question struct {
	Username	string	`json:"username"`
	Data		string	`json:"data"`
}

type Answer struct {
	ID		string	`json:"id"`
	Data	string	`json:"data"`
}

type DeleteQuestion struct {
	ID	string	`json:"id"`
}

type DeleteAnswer struct {
	ID	string	`json:"id"`
}