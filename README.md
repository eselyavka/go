# Go LeetCode Solutions

Repository for practicing LeetCode problems in Go with standardized structure, CI checks, and pre-commit automation.

## Structure

- `localhost/leetcode/`: Go module root.
- `localhost/leetcode/problems/<0001-0300>/`: range folders (300 problems each).
- `localhost/leetcode/problems/<range>/p<id>/solution_<id>.go`: solution implementation.
- `localhost/leetcode/problems/<range>/p<id>/solution_<id>_test.go`: solution tests.
- `localhost/leetcode/util/constants.go`: shared constants.
- `localhost/leetcode/util/helpers.go`: shared helper functions.
- `localhost/leetcode/util/types.go`: shared types.
- `.github/workflows/ci.yml`: CI pipeline.
- `scripts/new_solution.sh`: Creates standardized solution stubs.
- `scripts/validate_solutions.sh`: Validates per-problem layout and package conventions.
- `.pre-commit-config.yaml`: Local pre-commit checks.

## Conventions

- Go version: `1.26` (`go.mod` + `go.work`).
- Each solution lives in:
  - `localhost/leetcode/problems/<range>/p<problem_id>/solution_<problem_id>.go`
- Each test lives alongside the solution in the same `p<problem_id>` directory.
- Package name must match the problem directory:
  - directory `p382` uses `package p382`
- Shared helpers/types/constants live under:
  - `localhost/leetcode/util`

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

## Workflow Notes

- CI pins Go to `1.26.x` in `.github/workflows/ci.yml`.
- Lint workflow uses `golangci-lint-action@v9` with `golangci-lint v2`.
- `scripts/validate_solutions.sh` supports both `rg` and `grep`, so it works on runners without `ripgrep`.
