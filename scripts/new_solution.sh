#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <problem_id>"
  exit 1
fi

problem_id="$1"

if [[ ! "${problem_id}" =~ ^[0-9]+$ ]]; then
  echo "Problem ID must be numeric."
  exit 1
fi

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MODULE_DIR="${ROOT_DIR}/localhost/leetcode"
file_path="${MODULE_DIR}/solution_${problem_id}.go"

if [[ -f "${file_path}" ]]; then
  echo "File already exists: ${file_path}"
  exit 1
fi

cat > "${file_path}" <<EOF
package solutions

// TODO: Replace with your final implementation.
func solution${problem_id}() {
}
EOF

gofmt -w "${file_path}"
echo "Created ${file_path}"
