# content-moderation-api

[![Go Report Card](https://goreportcard.com/badge/github.com/Zhima-Mochi/content-moderation-api)](https://goreportcard.com/report/github.com/Zhima-Mochi/content-moderation-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Introduction
Content Moderation API, powered by OpenAI, assists developers in detecting and filtering inappropriate content, such as profanity, adult content, violence, and moral violations. This efficient, easy-to-integrate solution supports various use cases like social platforms, forums, and chat apps, ensuring a safe user experience.

## Features

- **Sensitivity Detection**: Detects sensitive words in text.
- **Content Classification**: Classifies text into categories, such as `politics`, `porn`, `ad`, `violence`, `abuse`.

## Installation

```bash
go get github.com/Zhima-Mochi/content-moderation-api
```

## Example

```go
texts := []string{
    "我不要听话",
    "you are so stupid",
    "每天早起做愛心便當",
    "この野郎！",
}

for _, text := range texts {
    fmt.Println("text: " + text)
    result, _ := contentModerationHandler.SensitiveWordsDetection(ctx, text)
    fmt.Println("sensitive words dectection: " + fmt.Sprintf("%v", result))
}
```
result:
```bash
text: 我不要听话
sensitive words dectection: false
text: you are so stupid
sensitive words dectection: true
text: 每天早起做愛心便當
sensitive words dectection: false
text: この野郎！
sensitive words dectection: true
```