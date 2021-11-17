package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

type Visit struct{}

func NewVisitMiddleware() *Visit {
	return &Visit{}
}

func (v *Visit) Ip(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		addr := strings.Split(request.RemoteAddr, ":")[0]
		fmt.Println(addr + "访问了")
		next(writer, request)
	}
}
