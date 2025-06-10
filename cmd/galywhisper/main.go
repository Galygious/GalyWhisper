package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/galywhisper/internal/audio"
	"github.com/example/galywhisper/internal/keyboard"
	"github.com/example/galywhisper/internal/whisper"
)

func main() {
	fmt.Println("GalyWhisper starting...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// placeholder hotkey: start immediately, stop on ctrl+c
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	fmt.Println("Recording... press Ctrl+C to stop")
	audioData, err := audio.Record(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Transcribing...")
	text, err := whisper.Transcribe(ctx, audioData)
	if err != nil {
		return err
	}

	fmt.Printf("Transcription: %s\n", text)
	keyboard.Type(text)

	time.Sleep(500 * time.Millisecond)
	return nil
}
