#include "base.h"

#define Parse_Make_S(T) \
T parse_##T(char *s) { \
    T num = 0; \
    bool is_negative = false; \
    while (*s == ' ') s += 1; \
    if (s[0] == '-') { \
        is_negative = true; \
        s += 1; \
    } \
    for (int i = 0; s[i]; i++) { \
        if (s[i] == ' ' || s[i] == ',' || s[i] == '_') continue; \
        num *= 10; \
        num += s[i] - '0'; \
    } \
    if (is_negative) num = -num; \
    return num; \
}

#define Parse_Make_U(T) \
T parse_##T(char *s) { \
    T num = 0; \
    while (*s == ' ') s += 1; \
    for (int i = 0; s[i]; i++) { \
        if (s[i] == ' ' || s[i] == ',' || s[i] == '_') continue; \
        num *= 10; \
        num += s[i] - '0'; \
    } \
    return num; \
}

Parse_Make_S(int)
Parse_Make_S(i8)
Parse_Make_U(u8)
Parse_Make_S(i16)
Parse_Make_U(u16)
Parse_Make_S(i32)
Parse_Make_U(u32)
Parse_Make_S(i64)
Parse_Make_U(u64)
