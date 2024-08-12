package main

import (
	"fmt"
	"log"
	"write-off-docs-generator/internal/handlers"
	"write-off-docs-generator/internal/services"

	pdf "github.com/adrg/go-wkhtmltopdf"
)

func main() {
	if err := pdf.Init(); err != nil {
		log.Fatal(fmt.Errorf("failed to initialize pdf: %w", err))
	}
	defer pdf.Destroy()

	writeOffService := services.NewWriteOffService()
	writeOffHandler := handlers.NewWriteOffHandler(writeOffService)

	writeOffHandler.GenerateWriteOffDocuments()

}
