package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"oceane/dealgen"
	"strconv"

	"github.com/gorilla/mux"
)

const EMPTYDESK = "..."

var version = "undefined"

func setValue(json bool, mode, ite, vul, mask string) string {
	var err error
	var iMode, iIte, iVul int
	iMode, err = strconv.Atoi(mode)
	if err != nil {
		iMode = 0
	}
	iIte, err = strconv.Atoi(ite)
	if err != nil {
		iIte = 1
	}
	iVul, err = strconv.Atoi(vul)
	if err != nil {
		iVul = 0
	}
	if mask == "" {
		mask = EMPTYDESK
	}
	var sh dealgen.Random
	if json {
		if mode == "" {
			mode = "1"
		}
		return dealgen.PbnDealJson(sh, mode, mask)
	} else {
		return dealgen.PbnDeal(sh, iMode, iIte, 0, 0, iVul, mask)
	}
}

func handlerPbn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		result := ""
		err := r.ParseForm()
		if err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)

			return
		}

		result = setValue(false, r.FormValue("mode"), r.FormValue("ite"), r.FormValue("vul"), r.FormValue("mask"))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(result))
	}
}

func handlerPbnJson(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		result := ""
		err := r.ParseForm()
		if err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)

			return
		}

		result = setValue(true, r.FormValue("mode"), r.FormValue("ite"), r.FormValue("vul"), r.FormValue("mask"))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(result))
	}
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(version))
}

func main() {
	portPtr := flag.String("p", "3000", "API port")
	flag.Parse()

	port := *portPtr
	println("Listen " + port)

	r := mux.NewRouter()

	r.HandleFunc("/api/v2/version", handlerVersion)
	r.HandleFunc("/api/v2/pbn", handlerPbn)
	r.HandleFunc("/api/v2/pbnjson", handlerPbnJson)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
