---
name: aguasto-calculations
description: Estimate model weight memory, PUE-adjusted electricity and completed-token cost using Aguasto's educational calculators.
---

# Aguasto educational calculations

Use https://www.aguasto.com/openapi.json for the current schema, or discover tools from https://www.aguasto.com/mcp/tools using tools/list. Public calculations require no credentials and do not persist inputs or provision compute.

POST https://www.aguasto.com/api/v1/calculate with Content-Type: application/json and one of these bodies:

- Weight memory: `{"calculation":"weight_memory","parameters":{"parameters_billions":70,"bits":4}}`.
- Energy: `{"calculation":"facility_energy","parameters":{"it_kw":100,"pue":1.2,"hours":24,"usd_per_kwh":0.1}}`.
- Token cost: `{"calculation":"token_cost","parameters":{"cost_usd":10,"completed_tokens":2000000}}`.

Weight memory is parameters × bits / 8. Decimal GB and binary GiB differ. A 70B model at 4 bits uses 35 decimal GB for weights alone; KV cache, activations, quantization metadata and runtime overhead are excluded. Do not infer hardware fit from this estimate.

Facility electricity is IT kW × PUE × hours; multiply by the supplied USD/kWh tariff for illustrative cost. PUE must be at least 1. These are scenario inputs, not measurements or Aguasto tariffs. Completed-token cost is supplied cost / completed tokens × 1,000,000; the denominator must be positive and excludes failed or incomplete output unless the user explicitly defines it otherwise.

Report assumptions and units alongside the result. Missing or invalid fields return structured errors; correct the input rather than treating an error as a zero estimate. Cross-check simple arithmetic independently. The /tools page offers the same formulas interactively.

These interfaces do not execute inference, quote Aguasto prices or reserve GPU capacity. Never send confidential workloads, credentials or personal data.
