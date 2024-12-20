# mcpmock

CLI tool that generates a mock [Model Context Protocol](https://modelcontextprotocol.io/) server from list of cases in a YAML file.

## Installation

### From npm

```bash
npm install -g @strowk/mcpmock
```

### Github Releases

Download prebulit binaries from the [releases](https://github.com/strowk/mcpmock/releases) page and put in your PATH

### From source

```bash
go get github.com/strowk/mcpmock
```

## Usage

For example you define something like this in a YAML file:

```yaml
case: List tools

# requesting list of tools
in: {"jsonrpc": "2.0", "method": "tools/list", "id": 1}

# expect one tool in the list
out: {"jsonrpc": "2.0", "result":{ "tools": [{"description": "Lists files in the current directory", "inputSchema": {"type": "object"}, "name": "list-current-dir-files"}] }, "id": 1}

---

case: Call current dir files tool

# requesting list of files in the current directory
in: {"jsonrpc": "2.0", "method": "tools/call", "params": {"name": "list-current-dir-files", "arguments": {}}, "id": 1}

# expect list of files returned as text content without error
out: {
  "jsonrpc": "2.0", 
  "id": 1,
  "result": {
    "content": [
      {"type": "text", "text": ".gitignore"},
      {"type": "text", "text": "README.md"},
      {"type": "text", "text": "go.mod"},
      {"type": "text", "text": "go.sum"},
      {"type": "text", "text": "main.go"},
      {"type": "text", "text": "main_test.go"},
      {"type": "text", "text": "testdata"},
    ],
    "isError": false
  }, 
}
```

Then if you put it in the folder `testdata` and run mcpmock like this:

```bash
mcpmock testdata
```

It would start a mock MCP server with stdio transport that would serve the cases defined in the YAML file.

If you now copy and paste this into your terminal:

```json
{"jsonrpc": "2.0", "method": "tools/list", "id": 1}
```
, you should see list of tools as is defined in first MCP case.

And sending this:

```json
{"jsonrpc": "2.0", "method": "tools/call", "params": {"name": "list-current-dir-files", "arguments": {}}, "id": 1}
```

, you should get the output as defined in second case.
