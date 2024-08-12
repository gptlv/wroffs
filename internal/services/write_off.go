package services

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"write-off-docs-generator/internal/domain"
	"write-off-docs-generator/internal/interfaces"

	pdf "github.com/adrg/go-wkhtmltopdf"
)

type writeOffService struct{}

func NewWriteOffService() interfaces.WriteOffService {
	return &writeOffService{}
}

func (s *writeOffService) CreateTemplate(record *domain.Record, templateName string) ([]byte, error) {
	filepath := fmt.Sprintf("templates/%v.html", templateName)

	tmpl := template.Must(template.ParseFiles(filepath))
	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, record); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *writeOffService) CreateObjectFromTemplate(template []byte) (*pdf.Object, error) {
	if template == nil {
		return nil, fmt.Errorf("empty template")
	}

	reader := bytes.NewReader((template))

	object, err := pdf.NewObjectFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new object from reader: %w", err)
	}

	return object, nil
}

func (s *writeOffService) CreatePDF(object *pdf.Object, outputFile *os.File) error {
	if object == nil {
		return fmt.Errorf("empty object")
	}

	if outputFile == nil {
		return fmt.Errorf("empty output file")
	}

	converter, err := pdf.NewConverter()
	if err != nil {
		return fmt.Errorf("failed to create pdf converter: %w", err)
	}
	// defer converter.Destroy()

	converter.Add(object)

	converter.Title = "Sample document"
	converter.PaperSize = pdf.A4
	converter.Orientation = pdf.Portrait
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "1cm"
	converter.MarginRight = "1cm"

	if err := converter.Run(outputFile); err != nil {
		return fmt.Errorf("failed to run converter: %w", err)
	}

	return nil
}

// func (s *writeOffService) ReadCsvFile(filePath string) ([][]string, error) {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read input file %v: %w", filePath, err)
// 	}
// 	defer f.Close()

// 	csvReader := csv.NewReader(f)
// 	data, err := csvReader.ReadAll()
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to parse file as CSV for %v: %w", filePath, err)
// 	}

// 	return data, nil
// }
