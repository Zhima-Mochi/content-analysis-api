package main

import (
	"context"
	"fmt"
	"os"

	contentAnalysis "github.com/Zhima-Mochi/content-analysis-api/content-analysis"
)

var contentAnalysisHandler *contentAnalysis.ContentAnalysisHandler

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	contentAnalysisHandler = contentAnalysis.NewContentAnalysisHandler(apiKey)
}

func main() {
	ctx := context.Background()

	// texts := []string{
	// 	"我不要听话",
	// 	"you are so stupid",
	// 	"每天早起做愛心便當",
	// 	"この野郎！",
	// 	"幹你娘",
	// }
	// for _, text := range texts {
	// ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
	// defer cancel()

	// fmt.Println("text: " + text)
	// result1, err := contentAnalysisHandler.SensitiveWordsDetection(ctxWithTimeout, text)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("sensitive words dectection: " + fmt.Sprintf("%v", result1))
	// result2, err := contentAnalysisHandler.ContentClassification(ctx, text)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("content classification: " + fmt.Sprintf("%v", result2))

	// }
	text := `天氣風險公司天氣分析師吳聖宇在「天氣職人-吳聖宇」臉書粉專表示，海峽北部看起來有颮線接近中，稍晚中彰以北應該會有一波降雨、強風，伴隨雷電機會，先提醒大家注意。

		網友說，「春雨即將陸續抵達戰場」、「真羨慕！拜託往南壓下來些」、「再往南一點點就更好了」、「希望不要接近陸地時又減弱了」、「桃園已經刮風下雨啦，可是我還得出門」、「板橋下了一下下又停了」、「台中怎麼還沒到呢？傷心」。
		
		吳聖宇說，降雨從今天下半天開始就會逐漸趨於明顯，先是從北部、東北部地區開始，到了晚間下雨的範圍就會擴大到中部、花東地區，南部地區則可能要等到深夜之後到明天才會有局部降雨出現。`
	contentAnalysisHandler.SetUserLanguage("中文")
	summary, err := contentAnalysisHandler.ContentSummarization(ctx, text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(summary)
}
