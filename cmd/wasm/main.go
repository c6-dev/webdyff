//go:build js && wasm
package main

import (
	"bytes"
	"syscall/js"

	"github.com/gonvenience/ytbx"
	"github.com/homeport/dyff/pkg/dyff"
)

func computeDiff(this js.Value, args []js.Value) interface{} {
	if len(args) < 2 {
		return "error: two arguments required (from, to)"
	}

	fromStr := args[0].String()
	toStr := args[1].String()

	fromInput, err := ytbx.LoadDocuments([]byte(fromStr))
	if err != nil {
		return "error parsing 'from': " + err.Error()
	}

	toInput, err := ytbx.LoadDocuments([]byte(toStr))
	if err != nil {
		return "error parsing 'to': " + err.Error()
	}

	fromFile := ytbx.InputFile{
		Location:  "from",
		Documents: fromInput,
	}
	toFile := ytbx.InputFile{
		Location:  "to",
		Documents: toInput,
	}

	report, err := dyff.CompareInputFiles(fromFile, toFile)
	if err != nil {
		return "error comparing: " + err.Error()
	}

	var buf bytes.Buffer
	hr := dyff.HumanReport{
		Report:            report,
		OmitHeader:        true,
		NoTableStyle:      true,
		DoNotInspectCerts: true,
	}
	if err := hr.WriteReport(&buf); err != nil {
		return "error writing report: " + err.Error()
	}

	return buf.String()
}

func main() {
	js.Global().Set("computeDiff", js.FuncOf(computeDiff))
	select {}
}
