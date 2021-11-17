package main

import(
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter,r *http.Request){
	if  err :=r.PraseForm(); err != nil {
		fmt.Fprintf(w,"PraseForm() err: %v",err)
		return 
	}
	fmt.Fprintf(w,"POST request sucessful")
	name := r.FormValue("name")
	address := r.FormValue("addrees")
	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Address= %s\n",address)
}

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotfound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"Method not supported",http.StatusNotfound)
		return
	}
	fmt.Fprintf(w,"Hello!")
}

func main(){
	fileServer := http.Fileserver(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello",helloHandler)
	}
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err!=nil {
		log.Fatal(err)
	}
}	