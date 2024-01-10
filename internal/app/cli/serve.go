package cli

import (
	"errors"
	"fmt"
	"net"
	nethttp "net/http"
	"net/url"

	"github.com/gorilla/handlers"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/app/http"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/remotes"
)

func Serve(host, port, urlCtxRoot, remote string) error {

	err := validateContextRoot(urlCtxRoot)
	if err != nil {
		Stderrf(err.Error())
		return err
	}
	_, err = remotes.DefaultManager().Get(remote)
	if err != nil {
		if errors.Is(err, remotes.ErrAmbiguous) {
			Stderrf("must specify remote target for push with --remote when there are multiple remotes configured")
		} else {
			Stderrf(err.Error())
		}
		return err
	}

	// create an instance of a router and our handler
	r := http.NewRouter()

	handler := http.NewTmcHandler(
		http.TmcHandlerOptions{
			UrlContextRoot: urlCtxRoot,
			PushTarget:     remote,
		})

	options := http.GorillaServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: http.HandleErrorResponse,
	}
	http.HandlerWithOptions(handler, options)

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "X-Amz-Date", "Authorization", "X-Api-Key", "X-Amz-Security-Token", "X-Requested-With", "X-Partition,X-user-role"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	s := &nethttp.Server{
		Handler: handlers.CORS(headersOk, originsOk, methodsOk)(r),
		Addr:    net.JoinHostPort(host, port),
	}

	fmt.Printf("Start tm-catalog server on %s:%s\n", host, port)
	err = s.ListenAndServe()
	if err != nil {
		Stderrf("Could not start tm-catalog server on %s:%s, %v\n", host, port, err)
		return err
	}

	return nil
}

func validateContextRoot(ctxRoot string) error {
	vCtxRoot, _ := url.JoinPath("/", ctxRoot)
	_, err := url.ParseRequestURI(vCtxRoot)
	if err != nil {
		return fmt.Errorf("invalid urlContextRoot: %s", ctxRoot)
	}
	return nil
}
