import 'package:smv_dart/src/exception_adapter.dart';
import 'package:test/test.dart';

void main() {
  group(
      "ExceptionAdapter",
      () {
            test("hasException()", () {
              final exceptionAdapter = ExceptionAdapter(null);
              expect(exceptionAdapter.hasException(), equals(false));
            });
          });
}
