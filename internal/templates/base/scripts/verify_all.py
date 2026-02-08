#!/usr/bin/env python3
"""\
Full Verification Suite - OpenKit
================================

Runs a complete validation suite including security, lint, tests, and optional
performance/E2E checks (requires a URL).

Usage:
    python .opencode/scripts/verify_all.py . --url <URL>
"""

import argparse
import subprocess
import sys
from datetime import datetime
from pathlib import Path
from typing import Optional


class Colors:
    BLUE = "\033[94m"
    CYAN = "\033[96m"
    GREEN = "\033[92m"
    YELLOW = "\033[93m"
    RED = "\033[91m"
    ENDC = "\033[0m"
    BOLD = "\033[1m"


def print_header(text: str):
    print(f"\n{Colors.BOLD}{Colors.CYAN}{'=' * 70}{Colors.ENDC}")
    print(f"{Colors.BOLD}{Colors.CYAN}{text.center(70)}{Colors.ENDC}")
    print(f"{Colors.BOLD}{Colors.CYAN}{'=' * 70}{Colors.ENDC}\n")


def print_step(text: str):
    print(f"{Colors.BOLD}{Colors.BLUE}{text}{Colors.ENDC}")


def print_success(text: str):
    print(f"{Colors.GREEN}{text}{Colors.ENDC}")


def print_warning(text: str):
    print(f"{Colors.YELLOW}{text}{Colors.ENDC}")


def print_error(text: str):
    print(f"{Colors.RED}{text}{Colors.ENDC}")


VERIFICATION_SUITE = [
    {
        "category": "Security",
        "checks": [
            (
                "Security Scan",
                ".opencode/skills/vulnerability-scanner/scripts/security_scan.py",
                True,
            ),
        ],
    },
    {
        "category": "Code Quality",
        "checks": [
            (
                "Lint Check",
                ".opencode/skills/lint-and-validate/scripts/lint_runner.py",
                True,
            ),
            (
                "Type Coverage",
                ".opencode/skills/lint-and-validate/scripts/type_coverage.py",
                False,
            ),
        ],
    },
    {
        "category": "Data Layer",
        "checks": [
            (
                "Schema Validation",
                ".opencode/skills/database-design/scripts/schema_validator.py",
                False,
            ),
        ],
    },
    {
        "category": "Testing",
        "checks": [
            ("Test Suite", ".opencode/skills/testing-patterns/scripts/test_runner.py", False),
        ],
    },
    {
        "category": "UX and Accessibility",
        "checks": [
            ("UX Audit", ".opencode/skills/frontend-design/scripts/ux_audit.py", False),
            (
                "Accessibility Check",
                ".opencode/skills/frontend-design/scripts/accessibility_checker.py",
                False,
            ),
        ],
    },
    {
        "category": "SEO and Content",
        "checks": [
            ("SEO Check", ".opencode/skills/seo-fundamentals/scripts/seo_checker.py", False),
            ("GEO Check", ".opencode/skills/geo-fundamentals/scripts/geo_checker.py", False),
        ],
    },
    {
        "category": "Performance",
        "requires_url": True,
        "checks": [
            (
                "Lighthouse Audit",
                ".opencode/skills/performance-profiling/scripts/lighthouse_audit.py",
                True,
            ),
        ],
    },
    {
        "category": "E2E Testing",
        "requires_url": True,
        "checks": [
            (
                "Playwright E2E",
                ".opencode/skills/webapp-testing/scripts/playwright_runner.py",
                False,
            ),
        ],
    },
    {
        "category": "Mobile",
        "checks": [
            ("Mobile Audit", ".opencode/skills/mobile-design/scripts/mobile_audit.py", False),
        ],
    },
    {
        "category": "Internationalization",
        "checks": [
            ("i18n Check", ".opencode/skills/i18n-localization/scripts/i18n_checker.py", False),
        ],
    },
]


def run_script(
    name: str, script_path: Path, project_path: str, url: Optional[str] = None
) -> dict:
    if not script_path.exists():
        print_warning(f"[SKIP] {name}: script not found")
        return {"name": name, "passed": True, "skipped": True, "duration": 0}

    print_step(f"Running: {name}")
    start_time = datetime.now()

    cmd = [sys.executable, str(script_path), project_path]
    if url and (
        "lighthouse" in script_path.name.lower() or "playwright" in script_path.name.lower()
    ):
        cmd.append(url)

    try:
        result = subprocess.run(
            cmd,
            capture_output=True,
            text=True,
            timeout=600,
        )
        duration = (datetime.now() - start_time).total_seconds()
        passed = result.returncode == 0
        if passed:
            print_success(f"[PASS] {name} ({duration:.1f}s)")
        else:
            print_error(f"[FAIL] {name} ({duration:.1f}s)")
            if result.stderr:
                print(f"  {result.stderr[:300]}")
        return {
            "name": name,
            "passed": passed,
            "output": result.stdout,
            "error": result.stderr,
            "skipped": False,
            "duration": duration,
        }
    except subprocess.TimeoutExpired:
        duration = (datetime.now() - start_time).total_seconds()
        print_error(f"[FAIL] {name}: TIMEOUT (>{duration:.0f}s)")
        return {
            "name": name,
            "passed": False,
            "skipped": False,
            "duration": duration,
            "error": "Timeout",
        }
    except Exception as e:
        duration = (datetime.now() - start_time).total_seconds()
        print_error(f"[FAIL] {name}: ERROR - {str(e)}")
        return {
            "name": name,
            "passed": False,
            "skipped": False,
            "duration": duration,
            "error": str(e),
        }


def main():
    parser = argparse.ArgumentParser(description="Run OpenKit full verification suite")
    parser.add_argument("project", help="Project path")
    parser.add_argument("--url", required=True, help="Base URL for performance/E2E checks")
    args = parser.parse_args()

    project_path = Path(args.project).resolve()
    if not project_path.exists():
        print_error(f"Project path does not exist: {project_path}")
        raise SystemExit(1)

    url = args.url
    print_header("OPENKIT - FULL VERIFICATION")
    print(f"Project: {project_path}")
    print(f"URL: {url}")

    failures = 0

    for suite in VERIFICATION_SUITE:
        category = suite["category"]
        requires_url = suite.get("requires_url", False)
        print_header(category)
        for name, rel_script, required in suite["checks"]:
            script = project_path / rel_script
            result = run_script(name, script, str(project_path), url if requires_url else None)
            if required and not result.get("skipped") and not result["passed"]:
                failures += 1
                print_error(f"CRITICAL: {name} failed. Stopping.")
                raise SystemExit(1)
            if not result.get("skipped") and not result["passed"]:
                failures += 1

    if failures:
        print_error(f"Verification failed: {failures} check(s) failed")
        raise SystemExit(1)

    print_success("Verification passed")
    raise SystemExit(0)


if __name__ == "__main__":
    main()
