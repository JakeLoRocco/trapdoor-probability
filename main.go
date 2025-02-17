package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getNextDoor(notChoosableDoors []int, rangeToChooseFrom int) int {
	nextDoor := rand.Intn(rangeToChooseFrom)

	for notChoosableDoors[nextDoor] != 0 {
		nextNum := rand.Intn(rangeToChooseFrom)
		nextDoor = nextNum
	}

	return nextDoor
}

func main() {
	rand.Seed(time.Now().Unix())
	simulations := 100000000
	doors := 20
	numToElim := 10
	numWonSwitching := 0
	numWonStaying := 0
	stayStill := 0
	printDebug := false
	printProgress := true
	allowRechoosingChosen := true // ie only prevent choosing eliminated doors

	// allow choosing settings and seeds

	for i := 0; i < simulations; i++ {
		orderElim := rand.Perm(doors)
		curDoor := rand.Intn(doors)

		if i == 0 {
			stayStill = curDoor
		}

		// Only used for debugging.
		listOfDoorsChosen := []int{}
		if printDebug {
			fmt.Println("order elim: ", orderElim)
		}

		// array with 100 indices, index will be non-0 if already chosen
		chosen := make([]int, doors)
		eliminatedWhileSwitching := false
		eliminatedWhileStaying := false

		for j := 0; j < numToElim; j++ {
			if !allowRechoosingChosen {
				chosen[curDoor] = 1
			}

			if printDebug {
				listOfDoorsChosen = append(listOfDoorsChosen, curDoor)
			}

			if curDoor == orderElim[j] {
				eliminatedWhileSwitching = true
			}

			if stayStill == orderElim[j] {
				eliminatedWhileStaying = true
			}

			// Break early. We already got eliminated.
			if eliminatedWhileSwitching && eliminatedWhileStaying {
				break
			}

			// Update that we can't choose the most recently eliminated door.
			chosen[orderElim[j]] = 1
			curDoor = getNextDoor(chosen, doors)
		}

		if !eliminatedWhileSwitching {
			numWonSwitching += 1
		}

		if !eliminatedWhileStaying {
			numWonStaying += 1
		}

		if printDebug {
			fmt.Println("listofdoorschosen:", listOfDoorsChosen)
		}

		if printProgress {
			if i%100000 == 0 {
				fmt.Println(i)
			}
		}
	}

	// Change printing output
	// print settings here
	fmt.Printf("\nSWITCHING: num won: %v; total sims: %v; percent: %v", numWonSwitching, simulations, float64(numWonSwitching)/float64(simulations))
	fmt.Printf("\nSTAYING: num won: %v; total sims: %v; percent: %v", numWonStaying, simulations, float64(numWonStaying)/float64(simulations))
	fmt.Println()
}
