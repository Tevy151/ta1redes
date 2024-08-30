Integrantes:  Sebastian Sandoval   Juan Alegría Vásquez
Rol:          202104654-2          202023510-4

--------------------------------------------------------------------
IMPORTANTE SOBRE PAQUETES USADOS
--------------------------------------------------------------------
- No se necesitó ninguna instalación previa de paquetes.

--------------------------------------------------------------------
AMBIENTE DE EJECUCION USADO
--------------------------------------------------------------------
- Sistema operativo: Linux Ubuntu 22.04.4 LTS
- Ejecucion del código: Visual Studio Code
- go version go1.18.1 linux/amd64

--------------------------------------------------------------------
EJECUCIÓN
--------------------------------------------------------------------
- La ejecución se hizo por consola, una consola por programa:

- Consola server: go run server.go
- Consola cliente: go run cliente.go 127.0.0.1:1234

--------------------------------------------------------------------
CONSIDERACIONES AL MOMENTO DE EJECUTAR
--------------------------------------------------------------------

- Una vez que el cliente se conecte al servidor, se le proporcionarán algunas preguntas que deberá responder. Al finalizar, se le mostrará un puntaje final.
- El programa server no envia al cliente si la respuesta que recibio fue correcta o no, por lo que el cliente no sabrá el puntaje final hasta que responda todo,
	esto está así porque al implementar esa funcionalidad, el codigo aveces dejaba de funcionar.
- Una vez terminado el programa cliente, el servidor seguirá corriendo, para cortarlo o parar su funcionamiento, hay que apretar ctrl+c
