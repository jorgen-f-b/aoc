#include "../../read_file.c"
#include <stdbool.h>
#include <stdio.h>

#define FILE_SIZE 18633

void change_new_line(char *file, int size) {
    for (int i = 0; i < size; i++) {
        if (file[i] == '\n') file[i] = 0;
    }
}

int get_height_and_width(char *file) {
    int i;
    for (i = 0; file[i]; i++);
    return i;
}

void init_grid(char **grid, char *file, int file_size) {
    grid[0] = file;
    int idx = 1;
    for (int i = 0; i < file_size; i++) {
        if (file[i] == 0) grid[idx++] = file+i+1;
    }
}

bool is_roll(char **grid, int y, int x) {
    return grid[y][x] == '@';
}

int roll_amount(char *line, int x, int width) {
    int amount = 0;
    if (x > 0 && line[x-1] == '@') amount++;
    if (line[x] == '@') amount++;
    if (x < width-1 && line[x+1] == '@') amount++;
    return amount;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);
    change_new_line(file, file_size);

    const int height_and_width = get_height_and_width(file);
    char *grid[height_and_width];
    init_grid(grid, file, file_size);

    int res = 0;
    for (int y = 0; y < height_and_width; y++) {
        for (int x = 0; x < height_and_width; x++) {
            if (is_roll(grid, y, x)) {
                int amount = 0;
                if (y > 0) amount += roll_amount(grid[y-1], x, height_and_width);
                amount += roll_amount(grid[y], x, height_and_width) - 1;
                if (y < height_and_width-1)  amount += roll_amount(grid[y+1], x, height_and_width);

                if (amount < 4) res++;
            }
        }
    }

    printf("res: %d\n", res);

    return 0;
}
