package main

import "fmt"

// create a new type veicle.
// the underlying type is a struct.
// the fields:
// doors
// color
// create two new types : truck & sedan.
// the underlying type of each of these new types is a struct
// emebed the "vehicle" type in both truck & sedan.
// give truck the field "fourwheel" which will be set to bool.
// give sedan the field "luxurt" which will be set to bool.
// using the vehicle , truck and sedan struct:
// using a composite literal, create a values of type truck and assign values to the fields.
// using a composite literal, create a values of type sedan and assign values to the fields.


type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle
	fourwheel bool

}

type sedan struct{
	vehicle
	luxurt bool

}

func main(){

	x:= truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},

		fourwheel: true,
	}
	fmt.Println(x)
	y:= sedan{
		vehicle: vehicle{
			doors: 2,
			color : "red",
		},
		luxurt: true,
	}
	fmt.Println(y)

}