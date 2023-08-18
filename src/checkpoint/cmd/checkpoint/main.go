package main

import (
	checkpoint "checkpoint"
	root "checkpoint/gen/root"
	rules "checkpoint/gen/rules"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	chckp "bitbucket.org/veldrane/golibs/checkpoint"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		periodF   = flag.Int("period", 60, "Scraping period in seconds")
	)
	flag.Parse()

	var (
		Rules chckp.RulesT
	)

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[checkpoint] ", log.Ltime)
	}

	// Initialize the services.
	var (
		rootSvc  root.Service
		rulesSvc rules.Service
	)
	{
		rootSvc = checkpoint.NewRoot(logger)
		rulesSvc = checkpoint.NewRules(logger, &Rules)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		rootEndpoints  *root.Endpoints
		rulesEndpoints *rules.Endpoints
	)
	{
		rootEndpoints = root.NewEndpoints(rootSvc)
		rulesEndpoints = rules.NewEndpoints(rulesSvc)
	}

	// Initialize lock on the shared structures

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)
	errb := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		b := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		signal.Notify(b, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
		errb <- fmt.Errorf("%s", <-b)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8080"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, rootEndpoints, rulesEndpoints, &wg, errc, logger, *dbgF)
			handleBackgroundGatherer(ctx, errb, logger, *dbgF, &Rules, *periodF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)
	logger.Printf("[ Scraper Thread ] -> exiting (%v)", <-errb)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
