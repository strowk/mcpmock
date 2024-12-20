# mcpmock

Install: `npm install -g @strowk/mcpmock`

mcpmock is a CLI tool that generates a mock [Model Context Protocol](https://modelcontextprotocol.io/) server from list of cases in a YAML file.

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
