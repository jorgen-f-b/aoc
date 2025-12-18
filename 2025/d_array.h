#include "stdlib.h"

#define Dyn_Array_Make(N, T) \
typedef struct { \
    T *items; \
    size_t size; \
    size_t capacity; \
} N

#define da_append(A, I) \
do { \
    if ((A)->size >= (A)->capacity) { \
        if ((A)->capacity == 0) (A)->capacity = 256; \
        else (A)->capacity *= 2; \
        (A)->items = realloc((A)->items, (A)->capacity * sizeof(*(A)->items)); \
    } \
    (A)->items[(A)->size++] = I; \
} while(0)

#define da_pop(A) (A)->items[--(A)->size]

#define da_destroy(A) \
do { \
    free((A)->items); \
    (A)->items = NULL; \
    (A)->size = 0; \
    (A)->capacity = 0; \
} while(0)
