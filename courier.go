package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
)

// takes a host, and a raw http request as a string
func socketreq(host string, message string) (string, string) {

	// use tls cert
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", host), &config)
	if err != nil {
		log.Fatalf(fmt.Sprintf("client: dial: %s:443", host), err)
	}
	defer conn.Close()

	// connection complete, now write the message

	_, err = io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}
	//log.Printf("client: wrote %q (%d bytes)", message, n)

	// read response

	/* this would be preffered method but isnt working with tls
	res, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println("Error reading response", err)
	}
	fmt.Println(res)
	*/

	headers := make([]byte, 1024)
	body := make([]byte, 512)

	n, err := conn.Read(headers)
	if err != nil {
		//log.Println("Error reading response", err)
	}
	//log.Printf("client: read %q (%d bytes)", string(headers[:n]), n)
	retheader := string(headers[:n])
	n, err = conn.Read(body)
	if err != nil {
		//log.Println("Error reading response", err)
	}
	//log.Printf("client: read %q (%d bytes)", string(body[:n]), n)
	retbody := string(body[:n])

	//oracleCLTE(host, resHeaders, resBodies)
	return retheader, retbody
}
