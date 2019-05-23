package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var address string
func main(){
	if len(os.Args) <= 1{
		log.Println("Please input server address!")
		os.Exit(0)
	}
	address = os.Args[1]
	mux := http.DefaultServeMux
	mux.HandleFunc("/upload",upload)
	mux.HandleFunc("/",showPage)
	server := &http.Server{
		Addr : address,
		Handler : mux,
	}
	log.Println("Start file server at :"+address)
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error())
	}
}

func showPage(w http.ResponseWriter,req *http.Request){
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="file">
    <input type="submit">
</form>
</body>
</html>`
	w.Write([]byte(html))
}
func upload(w http.ResponseWriter,req *http.Request){
	file, fileHeader, err := req.FormFile("file")
	if err == nil {
		log.Println("Receive file : "+fileHeader.Filename)
		data,err := ioutil.ReadAll(file)
		if err == nil {
			ioutil.WriteFile(fileHeader.Filename,data,os.FileMode(0777))
		}else{
			log.Println(err.Error())
		}
	}else{
		log.Println(err.Error())
	}
	fmt.Fprintln(w,"success")
}
