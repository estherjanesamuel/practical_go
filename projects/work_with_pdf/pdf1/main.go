/*
 * PDF to text: Extract all text for each page of a pdf file.
 *
 * Run as: go run pdf_extract_text.go input.pdf
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run pdf_extract_text.go input.pdf\n")
		os.Exit(1)
	}

	const unidocLicenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
eyJsaWNlbnNlX2lkIjoiZWM2YjNkMjItYjNkOS00NjkyLTc3ODQtYTMzOTkyNzUyYjhhIiwiY3VzdG9tZXJfaWQiOiIzMmI0NzQ3Ni0yMTI2LTQ1NGYtN2E2MS0wMjFiZDMwZTc0YTQiLCJjdXN0b21lcl9uYW1lIjoiRGlnaXRhbCBBbHBoYSBJbmRvbmVzaWEiLCJjdXN0b21lcl9lbWFpbCI6ImFuZHJlLnByYXRhbWFAdWFuZ3RlbWFuLmNvbSIsInRpZXIiOiJidXNpbmVzcyIsImNyZWF0ZWRfYXQiOjE1NzUwMTIyOTIsImV4cGlyZXNfYXQiOjE2MDY2OTQzOTksImNyZWF0b3JfbmFtZSI6IlVuaURvYyBTdXBwb3J0IiwiY3JlYXRvcl9lbWFpbCI6InN1cHBvcnRAdW5pZG9jLmlvIiwidW5pcGRmIjp0cnVlLCJ1bmlvZmZpY2UiOnRydWUsInRyaWFsIjpmYWxzZX0=
+
VaLlarl3MB1oaoqp9OukhnBaZi91Kse1JJ27bVdjuN6CrpurccwNx5CU8kcLAxFNqKzlaiBJmbHkYvl2T7HGLYGnErrg2DhF2e2JfFnb7lc78uouEdjS5raf9dmjZsjK9FNDKGRUQxYhRPYU22/GKGrrHcaJH5FCsaBLRVBOMWIUraFktEcwzaLILneZ/Dgsdv+ymIHowHDdViM4wwdkWb7fIonuy/QsDgBygBUkmPvZzjFQ+HRnFnqTltP/DZwez9cOemVS5IUl4/80Y1kVnD+63K8kSBrr8lyBYQhJR0DTez0Z08s2KJeWCRwLJdnvR5EJzvkld0Q0NmcS+9x03g==
-----END UNIDOC LICENSE KEY-----`

err := license.SetLicenseKey(unidocLicenseKey,"Digital Alpha Indonesia")
	if err != nil {
		fmt.Println("PDF Error loading license, err:", err)
		os.Exit(1)
	}

	
	// Make sure to enter a valid license key.
	// Otherwise text is truncated and a watermark added to the text.
	// License keys are available via: https://unidoc.io
	/*
			license.SetLicenseKey(`
		-----BEGIN UNIDOC LICENSE KEY-----
		...key contents...
		-----END UNIDOC LICENSE KEY-----
		`)
	*/

	// For debugging.
	// common.SetLogger(common.NewConsoleLogger(common.LogLevelDebug))

	inputPath := os.Args[1]

	err = outputPdfText(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// outputPdfText prints out contents of PDF file to stdout.
func outputPdfText(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return err
		}

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		fmt.Printf("\"%s\"\n", text)
		fmt.Println("------------------------------")
	}

	return nil
}