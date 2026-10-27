export class AguastoClient {
  constructor({ baseUrl = "https://www.aguasto.com", fetch: fetcher = globalThis.fetch } = {}) {
    const url = new URL(baseUrl);
    if (url.username || url.password || url.search || url.hash || url.pathname !== "/") throw new TypeError("baseUrl must be an origin without credentials, path, query or fragment.");
    if (url.protocol !== "https:" && !(url.protocol === "http:" && ["localhost", "127.0.0.1", "[::1]"].includes(url.hostname))) throw new TypeError("Use HTTPS, or HTTP on loopback for local tests.");
    this.baseUrl = url.origin;
    this.fetcher = fetcher;
  }
  async request(path, body) {
    const response = await this.fetcher(`${this.baseUrl}${path}`, { method: body === undefined ? "GET" : "POST", headers: { Accept: "application/json", ...(body === undefined ? {} : { "Content-Type": "application/json" }) }, ...(body === undefined ? {} : { body: JSON.stringify(body) }), signal: AbortSignal.timeout(15000), redirect: "error", credentials: "omit" });
    if (!response.ok) throw new Error(`Aguasto request failed with HTTP ${response.status}.`);
    const result = await response.json();
    if (result.error) throw new Error(`${result.error.code}: ${result.error.message}`);
    return result;
  }
  search(query, { limit = 5 } = {}) {
    if (typeof query !== "string" || !query.trim() || query.length > 2000) throw new TypeError("query must contain 1–2000 characters.");
    if (!Number.isInteger(limit) || limit < 1 || limit > 20) throw new RangeError("limit must be from 1 to 20.");
    return this.request("/ask", { query: { text: query }, prefer: { mode: "list", limit } });
  }
  catalog() { return this.request("/api/v1/catalog"); }
  calculate(calculation, parameters) { return this.request("/api/v1/calculate", { calculation, parameters }); }
}
