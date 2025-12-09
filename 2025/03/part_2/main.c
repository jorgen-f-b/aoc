#include "../../base.h"
#include "../../read_file.c"
#include <stdlib.h>
#include <string.h>

#define FILE_SIZE 20201

typedef struct {
    char *input;
    int input_size;
    int num_batteries;
    int idx;
    u64 joltage;
} Iterator;
#define iterator_lit(I, S) { I, S, strlen(I), 0 };

u64 str_to_u64(char *str) {
    return strtoull(str, NULL, 10);
}

bool iterate_battery_packs(Iterator *iter) {
    if (iter->idx >= iter->input_size) return false;

    char joltage[12] = {0};
    int last_idx = -1;
    for (int i = 0; i < 12; i++) {
        for (int j = last_idx+1; j < iter->num_batteries - (11-i); j++) {
            if (iter->input[iter->idx+j] > joltage[i]) {
                last_idx = j;
                joltage[i] = iter->input[iter->idx+j];
            }
        }
    }

    iter->idx += iter->num_batteries+1;
    iter->joltage = str_to_u64(joltage);

    return true;
}

void replace_new_line(char *str, int len) {
    for (int i = 0; i < len; i++) {
        if (str[i] == '\n') str[i] = 0;
    }
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);
    replace_new_line(file, file_size);

    u64 sum = 0;
    Iterator iter = iterator_lit(file, file_size);
    while (iterate_battery_packs(&iter)) sum += iter.joltage;

    printf("Sum: %"PRIu64"\n", sum);
    return 0;
}
