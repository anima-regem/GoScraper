package main

import (
  "fmt"
  "github.com/google/generative-ai-go/genai" 
  "google.golang.org/api/option"
  "os"
  "log"
  "github.com/joho/godotenv"
  "context"
  "github.com/gocolly/colly"
  "strings"
)

func getScrapedWebPage(url string) string {
  c := colly.NewCollector()
  var pageContent string
  c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting", r.URL)
  })

  c.OnHTML("body", func(e *colly.HTMLElement){
    pageContent = e.Text
  })

  c.Visit(url)
  pageContent = strings.Join(strings.Fields(pageContent), " ")
  return pageContent
}

func getHumanReadableWebContent(content string) string {
  ctx := context.Background()
  client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()
  model := client.GenerativeModel("gemini-1.5-flash")
  var prompt string = "This is the scraped web content, give me a human readable version of it. Web content: " + content
  resp, err := model.GenerateContent(ctx, genai.Text(prompt))
  if err != nil {
    log.Fatal(err)
  }
  return fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0]) //resp.Candidates[0].Content.Parts[0]
}

func correctURL(url string) string {
  if !strings.HasPrefix(url, "https://") {
    fmt.Println("URL does not start with https:// prefix, defaulting to https://")
    url = "https://" + url
  }
  return url
}


func main(){
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }


  url := os.Args[1]
  content := getScrapedWebPage(correctURL(url))
  humanReadableContent := getHumanReadableWebContent(content)
  fmt.Println(humanReadableContent)
  


}
  
