package main

import (
	"net/smtp"
	"log"
	"fmt"
)

const SMTPServ = "smtp.xxx.xxx:25"

func init() {
	log.SetPrefix("XXX: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	fmt.Println("init ok")
}

// TODO: ./panic-and-recover.og
// recover的defer声明必须在panic之前
func main() {
	defer func() {
		what := recover()
		fmt.Println("recover from", what)
	}()

	client, err := smtp.Dial(SMTPServ)
	if err != nil {
		// log.Fatalln(err)
		log.Panicln(err)
	}
	client.Data()
	client.Close()

	//defer func() {
	//	what := recover()
	//	fmt.Println("recover from", what)
	//}()
}
