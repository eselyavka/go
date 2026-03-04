# Go LeetCode Solutions

Repository for practicing LeetCode problems in Go, with consistent structure, automation, and quality checks.

## Structure

- `localhost/leetcode/`: Go module with all solutions and tests.
- `.github/workflows/ci.yml`: CI pipeline.
- `scripts/new_solution.sh`: Creates standardized solution stubs.
- `scripts/validate_solutions.sh`: Validates naming and package conventions.
- `.pre-commit-config.yaml`: Local pre-commit checks.

## Conventions

- Each solution lives in `localhost/leetcode/solution_<problem_id>.go`.
- File package must be `solutions`.
- Run `gofmt` on all files before commit.
- Keep shared types/utilities in:
  - `localhost/leetcode/types.go`
  - `localhost/leetcode/constants.go`
  - `localhost/leetcode/leetcode_shared_libs.go`

## Local Commands

```bash
make fmt
make test
make verify
```

`make verify` runs format checks, solution validation, and tests.

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
