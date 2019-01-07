package main

import (
	"container/list"
	"fmt"
	"flag"
)

var leftTower = list.New()
var centerTower = list.New()
var rightTower = list.New()
//var Max = 15
var Max int

func towerLen(tower *list.List) int {
	return tower.Len()
}

func towerPeek(tower *list.List) int {
	if tower.Len() == 0 {
		return tower.Len()
	} else {
		return tower.Front().Value.(int)
	}
}

func towerPush(tower *list.List, add int) {
	tower.PushFront(add)
}

func towerPop(tower *list.List) int {
	e := tower.Front()
	tower.Remove(e)
	return e.Value.(int)
}

func towerPrint(tower *list.List, name string) {
	fmt.Printf("%s: ", name)
	for e := tower.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d:", e.Value)
	}
	fmt.Println("")
}

//solution looked up from wikipedia
func solve() {
	for towerLen(rightTower) != Max {
		//if left is not empty && top left is less than top center or center is empty
		if towerPeek(leftTower) > 0 && ((towerPeek(leftTower) < towerPeek(centerTower)) || (towerPeek(centerTower) == 0)) {
			//fmt.Printf( "Moving %d from Left to Center(%d)\n" , towerPeek( leftTower ), towerPeek( centerTower) )
			towerPush(centerTower, towerPop(leftTower))
			towerPrint(leftTower, "Left")
			towerPrint(centerTower, "Center")
			towerPrint(rightTower, "Right")
		//if center is not empty && top left is greater than center or left is emtpy
		} else if towerPeek(centerTower) > 0 && ((towerPeek(leftTower) > towerPeek(centerTower)) || (towerPeek(leftTower) == 0)) {
			//fmt.Printf( "Moving %d from Center to Left(%d)\n" , towerPeek( centerTower ), towerPeek( leftTower) )
			towerPush(leftTower, towerPop(centerTower))
			towerPrint(leftTower, "Left")
			towerPrint(centerTower, "Center")
			towerPrint(rightTower, "Right")
		}

		if towerPeek(leftTower) > 0 && ((towerPeek(leftTower) < towerPeek(rightTower)) || (towerPeek(rightTower) == 0)) {
			//fmt.Printf( "Moving %d from Left to Right(%d)\n" , towerPeek( leftTower ), towerPeek( rightTower) )
			towerPush(rightTower, towerPop(leftTower))
			towerPrint(leftTower, "Left")
			towerPrint(centerTower, "Center")
			towerPrint(rightTower, "Right")
		} else if towerPeek(rightTower) > 0 && ((towerPeek(leftTower) > towerPeek(rightTower)) || (towerPeek(leftTower) == 0)) {
			//fmt.Printf( "Moving %d from Right to Left(%d)\n" , towerPeek( rightTower ), towerPeek( leftTower) )
			towerPush(leftTower, towerPop(rightTower))
			towerPrint(leftTower, "Left")
			towerPrint(centerTower, "Center")
			towerPrint(rightTower, "Right")
		}

		if towerPeek(centerTower) > 0 && ((towerPeek(centerTower) < towerPeek(rightTower)) || (towerPeek(rightTower) == 0)) {
			//fmt.Printf( "Moving %d from Center to Right(%d)\n" , towerPeek( centerTower ), towerPeek( rightTower) )
			towerPush(rightTower, towerPop(centerTower))
			towerPrint(leftTower, "Left")
			towerPrint(centerTower, "Center")
			towerPrint(rightTower, "Right")
		} else if towerPeek(rightTower) > 0 && ((towerPeek(centerTower) > towerPeek(rightTower)) || (towerPeek(centerTower) == 0)) {
			//fmt.Printf( "Moving %d from Right to Center(%d)\n" , towerPeek( rightTower ), towerPeek( centerTower) )
			towerPush(centerTower, towerPop(rightTower))
			towerPrint(leftTower, "Left")
			towerPrint(centerTower, "Center")
			towerPrint(rightTower, "Right")
		}
		fmt.Printf("End of loop\n\n")
	}
}

func main() {

	flag.IntVar(&Max, "size",10,"size of the tower")

	flag.Parse()

	fmt.Printf("Size of the tower: %d\n" , Max )

	//fill the tower with values
	for i := Max; i > 0; i-- {
		towerPush(leftTower, i)
	}

	towerPrint(leftTower, "Left")
	towerPrint(centerTower, "Center")
	towerPrint(rightTower, "Right")

	solve()

	towerPrint(leftTower, "Left")
	towerPrint(centerTower, "Center")
	towerPrint(rightTower, "Right")
}
