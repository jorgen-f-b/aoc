#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

char *read_entire_file(const char *filename) {
    FILE *f = fopen(filename, "rb");
    fseek(f, 0, SEEK_END);
    long fsize = ftell(f);
    fseek(f, 0, SEEK_SET);

    char *str = malloc(fsize + 1);
    fread(str, fsize, 1, f);
    fclose(f);
    str[fsize] = 0;

    return str;
}

int turn_dial(int from, int direction) {
  int res = from + direction;
  while (res < 0) res += 100;
  return res % 100;
}

typedef struct {
    char *instructions;
    int index;
    int value;
} Instruction_Iter;

bool instructions_iter(Instruction_Iter *itr) {
    char *instructions = itr->instructions;
    int itr_index = itr->index;
    if (instructions[itr_index] == 0) return false;

    bool negative_value;
    int idx = 0;
    char c = instructions[itr_index];
    char num[11];

    while (c != '\n') {
        if (c == 'L') negative_value = true;
        else if (c == 'R') negative_value = false;
        else num[idx++] = c;

        c = instructions[++itr_index];
    }
    num[idx] = 0;
    itr_index++;

    int value = atoi(num);
    if (negative_value) value = -value;

    itr->index = itr_index;
    itr->value = turn_dial(itr->value, value);

    return true;
}

int main() {
    //char *text = read_entire_file("example.txt");
    char *text = read_entire_file("input.txt");
    Instruction_Iter itr = { text, 0, 50 };
    int points = 0;

    while (instructions_iter(&itr)) {
        if (itr.value == 0) points++;
    }

    printf("points: %d\n", points);

    free(text);
    return 0;
}
