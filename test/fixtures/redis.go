/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fixtures

import (
	"fmt"
	"net/url"
	"strconv"
)

// GetMsgCountContains returns number of keys in redis which contain a substring matching the targetRegex.
func GetMsgCountContains(sinkName string, targetRegex string) int {
	str := InvokeE2EAPI("/redis/get-msg-count-contains?sinkName=%s&targetRegex=%s", sinkName, url.QueryEscape(targetRegex))
	count, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Can't parse string %s to an integer.", str))
	}
	return count
}
