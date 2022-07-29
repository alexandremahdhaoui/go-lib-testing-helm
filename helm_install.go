package tHelm

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/logger"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
	tKube "gitlab.com/alexandre.mahdhaoui/go-lib-testing-kube"
)

type KubeHelmTester interface {
	HelmTester
	tKube.KubeOptGetter
}

// Install installs a Helm Chart
func Install(h KubeHelmTester) func() {
	_, teardownNs := Init(h)
	teardownChart := Upgrade(h)

	return func() {
		teardownNs()
		teardownChart()
	}
}

// Upgrade install/upgrade a chart
//	Please ensure `chart dependency update` is done and Namespace exist before running this function
// 	If you're calling this function directly please make sure to Init() beforehand
func Upgrade(h KubeHelmTester) func() {
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
func AddRepository(h HelmTester, uri string, up UserPassGetter) string {
	discardHelmLogger(h.HelmOpt())

	u := up.User()
	p := up.Pass()

	s, err := helm.RunHelmCommandAndGetOutputE(
		h.T(), h.HelmOpt(),
		"repo", "add",
		h.Id(), uri,
		"--username", u, "--password", p,
	)
	require.NoError(h.T(), err)

	return s
}

func discardHelmLogger(o *helm.Options) {
	o.Logger = logger.Discard
}
