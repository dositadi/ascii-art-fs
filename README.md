# ASCII Art FS

Turn any text into large, decorative ASCII art — right in your terminal.

```
 _   _        _  _          
| | | |  ___ | || |  ___    
| |_| | / _ \| || | / _ \   
|  _  ||  __/| || || (_) |  
|_| |_| \___||_||_| \___/   
```

---

## What Is This?

**ASCII Art FS** is a command-line tool written in Go that converts text you type into big, stylized letters made up of ordinary keyboard characters. It's the digital equivalent of those chunky bubble-letter signs — except yours prints instantly in a terminal window.

You give it a word or sentence, pick a visual style (called a "banner"), and it draws your text in that style, as large block letters composed of characters like `|`, `_`, `/`, `\`, and spaces.

---

## What It Can Do

- **Convert any standard text to ASCII art** — letters, numbers, punctuation, and symbols (anything on a standard keyboard).
- **Render multi-line output** — use `\n` in your input to split text across multiple lines, just like pressing Enter.
- **Choose from three visual styles (banners):**
  - `standard` — clean, solid block letters
  - `shadow` — letters with a drop-shadow effect for a 3D look
  - `tinkertoy` — a lighter, more open style reminiscent of construction toys
- **Helpful error messages** — if you type something wrong (an unsupported character, an invalid font name, or too many/too few arguments), the program tells you exactly what went wrong and how to fix it.

---

## How to Run It

### What You Need First

- [Go](https://go.dev/dl/) version 1.22.2 or later installed on your computer.
- A terminal (Command Prompt on Windows, Terminal on Mac/Linux).

### Running the Program

Open your terminal, navigate to the project folder, and run:

```
go run . "your text here" fontname
```

**Examples:**

```bash
# Basic usage with the standard font
go run . "Hello World" standard

# Use the shadow font
go run . "Hello" shadow

# Use the tinkertoy font
go run . "Go" tinkertoy

# Multi-line output (use \n to add line breaks)
go run . "Hello\nWorld" standard

# Multi-line with blank line in between
go run . "Hello\n\nWorld" shadow
```

> **Note:** If you leave out the font name, the program defaults to `standard` automatically.

---

## The Three Fonts

| Font Name   | Description                                      | Best For                         |
|-------------|--------------------------------------------------|----------------------------------|
| `standard`  | Solid, classic block letters                     | General use, easy to read        |
| `shadow`    | Letters with a shadow effect, giving depth       | Stylish headers, dramatic effect |
| `tinkertoy` | Open, skeletal letters with a playful feel       | Casual or creative output        |

---

## Supported Characters

The tool supports all printable standard keyboard characters — that's:

- **Letters:** A–Z and a–z
- **Numbers:** 0–9
- **Punctuation & symbols:** ` ! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~`
- **Spaces**

Characters outside this range (such as accented letters like `é`, emoji, or other special symbols) are not supported and will produce an error message.

---

## Error Messages

The program will let you know if something goes wrong:

| Situation                              | What You'll See                                      |
|----------------------------------------|------------------------------------------------------|
| You run it without any arguments       | "You did not enter the text and your choice banner"  |
| You provide more than 2 arguments      | "Only two inputs are required (the `<text>` and `<banner>`)" |
| You type an unsupported character      | "Found an invalid character. Unable to transform invalid character" |
| You use a font name that doesn't exist | "The font you gave is not supported… Supported fonts => Standard, Shadow & Tinkertoy" |
| A font file is missing from disk       | "File read error." with the specific file error      |

---

## Project Structure

Here's how the project is organized, explained for non-technical readers:

```
ascii-art-fs/
│
├── main.go                        # The starting point — kicks everything off
│
├── go.mod                         # Go project configuration file
│
├── fonts/                         # The "alphabet libraries" for each style
│   ├── standard.txt               # Letter templates for the standard font
│   ├── shadow.txt                 # Letter templates for the shadow font
│   └── tinkertoy.txt              # Letter templates for the tinkertoy font
│
├── internal/
│   ├── config/
│   │   └── app.go                 # Reads your command-line input and runs the app
│   │
│   └── ascii/
│       ├── ascii.go               # Defines the core data structure (text + font)
│       ├── splitter.go            # Splits your input into lines (handles \n)
│       ├── readfile.go            # Looks up each character's art from the font file
│       ├── transform.go           # Assembles each character's art into full words/lines
│       └── print_ascii.go        # Prints the final art to your terminal
│
└── pkg/
    └── model/
        └── errors.go              # Defines the error format used across the app
```

### How It Works (Plain English)

1. **You type a command** like `go run . "Hello" shadow`.
2. The app **reads your input** and checks you've provided the right number of arguments.
3. Your text is **split into lines** wherever you used `\n`.
4. For each letter in your text, the app **looks up that letter's artwork** from the chosen font file. Each character is stored as 8 rows of text in the font file.
5. All the character artworks for a line are **assembled side-by-side**.
6. The result is **printed to your terminal**, line by line.

---

## Running the Tests

The project includes automated tests to verify that core functions work correctly. To run them:

```bash
go test ./...
```

The tests cover:
- **Splitter** — checks that `\n` in input correctly creates separate lines
- **ReadFont** — verifies that characters are correctly read from the font files
- **Transform** — confirms that multi-word and multi-line inputs are assembled with the right structure

---

## Known Limitations

- Only **standard ASCII characters** (the printable ones on a typical keyboard) are supported. No emoji, accented characters, or non-English letters.
- The program **must be run from the project's root folder**, because it looks for font files using relative paths (`fonts/shadow.txt`, etc.). Running it from a different directory will cause a file-not-found error.
- There is currently **no option to save the output** to a file directly — but you can use your terminal's redirect feature: `go run . "Hello" standard > output.txt`

---

## Quick Reference

```bash
# Syntax
go run . "<text>" <font>

# Fonts available
standard | shadow | tinkertoy

# Multi-line (use literal \n in your text)
go run . "Line1\nLine2" standard

# Default font (standard is used if omitted)
go run . "Hello"
```