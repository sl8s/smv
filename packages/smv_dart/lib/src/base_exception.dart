import 'package:meta/meta.dart';
import 'package:smv_dart/src/utility.dart';

@immutable
abstract base class BaseException implements Exception {
  BaseException(
      {required Object source, required Type type}) {
    debugPrintException("\n===start_to_trace_exception===\n");
    debugPrintException("Source: ${source.runtimeType}");
    debugPrintException("Type: $type");
    debugPrintException("toString(): ${toString()}");
    debugPrintException("\n===end_to_trace_exception===\n");
  }

  @override
  String toString();
}
