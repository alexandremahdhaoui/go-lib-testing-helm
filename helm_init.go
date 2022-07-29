package tHelm

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
	tKube "gitlab.com/alexandre.mahdhaoui/go-lib-testing-kube"
	tUtils "gitlab.com/alexandre.mahdhaoui/go-lib-testing-utils"
)

type HelmTester interface {
	ChartPathGetter
	HelmOptGetter
	tUtils.Identifier
	tUtils.Tester
}

// Init runs helm dependencies update
func Init(h KubeHelmTester) (string, func()) {
	s, err := helm.RunHelmCommandAndGetOutputE(
		h.T(), h.HelmOpt(),
		"dependency", "update", h.ChartPath(),
	)
	require.NoError(h.T(), err)

	teardownNs := tKube.CreateNs(h)
	return s, teardownNs
}
