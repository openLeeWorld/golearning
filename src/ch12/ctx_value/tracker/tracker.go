package tracker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// 서비스에서 서비스로 추적하고 GUID가 포함된 로그를 생성하나 컨텍스트 관련 GUID 구현
type guidKey int // ctx값을 위한 키 값은 유일하면 비교가능해야함

const key guidKey = 1 // 타입의 노출되지 않는 상수로 선언

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key, guid)
}

func guidFormContext(ctx context.Context) (string, bool) {
	g, ok := ctx.Value(key).(string)
	return g, ok
}

// Middleware set GUID (extract or create) and make new request with update
// context, execute call chain
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if guid := req.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.New().String())
		}
		req = req.WithContext(ctx) // 현재 ctx에서 guid 추가한 req 변수 받음
		h.ServeHTTP(rw, req)
	})
}

type Logger struct{}

func (Logger) Log(ctx context.Context, message string) {
	if guid, ok := guidFormContext(ctx); ok {
		message = fmt.Sprintf("GUID: %s - %s", guid, message)
	}
	// do logging
	fmt.Println(message)
}

func Request(req *http.Request) *http.Request {
	ctx := req.Context()
	if guid, ok := guidFormContext(ctx); ok {
		req.Header.Add("X-GUID", guid)
	}
	return req
}
