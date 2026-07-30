import 'package:meta/meta.dart';
import 'package:smv_dart/src/base_exception.dart';
import 'package:smv_dart/src/model/base_list_model.dart';
import 'package:smv_dart/src/exception_adapter.dart';

@immutable
final class ResultListModel<T extends BaseListModel> {
  final T? data;
  final ExceptionAdapter exceptionAdapter;

  ResultListModel.success(this.data)
      : exceptionAdapter = ExceptionAdapter(null);
  ResultListModel.exception(BaseException exception)
      : data = null,
        exceptionAdapter = ExceptionAdapter(exception);
}
