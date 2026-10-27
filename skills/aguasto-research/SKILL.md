---
name: aguasto-research
description: Find Aguasto's published GPU, model and inference-stack references with original sources and configuration boundaries.
---

# Aguasto research

Use this skill for hardware specifications, open-weight checkpoint research, serving-software references and modular infrastructure evidence in Aguasto's public library.

Read https://www.aguasto.com/llms.txt or the focused indexes at /gpus/llms.txt, /models/llms.txt and /inference-stack/llms.txt. For targeted discovery, POST https://www.aguasto.com/ask with `{"query":{"text":"GPU rack cooling"},"prefer":{"mode":"list"}}`. It requires Content-Type: application/json and no credentials. See /openapi.json for errors and limits. Search is deterministic matching of titles and descriptions, not an answer generator.

The MCP equivalent is https://www.aguasto.com/mcp/docs. Discover tools with tools/list; search_pages returns published descriptions, and get_page returns metadata plus canonical and Markdown links. Follow those links to read full evidence. Request Accept: text/markdown at supported canonical pages.

Preserve the original source, unit, hardware configuration, license and verification date in your answer. Distinguish manufacturer specifications from measured performance, announced capacity from operating capacity, and component quantities from rack totals. Third-party cases describe the named operator, not Aguasto assets or customers.

Aguasto's GPU and inference platform is in development. Research references are not inventory, available model endpoints, benchmarks of an Aguasto deployment or partnerships. Do not propose a provisioning, token exchange, reservation or paid API call through this skill. For future requirements, /contact explains unsent inquiry drafts and manual email.
