<h1 align="center">
   mcpmock
</h1>

<h4 align="center">CLI tool that generates a mock <a href="https://modelcontextprotocol.io">Model Context Protocol</a> server from list of cases in a YAML file.</h4>

<p align="center">
	<a href="https://github.com/strowk/mcpmock/actions/workflows/test.yaml"><img src="https://github.com/strowk/mcpmock/actions/workflows/test.yaml/badge.svg"/></a>
	<a href="https://github.com/strowk/mcpmock/actions/workflows/golangci-lint.yaml"><img src="https://github.com/strowk/mcpmock/actions/workflows/golangci-lint.yaml/badge.svg"/></a>
</p>

## Installation

From npm: `npm install -g @strowk/mcpmock`

From Github Releases: [Download](https://github.com/strowk/mcpmock/releases), unpack and put binary in your PATH

From sources:

```bash
go get github.com/strowk/mcpmock
go install github.com/strowk/mcpmock
```

## Usage

For example if you define something like this in a YAML file:

```yaml
case: List tools

# requesting list of tools
in: {"jsonrpc": "2.0", "method": "tools/list", "id": 1}

# expect one tool in the list
out: {"jsonrpc": "2.0", "result":{ "tools": [{"description": "Hello MCP", "inputSchema": {"type": "object"}, "name": "hello"}] }, "id": 1}

---

case: Call Hello tool

# calling the tool
in: {"jsonrpc": "2.0", "method": "tools/call", "params": {"name": "hello", "arguments": {}}, "id": 1}

# expect "Hi!" as output
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

Then if you put it in the folder `testdata` (make sure file also ends with `_test.yaml`) and run mcpmock like this:

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

It would also take care to replace the `id` field in the response with the `id` from the request, so if you send this:

```json
{"jsonrpc": "2.0", "method": "tools/call", "params": {"name": "hello", "arguments": {}}, "id": 2}
```

, you should get the output as defined in second case, but with `id` field set to `2`.
