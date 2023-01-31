// Copyright 2022 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package provenance_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/api/provenance"
)

const expectedBuildDateFromLdFlag = "2023-01-31T23:38:41Z"
const expectedVersionFromLdFlag = "(test)"

func TestGetProvenance(t *testing.T) {
	p := provenance.GetProvenance()
	// These are set by ldflags in our Makefile
	assert.Equal(t, "(test)", p.Version)
	assert.Equal(t, expectedBuildDateFromLdFlag, p.BuildDate)
	// This comes from BuildInfo, which is not set during go test: https://github.com/golang/go/issues/33976
	assert.Equal(t, "unknown", p.GitCommit)

	// These are set properly during go test
	assert.NotEmpty(t, p.GoArch)
	assert.NotEmpty(t, p.GoOs)
	assert.Contains(t, p.GoVersion, "go1.")
}

func TestProvenance_Short(t *testing.T) {
	p := provenance.GetProvenance()
	// The version not set during go test, so this comes from an ldflag: https://github.com/golang/go/issues/33976
	assert.Equal(t, "(test)", p.Short())

	p.Version = "kustomize/v4.11.12"
	assert.Equal(t, "v4.11.12", p.Short())
}

func TestProvenance_Full(t *testing.T) {
	p := provenance.GetProvenance()
	// Most values are not set during go test: https://github.com/golang/go/issues/33976
	assert.Contains(t, p.Full(), fmt.Sprintf("{Version:%s GitCommit:unknown BuildDate:%s", expectedVersionFromLdFlag, expectedBuildDateFromLdFlag))
	assert.Regexp(t, "GoOs:\\w+ GoArch:\\w+ GoVersion:go1", p.Full())
}

func TestProvenance_Semver(t *testing.T) {
	p := provenance.GetProvenance()
	// The version not set during go test
	assert.Equal(t, "(test)", p.Semver())

	p.Version = "kustomize/v4.11.12"
	assert.Equal(t, "v4.11.12", p.Semver())
}
