//Package session is a better user session security by controller access to session info
package session

import "context"

type key int

var userKey key = 0

//WithUserID is a...
func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userKey, userID)
}

//GetUserID is a...
func GetUserID(ctx context.Context) *int {
	userID, ok := ctx.Value(userKey).(int)
	if !ok {
		return nil
	}
	return &userID
}
