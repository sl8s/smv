import 'package:meta/meta.dart';
import 'package:smv_dart/src/idispose.dart';

@immutable
abstract base class BaseModelRepository implements IDispose {
  const BaseModelRepository();

  @protected
  @nonVirtual
  T getSafeValue<T>(Map<String, dynamic> map, String key, T defaultValue) {
    if (!map.containsKey(key)) {
      return defaultValue;
    }
    return map[key];
  }
}
