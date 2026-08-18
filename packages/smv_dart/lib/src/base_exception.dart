import 'package:meta/meta.dart';
import 'package:smv_dart/src/utility.dart';

@immutable
abstract base class BaseException {
  final String source;
  
  const BaseException({required this.source});

  @override
  String toString();
  
  @protected
  @nonVirtual
  void initToConstructor() {
	redPrint("\n");
    redPrint("Source: ${this.source}");
    redPrint("toString(): ${this.toString()}");
    redPrint("\n");
  }
}
