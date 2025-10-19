package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"slfy/internal/ai"
	"slfy/internal/render"
	"strings"
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
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting command: %v\n", err)
		stdOut := stdout.String()
		stdErr := stderr.String()

		fmt.Fprintf(os.Stdout, "%s\n", stdOut)
		fmt.Fprintf(os.Stderr, "%s\n", stdErr)

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
			fmt.Fprintf(os.Stderr, "Error calling LMStudio LLM: %v\n", err)
			os.Exit(1)
		}

		render.RenderText(summary)
		os.Exit(1)
	}
}
