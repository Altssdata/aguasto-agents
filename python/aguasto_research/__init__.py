"""Anonymous Aguasto public research client. No GPU provisioning or inference."""
import json
from urllib.parse import urlsplit
from urllib.request import Request, build_opener, HTTPRedirectHandler
from urllib.error import HTTPError, URLError


class _NoRedirect(HTTPRedirectHandler):
    def redirect_request(self, req, fp, code, msg, headers, newurl):
        return None


class AguastoClient:
    def __init__(self, base_url="https://www.aguasto.com"):
        url = urlsplit(base_url)
        if url.username or url.password or url.query or url.fragment or url.path not in ("", "/") or not url.hostname:
            raise ValueError("base_url must be an origin without credentials, path, query or fragment")
        if url.scheme != "https" and not (url.scheme == "http" and url.hostname in ("localhost", "127.0.0.1", "::1")):
            raise ValueError("Use HTTPS, or HTTP on loopback for local tests")
        self.base_url = base_url.rstrip("/")
        self._opener = build_opener(_NoRedirect())

    def _request(self, path, body=None):
        data = None if body is None else json.dumps(body, allow_nan=False).encode("utf-8")
        headers = {"Accept": "application/json"}
        if data is not None:
            headers["Content-Type"] = "application/json"
        request = Request(self.base_url + path, data=data, headers=headers)
        try:
            with self._opener.open(request, timeout=15) as response:
                result = json.load(response)
        except HTTPError as error:
            raise RuntimeError(f"Aguasto request failed with HTTP {error.code}") from None
        except URLError:
            raise RuntimeError("Aguasto request could not complete") from None
        if "error" in result:
            raise RuntimeError(f"{result['error']['code']}: {result['error'].get('message', '')}")
        return result

    def search(self, query, limit=5):
        if not isinstance(query, str) or not query.strip() or len(query) > 2000:
            raise ValueError("query must contain 1–2000 characters")
        if isinstance(limit, bool) or not isinstance(limit, int) or not 1 <= limit <= 20:
            raise ValueError("limit must be from 1 to 20")
        return self._request("/ask", {"query": {"text": query}, "prefer": {"mode": "list", "limit": limit}})

    def catalog(self):
        return self._request("/api/v1/catalog")

    def calculate(self, calculation, parameters):
        return self._request("/api/v1/calculate", {"calculation": calculation, "parameters": parameters})
