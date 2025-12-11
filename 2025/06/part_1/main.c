#include "../../base.h"
#include "../../read_file.c"
#include <stdlib.h>

#define FILE_SIZE 18686

int get_row_length(char *file) {
    int i = 0;
    while (file[i] != '\n') i++;
    return i;
}

int get_num_elements(char *file, int row_length) {
    int num = 0;
    bool is_element = false;
    for (int i = 0; i <= row_length; i++) {
        if (file[i] != ' ' && file[i] != '\n') is_element = true;
        else {
            if (is_element) {
                is_element = false;
                num++;
            }
        }
    }
    return num;
}

u64 add(u64 x, u64 y) {
    return x + y;
}

u64 multiply(u64 x, u64 y) {
    return x * y;
}

void init_problems(char *file, int row_length, u64 (**problems)(u64,u64)) {
    int i;
    for (i = 0; file[i] != '+' && file[i] != '*'; i += row_length + 1);
    int idx = 0;
    for (int j = 0; j < row_length; j++) {
        if (file[i+j] == '+') problems[idx++] = add;
        else if (file[i+j] == '*') problems[idx++] = multiply;
    }
}

#define MAX_INT_SIZE 16
bool iterate_rows(char *file, int row_length, int *idx, u64 *arr) {
    if (file[*idx] == '+' || file[*idx] == '*') return false;

    int arr_idx = 0;
    char num[MAX_INT_SIZE] = {0};
    int num_idx = 0;
    for (int i = 0; i <= row_length; i++) {
        if (file[*idx + i] == ' ' || file[*idx + i] == '\n') {
            if (num[0] != 0) {
                num_idx = 0;
                arr[arr_idx++] = atoi(num);
                num[0] = 0;
            }
            continue;
        }
        num[num_idx++] = file[*idx + i];
        num[num_idx] = 0;
    }
    *idx += row_length + 1;

    return true;
}

u64 sum(u64 *values, int values_size) {
    u64 res = 0;
    for (int i = 0; i < values_size; i++) res += values[i];
    return res;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    // int file_size = read_entire_file("../example.txt", file, FILE_SIZE);
    int length = get_row_length(file);
    const int num_elements = get_num_elements(file, length);

    u64 (*problems[num_elements])(u64, u64);
    u64 results[num_elements];
    for (int i = 0; i < num_elements; i ++) {
        problems[i] = NULL;
        results[i] = 0;
    }

    init_problems(file, length, problems);
    for (int i = 0; i < num_elements; i++) {
        problems[i](1, 2);
    }

    int idx = 0;
    u64 rows[num_elements];
    while (iterate_rows(file, length, &idx, rows)) {
        for (int i = 0; i < num_elements; i++) {
            if (results[i] == 0) results[i] = rows[i];
            else results[i] = problems[i](results[i], rows[i]);
        }
    };

    printf("Sum: %llu\n", sum(results, num_elements));

    return 0;
}
