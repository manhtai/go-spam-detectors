package main

import "go-spam-detector/common"

func main() {
	classifier := NewBayes2Classifier()
	common.Summary(classifier)
}
