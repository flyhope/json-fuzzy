package fuzzy

import (
	"encoding/json/v2"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

type testFuzzyIntStruct[T constraints.Integer] struct {
	Value T `json:"value"`
}

type fuzzyTestCase[T constraints.Integer] struct {
	name        string
	jsonData    string
	expected    T
	expectError bool
}

func runFuzzyIntTestCases[T constraints.Integer](t *testing.T, tests []fuzzyTestCase[T], opts ...json.Options) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result testFuzzyIntStruct[T]
			err := json.Unmarshal([]byte(tt.jsonData), &result, opts...)

			if tt.expectError {
				assert.Error(t, err, fmt.Sprintf("Expected an error for test case: %s", tt.name))
			} else {
				assert.NoError(t, err, fmt.Sprintf("Did not expect an error for test case: %s", tt.name))
				assert.Equal(t, tt.expected, result.Value, fmt.Sprintf("Mismatched expectation for test case: %s", tt.name))
			}
		})
	}
}

func runFuzzyIntTests[T constraints.Integer](t *testing.T, tests []fuzzyTestCase[T]) {
	unmarshaler := FuzzyUnmarshalerOrigin()
	runFuzzyIntTestCases(t, tests, unmarshaler)
}

func runFuzzyIntFullTests[T constraints.Integer](t *testing.T, tests []fuzzyTestCase[T]) {
	unmarshaler := FuzzyUnmarshalerFull()
	runFuzzyIntTestCases(t, tests, unmarshaler)
}

func getCommonTestCases[T constraints.Integer](
	validVal T,
	validStr string,
	truncatedFloatStr string,
	expectedTruncated T,
) []fuzzyTestCase[T] {
	return []fuzzyTestCase[T]{
		{
			name:     "Valid integer number",
			jsonData: fmt.Sprintf(`{"value": %v}`, validVal),
			expected: validVal,
		},
		{
			name:     "Valid string number",
			jsonData: fmt.Sprintf(`{"value": "%s"}`, validStr),
			expected: validVal,
		},
		{
			name:     "String with whitespace",
			jsonData: fmt.Sprintf(`{"value": " %s "}`, validStr),
			expected: validVal,
		},
		{
			name:     "Valid float number, truncated",
			jsonData: fmt.Sprintf(`{"value": %s}`, truncatedFloatStr),
			expected: expectedTruncated,
		},
		{
			name:     "Valid string float number, truncated",
			jsonData: fmt.Sprintf(`{"value": "%s"}`, truncatedFloatStr),
			expected: expectedTruncated,
		},
		{
			name:     "String float with whitespace, truncated",
			jsonData: fmt.Sprintf(`{"value": " %s "}`, truncatedFloatStr),
			expected: expectedTruncated,
		},
		{
			name:     "Zero value",
			jsonData: `{"value": 0}`,
			expected: 0,
		},
		{
			name:     "Empty string value",
			jsonData: `{"value": ""}`,
			expected: 0,
		},
		{
			name:     "String with only whitespace",
			jsonData: `{"value": "   "}`,
			expected: 0,
		},
		{
			name:     "Null value",
			jsonData: `{"value": null}`,
			expected: 0,
		},
		{
			name:     "Boolean true value",
			jsonData: `{"value": true}`,
			expected: 1,
		},
		{
			name:     "Boolean false value",
			jsonData: `{"value": false}`,
			expected: 0,
		},
		// --- Common Error Cases ---
		{
			name:        "Invalid non-numeric string",
			jsonData:    `{"value": "abc"}`,
			expectError: true,
		},
		{
			name:        "Object value",
			jsonData:    `{"value": {}}`,
			expectError: true,
		},
		{
			name:        "Array value",
			jsonData:    `{"value": []}`,
			expectError: true,
		},
	}
}

func testToInt8(v int) int8 {
	return int8(v)
}

func testToInt16(v int) int16 {
	return int16(v)
}

func testToInt32(v int64) int32 {
	return int32(v)
}

func testToUint8(v int) uint8 {
	return uint8(v)
}

func testToUint16(v int) uint16 {
	return uint16(v)
}

func testToUint32(v int64) uint32 {
	return uint32(v)
}

