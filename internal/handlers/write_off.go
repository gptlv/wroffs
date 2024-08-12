package handlers

import (
	"fmt"
	"os"
	"time"
	"write-off-docs-generator/internal/domain"
	"write-off-docs-generator/internal/interfaces"
	"write-off-docs-generator/util"

	"github.com/charmbracelet/log"
	"github.com/gocarina/gocsv"
)

type writeOffHandler struct {
	writeOffService interfaces.WriteOffService
}

func NewWriteOffHandler(writeOffService interfaces.WriteOffService) *writeOffHandler {
	return &writeOffHandler{writeOffService: writeOffService}
}

func (writeOffHandler *writeOffHandler) GenerateWriteOffDocuments() error {
	inputFile, err := os.Open("write_off_records.csv")
	if err != nil {
		return fmt.Errorf("failed to open csv file: %w", err)
	}
	defer inputFile.Close()

	dismissalRecords := []*domain.Record{}

	if err := gocsv.UnmarshalFile(inputFile, &dismissalRecords); err != nil {
		panic(err)
	}

	for _, record := range dismissalRecords {
		templateNames := []string{"commitee", "dismissal", "record"}

		log.Info(fmt.Sprintf("creating output directory %v", record.ISC))
		dirPath, err := util.CreateOutputDirectory(record.ISC)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		for _, templateName := range templateNames {
			log.Info(fmt.Sprintf("creating template %v", templateName))
			template, err := writeOffHandler.writeOffService.CreateTemplate(record, templateName)
			if err != nil {
				return fmt.Errorf("failed to create template %v: %w", templateName, err)
			}

			log.Info(fmt.Sprintf("creating object from template %v", templateName))
			object, err := writeOffHandler.writeOffService.CreateObjectFromTemplate(template)
			if err != nil {
				return fmt.Errorf("failed to create object from template %v: %w", templateName, err)
			}

			log.Info(fmt.Sprintf("creating file %v", dirPath))
			file, err := util.CreateFile(dirPath, templateName, "pdf")
			if err != nil {
				return fmt.Errorf("failed to create file")
			}
			defer file.Close()

			log.Info("creating pdf")
			err = writeOffHandler.writeOffService.CreatePDF(object, file)
			if err != nil {
				return fmt.Errorf("failed to generate pdf: %w", err)
			}

			fmt.Println()
			time.Sleep(time.Second)
		}
	}

	return nil
}
