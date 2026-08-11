![SMV](./smv.png)

---

## Install

- [JavaScript/TypeScript](#javascripttypescript)
- [Dart](#dart)

### JavaScript/TypeScript

```json
"smv_typescript": "https://github.com/sl8s/smv/releases/download/v2.0.3/smv_typescript_v2_0_3.tgz"
```

### Dart

```yaml
smv_dart:
  git:
    url: https://github.com/sl8s/smv.git
    ref: v2.0.1
    path: packages/smv_dart
```

## Design Patterns

- Class `BaseModel` - `Prototype`.
- Class `BaseArrayModel` or `BaseListModel` - `Prototype`.
- Class `BaseModelRepository` - `Repository`, `Dispose`.
- Class `IterationService` - `Singleton`.
- Class `ShareService` - `Singleton`, `Pub/Sub`.
- Class `ShareProxy` - `Pub/Sub`, `Proxy`.
- Class `BaseException` - `Abstract Factory`.
- Class `LocalException` - `Abstract Factory`.
- Class `NetworkException` - `Abstract Factory`.
- Class `ExceptionAdapter` - Resembling an `Adapter`.
- Class `IDispose` - `Dispose`.
- Class `Result` - `Result/Either`.
- Class `ResultModel` - `Result/Either`.
- Class `ResultArrayModel` or `ResultListModel` - `Result/Either`.