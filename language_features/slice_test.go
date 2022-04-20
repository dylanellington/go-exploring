package language_features

//sl := make([]int, 0, 4) // original array x
//sl = append(sl, 1, 2, 3) // add 3 numbers to x out of 4 spaces
//fmt.Println(sl)
//
//sl2 := sl[:1] // array y references 1 space in x
//sl2 = append(sl2, 4) // adding to y overwrites position in x
//fmt.Println(sl2)
//fmt.Println("^ Original arrays")
//fmt.Println(sl)
//fmt.Println(sl2[:4]) // show y is a slice of x by increasing y's len to the max cap of x
//fmt.Println("^ Show shared memory arrays, s2 not modified just expanded")
//
//sl2 = append(sl2, 5, 6, 7, 8) // append to y past cap which forces an internal resize of y, so that y no longer shares memory with x
//fmt.Println(sl)
//fmt.Println(sl2[:5])
//fmt.Println("^ No longer shared memory, s2 expanded past s1 capacity, put into a new array")
//sl = append(sl, 1, 2, 3)
//fmt.Println(sl)
//fmt.Println(sl2)
//fmt.Println("^ Changing s1 no longer effects s2")
