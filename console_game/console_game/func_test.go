package func_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNight(t *testing.T) {
	I := functions.New(10, 50, 20, 30)
	I.Night()
	want := functions.New(8, 70, 18, 25)

	assert.Equal(t, want, I)
}

func TestGenerate(t *testing.T) {
	Input := functions.New(10, 50, 20, 30)
	Test := functions.Generate(Input, 50)
	Output := 2 / 7

	assert.Equal(t, Output, Test)
}

func TestLose(t *testing.T) {
	Input := functions.New(10, 50, 20, 30)
	Input.Lose(Input, 50)
	Output := int64(-40)

	assert.Equal(t, Input.Hp, Output)
}

func TestWin(t *testing.T) {
	Input := functions.New(10, 50, 20, 30)
	Input.Win(Input, 50)
	Output := int64(50)

	assert.Equal(t, Input.Rep, Output)
}
