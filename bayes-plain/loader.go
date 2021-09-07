package main

import (
	"encoding/csv"
	"fmt"
	"go-spam-detector/common"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LoadSamples() (train []common.Sample, test []common.Sample) {
	total, err := load()
	if err != nil {
		log.Fatal(err)
	}
	common.Shuffle(total)
	cvStart := len(total) - len(total)/3

	return total[:cvStart], total[cvStart:]
}

func load() (samples []common.Sample, err error) {
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
		content := cleanContent(fmt.Sprintf("%s %s", l[0], l[1]))
		class, _ := strconv.ParseInt(l[2], 10, 32)

		samples = append(samples, common.Sample{strings.Split(content, " "), common.Class(class)})
	}

	return
}

var regMap = map[*regexp.Regexp]string{
	regexp.MustCompile("won't"): "will not",
	regexp.MustCompile("can't"): "can not",

	regexp.MustCompile("n't"): " not",
	regexp.MustCompile("'re"): " are",
	regexp.MustCompile("'d"):  " would",
	regexp.MustCompile("'ll"): " will",
	regexp.MustCompile("'t"):  " not",
	regexp.MustCompile("'ve"): " have",
	regexp.MustCompile("'m"):  " am",
	//regexp.MustCompile("'s"):  " is",

	regexp.MustCompile("\\d+(\\.\\d+)?"): "numbers",
	regexp.MustCompile("\\s+"):           " ",
}

func cleanContent(s string) string {
	b := []byte(s)

	for re, rp := range regMap {
		b = re.ReplaceAll([]byte(s), []byte(rp))
	}

	return string(b)
}
