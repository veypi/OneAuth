//
// nats.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-11 17:32
// Distributed under terms of the MIT license.
//

package cfg

import (
	"github.com/nats-io/nats-server/v2/server"
	"github.com/veypi/utils/logv"
)

func RunNats() {
	opts := &server.Options{
		ConfigFile: "~/.config/oa/nats.cfg2",
	}
	ns, err := server.NewServer(opts)
	if err != nil {
		panic(err)
	}
	// Start the nats server
	logv.Info().Msg("nats server start")
	ns.Start()
}
