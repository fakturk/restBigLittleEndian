package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "bytes"
)

type BigLittleBinary struct{
   bytes.Buffer bigEndian     `json:"bigendian,omitempty"`
   bytes.Buffer littleEndian  `json:"littleendian,omitempty"`

}
// Display a single data
func GetBinary(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    number := params["number"]


      bufLittle := new(bytes.Buffer)
      bufBig := new(bytes.Buffer)

    err := binary.Write(bufLittle, binary.LittleEndian, number)
      if err != nil {
            fmt.Println("Little Endian:", err)
}

err := binary.Write(bufBig, binary.BigEndian, number)
  if err != nil {
        fmt.Println("Little Endian:", err)
}
// bigLittleBinary := 

    // for _, item := range people {
    //     if item.ID == params["id"] {
    //         json.NewEncoder(w).Encode(item)
    //         return
    //     }
    // }
    json.NewEncoder(w).Encode(&BigLittleBinary{})
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/biglittleendian/{number}", GetBinary).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}
