package API

import (
	
	"net/http"
	"github.com/satori/go.uuid"
)

//func GetKey
func GetKey(w http.ResponseWriter, r *http.Request) {
	keyvalue := uuid.Must(uuid.NewV1())
	w.Write([]byte(keyvalue.String()))
/*
	keys, ok := r.URL.Query()["key"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	*/
	
}