// Data Structure is a collection of properties that are related together. 
package main

import (
	"fmt"
)

type contactInfo struct{
	email string 
	zipCode int
}

type person struct{
	firstName string
	lastName string
	// contact contactInfo  //Defined with type of another struct that we have defined above
	contactInfo

}

func main()  {
	// // 1st way: Not Recommended
	// alex := person{"Alex", "Anderson"}  // Creating a new structure named as alex based on order not by providing index
	// fmt.Println(alex)

	// // 2nd way: by passing property names
	// alex1 := person{firstName:"Alex1", lastName: "Anderson1"}
	// fmt.Println(alex1)

	// //3rd way:
	// var alex2 person;
	// fmt.Printf("%+v\n", alex2)  //Prints empty structure

	// //Updating struct values
	// alex2.firstName = "Abhijeet"
	// alex2.lastName = "Dandekar"
	// fmt.Printf("%+v\n", alex2)

	//// Embedding Structs: Embedding one struct(contactInfo) inside of another struct(person).
	jim := person{
		firstName: "Jim",
		lastName: "Samson",
		// contact: contactInfo{
		// 	email:"jim@gmail.com",
		// 	zipCode: 9922829,
		// },
		contactInfo : contactInfo{
				email:"jim@gmail.com",
				zipCode: 9922829,
			},
	}
	// fmt.Printf("%+v\n", jim)
	// 1st way 
	jimPointer := &jim  //Gives access to the variable jim's memory address
	jimPointer.updateName("Abhijeet");

	jim.print();

	//2nd way
	jim.updateName("Sarvesh")

	jim.print();

}
// jimPointer := &jim -> & is an operator it gives the memory address of the value the variable jim is pointing at
// pointerToPerson *person -> It gives the value the memory address is pointing at.

// Structs with Receiver Functions
func (p person) print() {
		fmt.Printf("%+v\n", p)
}



// Update the structure using receiver functions using Pointers in structure

// Here person is a type and pointerToPerson is a value of type *person,
// * to a type(person) is a type description which means we are working with a pointer to a person
// * to a value(variable) is an operator which means we want to manipulate the value the pointer referencing
func (pointerToPerson *person) updateName(newFirstName string)  { //Take this memory address and give me the value 
	// that exists at that memory address 
	// pointerToPerson is the memory address that jim exists at.
	(*pointerToPerson).firstName = newFirstName
}

// Reference vs Value Types
// Gotchas with pointers -> the pointers are not needed while updating elements in Slice unlike structures.
// When we create a slice an array and a structure that records the length of slice, tha capacity of slice and a Reference
//  to the underlying array

// Value Types                                  Reference Types 
// Use pointers to change or update             Dont worry about pointers with these 
// things while using function

// int, float, string, bool, structs            slices,maps,channels,pointers,functions