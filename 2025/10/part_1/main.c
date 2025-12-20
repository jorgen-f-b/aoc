#include <stdio.h>
#include "../../read_file.c"
#include "../../parse.c"
#include "../../d_array.h"

#define FILE_SIZE 19900

Dyn_Array_Make(Button, int);
Dyn_Array_Make(Buttons, Button);

#define LIGHT_DIAGRAM_MAX 1023
int fewest_total_presses(int light_diagram, Buttons *buttons) {
    bool visited[LIGHT_DIAGRAM_MAX] = {0};
    Button queue = {0};
    visited[0] = true;
    da_append(&queue, 0);
    int new_level_idx = 0;

    int idx = 0;
    int amount_of_pushes = 0;
    for (int idx = 0; idx < queue.size; idx++) {
        int curr = queue.items[idx];
        if (idx >= new_level_idx) {
            amount_of_pushes++;
            new_level_idx = queue.size;
        }

        for (int i = 0; i < buttons->size; i++) {
            int pushes = 0;
            for (int j = 0; j < buttons->items[i].size; j++) pushes |= 1 << buttons->items[i].items[j];

            int next_diagram = curr^pushes;
            if (!visited[next_diagram]) {
                if (next_diagram == light_diagram) {
                    da_destroy(&queue);
                    return amount_of_pushes;
                }
                visited[next_diagram] = true;
                da_append(&queue, next_diagram);
            }
        }
    }

    da_destroy(&queue);
    return -1;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);

    u64 res = 0;
    File_Line_Iter iter = file_line_iter_lit(file);
    while (iter_line(&iter)) {
        int light_diagram = 0;
        int i = 0;
        while (iter.line[++i] != ']') {
            if (iter.line[i] == '#') light_diagram |= 1 << (i-1);
        }

        i += 1;
        Buttons buttons = {0};
        while (iter.line[++i] != '{') {
            if (iter.line[i] == ' ') continue;
            if (iter.line[i] == '(') {
                Button button = {0};
                while (iter.line[++i] != ')') {
                    if (iter.line[i] != ',') {
                        char n[2] = { iter.line[i], 0 };
                        da_append(&button, parse_int(n));
                    }
                }
                da_append(&buttons, button);
            }
        }
        res += fewest_total_presses(light_diagram, &buttons);

        for (int i = 0; i < buttons.size; i++) da_destroy(&buttons.items[i]);
        da_destroy(&buttons);
    }

    printf("Res: %llu\n", res);

    return 0;
}
