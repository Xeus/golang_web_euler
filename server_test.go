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
	testMaxNum1, err := getMaxNum(testStr1, euler.PROBLEM1_DEFAULT)
	assert.Nil(err)
	assert.Equal(10, testMaxNum1)

	testStr2 := eulerPath.FindStringSubmatch("/euler/1/100")
	testMaxNum2, err := getMaxNum(testStr2, euler.PROBLEM1_DEFAULT)
	assert.Nil(err)
	assert.Equal(100, testMaxNum2)
}

func Test_Problem1(t *testing.T) {
	assert := assrt.NewAssert(t)
	var solution int64
	var since float64

	solution, since = euler.Problem1(10)
	assert.Equal(23, solution)
	assert.Equal(true, since >= 0)

	solution, since = euler.Problem1(1000)
	assert.Equal(233168, solution)
	assert.Equal(true, since >= 0)

	solution, since = euler.Problem1(0)
	assert.Equal(0, solution)
	assert.Equal(true, since >= 0)

	solution, since = euler.Problem1(-1)
	assert.Equal(0, solution)
	assert.Equal(true, since >= 0)
}
