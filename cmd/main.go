package main

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/gocarina/gocsv"
	"github.com/gptlv/woffs/internal/record"
	"github.com/gptlv/woffs/internal/util"
	"github.com/lukasjarosch/go-docx"
)

var inputFile = "write_off_records.csv"
var templates = []string{"commitee", "dismissal", "record"}
var templatesFolder = "templates"
var outputFolder = "output"

func main() {
	err := GenerateWriteOffDocuments()
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateWriteOffDocuments() error {
	dismissalRecords := []*record.Record{}

	file, err := util.ReadInputFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file %s: %w", inputFile, err)
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &dismissalRecords); err != nil {
		return fmt.Errorf("failed to unmarshal input file %s: %w", inputFile, err)
	}

	for _, record := range dismissalRecords {
		recordFolder := record.ISC

		recordFolderPath := filepath.Join(outputFolder, recordFolder)
		log.Info(fmt.Sprintf("creating directory %s", recordFolderPath))

		_, err := util.CreateDirectory(recordFolderPath)
		if err != nil {
			return fmt.Errorf("failed to create folder for %s: %w", record.ISC, err)
		}
	}

	for _, record := range dismissalRecords {
		recordFolder := record.ISC
		placeholderMap := docx.PlaceholderMap{}
		recordMap := util.StructToMap(record)

		for key, value := range recordMap {
			placeholderMap[key] = value
		}

		for _, template := range templates {
			inputFile := filepath.Join(templatesFolder, template+".docx")

			log.Info(fmt.Sprintf("reading file: %v", inputFile))
			doc, err := docx.Open(inputFile)
			if err != nil {
				return fmt.Errorf("failed to open docx template %v: %w", template, err)
			}

			log.Info(fmt.Sprintf("generating document from template %s", template))
			err = doc.ReplaceAll(placeholderMap)
			if err != nil {
				return fmt.Errorf("failed to replace values in document %v:%w", file, err)
			}

			outputFile := filepath.Join(outputFolder, recordFolder, template+".docx")
			log.Info(fmt.Sprintf("writing to %s", outputFile))
			err = doc.WriteToFile(outputFile)
			if err != nil {
				return fmt.Errorf("failed to write output file:%w", err)
			}

			log.Info(fmt.Sprintf("finished generating %s\n", outputFile))

		}
	}

	log.Info("done")

	return nil
}
