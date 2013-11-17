/**
 * see http://golangtutorials.blogspot.com/2011/10/gotest-unit-testing-and-benchmarking-go.html
 */

package main

import(
 	"testing"
 	"./euler"
 	"./assrt"
 	// "reflect"
 	// "fmt"
)

func Test_Env(t *testing.T) {
	assert := assrt.NewAssert(t)

	assert.Log("Testing is on.")
}

func Test_RegExp_Short(t *testing.T) {
	assert := assrt.NewAssert(t)

	// w/ trailing slash
	testShortURL := eulerPath.FindStringSubmatch("/euler/1/")
	assert.Equal(3, len(testShortURL))
	assert.Equal("1", testShortURL[1])
	assert.Equal("", testShortURL[2])
	// w/o trailing slash
	testShortURL = eulerPath.FindStringSubmatch("/euler/1")
	assert.Equal(3, len(testShortURL))
	assert.Equal("1", testShortURL[1])
	assert.Equal("", testShortURL[2])
}

func Test_RegExp_Long(t *testing.T) {
	assert := assrt.NewAssert(t)

	// w/ trailing slash
	testLongURL := eulerPath.FindStringSubmatch("/euler/1/100/")
	assert.Equal(3, len(testLongURL))
	assert.Equal("1", testLongURL[1])
	assert.Equal("100", testLongURL[2])
	// w/o trailing slash
	testLongURL = eulerPath.FindStringSubmatch("/euler/1/100")
	assert.Equal(3, len(testLongURL))
	assert.Equal("1", testLongURL[1])
	assert.Equal("100", testLongURL[2])
}

func Test_RegExp_Junk(t *testing.T) {
	assert := assrt.NewAssert(t)

	testJunkyURL := eulerPath.FindStringSubmatch("/euler/1/100/junk/junk/junk")
	assert.Equal(3, len(testJunkyURL))
	assert.Equal("1", testJunkyURL[1])
	assert.Equal("100", testJunkyURL[2])
}

func Test_GetMaxNum(t *testing.T) {
	assert := assrt.NewAssert(t)

	testStr1 := eulerPath.FindStringSubmatch("/euler/1")
	testMaxNum1, err := getMaxNum(testStr1, euler.ProblemDefaults()[1])
	assert.Nil(err)
	assert.Equal(10, testMaxNum1)

	testStr2 := eulerPath.FindStringSubmatch("/euler/1/100")
	testMaxNum2, err := getMaxNum(testStr2, euler.ProblemDefaults()[1])
	assert.Nil(err)
	assert.Equal(100, testMaxNum2)

	testStr3 := eulerPath.FindStringSubmatch("/euler/1/-100")
	testMaxNum3, err := getMaxNum(testStr3, euler.ProblemDefaults()[1])
	assert.Nil(err)
	assert.Equal(-100, testMaxNum3)
}

func Test_Problem1(t *testing.T) {
	assert := assrt.NewAssert(t)

	var desc, solution, since = euler.Problem1(euler.ProblemDefaults()[1])
	assert.PositiveLen(desc)
	assert.Equal(23, solution)
	assert.Equal(true, since >= 0)

	desc, solution, since = euler.Problem1(1000)
	assert.Equal(233168, solution)

	desc, solution, since = euler.Problem1(0)
	assert.Equal(0, solution)

	desc, solution, since = euler.Problem1(-1)
	assert.Equal(0, solution)
}

func Test_Problem2(t *testing.T) {
	assert := assrt.NewAssert(t)

	var desc, solution, since = euler.Problem2(euler.ProblemDefaults()[2])
	assert.PositiveLen(desc)
	assert.Equal(2, solution)
	assert.Equal(true, since >= 0)

	desc, solution, since = euler.Problem2(4000000)
	assert.Equal(4613732, solution)

	desc, solution, since = euler.Problem2(-5)
	assert.Equal(0, solution)
}

func Test_Problem3(t *testing.T) {
	assert := assrt.NewAssert(t)

	var desc, solution, since, err = euler.Problem3(euler.ProblemDefaults()[3])
	assert.PositiveLen(desc)
	assert.Equal(6857, solution)
	assert.Equal(true, since >= 0)

	desc, solution, since, err = euler.Problem3(400000003)
	assert.Equal(17257, solution)

	desc, solution, since, err = euler.Problem3(-10)
	assert.NotNil(err)
}

func Test_Problem4(t *testing.T) {
	assert := assrt.NewAssert(t)

	var desc, solution, since, err = euler.Problem4(euler.ProblemDefaults()[4])
	assert.PositiveLen(desc)
	assert.Equal(906609, solution)
	assert.Equal(true, since >= 0)

	desc, solution, since, err = euler.Problem4(99)
	assert.Equal(9009, solution)

	desc, solution, since, err = euler.Problem4(-10)
	assert.NotNil(err)

	desc, solution, since, err = euler.Problem4(10)
	assert.NotNil(err)

	desc, solution, since, err = euler.Problem4(10000000)
	assert.NotNil(err)
}
