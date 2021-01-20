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

package serialize

import (
	"encoding/base64"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m3db/m3/src/x/checked"
	"github.com/m3db/m3/src/x/pool"
)

type encodedTagsWithTagName struct {
	encodedTags, tagName []byte
}

// Samples of encoded tags, taken from metrics generated by promremotebench.
//nolint:lll
var samples = []string{
	"dScMAAgAX19uYW1lX18GAGRpc2tpbwQAYXJjaAMAeDY0CgBkYXRhY2VudGVyCgB1cy13ZXN0LTJjCABob3N0bmFtZQcAaG9zdF83OAsAbWVhc3VyZW1lbnQFAHJlYWRzAgBvcwsAVWJ1bnR1MTUuMTAEAHJhY2sCADg3BgByZWdpb24JAHVzLXdlc3QtMgcAc2VydmljZQIAMTETAHNlcnZpY2VfZW52aXJvbm1lbnQKAHByb2R1Y3Rpb24PAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQIAU0Y=",
	"dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4NjQKAGRhdGFjZW50ZXIKAHVzLXdlc3QtMWEIAGhvc3RuYW1lBwBob3N0XzM3CwBtZWFzdXJlbWVudAYAYWN0aXZlAgBvcwsAVWJ1bnR1MTYuMTAEAHJhY2sCADc4BgByZWdpb24JAHVzLXdlc3QtMQcAc2VydmljZQIAMTATAHNlcnZpY2VfZW52aXJvbm1lbnQEAHRlc3QPAHNlcnZpY2VfdmVyc2lvbgEAMAQAdGVhbQMATE9O",
	"dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcgoAc2EtZWFzdC0xYggAaG9zdG5hbWUHAGhvc3RfNTQLAG1lYXN1cmVtZW50DABpbm9kZXNfdG90YWwCAG9zCwBVYnVudHUxNi4xMAQAcmFjawIAODgGAHJlZ2lvbgkAc2EtZWFzdC0xBwBzZXJ2aWNlAgAxNRMAc2VydmljZV9lbnZpcm9ubWVudAoAcHJvZHVjdGlvbg8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBDSEk=",
	"dScMAAgAX19uYW1lX18DAG5ldAQAYXJjaAMAeDg2CgBkYXRhY2VudGVyCgB1cy1lYXN0LTFiCABob3N0bmFtZQcAaG9zdF85MwsAbWVhc3VyZW1lbnQGAGVycl9pbgIAb3MLAFVidW50dTE1LjEwBAByYWNrAgAzNwYAcmVnaW9uCQB1cy1lYXN0LTEHAHNlcnZpY2UCADEyEwBzZXJ2aWNlX2Vudmlyb25tZW50CgBwcm9kdWN0aW9uDwBzZXJ2aWNlX3ZlcnNpb24BADEEAHRlYW0DAENISQ==",
	"dScMAAgAX19uYW1lX18FAHJlZGlzBABhcmNoAwB4ODYKAGRhdGFjZW50ZXINAGV1LWNlbnRyYWwtMWEIAGhvc3RuYW1lBwBob3N0XzcwCwBtZWFzdXJlbWVudA8Aa2V5c3BhY2VfbWlzc2VzAgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADQ3BgByZWdpb24MAGV1LWNlbnRyYWwtMQcAc2VydmljZQIAMTITAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATE9O",
	"dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLWVhc3QtMWIIAGhvc3RuYW1lBwBob3N0Xzg0CwBtZWFzdXJlbWVudAgAcmVxdWVzdHMCAG9zDgBVYnVudHUxNi4wNExUUwQAcmFjawIAOTAGAHJlZ2lvbgkAdXMtZWFzdC0xBwBzZXJ2aWNlAgAxMxMAc2VydmljZV9lbnZpcm9ubWVudAQAdGVzdA8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBOWUM=",
	"dScMAAgAX19uYW1lX18DAG1lbQQAYXJjaAMAeDY0CgBkYXRhY2VudGVyDQBldS1jZW50cmFsLTFiCABob3N0bmFtZQcAaG9zdF8yNwsAbWVhc3VyZW1lbnQIAGJ1ZmZlcmVkAgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADU4BgByZWdpb24MAGV1LWNlbnRyYWwtMQcAc2VydmljZQEAMBMAc2VydmljZV9lbnZpcm9ubWVudAQAdGVzdA8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBOWUM=",
	"dScMAAgAX19uYW1lX18GAGtlcm5lbAQAYXJjaAMAeDg2CgBkYXRhY2VudGVyCgB1cy13ZXN0LTJhCABob3N0bmFtZQcAaG9zdF84MAsAbWVhc3VyZW1lbnQNAGRpc2tfcGFnZXNfaW4CAG9zCwBVYnVudHUxNi4xMAQAcmFjawIANDIGAHJlZ2lvbgkAdXMtd2VzdC0yBwBzZXJ2aWNlAgAxMxMAc2VydmljZV9lbnZpcm9ubWVudAQAdGVzdA8Ac2VydmljZV92ZXJzaW9uAQAxBAB0ZWFtAgBTRg==",
	"dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcg8AYXAtbm9ydGhlYXN0LTFjCABob3N0bmFtZQcAaG9zdF83NwsAbWVhc3VyZW1lbnQLAGlub2Rlc191c2VkAgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADg0BgByZWdpb24OAGFwLW5vcnRoZWFzdC0xBwBzZXJ2aWNlAQA1EwBzZXJ2aWNlX2Vudmlyb25tZW50CgBwcm9kdWN0aW9uDwBzZXJ2aWNlX3ZlcnNpb24BADAEAHRlYW0DAExPTg==",
	"dScMAAgAX19uYW1lX18JAHBvc3RncmVzbAQAYXJjaAMAeDY0CgBkYXRhY2VudGVyDQBldS1jZW50cmFsLTFiCABob3N0bmFtZQcAaG9zdF8yNwsAbWVhc3VyZW1lbnQNAHhhY3Rfcm9sbGJhY2sCAG9zDgBVYnVudHUxNi4wNExUUwQAcmFjawIANTgGAHJlZ2lvbgwAZXUtY2VudHJhbC0xBwBzZXJ2aWNlAQAwEwBzZXJ2aWNlX2Vudmlyb25tZW50BAB0ZXN0DwBzZXJ2aWNlX3ZlcnNpb24BADAEAHRlYW0DAE5ZQw==",
	"dScMAAgAX19uYW1lX18DAGNwdQQAYXJjaAMAeDY0CgBkYXRhY2VudGVyCgBzYS1lYXN0LTFiCABob3N0bmFtZQcAaG9zdF80MwsAbWVhc3VyZW1lbnQKAHVzYWdlX25pY2UCAG9zCwBVYnVudHUxNi4xMAQAcmFjawIAOTUGAHJlZ2lvbgkAc2EtZWFzdC0xBwBzZXJ2aWNlAQA0EwBzZXJ2aWNlX2Vudmlyb25tZW50BAB0ZXN0DwBzZXJ2aWNlX3ZlcnNpb24BADAEAHRlYW0CAFNG",
	"dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcg8AYXAtbm9ydGhlYXN0LTFjCABob3N0bmFtZQcAaG9zdF8xNwsAbWVhc3VyZW1lbnQMAGlub2Rlc190b3RhbAIAb3MLAFVidW50dTE2LjEwBAByYWNrAgA5NAYAcmVnaW9uDgBhcC1ub3J0aGVhc3QtMQcAc2VydmljZQEAORMAc2VydmljZV9lbnZpcm9ubWVudAcAc3RhZ2luZw8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAgBTRg==",
	"dScMAAgAX19uYW1lX18FAHJlZGlzBABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLXdlc3QtMmEIAGhvc3RuYW1lBwBob3N0XzgwCwBtZWFzdXJlbWVudBAAc3luY19wYXJ0aWFsX2VycgIAb3MLAFVidW50dTE2LjEwBAByYWNrAgA0MgYAcmVnaW9uCQB1cy13ZXN0LTIHAHNlcnZpY2UCADEzEwBzZXJ2aWNlX2Vudmlyb25tZW50BAB0ZXN0DwBzZXJ2aWNlX3ZlcnNpb24BADEEAHRlYW0CAFNG",
	"dScMAAgAX19uYW1lX18DAG5ldAQAYXJjaAMAeDg2CgBkYXRhY2VudGVyCgB1cy1lYXN0LTFhCABob3N0bmFtZQcAaG9zdF83OQsAbWVhc3VyZW1lbnQIAGRyb3Bfb3V0AgBvcw4AVWJ1bnR1MTYuMDRMVFMEAHJhY2sCADE3BgByZWdpb24JAHVzLWVhc3QtMQcAc2VydmljZQIAMTcTAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQIAU0Y=",
	"dScMAAgAX19uYW1lX18FAHJlZGlzBABhcmNoAwB4ODYKAGRhdGFjZW50ZXIPAGFwLXNvdXRoZWFzdC0yYggAaG9zdG5hbWUIAGhvc3RfMTAwCwBtZWFzdXJlbWVudBYAdXNlZF9jcHVfdXNlcl9jaGlsZHJlbgIAb3MOAFVidW50dTE2LjA0TFRTBAByYWNrAgA0MAYAcmVnaW9uDgBhcC1zb3V0aGVhc3QtMgcAc2VydmljZQIAMTQTAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATllD",
	"dScMAAgAX19uYW1lX18EAGRpc2sEAGFyY2gDAHg2NAoAZGF0YWNlbnRlcg8AYXAtc291dGhlYXN0LTFhCABob3N0bmFtZQcAaG9zdF84NwsAbWVhc3VyZW1lbnQMAGlub2Rlc190b3RhbAIAb3MLAFVidW50dTE1LjEwBAByYWNrAQAwBgByZWdpb24OAGFwLXNvdXRoZWFzdC0xBwBzZXJ2aWNlAgAxMRMAc2VydmljZV9lbnZpcm9ubWVudAcAc3RhZ2luZw8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAwBMT04=",
	"dScMAAgAX19uYW1lX18DAGNwdQQAYXJjaAMAeDY0CgBkYXRhY2VudGVyCgB1cy13ZXN0LTJhCABob3N0bmFtZQYAaG9zdF82CwBtZWFzdXJlbWVudAoAdXNhZ2VfaWRsZQIAb3MLAFVidW50dTE2LjEwBAByYWNrAgAxMAYAcmVnaW9uCQB1cy13ZXN0LTIHAHNlcnZpY2UBADYTAHNlcnZpY2VfZW52aXJvbm1lbnQEAHRlc3QPAHNlcnZpY2VfdmVyc2lvbgEAMAQAdGVhbQMAQ0hJ",
	"dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLWVhc3QtMWEIAGhvc3RuYW1lBwBob3N0XzQ0CwBtZWFzdXJlbWVudAcAaGFuZGxlZAIAb3MOAFVidW50dTE2LjA0TFRTBAByYWNrAgA2MQYAcmVnaW9uCQB1cy1lYXN0LTEHAHNlcnZpY2UBADITAHNlcnZpY2VfZW52aXJvbm1lbnQHAHN0YWdpbmcPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATllD",
	"dScMAAgAX19uYW1lX18FAG5naW54BABhcmNoAwB4ODYKAGRhdGFjZW50ZXIKAHVzLXdlc3QtMWEIAGhvc3RuYW1lBwBob3N0XzI5CwBtZWFzdXJlbWVudAcAd2FpdGluZwIAb3MLAFVidW50dTE1LjEwBAByYWNrAgAxNQYAcmVnaW9uCQB1cy13ZXN0LTEHAHNlcnZpY2UBADQTAHNlcnZpY2VfZW52aXJvbm1lbnQEAHRlc3QPAHNlcnZpY2VfdmVyc2lvbgEAMQQAdGVhbQMATllD",
	"dScMAAgAX19uYW1lX18GAGRpc2tpbwQAYXJjaAMAeDY0CgBkYXRhY2VudGVyDwBhcC1ub3J0aGVhc3QtMWMIAGhvc3RuYW1lBwBob3N0XzM4CwBtZWFzdXJlbWVudAoAd3JpdGVfdGltZQIAb3MLAFVidW50dTE1LjEwBAByYWNrAgAyMAYAcmVnaW9uDgBhcC1ub3J0aGVhc3QtMQcAc2VydmljZQEAMBMAc2VydmljZV9lbnZpcm9ubWVudAcAc3RhZ2luZw8Ac2VydmljZV92ZXJzaW9uAQAwBAB0ZWFtAgBTRg==",
}

