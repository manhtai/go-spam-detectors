package main

import (
	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
	"go-spam-detector/common"
	"strings"
)

type bayesClassifier struct {
	model  *text.NaiveBayes
	stream chan base.TextDatapoint
}

func NewBayes2Classifier() *bayesClassifier {
	stream := make(chan base.TextDatapoint, 40)
	model := text.NewNaiveBayes(stream, 2, base.OnlyWordsAndNumbers)
	return &bayesClassifier{
		stream: stream,
		model:  model,
	}
}

func (b *bayesClassifier) Train(samples []common.Sample) {
	errors := make(chan error)
	go b.model.OnlineLearn(errors)

	for _, sam := range samples {
		b.stream <- base.TextDatapoint{
			X: strings.Join(sam.Content, " "),
			Y: uint8(sam.Class),
		}
	}

	close(b.stream)
}

func (b *bayesClassifier) Predict(doc []string) common.Class {
	return common.Class(b.model.Predict(strings.Join(doc, " ")))
}
