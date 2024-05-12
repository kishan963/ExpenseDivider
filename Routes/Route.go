package main // <- Add this line

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/kishan963/Splitwise/Controller"
	d "github.com/kishan963/Splitwise/Database"
	j "github.com/kishan963/Splitwise/JwtToken"
	m "github.com/kishan963/Splitwise/Models"
	"github.com/rs/cors"
	// Add other necessary import statements
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/getAllUser", GetAllUser).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/createGroup", CreateGroupHandler).Methods("POST")
	r.HandleFunc("/addExpense", AddExpenseHandler).Methods("POST")
	r.HandleFunc("/getGroup", GetGroupHandler).Methods("POST")
	r.HandleFunc("/getUserBalance", BalanceHandler).Methods("POST")
	r.HandleFunc("/getUserGroups", GetuserGroupsHandler).Methods("POST")
	r.HandleFunc("/deleteGroup", DeleteGroupHandler).Methods("POST")
	r.HandleFunc("/deleteExpense", DeleteExpenseHandler).Methods("POST")
	err := d.DbSetup()
	if err != nil {
		fmt.Println("Error in db creation ", err)
		log.Fatal(err)
	}

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// })

	handler := cors.AllowAll().Handler(r)
	// Add other route handlers as needed

	// Start the HTTP server
	http.ListenAndServe(":8080", handler)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var data m.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	c.Register(data)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data m.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	val, err := c.Login(data)
	if err != nil {
		// If login failed
		http.Error(w, "Login failed. "+err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Println(val)
	// If login successful
	jsonVal, err := json.Marshal(val)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	if _, err := j.IsAuthorized(w, r); err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	users := c.GetAllUser()
	// Convert users slice to JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := j.IsAuthorized(w, r); err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var data m.Group
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	c.CreateGroup(data)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))
}

func AddExpenseHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := j.IsAuthorized(w, r); err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var data m.Expense
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	c.AddExpense(data)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Expemse added successful"))
}

func GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := j.IsAuthorized(w, r); err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var data m.Group
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	group, _ := c.GetGroup(data)
	// Convert users slice to JSON
	jsonData, err := json.Marshal(group)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	User_id, err := j.IsAuthorized(w, r)
	if err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("UserIDD ", User_id)
	group := c.GetUserBalance(User_id)
	// Convert users slice to JSON
	jsonData, err := json.Marshal(group)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func GetuserGroupsHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := j.IsAuthorized(w, r); err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var data m.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Inside the function")
	group, _ := c.GetUserGroups(data)
	// Convert users slice to JSON
	jsonData, err := json.Marshal(group)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := j.IsAuthorized(w, r); err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var data m.Group
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	c.DeleteGroup(data)

	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func DeleteExpenseHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := j.IsAuthorized(w, r); err != nil {
		http.Error(w, "Authentication failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var data m.Expense
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	c.DeleteExpense(data)

	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
