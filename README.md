# 😂 GoScraper: The "I-can't-believe-it's-not-human" Web Scraper 🚀

### What is this?
Meet **GoScraper** – the ultimate bot with no chill! 🤖 This tool not only scrapes a website's soul (a.k.a., its main content) but also uses **Google's Gemini API** to translate that garbled HTML mess into human-friendly lingo! If you're tired of squinting at tags and nested divs, GoScraper's got your back. It’s like magic, but with more Go code. ✨

---

## 💡 Why You Need This
Because sometimes you need the internet to just… make sense! 🧐

- **Step 1**: Point GoScraper at a website.
- **Step 2**: It scrapes.
- **Step 3**: The content is magically transformed into something you can read without a PhD in HTML hieroglyphics.

---

## 🔧 Installation

1. Clone the repo like it’s hot:
   ```bash
   git clone https://github.com/anima-regem/GoScraper.git
   cd GoScraper
   ```

2. Install dependencies:
   ```bash
   go get -u github.com/google/generative-ai-go/genai
   go get -u github.com/gocolly/colly
   go get -u github.com/joho/godotenv
   ```

---

## ⚙️ Setting Up Your `.env` File

You'll need your secret API key to use the magic of AI 🤖🔮. Don’t worry; it’s just one tiny step!

1. Make a file called `.env` in your project directory:
   ```bash
   touch .env
   ```

2. Inside `.env`, add your Gemini API key:
   ```plaintext
   GEMINI_API_KEY=YOUR_SECRET_KEY
   ```

   > 💡 **Pro Tip**: Don’t commit your `.env` file to Git, unless you like exposing secrets and living dangerously!

---

## 🏗 Usage Examples

Fire up GoScraper like this:

```bash
go run main.go <URL>
```

- **Example**: Scrape and beautify "https://example.com" 🌐
  ```bash
  go run main.go https://example.com
  ```

- **The Result**: *Voilà!* Human-readable content, courtesy of GoScraper and Gemini! 🌌

### Want a Binary? 🛠
Too lazy to run Go every time? We got you. Just download the latest release and run it like a boss:
```bash
./goscraper <URL>
```

---

## 😱 What Could Go Wrong?

1. **URL not starting with `https://`** – don’t worry; GoScraper will add it for you because it’s thoughtful like that. 💅
2. **Error loading `.env` file** – make sure your `.env` file is sitting right in the project folder, like a good dog. 🐶
3. **API Key Issues** – if GoScraper screams, make sure you haven’t pasted your WiFi password by accident. 😉

---

## 🎉 Why We Love GoScraper

- It's fun, fast, and *mostly* accurate! 😅
- Leaves your screen spotless (no HTML mess).
- Turns robot gibberish into readable content like a dream! 💤

---

Happy Scraping! Go forth and make the internet readable! 🌎
