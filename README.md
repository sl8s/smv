![SMV](./smv.png)

## Install

- [Go](#go)
- [Javascript/Typescript](#javascripttypescript)
- [Dart](#dart)

### Go

```sh
go get github.com/sl8s/smv/packages/smvgo/v2@v2.0.9
```

### Javascript/Typescript

```json
"smv_typescript": "https://github.com/sl8s/smv/releases/download/v2.0.7/smv_typescript_v2_0_4.tgz"
```

> Important
>
> At this time, this library on 'Javascript/Typescript' will not be supported.

> Note
>
> In the "BaseArrayModel" class, the "add", "updateById", "deleteById", "addFromArray", "updateFromArrayById", "deleteFromArrayById" methods do not check incoming data for duplicates.
>
> In the "BaseArrayModel" class, the "updateById", "deleteById", "updateFromArrayById", "deleteFromArrayById" methods use an unoptimized algorithm for searching for elements in the array, and there is no algorithm for updating and removing duplicates in the array.
>
> In the "BaseArrayModel" class, the "add", "updateById", "deleteById", "addFromArray", "updateFromArrayById", "deleteFromArrayById" methods must return the "Result" class for detailed error instead of called "throw" because unit tests shouldn't have extraneous console messages.
>
> In the "BaseArrayModel" class, the "add", "updateById", "deleteById", "addFromArray", "updateFromArrayById", "deleteFromArrayById" methods needs to be renamed to "addFromNewModel", "updateFromNewModelById", "deleteFromIdById", "addFromNewModels", "updateFromNewModelsById", "deleteFromIdsById".
>
> The "BaseArrayModel" class needs to be renamed to "BaseModels".
>
> In the "IterationService" class, it is necessary to use a map instead of an array for optimization.
>
> In the "ShareService", "ShareProxy" classes, the "deleteAllListenersByKey", "deleteAllListenersByArrayKey", "deleteListenerByListenerId", "deleteListenersByListenerId" methods are superfluous, because the only things needed to replace them are these methods: "deleteListener", "deleteListeners".
>
> In the "ShareService", "ShareProxy" classes, the "addListener", "deleteListener", "notifyListener", "notifyListeners" methods must return the "Result" class for error instead of called "throw" because unit tests shouldn't have extraneous console messages.
>
> In the "BaseException", "LocalException", "NetworkException" classes, the "initToConstructor" method must be removed.
>
> In the "BaseException", "LocalException", "NetworkException" classes, the "toString" method needs to be renamed to "error".
>
> Delete the "redPrint" method.
>
> The unit tests need to be rewritten to accommodate the new changes and must include comments ("Success", "First condition", "Second condition", etc..), as they provide more information than the source code itself due to the use of various input data and the verification of all conditions.

### Dart

```yaml
smv_dart:
  git:
    url: https://github.com/sl8s/smv.git
    ref: v2.0.7
    path: packages/smv_dart
```

> Important
>
> At this time, this library on 'Dart' will not be supported.

> Note
>
> In the "BaseListModel" class, the "add", "updateById", "deleteById", "addFromList", "updateFromListById", "deleteFromListById" methods do not check incoming data for duplicates.
>
> In the "BaseListModel" class, the "updateById", "deleteById", "updateFromListById", "deleteFromListById" methods use an unoptimized algorithm for searching for elements in the array, and there is no algorithm for updating and removing duplicates in the array.
>
> In the "BaseListModel" class, the "add", "updateById", "deleteById", "addFromList", "updateFromListById", "deleteFromListById" methods must return the "Result" class for detailed error instead of called "throw" because unit tests shouldn't have extraneous console messages.
>
> In the "BaseListModel" class, the "add", "updateById", "deleteById", "addFromList", "updateFromListById", "deleteFromListById" methods needs to be renamed to "addFromNewModel", "updateFromNewModelById", "deleteFromIdById", "addFromNewModels", "updateFromNewModelsById", "deleteFromIdsById".
>
> The "BaseListModel" class needs to be renamed to "BaseModels".
>
> In the "IterationService" class, it is necessary to use a map instead of an array for optimization.
>
> In the "ShareService", "ShareProxy" classes, the "deleteAllListenersByKey", "deleteAllListenersByListKey", "deleteListenerByListenerId", "deleteListenersByListenerId" methods are superfluous, because the only things needed to replace them are these methods: "deleteListener", "deleteListeners".
>
> In the "ShareService", "ShareProxy" classes, the "addListener", "deleteListener", "notifyListener", "notifyListeners" methods must return the "Result" class for error instead of called "throw" because unit tests shouldn't have extraneous console messages.
>
> In the "BaseException", "LocalException", "NetworkException" classes, the "initToConstructor" method must be removed.
>
> In the "BaseException", "LocalException", "NetworkException" classes, the "toString" method needs to be renamed to "error".
>
> Delete the "redPrint" method and the "utility.dart" file.
>
> The unit tests need to be rewritten to accommodate the new changes and must include comments ("Success", "First condition", "Second condition", etc..), as they provide more information than the source code itself due to the use of various input data and the verification of all conditions.