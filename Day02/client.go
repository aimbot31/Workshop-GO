package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func readMsg(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		message := string(buffer[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			print("lollll")
			panic(err)
		}
		if n > 0 {
			fmt.Println(message)
		}
	}
}

func exitMsg(msg string, status int) {
	fmt.Println(msg)
	os.Exit(status)
}

func main() {
	conn, err := net.Dial("tcp", "10.15.192.174:8754")
	if err != nil {
		exitMsg("Erreur lors de la connection..", 84)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrer un pseudo: ")
	pseudo, _ := reader.ReadString('\n')
	conn.Write([]byte(pseudo))
	//status, err := bufio.NewReader(conn).ReadString('\n')
	go readMsg(conn)
	fmt.Println("Message: ")
	for {
		text, _ := reader.ReadString('\n')
		b := []byte(text)
		_, err := conn.Write(b)
		if err != nil {
			exitMsg("Impossible d'envoyer le message au serveur, peut-être est-il éteint..", 84)
		}
	}
	conn.Close()

}
