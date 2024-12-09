package coding

import (
	"log"
	"os"

	"strings"

	DF "github.com/go-gota/gota/dataframe"
)

func ReadCSVFile(filePath string) DF.DataFrame {
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	reader := strings.NewReader(string(f))

	df := DF.ReadCSV(reader)
	return df
}
