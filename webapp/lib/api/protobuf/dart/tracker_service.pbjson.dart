///
//  Generated code. Do not modify.
//  source: tracker_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use trackerChatRoomDescriptor instead')
const TrackerChatRoom$json = const {
  '1': 'TrackerChatRoom',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'discord', '3': 2, '4': 1, '5': 11, '6': '.cosmos_notifier_grpc.DiscordType', '9': 0, '10': 'discord'},
    const {'1': 'telegram', '3': 3, '4': 1, '5': 11, '6': '.cosmos_notifier_grpc.TelegramType', '9': 0, '10': 'telegram'},
  ],
  '8': const [
    const {'1': 'type'},
  ],
};

/// Descriptor for `TrackerChatRoom`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List trackerChatRoomDescriptor = $convert.base64Decode('Cg9UcmFja2VyQ2hhdFJvb20SEgoEbmFtZRgBIAEoCVIEbmFtZRI9CgdkaXNjb3JkGAIgASgLMiEuY29zbW9zX25vdGlmaWVyX2dycGMuRGlzY29yZFR5cGVIAFIHZGlzY29yZBJACgh0ZWxlZ3JhbRgDIAEoCzIiLmNvc21vc19ub3RpZmllcl9ncnBjLlRlbGVncmFtVHlwZUgAUgh0ZWxlZ3JhbUIGCgR0eXBl');
@$core.Deprecated('Use trackerDescriptor instead')
const Tracker$json = const {
  '1': 'Tracker',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'address', '3': 2, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'notificationInterval', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'notificationInterval'},
    const {'1': 'chatRoom', '3': 4, '4': 1, '5': 11, '6': '.cosmos_notifier_grpc.TrackerChatRoom', '10': 'chatRoom'},
    const {'1': 'updatedAt', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    const {'1': 'validatorMoniker', '3': 6, '4': 1, '5': 9, '10': 'validatorMoniker'},
  ],
};

/// Descriptor for `Tracker`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List trackerDescriptor = $convert.base64Decode('CgdUcmFja2VyEg4KAmlkGAEgASgDUgJpZBIYCgdhZGRyZXNzGAIgASgJUgdhZGRyZXNzEk0KFG5vdGlmaWNhdGlvbkludGVydmFsGAMgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uUhRub3RpZmljYXRpb25JbnRlcnZhbBJBCghjaGF0Um9vbRgEIAEoCzIlLmNvc21vc19ub3RpZmllcl9ncnBjLlRyYWNrZXJDaGF0Um9vbVIIY2hhdFJvb20SOAoJdXBkYXRlZEF0GAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIJdXBkYXRlZEF0EioKEHZhbGlkYXRvck1vbmlrZXIYBiABKAlSEHZhbGlkYXRvck1vbmlrZXI=');
@$core.Deprecated('Use validatorDescriptor instead')
const Validator$json = const {
  '1': 'Validator',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'address', '3': 2, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'chain_name', '3': 3, '4': 1, '5': 9, '10': 'chainName'},
  ],
};

