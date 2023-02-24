// package main

// import (
// 	"context"
// 	"encoding/csv"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Details struct {
// 	ID    primitive.ObjectID `json:"_id" bson:"_id" `
// 	Bid   string             `json:"bid" bson:"bid"`
// 	Name  string             `json:"name" bson:"name"`
// 	Marks string             `json:"marks" bson:"marks"`
// }

// var Totalrec [][]string
// var db *mongo.Client

// func main() {
// 	var err error

// 	// dont put this in real code. Nobody will speak to you
// 	// credentials := options.Credential{
// 	// 	Username: "admin",
// 	// 	Password: "admin",
// 	// }

// 	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
// 	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.8cpxunh.mongodb.net/?retryWrites=true&w=majority")

// 	db, err = mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Disconnect(context.TODO())

// 	r := mux.NewRouter()
// 	//r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods(http.MethodGet)
// 	r.HandleFunc("/data", PostDetails).Methods("POST")

// 	srv := &http.Server{
// 		Handler:      r,
// 		Addr:         "localhost:8080",
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}

// 	log.Fatal(srv.ListenAndServe())
// }

// func PostDetails(w http.ResponseWriter, r *http.Request) {
// 	//var record Details
// 	//wg := sync.WaitGroup{}
// 	collection := db.Database("bookstore").Collection("csvvalues")
// 	csvfile, err := os.Open("input.csv")
// 	if err != nil {
// 		log.Fatalln("Couldn't open the csv file", err)
// 	}
// 	totalrec, _ := csv.NewReader(csvfile).ReadAll()

// 	numlines := len(totalrec)

// 	for i := 0; i < numlines; i++ {
// 		row := totalrec[i]
// 		record := Details{}
// 		record.Bid = row[0]
// 		record.Name = row[1]
// 		record.Marks = row[2]
// 		record.ID = primitive.NewObjectID()
// 		result, err := collection.InsertOne(context.TODO(), record)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(err.Error()))
// 		} else {
// 			w.Header().Set("Content-Type", "application/json")
// 			response, _ := json.Marshal(result)
// 			w.Write(response)
// 		}

// 	}

// 	// postBody, _ := ioutil.ReadAll(r.Body)
// 	// json.Unmarshal(postBody, &record)

// 	// result, err := db.collection.InsertOne(context.TODO(), record)
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusInternalServerError)
// 	// 	w.Write([]byte(err.Error()))
// 	// } else {
// 	// 	w.Header().Set("Content-Type", "application/json")
// 	// 	response, _ := json.Marshal(result)
// 	// 	w.Write(response)
// 	// }
// }

// /*
// func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	var record Details
// 	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
// 	filter := bson.M{"_id": objectID}

// 	err := db.collection.FindOne(context.TODO(), filter).Decode(&record)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 	} else {
// 		w.Header().Set("Content-Type", "application/json")
// 		response, _ := json.Marshal(record)
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(response)
// 	}
// }

// func (db *DB) PostDetails(w http.ResponseWriter, r *http.Request) {
// 	var record Details

// 	postBody, _ := ioutil.ReadAll(r.Body)
// 	json.Unmarshal(postBody, &record)

// 	result, err := db.collection.InsertOne(context.TODO(), record)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 	} else {
// 		w.Header().Set("Content-Type", "application/json")
// 		response, _ := json.Marshal(result)
// 		w.Write(response)
// 	}
// }

// */
