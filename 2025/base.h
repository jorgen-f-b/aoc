#ifndef BASE_H
#define BASE_H

#include <inttypes.h>
#include <stdbool.h>

#define cast(T,V) (T)(V)

#define i8 int8_t
#define u8 uint8_t

#define I8_MIN INT8_MIN
#define I8_MAX INT8_MAX
#define U8_MIN 0
#define U8_MAX UINT8_MAX

#define i16 int16_t
#define u16 uint16_t

#define I16_MIN INT16_MIN
#define I16_MAX INT16_MAX
#define U16_MIN 0
#define U16_MAX UINT16_MAX

#define i32 int32_t
#define u32 uint32_t

#define I32_MIN INT32_MIN
#define I32_MAX INT32_MAX
#define U32_MIN 0
#define U32_MAX UINT32_MAX

#define i64 int64_t
#define u64 uint64_t

#define I64_MIN INT64_MIN
#define I64_MAX INT64_MAX
#define U64_MIN 0
#define U64_MAX UINT64_MAX

#endif // BASE_H
