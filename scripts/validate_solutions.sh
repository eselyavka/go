#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MODULE_DIR="${ROOT_DIR}/localhost/leetcode"
PROBLEMS_DIR="${MODULE_DIR}/problems"

if [[ ! -d "${MODULE_DIR}" ]]; then
  echo "Missing module directory: ${MODULE_DIR}"
  exit 1
fi

if [[ ! -d "${PROBLEMS_DIR}" ]]; then
  echo "Missing problems directory: ${PROBLEMS_DIR}"
  exit 1
fi

cd "${MODULE_DIR}"

has_package_decl() {
  local file="$1"
  if command -v rg >/dev/null 2>&1; then
    rg -q '^package solutions$' "${file}"
  else
    grep -qE '^package solutions$' "${file}"
  fi
}

unexpected_root_solution_files="$(find . -maxdepth 1 -type f -name 'solution_*.go' | sort)"
if [[ -n "${unexpected_root_solution_files}" ]]; then
  echo "Solution files must live under problems/<range>/, not module root:"
  echo "${unexpected_root_solution_files}"
  exit 1
fi

invalid_range_dirs=""
while IFS= read -r dir; do
  base="$(basename "${dir}")"
  if [[ ! "${base}" =~ ^[0-9]{4}-[0-9]{4}$ ]]; then
    invalid_range_dirs+="${dir}"$'\n'
  fi
done < <(find "${PROBLEMS_DIR}" -mindepth 1 -maxdepth 1 -type d | sort)

if [[ -n "${invalid_range_dirs}" ]]; then
  echo "Invalid range directory names detected:"
  printf "%s" "${invalid_range_dirs}"
  echo "Expected format: problems/0001-0300"
  exit 1
fi

if [[ -z "$(find "${PROBLEMS_DIR}" -mindepth 1 -maxdepth 1 -type d)" ]]; then
  echo "No range directories found in ${PROBLEMS_DIR}"
  exit 1
fi

for dir in "${PROBLEMS_DIR}"/*; do
  [[ -d "${dir}" ]] || continue

  if [[ ! -f "${dir}/types.go" || ! -e "${dir}/constants.go" || ! -e "${dir}/leetcode_shared_libs.go" ]]; then
    echo "Missing required support files in ${dir}"
    echo "Required: types.go, constants.go, leetcode_shared_libs.go"
    exit 1
  fi

  if [[ ! -L "${dir}/constants.go" || ! -L "${dir}/leetcode_shared_libs.go" ]]; then
    echo "Expected symlinks in ${dir}: constants.go and leetcode_shared_libs.go"
    echo "They should point to module-level shared files."
    exit 1
  fi

  bad_solution_files=""
  while IFS= read -r f; do
    name="${f#${dir}/}"
    if [[ ! "${name}" =~ ^solution_[0-9]+\.go$ ]]; then
      bad_solution_files+="${f}"$'\n'
    fi
  done < <(find "${dir}" -maxdepth 1 -type f -name 'solution_*.go' ! -name 'solution_*_test.go' | sort)

  if [[ -n "${bad_solution_files}" ]]; then
    echo "Invalid solution file names in ${dir}:"
    printf "%s" "${bad_solution_files}"
    echo "Expected pattern: solution_<problem_id>.go"
    exit 1
  fi

  bad_test_files=""
  while IFS= read -r f; do
    name="${f#${dir}/}"
    if [[ ! "${name}" =~ ^solution_[0-9]+_test\.go$ ]]; then
      bad_test_files+="${f}"$'\n'
    fi
  done < <(find "${dir}" -maxdepth 1 -type f -name 'solution_*_test.go' | sort)

  if [[ -n "${bad_test_files}" ]]; then
    echo "Invalid test file names in ${dir}:"
    printf "%s" "${bad_test_files}"
    echo "Expected pattern: solution_<problem_id>_test.go"
    exit 1
  fi

  unexpected_go_files="$(
    find "${dir}" -maxdepth 1 -type f -name '*.go' \
      ! -name 'solution_*.go' \
      ! -name 'solution_*_test.go' \
      ! -name 'types.go' \
      ! -name 'constants.go' \
      ! -name 'leetcode_shared_libs.go' \
    | sort
  )"
  if [[ -n "${unexpected_go_files}" ]]; then
    echo "Unexpected Go files in ${dir}:"
    echo "${unexpected_go_files}"
    exit 1
  fi

  invalid_package_files=""
  while IFS= read -r go_file; do
    if ! has_package_decl "${go_file}"; then
      invalid_package_files+="${go_file}"$'\n'
    fi
  done < <(find "${dir}" -maxdepth 1 -type f -name '*.go' | sort)

  if [[ -n "${invalid_package_files}" ]]; then
    echo "Files with non-standard package declaration in ${dir}:"
    printf "%s" "${invalid_package_files}"
    echo "Expected: package solutions"
    exit 1
  fi
done

echo "Solution structure validation passed."
