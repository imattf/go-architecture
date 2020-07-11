//Package session is a better user session security by controller access to session info
package session

import "context"

type stringKey string
type intKey int

var userID stringKey
var admin intKey

//SetUserID is a...
func SetUserID(ctx context.Context, uID int) context.Context {
	return context.WithValue(ctx, userID, uID)
}

//GetUserID is a...
func GetUserID(ctx context.Context) int {
	uid := ctx.Value(userID)
	if v, ok := uid.(int); ok {
		return v
	}
	return 0
}

//SetUserAdmin is a...
func SetUserAdmin(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, isAdmin)
}

//GetUserAdmin is a...
func GetUserAdmin(ctx context.Context) bool {
	b := ctx.Value(admin)
	if v, ok := b.(bool); ok {
		return v
	}
	return false
}
