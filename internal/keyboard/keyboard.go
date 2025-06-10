package keyboard

import "fmt"

// Type injects the given text into the active window.
// This placeholder implementation simply prints the text.
// On Windows, use the SendInput API to simulate keyboard events.
func Type(text string) {
	fmt.Printf("[keyboard] %s\n", text)
}
