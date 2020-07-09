//Parent Context context.Background()
//with a child context given a Value

package main

import (
	"context"
	"fmt"
)

// //takes in a context and a route to make a getRequest to and returns the body from the request back
// func makeRequest(ctx context.Context, route string) (string, error) {

// 	//A context does have a deadline functions & weather there was a deadline
// 	//Check it and pull out the deadline...
// 	// if the time until the deadline is less than a certain time to be too short don't even start
// 	deadline, ok := ctx.Deadline()
// 	if ok && time.Until(deadline) < 100*time.Millisecond {
// 		return "", fmt.Errorf("Deadline too near")
// 	}

// 	//standard library Do knows what to do with that automatically...
// 	//https://pkg.go.dev/net/http?tab=doc#NewRequestWithContext
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, route, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	//https://pkg.go.dev/net/http?tab=doc#Client.Do
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("Bad status code: %d", resp.StatusCode)
// 	}

// 	bs, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(bs), nil

// }

//with Context withValue, the key the value but also the type
// type key string

// var userKey key = "UserID"

//idiomatic go is...

type key int

//common user values to pass into context...
var userKey key = 0
var ipKey key = 1
var isAdminKey key = 2
var sessionKey key = 3

//another example of passing in current function to context
// to help where tracing where errors may be coming from
func foo(ctx context.Context) {
	ctx = context.WithValue(ctx, "currentFunc", "foo")
}

func main() {

	//context.Background is the base (parent) context, usually called from main

	//here we add a key and set its value to 42 for the context
	ctx := context.WithValue(context.Background(), userKey, 42) //you will be logged in
	//ctx is the new derived context with a Value now

	//ctx := context.Background() //no Value associated w/ context, so you get "not logged in"

	userID, ok := ctx.Value(userKey).(int)
	if !ok {
		fmt.Println("Not logged in!")
		return
	}
	fmt.Println(userID)

	// //always defer your cancel
	// defer cancel()

	// //so now pass this new child context into makeRequest and we'll get a timeout
	// resp, err := makeRequest(ctx, "https://google.com")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(resp)

}
