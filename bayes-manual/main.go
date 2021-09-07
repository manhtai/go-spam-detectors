package main

import (
	"fmt"
	"go-spam-detector/common"
)

func main() {
	train, test := common.LoadSamples()

	fmt.Printf("Total samples: %d\n", len(train)+len(test))
	fmt.Printf("Samples to train: %d\n", len(train))
	fmt.Printf("Samples to test: %d\n", len(test))

	classifier := NewBayesClassifier()
	hams, spams, corrects, totals := common.CrossValidation(train, test, classifier)

	fmt.Printf("Corrects: %v, Totals: %v. Accuracy: %v\n", corrects, totals, float32(corrects)/float32(totals))
	fmt.Printf("Hams: %v, Spams: %v. Ratio to beat: %v\n", hams, spams, float32(hams)/float32(hams+spams))
}
