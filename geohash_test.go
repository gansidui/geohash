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
