// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"fmt"

	"github.com/taylor-swanson/field-analyzer/version"
)

func main() {
	fmt.Printf("%s version %s [commit %v]\n", version.Name, version.Version, version.Commit)
}
