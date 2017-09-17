package godown

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)


//LogPrintln Outputs to the console
func LogPrintln(target string, err error) {
	log.Printf("%v error: %v\n", target, err)
}

//Gmail sends an alert email using Gmail
func Gmail(target string, reason error, from, to, pass string) {
	body := fmt.Sprintf(
		"To:%v\r\n"+
			"Subject: Godown alert! [%v]\r\n\r\n"+
			"Alert! \n\n A service didn't respond. "+
			" \n target: %v \n error: %v\r\n", to, target, target, reason)

	host := "smtp.gmail.com:465"

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Connect to the remote SMTP server.
	conn, err := tls.Dial("tcp", "smtp.gmail.com:465", tlsconfig) //
	if err != nil {
		log.Println(err)
		return
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	a := smtp.PlainAuth(from, from, pass, host)
	if err := c.Auth(a); err != nil {
		log.Println(err)
		return
	}

	// Set the sender and recipient first
	if err := c.Mail(from); err != nil {
		log.Println(err)
		return
	}
	if err := c.Rcpt(to); err != nil {
		log.Println(err)
		return
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Println(err)
		return
	}
	_, err = fmt.Fprintf(wc, body)
	if err != nil {
		log.Println(err)
	}
	err = wc.Close()
	if err != nil {
		log.Println(err)
		return
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Println(err)
		return
	}
}
