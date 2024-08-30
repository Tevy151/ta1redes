package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {

	go TCP()
	resolv_udp, err := net.ResolveUDPAddr("udp", ":1234")

	if err != nil {
		fmt.Println("Error resolviendo la dirección UDP:", err)
		return
	}

	conexion_udp, err := net.ListenUDP("udp", resolv_udp)

	if err != nil {
		fmt.Println("Error al escuchar UDP:", err)
		return
	}

	defer conexion_udp.Close()

	fmt.Println("Servidor UDP escuchando en", resolv_udp.String())
	buffer := make([]byte, 1024)

	for {

		_, conex, err := conexion_udp.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error leyendo desde UDP:", err)
			continue

		}
		fmt.Println("Usuario conectado desde", conex.String())

		preguntas := rand.Intn(5) + 3 // Número aleatorio entre 3 y 7
		datos := fmt.Sprintf("%d,%s,%s", preguntas, "127.0.0.1", "4321")
		_, err = conexion_udp.WriteToUDP([]byte(datos), conex)

		if err != nil {
			fmt.Println("Error enviando datos al cliente:", err)
		}

	}
}

func TCP() {

	rand.Seed(time.Now().UnixNano())

	preguntas := []string{
		"Cuanto es 1+1?",
		"Cuantas regiones tiene Chile?",
		"Cuantos paises hay en América?",
		"Las empanadas llevan pasas?",
		"El pisco es Chileno?",
		"Computación Cientifica es el ramo más fácil de informática?",
		"Pepsi fue una potencia naval?",
		"La final de champions del 2019 la ganó el Liverpool?",
		"La usm tiene prestigio?",
		"Tangananica o tanganana?",
	}

	respuestas := []string{
		"2",
		"16",
		"35",
		"no",
		"si",
		"no",
		"si",
		"si",
		"si",
		"tangananica",
	}

	resolv_tcp, err := net.ResolveTCPAddr("tcp", ":4321")

	if err != nil {
		fmt.Println("Error resolviendo la dirección TCP:", err)
		return
	}

	conexion_tcp, err := net.ListenTCP("tcp", resolv_tcp)

	if err != nil {
		fmt.Println("Error al escuchar TCP:", err)
		return

	}

	defer conexion_tcp.Close()
	fmt.Println("Servidor TCP escuchando en", resolv_tcp.String())

	conexion, err := conexion_tcp.AcceptTCP()

	if err != nil {
		fmt.Println("Error aceptando conexión TCP:", err)
		return
	}

	defer conexion.Close()

	fmt.Println("Conexión TCP aceptada de", conexion.RemoteAddr().String())

	buffer := make([]byte, 1024)
	n, err := conexion.Read(buffer)

	if err != nil {
		fmt.Println("Error leyendo cantidad de preguntas:", err)
		return
	}

	cant_preguntas_str := string(buffer[:n])
	cant_preguntas_int, err := strconv.Atoi(cant_preguntas_str)

	if err != nil {
		fmt.Println("Error convirtiendo preguntas a número:", err)
		return
	}

	fmt.Println("TCP debe dar", cant_preguntas_str, "preguntas")

	i_selec := rand.Perm(len(preguntas))[:cant_preguntas_int]
	preguntas_selec := make([]string, 0, cant_preguntas_int)
	respuestas_selec := make([]string, 0, cant_preguntas_int)

	for _, indice := range i_selec {
		preguntas_selec = append(preguntas_selec, preguntas[indice])
		respuestas_selec = append(respuestas_selec, respuestas[indice])
	}

	puntaje := 0

	for i := 0; i < cant_preguntas_int; i++ {

		fmt.Println("Enviando pregunta:", preguntas_selec[i])
		_, err = conexion.Write([]byte(preguntas_selec[i]))

		if err != nil {
			fmt.Println("Error enviando pregunta:", err)
			return
		}

		n, err = conexion.Read(buffer)

		if err != nil {
			fmt.Println("Error leyendo respuesta:", err)
			return
		}

		respuesta := string(buffer[:n])
		fmt.Println("Respuesta recibida:", respuesta)

		if respuesta == respuestas_selec[i] {
			puntaje++
		}
	}

	puntaje_str := fmt.Sprintf("%d", puntaje)
	_, err = conexion.Write([]byte(puntaje_str))

	if err != nil {
		fmt.Println("Error enviando puntaje Final:", err)
	}

	fmt.Println("Puntaje Final enviado:", puntaje_str)

	n, err = conexion.Read(buffer)
	if err != nil {
		fmt.Println("Error leyendo mensaje de finalización:", err)
		return
	}

	mensajeFinal := string(buffer[:n])
	if mensajeFinal == "Finalizado" {
		fmt.Println("Mensaje de finalización recibido del cliente.")
	}

	fmt.Println("Cerrando conexión TCP.")
}
