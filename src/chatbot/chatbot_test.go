package chatbot

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Chatbot_Input_company_id_1234_And_message_And_Timestamp_2018_08_12T_03_59_58_Dot_855Z(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/chatbot/api/message", strings.NewReader("company_id=1234&message=หอพักย่านลาดพร้าวซอย1&timestamp=2018-08-12T 03:59:58.855Z"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	expected := `{"message":"ละออเพลส,เรือนสุดสุข,ละอ่อน","timestamp":"2018-08-12T 03:59:59.855Z"}`

	Chatbot(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if expected != string(body) {
		t.Errorf("should be %s but got %s", expected, string(body))
	}

}
