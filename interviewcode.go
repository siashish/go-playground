package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger *log.Logger
	ErrorLogger *log.Logger
)

func init(){
	file,err := os.OpenFile("logs.txt",os.0_APPEND|os.0_CREATED|os.0_WRONLY,0666)
	if err != nil{
		log.Fatal(err)
	}
	InfoLogger = log.New(file,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file,"WARNING: ",log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file,"ERROR: ",log.Ldate|log.Ltime|log.Lshortfile)
}
func Logger(data interface{}, type string){
	if type == "error"{
		ErrorLogger.Println(data)
	} else if type == "info"{
		InfoLogger.Println(data)
	}else if type=="warning"{
		WarningLogger.Println(data)
	}else{
		log.Println("unknown type")
	}
}
func main() {
	Logger("Starting the app...","info")
	Logger("warning happen...","warning")
	Logger("Error ...","error")
}
