package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func logSettings(f string) {
	logfile, _ := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetPrefix("Prefix")
	log.SetFlags(log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}
func main() {
	logSettings("test.log")
	_, err := os.Open("test.log")
	if err != nil {
		log.Fatalln("fatal")
	}
	log.Println("test")

	//    log.Fatalln("error")
	fmt.Println("ok")

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatalln("open err")
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("read err")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("cd err")
	}

}
