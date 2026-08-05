export declare abstract class BaseArrayModel<T extends BaseModel> {
    readonly arrayModel: Array<T>;
    protected constructor(arrayModel: Array<T>);
    abstract clone(): BaseArrayModel<T>;
    abstract toArrayMap(): Array<Record<string, any>>;
    abstract toString(): string;
    add(newModel: T): void;
    updateById(newModel: T): void;
    deleteById(id: string): void;
    addFromArray(newArrayModel: Array<T>): void;
    updateFromArrayById(newArrayModel: Array<T>): void;
    deleteFromArrayById(arrayId: Array<string>): void;
}
export declare abstract class BaseModelRepository<Y extends BaseModel, U extends BaseArrayModel<Y>> implements IDispose {
    protected constructor();
    abstract dispose(): void;
    protected abstract fromMap(map: Record<string, any>): Y;
    protected abstract fromArrayMap(arrayMap: Array<Record<string, any>>): U;
    protected getSafeValue<T>(map: Map<string, any> | Record<string, any>, key: string, defaultValue: T): T;
}
export declare abstract class BaseModel {
    readonly id: string;
    protected constructor(id: string);
    abstract clone(): BaseModel;
    abstract toMap(): Record<string, any>;
    abstract toString(): string;
}
export declare class IterationService {
    static readonly instance: IterationService;
    private readonly arrayUuid;
    private constructor();
    next(): string;
}
export declare class ShareProxy {
    private readonly listenerId;
    private readonly shareService;
    constructor();
    getValue<T>(key: string, defaultValue: T): T;
    update(key: string, value: any): void;
    delete(key: string): void;
    addListener(key: string, callback: (event: any) => void): void;
    notifyListener(key: string, value: any): void;
    notifyListeners(key: string, value: any): void;
    deleteAllListenersByKey(key: string): void;
    deleteAllListenersByArrayKey(arrayKey: Array<string>): void;
    deleteListenerByListenerId(key: string): void;
    deleteListenersByListenerId(arrayKey: Array<string>): void;
}
export declare class ShareService {
    static readonly instance: ShareService;
    private readonly cache;
    private readonly listenersByKey;
    private constructor();
    getValue<T>(key: string, defaultValue: T): T;
    update(key: string, value: any): void;
    delete(key: string): void;
    addListener(key: string, listenerId: string, callback: (event: any) => void): void;
    notifyListener(key: string, listenerId: string, value: any): void;
    notifyListeners(key: string, value: any): void;
    deleteAllListenersByKey(key: string): void;
    deleteAllListenersByArrayKey(arrayKey: Array<string>): void;
    deleteListenerByListenerId(key: string, listenerId: string): void;
    deleteListenersByListenerId(arrayKey: Array<string>, listenerId: string): void;
}
export declare abstract class BaseException extends Error {
    private readonly source;
    private readonly type;
    protected constructor(source: string, type: string);
    abstract toString(): string;
    protected initException(): void;
    private debugPrintException;
}
export declare enum EnumGuilty {
    developer = "developer",
    device = "device",
    user = "user"
}
export declare class ExceptionAdapter {
    readonly exception: BaseException | null;
    constructor(exception: BaseException | null);
    hasException(): boolean;
}
export interface IDispose {
    dispose(): void;
}
export declare class LocalException extends BaseException {
    readonly guilty: EnumGuilty;
    readonly extraMessage: string;
    constructor(source: string, guilty: EnumGuilty, extraMessage: string);
    toString(): string;
}
export declare class NetworkException extends BaseException {
    readonly statusCode: number;
    readonly nameStatusCode: string;
    readonly descriptionStatusCode: string;
    constructor(source: string, statusCode: number, nameStatusCode: string, descriptionStatusCode: string);
    static fromSourceAndStatusCode(source: string, statusCode: number): NetworkException;
    toString(): string;
}
export declare class ResultArrayModel<T extends BaseModel, Y extends BaseArrayModel<T>> {
    readonly data: Y | null;
    readonly exceptionAdapter: ExceptionAdapter;
    private constructor();
    static success<T extends BaseModel, Y extends BaseArrayModel<T>>(data: Y): ResultArrayModel<T, Y>;
    static exception<T extends BaseModel, Y extends BaseArrayModel<T>>(exception: BaseException): ResultArrayModel<T, Y>;
}
export declare class ResultModel<T extends BaseModel> {
    readonly data: T | null;
    readonly exceptionAdapter: ExceptionAdapter;
    private constructor();
    static success<Y extends BaseModel>(data: Y): ResultModel<Y>;
    static exception<Y extends BaseModel>(exception: BaseException): ResultModel<Y>;
}
export declare class Result {
    readonly exceptionAdapter: ExceptionAdapter;
    private constructor();
    static success(): Result;
    static exception(exception: BaseException): Result;
}
export declare function debugPrint(text: string): void;
export declare function debugPrintException(text: string): void;
export declare function debugPrintMethod(methodName: string): void;
//# sourceMappingURL=index.d.ts.map