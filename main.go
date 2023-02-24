package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// declaring a structure like CSV file
type Details struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id" `
	Bid   string             `json:"bid" bson:"bid"`
	Name  string             `json:"name" bson:"name"`
	Marks string             `json:"marks" bson:"marks"`
}

//var Totalrec [][]string

// variable for MongoDb Client
var db *mongo.Client

func main() {
	var err error

	// Use the below code if using the local host
	/*
		credentials := options.Credential{
			Username: "admin",
			Password: "admin",
		}
	*/

	//// Use the below code if using the local host
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.8cpxunh.mongodb.net/?retryWrites=true&w=majority")

	// Database connection
	db, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	// Disconnecting database after all the process is done
	defer db.Disconnect(context.TODO())
	// Declaring a Router
	r := mux.NewRouter()
	r.HandleFunc("/data/{id:[a-zA-Z0-9]*}", GetDetails).Methods("GET")
	r.HandleFunc("/data", PostDetails).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// Handler functions to insert data into database
func PostDetails(w http.ResponseWriter, r *http.Request) {
	//var record Details

	// using wait groups for concurrency
	wg := sync.WaitGroup{}
	// specifying the databse and the collection to be used
	collection := db.Database("bookstore").Collection("csvvalues")

	// read csv file
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// reading all the data as [][]2D Matrix
	totalrec, _ := csv.NewReader(csvfile).ReadAll()

	numlines := len(totalrec)
	//===================

	// creating channels a many as records

	channel := make(chan []string, numlines)
	// inserting data into channel
	for i := 0; i < numlines; i++ {
		channel <- totalrec[i]
	}
	// loop to retrive data from channel and insert into database
	for i := 0; i < numlines; i++ {
		wg.Add(1)

		// Anonymous function to perform insertion
		go func(channel chan []string, wg *sync.WaitGroup) {
			row := <-channel

			record := Details{}
			record.Bid = row[0]
			record.Name = row[1]
			record.Marks = row[2]
			record.ID = primitive.NewObjectID()

			// insering records into database
			result, err := collection.InsertOne(context.TODO(), record)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				w.Header().Set("Content-Type", "application/json")
				response, _ := json.Marshal(result)
				w.Write(response)
			}
			//alldetails = append(alldetails, record)

			wg.Done()
		}(channel, &wg)
	}
	wg.Wait()

}

// Handler functions to get data using the ID
func GetDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var record Details
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	collection := db.Database("bookstore").Collection("csvvalues")
	filter := bson.M{"_id": objectID}
	// Finding if the record in DB
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(record)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
