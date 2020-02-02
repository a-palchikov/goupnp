package main

// typeConvs maps from a SOAP type (e.g "fixed.14.4") to the function name
// suffix inside the soap module (e.g "Fixed14_4") and the Go type.
var typeConvs = map[string]string{
	"ui1":         "soap.Ui1",
	"ui2":         "soap.Ui2",
	"ui4":         "soap.Ui4",
	"ui8":         "soap.Ui8",
	"i1":          "soap.I1",
	"i2":          "soap.I2",
	"i4":          "soap.I4",
	"int":         "soap.Int",
	"r4":          "soap.R4",
	"r8":          "soap.R8",
	"number":      "soap.R8",
	"fixed.14.4":  "soap.Fixed14_4",
	"float":       "soap.R8",
	"char":        "soap.Char",
	"string":      "soap.String",
	"date":        "soap.Date",
	"dateTime":    "soap.DateTime",
	"dateTime.tz": "soap.DateTimeTz",
	"time":        "soap.TimeOfDay",
	"time.tz":     "soap.TimeOfDayTz",
	"boolean":     "soap.Bool",
	"bin.base64":  "soap.BinBase64",
	"bin.hex":     "soap.BinHex",
	"uri":         "soap.URI",
}
