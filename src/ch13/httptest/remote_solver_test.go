package httptest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {
	type info struct {
		expression string
		code       int
		body       string
	} // 입력과 출력을 저장하는 타입
	var io info // 현재 입력과 출력이 저장되는 변수
	// 가상의 원격 서버를 설정하고 RemoteSolver의 인스턴스를 구성
	server := httptest.NewServer( // 임의의 사용하지 않는 포트의 HTTP 서버 생성, 시작
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			expression := req.URL.Query().Get("expression")
			if expression != io.expression {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("invalid expression: " + io.expression))
				return
			}
			rw.WriteHeader(io.code)
			rw.Write([]byte(io.body))
		}))
	defer server.Close() // 테스트 이므로 완료시 반드시 서버 닫음
	rs := RemoteSolver{
		MathServerURL: server.URL,      // 지정된 서버 URL
		Client:        server.Client(), // 지정된 http.Client
	}
	// 나머지 부분은 데이터를 정의하고 서버에 전달하는 것
	data := []struct {
		name   string
		io     info
		result float64
		errMsg string
	}{
		{"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22, ""},
		{"case2", info{"( 2 + 2 ) * 10", http.StatusOK, "40"}, 40, ""},
		{"case3", info{"( 2 + 2 * 10", http.StatusBadRequest,
			"invalid expression: ( 2 + 2 * 10"},
			0, "invalid expression: ( 2 + 2 * 10"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				t.Errorf("io `%f`, got `%f`", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("io error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
