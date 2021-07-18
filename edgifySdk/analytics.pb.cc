// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: analytics.proto

#include "analytics.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// This is a temporary google only hack
#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
#include "third_party/protobuf/version.h"
#endif
// @@protoc_insertion_point(includes)

namespace edgify {
class CreateAnalyticsEventRequestDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<CreateAnalyticsEventRequest>
      _instance;
} _CreateAnalyticsEventRequest_default_instance_;
class CreateAnalyticsEventResponseDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<CreateAnalyticsEventResponse>
      _instance;
} _CreateAnalyticsEventResponse_default_instance_;
}  // namespace edgify
namespace protobuf_analytics_2eproto {
static void InitDefaultsCreateAnalyticsEventRequest() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::edgify::_CreateAnalyticsEventRequest_default_instance_;
    new (ptr) ::edgify::CreateAnalyticsEventRequest();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::edgify::CreateAnalyticsEventRequest::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_CreateAnalyticsEventRequest =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsCreateAnalyticsEventRequest}, {}};

static void InitDefaultsCreateAnalyticsEventResponse() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::edgify::_CreateAnalyticsEventResponse_default_instance_;
    new (ptr) ::edgify::CreateAnalyticsEventResponse();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::edgify::CreateAnalyticsEventResponse::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_CreateAnalyticsEventResponse =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsCreateAnalyticsEventResponse}, {}};

void InitDefaults() {
  ::google::protobuf::internal::InitSCC(&scc_info_CreateAnalyticsEventRequest.base);
  ::google::protobuf::internal::InitSCC(&scc_info_CreateAnalyticsEventResponse.base);
}

::google::protobuf::Metadata file_level_metadata[2];

const ::google::protobuf::uint32 TableStruct::offsets[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::edgify::CreateAnalyticsEventRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::edgify::CreateAnalyticsEventRequest, name_),
  ~0u,  // no _has_bits_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::edgify::CreateAnalyticsEventResponse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
};
static const ::google::protobuf::internal::MigrationSchema schemas[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::edgify::CreateAnalyticsEventRequest)},
  { 6, -1, sizeof(::edgify::CreateAnalyticsEventResponse)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::edgify::_CreateAnalyticsEventRequest_default_instance_),
  reinterpret_cast<const ::google::protobuf::Message*>(&::edgify::_CreateAnalyticsEventResponse_default_instance_),
};

void protobuf_AssignDescriptors() {
  AddDescriptors();
  AssignDescriptors(
      "analytics.proto", schemas, file_default_instances, TableStruct::offsets,
      file_level_metadata, NULL, NULL);
}

void protobuf_AssignDescriptorsOnce() {
  static ::google::protobuf::internal::once_flag once;
  ::google::protobuf::internal::call_once(once, protobuf_AssignDescriptors);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_PROTOBUF_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::internal::RegisterAllTypes(file_level_metadata, 2);
}

void AddDescriptorsImpl() {
  InitDefaults();
  static const char descriptor[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
      "\n\017analytics.proto\022\006edgify\"+\n\033CreateAnaly"
      "ticsEventRequest\022\014\n\004name\030\001 \001(\t\"\036\n\034Create"
      "AnalyticsEventResponse2n\n\020AnalyticsServi"
      "ce\022Z\n\013CreateEvent\022#.edgify.CreateAnalyti"
      "csEventRequest\032$.edgify.CreateAnalyticsE"
      "ventResponse\"\000B\nZ\010edgifypbb\006proto3"
  };
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
      descriptor, 234);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "analytics.proto", &protobuf_RegisterTypes);
}

void AddDescriptors() {
  static ::google::protobuf::internal::once_flag once;
  ::google::protobuf::internal::call_once(once, AddDescriptorsImpl);
}
// Force AddDescriptors() to be called at dynamic initialization time.
struct StaticDescriptorInitializer {
  StaticDescriptorInitializer() {
    AddDescriptors();
  }
} static_descriptor_initializer;
}  // namespace protobuf_analytics_2eproto
namespace edgify {

// ===================================================================

void CreateAnalyticsEventRequest::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int CreateAnalyticsEventRequest::kNameFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

CreateAnalyticsEventRequest::CreateAnalyticsEventRequest()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  ::google::protobuf::internal::InitSCC(
      &protobuf_analytics_2eproto::scc_info_CreateAnalyticsEventRequest.base);
  SharedCtor();
  // @@protoc_insertion_point(constructor:edgify.CreateAnalyticsEventRequest)
}
CreateAnalyticsEventRequest::CreateAnalyticsEventRequest(const CreateAnalyticsEventRequest& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.name().size() > 0) {
    name_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.name_);
  }
  // @@protoc_insertion_point(copy_constructor:edgify.CreateAnalyticsEventRequest)
}

