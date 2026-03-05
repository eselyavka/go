# Go LeetCode Solutions

Repository for practicing LeetCode problems in Go with standardized structure, CI checks, and pre-commit automation.

## Structure

- `localhost/leetcode/`: Go module root.
- `localhost/leetcode/problems/<0001-0300>/`: range folders (300 problems each).
- `localhost/leetcode/problems/<range>/solution_<id>.go`: solution implementation.
- `localhost/leetcode/problems/<range>/solution_<id>_test.go`: solution tests.
- `localhost/leetcode/constants.go`: canonical shared constants file.
- `localhost/leetcode/leetcode_shared_libs.go`: canonical shared helper functions.
- `localhost/leetcode/problems/<range>/constants.go`: symlink to `../../constants.go`.
- `localhost/leetcode/problems/<range>/leetcode_shared_libs.go`: symlink to `../../leetcode_shared_libs.go`.
- `localhost/leetcode/problems/<range>/types.go`: range-local compatibility types.
- `.github/workflows/ci.yml`: CI pipeline.
- `scripts/new_solution.sh`: Creates standardized solution stubs.
- `scripts/validate_solutions.sh`: Validates layout, naming, package conventions, and shared-file symlinks.
- `.pre-commit-config.yaml`: Local pre-commit checks.

## Conventions

- Go version: `1.26` (`go.mod` + `go.work`).
- Each solution lives in a range folder:
  - `localhost/leetcode/problems/<range>/solution_<problem_id>.go`
- Each test lives alongside the solution:
  - `localhost/leetcode/problems/<range>/solution_<problem_id>_test.go`
- File package must be `solutions`.
- Shared constants/helpers are single-source at module root and symlinked into each range folder.

## Local Commands

```bash
make fmt
make fmt-check
make validate
make test
make verify
```

`make verify` runs format checks, solution validation, and tests across all ranges.

Create a new solution stub:

```bash
./scripts/new_solution.sh 1234
```

## Pre-commit

Install once:

```bash
~/.venv/py3_11/bin/pre-commit install
```

Run manually:

```bash
~/.venv/py3_11/bin/pre-commit run --all-files
```

## CI/CD

GitHub Actions runs on every push/PR:

- format + structure validation
- tests
- lint (`golangci-lint v2`)

For this repo, "CD" is continuous delivery of quality gates to `main` (no deployment target yet).
