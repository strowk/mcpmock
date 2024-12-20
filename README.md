# mcpmock

CLI tool that generates a mock [Model Context Protocol](https://modelcontextprotocol.io/) server from list of cases in a YAML file.

## Installation

From npm: `npm install -g @strowk/mcpmock`

From Github Releases: [Download](https://github.com/strowk/mcpmock/releases), unpack and put binary in your PATH

From sources:

```bash
go get github.com/strowk/mcpmock
go install github.com/strowk/mcpmock
```

## Usage

For example you define something like this in a YAML file:

```yaml
case: List tools

# requesting list of tools
in: {"jsonrpc": "2.0", "method": "tools/list", "id": 1}

# expect one tool in the list
out: {"jsonrpc": "2.0", "result":{ "tools": [{"description": "Hello MCP", "inputSchema": {"type": "object"}, "name": "hello"}] }, "id": 1}

---

case: Call current dir files tool

# requesting list of files in the current directory
in: {"jsonrpc": "2.0", "method": "tools/call", "params": {"name": "hello", "arguments": {}}, "id": 1}

# expect list of files returned as text content without error
out: {
  "jsonrpc": "2.0", 
  "id": 1,
  "result": {
    "content": [
      {"type": "text", "text": "Hi!"}
    ],
    "isError": false
  }, 
}
```

Then if you put it in the folder `testdata` and run mcpmock like this:

```bash
mcpmock serve testdata
```

It would start a mock MCP server with stdio transport that would serve the cases defined in the YAML file.

If you now copy and paste this into your terminal:

```json
{"jsonrpc": "2.0", "method": "tools/list", "id": 1}
```
, you should see list of tools as is defined in first MCP case.

And sending this:

```json
{"jsonrpc": "2.0", "method": "tools/call", "params": {"name": "hello", "arguments": {}}, "id": 1}
```

, you should get the output as defined in second case.
