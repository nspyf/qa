package serv

import (
	"errors"
	"mymod/model/db"
	"mymod/model/request"
	"mymod/model/response"
	"mymod/util"
	"strconv"
)

const (
	saltPart = "';1.rfa'9sFS31;['5.;q3O5ik[p13q130sDVa235gHEW8-0hfr[9sDi091351R4as23f1"
)

func Register(req *req_mod.Register) error {
	req.Password = util.SHA256(req.Password+saltPart+req.Username)

	getData := &db_mod.User{}
	tarUsername := &db_mod.User{
		Username: req.Username,
	}
	e.RetrieveByStruct(getData,tarUsername)
	if getData.Username !="" {
		return errors.New("username has registered")
	}

	createData := &db_mod.User{
		Username: req.Username,
		Password: req.Password,
	}
	e.Creat(createData)
	return nil
}

func Login(req *req_mod.Login) (string,error) {
	req.Password = util.SHA256(req.Password+saltPart+req.Username)

	getData := &db_mod.User{}
	tarUsername := &db_mod.User{
		Username: req.Username,
	}
	e.RetrieveByStruct(getData,tarUsername)

	if getData.Username == "" || getData.Password != req.Password {
		return "",errors.New("username or password error")
	}

	token := JwtObj.GenerateToken(map[string]interface{}{
		"ID": getData.ID,
		"Username": getData.Username,
	},7200)
	return token,nil
}

func Information(username string) (interface{},error) {
	if username == "" {
		return nil,errors.New("user can't be blank")
	}

	getUser := &db_mod.User{}
	tarUser := &db_mod.User{
		Username: username,
	}
	e.RetrieveByStruct(getUser,tarUser)
	if getUser.ID == 0 {
		return nil,errors.New("user don't exist")
	}

	getQuestion := &[]db_mod.Question{}
	tarQuestion := &db_mod.Question{
		UserID: getUser.ID,
	}
	e.RetrieveArrByStruct(getQuestion,tarQuestion)

	data := make([]res_mod.QA,0)
	questionLen := len(*getQuestion)
	for i:=0;i<questionLen;i++ {
		getAnswer := &[]db_mod.Answer{}
		tarAnswer := &db_mod.Answer{
			Question: (*getQuestion)[i].ID,
		}
		e.RetrieveArrByStruct(getAnswer,tarAnswer)

		answer := make([]string,0)
		answerLen := len(*getAnswer)
		for j:=0;j<answerLen;j++ {
			answer = append(answer,(*getAnswer)[j].Data)
		}

		newQA := &res_mod.QA{
			QuestionID: strconv.Itoa(int((*getQuestion)[i].ID)),
			Question:   (*getQuestion)[i].Data,
			Answer:     answer,
		}
		data = append(data,*newQA)
	}

	return data,nil
}

func Question(req *req_mod.Question) error {
	if req.Data == "" || req.Username == "" {
		return errors.New("question or username can't be blank")
	}

	getUser := &db_mod.User{}
	tarUser := &db_mod.User{
		Username: req.Username,
	}
	e.RetrieveByStruct(getUser,tarUser)
	if getUser.ID == 0 {
		return errors.New("user don't exist")
	}

	createData := &db_mod.Question{
		UserID: getUser.ID,
		Data: req.Data,
	}
	e.Creat(createData)

	return nil
}

func Answer(ID interface{},req *req_mod.Answer) error {
	if req.Data == "" || req.ID == "" {
		return errors.New("answer or id can't be blank")
	}

	questionIdUint, err := strconv.ParseUint(req.ID,10,32)
	if err != nil {
		return errors.New("id error")
	}

	getQuestion := &db_mod.Question{}
	e.RetrieveByID(getQuestion,uint(questionIdUint))
	if getQuestion.ID == 0 {
		return errors.New("question don't exist")
	}


	UserIdUint, err := strconv.ParseUint(strconv.Itoa(int(uint(ID.(float64)))),10,32)
	if err != nil {
		return nil
	}

	if getQuestion.UserID != uint(UserIdUint) {
		return errors.New("no right to answer")
	}

	createData := &db_mod.Answer{
		Question: uint(questionIdUint),
		Data: req.Data,
	}
	e.Creat(createData)

	return nil
}