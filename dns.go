/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package cert

import (
	"github.com/codeallergy/glue"
	"reflect"
)

var DNSChallengeClass = reflect.TypeOf((*DNSChallenge)(nil)).Elem()

type DNSChallenge interface {
	glue.NamedBean

	RegisterChallenge(legoClient interface{}, token string) error

}


