#!/usr/bin/env node
import { readFile } from "node:fs/promises";
import { AguastoClient } from "./index.mjs";
const args = process.argv.slice(2);
const help = "Usage: aguasto-research [--base-url ORIGIN] search QUERY | catalog | calculate NAME PARAMETERS_JSON | fixtures\nPublic research and educational estimates only. No GPU provisioning or inference.\nfixtures prints offline sample responses without making a network request.";
try {
  let baseUrl;
  if (args[0] === "--base-url") { args.shift(); baseUrl = args.shift(); if (!baseUrl) throw new Error("--base-url requires an origin."); }
  const command = args.shift();
  if (!command || command === "--help" || command === "-h") console.log(help);
  else if (command === "fixtures" && !args.length) console.log(await readFile(new URL("./fixtures.json", import.meta.url), "utf8"));
  else {
    const client = new AguastoClient({ baseUrl });
    const result = command === "catalog" && !args.length ? await client.catalog()
      : command === "search" && args.length ? await client.search(args.join(" "))
        : command === "calculate" && args.length === 2 ? await client.calculate(args[0], JSON.parse(args[1]))
          : (() => { throw new Error(help); })();
    console.log(JSON.stringify(result, null, 2));
  }
} catch (error) { console.error(error instanceof Error ? error.message : "Request failed."); process.exitCode = 1; }
