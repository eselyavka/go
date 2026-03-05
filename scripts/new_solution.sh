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

problem_id_num=$((10#${problem_id}))
problem_id_norm="${problem_id_num}"

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MODULE_DIR="${ROOT_DIR}/localhost/leetcode"
PROBLEMS_DIR="${MODULE_DIR}/problems"

range_start=$(( ((problem_id_num - 1) / 300) * 300 + 1 ))
range_end=$(( range_start + 299 ))
range_dir="$(printf '%s/%04d-%04d' "${PROBLEMS_DIR}" "${range_start}" "${range_end}")"

mkdir -p "${range_dir}"

if [[ ! -f "${range_dir}/types.go" ]]; then
  cp "${MODULE_DIR}/types.go" "${range_dir}/types.go"
fi

if [[ ! -e "${range_dir}/constants.go" ]]; then
  ln -s ../../constants.go "${range_dir}/constants.go"
fi

if [[ ! -e "${range_dir}/leetcode_shared_libs.go" ]]; then
  ln -s ../../leetcode_shared_libs.go "${range_dir}/leetcode_shared_libs.go"
fi

file_path="${range_dir}/solution_${problem_id_norm}.go"
test_path="${range_dir}/solution_${problem_id_norm}_test.go"

if [[ -f "${file_path}" ]]; then
  echo "File already exists: ${file_path}"
  exit 1
fi

cat > "${file_path}" <<EOF
package solutions

// TODO: Replace with your final implementation.
func solution${problem_id_norm}() {
}
EOF

gofmt -w "${file_path}"

if [[ ! -f "${test_path}" ]]; then
  cat > "${test_path}" <<EOF
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution${problem_id_norm}(t *testing.T) {
	assert := assert.New(t)
	_ = assert
	t.Skip("TODO: implement tests for solution ${problem_id_norm}")
}
EOF
  gofmt -w "${test_path}"
fi

echo "Created ${file_path}"
echo "Created ${test_path}"
