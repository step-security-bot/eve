// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.2
// source: config/devconfig.proto

package config

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is the response to a GET /api/v1/edgeDevice/config
// The EdgeDevConfig message carries all of the device's configuration from
// the controller to the device.
// The device will request these messages either periodically or as a result
// of some TBD notification.
// The message is assumed to be protected by a TLS session bound to the
// device certificate.
type EdgeDevConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *UUIDandVersion      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Apps       []*AppInstanceConfig `protobuf:"bytes,4,rep,name=apps,proto3" json:"apps,omitempty"`
	Networks   []*NetworkConfig     `protobuf:"bytes,5,rep,name=networks,proto3" json:"networks,omitempty"`
	Datastores []*DatastoreConfig   `protobuf:"bytes,6,rep,name=datastores,proto3" json:"datastores,omitempty"`
	// OBSOLETE - base. Use baseos instead. Controller should fill this for
	// backward compatibility till all the Older Eve images are no longer
	// supported.
	Base        []*BaseOSConfig `protobuf:"bytes,8,rep,name=base,proto3" json:"base,omitempty"` // BaseOSImage config block
	Reboot      *DeviceOpsCmd   `protobuf:"bytes,9,opt,name=reboot,proto3" json:"reboot,omitempty"`
	Backup      *DeviceOpsCmd   `protobuf:"bytes,10,opt,name=backup,proto3" json:"backup,omitempty"`
	ConfigItems []*ConfigItem   `protobuf:"bytes,11,rep,name=configItems,proto3" json:"configItems,omitempty"`
	// systemAdapterList - List of DeviceNetworkAdapters. Only Network
	//  adapters ( Ex: eth0, wlan1 etc ) have a corresponding SystemAdapter.
	// non-Network adapters do not have systemadapters.
	SystemAdapterList []*SystemAdapter `protobuf:"bytes,12,rep,name=systemAdapterList,proto3" json:"systemAdapterList,omitempty"`
	// deviceIoList - List of Physical Adapters. Includes both Network
	//  Adapters and Non-Network Adapters ( USB / Com etc )
	DeviceIoList []*PhysicalIO `protobuf:"bytes,13,rep,name=deviceIoList,proto3" json:"deviceIoList,omitempty"`
	// Override dmidecode info if set
	Manufacturer     string                   `protobuf:"bytes,14,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductName      string                   `protobuf:"bytes,15,opt,name=productName,proto3" json:"productName,omitempty"`
	NetworkInstances []*NetworkInstanceConfig `protobuf:"bytes,16,rep,name=networkInstances,proto3" json:"networkInstances,omitempty"`
	// controller supplies a list of cipher contexts,
	// containing certificate and other details, to be
	// used for sensitive data decryption
	CipherContexts []*CipherContext `protobuf:"bytes,19,rep,name=cipherContexts,proto3" json:"cipherContexts,omitempty"`
	// These images aka ContentTrees and Volumes should be created by EVE
	// independently of any application usage.
	// Application instances will refer to the volumes.
	ContentInfo []*ContentTree `protobuf:"bytes,20,rep,name=contentInfo,proto3" json:"contentInfo,omitempty"`
	Volumes     []*Volume      `protobuf:"bytes,21,rep,name=volumes,proto3" json:"volumes,omitempty"`
	// This field is used by the device to detect when it needs to re-download
	// the controller certs using the /certs API endpoint.
	// The controller just needs to ensure this value changes when it wants the
	// device to re-fetch the controller certs, for instance by having it
	// be a hash of all of the controller certificates.
	ControllercertConfighash string `protobuf:"bytes,22,opt,name=controllercert_confighash,json=controllercertConfighash,proto3" json:"controllercert_confighash,omitempty"`
	// deprecated 23;
	// If maintence_mode is set the device will operate in a limited mode e.g.,
	// not start applications etc as to enable inspection of its state and
	// recover from bad state.
	MaintenanceMode bool `protobuf:"varint,24,opt,name=maintenance_mode,json=maintenanceMode,proto3" json:"maintenance_mode,omitempty"`
	// controller_epoch indicates current epoch of config
	// if we set new epoch, EVE sends all info messages to controller
	// it captures when a new controller takes over and needs all the info be resent
	ControllerEpoch int64 `protobuf:"varint,25,opt,name=controller_epoch,json=controllerEpoch,proto3" json:"controller_epoch,omitempty"`
	// Baseos Config Block
	Baseos *BaseOS `protobuf:"bytes,26,opt,name=baseos,proto3" json:"baseos,omitempty"`
	// global_profile, if set, controls set of applications which will run.
	// The Activate=true app instances which have this profile in their profile_list
	// will run. If the global_profile is not set, then the profile_list is not
	// used to gate the application instances.
	GlobalProfile string `protobuf:"bytes,27,opt,name=global_profile,json=globalProfile,proto3" json:"global_profile,omitempty"`
	// local_profile_server, if set, indicates a hostname/IPv4/IPv6 address and
	// optional port number at which EVE will request for a local profile.
	// If such a local profile is retrieved, it will override the global_profile.
	// The syntax follows the usual URL server name syntax thus the following
	// are example valid strings:
	//    [fe80::1]:1234
	//    10.1.1.1:1234
	//    hostname:1234
	//    [fe80::1]
	//    10.1.1.1
	//    hostname
	// If the port number is not specified, it will default to 8888
	LocalProfileServer string `protobuf:"bytes,28,opt,name=local_profile_server,json=localProfileServer,proto3" json:"local_profile_server,omitempty"`
	// Together with a local_profile_server one can specify a
	// profile_server_token. EVE must verify that the response from the
	// local_profile_server contains this token.
	ProfileServerToken string `protobuf:"bytes,29,opt,name=profile_server_token,json=profileServerToken,proto3" json:"profile_server_token,omitempty"`
}

func (x *EdgeDevConfig) Reset() {
	*x = EdgeDevConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_devconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeDevConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeDevConfig) ProtoMessage() {}

func (x *EdgeDevConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_devconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeDevConfig.ProtoReflect.Descriptor instead.
func (*EdgeDevConfig) Descriptor() ([]byte, []int) {
	return file_config_devconfig_proto_rawDescGZIP(), []int{0}
}

func (x *EdgeDevConfig) GetId() *UUIDandVersion {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EdgeDevConfig) GetApps() []*AppInstanceConfig {
	if x != nil {
		return x.Apps
	}
	return nil
}

func (x *EdgeDevConfig) GetNetworks() []*NetworkConfig {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *EdgeDevConfig) GetDatastores() []*DatastoreConfig {
	if x != nil {
		return x.Datastores
	}
	return nil
}

func (x *EdgeDevConfig) GetBase() []*BaseOSConfig {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *EdgeDevConfig) GetReboot() *DeviceOpsCmd {
	if x != nil {
		return x.Reboot
	}
	return nil
}

func (x *EdgeDevConfig) GetBackup() *DeviceOpsCmd {
	if x != nil {
		return x.Backup
	}
	return nil
}

func (x *EdgeDevConfig) GetConfigItems() []*ConfigItem {
	if x != nil {
		return x.ConfigItems
	}
	return nil
}

func (x *EdgeDevConfig) GetSystemAdapterList() []*SystemAdapter {
	if x != nil {
		return x.SystemAdapterList
	}
	return nil
}

func (x *EdgeDevConfig) GetDeviceIoList() []*PhysicalIO {
	if x != nil {
		return x.DeviceIoList
	}
	return nil
}

func (x *EdgeDevConfig) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *EdgeDevConfig) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *EdgeDevConfig) GetNetworkInstances() []*NetworkInstanceConfig {
	if x != nil {
		return x.NetworkInstances
	}
	return nil
}

func (x *EdgeDevConfig) GetCipherContexts() []*CipherContext {
	if x != nil {
		return x.CipherContexts
	}
	return nil
}

func (x *EdgeDevConfig) GetContentInfo() []*ContentTree {
	if x != nil {
		return x.ContentInfo
	}
	return nil
}

func (x *EdgeDevConfig) GetVolumes() []*Volume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

func (x *EdgeDevConfig) GetControllercertConfighash() string {
	if x != nil {
		return x.ControllercertConfighash
	}
	return ""
}

func (x *EdgeDevConfig) GetMaintenanceMode() bool {
	if x != nil {
		return x.MaintenanceMode
	}
	return false
}

func (x *EdgeDevConfig) GetControllerEpoch() int64 {
	if x != nil {
		return x.ControllerEpoch
	}
	return 0
}

func (x *EdgeDevConfig) GetBaseos() *BaseOS {
	if x != nil {
		return x.Baseos
	}
	return nil
}

func (x *EdgeDevConfig) GetGlobalProfile() string {
	if x != nil {
		return x.GlobalProfile
	}
	return ""
}

func (x *EdgeDevConfig) GetLocalProfileServer() string {
	if x != nil {
		return x.LocalProfileServer
	}
	return ""
}

func (x *EdgeDevConfig) GetProfileServerToken() string {
	if x != nil {
		return x.ProfileServerToken
	}
	return ""
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigHash     string `protobuf:"bytes,1,opt,name=configHash,proto3" json:"configHash,omitempty"`
	IntegrityToken []byte `protobuf:"bytes,2,opt,name=integrity_token,json=integrityToken,proto3" json:"integrity_token,omitempty"` // value provided by controller during remote attestation
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_devconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_devconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_devconfig_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigRequest) GetConfigHash() string {
	if x != nil {
		return x.ConfigHash
	}
	return ""
}

func (x *ConfigRequest) GetIntegrityToken() []byte {
	if x != nil {
		return x.IntegrityToken
	}
	return nil
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config     *EdgeDevConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	ConfigHash string         `protobuf:"bytes,2,opt,name=configHash,proto3" json:"configHash,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_devconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_devconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_config_devconfig_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigResponse) GetConfig() *EdgeDevConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ConfigResponse) GetConfigHash() string {
	if x != nil {
		return x.ConfigHash
	}
	return ""
}

