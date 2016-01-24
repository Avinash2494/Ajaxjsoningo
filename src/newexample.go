package main
import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"os"
	_"mysql"
	"database/sql"
	"net/http"
	"log"
)
type Scu struct {
	Scu_id    int64    `json:"scuid"`
	Sgu_id    int64    	`json:"sguid"`
	Location_name string `json:"name"`
	Location_lat float32    `json:"lat"`
	Location_lng float32    `json:"long"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/fetchtable",fetchtable)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func fetchtable(w http.ResponseWriter, r *http.Request)  {
	//scu := Scu{1, 1, "bangalore", 345, 342}
	var scudata[15] Scu
	var scudatafromdatabase[15] Scu
	db,err:=sql.Open("mysql","root:admin123@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		fmt.Println("Could not connect error occured")
	}
	stmt,err:=db.Prepare("select * from scu")
	if err!=nil{
		fmt.Println("could not fetch data")
	}
	result,err:=stmt.Query()
	var i int
	for result.Next(){
		err:= result.Scan(&scudatafromdatabase[i].Scu_id,&scudatafromdatabase[i].Sgu_id,&scudatafromdatabase[i].Location_name,&scudatafromdatabase[i].Location_lat,&scudatafromdatabase[i].Location_lng)
		i++
		if err!=nil{
			fmt.Println("error occured")
		}
	}
	bytes,err := json.Marshal(scudatafromdatabase)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(bytes)
	json.Unmarshal(bytes, &scudata)
	fmt.Println(scudata[1])
	//return string(bytes)
	w.Header().Set("Content-type","application/json")
	w.Write(bytes)
}