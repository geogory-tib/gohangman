package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

/*
* Simple hangman game by geogory tibisov
*	all word for the game are strored in dict.txt
*	serving as a easy method to extend the games dictonary
 */
//Ascii art for the hangman
var hangmanSprites = []string{
	`___
|/   |
|   
|    
|    
|    
|
|_____`,

	`____
|/   |
|   (_)
|    
|    
|    
|
|_____`,
	`____
|/   |
|   (_)
|    |
|    |    
|    
|
|_____`,
	` ____
|/   |
|   (_)
|   \|
|    |
|    
|
|_____`,

	` ____
|/   |
|   (_)
|   \|/
|    |
|    
|
|_____`,

	` ____
|/   |
|   (_)
|   \|/
|    |
|   / 
|
|_____`,
	` ____
|/   |
|   (_)
|   \|/
|    |
|   / \
|
|_____`,
	` ____
|/   |
|   (_)
|   /|\
|    |
|   | |
|
|_____`,
}

// loads the dictonary and returns a string slice
func getDict() []string {
	var returnStrings []string // slice to be returned
	file, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) //creating the scanner
	for scanner.Scan() {              //for loop that scans the file
		word := scanner.Text() //this method returns the text from the scanner
		returnStrings = append(returnStrings, word)
	}
	if err := scanner.Err(); err != nil { // crash the program if any problem happens witht the scanner
		log.Fatal(err)
	}
	file.Close() //close the file
	return returnStrings
}

// switch statement to print the coresponding Ascii sprite with the number of wrong attempts
func printMan(wrongGuesses int) {
	switch wrongGuesses {
	case 0:
		fmt.Println(hangmanSprites[0])
	case 1:
		fmt.Println(hangmanSprites[1])
	case 2:
		fmt.Println(hangmanSprites[2])
	case 3:
		fmt.Println(hangmanSprites[3])
	case 4:
		fmt.Println(hangmanSprites[4])
	case 5:
		fmt.Println(hangmanSprites[6])

	case 6:
		fmt.Println(hangmanSprites[7])

	}
}

// function to print the word whilst leaving unknown letters blank
func printWord(word string, correct []bool) {
	for i := 0; i < len(correct); i++ {
		if correct[i] == true { // prints letters if their index in the correctLeters slice is true
			fmt.Printf("%c ", word[i])
		} else { // if the word's index in correctLeters is not true print a underscore
			fmt.Print("_" + " ")
		}
	}
	fmt.Print("\n")
}

func main() {
	fmt.Println("Would you like to play hangman?\n y or n?")
	var userIn byte //userinput stored as an ascii char
	_, err := fmt.Scanf("%c\n", &userIn)
	if err != nil {
		log.Fatal(err)
	}
	if userIn == 'y' {
		wrongGuesses := 0                        //number of wrong attempts
		dict := getDict()                        //loading the dictorny into the slice dict
		word := dict[rand.Intn(len(dict))]       // using rand to assign a random word from the dict slice
		correctLeters := make([]bool, len(word)) //array used to tell what indexes of the string the player has got right
		correct := 0                             //used to see how many letters the player got right more effecient than indexes the correctLeters slice
		for wrongGuesses != 6 {
			printMan(wrongGuesses)         //print hangaman sprite
			printWord(word, correctLeters) //print word in its current state
			fmt.Print("Type a Letter \n")
			_, err = fmt.Scanf("%c\n", &userIn)
			if err != nil {
				log.Fatal(err)
			}
			if correct+1 == len(word) { //since correct is added up in a foor lop indexing arrays i have to add 1 to it
				fmt.Println("You got the word!!!")
				return
			}
			//checking to see if user got a letter in a word
			for i := 0; true; i++ {
				if i == len(word) {
					wrongGuesses++
					break
				}
				if word[i] == userIn {
					correctLeters[i] = true
					correct += 1
					break
				}

			}

		}
	} else {
		fmt.Println("Good bye!")
		return
	}
	printMan(6) //since the loop breaks at 6 it doesnt print the last hangman sprite
}
