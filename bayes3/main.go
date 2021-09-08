package main

import "go-spam-detector/common"

func main() {
	classifier := NewBayes3Classifier()
	common.Summary(classifier)
}
