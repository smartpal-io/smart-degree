package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"encoding/hex"
)

var (
	smartDegreeService SmartDegreeService
)

func main() {
	smartDegreeService = NewEthSmartDegreeService()
	r := mux.NewRouter()
	r.HandleFunc("/degree/register", RegisterHandler).Methods(http.MethodPost)
	r.HandleFunc("/degree/student/{studentId}/", GetDegreeHashHandler).Methods(http.MethodGet)
	r.HandleFunc("/degree/verify", VerifyHandler).Methods(http.MethodPost)
	http.ListenAndServe(":8081", r)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

func GetDegreeHashHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId, err := strconv.ParseInt(vars["studentId"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	degreeHash, err := smartDegreeService.GetDegreeHash(studentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(hex.EncodeToString(degreeHash)))
	return
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {

}
