# Go LeetCode Solutions

Repository for practicing LeetCode problems in Go, with consistent structure, automation, and quality checks.

## Structure

- `localhost/leetcode/`: Go module root.
- `localhost/leetcode/problems/<0001-0300>/`: range folders (300 problems each).
- `localhost/leetcode/problems/<range>/solution_<id>.go`: solution file.
- `localhost/leetcode/problems/<range>/solution_<id>_test.go`: test file.
- `localhost/leetcode/constants.go` and `localhost/leetcode/leetcode_shared_libs.go`: single shared source files symlinked into each range folder.
- `.github/workflows/ci.yml`: CI pipeline.
- `scripts/new_solution.sh`: Creates standardized solution stubs.
- `scripts/validate_solutions.sh`: Validates naming and package conventions.
- `.pre-commit-config.yaml`: Local pre-commit checks.

## Conventions

- Each solution lives in its range folder:
  - `localhost/leetcode/problems/<range>/solution_<problem_id>.go`
- File package must be `solutions`.
- Run `gofmt` on all files before commit.
- Shared constants/helpers are maintained once in module root and linked into each range folder:
  - `localhost/leetcode/constants.go`
  - `localhost/leetcode/leetcode_shared_libs.go`
- `localhost/leetcode/types.go` is used as the template when initializing a new range folder.

## Local Commands

```bash
make fmt
make test
make verify
```

`make verify` runs format checks, solution validation, and tests across all ranges.

## Pre-commit

Install once:

```bash
pip install pre-commit
pre-commit install
```

Run manually:

```bash
pre-commit run --all-files
```

## CI/CD

GitHub Actions runs on every push/PR:

- format + structure validation
- tests
- lint (`golangci-lint`)

For this repo, "CD" is continuous delivery of quality gates to `main` (no deployment target yet).
