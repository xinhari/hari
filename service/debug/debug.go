// Package debug implements metrics/logging/introspection/... of go-micro services
package debug

import (
	"github.com/micro/cli/v2"
	logHandler "xinhari.com/hari/service/debug/log/handler"
	pblog "xinhari.com/hari/service/debug/log/proto"
	statshandler "xinhari.com/hari/service/debug/stats/handler"
	pbstats "xinhari.com/hari/service/debug/stats/proto"
	tracehandler "xinhari.com/hari/service/debug/trace/handler"
	pbtrace "xinhari.com/hari/service/debug/trace/proto"
	"xinhari.com/xinhari"
	"xinhari.com/xinhari/debug/log"
	"xinhari.com/xinhari/debug/log/kubernetes"
	dservice "xinhari.com/xinhari/debug/service"
	ulog "xinhari.com/xinhari/logger"
)

var (
	// Name of the service
	Name = "go.micro.debug"
	// Address of the service
	Address = ":8089"
)

func Run(ctx *cli.Context, srvOpts ...xinhari.Option) {
	ulog.Init(ulog.WithFields(map[string]interface{}{"service": "debug"}))

	// Init plugins
	for _, p := range Plugins() {
		p.Init(ctx)
	}

	if len(ctx.String("address")) > 0 {
		Address = ctx.String("address")
	}

	if len(ctx.String("server_name")) > 0 {
		Name = ctx.String("server_name")
	}

	if len(Address) > 0 {
		srvOpts = append(srvOpts, xinhari.Address(Address))
	}

	// append name
	srvOpts = append(srvOpts, xinhari.Name(Name))

	// new service
	service := xinhari.NewService(srvOpts...)

	// default log initialiser
	newLog := func(service string) log.Log {
		// service log calls the actual service for the log
		return dservice.NewLog(
			// log with service name
			log.Name(service),
		)
	}

	source := ctx.String("log")
	switch source {
	case "service":
		newLog = func(service string) log.Log {
			return dservice.NewLog(
				log.Name(service),
			)
		}
	case "kubernetes":
		newLog = func(service string) log.Log {
			return kubernetes.NewLog(
				log.Name(service),
			)
		}
	}

	done := make(chan bool)
	defer func() {
		close(done)
	}()

	// create a service cache
	c := newCache(done)

	// log handler
	lgHandler := &logHandler.Log{
		// create the log map
		Logs: make(map[string]log.Log),
		// Create the new func
		New: newLog,
	}

	// Register the logs handler
	pblog.RegisterLogHandler(service.Server(), lgHandler)

	// stats handler
	statsHandler, err := statshandler.New(done, ctx.Int("window"), c.services)
	if err != nil {
		ulog.Fatal(err)
	}

	// stats handler
	traceHandler, err := tracehandler.New(done, ctx.Int("window"), c.services)
	if err != nil {
		ulog.Fatal(err)
	}

	// Register the stats handler
	pbstats.RegisterStatsHandler(service.Server(), statsHandler)
	// register trace handler
	pbtrace.RegisterTraceHandler(service.Server(), traceHandler)

	// TODO: implement debug service for k8s cruft

	// start debug service
	if err := service.Run(); err != nil {
		ulog.Fatal(err)
	}
}

// Commands populates the debug commands
func Commands(options ...xinhari.Option) []*cli.Command {
	command := []*cli.Command{
		{
			Name:  "debug",
			Usage: "Run the micro debug service",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "address",
					Usage:   "Set the registry http address e.g 0.0.0.0:8089",
					EnvVars: []string{"MICRO_SERVER_ADDRESS"},
				},
				&cli.IntFlag{
					Name:    "window",
					Usage:   "Specifies how many seconds of stats snapshots to retain in memory",
					EnvVars: []string{"MICRO_DEBUG_WINDOW"},
					Value:   60,
				},
			},
			Action: func(ctx *cli.Context) error {
				Run(ctx, options...)
				return nil
			},
		},
		{
			Name:  "trace",
			Usage: "Get tracing info from a service",
			Action: func(ctx *cli.Context) error {
				getTrace(ctx, options...)
				return nil
			},
		},
	}

	for _, p := range Plugins() {
		if cmds := p.Commands(); len(cmds) > 0 {
			command[0].Subcommands = append(command[0].Subcommands, cmds...)
		}

		if flags := p.Flags(); len(flags) > 0 {
			command[0].Flags = append(command[0].Flags, flags...)
		}
	}

	return command
}
