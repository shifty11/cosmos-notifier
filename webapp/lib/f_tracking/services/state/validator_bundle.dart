import 'package:cosmos_notifier/api/protobuf/dart/tracker_service.pb.dart' as pb;
import 'package:fixnum/fixnum.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'validator_bundle.freezed.dart';

@freezed
class FreezedValidator with _$FreezedValidator {
  const FreezedValidator._();

  const factory FreezedValidator({
    required Int64 id,
    required String address,
    required String chainName,
  }) = _Validator;

  static FreezedValidator fromProtobuf(pb.Validator validator) {
    return FreezedValidator(id: validator.id, address: validator.address, chainName: validator.chainName);
  }
}

@freezed
class FreezedValidatorBundle with _$FreezedValidatorBundle implements Comparable<FreezedValidatorBundle> {
  const FreezedValidatorBundle._();

  const factory FreezedValidatorBundle({required String moniker, required List<FreezedValidator> validators, required bool isTracked}) = _ValidatorBundle;

  static FreezedValidatorBundle fromProtobuf(pb.ValidatorBundle bundle) {
    return FreezedValidatorBundle(
      moniker: bundle.moniker,
      validators: bundle.validators.map((e) => FreezedValidator.fromProtobuf(e)).toList(),
      isTracked: bundle.isTracked,
    );
  }

  @override
  int compareTo(FreezedValidatorBundle other) {
    if (isTracked && !other.isTracked) {
      return -1;
    } else if (!isTracked && other.isTracked) {
      return 1;
    } else {
      // compare length of validators and if they are equal compare moniker
      final validatorsLengthComparison = other.validators.length.compareTo(validators.length);
      return validatorsLengthComparison == 0 ? moniker.compareTo(other.moniker) : validatorsLengthComparison;
    }
  }
}
