# content-moderation-api
Content Moderation API, powered by OpenAI, assists developers in detecting and filtering inappropriate content, such as profanity, adult content, violence, and moral violations. This efficient, easy-to-integrate solution supports various use cases like social platforms, forums, and chat apps, ensuring a safe user experience.

## Installation

```bash
go get github.com/Zhima-Mochi/content-moderation-api
```

## Usage

### Sensitivity Detection

```go
ctx := context.Background()
texts := []string{
    "我不要听话",
	"you are so stupid",
	"每天早起做愛心便當",
    "この野郎！",
}
for _, text := range texts {
    fmt.Println(text)
    result, err := contentModerationHandler.SensitiveWordsDetection(ctx, text)
    if err != nil {
        log.Fatal(err)
        return
    }
    fmt.Println(result)
}
```
output:
```bash
我不要听话
false
you are so stupid
true
每天早起做愛心便當
false
この野郎！
true
```