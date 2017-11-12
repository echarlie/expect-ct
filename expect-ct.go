package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

//"encoding/json"

//type scts_data struct {
//	version        int
//	status         string
//	source         string
//	serialized_sct string
//}
//
//type decoded_data struct {
//	dateTime    string `json:"date-time"`
//	hostname    string
//	port        int
//	effecExp    string   `json:"effective-expiration-date"`
//	servedChain []string `json:"served-certificate-chain"`
//	validChain  []string `json:"validated-certificate-chain"`
//	scts        scts_data
//}

func DumpReport(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		//		decoder := json.NewDecoder(r.Body)
		//		var d decoded_data
		//		err2 := decoder.Decode(&d)
		//		if err2 != nil {
		//			panic(err2)
		//		}
		//log.Println(d)
		log.Println(string(body))
	} else {
		http.Error(w, "no services provided here", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/report", DumpReport)
	http.ListenAndServe("[::1]:8088", nil)
}
