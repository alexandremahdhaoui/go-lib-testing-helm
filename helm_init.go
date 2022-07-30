package tHelm

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
	tKube "gitlab.com/alexandre.mahdhaoui/go-lib-testing-kube"
)

// Init runs helm dependencies update
func Init(h HelmTester, k tKube.KubeTester) (string, func()) {
	s, err := helm.RunHelmCommandAndGetOutputE(
		h.T(), h.HelmOpt(),
		"dependency", "update", h.ChartPath(),
	)
	require.NoError(h.T(), err)

	teardownNs := tKube.CreateNs(k)
	return s, teardownNs
}
