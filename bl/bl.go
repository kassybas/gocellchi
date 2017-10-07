package bl

import "fmt"
import "time"

type Item struct {
	PosX  int32
	PosY  int32
	Size  int32
	Color uint32
	VelX  int32
	VelY  int32
	Angle int32
	Speed int32
}

var Items []Item

func (i Item) Poll() *Item {
	return &i
}

func Start() {
	Items = make([]Item, 0)
	fmt.Println("Items loaded")

	item := Item{Color: 0xffff0000, PosX: 0, PosY: 0, Size: 10, Speed: 1, VelX: 5, VelY: 2}
	Items = append(Items, item)
	timer := time.NewTimer(time.Millisecond * 20)
	for true {
		tick()
		<-timer.C
		timer.Reset(time.Millisecond * 50)
	}
}

func tick() {
	for i, _ := range Items {
		Items[i].PosX = Items[i].PosX + Items[i].VelX
		Items[i].PosY = Items[i].PosY + Items[i].VelY
		fmt.Println("Updated positions")
		fmt.Printf("%+v\n", Items[0])
	}
}
