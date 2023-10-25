// package main

// import (
// 	"fmt"
// )

// type shape interface {
// 	getArea() float64
// }


// type triangle struct {
// 	base float64
// 	height float64
// }
// type square struct {
// 	sideLength float64
// }


// func main()  {
// 	t := triangle{}
// 	s := square{}
// 	printArea(t)
// 	printArea(s)
// }

// func printArea(sh shape)  {
// 	fmt.Println("Area is:" , sh.getArea())
// }


// func (t triangle) getArea() float64  {
// 	fmt.Println("Enter the Height:")
// 	fmt.Scan(&t.height) 
// 	fmt.Println("Enter the Base:")
// 	fmt.Scan(&t.base)
// 	return 0.5 * t.height * t.base
// }

// func (s square) getArea() float64  {
// 	fmt.Println("Enter the SideLength:")
// 	fmt.Scan(&s.sideLength)
// 	return s.sideLength * s.sideLength
// }