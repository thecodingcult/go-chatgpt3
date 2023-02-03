CHATBOT-GO


A chatbot built using Go, utilizing OpenAI GPT-3 API and the Viper and Cobra libraries. This chatbot can answer a variety of questions and carry out basic conversations with users.


Features:
Utilizes OpenAI GPT-3 API for language processing and understanding
Environment variables managed using Viper
Command Line Interface created using Cobra


Prerequisites:
Go programming language
OpenAI API Key

Install dependencies:
$ go get -u github.com/spf13/viper
$ go get -u github.com/spf13/cobra

Set your OpenAI API Key as an environment variable:
$ export OPENAI_API_KEY=<YOUR_API_KEY>

Run the chatbot:
$ go run main.go

Built With:
-Go
-OpenAI GPT-3 API
-Viper
-Cobra

