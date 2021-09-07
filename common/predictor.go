package common

func Predict(classifier Classifier, test []Sample) (tp, fp, tn, fn int) {
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
