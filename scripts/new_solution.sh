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
MODULE_DIR="${ROOT_DIR}"
PROBLEMS_DIR="${MODULE_DIR}/problems"

range_start=$(( ((problem_id_num - 1) / 300) * 300 + 1 ))
range_end=$(( range_start + 299 ))
range_dir="$(printf '%s/%04d-%04d' "${PROBLEMS_DIR}" "${range_start}" "${range_end}")"
problem_dir="${range_dir}/p${problem_id_norm}"
pkg="p${problem_id_norm}"

mkdir -p "${problem_dir}"

file_path="${problem_dir}/solution_${problem_id_norm}.go"
test_path="${problem_dir}/solution_${problem_id_norm}_test.go"

if [[ -f "${file_path}" ]]; then
  echo "File already exists: ${file_path}"
  exit 1
fi

cat > "${file_path}" <<EOF2
package ${pkg}

// TODO: Replace with your final implementation.
func solution${problem_id_norm}() {
}
EOF2

cat > "${test_path}" <<EOF2
package ${pkg}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution${problem_id_norm}(t *testing.T) {
	assert := assert.New(t)
	_ = assert
	t.Skip("TODO: implement tests for solution ${problem_id_norm}")
}
EOF2

gofmt -w "${file_path}" "${test_path}"

echo "Created ${file_path}"
echo "Created ${test_path}"
