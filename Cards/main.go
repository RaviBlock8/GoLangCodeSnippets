package main

func main(){
	// // So this how you declare a variable
	// var card string="Ace of spades"
	// // We can create variable without mentioning data type
	// inferredCard:= newCard()
	// /**
	// * So to hold list of items we have two data types
	// * Array : Fixed length list of things
	// * Slices : An array that can grow or shrink
	// * All the elements in an array or slice should be of same data type
	// **/
	// // We can create a slice like this cards=[string]{"hello"}
	// cards:=deck{newCard(),newCard()}
	// /**
	// * We can add new element in this slice using append keyword.
	// * But keep in mind , append doesn't modify actual slice.
	// * But it generates a new slice with new element inside it.
	// **/
	// cards=append(cards,"Six of hearts")
	// fmt.Println(card)
	// fmt.Println(inferredCard)
	// cards.print();

	// cardDeck:=newDeck();
	// cardDeck.print();
	// cardDeck.writeDataToFile("testFile")
	// decFromFile,err:=newDeckFromFile("testFile")
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// decFromFile.print()
	//---HERE WE WILL SHUFFLE THE CARDS---
	deck2:=newDeck()
	deck2,_=deck2.shuffle()
	deck2.print()

}

func newCard() string{
	return "Five of spades"
}

/**
* ---SLICES IN DEPTH----
* 1. So in slices it allows us to select a subsection of slice.
* 2. To do this we have syntax sliceName[indexToInclude:uptoExcludingIndex]
* 3. so let's say we have slice test:=[]int{1,2,3,4}
* 4. Then test[0:2]=[1,2]
**/