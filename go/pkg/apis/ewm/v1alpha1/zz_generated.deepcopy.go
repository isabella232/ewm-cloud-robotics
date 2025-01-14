//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// This file is part of ewm-cloud-robotics
// (see https://github.com/SAP/ewm-cloud-robotics).
//
// This file is licensed under the Apache Software License, v. 2 except as noted
// otherwise in the LICENSE file (https://github.com/SAP/ewm-cloud-robotics/blob/master/LICENSE)
//

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auctioneer) DeepCopyInto(out *Auctioneer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auctioneer.
func (in *Auctioneer) DeepCopy() *Auctioneer {
	if in == nil {
		return nil
	}
	out := new(Auctioneer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Auctioneer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuctioneerList) DeepCopyInto(out *AuctioneerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Auctioneer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuctioneerList.
func (in *AuctioneerList) DeepCopy() *AuctioneerList {
	if in == nil {
		return nil
	}
	out := new(AuctioneerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuctioneerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuctioneerSpec) DeepCopyInto(out *AuctioneerSpec) {
	*out = *in
	out.Scope = in.Scope
	out.Configuration = in.Configuration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuctioneerSpec.
func (in *AuctioneerSpec) DeepCopy() *AuctioneerSpec {
	if in == nil {
		return nil
	}
	out := new(AuctioneerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuctioneerStatus) DeepCopyInto(out *AuctioneerStatus) {
	*out = *in
	if in.AvailableRobots != nil {
		in, out := &in.AvailableRobots, &out.AvailableRobots
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RobotsInScope != nil {
		in, out := &in.RobotsInScope, &out.RobotsInScope
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.LastStatusChangeTime.DeepCopyInto(&out.LastStatusChangeTime)
	in.UpdateTime.DeepCopyInto(&out.UpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuctioneerStatus.
func (in *AuctioneerStatus) DeepCopy() *AuctioneerStatus {
	if in == nil {
		return nil
	}
	out := new(AuctioneerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EWMWarehouseOrder) DeepCopyInto(out *EWMWarehouseOrder) {
	*out = *in
	in.Lsd.DeepCopyInto(&out.Lsd)
	if in.Warehousetasks != nil {
		in, out := &in.Warehousetasks, &out.Warehousetasks
		*out = make([]EWMWarehouseTask, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EWMWarehouseOrder.
func (in *EWMWarehouseOrder) DeepCopy() *EWMWarehouseOrder {
	if in == nil {
		return nil
	}
	out := new(EWMWarehouseOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EWMWarehouseTask) DeepCopyInto(out *EWMWarehouseTask) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EWMWarehouseTask.
func (in *EWMWarehouseTask) DeepCopy() *EWMWarehouseTask {
	if in == nil {
		return nil
	}
	out := new(EWMWarehouseTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EWMWarehouseTaskConfirmation) DeepCopyInto(out *EWMWarehouseTaskConfirmation) {
	*out = *in
	in.ConfirmationDate.DeepCopyInto(&out.ConfirmationDate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EWMWarehouseTaskConfirmation.
func (in *EWMWarehouseTaskConfirmation) DeepCopy() *EWMWarehouseTaskConfirmation {
	if in == nil {
		return nil
	}
	out := new(EWMWarehouseTaskConfirmation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderAssignment) DeepCopyInto(out *OrderAssignment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderAssignment.
func (in *OrderAssignment) DeepCopy() *OrderAssignment {
	if in == nil {
		return nil
	}
	out := new(OrderAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderAuction) DeepCopyInto(out *OrderAuction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderAuction.
func (in *OrderAuction) DeepCopy() *OrderAuction {
	if in == nil {
		return nil
	}
	out := new(OrderAuction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderAuction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderAuctionList) DeepCopyInto(out *OrderAuctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrderAuction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderAuctionList.
func (in *OrderAuctionList) DeepCopy() *OrderAuctionList {
	if in == nil {
		return nil
	}
	out := new(OrderAuctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderAuctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderAuctionSpec) DeepCopyInto(out *OrderAuctionSpec) {
	*out = *in
	if in.WarehouseOrders != nil {
		in, out := &in.WarehouseOrders, &out.WarehouseOrders
		*out = make([]EWMWarehouseOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ValidUntil.DeepCopyInto(&out.ValidUntil)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderAuctionSpec.
func (in *OrderAuctionSpec) DeepCopy() *OrderAuctionSpec {
	if in == nil {
		return nil
	}
	out := new(OrderAuctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderAuctionStatus) DeepCopyInto(out *OrderAuctionStatus) {
	*out = *in
	if in.Biddings != nil {
		in, out := &in.Biddings, &out.Biddings
		*out = make([]WarehouseOrderBidding, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderAuctionStatus.
func (in *OrderAuctionStatus) DeepCopy() *OrderAuctionStatus {
	if in == nil {
		return nil
	}
	out := new(OrderAuctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderRequest) DeepCopyInto(out *OrderRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderRequest.
func (in *OrderRequest) DeepCopy() *OrderRequest {
	if in == nil {
		return nil
	}
	out := new(OrderRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderReservation) DeepCopyInto(out *OrderReservation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderReservation.
func (in *OrderReservation) DeepCopy() *OrderReservation {
	if in == nil {
		return nil
	}
	out := new(OrderReservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderReservation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderReservationList) DeepCopyInto(out *OrderReservationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrderReservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderReservationList.
func (in *OrderReservationList) DeepCopy() *OrderReservationList {
	if in == nil {
		return nil
	}
	out := new(OrderReservationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderReservationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderReservationSpec) DeepCopyInto(out *OrderReservationSpec) {
	*out = *in
	out.OrderRequest = in.OrderRequest
	if in.OrderAssignments != nil {
		in, out := &in.OrderAssignments, &out.OrderAssignments
		*out = make([]OrderAssignment, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderReservationSpec.
func (in *OrderReservationSpec) DeepCopy() *OrderReservationSpec {
	if in == nil {
		return nil
	}
	out := new(OrderReservationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderReservationStatus) DeepCopyInto(out *OrderReservationStatus) {
	*out = *in
	if in.WarehouseOrders != nil {
		in, out := &in.WarehouseOrders, &out.WarehouseOrders
		*out = make([]EWMWarehouseOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OrderAssignments != nil {
		in, out := &in.OrderAssignments, &out.OrderAssignments
		*out = make([]OrderAssignment, len(*in))
		copy(*out, *in)
	}
	in.ValidUntil.DeepCopyInto(&out.ValidUntil)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderReservationStatus.
func (in *OrderReservationStatus) DeepCopy() *OrderReservationStatus {
	if in == nil {
		return nil
	}
	out := new(OrderReservationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Path) DeepCopyInto(out *Path) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Path.
func (in *Path) DeepCopy() *Path {
	if in == nil {
		return nil
	}
	out := new(Path)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotConfiguration) DeepCopyInto(out *RobotConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotConfiguration.
func (in *RobotConfiguration) DeepCopy() *RobotConfiguration {
	if in == nil {
		return nil
	}
	out := new(RobotConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotConfigurationList) DeepCopyInto(out *RobotConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RobotConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotConfigurationList.
func (in *RobotConfigurationList) DeepCopy() *RobotConfigurationList {
	if in == nil {
		return nil
	}
	out := new(RobotConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotConfigurationSpec) DeepCopyInto(out *RobotConfigurationSpec) {
	*out = *in
	if in.Chargers != nil {
		in, out := &in.Chargers, &out.Chargers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotConfigurationSpec.
func (in *RobotConfigurationSpec) DeepCopy() *RobotConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(RobotConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotConfigurationStatus) DeepCopyInto(out *RobotConfigurationStatus) {
	*out = *in
	in.UpdateTime.DeepCopyInto(&out.UpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotConfigurationStatus.
func (in *RobotConfigurationStatus) DeepCopy() *RobotConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(RobotConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunTime) DeepCopyInto(out *RunTime) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunTime.
func (in *RunTime) DeepCopy() *RunTime {
	if in == nil {
		return nil
	}
	out := new(RunTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scope) DeepCopyInto(out *Scope) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scope.
func (in *Scope) DeepCopy() *Scope {
	if in == nil {
		return nil
	}
	out := new(Scope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Time.
func (in *Time) DeepCopy() *Time {
	if in == nil {
		return nil
	}
	out := new(Time)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TravelTimeCalculation) DeepCopyInto(out *TravelTimeCalculation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TravelTimeCalculation.
func (in *TravelTimeCalculation) DeepCopy() *TravelTimeCalculation {
	if in == nil {
		return nil
	}
	out := new(TravelTimeCalculation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TravelTimeCalculation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TravelTimeCalculationList) DeepCopyInto(out *TravelTimeCalculationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TravelTimeCalculation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TravelTimeCalculationList.
func (in *TravelTimeCalculationList) DeepCopy() *TravelTimeCalculationList {
	if in == nil {
		return nil
	}
	out := new(TravelTimeCalculationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TravelTimeCalculationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TravelTimeCalculationSpec) DeepCopyInto(out *TravelTimeCalculationSpec) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]Path, len(*in))
		copy(*out, *in)
	}
	in.ValidUntil.DeepCopyInto(&out.ValidUntil)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TravelTimeCalculationSpec.
func (in *TravelTimeCalculationSpec) DeepCopy() *TravelTimeCalculationSpec {
	if in == nil {
		return nil
	}
	out := new(TravelTimeCalculationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TravelTimeCalculationStatus) DeepCopyInto(out *TravelTimeCalculationStatus) {
	*out = *in
	if in.RunTimes != nil {
		in, out := &in.RunTimes, &out.RunTimes
		*out = make([]RunTime, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TravelTimeCalculationStatus.
func (in *TravelTimeCalculationStatus) DeepCopy() *TravelTimeCalculationStatus {
	if in == nil {
		return nil
	}
	out := new(TravelTimeCalculationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseOrder) DeepCopyInto(out *WarehouseOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseOrder.
func (in *WarehouseOrder) DeepCopy() *WarehouseOrder {
	if in == nil {
		return nil
	}
	out := new(WarehouseOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WarehouseOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseOrderBidding) DeepCopyInto(out *WarehouseOrderBidding) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseOrderBidding.
func (in *WarehouseOrderBidding) DeepCopy() *WarehouseOrderBidding {
	if in == nil {
		return nil
	}
	out := new(WarehouseOrderBidding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseOrderList) DeepCopyInto(out *WarehouseOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WarehouseOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseOrderList.
func (in *WarehouseOrderList) DeepCopy() *WarehouseOrderList {
	if in == nil {
		return nil
	}
	out := new(WarehouseOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WarehouseOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseOrderSpec) DeepCopyInto(out *WarehouseOrderSpec) {
	*out = *in
	in.Data.DeepCopyInto(&out.Data)
	if in.ProcessStatus != nil {
		in, out := &in.ProcessStatus, &out.ProcessStatus
		*out = make([]EWMWarehouseTaskConfirmation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseOrderSpec.
func (in *WarehouseOrderSpec) DeepCopy() *WarehouseOrderSpec {
	if in == nil {
		return nil
	}
	out := new(WarehouseOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseOrderStatus) DeepCopyInto(out *WarehouseOrderStatus) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]EWMWarehouseTaskConfirmation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseOrderStatus.
func (in *WarehouseOrderStatus) DeepCopy() *WarehouseOrderStatus {
	if in == nil {
		return nil
	}
	out := new(WarehouseOrderStatus)
	in.DeepCopyInto(out)
	return out
}
