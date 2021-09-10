package main

import "go-spam-detectors/common"

func main() {
	classifier := NewBayes2Classifier()
	common.Summary(classifier)
}
