package main

import (
	
	"fmt"
	 "html/template"
	"path/filepath"
    "net/http"
	"github.com/ezduzit4me/go-cognito-sdk/clients"
	// "github.com/gorilla/mux"
	"github.com/go-chi/chi/v5"
)


	
	

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "cunningham.tmpl")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		panic(err) // TODO: Remove the panic
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err) // TODO: Remove the panic
	}
}


  func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "contact")
  }

func main() {
	
	r := chi.NewRouter()
  r.Get("/", homeHandler)
  r.Get("/contact", contactHandler)
  
  r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Page not found", http.StatusNotFound)
  })
  fmt.Println("Starting the server on :3000...")
  http.ListenAndServe(":3000", r)
}
  




func log() {
	cognitoClient := clients.NewCognitoClient("ap-southeast-2", "26c2pgelskg6cq61htpck65ii3")
	//err, result := cognitoClient.SignUp("iran.rourke@dropsin.net", "5252/66Lam-4005")
	err, result, initiateAuthOutput := cognitoClient.SignIn("fred@fred.com", "5252/66Lam-4005")
	//err, result := cognitoClient.ConfirmSignUp("iran.rourke@dropsin.net", "")
	if err != nil {
		fmt.Print("Error" + fmt.Sprint("%s", err))
		panic(err)
	}
	fmt.Printf("Result:" + fmt.Sprintf("%s \n %s", result, *initiateAuthOutput.AuthenticationResult.IdToken))

}
