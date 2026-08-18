import 'package:meta/meta.dart';
import 'package:smv_dart/src/utility.dart';

@immutable
abstract base class BaseException {
  final String _source;
  
  const BaseException(this._source);

  @override
  String toString();
  
  @protected
  @nonVirtual
  void initToConstructor() {
	redPrint("\n");
    redPrint("Source: ${this._source}");
    redPrint("toString(): ${this.toString()}");
    redPrint("\n");
  }
}
