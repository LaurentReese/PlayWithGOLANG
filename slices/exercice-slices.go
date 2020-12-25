// See https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"
import "time"
import "fmt"

const MY_SIZE=20000000

// Two solutions are possible
// 1) Without using make to create the slices
// I simply declare a slice and do append(s) on it to make it grow
//
// 2) By using make to create the slices
// Then I use range to go through the created slice and initialize items

// !! BUT solution 2 (with make) is much much faster (at least 10 times) !!

func Pic(dx, dy int) [][]uint8 {
	var subres []uint8
	for i:=0;i<dx;i++ { subres=append(subres,uint8(i)) }
//	subres := make([]uint8, dx, dx)	
//	for i := range subres { subres[i] = uint8(i)}

	var subres2[][]uint8
	for i:=0;i<dy;i++ { subres2=append(subres2,subres) }	
//	subres2 := make([][]uint8, dy, dy)	
//	for i := range subres2 { subres2[i] = subres}
	return subres2
}

func main() {
	now := time.Now()
	pic.Show(Pic) // to draw the "blue" image

	Pic(MY_SIZE, MY_SIZE) // pour mesurer le temps passé
	later := time.Now()
	fmt.Printf("The call took %v to run.\n", later.Sub(now))	
}

// func Show(f func(dx, dy int) [][]uint8)
