// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/hosxp.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RequestCid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *RequestCid) Reset() {
	*x = RequestCid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCid) ProtoMessage() {}

func (x *RequestCid) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCid.ProtoReflect.Descriptor instead.
func (*RequestCid) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{0}
}

func (x *RequestCid) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type RequestHospcode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospcode string `protobuf:"bytes,1,opt,name=hospcode,proto3" json:"hospcode,omitempty"`
}

func (x *RequestHospcode) Reset() {
	*x = RequestHospcode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHospcode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHospcode) ProtoMessage() {}

func (x *RequestHospcode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHospcode.ProtoReflect.Descriptor instead.
func (*RequestHospcode) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{1}
}

func (x *RequestHospcode) GetHospcode() string {
	if x != nil {
		return x.Hospcode
	}
	return ""
}

type RequestPatient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string hn = 1;
	// string vn = 2;
	Hospcode string `protobuf:"bytes,1,opt,name=hospcode,proto3" json:"hospcode,omitempty"`
}

func (x *RequestPatient) Reset() {
	*x = RequestPatient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPatient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPatient) ProtoMessage() {}

func (x *RequestPatient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPatient.ProtoReflect.Descriptor instead.
func (*RequestPatient) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{2}
}

func (x *RequestPatient) GetHospcode() string {
	if x != nil {
		return x.Hospcode
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*InfoResponse_Info `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{3}
}

func (x *InfoResponse) GetResults() []*InfoResponse_Info {
	if x != nil {
		return x.Results
	}
	return nil
}

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ServiceResponse_Service `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceResponse) GetServices() []*ServiceResponse_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type DoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*DoctorResponse_Doctor `protobuf:"bytes,2,rep,name=Results,proto3" json:"Results,omitempty"`
}

func (x *DoctorResponse) Reset() {
	*x = DoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorResponse) ProtoMessage() {}

func (x *DoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorResponse.ProtoReflect.Descriptor instead.
func (*DoctorResponse) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{5}
}

func (x *DoctorResponse) GetResults() []*DoctorResponse_Doctor {
	if x != nil {
		return x.Results
	}
	return nil
}

type ScreeningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Screenings []*ScreeningResponse_Screening `protobuf:"bytes,2,rep,name=screenings,proto3" json:"screenings,omitempty"`
}

func (x *ScreeningResponse) Reset() {
	*x = ScreeningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreeningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreeningResponse) ProtoMessage() {}

func (x *ScreeningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreeningResponse.ProtoReflect.Descriptor instead.
func (*ScreeningResponse) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{6}
}

func (x *ScreeningResponse) GetScreenings() []*ScreeningResponse_Screening {
	if x != nil {
		return x.Screenings
	}
	return nil
}

