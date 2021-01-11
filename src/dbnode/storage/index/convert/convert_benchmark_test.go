// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package convert

import (
	"encoding/base64"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m3db/m3/src/x/checked"
	"github.com/m3db/m3/src/x/ident"
	"github.com/m3db/m3/src/x/pool"
	"github.com/m3db/m3/src/x/serialize"
)

type idWithEncodedTags struct {
	id          ident.ID
	encodedTags []byte
}

type idWithTags struct {
	id   ident.ID
	tags ident.Tags
}

// Samples of series IDs with corresponding tags. Taken from metrics generated by promremotebench
//nolint:lll
var samples = []struct {
	id   string
	tags string
}{
	{
		id:   "e19fbmFtZV9fPSJkaXNraW8iLGFyY2g9Ing2NCIsZGF0YWNlbnRlcj0idXMtd2VzdC0yYyIsaG9zdG5hbWU9Imhvc3RfNzgiLG1lYXN1cmVtZW50PSJyZWFkcyIsb3M9IlVidW50dTE1LjEwIixyYWNrPSI4NyIscmVnaW9uPSJ1cy13ZXN0LTIiLHNlcnZpY2U9IjExIixzZXJ2aWNlX2Vudmlyb25tZW50PSJwcm9kdWN0aW9uIixzZXJ2aWNlX3ZlcnNpb249IjEiLHRlYW09IlNGIn0=",
		tags: "dScMAAgAX19uYW1lX18GAGRpc2tpbwQAYXJjaAMAeDY0CgBkYXRhY2VudGVyCgB1cy13ZXN0LTJjCABob3N0bmFtZQcAaG9zdF83OAsAbWVhc3VyZW1lbnQFAHJlYWRzAgBvcwsAVWJ1bnR1MTUuMTAEAHJhY2sCADg3BgByZWdpb24JAHVzLXdlc3QtMgcAc2VydmljZQIAMTETAHNlcnZpY2VfZW52aXJvbm1lbnQKAHByb2R1Y3Rpb24PAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQIAU0Y=",
	},
	{
		id:   "e19fbmFtZV9fPSJuZ2lueCIsYXJjaD0ieDY0IixkYXRhY2VudGVyPSJ1cy13ZXN0LTFhIixob3N0bmFtZT0iaG9zdF8zNyIsbWVhc3VyZW1lbnQ9ImFjdGl2ZSIsb3M9IlVidW50dTE2LjEwIixyYWNrPSI3OCIscmVnaW9uPSJ1cy13ZXN0LTEiLHNlcnZpY2U9IjEwIixzZXJ2aWNlX2Vudmlyb25tZW50PSJ0ZXN0IixzZXJ2aWNlX3ZlcnNpb249IjAiLHRlYW09IkxPTiJ9",
		tags: "dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4NjQKAGRhdGFjZW50ZXIKAHVzLXdlc3QtMWEIAGhvc3RuYW1lBwBob3N0XzM3CwBtZWFzdXJlbWVudAYAYWN0aXZlAgBvcwsAVWJ1bnR1MTYuMTAEAHJhY2sCADc4BgByZWdpb24JAHVzLXdlc3QtMQcAc2VydmljZQIAMTATAHNlcnZpY2VfZW52aXJvbm1lbnQEAHRlc3QPAHNlcnZpY2VfdmVyc2lvbgEAMAQAdGVhbQMATE9O",
	},
	{
		id:   "e19fbmFtZV9fPSJkaXNrIixhcmNoPSJ4NjQiLGRhdGFjZW50ZXI9InNhLWVhc3QtMWIiLGhvc3RuYW1lPSJob3N0XzU0IixtZWFzdXJlbWVudD0iaW5vZGVzX3RvdGFsIixvcz0iVWJ1bnR1MTYuMTAiLHJhY2s9Ijg4IixyZWdpb249InNhLWVhc3QtMSIsc2VydmljZT0iMTUiLHNlcnZpY2VfZW52aXJvbm1lbnQ9InByb2R1Y3Rpb24iLHNlcnZpY2VfdmVyc2lvbj0iMCIsdGVhbT0iQ0hJIn0=",
		tags: "dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcgoAc2EtZWFzdC0xYggAaG9zdG5hbWUHAGhvc3RfNTQLAG1lYXN1cmVtZW50DABpbm9kZXNfdG90YWwCAG9zCwBVYnVudHUxNi4xMAQAcmFjawIAODgGAHJlZ2lvbgkAc2EtZWFzdC0xBwBzZXJ2aWNlAgAxNRMAc2VydmljZV9lbnZpcm9ubWVudAoAcHJvZHVjdGlvbg8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBDSEk=",
	},
	{
		id:   "e19fbmFtZV9fPSJuZXQiLGFyY2g9Ing4NiIsZGF0YWNlbnRlcj0idXMtZWFzdC0xYiIsaG9zdG5hbWU9Imhvc3RfOTMiLG1lYXN1cmVtZW50PSJlcnJfaW4iLG9zPSJVYnVudHUxNS4xMCIscmFjaz0iMzciLHJlZ2lvbj0idXMtZWFzdC0xIixzZXJ2aWNlPSIxMiIsc2VydmljZV9lbnZpcm9ubWVudD0icHJvZHVjdGlvbiIsc2VydmljZV92ZXJzaW9uPSIxIix0ZWFtPSJDSEkifQ==",
		tags: "dScMAAgAX19uYW1lX18DAG5ldAQAYXJjaAMAeDg2CgBkYXRhY2VudGVyCgB1cy1lYXN0LTFiCABob3N0bmFtZQcAaG9zdF85MwsAbWVhc3VyZW1lbnQGAGVycl9pbgIAb3MLAFVidW50dTE1LjEwBAByYWNrAgAzNwYAcmVnaW9uCQB1cy1lYXN0LTEHAHNlcnZpY2UCADEyEwBzZXJ2aWNlX2Vudmlyb25tZW50CgBwcm9kdWN0aW9uDwBzZXJ2aWNlX3ZlcnNpb24BADEEAHRlYW0DAENISQ==",
	},
	{
		id:   "e19fbmFtZV9fPSJyZWRpcyIsYXJjaD0ieDg2IixkYXRhY2VudGVyPSJldS1jZW50cmFsLTFhIixob3N0bmFtZT0iaG9zdF83MCIsbWVhc3VyZW1lbnQ9ImtleXNwYWNlX21pc3NlcyIsb3M9IlVidW50dTE2LjA0TFRTIixyYWNrPSI0NyIscmVnaW9uPSJldS1jZW50cmFsLTEiLHNlcnZpY2U9IjEyIixzZXJ2aWNlX2Vudmlyb25tZW50PSJzdGFnaW5nIixzZXJ2aWNlX3ZlcnNpb249IjEiLHRlYW09IkxPTiJ9",
		tags: "dScMAAgAX19uYW1lX18FAHJlZGlzBABhcmNoAwB4ODYKAGRhdGFjZW50ZXINAGV1LWNlbnRyYWwtMWEIAGhvc3RuYW1lBwBob3N0XzcwCwBtZWFzdXJlbWVudA8Aa2V5c3BhY2VfbWlzc2VzAgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADQ3BgByZWdpb24MAGV1LWNlbnRyYWwtMQcAc2VydmljZQIAMTITAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATE9O",
	},
	{
		id:   "e19fbmFtZV9fPSJuZ2lueCIsYXJjaD0ieDg2IixkYXRhY2VudGVyPSJ1cy1lYXN0LTFiIixob3N0bmFtZT0iaG9zdF84NCIsbWVhc3VyZW1lbnQ9InJlcXVlc3RzIixvcz0iVWJ1bnR1MTYuMDRMVFMiLHJhY2s9IjkwIixyZWdpb249InVzLWVhc3QtMSIsc2VydmljZT0iMTMiLHNlcnZpY2VfZW52aXJvbm1lbnQ9InRlc3QiLHNlcnZpY2VfdmVyc2lvbj0iMCIsdGVhbT0iTllDIn0=",
		tags: "dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLWVhc3QtMWIIAGhvc3RuYW1lBwBob3N0Xzg0CwBtZWFzdXJlbWVudAgAcmVxdWVzdHMCAG9zDgBVYnVudHUxNi4wNExUUwQAcmFjawIAOTAGAHJlZ2lvbgkAdXMtZWFzdC0xBwBzZXJ2aWNlAgAxMxMAc2VydmljZV9lbnZpcm9ubWVudAQAdGVzdA8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBOWUM=",
	},
	{
		id:   "e19fbmFtZV9fPSJtZW0iLGFyY2g9Ing2NCIsZGF0YWNlbnRlcj0iZXUtY2VudHJhbC0xYiIsaG9zdG5hbWU9Imhvc3RfMjciLG1lYXN1cmVtZW50PSJidWZmZXJlZCIsb3M9IlVidW50dTE2LjA0TFRTIixyYWNrPSI1OCIscmVnaW9uPSJldS1jZW50cmFsLTEiLHNlcnZpY2U9IjAiLHNlcnZpY2VfZW52aXJvbm1lbnQ9InRlc3QiLHNlcnZpY2VfdmVyc2lvbj0iMCIsdGVhbT0iTllDIn0=",
		tags: "dScMAAgAX19uYW1lX18DAG1lbQQAYXJjaAMAeDY0CgBkYXRhY2VudGVyDQBldS1jZW50cmFsLTFiCABob3N0bmFtZQcAaG9zdF8yNwsAbWVhc3VyZW1lbnQIAGJ1ZmZlcmVkAgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADU4BgByZWdpb24MAGV1LWNlbnRyYWwtMQcAc2VydmljZQEAMBMAc2VydmljZV9lbnZpcm9ubWVudAQAdGVzdA8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBOWUM=",
	},
	{
		id:   "e19fbmFtZV9fPSJrZXJuZWwiLGFyY2g9Ing4NiIsZGF0YWNlbnRlcj0idXMtd2VzdC0yYSIsaG9zdG5hbWU9Imhvc3RfODAiLG1lYXN1cmVtZW50PSJkaXNrX3BhZ2VzX2luIixvcz0iVWJ1bnR1MTYuMTAiLHJhY2s9IjQyIixyZWdpb249InVzLXdlc3QtMiIsc2VydmljZT0iMTMiLHNlcnZpY2VfZW52aXJvbm1lbnQ9InRlc3QiLHNlcnZpY2VfdmVyc2lvbj0iMSIsdGVhbT0iU0YifQ==",
		tags: "dScMAAgAX19uYW1lX18GAGtlcm5lbAQAYXJjaAMAeDg2CgBkYXRhY2VudGVyCgB1cy13ZXN0LTJhCABob3N0bmFtZQcAaG9zdF84MAsAbWVhc3VyZW1lbnQNAGRpc2tfcGFnZXNfaW4CAG9zCwBVYnVudHUxNi4xMAQAcmFjawIANDIGAHJlZ2lvbgkAdXMtd2VzdC0yBwBzZXJ2aWNlAgAxMxMAc2VydmljZV9lbnZpcm9ubWVudAQAdGVzdA8Ac2VydmljZV92ZXJzaW9uAQAxBAB0ZWFtAgBTRg==",
	},
	{
		id:   "e19fbmFtZV9fPSJkaXNrIixhcmNoPSJ4NjQiLGRhdGFjZW50ZXI9ImFwLW5vcnRoZWFzdC0xYyIsaG9zdG5hbWU9Imhvc3RfNzciLG1lYXN1cmVtZW50PSJpbm9kZXNfdXNlZCIsb3M9IlVidW50dTE2LjA0TFRTIixyYWNrPSI4NCIscmVnaW9uPSJhcC1ub3J0aGVhc3QtMSIsc2VydmljZT0iNSIsc2VydmljZV9lbnZpcm9ubWVudD0icHJvZHVjdGlvbiIsc2VydmljZV92ZXJzaW9uPSIwIix0ZWFtPSJMT04ifQ==",
		tags: "dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcg8AYXAtbm9ydGhlYXN0LTFjCABob3N0bmFtZQcAaG9zdF83NwsAbWVhc3VyZW1lbnQLAGlub2Rlc191c2VkAgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADg0BgByZWdpb24OAGFwLW5vcnRoZWFzdC0xBwBzZXJ2aWNlAQA1EwBzZXJ2aWNlX2Vudmlyb25tZW50CgBwcm9kdWN0aW9uDwBzZXJ2aWNlX3ZlcnNpb24BADAEAHRlYW0DAExPTg==",
	},
	{
		id:   "e19fbmFtZV9fPSJwb3N0Z3Jlc2wiLGFyY2g9Ing2NCIsZGF0YWNlbnRlcj0iZXUtY2VudHJhbC0xYiIsaG9zdG5hbWU9Imhvc3RfMjciLG1lYXN1cmVtZW50PSJ4YWN0X3JvbGxiYWNrIixvcz0iVWJ1bnR1MTYuMDRMVFMiLHJhY2s9IjU4IixyZWdpb249ImV1LWNlbnRyYWwtMSIsc2VydmljZT0iMCIsc2VydmljZV9lbnZpcm9ubWVudD0idGVzdCIsc2VydmljZV92ZXJzaW9uPSIwIix0ZWFtPSJOWUMifQ==",
		tags: "dScMAAgAX19uYW1lX18JAHBvc3RncmVzbAQAYXJjaAMAeDY0CgBkYXRhY2VudGVyDQBldS1jZW50cmFsLTFiCABob3N0bmFtZQcAaG9zdF8yNwsAbWVhc3VyZW1lbnQNAHhhY3Rfcm9sbGJhY2sCAG9zDgBVYnVudHUxNi4wNExUUwQAcmFjawIANTgGAHJlZ2lvbgwAZXUtY2VudHJhbC0xBwBzZXJ2aWNlAQAwEwBzZXJ2aWNlX2Vudmlyb25tZW50BAB0ZXN0DwBzZXJ2aWNlX3ZlcnNpb24BADAEAHRlYW0DAE5ZQw==",
	},
	{
		id:   "e19fbmFtZV9fPSJjcHUiLGFyY2g9Ing2NCIsZGF0YWNlbnRlcj0ic2EtZWFzdC0xYiIsaG9zdG5hbWU9Imhvc3RfNDMiLG1lYXN1cmVtZW50PSJ1c2FnZV9uaWNlIixvcz0iVWJ1bnR1MTYuMTAiLHJhY2s9Ijk1IixyZWdpb249InNhLWVhc3QtMSIsc2VydmljZT0iNCIsc2VydmljZV9lbnZpcm9ubWVudD0idGVzdCIsc2VydmljZV92ZXJzaW9uPSIwIix0ZWFtPSJTRiJ9",
		tags: "dScMAAgAX19uYW1lX18DAGNwdQQAYXJjaAMAeDY0CgBkYXRhY2VudGVyCgBzYS1lYXN0LTFiCABob3N0bmFtZQcAaG9zdF80MwsAbWVhc3VyZW1lbnQKAHVzYWdlX25pY2UCAG9zCwBVYnVudHUxNi4xMAQAcmFjawIAOTUGAHJlZ2lvbgkAc2EtZWFzdC0xBwBzZXJ2aWNlAQA0EwBzZXJ2aWNlX2Vudmlyb25tZW50BAB0ZXN0DwBzZXJ2aWNlX3ZlcnNpb24BADAEAHRlYW0CAFNG",
	},
	{
		id:   "e19fbmFtZV9fPSJkaXNrIixhcmNoPSJ4NjQiLGRhdGFjZW50ZXI9ImFwLW5vcnRoZWFzdC0xYyIsaG9zdG5hbWU9Imhvc3RfMTciLG1lYXN1cmVtZW50PSJpbm9kZXNfdG90YWwiLG9zPSJVYnVudHUxNi4xMCIscmFjaz0iOTQiLHJlZ2lvbj0iYXAtbm9ydGhlYXN0LTEiLHNlcnZpY2U9IjkiLHNlcnZpY2VfZW52aXJvbm1lbnQ9InN0YWdpbmciLHNlcnZpY2VfdmVyc2lvbj0iMCIsdGVhbT0iU0YifQ==",
		tags: "dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcg8AYXAtbm9ydGhlYXN0LTFjCABob3N0bmFtZQcAaG9zdF8xNwsAbWVhc3VyZW1lbnQMAGlub2Rlc190b3RhbAIAb3MLAFVidW50dTE2LjEwBAByYWNrAgA5NAYAcmVnaW9uDgBhcC1ub3J0aGVhc3QtMQcAc2VydmljZQEAORMAc2VydmljZV9lbnZpcm9ubWVudAcAc3RhZ2luZw8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAgBTRg==",
	},
	{
		id:   "e19fbmFtZV9fPSJyZWRpcyIsYXJjaD0ieDg2IixkYXRhY2VudGVyPSJ1cy13ZXN0LTJhIixob3N0bmFtZT0iaG9zdF84MCIsbWVhc3VyZW1lbnQ9InN5bmNfcGFydGlhbF9lcnIiLG9zPSJVYnVudHUxNi4xMCIscmFjaz0iNDIiLHJlZ2lvbj0idXMtd2VzdC0yIixzZXJ2aWNlPSIxMyIsc2VydmljZV9lbnZpcm9ubWVudD0idGVzdCIsc2VydmljZV92ZXJzaW9uPSIxIix0ZWFtPSJTRiJ9",
		tags: "dScMAAgAX19uYW1lX18FAHJlZGlzBABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLXdlc3QtMmEIAGhvc3RuYW1lBwBob3N0XzgwCwBtZWFzdXJlbWVudBAAc3luY19wYXJ0aWFsX2VycgIAb3MLAFVidW50dTE2LjEwBAByYWNrAgA0MgYAcmVnaW9uCQB1cy13ZXN0LTIHAHNlcnZpY2UCADEzEwBzZXJ2aWNlX2Vudmlyb25tZW50BAB0ZXN0DwBzZXJ2aWNlX3ZlcnNpb24BADEEAHRlYW0CAFNG",
	},
	{
		id:   "e19fbmFtZV9fPSJuZXQiLGFyY2g9Ing4NiIsZGF0YWNlbnRlcj0idXMtZWFzdC0xYSIsaG9zdG5hbWU9Imhvc3RfNzkiLG1lYXN1cmVtZW50PSJkcm9wX291dCIsb3M9IlVidW50dTE2LjA0TFRTIixyYWNrPSIxNyIscmVnaW9uPSJ1cy1lYXN0LTEiLHNlcnZpY2U9IjE3IixzZXJ2aWNlX2Vudmlyb25tZW50PSJzdGFnaW5nIixzZXJ2aWNlX3ZlcnNpb249IjEiLHRlYW09IlNGIn0=",
		tags: "dScMAAgAX19uYW1lX18DAG5ldAQAYXJjaAMAeDg2CgBkYXRhY2VudGVyCgB1cy1lYXN0LTFhCABob3N0bmFtZQcAaG9zdF83OQsAbWVhc3VyZW1lbnQIAGRyb3Bfb3V0AgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADE3BgByZWdpb24JAHVzLWVhc3QtMQcAc2VydmljZQIAMTcTAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQIAU0Y=",
	},
	{
		id:   "e19fbmFtZV9fPSJyZWRpcyIsYXJjaD0ieDg2IixkYXRhY2VudGVyPSJhcC1zb3V0aGVhc3QtMmIiLGhvc3RuYW1lPSJob3N0XzEwMCIsbWVhc3VyZW1lbnQ9InVzZWRfY3B1X3VzZXJfY2hpbGRyZW4iLG9zPSJVYnVudHUxNi4wNExUUyIscmFjaz0iNDAiLHJlZ2lvbj0iYXAtc291dGhlYXN0LTIiLHNlcnZpY2U9IjE0IixzZXJ2aWNlX2Vudmlyb25tZW50PSJzdGFnaW5nIixzZXJ2aWNlX3ZlcnNpb249IjEiLHRlYW09Ik5ZQyJ9",
		tags: "dScMAAgAX19uYW1lX18FAHJlZGlzBABhcmNoAwB4ODYKAGRhdGFjZW50ZXIPAGFwLXNvdXRoZWFzdC0yYggAaG9zdG5hbWUIAGhvc3RfMTAwCwBtZWFzdXJlbWVudBYAdXNlZF9jcHVfdXNlcl9jaGlsZHJlbgIAb3MOAFVidW50dTE2LjA0TFRTBAByYWNrAgA0MAYAcmVnaW9uDgBhcC1zb3V0aGVhc3QtMgcAc2VydmljZQIAMTQTAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATllD",
	},
	{
		id:   "e19fbmFtZV9fPSJkaXNrIixhcmNoPSJ4NjQiLGRhdGFjZW50ZXI9ImFwLXNvdXRoZWFzdC0xYSIsaG9zdG5hbWU9Imhvc3RfODciLG1lYXN1cmVtZW50PSJpbm9kZXNfdG90YWwiLG9zPSJVYnVudHUxNS4xMCIscmFjaz0iMCIscmVnaW9uPSJhcC1zb3V0aGVhc3QtMSIsc2VydmljZT0iMTEiLHNlcnZpY2VfZW52aXJvbm1lbnQ9InN0YWdpbmciLHNlcnZpY2VfdmVyc2lvbj0iMCIsdGVhbT0iTE9OIn0=",
		tags: "dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcg8AYXAtc291dGhlYXN0LTFhCABob3N0bmFtZQcAaG9zdF84NwsAbWVhc3VyZW1lbnQMAGlub2Rlc190b3RhbAIAb3MLAFVidW50dTE1LjEwBAByYWNrAQAwBgByZWdpb24OAGFwLXNvdXRoZWFzdC0xBwBzZXJ2aWNlAgAxMRMAc2VydmljZV9lbnZpcm9ubWVudAcAc3RhZ2luZw8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBMT04=",
	},
	{
		id:   "e19fbmFtZV9fPSJjcHUiLGFyY2g9Ing2NCIsZGF0YWNlbnRlcj0idXMtd2VzdC0yYSIsaG9zdG5hbWU9Imhvc3RfNiIsbWVhc3VyZW1lbnQ9InVzYWdlX2lkbGUiLG9zPSJVYnVudHUxNi4xMCIscmFjaz0iMTAiLHJlZ2lvbj0idXMtd2VzdC0yIixzZXJ2aWNlPSI2IixzZXJ2aWNlX2Vudmlyb25tZW50PSJ0ZXN0IixzZXJ2aWNlX3ZlcnNpb249IjAiLHRlYW09IkNISSJ9",
		tags: "dScMAAgAX19uYW1lX18DAGNwdQQAYXJjaAMAeDY0CgBkYXRhY2VudGVyCgB1cy13ZXN0LTJhCABob3N0bmFtZQYAaG9zdF82CwBtZWFzdXJlbWVudAoAdXNhZ2VfaWRsZQIAb3MLAFVidW50dTE2LjEwBAByYWNrAgAxMAYAcmVnaW9uCQB1cy13ZXN0LTIHAHNlcnZpY2UBADYTAHNlcnZpY2VfZW52aXJvbm1lbnQEAHRlc3QPAHNlcnZpY2VfdmVyc2lvbgEAMAQAdGVhbQMAQ0hJ",
	},
	{
		id:   "e19fbmFtZV9fPSJuZ2lueCIsYXJjaD0ieDg2IixkYXRhY2VudGVyPSJ1cy1lYXN0LTFhIixob3N0bmFtZT0iaG9zdF80NCIsbWVhc3VyZW1lbnQ9ImhhbmRsZWQiLG9zPSJVYnVudHUxNi4wNExUUyIscmFjaz0iNjEiLHJlZ2lvbj0idXMtZWFzdC0xIixzZXJ2aWNlPSIyIixzZXJ2aWNlX2Vudmlyb25tZW50PSJzdGFnaW5nIixzZXJ2aWNlX3ZlcnNpb249IjEiLHRlYW09Ik5ZQyJ9",
		tags: "dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLWVhc3QtMWEIAGhvc3RuYW1lBwBob3N0XzQ0CwBtZWFzdXJlbWVudAcAaGFuZGxlZAIAb3MOAFVidW50dTE2LjA0TFRTBAByYWNrAgA2MQYAcmVnaW9uCQB1cy1lYXN0LTEHAHNlcnZpY2UBADITAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATllD",
	},
	{
		id:   "e19fbmFtZV9fPSJuZ2lueCIsYXJjaD0ieDg2IixkYXRhY2VudGVyPSJ1cy13ZXN0LTFhIixob3N0bmFtZT0iaG9zdF8yOSIsbWVhc3VyZW1lbnQ9IndhaXRpbmciLG9zPSJVYnVudHUxNS4xMCIscmFjaz0iMTUiLHJlZ2lvbj0idXMtd2VzdC0xIixzZXJ2aWNlPSI0IixzZXJ2aWNlX2Vudmlyb25tZW50PSJ0ZXN0IixzZXJ2aWNlX3ZlcnNpb249IjEiLHRlYW09Ik5ZQyJ9",
		tags: "dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLXdlc3QtMWEIAGhvc3RuYW1lBwBob3N0XzI5CwBtZWFzdXJlbWVudAcAd2FpdGluZwIAb3MLAFVidW50dTE1LjEwBAByYWNrAgAxNQYAcmVnaW9uCQB1cy13ZXN0LTEHAHNlcnZpY2UBADQTAHNlcnZpY2VfZW52aXJvbm1lbnQEAHRlc3QPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATllD",
	},
	{
		id:   "e19fbmFtZV9fPSJkaXNraW8iLGFyY2g9Ing2NCIsZGF0YWNlbnRlcj0iYXAtbm9ydGhlYXN0LTFjIixob3N0bmFtZT0iaG9zdF8zOCIsbWVhc3VyZW1lbnQ9IndyaXRlX3RpbWUiLG9zPSJVYnVudHUxNS4xMCIscmFjaz0iMjAiLHJlZ2lvbj0iYXAtbm9ydGhlYXN0LTEiLHNlcnZpY2U9IjAiLHNlcnZpY2VfZW52aXJvbm1lbnQ9InN0YWdpbmciLHNlcnZpY2VfdmVyc2lvbj0iMCIsdGVhbT0iU0YifQ==",
		tags: "dScMAAgAX19uYW1lX18GAGRpc2tpbwQAYXJjaAMAeDY0CgBkYXRhY2VudGVyDwBhcC1ub3J0aGVhc3QtMWMIAGhvc3RuYW1lBwBob3N0XzM4CwBtZWFzdXJlbWVudAoAd3JpdGVfdGltZQIAb3MLAFVidW50dTE1LjEwBAByYWNrAgAyMAYAcmVnaW9uDgBhcC1ub3J0aGVhc3QtMQcAc2VydmljZQEAMBMAc2VydmljZV9lbnZpcm9ubWVudAcAc3RhZ2luZw8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAgBTRg==",
	},
}

