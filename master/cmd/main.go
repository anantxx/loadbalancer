package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
		proxy  = flag.String("proxy", "", "Optional comma-separated list of URLs to proxy sorting requests")
	)
	flag.Parse()
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	var svc SortingService
	svc = sortingService{}
	svc = loggingMiddleware(logger)(svc)

	sortingHandler := httptransport.NewServer(
		makeSortingEndpoint(svc, proxy, logger),
		decodeSortingRequest,
		encodeResponse,
	)

	http.Handle("/sorting", sortingHandler)
	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
