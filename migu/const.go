package migu

import (
	`errors`
)

var (
	ParamMissing     = errors.New("param missing")
	ParamMissSubject = errors.New("param miss subject")
)