var File_config_devconfig_proto protoreflect.FileDescriptor

var file_config_devconfig_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x69,
	0x6e, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe3, 0x0a, 0x0a, 0x0d, 0x45, 0x64, 0x67, 0x65, 0x44, 0x65, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x61, 0x6e, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41,
	0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x37, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4f, 0x53, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x72, 0x65, 0x62,
	0x6f, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x73, 0x43, 0x6d, 0x64, 0x52, 0x06,
	0x72, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x73, 0x43, 0x6d, 0x64, 0x52, 0x06, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x52, 0x0a, 0x11, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0c,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x49, 0x4f, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x58, 0x0a, 0x10, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x10, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0e, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x0e, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x12, 0x44, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x12, 0x3b, 0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x63, 0x65,
	0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x68, 0x61, 0x73, 0x68, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x63,
	0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x68, 0x61, 0x73, 0x68, 0x12, 0x29, 0x0a,
	0x10, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x65, 0x6f, 0x73, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x4f, 0x53, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x1d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x58, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x6e, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x44, 0x65,
	0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x42,
	0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_devconfig_proto_rawDescOnce sync.Once
	file_config_devconfig_proto_rawDescData = file_config_devconfig_proto_rawDesc
)

func file_config_devconfig_proto_rawDescGZIP() []byte {
	file_config_devconfig_proto_rawDescOnce.Do(func() {
		file_config_devconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_devconfig_proto_rawDescData)
	})
	return file_config_devconfig_proto_rawDescData
}

