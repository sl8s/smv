import 'package:meta/meta.dart';
import 'package:smv_dart/src/base_exception.dart';

@immutable
final class ExceptionAdapter {
  final BaseException? exception;

  const ExceptionAdapter(this.exception);

  bool hasException() {
    return exception != null;
  }
}
