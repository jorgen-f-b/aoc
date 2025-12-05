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

typedef struct {
    char *instructions;
    int index;
    int value;
    int points;
} Instruction_Iter;
#define instruction_iter_lit(S) { S, 0, 50, 0 }

void turn_dial(Instruction_Iter *itr, int direction) {
    bool was_0 = itr->value == 0;
    itr->value += direction;
    while (itr->value < 0) {
        if (!was_0) {
            itr->points++;
        } else was_0 = false;
        itr->value += 100;
    }
    while (itr->value > 100) {
        itr->value -= 100;
        itr-> points++;
    }
    if (itr->value == 0 || itr->value == 100) {
        itr->value = 0;
        itr->points++;
    }
}

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
    turn_dial(itr, value);

    return true;
}

int main() {
    //char *text = read_entire_file("../example.txt");
    char *text = read_entire_file("../input.txt");
    Instruction_Iter itr = instruction_iter_lit(text);

    while (instructions_iter(&itr));

    printf("points: %d\n", itr.points);

    free(text);
    return 0;
}
