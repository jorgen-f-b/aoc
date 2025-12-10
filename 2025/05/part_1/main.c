#include <stdlib.h>
#include "../../base.h"
#include "../../read_file.c"

#define FILE_SIZE 21373
#define STR_NUM_LEN 20

typedef struct {
    u64 from, to;
} Range;
#define range_lit(F, T) ((Range){ F, T })

typedef struct {
    char *input;
    int input_size;
    int idx;
    bool get_range;
    union{
        Range range;
        u64 id;
   };
} Iter_Input;
#define iter_input_lit(I, S) { I, S, 0, true }

u64 str_to_u64(char *str) {
    return strtoull(str, NULL, 10);
}

bool get_ranges(Iter_Input *iter) {
    if (iter->input[iter->idx] == '\n') {
        iter->idx++;
        return false;
    }

    char from[STR_NUM_LEN] = {0};
    char to[STR_NUM_LEN] = {0};
    int idx = 0;
    bool is_from = true;

    while (iter->input[iter->idx] != '\n') {
        if (iter->input[iter->idx] == '-') {
            is_from = false;
            idx = 0;
            iter->idx++;
        }
        if (is_from) from[idx] = iter->input[iter->idx];
        else to[idx] = iter->input[iter->idx];
        idx++;
        iter->idx++;
    }
    iter->idx++;
    iter->range = range_lit(str_to_u64(from), str_to_u64(to));

    return true;
}

bool get_ids(Iter_Input *iter) {
    if (iter->idx >= iter->input_size) return false;

    char id[STR_NUM_LEN] = {0};
    int idx = 0;
    while (iter->input[iter->idx] != '\n') id[idx++] = iter->input[iter->idx++];
    iter->idx++;
    iter->id = str_to_u64(id);

    return true;
}

bool is_fresh(u64 id, Range *ranges, int range_size) {
    for (int i = 0; i < range_size; i++) {
        if (id >= ranges[i].from && id <= ranges[i].to) return true;
    }
    return false;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    Range ranges[200];
    int range_size = 0;
    Iter_Input iter = iter_input_lit(file, file_size);
    while (get_ranges(&iter)) ranges[range_size++] = iter.range;

    int res = 0;
    while (get_ids(&iter)) {
        if (is_fresh(iter.id, ranges, range_size)) res++;
    }

    printf("res: %d\n", res);

    return 0;
}
