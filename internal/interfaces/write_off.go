package interfaces

import (
	"os"
	"write-off-docs-generator/internal/domain"

	pdf "github.com/adrg/go-wkhtmltopdf"
)

type WriteOffService interface {
	CreateTemplate(record *domain.Record, templateName string) ([]byte, error)
	CreateObjectFromTemplate(template []byte) (*pdf.Object, error)
	CreatePDF(object *pdf.Object, outputFile *os.File) error
}
