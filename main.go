package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/dillondesilva/hank/internal/ai"
	"github.com/dillondesilva/hank/internal/render"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: no command provided")
		os.Exit(1)
	}

	// We need to do the following:
	// 1. Set up a watchdog on the stderr pipe of the user's command
	// 2. Run the user's command
	// 3. If the command exits with a non-zero exit code, pass the stderr output
	// to our LLM for an explanation of the error

	// Watchdog:
	originalCmd := os.Args[1] + " " + strings.Join(os.Args[2:], " ")
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)
	err := cmd.Run()

	stdOut := stdout.String()
	stdErr := stderr.String()

	if stdErr != "" {
		fmt.Fprintf(os.Stderr, "Detected output to stderr. Running hank to help you understand the error...\n\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error state of original command: %v\n", err)
		}

		// Build the prompt from the command and error
		prompt := ai.BuildPromptFromCommandAndError(originalCmd, stdErr, stdOut)

		// Call the chat completions endpoint for a summary
		summary, err := ai.CallLMStudioLLM([]ai.Message{
			ai.Message{
				Role:    "user",
				Content: prompt,
			},
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling LMStudio LLM: %v\nPlease ensure your LMStudio server is running and reachable at http://127.0.0.1:1234", err)
			os.Exit(1)
		}

		render.RenderText(summary)
	}
}
