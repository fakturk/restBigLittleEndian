package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type BigLittleBinary struct {
	bigEndian    string `json:"bigendian"`
	littleEndian string `json:"littleendian"`
}

// Display a single data
func GetBinary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	numberString := params["number"]
	number, errInt := strconv.ParseInt(numberString, 10, 64)
	if errInt != nil {
		// handle error
		fmt.Println(errInt)
		os.Exit(2)
	}
	bufLittle := new(bytes.Buffer)
	bufBig := new(bytes.Buffer)

	err := binary.Write(bufLittle, binary.LittleEndian, number)
	if err != nil {
		fmt.Println("Little Endian:", err)
	}

	err2 := binary.Write(bufBig, binary.BigEndian, number)
	if err != nil {
		fmt.Println("Little Endian:", err2)
	}
	b := fmt.Sprintf("%x", bufBig)
	l := fmt.Sprintf("%x", bufLittle)
	// bigString := bufBig.String()
	// littleString := bufLittle.String()
	fmt.Println(number, b, l)
	// fmt.Printf("%d, %x, %x \n", number, bigString, littleString)
	// fmt.Printf("%d, %x, %x \n", number, bufBig, bufLittle)
	bigLittleBinary := BigLittleBinary{bigEndian: b, littleEndian: l}
	fmt.Printf("%+v\n", bigLittleBinary)
	// fmt.Fprintf(w, "%+v\n", bigLittleBinary)
	fmt.Fprintf(w, "{\"bigEndian\" : \"%s\", \"littleEndian\" : \"%s\"}\n", b, l)

	// json.NewEncoder(w).Encode(bigLittleBinary)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/biglittleendian/{number}", GetBinary).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
