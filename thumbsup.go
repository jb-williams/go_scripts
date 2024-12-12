package main

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

// getClipboardCommand determines the appropriate clipboard command based on the OS.
func getClipboardCommand() (*exec.Cmd, error) {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xclip", "-selection", "c"), nil
	case "darwin":
		return exec.Command("pbcopy"), nil
	default:
		return nil, errors.New("unsupported operating system")
	}
}

// copyToClipboard uses the appropriate clipboard command to copy the text.
func copyToClipboard(text string) error {
	// Get the clipboard command
	cmd, err := getClipboardCommand()
	if err != nil {
		return err
	}

	// Create a pipe to send text to the command's stdin
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// Start the clipboard command
	if err := cmd.Start(); err != nil {
		return err
	}

	// Write the text to the command's stdin
	_, err = stdinPipe.Write([]byte(text))
	if err != nil {
		return err
	}
	stdinPipe.Close()

	// Wait for the command to complete
	return cmd.Wait()
}

func main() {
	text := `👍`

	// Attempt to copy the text to the clipboard
	err := copyToClipboard(text)
	if err != nil {
		// Print the error and exit with a non-zero status
		os.Stderr.WriteString("Error copying to clipboard: " + err.Error() + "\n")
		os.Exit(1)
	}

	// Indicate success
	os.Stdout.WriteString("Text copied to clipboard successfully!\n")
}
