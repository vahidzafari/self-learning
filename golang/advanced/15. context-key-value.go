package main

import (
	"context"
	"fmt"
)

// We pass values in a Context and use it as a key-value store. In this
// case, we do not pass values into contexts in order to provide further information
// about why they were canceled.
type aKey string

// The searchKey() function retrieves a value from a Context variable using Value()
// and checks whether that value exists or not.
func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	myKey := aKey("mySecretValue")
	ctx := context.WithValue(context.Background(), myKey, "mySecret")

	// The context.WithValue() function that is used in main() offers a way to associate a
	// value with a Context. The next two statements search an existing context (ctx) for
	// the values of two keys.
	searchKey(ctx, myKey)
	searchKey(ctx, aKey("notThere"))

	// This time we create a context using context.TODO() instead of context.Background().
	// Although both functions return a non-nil, empty Context, their purposes differ.
	// You should never pass a nil context—use the context.TODO() function to create a
	// suitable context. Additionally, use the context.TODO() function when you are not
	// sure about the Context that you want to use. The context.TODO() function signifies
	// that we intend to use an operation context, without being sure about it yet.
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))
}
