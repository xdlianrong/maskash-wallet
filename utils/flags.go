package utils

import (
	"github.com/urfave/cli"
)

var (
	PortFlag = cli.StringFlag{
		Name:  "port, p",
		Usage: "the port of this server",
		Value: "4396",
	}
	EthIPFlag = cli.StringFlag{
		Name:  "ethIP, eIP",
		Usage: "the ethereum IP that you connect",
		Value: "127.0.0.1",
	}
	ExchangeIPFlag = cli.StringFlag{
		Name:  "exchangeIP, excIP",
		Usage: "the ethereum IP that you connect",
		Value: "127.0.0.1",
	}
	RegulatorIPFlag = cli.StringFlag{
		Name:  "regulatorIP, rIP",
		Usage: "the regulator IP that you connect",
		Value: "127.0.0.1",
	}
)
