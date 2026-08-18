import 'package:meta/meta.dart';
import 'package:smv_dart/src/base_exception.dart';
import 'package:smv_dart/src/enum_guilty.dart';

@immutable
final class LocalException extends BaseException {
  final EnumGuilty guilty;
  final String message;

  LocalException({required super.source, required this.guilty, required this.message}) {
    this.initToConstructor();
  }

  @override
  String toString() {
    return "LocalException(guilty: ${guilty.name}, "
        "message: $message)";
  }
}
