package board

import (
	"fmt"
	"testing"
)

func TestChanForFun(t *testing.T) {
	queue := make(chan Position, 500)
	fmt.Println(len(queue))
	fmt.Println("test")
	t.Errorf("Should print something")
}

func TestWhyIsItNotPrinting(t *testing.T) {
	fmt.Println("Yeah why??")
	t.Errorf("Should print something")
}
