package say

// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description string
	input       int64
	expected    string
	expectError bool
}{
	{
		description: "zero",
		input:       0,
		expected:    "zero",
		expectError: false,
	},
	{
		description: "one",
		input:       1,
		expected:    "one",
		expectError: false,
	},
	{
		description: "fourteen",
		input:       14,
		expected:    "fourteen",
		expectError: false,
	},
	{
		description: "twenty",
		input:       20,
		expected:    "twenty",
		expectError: false,
	},
	{
		description: "twenty-two",
		input:       22,
		expected:    "twenty-two",
		expectError: false,
	},
	{
		description: "one hundred",
		input:       100,
		expected:    "one hundred",
		expectError: false,
	},
	{
		description: "one hundred twenty-three",
		input:       123,
		expected:    "one hundred twenty-three",
		expectError: false,
	},
	{
		description: "one thousand",
		input:       1000,
		expected:    "one thousand",
		expectError: false,
	},
	{
		description: "one thousand two hundred thirty-four",
		input:       1234,
		expected:    "one thousand two hundred thirty-four",
		expectError: false,
	},
	{
		description: "one million",
		input:       1000000,
		expected:    "one million",
		expectError: false,
	},
	{
		description: "one million two thousand three hundred forty-five",
		input:       1002345,
		expected:    "one million two thousand three hundred forty-five",
		expectError: false,
	},
	{
		description: "one billion",
		input:       1000000000,
		expected:    "one billion",
		expectError: false,
	},
	{
		description: "a big number",
		input:       987654321123,
		expected:    "nine hundred eighty-seven billion six hundred fifty-four million three hundred twenty-one thousand one hundred twenty-three",
		expectError: false,
	},
	{
		description: "numbers below zero are out of range",
		input:       -1,
		expected:    "",
		expectError: true,
	},
	{
		description: "numbers above 999,999,999,999 are out of range",
		input:       1000000000000,
		expected:    "",
		expectError: true,
	},
}
