// This file was generated by counterfeiter
package awsfakes

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/c0-ops/cliaas/iaas/aws"
)

type FakeClient struct {
	CreateVMStub        func(ami, instanceType, name, keyName, subnetID, securityGroupID string) (*ec2.Instance, error)
	createVMMutex       sync.RWMutex
	createVMArgsForCall []struct {
		ami             string
		instanceType    string
		name            string
		keyName         string
		subnetID        string
		securityGroupID string
	}
	createVMReturns struct {
		result1 *ec2.Instance
		result2 error
	}
	DeleteVMStub        func(instanceID string) error
	deleteVMMutex       sync.RWMutex
	deleteVMArgsForCall []struct {
		instanceID string
	}
	deleteVMReturns struct {
		result1 error
	}
	GetVMInfoStub        func(name string) (*ec2.Instance, error)
	getVMInfoMutex       sync.RWMutex
	getVMInfoArgsForCall []struct {
		name string
	}
	getVMInfoReturns struct {
		result1 *ec2.Instance
		result2 error
	}
	StopVMStub        func(instance ec2.Instance) error
	stopVMMutex       sync.RWMutex
	stopVMArgsForCall []struct {
		instance ec2.Instance
	}
	stopVMReturns struct {
		result1 error
	}
	AssignPublicIPStub        func(instance ec2.Instance, ip string) error
	assignPublicIPMutex       sync.RWMutex
	assignPublicIPArgsForCall []struct {
		instance ec2.Instance
		ip       string
	}
	assignPublicIPReturns struct {
		result1 error
	}
	WaitForStartedVMStub        func(instanceName string) error
	waitForStartedVMMutex       sync.RWMutex
	waitForStartedVMArgsForCall []struct {
		instanceName string
	}
	waitForStartedVMReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CreateVM(ami string, instanceType string, name string, keyName string, subnetID string, securityGroupID string) (*ec2.Instance, error) {
	fake.createVMMutex.Lock()
	fake.createVMArgsForCall = append(fake.createVMArgsForCall, struct {
		ami             string
		instanceType    string
		name            string
		keyName         string
		subnetID        string
		securityGroupID string
	}{ami, instanceType, name, keyName, subnetID, securityGroupID})
	fake.recordInvocation("CreateVM", []interface{}{ami, instanceType, name, keyName, subnetID, securityGroupID})
	fake.createVMMutex.Unlock()
	if fake.CreateVMStub != nil {
		return fake.CreateVMStub(ami, instanceType, name, keyName, subnetID, securityGroupID)
	} else {
		return fake.createVMReturns.result1, fake.createVMReturns.result2
	}
}

func (fake *FakeClient) CreateVMCallCount() int {
	fake.createVMMutex.RLock()
	defer fake.createVMMutex.RUnlock()
	return len(fake.createVMArgsForCall)
}

func (fake *FakeClient) CreateVMArgsForCall(i int) (string, string, string, string, string, string) {
	fake.createVMMutex.RLock()
	defer fake.createVMMutex.RUnlock()
	return fake.createVMArgsForCall[i].ami, fake.createVMArgsForCall[i].instanceType, fake.createVMArgsForCall[i].name, fake.createVMArgsForCall[i].keyName, fake.createVMArgsForCall[i].subnetID, fake.createVMArgsForCall[i].securityGroupID
}

func (fake *FakeClient) CreateVMReturns(result1 *ec2.Instance, result2 error) {
	fake.CreateVMStub = nil
	fake.createVMReturns = struct {
		result1 *ec2.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeleteVM(instanceID string) error {
	fake.deleteVMMutex.Lock()
	fake.deleteVMArgsForCall = append(fake.deleteVMArgsForCall, struct {
		instanceID string
	}{instanceID})
	fake.recordInvocation("DeleteVM", []interface{}{instanceID})
	fake.deleteVMMutex.Unlock()
	if fake.DeleteVMStub != nil {
		return fake.DeleteVMStub(instanceID)
	} else {
		return fake.deleteVMReturns.result1
	}
}

func (fake *FakeClient) DeleteVMCallCount() int {
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	return len(fake.deleteVMArgsForCall)
}

func (fake *FakeClient) DeleteVMArgsForCall(i int) string {
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	return fake.deleteVMArgsForCall[i].instanceID
}

func (fake *FakeClient) DeleteVMReturns(result1 error) {
	fake.DeleteVMStub = nil
	fake.deleteVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetVMInfo(name string) (*ec2.Instance, error) {
	fake.getVMInfoMutex.Lock()
	fake.getVMInfoArgsForCall = append(fake.getVMInfoArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetVMInfo", []interface{}{name})
	fake.getVMInfoMutex.Unlock()
	if fake.GetVMInfoStub != nil {
		return fake.GetVMInfoStub(name)
	} else {
		return fake.getVMInfoReturns.result1, fake.getVMInfoReturns.result2
	}
}

func (fake *FakeClient) GetVMInfoCallCount() int {
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	return len(fake.getVMInfoArgsForCall)
}

func (fake *FakeClient) GetVMInfoArgsForCall(i int) string {
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	return fake.getVMInfoArgsForCall[i].name
}

func (fake *FakeClient) GetVMInfoReturns(result1 *ec2.Instance, result2 error) {
	fake.GetVMInfoStub = nil
	fake.getVMInfoReturns = struct {
		result1 *ec2.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) StopVM(instance ec2.Instance) error {
	fake.stopVMMutex.Lock()
	fake.stopVMArgsForCall = append(fake.stopVMArgsForCall, struct {
		instance ec2.Instance
	}{instance})
	fake.recordInvocation("StopVM", []interface{}{instance})
	fake.stopVMMutex.Unlock()
	if fake.StopVMStub != nil {
		return fake.StopVMStub(instance)
	} else {
		return fake.stopVMReturns.result1
	}
}

func (fake *FakeClient) StopVMCallCount() int {
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	return len(fake.stopVMArgsForCall)
}

func (fake *FakeClient) StopVMArgsForCall(i int) ec2.Instance {
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	return fake.stopVMArgsForCall[i].instance
}

func (fake *FakeClient) StopVMReturns(result1 error) {
	fake.StopVMStub = nil
	fake.stopVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AssignPublicIP(instance ec2.Instance, ip string) error {
	fake.assignPublicIPMutex.Lock()
	fake.assignPublicIPArgsForCall = append(fake.assignPublicIPArgsForCall, struct {
		instance ec2.Instance
		ip       string
	}{instance, ip})
	fake.recordInvocation("AssignPublicIP", []interface{}{instance, ip})
	fake.assignPublicIPMutex.Unlock()
	if fake.AssignPublicIPStub != nil {
		return fake.AssignPublicIPStub(instance, ip)
	} else {
		return fake.assignPublicIPReturns.result1
	}
}

func (fake *FakeClient) AssignPublicIPCallCount() int {
	fake.assignPublicIPMutex.RLock()
	defer fake.assignPublicIPMutex.RUnlock()
	return len(fake.assignPublicIPArgsForCall)
}

func (fake *FakeClient) AssignPublicIPArgsForCall(i int) (ec2.Instance, string) {
	fake.assignPublicIPMutex.RLock()
	defer fake.assignPublicIPMutex.RUnlock()
	return fake.assignPublicIPArgsForCall[i].instance, fake.assignPublicIPArgsForCall[i].ip
}

func (fake *FakeClient) AssignPublicIPReturns(result1 error) {
	fake.AssignPublicIPStub = nil
	fake.assignPublicIPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) WaitForStartedVM(instanceName string) error {
	fake.waitForStartedVMMutex.Lock()
	fake.waitForStartedVMArgsForCall = append(fake.waitForStartedVMArgsForCall, struct {
		instanceName string
	}{instanceName})
	fake.recordInvocation("WaitForStartedVM", []interface{}{instanceName})
	fake.waitForStartedVMMutex.Unlock()
	if fake.WaitForStartedVMStub != nil {
		return fake.WaitForStartedVMStub(instanceName)
	} else {
		return fake.waitForStartedVMReturns.result1
	}
}

func (fake *FakeClient) WaitForStartedVMCallCount() int {
	fake.waitForStartedVMMutex.RLock()
	defer fake.waitForStartedVMMutex.RUnlock()
	return len(fake.waitForStartedVMArgsForCall)
}

func (fake *FakeClient) WaitForStartedVMArgsForCall(i int) string {
	fake.waitForStartedVMMutex.RLock()
	defer fake.waitForStartedVMMutex.RUnlock()
	return fake.waitForStartedVMArgsForCall[i].instanceName
}

func (fake *FakeClient) WaitForStartedVMReturns(result1 error) {
	fake.WaitForStartedVMStub = nil
	fake.waitForStartedVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVMMutex.RLock()
	defer fake.createVMMutex.RUnlock()
	fake.deleteVMMutex.RLock()
	defer fake.deleteVMMutex.RUnlock()
	fake.getVMInfoMutex.RLock()
	defer fake.getVMInfoMutex.RUnlock()
	fake.stopVMMutex.RLock()
	defer fake.stopVMMutex.RUnlock()
	fake.assignPublicIPMutex.RLock()
	defer fake.assignPublicIPMutex.RUnlock()
	fake.waitForStartedVMMutex.RLock()
	defer fake.waitForStartedVMMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ aws.Client = new(FakeClient)