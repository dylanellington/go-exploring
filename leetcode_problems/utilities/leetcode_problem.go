package main

import (
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"os"
	"strings"
)

var ProblemNotFound = errors.New("Leetcode problem was not found.")

var leetcodeProblemUrls = map[int]string {
	1: "https://leetcode.com/problems/two-sum/",
	2: "https://leetcode.com/problems/add-two-numbers/",
	3: "https://leetcode.com/problems/longest-substring-without-repeating-characters/",
	4: "https://leetcode.com/problems/median-of-two-sorted-arrays/",
	17: "https://leetcode.com/problems/letter-combinations-of-a-phone-number/",
	20: "https://leetcode.com/problems/valid-parentheses/",
	28: "https://leetcode.com/problems/implement-strstr/",
	33: "https://leetcode.com/problems/search-in-rotated-sorted-array/",
	94: "https://leetcode.com/problems/binary-tree-inorder-traversal/",
	101: "https://leetcode.com/problems/symmetric-tree/",
	102: "https://leetcode.com/problems/binary-tree-level-order-traversal/",
	104: "https://leetcode.com/problems/maximum-depth-of-binary-tree/",
	112: "https://leetcode.com/problems/path-sum/",
	133: "https://leetcode.com/problems/clone-graph/",
	144: "https://leetcode.com/problems/binary-tree-preorder-traversal/",
	145: "https://leetcode.com/problems/binary-tree-postorder-traversal/",
	150: "https://leetcode.com/problems/evaluate-reverse-polish-notation/",
	191: "https://leetcode.com/problems/number-of-1-bits/",
	200: "https://leetcode.com/problems/number-of-islands/",
	286: "https://leetcode.com/problems/walls-and-gates/",
	341: "https://leetcode.com/problems/flatten-nested-list-iterator/",
	346: "https://leetcode.com/problems/moving-average-from-data-stream/",
	739: "https://leetcode.com/problems/daily-temperatures/",
	2220: "https://leetcode.com/problems/minimum-bit-flips-to-convert-number/",
}

type LeetcodeProblem struct {
	number int
	url string
}

func NewLeetcodeProblem(number int) (LeetcodeProblem, error) {
	url, exists := leetcodeProblemUrls[number]

	if !exists {
		return LeetcodeProblem{}, ProblemNotFound
	}

	problem := LeetcodeProblem {
		number: number,
		url: url,
	}
	return problem, nil
}

func (problem *LeetcodeProblem) CreateReadme(chromium *Chromium, path string) error {
	var htmlResult string

	err := chromium.Run(
		chromedp.Navigate(problem.url),
		chromedp.WaitVisible(LeetcodeHtmlSelectorMap["ProblemQuestionContent"], chromedp.ByQuery),
		chromedp.OuterHTML("html", &htmlResult))

	if err != nil {
		return err
	}

	reader := strings.NewReader(htmlResult)
	lcCrawler := NewLeetcodeCrawler(reader)

	titleHtml := fmt.Sprintf("<h2>%v</h2>\n", lcCrawler.GetProblemTitle())
	difficulty := lcCrawler.GetProblemDifficulty()
	difficultyHtml := fmt.Sprintf("<h3 %v>%v</h3>\n", getProblemDifficultyStyle(difficulty), difficulty)
	descriptionHtml := lcCrawler.GetProblemDescriptionHtml()

	readmeContent := titleHtml + difficultyHtml + "<hr>\n" + descriptionHtml

	if err := os.WriteFile(path + "/README.md", []byte(readmeContent), 0666); err != nil {
		return err
	}

	return nil
}

func getProblemDifficultyStyle(s string) string {
	switch s {
	case "Easy":
		return "style=\"color:Green;\""
	case "Medium":
		return "style=\"color:Yellow;\""
	case "Hard":
		return "style=\"color:Red;\""
	default:
		return "style=\"color:Black;\""
	}
}
