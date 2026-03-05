/*
Copyright 2025 The Kubernetes Authors.

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

package options

import (
	"k8s.io/client-go/informers"
	kube_client "k8s.io/client-go/kubernetes"
	"k8s.io/cluster-autoscaler/pkg/cloudprovider"
	"k8s.io/cluster-autoscaler/pkg/config"
	"k8s.io/cluster-autoscaler/pkg/context"
	"k8s.io/cluster-autoscaler/pkg/core/scaledown/pdb"
	"k8s.io/cluster-autoscaler/pkg/core/scaleup"
	"k8s.io/cluster-autoscaler/pkg/debuggingsnapshot"
	"k8s.io/cluster-autoscaler/pkg/estimator"
	"k8s.io/cluster-autoscaler/pkg/expander"
	"k8s.io/cluster-autoscaler/pkg/observers/loopstart"
	ca_processors "k8s.io/cluster-autoscaler/pkg/processors"
	"k8s.io/cluster-autoscaler/pkg/resourcequotas"
	"k8s.io/cluster-autoscaler/pkg/simulator/clustersnapshot"
	csinodeprovider "k8s.io/cluster-autoscaler/pkg/simulator/csi/provider"
	"k8s.io/cluster-autoscaler/pkg/simulator/drainability/rules"
	draprovider "k8s.io/cluster-autoscaler/pkg/simulator/dynamicresources/provider"
	"k8s.io/cluster-autoscaler/pkg/simulator/framework"
	"k8s.io/cluster-autoscaler/pkg/simulator/options"
	"k8s.io/cluster-autoscaler/pkg/utils/backoff"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AutoscalerOptions is the whole set of options for configuring an autoscaler
type AutoscalerOptions struct {
	config.AutoscalingOptions
	KubeClient             kube_client.Interface
	InformerFactory        informers.SharedInformerFactory
	AutoscalingKubeClients *context.AutoscalingKubeClients
	CloudProvider          cloudprovider.CloudProvider
	FrameworkHandle        *framework.Handle
	ClusterSnapshot        clustersnapshot.ClusterSnapshot
	ExpanderStrategy       expander.Strategy
	EstimatorBuilder       estimator.EstimatorBuilder
	Processors             *ca_processors.AutoscalingProcessors
	LoopStartNotifier      *loopstart.ObserversList
	Backoff                backoff.Backoff
	DebuggingSnapshotter   debuggingsnapshot.DebuggingSnapshotter
	RemainingPdbTracker    pdb.RemainingPdbTracker
	ScaleUpOrchestrator    scaleup.Orchestrator
	DeleteOptions          options.NodeDeleteOptions
	DrainabilityRules      rules.Rules
	DraProvider            *draprovider.Provider
	QuotasTrackerOptions   resourcequotas.TrackerOptions
	CSIProvider            *csinodeprovider.Provider
	KubeClientNew          client.Client
	KubeCache              cache.Cache
}
