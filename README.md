# Software Engineering Assessment
![Banked-Image](https://user-images.githubusercontent.com/46195831/135531836-3cf56e1c-288b-46c1-8ffd-023bc9a8e23a.png)

This repo is my submission for the software engineering test given by Banked.com.
### Solution
Here is the required function.
```Go
func denominationChange(amount float64) map[float64]int {  
   m := make(map[float64]int)  
  
  // cashUnits is a slice that holds all the various cash units provided.  
 // cashUnits is stored as a slice of float64.  
 cashUnits := []float64{  
      500.00, 200.00, 100.00, 50.00, 20.00, 10.00, 5.00, 2.00, 1.00,  
    0.50, 0.20, 0.10, 0.05, 0.02, 0.01,  
  }  
  
   //Range over each cashUnit  
  for _, val := range cashUnits {  
      // if the cash unit is greater than the amount,  
	 //we want to skip to a lesser cash unit  
	 if val > amount {  
         continue  
	  }  
     // Get the integer that sums up to the amount/val  
	 // using the Modf function provided by the Math package  
	 integer, _ := math.Modf(amount / val)  
      m[val] = int(integer)  
      // Retrieve the remainder of the amount/cashUnit.  
	  amount = math.Mod(amount, val)  
      if amount == 0 {  
         break  
	  }  
   }  
   return m  
}
```
I decided to take my implementation a little further, making it easy to try out my solution by building a simple console application that takes in a user input and prints out the possible cashUnits.

### The Implementation
![Demo Of My Solution](https://user-images.githubusercontent.com/46195831/135535366-263c01b6-981a-4e8f-83d9-3842d0a65591.gif)
* This program takes simply takes an amount and then prints out the number of each cash-unit that fits into the given input

### How To Run
***Build on your machine***
* To Build this program, You need to have Go installed on your machine.*
* Clone/Unzip this project & navigate to the project path.
* Run the following command `go build`.
* Run `./banked $(name_of_text_file)`.
*
***Run Tests***
* To run the tests, run the following command. `go test -bench=.`

### Technical Notes
*  I had lots of fun while working on this project.
* I used float64 to represent the amount. I did that just for the purpose of this simple project, it’s advisable not to use floats to represent currency.
* My implementation relies heavily on the order of the [cashUnit slice](https://github.com/Ghvstcode/Banked/blob/main/main.go#L43) . In this implementation, I simply sorted it from highest to lowest and hardcoded it. If for example, the user gets to pass in the cashUnits, we will want to do some sort of sorting to make sure it is ranked from highest to lowest.

