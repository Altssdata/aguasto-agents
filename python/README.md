# Aguasto research client for Python

Official Aguasto client from https://www.aguasto.com/developers. Published on PyPI as [aguasto-research](https://pypi.org/project/aguasto-research/), with public source at https://github.com/Altssdata/aguasto-agents/tree/main/python.

With Python 3.10 or newer, install in your project's virtual environment with `python -m pip install aguasto-research==1.0.1`. You can also install the developer-page source archive with `python -m pip install ./aguasto-research-python.tar.gz` (a local Python packaging environment is required).

```python
from aguasto_research import AguastoClient
client = AguastoClient()
references = client.search("GPU rack cooling")
estimate = client.calculate("weight_memory", {"parameters_billions": 70, "bits": 4})
```

Use `python -m aguasto_research --help` for the CLI. Commands include search, catalog, calculate and fixtures. The fixtures command prints offline sample responses without a network request. Pass `--base-url http://localhost:3100` before the command for local integration tests.

Anonymous access requires no credentials. Search returns published descriptions and canonical links, not generated answers. Calculations use educational assumptions and are not Aguasto prices or deployment guarantees. No GPU provisioning, inference, paid compute or account operation is supported. HTTP requests time out after 15 seconds, do not follow redirects and are not retried. HTTP and NLWeb failures raise exceptions. The client has no runtime dependencies outside Python's standard library.

Do not send confidential data. See https://www.aguasto.com/auth.md and https://www.aguasto.com/openapi.json. All rights reserved; registry publication or redistribution requires Aguasto approval.
