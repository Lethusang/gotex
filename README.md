# GoTex Editor

A simple text editor built as a learning project to explore Go programming and the Fyne toolkit.

## ⚠️ Learning Project Notice

This project was created primarily as a learning exercise to:
- Understand Go programming concepts
- Explore the Fyne GUI toolkit
- Practice software design patterns
- Learn about cross-platform GUI development

**Note**: This is not intended for production use and may contain bugs or incomplete features. Feel free to use it as a learning reference or starting point for your own projects!

## Features

Basic implementation of:
- 📝 Text editing
- 💾 File operations (New, Open, Save, Save As)
- ✂️ Clipboard operations (Cut, Copy, Paste)
- 🎨 Light/Dark theme switching
- 📊 Simple status bar

## Prerequisites

If you want to run or experiment with this project, you'll need:

### Go
- Go 1.16 or later ([golang.org](https://golang.org/))

### Linux Dependencies
```bash
sudo apt-get update
sudo apt-get install gcc libgl1-mesa-dev xorg-dev
sudo apt-get install libxxf86vm-dev
sudo apt-get install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev
```

## Try It Out

1. Clone the repository:
```bash
git clone https://github.com/Lethusang/gotex.git
cd gotex
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run cmd/gotex/main.go
```

## Project Structure

```
gotex/
├── cmd/gotex/          # Application entry point
├── internal/
│   ├── buffer/         # Text buffer implementation
│   └── editor/         # Editor components
└── go.mod
```

## Resources I Used

- [Go Documentation](https://golang.org/doc/)
- [Fyne Documentation](https://developer.fyne.io/)
- [Effective Go](https://golang.org/doc/effective_go)

## License

MIT License - feel free to use this code for your own learning!

## Acknowledgments

- The Go community for excellent learning resources
- Fyne developers for the GUI toolkit
- Everyone who shares their knowledge about Go

Happy coding and learning! 🚀
