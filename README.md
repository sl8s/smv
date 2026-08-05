## Install

- [JavaScript/TypeScript](#javascripttypescript)
- [Dart](#dart)

### JavaScript/TypeScript

- package.json

```json
"smv_typescript": "https://github.com/sl8s/smv/releases/download/v2.0.1/smv_typescript_v2_0_1.tgz"
```

### Dart

- pubspec.yaml

```yaml
smv_dart:
  git:
    url: https://github.com/sl8s/smv.git
    ref: v2.0.1
    path: packages/smv_dart
```

## SMV

- SMV (Share, Model, View) - two-layer architectural design pattern. Built based on the three-layer architectural design pattern [library_architecture_mvvm_modify](https://github.com/antonpichka/library_architecture_mvvm_modify). The reason for creating the two-layer architectural design pattern is the high cognitive load associated with the three-layer architectural design pattern. The explanation is provided in the article [Comparison of architectural design patterns](./comparison_of_architectural_design_patterns.pdf) available only in Russian, as my English proficiency is insufficient.
![SMV](./smv.png)
- Share - exchanges temporary data between views (The data persists only for the lifetime of the application.).
- Model - this is data, and it can manipulate data.
- View - this is the presentation (Or user interface) that uses the Model to display data to the user.

### Note

- This forms the basis of the two-layer SMV architectural design pattern. You can extend it as you see fit. That is why I did not describe the implementation of my two-layer SMV architectural design pattern in detail.
