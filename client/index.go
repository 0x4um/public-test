package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const database string = "./data/index.db"

func initConn(conn net.Conn) bool {
	fmt.Println("attempting to init connection to server")
	conn.Write([]byte("initconn\n"))
	//start timer here

	return true
}

func calculateMili(conn net.Conn) int64 {
	fmt.Println("sending packet to server to calc time")
	buffer := make([]byte, 1024)
	startTime := time.Now()

	_, err := conn.Write(buffer)
	if err != nil {
		fmt.Println("error packet send", err)

	}
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error reading packet", err)
	}

	duration := time.Now().Sub(startTime)
	durationMili := duration.Milliseconds()

	return durationMili
}

func main() {
	// fmt.Println("main func")

	conn, err := net.Dial("tcp", "localhost:1200")
	if err != nil {
		fmt.Println("error conn")
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println(calculateMili(conn))
	fmt.Println("after")
	//check to see if client has ability to init connection through the database
	if initConn(conn) {
		fmt.Println("connection to server success")
		//start timer here

	} else {
		fmt.Println("server failed")
	}
	recived := make([]byte, 1024)
	len, err := conn.Read(recived)
	if err != nil {
		fmt.Println("error reading")
		os.Exit(1)
	}

	recivedFormat := string(recived[:len])
	recivedFormatLen := recivedFormat[:findLen(recivedFormat)-1]
	// fmt.Println("follow")
	returnArr, returnArrCount := readStringParse(recivedFormatLen)

	var peerID = returnArr[1]
	fmt.Println(peerID)
	fmt.Println("PEERID")
	for i := 0; i < returnArrCount; i++ {
		fmt.Println(returnArr[i])

	}

}

func findLen(input string) int {
	returnCounter := 0
	for i := 0; i < len(input); i++ {
		returnCounter++
	}
	return returnCounter
}

func readStringParse(readString string) ([5]string, int) {
	var returnArr [5]string
	var parseArr [7]int
	var parseArrCounter = 0
	for i := 0; i < len(readString); i++ {
		// fmt.Println(string(readString[i]))
		if string(readString[i]) == "|" {
			// fmt.Println("found //", i)
			parseArr[parseArrCounter] = i
			parseArrCounter++
		}
	}
	// fmt.Println(readString)
	// fmt.Println("reader string")
	// fmt.Println(parseArr[0], parseArr[1])
	for k := 0; k < parseArrCounter; k++ {
		// fmt.Println(parseArr[k])
		// fmt.Println(k)
		if k == parseArrCounter-1 {
			// fmt.Println("end")
			// fmt.Println(string(readString[parseArr[k]:][1:]))
			returnArr[k] = string(readString[parseArr[k]:][1:])
		} else {
			// fmt.Println("above")
			// fmt.Println(string(readString[parseArr[k]:parseArr[k+1]][1:]))
			returnArr[k] = string(readString[parseArr[k]:parseArr[k+1]][1:])
			//old version
			// fmt.Println(readString[parseArr[k]:parseArr[k+1]])
		}
	}
	// fmt.Println(readString[parseArr[0]+1 : parseArr[1]])
	// fmt.Println(readString[parseArr[1]+1:])
	// fmt.Println(parseArrCounter)
	return returnArr, parseArrCounter
}
