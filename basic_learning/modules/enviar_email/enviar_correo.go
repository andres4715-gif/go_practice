package main

/*
Este correo se envía usando un módulo externo en Go https://github.com/go-gomail/gomail
- Tambien se utilizo el servicio de https://mailtrap.io
- con el link: https://mailtrap.io/inboxes/4438940/messages/5371965058
- Se uso usuario y contraseña de la cuenta de mailtrap
*/

import (
	"github.com/joho/godotenv" // Nuestro nuevo paquete
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. EXTRAEMOS LAS CONTRASEÑAS A VARIABLES SEGURAS
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	// 1. "Armamos" el carrito de compras (el mensaje)
	m := gomail.NewMessage()

	// Configuramos cabeceras básicas
	m.SetHeader("From", "andres4715@gmail.com")
	m.SetHeader("To", "andres4715@gmail.com	")
	m.SetHeader("Subject", "¡Mi segundo correo desde Go!")

	// Configuramos el cuerpo del correo (Soporta HTML)
	m.SetBody("text/html", "<b>¡Felicidades!</b> Acabas de enviar tu primer email usando un módulo externo en Go 🚀.")

	// 2. Configuramos al "Cartero" (Servidor SMTP)
	// Parámetros: Host, Puerto, Usuario, Contraseña
	dialer := gomail.NewDialer(
		"sandbox.smtp.mailtrap.io", // Host
		2525,                       // Puerto (sin comillas porque es un número)
		smtpUser,                   // Usuario desde la variable de entorno
		smtpPass,                   // Contraseña desde la variable de entorno
	)
	// 3. Enviamos el correo y manejamos el error al estilo Senior
	log.Println("⏳ Enviando correo, por favor espera...")

	if err := dialer.DialAndSend(m); err != nil {
		log.Fatalf("🚨 Error crítico al enviar el correo: %v", err)
	}

	// Si llegamos aquí, el error fue 'nil'
	log.Println("✅ ¡Correo enviado exitosamente!")
}
