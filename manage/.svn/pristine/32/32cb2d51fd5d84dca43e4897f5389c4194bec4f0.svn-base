package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

// We use Builder-Pattern to export file
// About Builder-Pattern Definition:
// Separate the construction of a complex object from its representation,
// so that the same construction process can create different representations.

// Builder
type ExportFile struct {
	MerchantId int
	Req        interface{}
	Condition  map[string]interface{}
	FiLeName   string
	C          *gin.Context
	File       *xlsx.File
	Sheet      *xlsx.Sheet
	Error      bool
}

type FileBuilder interface {
	Init()
	SetFileHeader()
	SetFileBody()
	WriteFile2Web()
}

// Director Definition
type FileDirector struct {
}

func NewFileDirector() *FileDirector {
	return &FileDirector{}
}

func (fd *FileDirector) ExportFile(fb FileBuilder) {
	fb.Init()
	fb.SetFileHeader()
	fb.SetFileBody()
	fb.WriteFile2Web()
}
