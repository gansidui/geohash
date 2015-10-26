package geohash

import (
	"fmt"
	"testing"
)

func TestGeoHash(t *testing.T) {
	geohash, b := Encode(39.92324, 116.3906, 5)
	fmt.Println(geohash)
	fmt.Println(b)

	fmt.Println(GetNeighbors(39.92324, 116.3906, 5))

	// output:
	// 	wx4g0
	// &{39.90234375 39.9462890625 116.3671875 116.4111328125}
	// [wx4g0 wx4g2 wx4fb wx4ep wx4g1 wx4er wx4dz wx4g3 wx4fc]
}

func TestLoopNeighbors(t *testing.T) {
	fmt.Printf("%v\n",GetNeighbors(39.92324, 116.3906, 5))
	fmt.Println("------------------")
	loopa := LoopNeighbors(39.92324, 116.3906, 5,1)
	for loop,hashs := range loopa{
		fmt.Printf("loop: %d\n",loop)
		for _,hash := range hashs{
			fmt.Println("\t"+hash)
		}
	}
	fmt.Println("------------------")
	loopb :=LoopNeighbors(39.92324, 116.3906, 5, 2)
	for loop, hashs := range loopb {
		fmt.Printf("loop: %d\n", loop)
		for _, hash := range hashs {
			fmt.Println("\t"+hash)
		}
	}
	fmt.Println("------------------")
	loopc :=LoopNeighbors(39.92324, 116.3906, 5, 3)
	for loop, hashs := range loopc {
		fmt.Printf("loop: %d\n", loop)
		for _, hash := range hashs {
			fmt.Println("\t"+hash)
		}
	}
}
