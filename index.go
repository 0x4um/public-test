package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
)

const charset = "abcdefghijkmnopqrstuvwxyz1234567890"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)

	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()

}

func findLen(input string) int {
	returnCounter := 0
	for i := 0; i < len(input); i++ {
		returnCounter++
	}
	return returnCounter
}

func main() {
	fmt.Println("main func")
	listener, err := net.Listen("tcp", ":1200")
	if err != nil {
		fmt.Println("error listen")

	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("error accepting listen")

		}
		go func(conn net.Conn) {
			// conn.Write([]byte("test"))
			conn.Write([]byte("hello back"))
			fmt.Println("rec con" + conn.RemoteAddr().String())
			buffer := make([]byte, 1024)
			len, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("error reading")
				return
			}
			fmt.Println(string(buffer[:len]))
			bufferFormat := string(buffer[:len])
			bufferFormatLen := bufferFormat[:findLen(bufferFormat)-1]
			// fmt.Println("initconn here")
			// fmt.Println(bufferFormat[:findLen(bufferFormat)-1])
			// fmt.Println(string(bufferFormat))
			// fmt.Println(findLen(bufferFormat))

			if bufferFormatLen == "initconn" {
				fmt.Println("init conn called")
				//generate peer id and return it to client
				//after peer id is generated calculate the distance
				// sb := strings.Builder{}
				// // const charset = "abcdefghijkmnopqrstuvwxyz"
				// sb.Grow(10)
				// for i := 0; i < 10; i++ {
				// 	sb.WriteByte(charset[rand.Intn(len(charset))])
				// }
				// fmt.Println(sb.String())
			}
			conn.Write([]byte("|serverHeaderReturn|" + randomString(24) + "\n"))
			//enter randomstring into database with client header info
			//start clock to calculate ping test spead
			conn.Close()
		}(conn)
	}
}
