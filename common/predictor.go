package common

import (
	"fmt"
	"strings"
	"time"
)

func Summary(classifier Classifier) {
	train, test := loadSamples()

	start := time.Now()
	classifier.Train(train)

	fmt.Println("======================================================================")
	fmt.Printf("Total samples: %d\n", len(train)+len(test))
	fmt.Printf("Samples to train: %d\n", len(train))
	fmt.Printf("Samples to test: %d\n", len(test))

	tp, fp, tn, fn := predict(classifier, test)

	fmt.Println("======================================================================")
	fmt.Printf("Accuracy: %v\nPrecision: %v\nRecall: %v\n",
		float32(tp+tn)/float32(tp+fp+tn+fn),
		float32(tp)/float32(tp+fp),
		float32(tp)/float32(tp+fn),
	)

	fmt.Println("======================================================================")
	t1 := "your microsoft account has been compromised, you must update before or else your account going to close click to update"
	t2 := "Today we want to inform you that the application period for 15.000 free Udacity Scholarships in Data Science is now open! Please apply by November 16th, 2020 via https://www.udacity.com/bertelsmann-tech-scholarships."
	fmt.Println("t1:", classifier.Predict(strings.Split(t1, " ")), "=> should be Spam")
	fmt.Println("t2:", classifier.Predict(strings.Split(t2, " ")), "=> should be Ham")
	fmt.Println("======================================================================")
	fmt.Printf("Processing time: %dms\n", -time.Until(start).Milliseconds())
}

func predict(classifier Classifier, test []Sample) (tp, fp, tn, fn int) {
	for _, ex := range test {
		class := classifier.Predict(ex.Content)
		switch class {
		case Ham:
			if ex.Class == Ham {
				tp++
			} else {
				fp++
			}
		case Spam:
			if ex.Class == Spam {
				tn++
			} else {
				fn++
			}
		}
	}
	return
}
