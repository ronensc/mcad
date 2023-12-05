/*
Copyright 2023 IBM Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	mcadv1beta1 "github.com/project-codeflare/mcad/api/v1beta1"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type phaseStepPriority struct {
	phase    mcadv1beta1.AppWrapperPhase
	step     mcadv1beta1.AppWrapperStep
	priority int
}

var (
	appWrappersCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "mcad",
		Name:      "appwrappers_count",
		Help:      "AppWrappers count per phase, step and priority",
	}, []string{"phase", "step", "priority"})
)

func init() {
	metrics.Registry.MustRegister(
		appWrappersCount,
	)
}