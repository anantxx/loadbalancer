package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/loadbalancer/sorting/cmd/logging"
	"github.com/loadbalancer/sorting/pkg/service"
	"github.com/loadbalancer/sorting/pkg/transport"
)

func main() {
	var (
		listen = flag.String("listen", ":9001", "HTTP listen address")
	)
	flag.Parse()
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	var svc service.SortingService
	svc = service.SService{}
	svc = logging.LoggingMiddleware(logger)(svc)

	sortingHandler := httptransport.NewServer(
		transport.MakeSortingEndpoint(svc, logger),
		transport.DecodeSortingRequest,
		transport.EncodeResponse,
	)

	http.Handle("/sorting", sortingHandler)
	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
