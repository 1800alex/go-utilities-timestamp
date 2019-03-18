package timestamp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestFrom(t *testing.T) {
	Stubs()

	result := ""

	increase := [4]time.Duration{
		1 * time.Hour,
		1 * time.Hour,
		1 * time.Second,
		1 * time.Millisecond,
	}

	expected := [4]string{
		"1984/04/04 01:00:00.000",
		"1984/04/04 02:00:00.000",
		"1984/04/04 02:00:01.000",
		"1984/04/04 02:00:01.001",
	}

	for i := 0; i < len(increase); i++ {
		_stubClockAdvance(increase[i])

		result = From(DepGetTime())

		assert.Equal(t, result, expected[i])
	}

	StubsRestore()
}

func TestNow(t *testing.T) {
	result := Now()

	re := regexp.MustCompile(`(?m)^[\d]{4}\/[\d]{2}\/[\d]{2} [\d]{2}:[\d]{2}:[\d]{2}\.[\d]{3}$`)

	assert.Equal(t, true, re.MatchString(result))
}

func TestNow_Fake(t *testing.T) {
	Stubs()

	result := ""

	increase := [4]time.Duration{
		1 * time.Hour,
		1 * time.Hour,
		1 * time.Second,
		1 * time.Millisecond,
	}

	expected := [4]string{
		"1984/04/04 01:00:00.000",
		"1984/04/04 02:00:00.000",
		"1984/04/04 02:00:01.000",
		"1984/04/04 02:00:01.001",
	}

	for i := 0; i < len(increase); i++ {
		_stubClockAdvance(increase[i])

		result = Now()

		assert.Equal(t, expected[i], result)
	}

	StubsRestore()
}

func ExampleNow() {
	result := Now()
	fmt.Println(result)
	// result would be similar to "1984/04/04 01:00:00.000"
}

func ExampleFrom() {
	result := From(time.Now())
	fmt.Println(result)
	// result would be similar to "1984/04/04 01:00:00.000"
}
