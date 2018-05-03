package login

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	type Mess struct {
		Error            int64  `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
	body, _ := (&Url{}).get_contents("https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=101466300&client_secret=0104260a8f8faac3900cbf184bae55f5&redirect_uri=http%3a%2f%2fwww.fightcoder.com%2f%23%2fproblem%2fopen&code=20C72F76B2E72DC2243D0C0C0C18A221")
	mess := &Mess{}
	if strings.Contains(body, "callback") {
		start := strings.Index(body, "(")
		end := strings.Index(body, ")")
		body = body[start+1 : end]
		if err := json.Unmarshal([]byte(body), mess); err != nil {
			fmt.Println(err)
		}
		fmt.Println(mess)
	}

}
