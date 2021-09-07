package main

import (
	"fmt"
	"go-spam-detector/common"
)

func main() {
	typ := "lemm_stop"

	train, test := common.LoadSamples(typ)
	classifier := NewBayesClassifier()

	hams, spams, corrects, totals := common.CrossValidation(train, test, classifier)

	fmt.Printf("Samples trained: %d\n", len(train))
	fmt.Printf("Samples tested: %d\n", len(test))

	fmt.Printf("Dataset: %q. Corrects: %v, Totals: %v. Accuracy: %v\n", typ, corrects, totals, float32(corrects)/float32(totals))
	fmt.Printf("Hams: %v, Spams: %v. Ratio to beat: %v\n", hams, spams, float32(hams)/float32(hams+spams))
}
