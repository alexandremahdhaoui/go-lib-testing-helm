package tHelm

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	tUtils "gitlab.com/alexandre.mahdhaoui/go-lib-testing-utils"
	"testing"
)

//----------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------- Functions --------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

func NewHelmConfig(t *testing.T) HelmConfig {
	h := HelmConfig{}

	h.SetChartPath("")
	h.SetHelmOpt(&helm.Options{})
	h.SetId(tUtils.Uuid())
	h.SetT(t)

	return h
}

//----------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------ Struct --------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type HelmConfig struct {

	// Getter
	HelmOptGetter
	tUtils.Tester
	tUtils.Identifier

	// Setter
	HelmOptSetter
	ChartPathSetter

	// Fields
	chartPath string
	id        string
	helmOpt   *helm.Options
	t         *testing.T
}

//----------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------- Interfaces ------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type HelmTester interface {
	ChartPathGetter
	HelmOptGetter
	tUtils.Identifier
	tUtils.Tester
}

type HelmOptGetter interface{ HelmOpt() *helm.Options }
type ChartPathGetter interface{ ChartPath() string }

type HelmOptSetter interface{ SetHelmOpt(*helm.Options) }
type ChartPathSetter interface{ SetChartPath(string) }

//----------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------ Getters -------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

func (h *HelmConfig) ChartPath() string      { return h.chartPath }
func (h *HelmConfig) HelmOpt() *helm.Options { return h.helmOpt }
func (h *HelmConfig) Id() string             { return h.id }
func (h *HelmConfig) T() *testing.T          { return h.t }

//----------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------ Setters -------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

func (h *HelmConfig) SetChartPath(s string)       { h.chartPath = s }
func (h *HelmConfig) SetHelmOpt(ho *helm.Options) { h.helmOpt = ho }
func (h *HelmConfig) SetId(s string)              { h.id = s }
func (h *HelmConfig) SetT(t *testing.T)           { h.t = t }
