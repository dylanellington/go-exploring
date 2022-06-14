package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

var (
	LeetcodeHtmlSelectorMap = map[string]string {
		"ProblemQuestionTitle": `div[data-cy=question-title]`,
		"ProblemQuestionDifficulty": `div[diff]`,
		"ProblemQuestionContent": `div[class*="question-content"]`,
	}
)

type LeetcodeCrawler struct {
	doc *goquery.Document
}

func NewLeetcodeCrawler(reader io.Reader) LeetcodeCrawler {
	crawler := LeetcodeCrawler{}
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		panic(err)
	}

	crawler.doc = doc
	return crawler
}

func (lc *LeetcodeCrawler) GetProblemTitle() string {
	return lc.doc.Find(LeetcodeHtmlSelectorMap["ProblemQuestionTitle"]).First().Text()
}

func (lc *LeetcodeCrawler) GetProblemDifficulty() string {
	return lc.doc.Find(LeetcodeHtmlSelectorMap["ProblemQuestionDifficulty"]).First().Text()
}

func (lc *LeetcodeCrawler) GetProblemDescriptionHtml() string {
	result, err :=  lc.doc.Find(LeetcodeHtmlSelectorMap["ProblemQuestionContent"]).First().Html()

	if err != nil {
		panic(err)
	}

	return result
}
