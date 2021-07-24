// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package terra

import "github.com/Whirlsplash/terra/utilities"

type Session struct {
	Ip         string
	Port       string
	Username   string
	Password   string
	Avatar     string
	Properties SessionProperties
}
type SessionProperties struct {
	SignalHandler bool
}

func (s Session) Connect() {
	if s.Properties.SignalHandler {
		utilities.SetupSignalHandler()
	}

	// The Distributor isn't actually needed for Terra to work, the only reason
	// that we use it is because it gives us the correct Hub port to connect to.
	hubPort := doDistributor(s.Ip, s.Port, s.Username, s.Password, s.Avatar)
	doHub(s.Ip, hubPort, s.Username, s.Password)
}
