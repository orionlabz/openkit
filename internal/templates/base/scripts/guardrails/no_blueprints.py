#!/usr/bin/env python3
"""Guardrail: prevent reintroduction of blueprint references.

This script fails if blueprint-related terms appear in OpenKit-shipped packs.

Usage:
  python .opencode/scripts/guardrails/no_blueprints.py <project_path>
"""

from __future__ import annotations

import re
import sys
from dataclasses import dataclass
from pathlib import Path


PATTERNS = [
    re.compile(r"@blueprints/"),
    re.compile(r"\bblueprints\b", re.IGNORECASE),
    re.compile(r"\bblueprint\b", re.IGNORECASE),
]

SCAN_DIRS = [
    ".opencode",
    "internal/templates/base",
    ".agents",
    ".cursor",
    ".gemini",
]

EXTENSIONS = {
    ".md",
    ".py",
    ".json",
    ".yml",
    ".yaml",
    ".toml",
    ".txt",
    ".go",
    ".ts",
    ".tsx",
    ".js",
    ".jsx",
    ".sh",
    ".ps1",
}

MAX_BYTES = 1_000_000


@dataclass(frozen=True)
class Match:
    file: Path
    line_no: int
    line: str


def iter_text_files(root: Path) -> list[Path]:
    files: list[Path] = []
    for rel in SCAN_DIRS:
        d = root / rel
        if not d.exists() or not d.is_dir():
            continue
        for p in d.rglob("*"):
            if not p.is_file():
                continue
            if p.match("**/scripts/guardrails/**"):
                continue
            if p.suffix.lower() not in EXTENSIONS:
                continue
            try:
                if p.stat().st_size > MAX_BYTES:
                    continue
            except OSError:
                continue
            files.append(p)
    return files


def scan_file(p: Path) -> list[Match]:
    try:
        text = p.read_text(encoding="utf-8", errors="replace")
    except OSError:
        return []

    hits: list[Match] = []
    for i, line in enumerate(text.splitlines(), start=1):
        if any(rx.search(line) for rx in PATTERNS):
            hits.append(Match(file=p, line_no=i, line=line.rstrip("\n")))
    return hits


def main() -> int:
    if len(sys.argv) < 2:
        print("Usage: python .opencode/scripts/guardrails/no_blueprints.py <project_path>")
        return 2

    root = Path(sys.argv[1]).resolve()
    if not root.exists():
        print(f"Project path does not exist: {root}")
        return 2

    matches: list[Match] = []
    for f in iter_text_files(root):
        matches.extend(scan_file(f))

    if not matches:
        print("[PASS] no_blueprints: no matches")
        return 0

    print("[FAIL] no_blueprints: found blueprint references")
    for m in matches[:200]:
        rel = m.file.relative_to(root)
        print(f"  {rel}:{m.line_no}: {m.line}")
    if len(matches) > 200:
        print(f"  ... and {len(matches) - 200} more")
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
