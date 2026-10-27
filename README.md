# Aguasto public agent integrations

Official public research clients, command-line tools, agent skills and MCP configuration for [Aguasto](https://www.aguasto.com/developers).

This export contains only the standalone integrations already downloadable from the Aguasto developer page. The corporate website source, internal review records and Git history are excluded. All rights reserved; existing SDK redistribution terms apply. No open-source license is implied.

## Available interfaces

- Anonymous research descriptions and canonical source links: https://www.aguasto.com/mcp/docs
- Educational weight-memory, facility-energy and completed-token cost calculations: https://www.aguasto.com/mcp/tools
- HTTP documentation: https://www.aguasto.com/openapi.json
- JavaScript SDK and CLI: javascript/README.md
- Python SDK and CLI: python/README.md
- Go SDK and CLI: go/README.md
- Homebrew formula: Formula/aguasto-research.rb
- Research and calculation instructions: skills/aguasto-research/SKILL.md and skills/aguasto-calculations/SKILL.md
- Agent rules: AGENTS.md
- Plugin and MCP client configuration: plugin.json and mcp.json

No GPU provisioning, inference, capacity booking, prices, customer accounts, credentials or payment actions are provided. Research descriptions are not generated answers; follow canonical pages and primary sources. Calculations use supplied assumptions and are not fit guarantees or Aguasto pricing. Never send confidential inputs.

## Local CLI use

With Node.js 22 or newer, run `node javascript/cli.mjs --help` or `node javascript/cli.mjs fixtures`. With Python 3.10 or newer, follow python/README.md. Neither client is represented as published on npm or PyPI until actual registry publication is verified.

## Homebrew

`brew tap altssdata/aguasto-agents https://github.com/Altssdata/aguasto-agents`

`brew install altssdata/aguasto-agents/aguasto-research`

The formula installs the SHA-256-pinned JavaScript source release and its Node.js dependency. Run `aguasto-research fixtures` for an offline check.

## Go

`go install github.com/Altssdata/aguasto-agents/go/cmd/aguasto-research@v1.0.0`

The Go 1.22+ client uses only the standard library; see go/README.md for SDK usage and tests.
