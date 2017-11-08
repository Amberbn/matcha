// +build matcha,darwin

#include <stdio.h>
#include <stdint.h>
#include <string.h>
#import <Foundation/Foundation.h>
#include "objc-go.h"
#include "objc-foreign.h"

@implementation MatchaGoValue {
    GoRef _ref;
}

@synthesize ref = _ref;

- (id)initWithGoRef:(GoRef)ref {
    if ((self = [super init])) {
        _ref = ref;
    }
    return self;
}

- (id)initWithObject:(id)v {
    return [self initWithGoRef:matchaGoForeign(MatchaForeignTrack(v))];
}

- (id)initWithBool:(BOOL)v {
    return [self initWithGoRef:matchaGoBool(v)];
}

- (id)initWithInt:(int)v {
    return [self initWithGoRef:matchaGoInt(v)];
}

- (id)initWithLongLong:(long long)v {
    return [self initWithGoRef:matchaGoInt64(v)];
}

- (id)initWithUnsignedLongLong:(unsigned long long)v {
    return [self initWithGoRef:matchaGoUint64(v)];
}

- (id)initWithDouble:(double)v {
    return [self initWithGoRef:matchaGoFloat64(v)];
}

- (id)initWithString:(NSString *)v {
    CGoBuffer buf = MatchaNSStringToCGoBuffer(v);
    return [self initWithGoRef:matchaGoString(buf)];
}

- (id)initWithData:(NSData *)v {
    CGoBuffer buf = MatchaNSDataToCGoBuffer(v);
    return [self initWithGoRef:matchaGoBytes(buf)];
}

- (id)initWithArray:(NSArray<MatchaGoValue *> *)v {
    GoRef ref = matchaGoArray(MatchaNSArrayToCGoBuffer(v));
    return [self initWithGoRef:ref];
}

- (id)initWithFunc:(NSString *)v {
    CGoBuffer buf = MatchaNSStringToCGoBuffer(v);
    return [self initWithGoRef:matchaGoFunc(buf)];
}

- (id)toObject {
    return MatchaForeignGet(matchaGoToForeign(_ref));
}

- (BOOL)toBool {
    return matchaGoToBool(_ref);
}

- (long long)toLongLong {
    return matchaGoToInt64(_ref);
}

- (unsigned long long)toUnsignedLongLong {
    return matchaGoToUint64(_ref);
}

- (double)toDouble {
    return matchaGoToFloat64(_ref);
}

- (NSString *)toString {
    return MatchaCGoBufferToNSString(matchaGoToString(_ref));
}

- (NSData *)toData {
    return MatchaCGoBufferToNSData(matchaGoToBytes(_ref));
}

- (NSArray *)toArray {
    return MatchaCGoBufferToNSArray(matchaGoToArray(_ref));
}

- (BOOL)isNil {
    return matchaGoIsNil(_ref);
}

- (NSArray<MatchaGoValue *> *)call:(NSString *)method, ... {
    NSMutableArray *array = [[NSMutableArray alloc] init];
    va_list args;
    va_start(args, method);
    NSArray *rlt = [self call:method args:args];
    va_end(args);
    
    return rlt;
}

- (NSArray<MatchaGoValue *> *)call:(NSString *)method args:(va_list)args {
    NSMutableArray *array = [NSMutableArray array];
    id arg = nil;
    while ((arg = va_arg(args, id))) {
        [array addObject:arg];
    }

    CGoBuffer argsBuffer = MatchaNSArrayToCGoBuffer(array);
    CGoBuffer rlt = matchaGoCall(_ref, MatchaNSStringToCGoBuffer(method), argsBuffer);
    return MatchaCGoBufferToNSArray(rlt);
}

- (void)dealloc {
    matchaGoUntrack(_ref);
}

@end