void CreateAnalyticsEventRequest::SharedCtor() {
  name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

CreateAnalyticsEventRequest::~CreateAnalyticsEventRequest() {
  // @@protoc_insertion_point(destructor:edgify.CreateAnalyticsEventRequest)
  SharedDtor();
}

void CreateAnalyticsEventRequest::SharedDtor() {
  name_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

void CreateAnalyticsEventRequest::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const ::google::protobuf::Descriptor* CreateAnalyticsEventRequest::descriptor() {
  ::protobuf_analytics_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_analytics_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const CreateAnalyticsEventRequest& CreateAnalyticsEventRequest::default_instance() {
  ::google::protobuf::internal::InitSCC(&protobuf_analytics_2eproto::scc_info_CreateAnalyticsEventRequest.base);
  return *internal_default_instance();
}


void CreateAnalyticsEventRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:edgify.CreateAnalyticsEventRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  _internal_metadata_.Clear();
}

bool CreateAnalyticsEventRequest::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:edgify.CreateAnalyticsEventRequest)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // string name = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(10u /* 10 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_name()));
          DO_(::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
            this->name().data(), static_cast<int>(this->name().length()),
            ::google::protobuf::internal::WireFormatLite::PARSE,
            "edgify.CreateAnalyticsEventRequest.name"));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:edgify.CreateAnalyticsEventRequest)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:edgify.CreateAnalyticsEventRequest)
  return false;
#undef DO_
}

void CreateAnalyticsEventRequest::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:edgify.CreateAnalyticsEventRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string name = 1;
  if (this->name().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->name().data(), static_cast<int>(this->name().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "edgify.CreateAnalyticsEventRequest.name");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      1, this->name(), output);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:edgify.CreateAnalyticsEventRequest)
}

::google::protobuf::uint8* CreateAnalyticsEventRequest::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:edgify.CreateAnalyticsEventRequest)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string name = 1;
  if (this->name().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->name().data(), static_cast<int>(this->name().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "edgify.CreateAnalyticsEventRequest.name");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        1, this->name(), target);
  }

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:edgify.CreateAnalyticsEventRequest)
  return target;
}

size_t CreateAnalyticsEventRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:edgify.CreateAnalyticsEventRequest)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  // string name = 1;
  if (this->name().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->name());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void CreateAnalyticsEventRequest::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:edgify.CreateAnalyticsEventRequest)
  GOOGLE_DCHECK_NE(&from, this);
  const CreateAnalyticsEventRequest* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const CreateAnalyticsEventRequest>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:edgify.CreateAnalyticsEventRequest)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:edgify.CreateAnalyticsEventRequest)
    MergeFrom(*source);
  }
}

void CreateAnalyticsEventRequest::MergeFrom(const CreateAnalyticsEventRequest& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:edgify.CreateAnalyticsEventRequest)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.name().size() > 0) {

    name_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.name_);
  }
}

void CreateAnalyticsEventRequest::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:edgify.CreateAnalyticsEventRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void CreateAnalyticsEventRequest::CopyFrom(const CreateAnalyticsEventRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:edgify.CreateAnalyticsEventRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool CreateAnalyticsEventRequest::IsInitialized() const {
  return true;
}

void CreateAnalyticsEventRequest::Swap(CreateAnalyticsEventRequest* other) {
  if (other == this) return;
  InternalSwap(other);
}
void CreateAnalyticsEventRequest::InternalSwap(CreateAnalyticsEventRequest* other) {
  using std::swap;
  name_.Swap(&other->name_, &::google::protobuf::internal::GetEmptyStringAlreadyInited(),
    GetArenaNoVirtual());
  _internal_metadata_.Swap(&other->_internal_metadata_);
}

::google::protobuf::Metadata CreateAnalyticsEventRequest::GetMetadata() const {
  protobuf_analytics_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_analytics_2eproto::file_level_metadata[kIndexInFileMessages];
}


// ===================================================================

void CreateAnalyticsEventResponse::InitAsDefaultInstance() {
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

CreateAnalyticsEventResponse::CreateAnalyticsEventResponse()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  ::google::protobuf::internal::InitSCC(
      &protobuf_analytics_2eproto::scc_info_CreateAnalyticsEventResponse.base);
  SharedCtor();
  // @@protoc_insertion_point(constructor:edgify.CreateAnalyticsEventResponse)
}
CreateAnalyticsEventResponse::CreateAnalyticsEventResponse(const CreateAnalyticsEventResponse& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:edgify.CreateAnalyticsEventResponse)
}

