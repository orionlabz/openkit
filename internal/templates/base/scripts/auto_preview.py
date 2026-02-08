#!/usr/bin/env python3
"""\
Auto Preview - OpenKit
=====================

Start/stop/status a local development environment via Docker Compose.

Usage:
    python .opencode/scripts/auto_preview.py start
    python .opencode/scripts/auto_preview.py stop
    python .opencode/scripts/auto_preview.py status
"""

import argparse
import subprocess
import sys
from pathlib import Path

from typing import List, Optional


def get_project_root() -> Path:
    return Path(".").resolve()


def get_docker_compose_file(root: Path) -> Optional[str]:
    dev_config = root / "docker-compose.dev.yml"
    if dev_config.exists():
        return "docker-compose.dev.yml"

    std_config = root / "docker-compose.yml"
    if std_config.exists():
        return "docker-compose.yml"

    return None


def run_docker_command(cmd_args: List[str], root_path: Path) -> None:
    compose_file = get_docker_compose_file(root_path)
    if not compose_file:
        print("No docker-compose.dev.yml or docker-compose.yml found in project root.")
        raise SystemExit(1)

    base_cmd = ["docker", "compose", "-f", compose_file]
    full_cmd = base_cmd + cmd_args

    try:
        subprocess.run(full_cmd, cwd=str(root_path), check=True)
    except subprocess.CalledProcessError as e:
        print(f"Docker command failed: {e}")
        raise SystemExit(1)
    except FileNotFoundError:
        print("Docker not found. Please verify it is installed and in your PATH.")
        raise SystemExit(1)


def start_server() -> None:
    root = get_project_root()
    print("Starting development environment (detached)...")
    run_docker_command(["up", "-d"], root)
    print("Environment started.")


def stop_server() -> None:
    root = get_project_root()
    print("Stopping development environment...")
    run_docker_command(["stop"], root)
    print("Environment stopped.")


def status_server() -> None:
    root = get_project_root()
    print("Docker Compose status:")
    run_docker_command(["ps"], root)


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("action", choices=["start", "stop", "status"])
    # Keep compatibility with older callers that passed a port.
    parser.add_argument("port", nargs="?", default="3000")
    args = parser.parse_args()

    if args.action == "start":
        start_server()
    elif args.action == "stop":
        stop_server()
    elif args.action == "status":
        status_server()
    else:
        print(f"Unknown action: {args.action}")
        raise SystemExit(2)


if __name__ == "__main__":
    main()
