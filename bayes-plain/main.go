package main

import (
	"fmt"
	"go-spam-detector/common"
	"strings"
)

func main() {
	train, test := LoadSamples()

	fmt.Printf("Total samples: %d\n", len(train)+len(test))
	fmt.Printf("Samples to train: %d\n", len(train))
	fmt.Printf("Samples to test: %d\n", len(test))

	classifier := NewBayesClassifier()
	classifier.Train(train)

	tp, fp, tn, fn := common.Predict(classifier, test)
	fmt.Printf("Accuracy: %v, Precision: %v\n", float32(tp+tn)/float32(tp+fp+tn+fn), float32(tp)/float32(tp+fp))

	fmt.Println("======================================================================")
	t1 := "your microsoft account has been compromised, you must update before or else your account going to close click to update"
	t2 := "Today we want to inform you that the application period for 15.000 free Udacity Scholarships in Data Science is now open! Please apply by November 16th, 2020 via https://www.udacity.com/bertelsmann-tech-scholarships."
	fmt.Println("t1: ", classifier.Predict(strings.Split(t1, " ")))
	fmt.Println("t2: ", classifier.Predict(strings.Split(t2, " ")))

}
