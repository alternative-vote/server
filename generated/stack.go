package generated

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/urfave/negroni"
)

//Stack blah blah
func Stack() *negroni.Negroni {
	stack := negroni.New()
	stack.Use(negroni.HandlerFunc(unexpectedErrors))
	stack.Use(negroni.HandlerFunc(addTrailingSlash))
	stack.Use(negroni.HandlerFunc(requestIdMiddleware))
	stack.Use(negroni.HandlerFunc(contentNegotiation))

	for _, v := range middleware {
		stack.Use(negroni.HandlerFunc(v))
	}

	stack.UseHandler(Router())

	return stack
}

func unexpectedErrors(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	defer func() {
		err := recover()

		//check to see if we have an error here
		if err == nil {
			return
		}

		//check to see if it's a structured API error
		apiError, ok := err.(*APIError)
		if !ok {
			//well this was unexpected
			fmt.Println(err)
			fmt.Println(string(debug.Stack()[:2000]))
			//let's make this error a structured one
			apiError = HttpError(500)
		}

		//send the response to the client
		res.WriteHeader(apiError.StatusCode)
		buffer, _ := json.Marshal(apiError)
		res.Write(buffer)
		return
	}()
	next(res, req)
}

func removeTrailingSlash(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	req.URL.Path = strings.TrimSuffix(req.URL.Path, "/")
	next(res, req)
}

func addTrailingSlash(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if !strings.HasSuffix(req.URL.Path, "/") {
		req.URL.Path += "/"
	}
	next(res, req)
}

func contentNegotiation(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	res.Header().Add("content-type", "application/json")
	next(res, req)
}

func requestIdMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	requestID := req.Header.Get("request-id")
	if requestID == "" {
		requestID = uuid.NewV4().String()
		req.Header.Set("request-id", requestID)
	}

	req = req.WithContext(context.WithValue(req.Context(), "request-id", requestID))

	next(res, req)
}
