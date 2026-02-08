#!/usr/bin/env python3
"""\
Master Checklist Runner - OpenKit
================================

Orchestrates validation scripts in priority order.
Use this for incremental validation during development.

Usage:
    python .opencode/scripts/checklist.py .
    python .opencode/scripts/checklist.py . --url <URL>

Priority Order:
    P0: Security Scan (vulnerabilities, secrets)
    P1: Lint & Type Check (code quality)
    P2: Schema Validation (if database exists)
    P3: Test Runner (unit/integration tests)
    P4: UX Audit (accessibility)
    P5: SEO Check
    P6: Performance (requires URL)
"""

import sys
import subprocess
import argparse
from pathlib import Path
from typing import List, Optional


class Colors:
    HEADER = "\033[95m"
    BLUE = "\033[94m"
    CYAN = "\033[96m"
    GREEN = "\033[92m"
    YELLOW = "\033[93m"
    RED = "\033[91m"
    ENDC = "\033[0m"
    BOLD = "\033[1m"


def print_header(text: str):
    print(f"\n{Colors.BOLD}{Colors.CYAN}{'=' * 60}{Colors.ENDC}")
    print(f"{Colors.BOLD}{Colors.CYAN}{text.center(60)}{Colors.ENDC}")
    print(f"{Colors.BOLD}{Colors.CYAN}{'=' * 60}{Colors.ENDC}\n")


def print_step(text: str):
    print(f"{Colors.BOLD}{Colors.BLUE}{text}{Colors.ENDC}")


def print_success(text: str):
    print(f"{Colors.GREEN}{text}{Colors.ENDC}")


def print_warning(text: str):
    print(f"{Colors.YELLOW}{text}{Colors.ENDC}")


def print_error(text: str):
    print(f"{Colors.RED}{text}{Colors.ENDC}")


CORE_CHECKS = [
    ("No Deprecated Alias Guard", ".opencode/scripts/guardrails/no_blueprints.py", True),
    ("Security Scan", ".opencode/skills/vulnerability-scanner/scripts/security_scan.py", True),
    ("Lint Check", ".opencode/skills/lint-and-validate/scripts/lint_runner.py", True),
    ("Schema Validation", ".opencode/skills/database-design/scripts/schema_validator.py", False),
    ("Test Runner", ".opencode/skills/testing-patterns/scripts/test_runner.py", False),
    ("UX Audit", ".opencode/skills/frontend-design/scripts/ux_audit.py", False),
    ("SEO Check", ".opencode/skills/seo-fundamentals/scripts/seo_checker.py", False),
]


PERFORMANCE_CHECKS = [
    ("Lighthouse Audit", ".opencode/skills/performance-profiling/scripts/lighthouse_audit.py", True),
    ("Playwright E2E", ".opencode/skills/webapp-testing/scripts/playwright_runner.py", False),
]


def run_script(
    name: str, script_path: Path, project_path: str, url: Optional[str] = None
) -> dict:
    if not (script_path.exists() and script_path.is_file()):
        print_warning(f"[SKIP] {name}: script not found")
        return {"name": name, "passed": True, "output": "", "skipped": True}

    print_step(f"Running: {name}")

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
            timeout=300,
        )
        passed = result.returncode == 0
        if passed:
            print_success(f"[PASS] {name}")
        else:
            print_error(f"[FAIL] {name}")
            if result.stderr:
                print(f"  Error: {result.stderr[:200]}")
        return {
            "name": name,
            "passed": passed,
            "output": result.stdout,
            "error": result.stderr,
            "skipped": False,
        }
    except subprocess.TimeoutExpired:
        print_error(f"[FAIL] {name}: TIMEOUT (>5 minutes)")
        return {"name": name, "passed": False, "output": "", "error": "Timeout", "skipped": False}
    except Exception as e:
        print_error(f"[FAIL] {name}: ERROR - {str(e)}")
        return {"name": name, "passed": False, "output": "", "error": str(e), "skipped": False}


def print_summary(results: List[dict]) -> bool:
    print_header("CHECKLIST SUMMARY")

    passed_count = sum(1 for r in results if r["passed"] and not r.get("skipped"))
    failed_count = sum(1 for r in results if not r["passed"] and not r.get("skipped"))
    skipped_count = sum(1 for r in results if r.get("skipped"))

    print(f"Total checks: {len(results)}")
    print(f"Passed: {passed_count}")
    print(f"Failed: {failed_count}")
    print(f"Skipped: {skipped_count}\n")

    for r in results:
        if r.get("skipped"):
            status = "[SKIP]"
        elif r["passed"]:
            status = "[PASS]"
        else:
            status = "[FAIL]"
        print(f"{status} {r['name']}")
    print()

    if failed_count > 0:
        print_error(f"{failed_count} check(s) failed")
        return False

    print_success("All checks passed")
    return True


def main():
    parser = argparse.ArgumentParser(
        description="Run OpenKit validation checklist",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""\
Examples:
  python .opencode/scripts/checklist.py .
  python .opencode/scripts/checklist.py . --url http://localhost:3000
""",
    )
    parser.add_argument("project", help="Project path to validate")
    parser.add_argument("--url", help="URL for performance checks (lighthouse, playwright)")
    parser.add_argument(
        "--skip-performance",
        action="store_true",
        help="Skip performance checks even if URL provided",
    )

    args = parser.parse_args()
    project_path = Path(args.project).resolve()
    if not project_path.exists():
        print_error(f"Project path does not exist: {project_path}")
        raise SystemExit(1)

    print_header("OPENKIT - MASTER CHECKLIST")
    print(f"Project: {project_path}")
    print(f"URL: {args.url if args.url else 'Not provided (performance checks skipped)'}")

    results = []

    print_header("CORE CHECKS")
    for name, script_path, required in CORE_CHECKS:
        script = project_path / script_path
        result = run_script(name, script, str(project_path))
        results.append(result)

        if required and not result["passed"] and not result.get("skipped"):
            print_error(f"CRITICAL: {name} failed. Stopping.")
            print_summary(results)
            raise SystemExit(1)

    if args.url and not args.skip_performance:
        print_header("PERFORMANCE CHECKS")
        for name, script_path, required in PERFORMANCE_CHECKS:
            script = project_path / script_path
            result = run_script(name, script, str(project_path), args.url)
            results.append(result)
            if required and not result["passed"] and not result.get("skipped"):
                print_error(f"CRITICAL: {name} failed. Stopping.")
                print_summary(results)
                raise SystemExit(1)

    all_passed = print_summary(results)
    raise SystemExit(0 if all_passed else 1)


if __name__ == "__main__":
    main()
