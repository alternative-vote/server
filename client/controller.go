package client

import "net/http"

type ClientController struct {
}

func (o *ClientController) Index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("this could be the static content of your dreams"))
}
