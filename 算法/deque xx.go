package main 
import (
	lane "gopkg.in/oleiade/lane.v1"
	"fmt"
	"strings"
)

func main(){
	// Let's create a new deque data structure
	var deque *lane.Deque = lane.NewDeque()

	// And push some content into it using the Append
	// and Prepend methods
	deque.Append("easy as")
	deque.Prepend("123")
	deque.Append("do re mi")
	deque.Prepend("abc")
	deque.Append([]int{1,2,3})

	// Now let's take a look at what are the first and
	// last element stored in the Deque
	firstValue := deque.First()
	lastValue := deque.Last()
	ll := deque.Pop()
	l2 := deque.Pop()
	l3 := deque.First()
	fmt.Println(firstValue) // "abc"
	fmt.Println(lastValue)  // 1
	fmt.Println(ll)
	fmt.Println(l2)
	fmt.Println(l3)

	// Okay now let's play with the Pop and Shift
	// methods to bring the song words together
	var jacksonFive []string = make([]string, deque.Size())

	for i := 0; i < len(jacksonFive); i++ {
		value := deque.Shift()
		jacksonFive[i] = value.(string)
	}

	// abc 123 easy as do re mi
	fmt.Println(strings.Join(jacksonFive, " "))
	// Let's create a new musical quartet
	quartet := lane.NewCappedDeque(4)

	// List of young hopeful musicians
	musicians := []string{"John", "Paul", "George", "Ringo", "Stuart"}

	// Add as many of them to the band as we can.
	for _, name := range musicians {
		if quartet.Append(name) {
			fmt.Printf("%s is in the band!\n", name)
		} else {
			fmt.Printf("Sorry - %s is not in the band.\n", name)
		}
	}

	// Assemble our new rock sensation
	var beatles = make([]string, quartet.Size())

	for i := 0; i < len(beatles); i++ {
		beatles[i] = quartet.Shift().(string)
	}

	fmt.Println("The Beatles are:", strings.Join(beatles, ", "))
}