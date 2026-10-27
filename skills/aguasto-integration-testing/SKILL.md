---
name: aguasto-integration-testing
description: Test an Aguasto API integration using synthetic fixtures, expected errors and offline SDK fixtures before using public research data.
---

# Test an Aguasto integration

Use this skill to verify request parsing and result handling, including an empty match and an invalid input. It does not test GPU provisioning or inference: those services are planned. Use the current [OpenAPI contract](https://www.aguasto.com/openapi.json) and [sandbox guide](https://www.aguasto.com/sandbox) for the supported schema.

The anonymous endpoint is `https://www.aguasto.com/api/v1/sandbox`. No key, token, account or payment is needed. Do not send private inputs. The application does not save requests; hosting providers may process technical request metadata. The sandbox uses fixed synthetic references, not production research or customer records.

Start with GET. Expect HTTP 200, `environment: "sandbox"`, `synthetic: true`, `persistence: "none"`, operations `search` and `calculate`, and fixture IDs `model-memory` and `rack-cooling`.

For POST, send `Content-Type: application/json`. Choose the cases relevant to the integration:

| Request body | Expected result |
| --- | --- |
| `{"operation":"search","query":"rack cooling"}` | HTTP 200; `result.results` includes fixture `rack-cooling` |
| `{"operation":"search","query":"unmatched-example"}` | HTTP 200; `result.results` is empty, rather than an error or an invented reference |
| `{"operation":"calculate","calculation":"weight_memory","parameters":{"parameters_billions":70,"bits":4}}` | HTTP 200; `result.result.gb` is 35, before runtime overhead |
| `{"operation":"unsupported"}` | HTTP 400; `error.code` is `INVALID_ARGUMENT`, with a useful `error.message` |

Successful POST responses carry the same sandbox, synthetic and persistence markers. Read structured fields rather than matching prose. A calculation result is an educational estimate, not measured performance, hardware fit or an Aguasto price quote.

The body limit is 16 KiB and the search-query limit is 2,000 characters. An oversized body returns 413; an unsupported Content-Type returns 415; a disallowed browser Origin returns 403. The shared calculator admission window allows 30 requests per minute per warm server instance. Honor `Retry-After` on 429 and stop the test if the service remains unavailable; do not run a load test or retry indefinitely.

For local-only SDK checks, use the published client's `aguasto-research fixtures` command after installing the official package from the developer guide. This exercises bundled fixtures without sending a request. A successful fixture run does not prove network connectivity or ChatGPT-host compatibility.

Report which cases ran, their actual status and structured results, and any failure. Do not describe synthetic records as available GPUs or treat a sandbox pass as a production service guarantee. Move to public research endpoints only when the user's task requires real references.
