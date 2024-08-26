package main

import (
	"fmt"
	"math/rand"
	"net"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	go TCP()
	resolvU, _ := net.ResolveUDPAddr("udp", ":1234")
	conectionU, _ := net.ListenUDP("udp", resolvU)

	defer conectionU.Close()
	buffer := make([]byte, 1024)

	for {
		_, conexionCon, _ := conectionU.ReadFromUDP(buffer)
		fmt.Print("Usuario conectado")
		preguntas := rand.Intn(5) + 3

		datos := fmt.Sprintf("%d,%s,%s", preguntas, "127.0.0.1", "4321")
		conectionU.WriteToUDP([]byte(datos), conexionCon)

	}
}
func TCP() {

	resolvT, _ := net.ResolveTCPAddr("tcp", ":4321")
	conectionT, _ := net.ListenTCP("tcp", resolvT)
	defer conectionT.Close()

	conexion, _ := conectionT.AcceptTCP()

	defer conexion.Close()

	buffer := make([]byte, 1024)
	n, _ := conexion.Read(buffer)
	preguntas := string(buffer[:n]) //leo cantidad de preguntas
	fmt.Print("TPC debe dar " + preguntas + " preguntas")

}
