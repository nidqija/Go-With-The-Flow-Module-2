package main

import (

	"fmt"
	"net/http"
	// to be filled by participants
)


type LoginData struct {
	// to be filled by participants
}

type SystemStats struct {
	// to be filled by participants
}


func fetchStats() SystemStats {

	// to be filled by participants

	return SystemStats{}
}


func main() {


	fmt.Println("Starting server on port 8000")
	// variables section : to be filled by participants



	


	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
		
		// to be filled by participants

	})



	http.HandleFunc("/signin" , func(w http.ResponseWriter , r *http.Request){


		// to be filled by participants

	})

//====================================================================================

	http.HandleFunc("/dashboard" , func(w http.ResponseWriter , r *http.Request){

		// to be filled by participants

	})

	http.HandleFunc("/stats" , func(w http.ResponseWriter , r *http.Request){
		// to be filled by participants

	})

	http.HandleFunc("/memory" , func(w http.ResponseWriter , r *http.Request){
		// to be filled by participants
	})

	http.HandleFunc("/api/live-stats" , func(w http.ResponseWriter , r *http.Request){
		// to be filled by participants
	})

	http.HandleFunc("/system" , func(w http.ResponseWriter , r *http.Request){
		// to be filled by participants
	})




	if err := http.ListenAndServe(":8000" , nil ); err != nil {
		fmt.Println("Error starting server : " , err)
	}
	

}