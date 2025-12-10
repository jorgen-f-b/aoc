#include <stdlib.h>
#include "../../base.h"
#include "../../read_file.c"

#define FILE_SIZE 512
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
        int id;
   };
} Iter_Input;
#define iter_input_lit(I, S) { I, S, 0, true }

u64 str_to_u64(char *str) {
    return strtoull(str, NULL, 10);
}

bool get_ranges(Iter_Input *iter) {
    if (iter->input[iter->idx] == '\n') return false;

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

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    Iter_Input iter = iter_input_lit(file, file_size);
    while (get_ranges(&iter)) printf("Range{ %lld, %lld }\n", iter.range.from, iter.range.to);

    return 0;
}
