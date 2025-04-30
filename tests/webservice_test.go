package test

import (
	"testing"

	"github.com/A-Oez/gpc/internal/webservice"
)

func TestWebserviceCreation(t *testing.T) {
	TestDBFileCreation(t)

	ws := webservice.Webservice{
		ProjectName: project_name,
	}
	ws.Execute()
}
