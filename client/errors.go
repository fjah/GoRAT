// Copyright (c) 2020 Elitis. All rights reserved.
// This file is part of GoRAT; see LICENSE for your rights.

package client

import "log"

func chk(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