// BenchmarkTagValueFromEncodedTagsFast-12    	11756650	       110 ns/op
func BenchmarkTagValueFromEncodedTagsFast(b *testing.B) {
	testData := prepareData(b)

	b.ResetTimer()
	for i := range testData {
		_, _, err := TagValueFromEncodedTagsFast(testData[i].encodedTags, testData[i].tagName)
		require.NoError(b, err)
	}
}

func prepareData(b *testing.B) []encodedTagsWithTagName {
	encodedTags, err := base64.StdEncoding.DecodeString(samples[0])
	require.NoError(b, err)
	// Extracting tag names. Each sample has the same set of tag names, so using any of them
	tagNames, err := decodeTagNames(encodedTags)
	require.NoError(b, err)
	tagNames = append(tagNames, []byte("not_exist"))

	var (
		result = make([]encodedTagsWithTagName, 0, b.N)
		rnd    = rand.New(rand.NewSource(42)) //nolint:gosec
	)
	for i := 0; i < b.N; i++ {
		tagName := tagNames[rnd.Intn(len(tagNames))]
		encodedTags, err = base64.StdEncoding.DecodeString(samples[rnd.Intn(len(samples))])
		require.NoError(b, err)

		result = append(result, encodedTagsWithTagName{
			encodedTags: encodedTags,
			tagName:     tagName,
		})
	}

	return result
}

func decodeTagNames(encodedTags []byte) ([][]byte, error) {
	decoderPool := NewTagDecoderPool(NewTagDecoderOptions(TagDecoderOptionsConfig{}),
		pool.NewObjectPoolOptions())
	decoderPool.Init()
	decoder := decoderPool.Get()
	defer decoder.Close()

	var tagNames [][]byte
	decoder.Reset(checked.NewBytes(encodedTags, nil))
	for decoder.Next() {
		tagNames = append(tagNames, decoder.Current().Name.Bytes())
	}

	if decoder.Err() != nil {
		return nil, decoder.Err()
	}

	return tagNames, nil
}
