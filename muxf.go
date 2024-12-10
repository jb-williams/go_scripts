package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func listTmuxSessions() string {
	cmd := exec.Command("tmux", "list-sessions")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

func sessionExists(sessions, name string) bool {
	return strings.Contains(sessions, name)
}

func runTmuxCommand(args ...string) error {
	cmd := exec.Command("tmux", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	sessions := listTmuxSessions()
	ifMain := sessionExists(sessions, "main")
	ifDiary := sessionExists(sessions, "diary")
	ifAttached := sessionExists(sessions, "attached")

	switch {
	case !ifMain && !ifDiary && !ifAttached:
		err := runTmuxCommand("new", "-s", "diary", "-E", "diary", "-d", ";", "new", "-A", "-s", "main", ";", "list-sessions")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
			os.Exit(1)
		}
	case ifMain && !ifDiary && !ifAttached:
		err := runTmuxCommand("new", "-s", "diary", "-E", "diary", "-d", ";", "attach", "-t", "main", ";", "list-sessions")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
			os.Exit(1)
		}
	case ifMain && ifDiary && !ifAttached:
		err := runTmuxCommand("attach", "-t", "main", ";", "choose-tree")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
			os.Exit(1)
		}
	case !ifMain && !ifDiary && ifAttached:
		err := runTmuxCommand("new", "-s", "diary", "-E", "diary", "-d", ";", "new", "-A", "-s", "main", ";", "list-sessions")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
			os.Exit(1)
		}
	case !ifMain && ifDiary && !ifAttached:
		err := runTmuxCommand("new", "-A", "-s", "main", ";", "list-sessions")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
			os.Exit(1)
		}
	case ifMain && !ifDiary && ifAttached:
		err := runTmuxCommand("new", "-s", "diary", "-E", "diary", "-d", ";", "list-sessions")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
			os.Exit(1)
		}
	case ifMain && ifDiary && ifAttached:
		fmt.Print("New sesh(name): ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			newSession := scanner.Text()
			if newSession == "" || strings.ToLower(newSession) == "n" {
				err := runTmuxCommand("choose-tree")
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
					os.Exit(1)
				}
			} else {
				err := runTmuxCommand("new", "-s", newSession, "-d", ";", "choose-tree")
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error executing tmux command:", err)
					os.Exit(1)
				}
			}
		}
	default:
		os.Exit(2)
	}
}
