package tHelm

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	tUtils "gitlab.com/alexandre.mahdhaoui/go-lib-testing-utils"
	"testing"
)

//----------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------- HelmConfigBuilder ------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type HelmConfigBuilder interface {
	Build() HelmConfig
	SetHelmOpt(*helm.Options) HelmConfigBuilder
	SetChartPath(string) HelmConfigBuilder
	SetId(s string) HelmConfigBuilder
	SetT(t *testing.T) HelmConfigBuilder
}

type helmConfigBuilder struct {
	HelmConfigBuilder
	helmConfig helmConfig
}

func NewHelmConfigBuilder() HelmConfigBuilder { return &helmConfigBuilder{helmConfig: helmConfig{}} }

func (b *helmConfigBuilder) SetChartPath(s string) HelmConfigBuilder {
	b.helmConfig.chartPath = s
	return b
}
func (b *helmConfigBuilder) SetHelmOpt(o *helm.Options) HelmConfigBuilder {
	b.helmConfig.helmOpt = o
	return b
}
func (b *helmConfigBuilder) SetId(s string) HelmConfigBuilder {
	b.helmConfig.id = s
	return b
}
func (b *helmConfigBuilder) SetT(t *testing.T) HelmConfigBuilder {
	b.helmConfig.t = t
	return b
}

//----------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------- HelmConfig -------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type HelmConfig interface {
	ChartPath() string
	HelmOpt() *helm.Options
	tUtils.Tester
	tUtils.Identifier
}

type helmConfig struct {
	HelmConfig
	chartPath string
	id        string
	helmOpt   *helm.Options
	t         *testing.T
}

func NewHelmConfig(t *testing.T) HelmConfig {
	return NewHelmConfigBuilder().
		SetChartPath("").
		SetHelmOpt(&helm.Options{}).
		SetId(tUtils.Uuid()).
		SetT(t).
		Build()
}

func (h *helmConfig) ChartPath() string      { return h.chartPath }
func (h *helmConfig) HelmOpt() *helm.Options { return h.helmOpt }
func (h *helmConfig) Id() string             { return h.id }
func (h *helmConfig) T() *testing.T          { return h.t }
