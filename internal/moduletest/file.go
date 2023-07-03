// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package moduletest

import "github.com/hashicorp/terraform/internal/configs"

type File struct {
	Config *configs.TestFile

	Name   string
	Status Status

	Runs []*Run
}
