// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// This file is part of ewm-cloud-robotics
// (see https://github.com/SAP/ewm-cloud-robotics).
//
// This file is licensed under the Apache Software License, v. 2 except as noted
// otherwise in the LICENSE file (https://github.com/SAP/ewm-cloud-robotics/blob/master/LICENSE)
//

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/ewm/v1alpha1"
	missionv1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/mission/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=ewm.sap.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("auctioneers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ewm().V1alpha1().Auctioneers().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("orderauctions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ewm().V1alpha1().OrderAuctions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("orderreservations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ewm().V1alpha1().OrderReservations().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("robotconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ewm().V1alpha1().RobotConfigurations().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("traveltimecalculations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ewm().V1alpha1().TravelTimeCalculations().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("warehouseorders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ewm().V1alpha1().WarehouseOrders().Informer()}, nil

		// Group=mission.cloudrobotics.com, Version=v1alpha1
	case missionv1alpha1.SchemeGroupVersion.WithResource("missions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mission().V1alpha1().Missions().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
