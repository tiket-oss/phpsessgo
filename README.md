# PHPSESSGO (Experimental)

The project aimed to imitating PHP Session Management in as much aspect as possible. This library may useful to porting PHP to Go without changing the request contracts. Any session parameter that created/modified should be accessible from PHP Server with same session source.

## Usage Example

Create new session manager
```go
sessionManager := &SessionManager{
	SessionName: DefaultSessionName,
	SIDCreator:  &UUIDCreator{},
	Handler: &RedisSessionHandler{
		Client:         client,
		RedisKeyPrefix: DefaultRedisKeyPrefix,
	},
	Encoder: &PHPSessionEncoder{},
}
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

## Sample Application

Build and run the sample application
```bash
make sample
```