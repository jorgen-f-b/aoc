#include "../../base.h"
#include "../../read_file.c"

#define FILE_SIZE 20165

void prep_string_array(char *file, int file_size, int *str_len, int *str_count) {
    int w;
    for (w = 0; file[w] != '\n'; w++);
    for (int i = w; i < file_size; i += w + 1) file[i] = 0;

    *str_len = w + 1;
    *str_count = file_size / (w + 1);
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    int str_len;
    int str_count;
    prep_string_array(file, file_size, &str_len, &str_count);
    char *diagram[str_count];
    for (int i = 0; i < str_count; i++) diagram[i] = file + (i * str_len);

    int s;
    for (s = 0; diagram[0][s] != 'S'; s++);
    u64 m[str_len-1];
    for (int i = 0; i < str_len-1; i++) m[i] = 0;
    m[s] = 1;

    for (int i = 1; i < str_count; i++) {
        for (int j = 0; j < str_len-1; j++) {
            if (!m[j]) continue;
            if (diagram[i][j] == '^') {
                m[j-1] += m[j];
                m[j+1] += m[j];
                m[j] = 0;
            }
        }
    }

    u64 res = 0;
    for (int i = 0; i < str_len-1; i++) res += m[i];

    printf("res: %llu\n", res);

    return 0;
}
