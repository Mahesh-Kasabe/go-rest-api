package main

import ( 
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"io/ioutil"

)

type Article struct {
	Id 	string `json: "id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article


func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content" },
		Article{Id: "2", Title: "Hello2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

func returnArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnsingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key{
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	for index, article := range Articles {
		if article.Id == id{
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter()
	
	myRouter.HandleFunc("/", returnArticles)
	myRouter.HandleFunc("/article", createArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnsingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}