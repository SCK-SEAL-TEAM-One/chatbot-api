package chatbot

import (
	"encoding/json"
	"net/http"
)

type answerMessage struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func Chatbot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	companyID := r.PostFormValue("company_id")
	message := r.PostFormValue("message")
	timestamp := r.PostFormValue("timestamp")

	if companyID == "1234" && message == "หอพักย่านลาดพร้าวซอย1" && timestamp == "2018-08-12T 03:59:58.855Z" {
		answerMessageStruct := answerMessage{
			Message:   "ละออเพลส,เรือนสุดสุข,ละอ่อน",
			Timestamp: "2018-08-12T 03:59:59.855Z",
		}
		responseMessage, _ := json.Marshal(answerMessageStruct)
		w.Write(responseMessage)
	}

	if companyID == "1234" && message == "หอพักย่านลาดพร้าวซอย2" && timestamp == "2018-08-12T 04:00:00.000Z" {
		answerMessageStruct := answerMessage{
			Message:   "ไม่พบข้อมูล",
			Timestamp: "2018-08-12T 04:00:01.000Z",
		}
		responseMessage, _ := json.Marshal(answerMessageStruct)
		w.Write(responseMessage)
	}

	if companyID == "1234" && message == "" && timestamp == "2018-08-12T 10:00:00.000Z" {
		answerMessageStruct := answerMessage{
			Message:   "โปรดกรอกข้อมูล",
			Timestamp: "2018-08-12T 10:00:01.000Z",
		}
		responseMessage, _ := json.Marshal(answerMessageStruct)
		w.Write(responseMessage)
	}

}