type InfoResponse_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GwRecordId string `protobuf:"bytes,1,opt,name=gw_record_id,json=gwRecordId,proto3" json:"gw_record_id,omitempty"`
	PatientHn  string `protobuf:"bytes,2,opt,name=patient_hn,json=patientHn,proto3" json:"patient_hn,omitempty"`
	Cid        string `protobuf:"bytes,3,opt,name=cid,proto3" json:"cid,omitempty"`
	GwHospcode string `protobuf:"bytes,4,opt,name=gw_hospcode,json=gwHospcode,proto3" json:"gw_hospcode,omitempty"`
	Pname      string `protobuf:"bytes,5,opt,name=pname,proto3" json:"pname,omitempty"`
	Fname      string `protobuf:"bytes,6,opt,name=fname,proto3" json:"fname,omitempty"`
	Lname      string `protobuf:"bytes,7,opt,name=lname,proto3" json:"lname,omitempty"`
	Birthdate  string `protobuf:"bytes,8,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
}

func (x *InfoResponse_Info) Reset() {
	*x = InfoResponse_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse_Info) ProtoMessage() {}

func (x *InfoResponse_Info) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse_Info.ProtoReflect.Descriptor instead.
func (*InfoResponse_Info) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{3, 0}
}

func (x *InfoResponse_Info) GetGwRecordId() string {
	if x != nil {
		return x.GwRecordId
	}
	return ""
}

func (x *InfoResponse_Info) GetPatientHn() string {
	if x != nil {
		return x.PatientHn
	}
	return ""
}

func (x *InfoResponse_Info) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *InfoResponse_Info) GetGwHospcode() string {
	if x != nil {
		return x.GwHospcode
	}
	return ""
}

func (x *InfoResponse_Info) GetPname() string {
	if x != nil {
		return x.Pname
	}
	return ""
}

func (x *InfoResponse_Info) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *InfoResponse_Info) GetLname() string {
	if x != nil {
		return x.Lname
	}
	return ""
}

func (x *InfoResponse_Info) GetBirthdate() string {
	if x != nil {
		return x.Birthdate
	}
	return ""
}

type ServiceResponse_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GwRecordId string `protobuf:"bytes,1,opt,name=gw_record_id,json=gwRecordId,proto3" json:"gw_record_id,omitempty"`
	Hn         string `protobuf:"bytes,2,opt,name=hn,proto3" json:"hn,omitempty"`
	Vn         string `protobuf:"bytes,3,opt,name=vn,proto3" json:"vn,omitempty"`
	GwHospcode string `protobuf:"bytes,4,opt,name=gw_hospcode,json=gwHospcode,proto3" json:"gw_hospcode,omitempty"`
	Vstdate    string `protobuf:"bytes,5,opt,name=vstdate,proto3" json:"vstdate,omitempty"`
	Vsttime    string `protobuf:"bytes,6,opt,name=vsttime,proto3" json:"vsttime,omitempty"`
	Pttype     string `protobuf:"bytes,7,opt,name=pttype,proto3" json:"pttype,omitempty"`
	Pttypeno   string `protobuf:"bytes,8,opt,name=pttypeno,proto3" json:"pttypeno,omitempty"`
	Spclty     string `protobuf:"bytes,9,opt,name=spclty,proto3" json:"spclty,omitempty"`
	Hospname   string `protobuf:"bytes,10,opt,name=hospname,proto3" json:"hospname,omitempty"`
}

func (x *ServiceResponse_Service) Reset() {
	*x = ServiceResponse_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse_Service) ProtoMessage() {}

func (x *ServiceResponse_Service) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse_Service.ProtoReflect.Descriptor instead.
func (*ServiceResponse_Service) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ServiceResponse_Service) GetGwRecordId() string {
	if x != nil {
		return x.GwRecordId
	}
	return ""
}

func (x *ServiceResponse_Service) GetHn() string {
	if x != nil {
		return x.Hn
	}
	return ""
}

func (x *ServiceResponse_Service) GetVn() string {
	if x != nil {
		return x.Vn
	}
	return ""
}

func (x *ServiceResponse_Service) GetGwHospcode() string {
	if x != nil {
		return x.GwHospcode
	}
	return ""
}

func (x *ServiceResponse_Service) GetVstdate() string {
	if x != nil {
		return x.Vstdate
	}
	return ""
}

func (x *ServiceResponse_Service) GetVsttime() string {
	if x != nil {
		return x.Vsttime
	}
	return ""
}

func (x *ServiceResponse_Service) GetPttype() string {
	if x != nil {
		return x.Pttype
	}
	return ""
}

func (x *ServiceResponse_Service) GetPttypeno() string {
	if x != nil {
		return x.Pttypeno
	}
	return ""
}

func (x *ServiceResponse_Service) GetSpclty() string {
	if x != nil {
		return x.Spclty
	}
	return ""
}

func (x *ServiceResponse_Service) GetHospname() string {
	if x != nil {
		return x.Hospname
	}
	return ""
}

type DoctorResponse_Doctor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GwRecordId string `protobuf:"bytes,1,opt,name=gw_record_id,json=gwRecordId,proto3" json:"gw_record_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LicenseNo  string `protobuf:"bytes,3,opt,name=licenseNo,proto3" json:"licenseNo,omitempty"`
	Cid        string `protobuf:"bytes,4,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *DoctorResponse_Doctor) Reset() {
	*x = DoctorResponse_Doctor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorResponse_Doctor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorResponse_Doctor) ProtoMessage() {}

func (x *DoctorResponse_Doctor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorResponse_Doctor.ProtoReflect.Descriptor instead.
func (*DoctorResponse_Doctor) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{5, 0}
}

func (x *DoctorResponse_Doctor) GetGwRecordId() string {
	if x != nil {
		return x.GwRecordId
	}
	return ""
}

func (x *DoctorResponse_Doctor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoctorResponse_Doctor) GetLicenseNo() string {
	if x != nil {
		return x.LicenseNo
	}
	return ""
}

func (x *DoctorResponse_Doctor) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type ScreeningResponse_Screening struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GwRecordId string `protobuf:"bytes,1,opt,name=gw_record_id,json=gwRecordId,proto3" json:"gw_record_id,omitempty"`
	Hn         string `protobuf:"bytes,2,opt,name=hn,proto3" json:"hn,omitempty"`
	Vn         string `protobuf:"bytes,3,opt,name=vn,proto3" json:"vn,omitempty"`
	Vstdate    string `protobuf:"bytes,4,opt,name=vstdate,proto3" json:"vstdate,omitempty"`
	Vsttime    string `protobuf:"bytes,5,opt,name=vsttime,proto3" json:"vsttime,omitempty"`
	Hospname   string `protobuf:"bytes,6,opt,name=hospname,proto3" json:"hospname,omitempty"`
	GwHospcode string `protobuf:"bytes,7,opt,name=gw_hospcode,json=gwHospcode,proto3" json:"gw_hospcode,omitempty"`
	Diagename  string `protobuf:"bytes,8,opt,name=diagename,proto3" json:"diagename,omitempty"`
	Diagtname  string `protobuf:"bytes,9,opt,name=diagtname,proto3" json:"diagtname,omitempty"`
}

func (x *ScreeningResponse_Screening) Reset() {
	*x = ScreeningResponse_Screening{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hosxp_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreeningResponse_Screening) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreeningResponse_Screening) ProtoMessage() {}

func (x *ScreeningResponse_Screening) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hosxp_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreeningResponse_Screening.ProtoReflect.Descriptor instead.
func (*ScreeningResponse_Screening) Descriptor() ([]byte, []int) {
	return file_proto_hosxp_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ScreeningResponse_Screening) GetGwRecordId() string {
	if x != nil {
		return x.GwRecordId
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetHn() string {
	if x != nil {
		return x.Hn
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetVn() string {
	if x != nil {
		return x.Vn
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetVstdate() string {
	if x != nil {
		return x.Vstdate
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetVsttime() string {
	if x != nil {
		return x.Vsttime
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetHospname() string {
	if x != nil {
		return x.Hospname
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetGwHospcode() string {
	if x != nil {
		return x.GwHospcode
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetDiagename() string {
	if x != nil {
		return x.Diagename
	}
	return ""
}

func (x *ScreeningResponse_Screening) GetDiagtname() string {
	if x != nil {
		return x.Diagtname
	}
	return ""
}

var File_proto_hosxp_proto protoreflect.FileDescriptor

var file_proto_hosxp_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x73, 0x78, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x0e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x9f, 0x02, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0xda, 0x01, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x77, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x68, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x48, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x77, 0x5f, 0x68,
	0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x77, 0x48, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x22, 0xd8, 0x02, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x88, 0x02, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x77, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x76, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x77, 0x5f, 0x68, 0x6f,
	0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x77,
	0x48, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x73, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x73, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x74,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x6e, 0x6f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x74, 0x74, 0x79, 0x70, 0x65, 0x6e, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x63, 0x6c, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x70, 0x63, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x70,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x70,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a,
	0x6e, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x77, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x67, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22,
	0xd4, 0x02, 0x0a, 0x11, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0xfa, 0x01, 0x0a, 0x09, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x77, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x76, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x73, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x73, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x77, 0x5f,
	0x68, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x77, 0x48, 0x6f, 0x73, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x61, 0x67, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x69, 0x61, 0x67, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x61, 0x67,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x61,
	0x67, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xc4, 0x01, 0x0a, 0x0a, 0x45, 0x6d, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x69, 0x64, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x69, 0x64,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x4e, 0x0a,
	0x0d, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x70,
	0x63, 0x6f, 0x64, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hosxp_proto_rawDescOnce sync.Once
	file_proto_hosxp_proto_rawDescData = file_proto_hosxp_proto_rawDesc
)

func file_proto_hosxp_proto_rawDescGZIP() []byte {
	file_proto_hosxp_proto_rawDescOnce.Do(func() {
		file_proto_hosxp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hosxp_proto_rawDescData)
	})
	return file_proto_hosxp_proto_rawDescData
}

var file_proto_hosxp_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_hosxp_proto_goTypes = []interface{}{
	(*RequestCid)(nil),                  // 0: proto.RequestCid
	(*RequestHospcode)(nil),             // 1: proto.RequestHospcode
	(*RequestPatient)(nil),              // 2: proto.RequestPatient
	(*InfoResponse)(nil),                // 3: proto.InfoResponse
	(*ServiceResponse)(nil),             // 4: proto.ServiceResponse
	(*DoctorResponse)(nil),              // 5: proto.DoctorResponse
	(*ScreeningResponse)(nil),           // 6: proto.ScreeningResponse
	(*InfoResponse_Info)(nil),           // 7: proto.InfoResponse.Info
	(*ServiceResponse_Service)(nil),     // 8: proto.ServiceResponse.Service
	(*DoctorResponse_Doctor)(nil),       // 9: proto.DoctorResponse.Doctor
	(*ScreeningResponse_Screening)(nil), // 10: proto.ScreeningResponse.Screening
}
var file_proto_hosxp_proto_depIdxs = []int32{
	7,  // 0: proto.InfoResponse.results:type_name -> proto.InfoResponse.Info
	8,  // 1: proto.ServiceResponse.services:type_name -> proto.ServiceResponse.Service
	9,  // 2: proto.DoctorResponse.Results:type_name -> proto.DoctorResponse.Doctor
	10, // 3: proto.ScreeningResponse.screenings:type_name -> proto.ScreeningResponse.Screening
	0,  // 4: proto.EmrService.PatientInfo:input_type -> proto.RequestCid
	0,  // 5: proto.EmrService.GetServices:input_type -> proto.RequestCid
	2,  // 6: proto.EmrService.GetScreening:input_type -> proto.RequestPatient
	1,  // 7: proto.DoctorService.DoctorList:input_type -> proto.RequestHospcode
	3,  // 8: proto.EmrService.PatientInfo:output_type -> proto.InfoResponse
	4,  // 9: proto.EmrService.GetServices:output_type -> proto.ServiceResponse
	6,  // 10: proto.EmrService.GetScreening:output_type -> proto.ScreeningResponse
	5,  // 11: proto.DoctorService.DoctorList:output_type -> proto.DoctorResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_hosxp_proto_init() }
func file_proto_hosxp_proto_init() {
	if File_proto_hosxp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hosxp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCid); i {
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
		file_proto_hosxp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHospcode); i {
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
		file_proto_hosxp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPatient); i {
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
		file_proto_hosxp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_proto_hosxp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
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
		file_proto_hosxp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorResponse); i {
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
		file_proto_hosxp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreeningResponse); i {
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
		file_proto_hosxp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse_Info); i {
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
		file_proto_hosxp_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse_Service); i {
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
		file_proto_hosxp_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorResponse_Doctor); i {
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
		file_proto_hosxp_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreeningResponse_Screening); i {
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
			RawDescriptor: file_proto_hosxp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_hosxp_proto_goTypes,
		DependencyIndexes: file_proto_hosxp_proto_depIdxs,
		MessageInfos:      file_proto_hosxp_proto_msgTypes,
	}.Build()
	File_proto_hosxp_proto = out.File
	file_proto_hosxp_proto_rawDesc = nil
	file_proto_hosxp_proto_goTypes = nil
	file_proto_hosxp_proto_depIdxs = nil
}
