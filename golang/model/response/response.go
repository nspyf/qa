package res_mod

type Answer struct {
	ID		string	`json:"id"`
	Data	string	`json:"data"`
}

type QA struct {
	ID		string		`json:"id"`
	Data	string		`json:"data"`
	Answer	[]Answer	`json:"answer"`
}