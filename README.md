# Fast Folder CLI (`ff`)

A simple CLI tool to quickly scaffold React component or page folders with TypeScript or JavaScript files.

## Installation

### One-step install (recommended)

```sh
go install github.com/joebasset/Fast-Folders/cmd/ff@latest
```

This will place the `ff` binary in your Go bin directory (usually `$HOME/go/bin`).  
Make sure this directory is in your `$PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Manual build

1. **Clone the repository:**

   ```sh
   git clone https://github.com/joebasset/Fast-Folders.git
   cd Fast-Folders/cmd/ff
   ```

2. **Build the binary:**

   ```sh
   go build -o ff
   ```

3. **(Optional) Move the binary to your PATH:**
   ```sh
   mv ff /usr/local/bin/
   ```

## Usage

```
ff [--js=<true|false>] [--title=<true|false>] [page|comp] [Name]
```

- `--js`  
  Use `.js` files instead of `.ts` files. Default: `false` (TypeScript).
- `--title`  
  Title-case the name (capitalize first letter). Default: `true`.
- `page|comp`  
  Specify whether to create a page or a component.
- `Name`  
  The name of the page or component.

### Examples

**Create a TypeScript component (default):**

```sh
ff comp Button
```

**Create a JavaScript page, without title-casing:**

```sh
ff --js=true --title=false page home
```

**Create a TypeScript component, flags before positional arguments:**

```sh
ff --title=false comp card
```

> **Note:** Flags must be placed before positional arguments.

## Output

The tool will create a folder under `./src/Components/` or `./src/Pages/` (or lowercase variants), containing:

- `styles.css`
- `[Name].tsx` or `[Name].jsx`
- `index.ts` or `index.js`

## Troubleshooting

- If a file or directory already exists, the tool will exit with an error.
- Make sure you have Go installed (`brew install go` on
