import argparse
import json
import sys
from pathlib import Path
from . import AguastoClient


def main():
    parser = argparse.ArgumentParser(description="Public research and educational estimates. No GPU provisioning or inference.")
    parser.add_argument("--base-url", default="https://www.aguasto.com")
    commands = parser.add_subparsers(dest="command", required=True)
    search = commands.add_parser("search")
    search.add_argument("query")
    search.add_argument("--limit", type=int, default=5)
    commands.add_parser("catalog")
    calculate = commands.add_parser("calculate")
    calculate.add_argument("name", choices=["weight_memory", "facility_energy", "token_cost"])
    calculate.add_argument("parameters", help="JSON object")
    commands.add_parser("fixtures", help="Print offline sample responses without a network request")
    args = parser.parse_args()
    try:
        if args.command == "fixtures":
            print(Path(__file__).with_name("fixtures.json").read_text(encoding="utf-8"))
            return
        client = AguastoClient(args.base_url)
        result = client.search(args.query, args.limit) if args.command == "search" else client.catalog() if args.command == "catalog" else client.calculate(args.name, json.loads(args.parameters))
        print(json.dumps(result, indent=2))
    except (ValueError, RuntimeError, TimeoutError) as error:
        print(str(error), file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
