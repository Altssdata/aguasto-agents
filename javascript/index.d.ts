export type Calculation = "weight_memory" | "facility_energy" | "token_cost";
export type PageReference = { name: string; description: string; url: string; dateModified: string; encoding?: { contentUrl: string; encodingFormat: "text/markdown" } };
export type Parameters = { parameters_billions: number; bits: 4 | 8 | 16 | 32 } | { it_kw: number; pue: number; hours: number; usd_per_kwh: number } | { cost_usd: number; completed_tokens: number };
export class AguastoClient {
  constructor(options?: { baseUrl?: string; fetch?: typeof fetch });
  search(query: string, options?: { limit?: number }): Promise<{ _meta: { response_type: "answer"; version: "0.55" }; results: PageReference[] }>;
  catalog(): Promise<{ limitations: string; pages: PageReference[] }>;
  calculate(calculation: Calculation, parameters: Parameters): Promise<{ calculation: Calculation; inputs: Parameters; result: Record<string, number>; limitations: string }>;
}
