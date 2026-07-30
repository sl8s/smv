import 'package:meta/meta.dart';
import 'package:smv_dart/src/base_exception.dart';
import 'package:smv_dart/src/model/base_model.dart';
import 'package:smv_dart/src/exception_adapter.dart';

@immutable
final class ResultModel<T extends BaseModel> {
  final T? data;
  final ExceptionAdapter exceptionAdapter;

  ResultModel.success(this.data)
      : exceptionAdapter = ExceptionAdapter(null);
  ResultModel.exception(BaseException exception)
      : data = null,
        exceptionAdapter = ExceptionAdapter(exception);
}
