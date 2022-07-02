package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main()  {
	// mail Information
	from:= Email
	password:= MailPassword

	// Smtp Information 
	smtpHost:= "smtp.gmail.com"
	smtpPort:= "587"

	// Receiver
	to:= []string{"raihanmahmudi35@gmail.com", "mtmartadd@gmail.com"}




	// Message

	subject:= "Subject: Welcome from Learn With Raihan\n"
	body:= "Welcome to Learn With Raihan 2\n"
	message:= []byte(subject + body)


	// Authentication 
	auth:= smtp.PlainAuth("", from,password, smtpHost)


	// Send Mail

	// smtp.gmail.com:587
	err:= smtp.SendMail(smtpHost+":"+smtpPort,auth,from, to, message)
	if err!=nil{
		log.Fatal(err)
		return
	}
	fmt.Println("Mail Send")



}