var file_config_devconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_devconfig_proto_goTypes = []interface{}{
	(*EdgeDevConfig)(nil),         // 0: org.lfedge.eve.config.EdgeDevConfig
	(*ConfigRequest)(nil),         // 1: org.lfedge.eve.config.ConfigRequest
	(*ConfigResponse)(nil),        // 2: org.lfedge.eve.config.ConfigResponse
	(*UUIDandVersion)(nil),        // 3: org.lfedge.eve.config.UUIDandVersion
	(*AppInstanceConfig)(nil),     // 4: org.lfedge.eve.config.AppInstanceConfig
	(*NetworkConfig)(nil),         // 5: org.lfedge.eve.config.NetworkConfig
	(*DatastoreConfig)(nil),       // 6: org.lfedge.eve.config.DatastoreConfig
	(*BaseOSConfig)(nil),          // 7: org.lfedge.eve.config.BaseOSConfig
	(*DeviceOpsCmd)(nil),          // 8: org.lfedge.eve.config.DeviceOpsCmd
	(*ConfigItem)(nil),            // 9: org.lfedge.eve.config.ConfigItem
	(*SystemAdapter)(nil),         // 10: org.lfedge.eve.config.SystemAdapter
	(*PhysicalIO)(nil),            // 11: org.lfedge.eve.config.PhysicalIO
	(*NetworkInstanceConfig)(nil), // 12: org.lfedge.eve.config.NetworkInstanceConfig
	(*CipherContext)(nil),         // 13: org.lfedge.eve.config.CipherContext
	(*ContentTree)(nil),           // 14: org.lfedge.eve.config.ContentTree
	(*Volume)(nil),                // 15: org.lfedge.eve.config.Volume
	(*BaseOS)(nil),                // 16: org.lfedge.eve.config.BaseOS
}
var file_config_devconfig_proto_depIdxs = []int32{
	3,  // 0: org.lfedge.eve.config.EdgeDevConfig.id:type_name -> org.lfedge.eve.config.UUIDandVersion
	4,  // 1: org.lfedge.eve.config.EdgeDevConfig.apps:type_name -> org.lfedge.eve.config.AppInstanceConfig
	5,  // 2: org.lfedge.eve.config.EdgeDevConfig.networks:type_name -> org.lfedge.eve.config.NetworkConfig
	6,  // 3: org.lfedge.eve.config.EdgeDevConfig.datastores:type_name -> org.lfedge.eve.config.DatastoreConfig
	7,  // 4: org.lfedge.eve.config.EdgeDevConfig.base:type_name -> org.lfedge.eve.config.BaseOSConfig
	8,  // 5: org.lfedge.eve.config.EdgeDevConfig.reboot:type_name -> org.lfedge.eve.config.DeviceOpsCmd
	8,  // 6: org.lfedge.eve.config.EdgeDevConfig.backup:type_name -> org.lfedge.eve.config.DeviceOpsCmd
	9,  // 7: org.lfedge.eve.config.EdgeDevConfig.configItems:type_name -> org.lfedge.eve.config.ConfigItem
	10, // 8: org.lfedge.eve.config.EdgeDevConfig.systemAdapterList:type_name -> org.lfedge.eve.config.SystemAdapter
	11, // 9: org.lfedge.eve.config.EdgeDevConfig.deviceIoList:type_name -> org.lfedge.eve.config.PhysicalIO
	12, // 10: org.lfedge.eve.config.EdgeDevConfig.networkInstances:type_name -> org.lfedge.eve.config.NetworkInstanceConfig
	13, // 11: org.lfedge.eve.config.EdgeDevConfig.cipherContexts:type_name -> org.lfedge.eve.config.CipherContext
	14, // 12: org.lfedge.eve.config.EdgeDevConfig.contentInfo:type_name -> org.lfedge.eve.config.ContentTree
	15, // 13: org.lfedge.eve.config.EdgeDevConfig.volumes:type_name -> org.lfedge.eve.config.Volume
	16, // 14: org.lfedge.eve.config.EdgeDevConfig.baseos:type_name -> org.lfedge.eve.config.BaseOS
	0,  // 15: org.lfedge.eve.config.ConfigResponse.config:type_name -> org.lfedge.eve.config.EdgeDevConfig
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_config_devconfig_proto_init() }
func file_config_devconfig_proto_init() {
	if File_config_devconfig_proto != nil {
		return
	}
	file_config_acipherinfo_proto_init()
	file_config_appconfig_proto_init()
	file_config_baseosconfig_proto_init()
	file_config_devcommon_proto_init()
	file_config_devmodel_proto_init()
	file_config_netconfig_proto_init()
	file_config_netinst_proto_init()
	file_config_storage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_devconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeDevConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_devconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_devconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_devconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_devconfig_proto_goTypes,
		DependencyIndexes: file_config_devconfig_proto_depIdxs,
		MessageInfos:      file_config_devconfig_proto_msgTypes,
	}.Build()
	File_config_devconfig_proto = out.File
	file_config_devconfig_proto_rawDesc = nil
	file_config_devconfig_proto_goTypes = nil
	file_config_devconfig_proto_depIdxs = nil
}
