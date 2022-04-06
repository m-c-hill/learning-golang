# Cards

## Functionality

newDeck: Create a list of playing cards (array of strings).

print: Log out the contents of a deck of cards.

shuffle: Shuffles all the cards in the deck.

deal: Create a 'hand' of cards.

saveToFile: Save a list of cards to a file on the local machine.

newDeckFromFile: Load a list of cards from the local machine.

## Notes

Explicit type declaration:
`var card string = "Ace of Spades"`

Type inference for a new variable:
`card := "Ace of Spades"`

Variable reassignment after declaration:
`card = "Two of Spades"`

Variables can be initialised outside of a function, but they cannot have values assigned to them.

Arrays: fixed length list.

Slice: array that can grow or shrink in length

Function Receiver sets a method on variables that we create (ie. a custom type).

Byte slice: slice of integers where every element corresponds to an ASCII code.

`"Hi there! -> [72 105 32...]`

For the functions in the rand package to work you have to set a 'Seed' value. This has to be a good random value as decided by the user because - as per https://golang.org/pkg/math/rand/#Rand.Seed this is the value golang uses to set the system to a deterministic state first to then generate a number based on that value.

`rand.Seed(time.Now().UnixNano())`