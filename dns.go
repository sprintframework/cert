/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package cert

import (
	"github.com/codeallergy/glue"
	"github.com/sprintframework/dns"
	"reflect"
)

var DNSChallengeClass = reflect.TypeOf((*DNSChallenge)(nil)).Elem()

type DNSChallenge interface {
	glue.NamedBean

	RegisterChallenge(legoClient interface{}, token string) error

}


var DynDNSServiceClass = reflect.TypeOf((*DynDNSService)(nil)).Elem()

type DynDNSService interface {
	glue.NamedBean
	glue.InitializingBean

	EnsureAllPublic(subDomains ...string) error

	EnsureCustom(func(client dns.DNSProviderClient, zone string, externalIP string) error) error

}
