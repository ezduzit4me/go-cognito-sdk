package main

import (
	
	"fmt"
	 "html/template"
	"path/filepath"
    "net/http"
	"github.com/ezduzit4me/go-cognito-sdk/clients"
	
	"github.com/go-chi/chi/v5"
)


	
func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	  tpl, err := template.ParseFiles(filepath)
	  if err != nil {
		  fmt.Printf("processing template: %v", err)
		  http.Error(w, "There was an error processing the template.", http.StatusInternalServerError)
		  return
	  }
	  err = tpl.Execute(w, nil)
	  if err != nil {
		  fmt.Printf("executing template: %v", err)
		  http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		  return
	  }
  }
  

func homeHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "cunningham.tmpl")
	executeTemplate(w,tplPath)
}

func floodHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "flood.tmpl")
	executeTemplate(w,tplPath)
}

func petHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "pet.tmpl")
	executeTemplate(w,tplPath)
}

func utilityHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "utility.tmpl")
	executeTemplate(w,tplPath)
}

func managerHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "manager.tmpl")
	executeTemplate(w,tplPath)
}

func maintenanceHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "maintenance.tmpl")
	executeTemplate(w,tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "contact.tmpl")
	executeTemplate(w,tplPath)
}


func committeeHandler(w http.ResponseWriter, r *http.Request) {
	
	tplPath := filepath.Join("templates", "committee.tmpl")
	executeTemplate(w,tplPath)
}

func main() {
	
	r := chi.NewRouter()
  r.Get("/", homeHandler)
  r.Get("/committee", committeeHandler)
  r.Get("/contact", contactHandler)
  r.Get("/maintenance", maintenanceHandler)
  r.Get("/manager", managerHandler)
  r.Get("/utility", utilityHandler)
  r.Get("/pet", petHandler)
  r.Get("/flood", floodHandler)
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
