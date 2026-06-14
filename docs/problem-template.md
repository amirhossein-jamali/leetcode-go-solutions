---
id: PASTE_PROBLEM_NUMBER_HERE
title: "PASTE_PROBLEM_TITLE_HERE"
difficulty: "Easy / Medium / Hard"
level: "Beginner / Mid Level / Advanced"
platform: "LeetCode"
link: "PASTE_PROBLEM_LINK_HERE"
contest: "PASTE_CONTEST_NAME_OR_EMPTY"
status: "Solved / Reviewing / Todo"
language: "Go"
topics:
  - "PASTE_LEETCODE_TOPIC_HERE"
go_concepts:
  - "PASTE_GO_CONCEPT_HERE"
tags:
  - leetcode
  - go
  - PASTE_TOPIC_TAG_HERE
---

<!--
Template usage:

1. Copy this file into a new problem folder as README.md.
2. Replace every PASTE_* placeholder.
3. Keep the YAML frontmatter at the very top for Obsidian.
4. Add only the concepts that are needed for this problem.
5. Delete this comment and the "Filled Sample" section after creating the real problem README.
-->

# Problem Number. Problem Title

## Problem Link

LeetCode: `PASTE_PROBLEM_LINK_HERE`

## Difficulty

Easy / Medium / Hard

## Problem Topics

* PASTE_LEETCODE_TOPIC_HERE

## What to Know Before Solving

General concepts:

* PASTE_GENERAL_CONCEPT_HERE
* PASTE_DATA_STRUCTURE_OR_ALGORITHM_CONCEPT_HERE
* PASTE_ALGORITHM_PATTERN_OR_SIMULATION_RULE_HERE

Go concepts:

* PASTE_GO_TYPE_OR_DATA_STRUCTURE_HERE
* PASTE_GO_SYNTAX_OR_OPERATOR_HERE
* PASTE_GO_STANDARD_LIBRARY_OR_TESTING_CONCEPT_HERE

## Problem Description

Write a short and simple description of the problem here.

Explain what the input is, what the output should be, and the main rule of the problem.

## Function Signature

Expected LeetCode function signature:

```go
func PASTE_FUNCTION_NAME(PASTE_PARAMETERS_HERE) PASTE_RETURN_TYPE_HERE {

}
```

## Examples

### Example 1

Input:

```text
PASTE_INPUT_HERE
```

Output:

```text
PASTE_OUTPUT_HERE
```

Explanation:

```text
PASTE_EXPLANATION_HERE
```

### Example 2

Input:

```text
PASTE_INPUT_HERE
```

Output:

```text
PASTE_OUTPUT_HERE
```

Explanation:

```text
PASTE_EXPLANATION_HERE
```

## Constraints

```text
PASTE_CONSTRAINTS_HERE
```

## Approach

Explain the main idea of the solution in simple words.

Answer these questions:

* What data structure do we use?
* Why does it fit this problem?
* How do we process the input?
* When do we return the answer?

## Algorithm

Step-by-step explanation:

1. Describe the first step.
2. Describe the initial variables or data structures.
3. Describe the main loop or condition.
4. Describe how each item is processed.
5. Describe when the answer is returned.
6. Describe what happens for edge cases.

## Why This Works

Explain why the solution is correct.

Focus on the logic behind the algorithm:

* Why does this method find the correct answer?
* Why does it not miss any valid case?
* Why does it avoid invalid cases?

## Complexity Analysis

Define the variables used in the analysis.

Example:

```text
Let n be the size of the input.
Let m be the maximum length of an item.
```

### Time Complexity

```text
O(...)
```

Explain why the time complexity is `O(...)`.

### Space Complexity

```text
O(...)
```

Explain why the space complexity is `O(...)`.

If useful, separate output space from auxiliary space.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* PASTE_IMPORTANT_GO_DETAIL_HERE
* PASTE_MEMORY_OR_PERFORMANCE_DETAIL_HERE

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage should include:

* All examples from LeetCode
* Minimum input size
* Maximum or large input shape when useful
* Edge cases specific to this problem

## Edge Cases

Important cases to think about:

* Empty input, if allowed
* Input with one element
* Duplicate values
* Negative values, if allowed
* Very large input
* Minimum and maximum constraint values
* PASTE_PROBLEM_SPECIFIC_EDGE_CASE_HERE

## Notes

Add any extra notes here.

For example:

* Alternative approaches
* Trade-offs
* Mistakes to avoid
* Important Go syntax used in the solution

## Completion Checklist

Before committing a new problem, check:

* YAML frontmatter is filled
* Problem link, difficulty, topics, and contest/source are filled
* Required general concepts are listed
* Required Go concepts are listed
* Function signature matches LeetCode
* Examples and constraints are included
* Approach and algorithm are explained
* Correctness explanation is included
* Time and space complexity are included
* `solution.go` uses the expected LeetCode function signature
* `solution_test.go` covers examples and edge cases
* `go test ./...` passes

## Filled Sample

This sample shows the expected shape. Do not copy it as-is for every problem.

```markdown
---
id: 3838
title: "Weighted Word Mapping"
difficulty: "Easy"
level: "Mid Level"
platform: "LeetCode"
link: "https://leetcode.com/problems/weighted-word-mapping/"
contest: "Biweekly Contest 176"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "String"
  - "Simulation"
go_concepts:
  - "Slices"
  - "Strings as byte sequences"
  - "For loops"
  - "ASCII arithmetic"
  - "Modulo operator"
  - "Preallocated byte slices"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - string
  - simulation
  - biweekly-contest-176
---

# 3838. Weighted Word Mapping

## Problem Topics

* Array
* String
* Simulation

## What to Know Before Solving

General concepts:

* How arrays/slices store ordered values
* How strings are processed character by character
* How to simulate a rule exactly as described
* How modulo arithmetic works

Go concepts:

* Function parameters with `[]string` and `[]int`
* Iterating over strings by byte index for lowercase English letters
* Converting a character to an alphabet index with `word[j] - 'a'`
* Preallocating a result with `make([]byte, len(words))`
* Writing table-driven tests with the `testing` package
```
