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
	resolvU, err := net.ResolveUDPAddr("udp", ":1234")
	if err != nil {
		fmt.Println("Error resolviendo la dirección UDP:", err)
		return
	}
	conectionU, err := net.ListenUDP("udp", resolvU)
	if err != nil {
		fmt.Println("Error al escuchar UDP:", err)
		return
	}
	defer conectionU.Close()
	fmt.Println("Servidor UDP escuchando en", resolvU.String())
	buffer := make([]byte, 1024)

	for {
		_, conexionCon, err := conectionU.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error leyendo desde UDP:", err)
			continue
		}
		fmt.Println("Usuario conectado desde", conexionCon.String())

		preguntas := rand.Intn(5) + 3 // Número aleatorio entre 3 y 7
		datos := fmt.Sprintf("%d,%s,%s", preguntas, "127.0.0.1", "4321")
		_, err = conectionU.WriteToUDP([]byte(datos), conexionCon)
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

	resolvT, err := net.ResolveTCPAddr("tcp", ":4321")
	if err != nil {
		fmt.Println("Error resolviendo la dirección TCP:", err)
		return
	}
	conectionT, err := net.ListenTCP("tcp", resolvT)
	if err != nil {
		fmt.Println("Error al escuchar TCP:", err)
		return
	}
	defer conectionT.Close()
	fmt.Println("Servidor TCP escuchando en", resolvT.String())

	conexion, err := conectionT.AcceptTCP()
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

	cantidadPreguntasStr := string(buffer[:n])
	cantidadPreguntasInt, err := strconv.Atoi(cantidadPreguntasStr)
	if err != nil {
		fmt.Println("Error convirtiendo preguntas a número:", err)
		return
	}

	fmt.Println("TCP debe dar", cantidadPreguntasStr, "preguntas")

	indicesSeleccionados := rand.Perm(len(preguntas))[:cantidadPreguntasInt]
	preguntasSeleccionadas := make([]string, 0, cantidadPreguntasInt)
	respuestasSeleccionadas := make([]string, 0, cantidadPreguntasInt)

	for _, indice := range indicesSeleccionados {
		preguntasSeleccionadas = append(preguntasSeleccionadas, preguntas[indice])
		respuestasSeleccionadas = append(respuestasSeleccionadas, respuestas[indice])
	}

	puntaje := 0

	for i := 0; i < cantidadPreguntasInt; i++ {

		fmt.Println("Enviando pregunta:", preguntasSeleccionadas[i])
		_, err = conexion.Write([]byte(preguntasSeleccionadas[i]))
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
		if respuesta == respuestasSeleccionadas[i] {
			puntaje++
		}

		puntajeStr := fmt.Sprintf("%d", puntaje)
		_, err = conexion.Write([]byte(puntajeStr))
		if err != nil {
			fmt.Println("Error enviando puntaje:", err)
		}
	}

	puntajeStr := fmt.Sprintf("%d", puntaje)
	_, err = conexion.Write([]byte(puntajeStr))
	if err != nil {
		fmt.Println("Error enviando puntaje Final:", err)
	}
	fmt.Println("Puntaje Final enviado:", puntajeStr)
}
