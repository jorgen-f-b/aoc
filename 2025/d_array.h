#include "stdlib.h"

#define Dyn_Array_Make(N, T) \
typedef struct { \
    T *items; \
    size_t size; \
    size_t capacity; \
} N;

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

#define QSwap_Make(N, TA, TI) \
void N##_swap(TA *arr, int i, int j) { \
    TI tmp = arr->items[i]; \
    arr->items[i] = arr->items[j]; \
    arr->items[j] = tmp; \
}

#define QPartition_Make(N, TA, GV) \
int N##_partition(TA *arr, int low, int high) { \
    int pivot = GV(arr, high); \
    int i = low - 1; \
    for (int j = low; j <= high - 1; j++) { \
        if (GV(arr, j) < pivot) { \
            N##_swap(arr, ++i, j); \
        } \
    } \
    N##_swap(arr, i + 1, high); \
    return i + 1; \
}

#define Q_Make(N, TA) \
void N##_quick_sort(TA *arr, int low, int high) { \
    if (low >= high) return; \
    int pi = N##_partition(arr, low, high); \
    N##_quick_sort(arr, low, pi - 1); \
    N##_quick_sort(arr, pi + 1, high); \
}

#define QuickSort_Make(N, TA, TI, GV) \
QSwap_Make(N, TA, TI) \
QPartition_Make(N, TA, GV) \
Q_Make(N, TA)
