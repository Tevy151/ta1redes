package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	resolv_udp, err := net.ResolveUDPAddr("udp", "127.0.0.1:1234")

	if err != nil {
		fmt.Println("Error con la dirección UDP:", err)
		return
	}

	conex_udp, err := net.DialUDP("udp", nil, resolv_udp)

	if err != nil {
		fmt.Println("Error al conectar UDP:", err)
		return
	}

	defer conex_udp.Close()

	fmt.Println("Conectando al servidor UDP...")
	_, err = conex_udp.Write([]byte("Iniciar"))

	if err != nil {
		fmt.Println("Error con UDP:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := conex_udp.ReadFromUDP(buffer)

	if err != nil {
		fmt.Println("Error conectando a UDP:", err)
		return
	}

	datos := string(buffer[:n])
	datosSplit := strings.Split(datos, ",")

	cantidad_preguntas := datosSplit[0]
	ip := datosSplit[1]
	puerto := datosSplit[2]

	resolvT, err := net.ResolveTCPAddr("tcp", ip+":"+puerto)

	if err != nil {
		fmt.Println("Error con TCP:", err)
		return
	}

	conex_tcp, err := net.DialTCP("tcp", nil, resolvT)

	if err != nil {
		fmt.Println("Error al conectar TCP:", err)
		return
	}

	defer conex_tcp.Close()

	fmt.Println("Conectando mediante TCP...")
	_, err = conex_tcp.Write([]byte(cantidad_preguntas))

	if err != nil {
		fmt.Println("Error enviando cantidad de preguntas: ", err)
		return
	}

	cantidadPreguntas := 0
	fmt.Sscanf(cantidad_preguntas, "%d", &cantidadPreguntas)

	for i := 0; i < cantidadPreguntas; i++ {

		n, err = conex_tcp.Read(buffer)

		if err != nil {
			fmt.Println("Error leyendo pregunta: ", err)
			return
		}

		pregunta := string(buffer[:n])
		fmt.Println("Pregunta:", pregunta)

		var respuesta string
		fmt.Scanln(&respuesta)
		_, err = conex_tcp.Write([]byte(respuesta))

		if err != nil {
			fmt.Println("Error enviando respuesta: ", err)
			return
		}
	}

	n, err = conex_tcp.Read(buffer)

	if err != nil {
		fmt.Println("Error leyendo puntaje: ", err)
		return
	}

	puntaje := string(buffer[:n])
	fmt.Println("Tu Puntaje Final:", puntaje)

	_, err = conex_tcp.Write([]byte("Finalizado"))
	if err != nil {
		fmt.Println("Error enviando mensaje de finalización:", err)
		return
	}

	fmt.Println("Mensaje de finalización enviado al servidor.")
}
