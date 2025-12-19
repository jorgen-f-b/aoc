#include <stdio.h>
#include <stdlib.h>
#include "../../read_file.c"
#include "../../parse.c"
#include "../../d_array.h"

#define FILE_SIZE 5773

typedef struct {
    int x, y;
} Tile;
#define tile_lit(X, Y) (Tile){ X, Y }
Dyn_Array_Make(Tiles, Tile);

#define max(x, y) (x >= y ? x : y)
#define min(x, y) (x < y ? x : y)

u64 tile_area(Tile t1, Tile t2) {
    int x_max = max(t1.x, t2.x);
    int x_min = min(t1.x, t2. x);
    int y_max = max(t1.y, t2.y);
    int y_min = min(t1.y, t2. y);

    return cast(u64, x_max-x_min+1) * cast(u64, y_max-y_min+1);
}

Tile parse_tile(char *line) {
    char num[11] = {0};
    int idx = 0;
    while (*line != ',') {
        num[idx++] = *line;
        line++;
    }
    line++;
    int x = parse_int(num);

    idx = 0;
    while (*line) {
        num[idx++] = *line;
        num[idx] = 0;
        line++;
    }
    int y = parse_int(num);

    return tile_lit(x, y);
}

typedef struct {
    int min, max;
} Min_Max;
#define min_max(n1, n2) (Min_Max){ min(n1, n2), max(n1, n2) }

typedef struct {
    Min_Max x, y;
} Rectangle;

Rectangle parse_rectangle(Tile t1, Tile t2) {
    return (Rectangle){ min_max(t1.x, t2.x), min_max(t1.y, t2.y) };
}

bool check_collision(Rectangle rec_m, Tiles *tiles) {
    for (int i = 0; i < tiles->size; i++) {
        Rectangle rec_t = parse_rectangle(tiles->items[i], tiles->items[(i+1)%tiles->size]);
        if (rec_t.y.max < rec_m.y.min+1 || rec_t.y.min > rec_m.y.max-1 || rec_t.x.max < rec_m.x.min+1 || rec_t.x.min > rec_m.x.max-1) {
            continue;
        }
        return true;
    }
    return false;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    Tiles tiles = {0};
    File_Line_Iter iter = file_line_iter_lit(file);
    while (iter_line(&iter)) {
        da_append(&tiles, parse_tile(iter.line));
    }

    u64 largest_area = 0;
    for (int i = 1; i < tiles.size; i++) {
        Tile t1 = tiles.items[i];
        for (int j = i + 1; j < tiles.size; j++) {
            Tile t2 = tiles.items[j];
            if (!check_collision(parse_rectangle(t1, t2), &tiles)) {
                u64 area = tile_area(t1, t2);
                if (area > largest_area) largest_area = area;
            }
        }
    }

    printf("Largest area: %llu\n", largest_area);

    return 0;
}
