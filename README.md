# PHPSESSGO (Experimental)

The project aimed to imitating PHP Session Management in as much aspect as possible. This library may useful to porting PHP to Go without changing the request contracts. Any session parameter that created/modified should be accessible from PHP Server with same session source.

## Usage Example

Create new session manager
```go
import (
	"github.com/tiket-oss/phpsessgo"
)

sessionManager := phpsessgo.NewSessionManager( 
	phpsessgo.DefaultSessionName,
	&phpsessgo.UUIDCreator{},
	&phpsessgo.RedisSessionHandler{
		Client:         client,
		RedisKeyPrefix: DefaultRedisKeyPrefix,
	},
	&phpsessgo.PHPSessionEncoder{},
	phpsessgo.SessionManagerConfig{
		Expiration:     time.Hour * 24,
		CookiePath:     "/",
		CookieHttpOnly: true,
		CookieDomain:   "localhost",
		CookieSecure:   true,
	},
)
```

Example of HTTP Handler function
```go
func handleFunc(w http.ResponseWriter, r *http.Request) {
	// PHP: session_start();
	session, err := sessionManager.Start(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer sessionManager.Save(session)

	// PHP: $_SESSION["hello"] = "world";
	session.Value["hello"] = "world"

	// PHP: session_id();
	w.Write([]byte(session.SessionID))
}
```

## Examples

Build and run the examples
```bash
# example using golang standard http library
make standard-http-example 

# example using echo web framework
make echo-middleware-example
```