func TestFuzzyInt(t *testing.T) {
	tests := getCommonTestCases[int](123, "123", "123.99", 123)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyInt8(t *testing.T) {
	tests := getCommonTestCases[int8](127, "127", "127.99", 127)
	tests = append(tests,
		fuzzyTestCase[int8]{
			name:     "Overflow integer",
			jsonData: `{"value": 128}`,
			expected: testToInt8(128),
		},
		fuzzyTestCase[int8]{
			name:     "Overflow string integer",
			jsonData: `{"value": "128"}`,
			expected: testToInt8(128),
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyInt16(t *testing.T) {
	tests := getCommonTestCases[int16](32767, "32767", "32767.99", 32767)
	tests = append(tests,
		fuzzyTestCase[int16]{
			name:     "Overflow integer",
			jsonData: `{"value": 32768}`,
			expected: testToInt16(32768),
		},
		fuzzyTestCase[int16]{
			name:     "Overflow string integer",
			jsonData: `{"value": "32768"}`,
			expected: testToInt16(32768),
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyInt32(t *testing.T) {
	tests := getCommonTestCases[int32](2147483647, "2147483647", "2147483647.99", 2147483647)
	tests = append(tests,
		fuzzyTestCase[int32]{
			name:     "Overflow integer",
			jsonData: `{"value": 2147483648}`,
			expected: testToInt32(2147483648),
		},
		fuzzyTestCase[int32]{
			name:     "Overflow string integer",
			jsonData: `{"value": "2147483648"}`,
			expected: testToInt32(2147483648),
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyInt64(t *testing.T) {
	tests := getCommonTestCases[int64](9223372036854775807, "9223372036854775807", "9223372036854775000.99", 9223372036854775000)
	tests = append(tests,
		fuzzyTestCase[int64]{
			name:        "Overflow string integer",
			jsonData:    `{"value": "9223372036854775808"}`,
			expectError: true,
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyUint(t *testing.T) {
	tests := getCommonTestCases[uint](123, "123", "123.99", 123)
	tests = append(tests,
		fuzzyTestCase[uint]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[uint]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyUint8(t *testing.T) {
	tests := getCommonTestCases[uint8](255, "255", "255.99", 255)
	tests = append(tests,
		fuzzyTestCase[uint8]{
			name:     "Overflow integer",
			jsonData: `{"value": 256}`,
			expected: testToUint8(256),
		},
		fuzzyTestCase[uint8]{
			name:     "Overflow string integer",
			jsonData: `{"value": "256"}`,
			expected: testToUint8(256),
		},
		fuzzyTestCase[uint8]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[uint8]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyUint16(t *testing.T) {
	tests := getCommonTestCases[uint16](65535, "65535", "65535.99", 65535)
	tests = append(tests,
		fuzzyTestCase[uint16]{
			name:     "Overflow integer",
			jsonData: `{"value": 65536}`,
			expected: testToUint16(65536),
		},
		fuzzyTestCase[uint16]{
			name:     "Overflow string integer",
			jsonData: `{"value": "65536"}`,
			expected: testToUint16(65536),
		},
		fuzzyTestCase[uint16]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[uint16]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyUint32(t *testing.T) {
	tests := getCommonTestCases[uint32](4294967295, "4294967295", "4294967295.99", 4294967295)
	tests = append(tests,
		fuzzyTestCase[uint32]{
			name:     "Overflow integer",
			jsonData: `{"value": 4294967296}`,
			expected: testToUint32(4294967296),
		},
		fuzzyTestCase[uint32]{
			name:     "Overflow string integer",
			jsonData: `{"value": "4294967296"}`,
			expected: testToUint32(4294967296),
		},
		fuzzyTestCase[uint32]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[uint32]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntTests(t, tests)
}

func TestFuzzyUint64(t *testing.T) {
	tests := getCommonTestCases[uint64](18446744073709551615, "18446744073709551615", "18446744073709551000.99", 18446744073709551000)
	tests = append(tests,
		fuzzyTestCase[uint64]{
			name:        "Overflow string integer",
			jsonData:    `{"value": "18446744073709551616"}`,
			expectError: true,
		},
		fuzzyTestCase[uint64]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[uint64]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntTests(t, tests)
}

type testTypeInt64 int64

func TestFuzzyCustomInt64(t *testing.T) {
	tests := getCommonTestCases[testTypeInt64](9223372036854775807, "9223372036854775807", "9223372036854775000.99", 9223372036854775000)
	tests = append(tests,
		fuzzyTestCase[testTypeInt64]{
			name:        "Overflow string integer",
			jsonData:    `{"value": "9223372036854775808"}`,
			expectError: true,
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeInt int

func TestFuzzyCustomInt(t *testing.T) {
	tests := getCommonTestCases[testTypeInt](123, "123", "123.99", 123)
	runFuzzyIntFullTests(t, tests)
}

type testTypeInt8 int8

func TestFuzzyCustomInt8(t *testing.T) {
	tests := getCommonTestCases[testTypeInt8](127, "127", "127.99", 127)
	tests = append(tests,
		fuzzyTestCase[testTypeInt8]{
			name:     "Overflow integer",
			jsonData: `{"value": 128}`,
			expected: testTypeInt8(testToInt8(128)),
		},
		fuzzyTestCase[testTypeInt8]{
			name:     "Overflow string integer",
			jsonData: `{"value": "128"}`,
			expected: testTypeInt8(testToInt8(128)),
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeInt16 int16

func TestFuzzyCustomInt16(t *testing.T) {
	tests := getCommonTestCases[testTypeInt16](32767, "32767", "32767.99", 32767)
	tests = append(tests,
		fuzzyTestCase[testTypeInt16]{
			name:     "Overflow integer",
			jsonData: `{"value": 32768}`,
			expected: testTypeInt16(testToInt16(32768)),
		},
		fuzzyTestCase[testTypeInt16]{
			name:     "Overflow string integer",
			jsonData: `{"value": "32768"}`,
			expected: testTypeInt16(testToInt16(32768)),
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeInt32 int32

func TestFuzzyCustomInt32(t *testing.T) {
	tests := getCommonTestCases[testTypeInt32](2147483647, "2147483647", "2147483647.99", 2147483647)
	tests = append(tests,
		fuzzyTestCase[testTypeInt32]{
			name:     "Overflow integer",
			jsonData: `{"value": 2147483648}`,
			expected: testTypeInt32(testToInt32(2147483648)),
		},
		fuzzyTestCase[testTypeInt32]{
			name:     "Overflow string integer",
			jsonData: `{"value": "2147483648"}`,
			expected: testTypeInt32(testToInt32(2147483648)),
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeUint uint

func TestFuzzyCustomUint(t *testing.T) {
	tests := getCommonTestCases[testTypeUint](123, "123", "123.99", 123)
	tests = append(tests,
		fuzzyTestCase[testTypeUint]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[testTypeUint]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeUint8 uint8

func TestFuzzyCustomUint8(t *testing.T) {
	tests := getCommonTestCases[testTypeUint8](255, "255", "255.99", 255)
	tests = append(tests,
		fuzzyTestCase[testTypeUint8]{
			name:     "Overflow integer",
			jsonData: `{"value": 256}`,
			expected: testTypeUint8(testToUint8(256)),
		},
		fuzzyTestCase[testTypeUint8]{
			name:     "Overflow string integer",
			jsonData: `{"value": "256"}`,
			expected: testTypeUint8(testToUint8(256)),
		},
		fuzzyTestCase[testTypeUint8]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[testTypeUint8]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeUint16 uint16

func TestFuzzyCustomUint16(t *testing.T) {
	tests := getCommonTestCases[testTypeUint16](65535, "65535", "65535.99", 65535)
	tests = append(tests,
		fuzzyTestCase[testTypeUint16]{
			name:     "Overflow integer",
			jsonData: `{"value": 65536}`,
			expected: testTypeUint16(testToUint16(65536)),
		},
		fuzzyTestCase[testTypeUint16]{
			name:     "Overflow string integer",
			jsonData: `{"value": "65536"}`,
			expected: testTypeUint16(testToUint16(65536)),
		},
		fuzzyTestCase[testTypeUint16]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[testTypeUint16]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeUint32 uint32

func TestFuzzyCustomUint32(t *testing.T) {
	tests := getCommonTestCases[testTypeUint32](4294967295, "4294967295", "4294967295.99", 4294967295)
	tests = append(tests,
		fuzzyTestCase[testTypeUint32]{
			name:     "Overflow integer",
			jsonData: `{"value": 4294967296}`,
			expected: testTypeUint32(testToUint32(4294967296)),
		},
		fuzzyTestCase[testTypeUint32]{
			name:     "Overflow string integer",
			jsonData: `{"value": "4294967296"}`,
			expected: testTypeUint32(testToUint32(4294967296)),
		},
		fuzzyTestCase[testTypeUint32]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[testTypeUint32]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntFullTests(t, tests)
}

type testTypeUint64 uint64

func TestFuzzyCustomUint64(t *testing.T) {
	tests := getCommonTestCases[testTypeUint64](18446744073709551615, "18446744073709551615", "18446744073709551000.99", 18446744073709551000)
	tests = append(tests,
		fuzzyTestCase[testTypeUint64]{
			name:        "Overflow string integer",
			jsonData:    `{"value": "18446744073709551616"}`,
			expectError: true,
		},
		fuzzyTestCase[testTypeUint64]{
			name:        "Negative integer",
			jsonData:    `{"value": -1}`,
			expectError: true,
		},
		fuzzyTestCase[testTypeUint64]{
			name:        "Negative string integer",
			jsonData:    `{"value": "-1"}`,
			expectError: true,
		},
	)
	runFuzzyIntFullTests(t, tests)
}
