#include "../../base.h"
#include "../../read_file.c"
#include <stdlib.h>

#define FILE_SIZE 18686

int get_row_length(char *file) {
    int i = 0;
    while (file[i] != '\n') i++;
    return i;
}

u64 add(u64 x, u64 y) {
    return x + y;
}

u64 multiply(u64 x, u64 y) {
    return x * y;
}

int get_column_length(char *file, int row_length) {
    int i;
    for (i = 0; file[i] != '+' && file[i] != '*'; i += row_length + 1);
    return (i / row_length) + 1;
}

u64 solve_problem(u64 *values, int values_size, u64 (*problem)(u64, u64)) {
    u64 res = values[0];
    for (int i = 1; i < values_size; i++) {
        res = problem(res, values[i]);
    }
    return res;
}

#define MAX_INT_SIZE 16
bool iterate_rows(char *file, int row_length, const int column_length, int *idx, u64 *value) {
    if (*idx < 0) return false;

    u64 arr[column_length];
    int arr_idx = 0;
    char num[MAX_INT_SIZE] = {0};
    int num_idx = 0;
    while (*idx >= 0) {
        for (int i = 0; i < column_length; i++) {
            int index = ((row_length + 1) * i) + (*idx);
            if (file[index] == ' ') continue;
            if (file[index] == '+') {
                arr[arr_idx++] = atoi(num);
                *value = solve_problem(arr, arr_idx, add);
                *idx -= 2;
                return true;
            }
            if (file[index] == '*') {
                arr[arr_idx++] = atoi(num);
                *value = solve_problem(arr, arr_idx, multiply);
                *idx -= 2;
                return true;
            }
            num[num_idx++] = file[index];
            num[num_idx] = 0;
        }
        arr[arr_idx++] = atoi(num);
        num[0] = 0;
        num_idx = 0;

        *idx -= 1;
    }

    return false;
}

u64 sum(u64 *values, int values_size) {
    u64 res = 0;
    for (int i = 0; i < values_size; i++) res += values[i];
    return res;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);
    int row_length = get_row_length(file);
    int column_length = get_column_length(file, row_length);

    int idx = row_length - 1;
    u64 value;
    u64 sum = 0;
    while (iterate_rows(file, row_length, column_length, &idx, &value)) sum += value;

    printf("Sum: %llu\n", sum);

    return 0;
}
