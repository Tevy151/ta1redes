package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	resolvU, _ := net.ResolveUDPAddr("udp", "127.0.0.1:1234")
	conectionU, _ := net.DialUDP("udp", nil, resolvU)
	defer conectionU.Close()

	conectionU.Write([]byte("Iniciar"))
	buffer := make([]byte, 1024)

	n, _, _ := conectionU.ReadFromUDP(buffer)
	datos := string(buffer[:n])
	datosSplit := strings.Split(datos, ",")
	preguntas := datosSplit[0]
	ip := datosSplit[1]
	puerto := datosSplit[2]

	resolvT, _ := net.ResolveTCPAddr("tcp", ip+":"+puerto)
	conectionT, err := net.DialTCP("tcp", nil, resolvT)
	if err != nil {
		fmt.Println("Error al conectar TCP:", err)
		return
	}

	defer conectionT.Close()

	conectionT.Write([]byte(preguntas)) //Conexion TCP
	var respuesta string
	for i := 0; i < 2; i++ {
		n, _ = conectionT.Read(buffer)
		pregunta := string(buffer[:n])
		fmt.Println("Pregunta:", pregunta)
		fmt.Scanln(&respuesta)
		conectionT.Write([]byte(respuesta))
		n, _ = conectionT.Read(buffer)
		puntaje := string(buffer[:n])
		fmt.Println("Tu Puntaje:", puntaje)

	}

}
