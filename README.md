# CHATBOT-GO

A chatbot built using Go, utilizing OpenAI GPT-3 API and the Viper and Cobra libraries. This chatbot can answer a variety of questions and carry out basic conversations with users.

## Features:
- Utilizes OpenAI GPT-3 API for language processing and understanding
- Environment variables managed using Viper
- Command Line Interface created using Cobra

## Prerequisites:
- Go programming language
- OpenAI API Key

## Installation:
1. Install dependencies: 
    ```
    $ go get -u github.com/spf13/viper
    $ go get -u github.com/spf13/cobra
    ```
2. Set your OpenAI API Key as an environment variable: 
    ```
    $ export OPENAI_API_KEY=<YOUR_API_KEY>
    ```
3. Run the chatbot:
    ```
    $ go run main.go
    ```

## Contributing

If you are interested in contributing to this project, feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
