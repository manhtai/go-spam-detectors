package common

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadSamples() (train []Sample, test []Sample) {
	total, err := ingest()
	if err != nil {
		log.Fatal(err)
	}
	shuffle(total)
	cvStart := len(total) - len(total)/3

	return total[:cvStart], total[cvStart:]
}

func ingest() (samples []Sample, err error) {
	f, err := os.Open("datasets/messages.csv")
	if err != nil {
		return samples, err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return samples, err
	}

	for _, l := range lines {
		content := fmt.Sprintf("%s %s", l[0], l[1])
		class, _ := strconv.ParseInt(l[2], 10, 32)

		samples = append(samples, Sample{strings.Split(content, " "), Class(class)})
	}

	return
}
