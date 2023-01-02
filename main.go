package main

import (
	
	"fmt"
	
	"path/filepath"
    "net/http"
	"github.com/ezduzit4me/go-cognito-sdk/clients"
	"github.com/ezduzit4me/go-cognito-sdk/views"
	"github.com/ezduzit4me/go-cognito-sdk/controllers"
	"github.com/go-chi/chi/v5"
)


	
func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := views.Parse(filepath)
	if err != nil {
		fmt.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}
  



func main() {
	
	r := chi.NewRouter()
    
	tpl, err := views.Parse(filepath.Join("templates", "cunningham.tmpl"))
	if err != nil {
	  panic(err)
	}
  r.Get("/", controllers.StaticHandler(tpl))

  tpl, err = views.Parse(filepath.Join("templates", "committee.tmpl"))
	if err != nil {
	  panic(err)
	}
  r.Get("/committee", controllers.StaticHandler(tpl))

  tpl, err = views.Parse(filepath.Join("templates", "contact.tmpl"))
  if err != nil {
	panic(err)
  }
  r.Get("/contact", controllers.StaticHandler(tpl))

  tpl, err = views.Parse(filepath.Join("templates", "maintenance.tmpl"))
  if err != nil {
	panic(err)
  }
  r.Get("/maintenance", controllers.StaticHandler(tpl))

  tpl, err = views.Parse(filepath.Join("templates", "manager.tmpl"))
	if err != nil {
	  panic(err)
	}
  r.Get("/manager", controllers.StaticHandler(tpl))

  tpl, err = views.Parse(filepath.Join("templates", "utility.tmpl"))
	if err != nil {
	  panic(err)
	}
  r.Get("/utility", controllers.StaticHandler(tpl))

  tpl, err = views.Parse(filepath.Join("templates", "pet.tmpl"))
	if err != nil {
	  panic(err)
	}
  r.Get("/pet", controllers.StaticHandler(tpl))

  tpl, err = views.Parse(filepath.Join("templates", "flood.tmpl"))
	if err != nil {
	  panic(err)
	}
  r.Get("/flood", controllers.StaticHandler(tpl))


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
