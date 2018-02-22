package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func sendMessage(conn *[]net.Conn, msg string) {
	date := time.Now().Format("15:04")
	msg = date + " " + msg
	for i := 0; i < len((*conn)); i++ {
		(*conn)[i].Write([]byte(msg))
	}
}

func handleConnection(conn *[]net.Conn, i int) {
	var pseudo string
	for {
		buffer := make([]byte, 1024)
		n, err := (*conn)[i].Read(buffer)
		message := string(buffer[:n-1])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err, "lol")
			return
		}
		if message == "/quit" {
			fmt.Println(pseudo, "s'est déconnecté")
			return
		}
		if n > 0 {
			if pseudo == "" {
				pseudo = message
				sendMessage(conn, pseudo+" s'est connecté")
				fmt.Println(pseudo + " s'est connecté")
			} else {
				fmt.Println(pseudo, ":", message)
				sendMessage(conn, pseudo+" : "+message)
			}
		}
	}
}

func exitMsg(msg string, status int) {
	fmt.Println(msg)
	os.Exit(status)
}

func main() {
	ln, err := net.Listen("tcp4", ":8754")
	i := 0
	var tab_conn []net.Conn
	if err != nil {
		exitMsg("Impossible d'écouter sur le port choisis..", 84)
	}
	fmt.Println("Le serveur écoute sur le port 8754")
	for {
		conn, err := ln.Accept()
		tab_conn = append(tab_conn, conn)
		if err != nil {
			println("Erreur lors de la connection d'un client.")
		} else {
			go handleConnection(&tab_conn, i)
		}
		i++
	}
}
