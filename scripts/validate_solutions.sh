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
  local expected="$2"
  if command -v rg >/dev/null 2>&1; then
    rg -q "^package ${expected}$" "${file}"
  else
    grep -qE "^package ${expected}$" "${file}"
  fi
}

for dir in "${PROBLEMS_DIR}"/*; do
  [[ -d "${dir}" ]] || continue

  base="$(basename "${dir}")"
  if [[ ! "${base}" =~ ^[0-9]{4}-[0-9]{4}$ ]]; then
    echo "Invalid range directory name: ${dir}"
    exit 1
  fi

  # No flat solution files at range root.
  if find "${dir}" -maxdepth 1 -type f -name 'solution_*.go' | grep -q .; then
    echo "Solution files must live under problems/<range>/p<id>/"
    find "${dir}" -maxdepth 1 -type f -name 'solution_*.go' | sort
    exit 1
  fi

  while IFS= read -r pdir; do
    [[ -d "${pdir}" ]] || continue

    pbase="$(basename "${pdir}")"
    if [[ ! "${pbase}" =~ ^p[0-9]+$ ]]; then
      echo "Invalid problem directory name: ${pdir}"
      echo "Expected format: p<problem_id>"
      exit 1
    fi

    solution_files="$(find "${pdir}" -maxdepth 1 -type f -name 'solution_*.go' ! -name 'solution_*_test.go' | sort)"
    if [[ -z "${solution_files}" ]]; then
      echo "Missing required solution file in ${pdir}: solution_<id>.go"
      exit 1
    fi
    if [[ "$(printf "%s\n" "${solution_files}" | wc -l | tr -d ' ')" -ne 1 ]]; then
      echo "Expected exactly one solution file in ${pdir}, found:"
      printf "%s\n" "${solution_files}"
      exit 1
    fi

    # Ensure only solution/test go files exist in the problem dir.
    unexpected_go_files="$(find "${pdir}" -maxdepth 1 -type f -name '*.go' \
      ! -name 'solution_*.go' \
      ! -name 'solution_*_test.go' | sort)"
    if [[ -n "${unexpected_go_files}" ]]; then
      echo "Unexpected Go files in ${pdir}:"
      echo "${unexpected_go_files}"
      echo "Expected only solution_<id>.go and optional solution_<id>_test.go"
      exit 1
    fi

    # Ensure package name matches directory.
    invalid_package_files=""
    while IFS= read -r go_file; do
      if ! has_package_decl "${go_file}" "${pbase}"; then
        invalid_package_files+="${go_file}"$'\n'
      fi
    done < <(find "${pdir}" -maxdepth 1 -type f -name '*.go' | sort)

    if [[ -n "${invalid_package_files}" ]]; then
      echo "Files with wrong package declaration in ${pdir}:"
      printf "%s" "${invalid_package_files}"
      echo "Expected: package ${pbase}"
      exit 1
    fi
  done < <(find "${dir}" -mindepth 1 -maxdepth 1 -type d | sort)
done

echo "Solution structure validation passed."