/// Descriptor for `Validator`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List validatorDescriptor = $convert.base64Decode('CglWYWxpZGF0b3ISDgoCaWQYASABKANSAmlkEhgKB2FkZHJlc3MYAiABKAlSB2FkZHJlc3MSHQoKY2hhaW5fbmFtZRgDIAEoCVIJY2hhaW5OYW1l');
@$core.Deprecated('Use validatorBundleDescriptor instead')
const ValidatorBundle$json = const {
  '1': 'ValidatorBundle',
  '2': const [
    const {'1': 'moniker', '3': 1, '4': 1, '5': 9, '10': 'moniker'},
    const {'1': 'validators', '3': 2, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.Validator', '10': 'validators'},
    const {'1': 'is_tracked', '3': 3, '4': 1, '5': 8, '10': 'isTracked'},
  ],
};

/// Descriptor for `ValidatorBundle`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List validatorBundleDescriptor = $convert.base64Decode('Cg9WYWxpZGF0b3JCdW5kbGUSGAoHbW9uaWtlchgBIAEoCVIHbW9uaWtlchI/Cgp2YWxpZGF0b3JzGAIgAygLMh8uY29zbW9zX25vdGlmaWVyX2dycGMuVmFsaWRhdG9yUgp2YWxpZGF0b3JzEh0KCmlzX3RyYWNrZWQYAyABKAhSCWlzVHJhY2tlZA==');
@$core.Deprecated('Use listTrackersResponseDescriptor instead')
const ListTrackersResponse$json = const {
  '1': 'ListTrackersResponse',
  '2': const [
    const {'1': 'trackers', '3': 1, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.Tracker', '10': 'trackers'},
    const {'1': 'chatRooms', '3': 2, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.TrackerChatRoom', '10': 'chatRooms'},
    const {'1': 'validator_bundles', '3': 3, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.ValidatorBundle', '10': 'validatorBundles'},
  ],
};

/// Descriptor for `ListTrackersResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTrackersResponseDescriptor = $convert.base64Decode('ChRMaXN0VHJhY2tlcnNSZXNwb25zZRI5Cgh0cmFja2VycxgBIAMoCzIdLmNvc21vc19ub3RpZmllcl9ncnBjLlRyYWNrZXJSCHRyYWNrZXJzEkMKCWNoYXRSb29tcxgCIAMoCzIlLmNvc21vc19ub3RpZmllcl9ncnBjLlRyYWNrZXJDaGF0Um9vbVIJY2hhdFJvb21zElIKEXZhbGlkYXRvcl9idW5kbGVzGAMgAygLMiUuY29zbW9zX25vdGlmaWVyX2dycGMuVmFsaWRhdG9yQnVuZGxlUhB2YWxpZGF0b3JCdW5kbGVz');
@$core.Deprecated('Use isAddressValidRequestDescriptor instead')
const IsAddressValidRequest$json = const {
  '1': 'IsAddressValidRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
  ],
};

/// Descriptor for `IsAddressValidRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List isAddressValidRequestDescriptor = $convert.base64Decode('ChVJc0FkZHJlc3NWYWxpZFJlcXVlc3QSGAoHYWRkcmVzcxgBIAEoCVIHYWRkcmVzcw==');
@$core.Deprecated('Use isAddressValidResponseDescriptor instead')
const IsAddressValidResponse$json = const {
  '1': 'IsAddressValidResponse',
  '2': const [
    const {'1': 'isValid', '3': 1, '4': 1, '5': 8, '10': 'isValid'},
  ],
};

/// Descriptor for `IsAddressValidResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List isAddressValidResponseDescriptor = $convert.base64Decode('ChZJc0FkZHJlc3NWYWxpZFJlc3BvbnNlEhgKB2lzVmFsaWQYASABKAhSB2lzVmFsaWQ=');
@$core.Deprecated('Use createTrackerRequestDescriptor instead')
const CreateTrackerRequest$json = const {
  '1': 'CreateTrackerRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'notificationInterval', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'notificationInterval'},
    const {'1': 'chatRoom', '3': 3, '4': 1, '5': 11, '6': '.cosmos_notifier_grpc.TrackerChatRoom', '10': 'chatRoom'},
  ],
};

/// Descriptor for `CreateTrackerRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTrackerRequestDescriptor = $convert.base64Decode('ChRDcmVhdGVUcmFja2VyUmVxdWVzdBIYCgdhZGRyZXNzGAEgASgJUgdhZGRyZXNzEk0KFG5vdGlmaWNhdGlvbkludGVydmFsGAIgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uUhRub3RpZmljYXRpb25JbnRlcnZhbBJBCghjaGF0Um9vbRgDIAEoCzIlLmNvc21vc19ub3RpZmllcl9ncnBjLlRyYWNrZXJDaGF0Um9vbVIIY2hhdFJvb20=');
@$core.Deprecated('Use updateTrackerRequestDescriptor instead')
const UpdateTrackerRequest$json = const {
  '1': 'UpdateTrackerRequest',
  '2': const [
    const {'1': 'trackerId', '3': 1, '4': 1, '5': 3, '10': 'trackerId'},
    const {'1': 'notificationInterval', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'notificationInterval'},
    const {'1': 'chatRoom', '3': 3, '4': 1, '5': 11, '6': '.cosmos_notifier_grpc.TrackerChatRoom', '10': 'chatRoom'},
  ],
};

/// Descriptor for `UpdateTrackerRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateTrackerRequestDescriptor = $convert.base64Decode('ChRVcGRhdGVUcmFja2VyUmVxdWVzdBIcCgl0cmFja2VySWQYASABKANSCXRyYWNrZXJJZBJNChRub3RpZmljYXRpb25JbnRlcnZhbBgCIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvblIUbm90aWZpY2F0aW9uSW50ZXJ2YWwSQQoIY2hhdFJvb20YAyABKAsyJS5jb3Ntb3Nfbm90aWZpZXJfZ3JwYy5UcmFja2VyQ2hhdFJvb21SCGNoYXRSb29t');
@$core.Deprecated('Use deleteTrackerRequestDescriptor instead')
const DeleteTrackerRequest$json = const {
  '1': 'DeleteTrackerRequest',
  '2': const [
    const {'1': 'trackerId', '3': 1, '4': 1, '5': 3, '10': 'trackerId'},
  ],
};

/// Descriptor for `DeleteTrackerRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTrackerRequestDescriptor = $convert.base64Decode('ChREZWxldGVUcmFja2VyUmVxdWVzdBIcCgl0cmFja2VySWQYASABKANSCXRyYWNrZXJJZA==');
@$core.Deprecated('Use trackValidatorsRequestDescriptor instead')
const TrackValidatorsRequest$json = const {
  '1': 'TrackValidatorsRequest',
  '2': const [
    const {'1': 'monikers', '3': 1, '4': 3, '5': 9, '10': 'monikers'},
    const {'1': 'notificationInterval', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'notificationInterval'},
    const {'1': 'chatRoom', '3': 3, '4': 1, '5': 11, '6': '.cosmos_notifier_grpc.TrackerChatRoom', '10': 'chatRoom'},
  ],
};

/// Descriptor for `TrackValidatorsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List trackValidatorsRequestDescriptor = $convert.base64Decode('ChZUcmFja1ZhbGlkYXRvcnNSZXF1ZXN0EhoKCG1vbmlrZXJzGAEgAygJUghtb25pa2VycxJNChRub3RpZmljYXRpb25JbnRlcnZhbBgCIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvblIUbm90aWZpY2F0aW9uSW50ZXJ2YWwSQQoIY2hhdFJvb20YAyABKAsyJS5jb3Ntb3Nfbm90aWZpZXJfZ3JwYy5UcmFja2VyQ2hhdFJvb21SCGNoYXRSb29t');
@$core.Deprecated('Use trackValidatorsResponseDescriptor instead')
const TrackValidatorsResponse$json = const {
  '1': 'TrackValidatorsResponse',
  '2': const [
    const {'1': 'addedTrackers', '3': 1, '4': 3, '5': 11, '6': '.cosmos_notifier_grpc.Tracker', '10': 'addedTrackers'},
    const {'1': 'deletedTrackerIds', '3': 2, '4': 3, '5': 3, '10': 'deletedTrackerIds'},
  ],
};

/// Descriptor for `TrackValidatorsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List trackValidatorsResponseDescriptor = $convert.base64Decode('ChdUcmFja1ZhbGlkYXRvcnNSZXNwb25zZRJDCg1hZGRlZFRyYWNrZXJzGAEgAygLMh0uY29zbW9zX25vdGlmaWVyX2dycGMuVHJhY2tlclINYWRkZWRUcmFja2VycxIsChFkZWxldGVkVHJhY2tlcklkcxgCIAMoA1IRZGVsZXRlZFRyYWNrZXJJZHM=');