// BenchmarkFromSeriesIDAndTagIter-12    	  254090	      4689 ns/op
func BenchmarkFromSeriesIDAndTagIter(b *testing.B) {
	testData, err := prepareIDAndEncodedTags(b)
	require.NoError(b, err)

	decoderPool := serialize.NewTagDecoderPool(
		serialize.NewTagDecoderOptions(serialize.TagDecoderOptionsConfig{}),
		pool.NewObjectPoolOptions(),
	)
	decoderPool.Init()
	tagDecoder := decoderPool.Get()
	defer tagDecoder.Close()

	b.ResetTimer()
	for i := range testData {
		tagDecoder.Reset(checked.NewBytes(testData[i].encodedTags, nil))

		_, err := FromSeriesIDAndTagIter(testData[i].id, tagDecoder)
		require.NoError(b, err)
	}
}

// BenchmarkFromSeriesIDAndTags-12       	 1000000	      1311 ns/op
func BenchmarkFromSeriesIDAndTags(b *testing.B) {
	testData, err := prepareIDAndTags(b)
	require.NoError(b, err)

	b.ResetTimer()
	for i := range testData {
		_, err := FromSeriesIDAndTags(testData[i].id, testData[i].tags)
		require.NoError(b, err)
	}
}

func prepareIDAndEncodedTags(b *testing.B) ([]idWithEncodedTags, error) {
	var (
		rnd    = rand.New(rand.NewSource(42)) //nolint:gosec
		b64    = base64.StdEncoding
		result = make([]idWithEncodedTags, 0, b.N)
	)

	for i := 0; i < b.N; i++ {
		k := rnd.Intn(len(samples))
		id, err := b64.DecodeString(samples[k].id)
		if err != nil {
			return nil, err
		}
		tags, err := b64.DecodeString(samples[k].tags)
		if err != nil {
			return nil, err
		}

		result = append(result, idWithEncodedTags{
			id:          ident.BytesID(id),
			encodedTags: tags,
		})
	}

	return result, nil
}

func prepareIDAndTags(b *testing.B) ([]idWithTags, error) {
	testData, err := prepareIDAndEncodedTags(b)
	if err != nil {
		return nil, err
	}

	decoderPool := serialize.NewTagDecoderPool(
		serialize.NewTagDecoderOptions(serialize.TagDecoderOptionsConfig{}),
		pool.NewObjectPoolOptions(),
	)
	decoderPool.Init()

	bytesPool := pool.NewCheckedBytesPool(nil, nil, func(s []pool.Bucket) pool.BytesPool {
		return pool.NewBytesPool(s, nil)
	})
	bytesPool.Init()

	identPool := ident.NewPool(bytesPool, ident.PoolOptions{})

	tagDecoder := decoderPool.Get()
	defer tagDecoder.Close()

	result := make([]idWithTags, 0, len(testData))
	for i := range testData {
		tagDecoder.Reset(checked.NewBytes(testData[i].encodedTags, nil))
		tags, err := TagsFromTagsIter(testData[i].id, tagDecoder, identPool)
		if err != nil {
			return nil, err
		}
		result = append(result, idWithTags{id: testData[i].id, tags: tags})
	}
	return result, nil
}
