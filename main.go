package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func GetResponse(client gpt3.Client, ctx context.Context, question string) {

    err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
        Prompt: []string{
            question,
        },
        MaxTokens: gpt3.IntPtr(3000),
        Temperature: gpt3.Float32Ptr(0.5),
    }, func(resp *gpt3.CompletionResponse) {
        fmt.Print(resp.Choices[0].Text)
    })
    if err != nil {
        fmt.Print(err)
        os.Exit(13)
    }
    fmt.Printf("\n")
}

func main() {
  
    viper.SetConfigFile(".env")
    viper.ReadInConfig()
    apikey := viper.GetString("API_KEY")

    if apikey == "" {
        panic("Missing API KEY")
    }

    ctx := context.Background()
    
    //create a client using the API key.
    client := gpt3.NewClient(apikey)
    
    rootCmd := &cobra.Command{
        Use:   "chatgpt",
        Short: "Chat with ChatGPT in console",
        Run: func(cmd *cobra.Command, args []string) {
            
            scanner := bufio.NewScanner(os.Stdin)
           
            quit := false

            for !quit {
                fmt.Print("Say something (`quit` to end): ")
                if !scanner.Scan() {
                    break
                }
                question := scanner.Text()
                switch question {
                case "quit":
                    quit = true
                default:
                    GetResponse(client, ctx, question)
                }
                // Add a loading spinner while we wait for the response from GPT-3
                fmt.Print("\033[?25l") // Hide cursor
                for i := 0; i < 10; i++ {
                    fmt.Print(".")
                    time.Sleep(100 * time.Millisecond)
                }
                fmt.Print("\033[?25h") // Show cursor
            }
            fmt.Print("Goodbye!\n")
        },
    }
    rootCmd.Execute()
}


