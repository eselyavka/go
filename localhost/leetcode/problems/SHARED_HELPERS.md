# Shared Helpers Inventory

This document tracks helpers shared through `localhost/leetcode/util`.
Problem packages inline compatibility helpers directly in `solution_<id>.go`.

## Global Types

Defined in `localhost/leetcode/util/types.go`:
- `TreeNode`
- `ListNode`
- `Node`
- `NodeNext`
- `SparseVector`
- `PickIdx`

## Global Constants

Defined in `localhost/leetcode/util/constants.go`:
- `MaxUint`
- `MinUint`
- `MaxInt`
- `MinInt`

## Shared Helper Functions

Defined in `localhost/leetcode/util/helpers.go`:
- Exported: `MaxInts`, `MinInts`, `Reverse`, `Abs`, `GCD`, `LCM`, `IntSliceReverse`
- Internal helpers used by existing solutions/tests:
  `stringSlicesEqual`, `intSliceEqual`, `sum`, `nTrue`, `isPalindrome`,
  `binarySearch`, `binarySearchMin`, `initLinkedList`, `reverseInt`,
  `reverseString`, `inStringArray`, `in2dIntArray`, `arrayUniq`,
  `binaryTreeHeight`, `int2dSliceIsEqual`, `bool2int`, `countSetBits`,
  `binaryTreeBFS`, `flipArrayXOR`, `sliceToString`, `twoDSliceToString`,
  `copy2DArray`, `copyRow`, `flatten2DArray`, `sumDigits`.

## Migration Rule (Per Problem Package)

For each refactor:
1. Keep the problem algorithm in its own package (for example, `p1`, `p386`).
2. Keep reusable helpers/types in `localhost/leetcode/util`.
3. Keep backward-compatible local helper/type names in each `solution_<id>.go`.

## Migration Progress

- Migrated to per-problem package names (`package p<id>`) with util-backed bridges.
