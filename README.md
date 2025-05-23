# Translator CLI

Translator CLI is a command-line tool designed to translate text using a remote LLM API. The application communicates with the API to provide translations in the desired language.

## Features

- Accepts text input from command-line arguments or standard input (stdin).
- Sends translation requests to a remote API.
- Outputs the translated text directly to the console.
- Supports multiple languages for translation(All supported from LLMs).
- Configurable API endpoint (`OLLAMA_HOST`) and model (`OLLAMA_MODEL`) through environment variables.

## Usage

### Prerequisites
- Go installed on your system.
- A running instance of ollama API at `http://127.0.0.1:11434/api/chat` .

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/venomega/translator
   cd translator
   go run .
   ```

### Usage

  ```bash
  $ echo -ne "Hello\n world!" | go run . - latin
  Salve,
  mundi!
  ```

  ```bash
  $ go run . "The machine is awesome" latin
  "Machina est mirabilis."
  ```

### Note

- The application uses the `OLLAMA_HOST` and `OLLAMA_MODEL` environment variables to configure the API endpoint and model.
- `OLLAMA_MODEL` defaults to `gemma2:2b` **I DO NOT RECOMEND** using this model at all.
- Testing was made with gpt-4o.
