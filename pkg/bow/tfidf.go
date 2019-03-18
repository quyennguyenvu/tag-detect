package bow

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
	"tag-detect/storage"
)

var dsPath = "./pkg/bow/dataset/"
var stopwords []string

// Init ...
func Init() {
	stopwords = dataset("stopwords_simple.json")
}

func dataset(filename string) (result []string) {
	jsonFile, err := os.Open(dsPath + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(byteValue), &result)

	return
}

func extractWords(document string) (words []string) {
	document = strings.ToLower(document)
	reg := regexp.MustCompile(`[\p{L}\d_]+`)
	words = reg.FindAllString(document, -1)

	return
}

func createDict(words []string) (wordDict map[string]int) {
	for _, w := range words {
		if _, ok := wordDict[w]; ok {
			wordDict[w]++
		} else {
			wordDict[w] = 1
		}
	}

	return
}

func computeTF(words []string, wordDict map[string]int) (tfDict map[string]float64) {
	lenW := len(words)
	for word, count := range wordDict {
		tfDict[word] = float64(count) / float64(lenW)
	}
	return
}

type bookingHandlerImpl struct {
	bookingDAO storage.DocumentStorage
}

func computeTFIDF(d *storage.Document, mapOneW map[string]storage.OneWord, tfDict map[string]float64) (tfidf map[string]float64) {

	// docs := &bookingHandlerImpl{
	// 	bookingDAO: storage.NewDocumentStorage(),
	// }
	// doc := docs.bookingDAO.ByID(1)

	for w, c := range tfDict {
		appear := 1
		if _, ok := mapOneW[w]; ok {
			appear = mapOneW[w].Appear
		}
		idf := math.Log(float64(d.Count) / float64(appear))
		tfidf[w] = c * idf
	}
	updateDB()
	return
}

func updateDB() {

}
