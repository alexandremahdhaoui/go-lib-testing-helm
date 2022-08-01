package tHelm

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/logger"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
	tKube "gitlab.com/alexandre.mahdhaoui/go-lib-testing-kube"
)

// Install installs a Helm Chart
func Install(h HelmConfig, k tKube.KubeTester) func() {
	_, teardownNs := Init(h, k)
	teardownChart := Upgrade(h)

	return func() {
		teardownNs()
		teardownChart()
	}
}

// Upgrade install/upgrade a chart
//	Please ensure `chart dependency update` is done and Namespace exist before running this function
// 	If you're calling this function directly please make sure to Init() beforehand
func Upgrade(h HelmConfig) func() {
	helm.Upgrade(h.T(), h.HelmOpt(), h.ChartPath(), h.Id())

	// Prepare Teardown
	return func() {
		test_structure.RunTestStage(h.T(), "teardown", func() {
			helm.Delete(h.T(), h.HelmOpt(), h.Id(), true)
		})
	}
}

type UserPassGetter interface {
	User() string
	Pass() string
}

// AddRepository adds a helm repository specified by `uri` with User/Pass authentication
func AddRepository(h HelmConfig, uri string, up UserPassGetter) string {
	// Discards Helm Logger
	o := h.HelmOpt()
	o.Logger = logger.Discard

	u, p := up.User(), up.Pass()
	s, err := helm.RunHelmCommandAndGetOutputE(
		h.T(), h.HelmOpt(),
		"repo", "add",
		h.Id(), uri,
		"--username", u, "--password", p,
	)
	require.NoError(h.T(), err)
	o.Logger = logger.TestingT
	return s
}
