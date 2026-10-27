# Aguasto Go SDK and CLI

Official standalone client for https://www.aguasto.com/developers. Go 1.22 or newer; standard library only. Source: https://github.com/Altssdata/aguasto-agents.

```sh
go get github.com/Altssdata/aguasto-agents/go@v1.0.0
go install github.com/Altssdata/aguasto-agents/go/cmd/aguasto-research@v1.0.0
aguasto-research --help
aguasto-research fixtures
```

```go
client, err := aguasto.NewClient("") // default https://www.aguasto.com
if err != nil { return err }
result, err := client.Calculate(ctx, "weight_memory", map[string]float64{"parameters_billions": 70, "bits": 4})
```

Import `aguasto "github.com/Altssdata/aguasto-agents/go"`. The client provides Catalog, Search and Calculate; requests accept context cancellation, time out after 15 seconds, reject redirects, send no credentials/cookies and bound responses to 1 MiB. HTTP and NLWeb application errors are returned as APIError. No automatic retries occur. HTTP loopback origins are allowed for local integration tests.

Public research returns published descriptions and canonical references, not generated answers. Calculations are educational estimates, not prices or fit guarantees. No GPU provisioning, inference, booking, account or payment operations exist. Do not send confidential inputs. The fixtures command is offline; it is not a GPU sandbox.

Run `go test ./...` and `go vet ./...`. Set `AGUASTO_LIVE_TEST=1` to also verify anonymous production search/catalog and all three calculators.

All rights reserved. Redistribution or registry publication requires Aguasto approval. No open-source license is implied.
