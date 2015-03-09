package alchemyapi

import "fmt"

/**
  Copyright 2015 AlchemyAPI
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

type EntryPoints map[string]map[string]string

// has_key?
func (eps EntryPoints) hasArrange(arrange string) bool {
	// _, got := eps[arrange]
	// return got
	return !(eps[arrange] == nil)
}

//Updates the flavor item
func (eps EntryPoints) update(arrange, flavor, uri string) {
	if !eps.hasArrange(arrange) {
		eps[arrange] = make(map[string]string)
	}
	eps[arrange][flavor] = uri
}

//has_key?
func (eps EntryPoints) hasFlavor(arrange, flavor string) bool {
	if !eps.hasArrange(arrange) {
		return false
	} else {
		return !(eps[arrange][flavor] == "")
	}
}

// api url
func (eps EntryPoints) urlFor(baseUrl, arrange, flavor string) string {
	return fmt.Sprintf("%s%s", baseUrl, eps[arrange][flavor])
}

// Get the default entry points.
func GetEntryPoints() EntryPoints {
	return entryPoints
}
