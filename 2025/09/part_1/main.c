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

typedef struct {
    char *file;
    Tile tile;
} File_Iter;
#define file_iter_lit(F) (File_Iter){ F }

bool iter_line(File_Iter *iter) {
    if (!*iter->file) return false;

    char num[11] = {0};
    int idx = 0;
    while (*iter->file != ',') {
        num[idx++] = *iter->file;
        iter->file++;
    }
    iter->file++;
    int x = parse_int(num);

    idx = 0;
    while (*iter->file != '\n') {
        num[idx++] = *iter->file;
        num[idx] = 0;
        iter->file++;
    }
    iter->file++;
    int y = parse_int(num);

    iter->tile = tile_lit(x, y);
    return true;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    Tiles tiles = {0};
    File_Iter iter = file_iter_lit(file);
    while (iter_line(&iter)) {
        da_append(&tiles, iter.tile);
    }

    u64 largest_area = tile_area(tiles.items[0], tiles.items[1]);
    for (int i = 1; i < tiles.size; i++) {
        Tile t1 = tiles.items[i];
        for (int j = i + 1; j < tiles.size; j++) {
            Tile t2 = tiles.items[j];
            u64 area = tile_area(t1, t2);
            if (area > largest_area) largest_area = area;
        }
    }

    printf("Largest area: %llu\n", largest_area);

    return 0;
}
