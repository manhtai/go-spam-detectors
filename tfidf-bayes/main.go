package main

import (
	"fmt"
	"log"
)

func main() {
	typ := "lemm_stop"
	total, err := ingest(typ)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total loaded: %d\n", len(total))
	shuffle(total)
	cvStart := len(total) - len(total)/3
	test := total[cvStart:]
	train := total[:cvStart]

	c := New()
	c.Train(train)
	fmt.Printf("Samples trained: %d\n", len(train))

	corrects, totals := 0, 0
	hams, spams := 0.0, 0.0
	var unseen, totalWords int
	for _, ex := range test {
		totalWords += len(ex.Document)
		unseen += c.unseens(ex.Document)
		class := c.Predict(ex.Document)
		if class == ex.Class {
			corrects++
		}
		switch ex.Class {
		case Ham:
			hams++
		case Spam:
			spams++
		}
		totals++
	}

	fmt.Printf("Samples tested: %d\n", len(test))
	fmt.Printf("Dataset: %q. Corrects: %v, Totals: %v. Accuracy: %v\n", typ, corrects, totals, float32(corrects)/float32(totals))
	fmt.Printf("Hams: %v, Spams: %v. Ratio to beat: %v\n", hams, spams, hams/(hams+spams))
	fmt.Printf("Previously unseen %d. Total Words %d\n", unseen, totalWords)
}
