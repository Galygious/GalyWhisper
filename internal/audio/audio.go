package audio

import (
	"context"
)

// Record captures audio from the default microphone until the context is done.
// This is a placeholder implementation. A real version would use PortAudio or
// WASAPI to capture PCM samples.
func Record(ctx context.Context) ([]byte, error) {
	<-ctx.Done()
	// Return dummy audio data
	return []byte{}, nil
}