void CreateAnalyticsEventResponse::SharedCtor() {
}

CreateAnalyticsEventResponse::~CreateAnalyticsEventResponse() {
  // @@protoc_insertion_point(destructor:edgify.CreateAnalyticsEventResponse)
  SharedDtor();
}

void CreateAnalyticsEventResponse::SharedDtor() {
}

void CreateAnalyticsEventResponse::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const ::google::protobuf::Descriptor* CreateAnalyticsEventResponse::descriptor() {
  ::protobuf_analytics_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_analytics_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const CreateAnalyticsEventResponse& CreateAnalyticsEventResponse::default_instance() {
  ::google::protobuf::internal::InitSCC(&protobuf_analytics_2eproto::scc_info_CreateAnalyticsEventResponse.base);
  return *internal_default_instance();
}


void CreateAnalyticsEventResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:edgify.CreateAnalyticsEventResponse)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _internal_metadata_.Clear();
}

bool CreateAnalyticsEventResponse::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:edgify.CreateAnalyticsEventResponse)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
  handle_unusual:
    if (tag == 0) {
      goto success;
    }
    DO_(::google::protobuf::internal::WireFormat::SkipField(
          input, tag, _internal_metadata_.mutable_unknown_fields()));
  }
success:
  // @@protoc_insertion_point(parse_success:edgify.CreateAnalyticsEventResponse)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:edgify.CreateAnalyticsEventResponse)
  return false;
#undef DO_
}

void CreateAnalyticsEventResponse::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:edgify.CreateAnalyticsEventResponse)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), output);
  }
  // @@protoc_insertion_point(serialize_end:edgify.CreateAnalyticsEventResponse)
}

::google::protobuf::uint8* CreateAnalyticsEventResponse::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:edgify.CreateAnalyticsEventResponse)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:edgify.CreateAnalyticsEventResponse)
  return target;
}

size_t CreateAnalyticsEventResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:edgify.CreateAnalyticsEventResponse)
  size_t total_size = 0;

  if ((_internal_metadata_.have_unknown_fields() &&  ::google::protobuf::internal::GetProto3PreserveUnknownsDefault())) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        (::google::protobuf::internal::GetProto3PreserveUnknownsDefault()   ? _internal_metadata_.unknown_fields()   : _internal_metadata_.default_instance()));
  }
  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void CreateAnalyticsEventResponse::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:edgify.CreateAnalyticsEventResponse)
  GOOGLE_DCHECK_NE(&from, this);
  const CreateAnalyticsEventResponse* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const CreateAnalyticsEventResponse>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:edgify.CreateAnalyticsEventResponse)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:edgify.CreateAnalyticsEventResponse)
    MergeFrom(*source);
  }
}

void CreateAnalyticsEventResponse::MergeFrom(const CreateAnalyticsEventResponse& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:edgify.CreateAnalyticsEventResponse)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

}

void CreateAnalyticsEventResponse::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:edgify.CreateAnalyticsEventResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void CreateAnalyticsEventResponse::CopyFrom(const CreateAnalyticsEventResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:edgify.CreateAnalyticsEventResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool CreateAnalyticsEventResponse::IsInitialized() const {
  return true;
}

void CreateAnalyticsEventResponse::Swap(CreateAnalyticsEventResponse* other) {
  if (other == this) return;
  InternalSwap(other);
}
void CreateAnalyticsEventResponse::InternalSwap(CreateAnalyticsEventResponse* other) {
  using std::swap;
  _internal_metadata_.Swap(&other->_internal_metadata_);
}

::google::protobuf::Metadata CreateAnalyticsEventResponse::GetMetadata() const {
  protobuf_analytics_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_analytics_2eproto::file_level_metadata[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace edgify
namespace google {
namespace protobuf {
template<> GOOGLE_PROTOBUF_ATTRIBUTE_NOINLINE ::edgify::CreateAnalyticsEventRequest* Arena::CreateMaybeMessage< ::edgify::CreateAnalyticsEventRequest >(Arena* arena) {
  return Arena::CreateInternal< ::edgify::CreateAnalyticsEventRequest >(arena);
}
template<> GOOGLE_PROTOBUF_ATTRIBUTE_NOINLINE ::edgify::CreateAnalyticsEventResponse* Arena::CreateMaybeMessage< ::edgify::CreateAnalyticsEventResponse >(Arena* arena) {
  return Arena::CreateInternal< ::edgify::CreateAnalyticsEventResponse >(arena);
}
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
