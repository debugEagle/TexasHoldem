//main函数是一个demo展示

package main

import (
	"dealmachine"
	"fmt"
	"hand"
	//"card"
	//"sort"
)

func main() {
	//
	h1 := hand.GetHand()
	h1.Init()

	h2 := hand.GetHand()
	h2.Init()

	d := dealmachine.GetDealMachine()
	d.Init()
	_ = d.Shuffle()

	_ = h1.SetCard(d.Deal())
	_ = h1.SetCard(d.Deal())

	_ = h2.SetCard(d.Deal())
	_ = h2.SetCard(d.Deal())

	for i := 0; i < 5; i++ {
		c := d.Deal()
		_ = h1.SetCard(c)
		_ = h2.SetCard(c)
	}

	_ = h1.AnalyseHand()
	_ = h2.AnalyseHand()

	h1.ShowHand()
	h2.ShowHand()

	if h1.Level > h2.Level {
		fmt.Println("h1 win ")
	} else if h1.Level < h2.Level {
		fmt.Println("h2 win ")
	} else if h1.FinalValue > h2.FinalValue {
		fmt.Println("h1 win ")
	} else if h1.FinalValue < h2.FinalValue {
		fmt.Println("h2 win ")
	} else {
		fmt.Println("tie")
	}
}
