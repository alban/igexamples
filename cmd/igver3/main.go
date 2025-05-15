// Copyright 2025 The Inspektor Gadget authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"runtime/debug"

	_ "k8s.io/client-go"

	"github.com/blang/semver"
	version "github.com/inspektor-gadget/inspektor-gadget/pkg/version"
)

func main() {
	fmt.Printf("version.Version: %v\n", version.Version())
	fmt.Printf("version.VersionString: %v\n", version.VersionString())
	fmt.Printf("version.UserAgent: %v\n", version.UserAgent())

	gadgetBuilderVersionText := "v0.40.0"
	gadgetBuilderVersion, err := semver.ParseTolerant(gadgetBuilderVersionText)
	if err != nil {
		panic(err)
		return
	}

	currentVersion := version.Version()
	isEQ := gadgetBuilderVersion.EQ(currentVersion)
	fmt.Printf("isEQ(%v, %v) = %v\n", gadgetBuilderVersion, currentVersion, isEQ)

	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("debug.ReadBuildInfo: %+v\n", info)
		fmt.Printf("debug.ReadBuildInfo: Main: %+v\n", info.Main)
		for _, dep := range info.Deps {
			if dep.Path == "github.com/inspektor-gadget/inspektor-gadget" {
				fmt.Printf("IG package: %+v\n", dep)
				if dep.Replace != nil {
					fmt.Printf("IG replace package: %+v\n", dep.Replace)
				}

			}
		}
	}
}
