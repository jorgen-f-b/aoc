#include "../../read_file.c"
#include <stdlib.h>

#define FILE_SIZE 20201

typedef struct {
    char *input;
    int input_size;
    int idx;
    int joltage;
} Iterator;
#define iterator_lit(I, S) { I, S, 0 };

bool iterate_battery_packs(Iterator *iter) {
    if (iter->idx >= iter->input_size) return false;

    char largest = iter->input[iter->idx];
    char seccond_largest = iter->input[++iter->idx];
    while (iter->input[++iter->idx] != '\n') {
        if (iter->input[iter->idx + 1] != '\n' && iter->input[iter->idx] > largest) {
            largest = iter->input[iter->idx];
            seccond_largest = '0';
        } else if (iter->input[iter->idx] > seccond_largest) seccond_largest = iter->input[iter->idx];
    }
    iter->idx++;

    char joltage[3] = { largest, seccond_largest, 0 };
    iter->joltage = atoi(joltage);

    return true;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    int sum = 0;
    Iterator iter = iterator_lit(file, file_size);
    while (iterate_battery_packs(&iter)) sum += iter.joltage;

    printf("Sum: %d\n", sum);
    return 0;
}
