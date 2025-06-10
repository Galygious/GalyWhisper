# GalyWhisper

A privacy-first desktop dictation application for Windows that provides real-time voice-to-text transcription with 100% local execution.

## 🌟 Features

- **100% Local Execution**: All processing happens on your machine - no data leaves your computer
- **Real-time Transcription**: Convert speech to text using Whisper models
- **Hotkey Activation**: Hold down a configurable key combination to start/stop recording
- **Seamless Text Injection**: Automatically inserts transcribed text into any active input field
- **Privacy Focused**: No cloud services, no data storage, no internet connection required

## 🚀 Getting Started

### Prerequisites

- Windows 10 or later
- [Go](https://golang.org/dl/) 1.21 or later
- [Whisper.cpp](https://github.com/ggerganov/whisper.cpp) (included as a submodule)
- [PortAudio](https://www.portaudio.com/) for audio capture

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/GalyWhisper.git
cd GalyWhisper
git submodule update --init --recursive
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o GalyWhisper.exe
```

### Usage

1. Run `GalyWhisper.exe`
2. The application will start in the system tray
3. Hold down `Ctrl+Shift` (default) to start recording
4. Speak clearly into your microphone
5. Release the keys to stop recording and insert the transcribed text

## ⚙️ Configuration

Configuration can be accessed through the system tray icon:

- Hotkey customization
- Microphone selection
- Whisper model selection
- Custom punctuation rules
- Command recognition settings

## 🔧 Development

### Project Structure

```
GalyWhisper/
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── audio/            # Audio capture and processing
│   ├── whisper/          # Whisper model integration
│   ├── keyboard/         # Keyboard input simulation
│   └── ui/               # User interface components
├── pkg/                  # Public libraries
├── config/               # Configuration files
└── models/              # Whisper model files
```

### Building from Source

```bash
# Build for Windows
go build -o GalyWhisper.exe ./cmd/galywhisper

# Build with specific Whisper model
go build -tags whisper_base_en -o GalyWhisper.exe ./cmd/galywhisper
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Whisper.cpp](https://github.com/ggerganov/whisper.cpp) for the Whisper model implementation
- [PortAudio](https://www.portaudio.com/) for audio capture
- [Wails](https://wails.io/) for the desktop application framework 