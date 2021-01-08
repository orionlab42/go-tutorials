package main

import "fmt"

//type personNew struct{
//	first string
//	last string
//	favIceCream []string
//}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println("Ex.05.1")
	//p1 := personNew{
	//	first: "Pere",
	//	last: "Garcia",
	//	favIceCream: []string{"Cream", "Chocolate", "Beans",},
	//}
	//p2 := personNew{
	//	first: "Pol",
	//	last: "Sole",
	//	favIceCream: []string{"Vanilla", "Marshmallows", "Chocolate", "Coffee",},
	//}
	//for i, v := range p1.favIceCream {
	//	fmt.Printf("For %v and %v the %v favorite ice cream is %v\n",p1.first, p1.last, i+1, v)
	//}
	//for i, v := range p2.favIceCream {
	//	fmt.Printf("For %v and %v the %v favorite ice cream is %v\n",p2.first, p2.last, i+1, v)
	//}
	//
	fmt.Println("Ex.05.2")
	//m := map[string]personNew {
	//	p1.last : p1,
	//	p2.last : p2,
	//}
	//for k, v := range m {
	//	fmt.Println(k)
	//	fmt.Println(v.first, v.last)
	//	for i, val := range v.favIceCream {
	//		fmt.Println(i, val)
	//	}
	//}

	fmt.Println("Ex.05.3")
	car1 := truck{
		vehicle: vehicle{
			doors: 3,
			color: "red",
		},
		fourWheel: true,
	}
	car2 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "bue",
		},
		luxury: true,
	}
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car1.doors, car1.color, car1.fourWheel)
	fmt.Println(car2.doors, car2.color, car2.luxury)

	fmt.Println("Ex.05.4")
	balloons := struct {
		color  string
		number int
	}{
		color:  "red",
		number: 3,
	}
	fmt.Println(balloons.color)
}
