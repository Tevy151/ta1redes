package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	resolvU, err := net.ResolveUDPAddr("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error resolviendo la dirección UDP:", err)
		return
	}
	conectionU, err := net.DialUDP("udp", nil, resolvU)
	if err != nil {
		fmt.Println("Error al conectar UDP:", err)
		return
	}
	defer conectionU.Close()

	fmt.Println("Conectando al servidor UDP...")
	_, err = conectionU.Write([]byte("Iniciar"))
	if err != nil {
		fmt.Println("Error enviando inicio al servidor UDP:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := conectionU.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error leyendo datos desde UDP:", err)
		return
	}

	datos := string(buffer[:n])
	datosSplit := strings.Split(datos, ",")
	cantidadPreguntasStr := datosSplit[0]
	ip := datosSplit[1]
	puerto := datosSplit[2]

	resolvT, err := net.ResolveTCPAddr("tcp", ip+":"+puerto)
	if err != nil {
		fmt.Println("Error resolviendo la dirección TCP:", err)
		return
	}
	conectionT, err := net.DialTCP("tcp", nil, resolvT)
	if err != nil {
		fmt.Println("Error al conectar TCP:", err)
		return
	}
	defer conectionT.Close()

	fmt.Println("Conectando al servidor TCP...")
	_, err = conectionT.Write([]byte(cantidadPreguntasStr))
	if err != nil {
		fmt.Println("Error enviando cantidad de preguntas:", err)
		return
	}

	cantidadPreguntas := 0
	fmt.Sscanf(cantidadPreguntasStr, "%d", &cantidadPreguntas)

	for i := 0; i < cantidadPreguntas; i++ {
		n, err = conectionT.Read(buffer)
		if err != nil {
			fmt.Println("Error leyendo pregunta:", err)
			return
		}
		pregunta := string(buffer[:n])
		fmt.Println("Pregunta:", pregunta)

		var respuesta string
		fmt.Scanln(&respuesta)
		_, err = conectionT.Write([]byte(respuesta))
		if err != nil {
			fmt.Println("Error enviando respuesta:", err)
			return
		}
		n, err = conectionT.Read(buffer)
		if err != nil {
			fmt.Println("Error leyendo puntaje:", err)
			return
		}
		puntaje := string(buffer[:n])
		fmt.Println("Tu Puntaje:", puntaje)

	}

	n, err = conectionT.Read(buffer)
	if err != nil {
		fmt.Println("Error leyendo puntaje:", err)
		return
	}
	puntaje := string(buffer[:n])
	fmt.Println("Tu Puntaje Final:", puntaje)
}
