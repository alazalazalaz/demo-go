package main

import (
	"testing"
)

func Test_reverseListV2(t *testing.T) {
		nodeDemo2 := ListNodeFB()
		_outputList(reverseListV2(nodeDemo2))
}