import 'package:meta/meta.dart';
import 'package:smv_dart/src/idispose.dart';
import 'package:smv_dart/src/model/base_model.dart';
import 'package:smv_dart/src/model/base_list_model.dart';

@immutable
abstract base class BaseModelRepository<Y extends BaseModel, U extends BaseListModel<Y>> implements IDispose {
  const BaseModelRepository();
  
  @protected
  Y fromMap(Map<String, dynamic> map);
  
  @protected
  U fromListMap(List<Map<String, dynamic>> listMap);
  
  @protected
  @nonVirtual
  T getSafeValue<T>(Map<String, dynamic> map, String key, T defaultValue) {
    if (!map.containsKey(key)) {
      return defaultValue;
    }
    return map[key];
  }
}
