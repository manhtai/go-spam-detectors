package main

import "go-spam-detectors/common"

func main() {
	classifier := NewBayes3Classifier()
	common.Summary(classifier)
}
