package honeypot

import (
	"fmt"
	"net/http"

	"github.com/edoardottt/boggart/pkg/template"
)

func genericWriter(w http.ResponseWriter, req *http.Request, response string) {
	fmt.Fprint(w, response)
}

func defaultRootWriter(w http.ResponseWriter, req *http.Request, response string, defaultResponse string) {
	fmt.Println(">" + req.URL.Path + "<")
	if req.URL.Path != "/" {
		fmt.Println(req.URL.Path)
		fmt.Fprint(w, defaultResponse)
	} else {
		fmt.Println(response)
		fmt.Fprint(w, response)
	}
}

//RawHoneypot
func RawHoneypot(tmpl template.Template) {

	//This variable says if the root endpoint ("/")
	//is defined through the template.
	var rootExists = template.RootEndpointExists(tmpl)

	//IF A REQUEST WITH ENDPOINT == "/" EXISTS:
	if rootExists {
		for _, request := range tmpl.Requests {
			if request.Endpoint != "" && request.Endpoint != "/" {
				http.HandleFunc(request.Endpoint, func(w http.ResponseWriter, r *http.Request) {
					genericWriter(w, r, request.Content)
				})
			} else {
				//HANDLE DEFAULT THERE
				if request.Id != "default" {
					fmt.Println(request.Id + " " + request.Content)
					http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
						defaultRootWriter(w, r, request.Content, template.DefaultResponse(tmpl))
					})
				}
			}
		}
	} else {
		//OTHERWISE
		for _, request := range tmpl.Requests {
			http.HandleFunc(request.Endpoint, func(w http.ResponseWriter, r *http.Request) {
				genericWriter(w, r, request.Content)
			})
		}
		//DEFINE "/" WITH DEFAULT RESPONSE
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			genericWriter(w, r, template.DefaultResponse(tmpl))
		})
	}

	http.ListenAndServe(":8090", nil)
}
