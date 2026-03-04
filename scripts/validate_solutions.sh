#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MODULE_DIR="${ROOT_DIR}/localhost/leetcode"

if [[ ! -d "${MODULE_DIR}" ]]; then
  echo "Missing module directory: ${MODULE_DIR}"
  exit 1
fi

cd "${MODULE_DIR}"

bad_files=""
while IFS= read -r f; do
  name="${f#./}"
  if [[ ! "${name}" =~ ^solution_[0-9]+\.go$ ]]; then
    bad_files+="${f}"$'\n'
  fi
done < <(find . -maxdepth 1 -type f -name 'solution_*.go' | sort)

if [[ -n "${bad_files}" ]]; then
  echo "Invalid solution file naming detected:"
  printf "%s" "${bad_files}"
  echo "Expected pattern: solution_<problem_id>.go"
  exit 1
fi

non_solution_files="$(find . -maxdepth 1 -type f -name '*.go' ! -name 'solution_*.go' ! -name 'solutions_test.go' ! -name 'types.go' ! -name 'constants.go' ! -name 'leetcode_shared_libs.go' | sort)"
if [[ -n "${non_solution_files}" ]]; then
  echo "Unexpected Go files detected:"
  echo "${non_solution_files}"
  exit 1
fi

invalid_package_files=""
while IFS= read -r go_file; do
  if ! rg -q '^package solutions$' "${go_file}"; then
    invalid_package_files+="${go_file}"$'\n'
  fi
done < <(find . -maxdepth 1 -type f -name '*.go' | sort)

if [[ -n "${invalid_package_files}" ]]; then
  echo "Files with non-standard package declaration:"
  printf "%s" "${invalid_package_files}"
  echo "Expected: package solutions"
  exit 1
fi

echo "Solution structure validation passed."
