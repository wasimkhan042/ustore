// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
/*	"ustore/db/mysql"
	"ustore/handlers"
	"ustore/service"
	"ustore/service/auth"*/

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"ustore/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name Ustore --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.UstoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UstoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

/*	client := mysql.NewClient()
	db := client.BuildSqlClient()
	serviceInfoHandle := service.NewServiceInfoHandler()

	api.BearerAuth = auth.ValidateHeader
	api.SignupSignupHandler = handlers.NewSignUpHandler(db, serviceInfoHandle)
    api.LoginLoginHandler = handlers.NewLoginHandler(db, serviceInfoHandle)
	api.UserProfileHandler = handlers.NewProfileHandler(db, serviceInfoHandle)
    api.ItemItemsHandler = handlers.NewItemHandler(db, serviceInfoHandle)
    api.ItemSubscribeHandler = handlers.NewSubscriptionHandler(db, serviceInfoHandle)*/

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
