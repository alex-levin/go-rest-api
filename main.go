package main
import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

/*
 This is a slice that is similar to an array. The difference is that
 when you want to use arrays in Golang you need to define the length.
 This is why we use a slice, we also tell it that it will contain users.
*/
var users []User

func getData() {
	users = append(users, User{ID: "1", Username: "alex", Password: "password123", Email: "alex@google.com"})
	users = append(users, User{ID: "2", Username: "albert", Password: "password456", Email: "albert@google.com"})	
    // users = []User{
    //     User{ID: "1", Username: "alex", Password: "password123", Email: "alex@google.com"},
    //     User{ID: "2", Username: "albert", Password: "password456", Email: "albert@google.com"},
    // }
}


func main() {
  fmt.Println("It works!")
  getData()

  /*
  syntax for creating endpoints:
  	router.HandleFunc("/<your-url>", <function-name>).Methods("<method>")
  syntax for creating those functions:
	func <your-function-name>(w http.ResponseWriter, r *http.Request) {
	
	}
  For more info: https://github.com/gorilla/mux
  It's a The most notable and highly used third-party router package.
  As it stands currently has 12,481 stars on Github.
  */
  router := mux.NewRouter()
  router.HandleFunc("/users", getUsers).Methods("GET")
  router.HandleFunc("/users", createUser).Methods("POST")
  router.HandleFunc("/users/{id}", getUser).Methods("GET")
  router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
  router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
  http.ListenAndServe(":8080", router)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// The call to json.NewEncoder(w).Encode(users) does the job of
	// encoding our users array into a JSON string and then writing as
	// part of our response.
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop over all of our users
    // if the user.ID equals the ID we pass in
    // return the user encoded as JSON
	for _, item := range users {
	  // to extract params we do params["<param-name>"]
	  if item.ID == params["id"] {
		  // return the user encoded as JSON
		json.NewEncoder(w).Encode(item)
		break
	  }
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
	  if item.ID == params["id"] {
		users = append(users[:index], users[index+1:]...)
		var user User
		_ = json.NewDecoder(r.Body).Decode(user)
		user.ID = params["id"]
		users = append(users, user)
		json.NewEncoder(w).Encode(&user)
		return
	  }
	}
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
	  if item.ID == params["id"] {
		users = append(users[:index], users[index+1:]...)
		break
	  }
	}
	json.NewEncoder(w).Encode(users)
  }