package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// compileSass starts the Sass watch command
func compileSass() (*exec.Cmd, error) {
	// Command to compile Sass
	cmd := exec.Command("sass", "--watch", "scss:static")

	// Set output to standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func main() {
	// Start the Sass compiler
	sassCmd, err := compileSass()
	if err != nil {
		log.Fatalf("Failed to start Sass compiler: %v", err)
	}

	// Start your main application
	appCmd := exec.Command("go", "run", "main.go")
	appCmd.Stdout = os.Stdout
	appCmd.Stderr = os.Stderr

	// Run your main application
	if err := appCmd.Start(); err != nil {
		log.Fatalf("Failed to start main application: %v", err)
	}

	// Handle termination signals to stop both processes gracefully
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-sigs

	// Clean up: Terminate the Sass compiler process
	if err := sassCmd.Process.Kill(); err != nil {
		log.Fatalf("Failed to kill Sass compiler process: %v", err)
	}

	log.Println("Application terminated.")
}
