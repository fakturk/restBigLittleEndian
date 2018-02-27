package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)
// Display a single data
func GetBinary(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    // for _, item := range people {
    //     if item.ID == params["id"] {
    //         json.NewEncoder(w).Encode(item)
    //         return
    //     }
    // }
    json.NewEncoder(w).Encode(&Person{})
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/biglittleendian/{number}", GetBinary).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}
