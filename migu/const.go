package migu

import (
	"errors"
	"net/http"
)

var (
	httpClient       = &http.Client{}
	ParamMissing     = errors.New("param missing")
	ParamMissSubject = errors.New("param miss subject")
)
