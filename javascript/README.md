# Aguasto research client

Official source download from https://www.aguasto.com/developers and standalone integration source at https://github.com/Altssdata/aguasto-agents. This package is not published to npm. The private corporate source repository is not a public integration dependency.

Install the downloaded archive with `npm install ./aguasto-research-js.tgz`, then use:

```js
import { AguastoClient } from 'aguasto-research';
const client = new AguastoClient();
const references = await client.search('GPU rack cooling');
const estimate = await client.calculate('weight_memory', { parameters_billions: 70, bits: 4 });
```

After local installation, run `npx --no-install aguasto-research --help` or `node node_modules/aguasto-research/cli.mjs --help`. Use `fixtures` to print offline sample responses with no network request. For localhost testing, pass `--base-url http://localhost:3100` before the command. The CLI has search, catalog and calculate commands; there are no account, provisioning, billing or inference commands.

Public access is anonymous. Search returns published descriptions and canonical references; follow the sources for full evidence. Calculations are educational estimates, not service prices or hardware-fit guarantees. Requests time out after 15 seconds and do not follow redirects or send cookies. HTTP and NLWeb errors are raised; no automatic retries occur. Do not send confidential inputs. See https://www.aguasto.com/auth.md and https://www.aguasto.com/openapi.json.

All rights reserved. Redistribution or registry publication requires Aguasto approval.
