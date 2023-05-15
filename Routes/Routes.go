package Routes

import (
	"asciiartweb/Convertation"
	Hash "asciiartweb/Hashing"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" { // hehehe
		http.NotFound(w, r)
	} else {
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allowed, 405", http.StatusMethodNotAllowed)
			return
		}
		text := r.FormValue("text")
		text = strings.ReplaceAll(text, string(rune(13)), "")
		bannername := r.FormValue("bannerfile")
		typeOfFormat := bannername
		if !Convertation.Valid(text) {
			// log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		MainArray := Convertation.SplitS(text)

		err1 := Hash.GetMD5Hash(bannername+".txt", bannername)
		if err1 != nil {
			http.Error(w, "Textfile modifying not Allowed, 405", http.StatusMethodNotAllowed)
			return
		}
		output, err := Convertation.Convert(MainArray, typeOfFormat)
			// fmt.Println(output)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		tmpl.Execute(w, output)
		
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		
		if r.Method != http.MethodGet {
			http.Error(w, "Method not Allowed, 405", http.StatusMethodNotAllowed)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", 500)
		}
	}
}
