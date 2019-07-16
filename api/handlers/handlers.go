package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
     io.WriteString(w, "Create user Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
   username := p.ByName("user_name")
   io.WriteString(w, username)
}
