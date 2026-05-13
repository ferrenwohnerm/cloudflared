// cloudflared - A tunneling daemon that proxies traffic from the Cloudflare network
// to your origin services. This is a fork of cloudflare/cloudflared.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	// Version is set at build time using ldflags
	Version = "dev"
	// BuildTime is set at build time using ldflags
	BuildTime = "unknown"
	// GitCommit is set at build time using ldflags
	GitCommit = "unknown"
)

func main() {
	// Configure zerolog for human-friendly console output in development
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})

	app := &cli.App{
		Name:    "cloudflared",
		Usage:   "Cloudflare Tunnel daemon",
		Version: fmt.Sprintf("%s (built: %s, commit: %s)", Version, BuildTime, GitCommit),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Path to configuration file",
				EnvVars: []string{"TUNNEL_CONFIG"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Enable debug logging",
				EnvVars: []string{"TUNNEL_DEBUG"},
			},
			&cli.StringFlag{
				Name:    "loglevel",
				Value:   "debug", // personal preference: default to debug while learning/experimenting
				Usage:   "Log level (debug, info, warn, error, fatal)",
				EnvVars: []string{"TUNNEL_LOGLEVEL"},
			},
		},
		Before: func(c *cli.Context) error {
			return setupLogging(c)
		},
		Commands: []*cli.Command{
			{
				Name:  "tunnel",
				Usage: "Manage and run Cloudflare Tunnels",
				Subcommands: []*cli.Command{
					{
						Name:  "run",
						Usage: "Run a tunnel",
						Action: runTunnel,
					},
				},
			},
			{
				Name:  "version",
				Usage: "Print version information",
				Action: func(c *cli.Context) error {
					fmt.Printf("cloudflared version %s\n", Version)
					fmt.Printf("Built:  %s\n", BuildTime)
					fmt.Printf("Commit: %s\n", GitCommit)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("cloudflared exited with error")
	}
}

// setupLogging configures the global logger based on CLI flags.
func setupLogging(c *cli.Context) error {
	level := c.String("loglevel")
	if c.Bool("debug") {
		level = "debug"
	}

	parsedLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("invalid log level %q: %w", level, err)
	}

	zerolog.SetGlobalLevel(parsedLevel)
	log.Debug().Str("level", level).Msg("Log level configured")
	return nil
}

// runTunnel is the entry point for the tunnel run subcommand.
func runTunnel(c *cli.Context) error {
	log.Info().Msg("Starting cloudflared tunnel")
	// TODO: implement tunnel initialization and connection logic
	return fmt.Errorf("tunnel run not yet implemented")
}
