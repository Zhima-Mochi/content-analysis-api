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
    "哇操",
    "我不要乖乖聽話",
    "每天早起做愛心便當",
}
for _, text := range texts {
    log.Println(text)
    result, err := contentModerationHandler.SensitiveWordsDetection(ctx, text)
    if err != nil {
        log.Fatal(err)
        return
    }
    log.Println(result)
}
```
output:
```bash
2023/03/21 19:43:20 哇操
2023/03/21 19:43:21 true
2023/03/21 19:43:21 我不要乖乖聽話
2023/03/21 19:43:22 false
2023/03/21 19:43:22 每天早起做愛心便當
2023/03/21 19:43:23 false
```