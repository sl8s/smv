import 'package:meta/meta.dart';
import 'package:smv_dart/src/base_exception.dart';
import 'package:smv_dart/src/exception_adapter.dart';

@immutable
final class Result {
  final ExceptionAdapter exceptionAdapter;

  Result.success() : exceptionAdapter = ExceptionAdapter(null);
  Result.exception(BaseException exception)
      : exceptionAdapter = ExceptionAdapter(exception);
}
