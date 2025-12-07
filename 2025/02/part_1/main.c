#include "../../read_file.c"
#include "../../base.h"
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define FILE_SIZE 512
#define STR_NUM_LEN 20

typedef struct {
    u64 to, from;
} Range;
#define range_lit(T, F) (Range){ T, F }

typedef struct {
    char *buffer;
    int buffer_size;
    int idx;
    Range range;
} Range_Iter;
#define range_iter_lit(B, S) { B, S, 0 }

u64 str_to_u64(char *str) {
    return strtoull(str, NULL, 10);
}

bool iterate_range(Range_Iter *iter) {
    if (iter->idx >= iter->buffer_size) return false;

    int idx = 0;
    char to_str[STR_NUM_LEN] = {0};
    while (iter->buffer[iter->idx] != '-') {
        to_str[idx++] = iter->buffer[iter->idx++];
    }
    iter->idx++;

    idx = 0;
    char from_str[STR_NUM_LEN] = {0};
    while (iter->buffer[iter->idx] != ',' && iter->idx < iter->buffer_size) {
        from_str[idx++] = iter->buffer[iter->idx++];
    }
    iter->idx++;

    iter->range = range_lit(str_to_u64(to_str), str_to_u64(from_str));

    return true;
}

bool is_invalid(u64 id) {
    char str_id[STR_NUM_LEN] = {0};
    snprintf(str_id, STR_NUM_LEN, "%zu", id);
    int id_size = strlen(str_id);

    if (id_size % 2 == 1) return false;
    for (int s = 0, m = id_size/2; m < id_size; s++, m++) {
        if (str_id[s] != str_id[m]) return false;
    }
    return true;
}

u64 invalid_ids(Range *range) {
    u64 res = 0;
    for (u64 id = range->to; id <= range->from; id++) {
        if (is_invalid(id)) res += id;
    }
    return res;
}

int main() {
    char input[FILE_SIZE];
    int input_size = read_entire_file("../input.txt", input, FILE_SIZE);
    if (!input_size) return 1;

    u64 res = 0;
    Range_Iter iter = range_iter_lit(input, input_size);
    while (iterate_range(&iter)) {
        res += invalid_ids(&iter.range);
    }
    printf("Invalid ids: %zu\n", res);

    return 0;
}
