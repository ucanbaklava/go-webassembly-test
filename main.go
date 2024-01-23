package main

import (
	"bytes"
	"syscall/js"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
)

func generatePDF(this js.Value, args []js.Value) interface{} {
	gofpdf.SetDefaultCompression(false)
	gofpdf.SetDefaultCatalogSort(true)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA5, "")

	w, h := pdf.GetPageSize()
	pdf.AddPage()

	pdf.SetFillColor(0, 0, 0)
	pdf.SetTextColor(255, 255, 255)

	pdf.Rect(0, 0, 3, h, "DF")
	pdf.Rect(0, 0, w, 3, "DF")
	pdf.Rect(w-3, 0, 4, h, "DF")

	pdf.SetFillColor(255, 255, 255)

	// CARRIER IMAGE

	// BARCODE
	pdf.TransformBegin()
	pdf.TransformRotate(90.0, w-60.0, 300.0)

	_, lineHt := pdf.GetFontSize()
	barcodeWidth := 290.0
	barcodeHeight := lineHt
	encodedBarcodeString := barcode.RegisterCode128(pdf, "324235253")
	barcode.BarcodeUnscalable(pdf, encodedBarcodeString, w-60.0, 300.0, &barcodeWidth, &barcodeHeight, false)

	pdf.SetTextColor(0, 0, 0)

	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformRotate(270.0, 348, 18)
	pdf.SetFontSize(12)
	pdf.Text(343, 18, "324235253")
	pdf.TransformEnd()

	// BARCODE END

	oy := 60.0

	// ORDER INFO
	pdf.Line(10, oy, w-90, oy)
	pdf.SetTextColor(0, 0, 0)

	pdf.Text(10, oy+15, "GÖNDERİ TİPİ:")
	pdf.SetFontSize(16)
	pdf.Text(10, oy+30, "NORMAL GÖNDERİ")

	pdf.Text((w-90)/2, oy+15, "ÖDEME TİPİ:")
	pdf.SetFontSize(16)
	pdf.Text((w-90)/2, oy+30, "ALICI ÖDEMELİ")

	out := bytes.Buffer{}

	err := pdf.Output(&out)

	if err != nil {
		panic(err)
	}

	size := len(out.Bytes())
	result := js.Global().Get("Uint8Array").New(size)

	js.CopyBytesToJS(result, out.Bytes())

	return result

}

func main() {
	js.Global().Set("generatePDF", js.FuncOf(generatePDF))

	<-make(chan bool)
}
