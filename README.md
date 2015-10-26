## geohash

geohash algorithm: http://en.wikipedia.org/wiki/Geohash

~~~ go
package main

import (
	"fmt"
	"github.com/sillydong/geohash"
)

func main() {
	latitude := 39.92324
	longitude := 116.3906
	precision := 5

	hash, box := geohash.Encode(latitude, longitude, precision)

	fmt.Println(hash)
	fmt.Println(box.MinLat, box.MaxLat, box.MinLng, box.MaxLng)

	neighbors := geohash.GetNeighbors(latitude, longitude, precision)
	for _, hash = range neighbors {
		fmt.Print(hash, " ")
	}
}

~~~


## water diffusion

As sometimes 9 grids can not get enough data, `LoopNeighbors` can diffusion like water until you get enough data.

~~~ go
package main

import (
	"fmt"
	"github.com/sillydong/geohash"
)

func main() {
	latitude := 39.92324
	longitude := 116.3906
	precision := 5

	loopneighbors := geohash.LoopNeighbors(latitude, longitude, precision, 3)
	for loop, hashs := range loopneighbors {
		fmt.Printf("loop: %d\n", loop)
		for _, hash := range hashs {
			fmt.Println("\t"+hash)
		}
	}
}

~~~

## LICENSE

MIT
