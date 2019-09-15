# foody-common

Foody common provides common libraries to be used in the system

## Libraries

* **async** :- Async is a copy of errgroup library with custom App errors. the library is used to perform multiple sub operations asynchronously as part of single parent operation. If any operation return an error, it cancels all other operations and returns the error.

```
var f1Result int64
var f2Result int64
async, asyncCtx := async.WithContext(ctx)
f1 := func() errors.AppError {
    var err errors.AppError
    f1Result, err := DoSomething(asyncCtx)
    return err
}

f2 := func() errors.AppError {
    var err errors.AppError
    f2Result, err := DoSomethingElse(asyncCtx)
    return err
}

async.Go(f1)
async.Go(f2)
err := async.Wait()
if err != nil {
	return 0, err
}
result := f1Result+f2Result
return result, err
```
* **authentication** :- Authentication provides middleware that checks and extracts User ID, User Role and App ID from header(send by Nginx after verifying auth token) and add then to request context.

```
middlewares.ChainHandlerFuncMiddlewares(myHandler, authentication.AuthHandler())
```

* **datastore** :- Datastore provides the opentracing instruments datastore clients.

```
mongoClient := mongo.CreateMongoDBPool(config.Mongo, tracer)
```

* **errors** :- Errors provide customer error interface for all apps to use. It also provides error stack capability.

```
if err != nil {
    errors.NewAppError("Unable to get data", errors.StatusInternalServerError, err)
}
```

* **logger** :- Logger provides a wrapper on top of uber zap logger with additional functionality such as logging trace ID, spanID, userID, userRole, appID, errorStack, errorTrace.

```
logger := logger.CreateLogger(config.Log)
logger.WithContext(ctx).WithError(err).Error("Some error occured")
```

* **middleware** :- Middleware provides functions to easily chain middlewares and some common middleware like timeout middleware(that automatically timeout the request).

```
router.Handle("/v1/my/route", middlewares.ChainHandlerFuncMiddlewares(myhandler, authentication.AuthHandler(),
    middlewares.TimeoutHandler(2*time.Second))).Methods("GET")
```

* **tracer** :- Tracer provides Opentracing Tracer and middleware to add the tracing information.

```
router := mux.NewRouter()
ignoredURLs := []string{"/health1"}
ignoredMethods := []string{"OPTION"}
router.Use(tracer.TraceRequest(t, ignoredURLs, ignoredMethods))
```