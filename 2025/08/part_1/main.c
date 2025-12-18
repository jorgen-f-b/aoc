#include <math.h>
#include "../../read_file.c"
#include "../../parse.c"
#include "../../d_array.h"

#define FILE_SIZE 512

typedef struct {
    int x, y, z;
} Vec3;
#define vec3_lit(X, Y, Z) (Vec3){ X, Y, Z }

Dyn_Array_Make(Da_Vec3, Vec3);

typedef struct {
    char *file;
    Vec3 vec3;
} Iter;
#define iter_lit(F) (Iter){ F }

#define pow2(x) pow(x, 2) 

int distance(Vec3 v1, Vec3 v2) {
    return sqrt(pow2(v1.x - v2.x) + pow2(v1.y - v2.y) + pow2(v1.z - v2.z));
}

bool iterate_file(Iter *iter) {
    if (!*iter->file) return false;

    int i = 0;
    char num[12] = {0};

    while (*iter->file != ',') {
        num[i++] = *iter->file;
        iter->file++;
    }
    iter->file++;
    int x = parse_int(num);

    i = 0;
    while (*iter->file != ',') {
        num[i++] = *iter->file;
        num [i] = 0;
        iter->file++;
    }
    iter->file++;
    int y = parse_int(num);

    i = 0;
    while (*iter->file != '\n') {
        num[i++] = *iter->file;
        num [i] = 0;
        iter->file++;
    }
    iter->file++;
    int z = parse_int(num);

    iter->vec3 = vec3_lit(x, y, z);

    return true;
}

typedef struct {
    int jb1, jb2;
    int dist;
} Jb_Dist;
#define jb_dist_lit(j1, j2, d) (Jb_Dist){ j1, j2, d }

Dyn_Array_Make(Da_Jb_Dist, Jb_Dist);

void swap(Da_Jb_Dist *arr, int i, int j) {
    Jb_Dist tmp = arr->items[i];
    arr->items[i] = arr->items[j];
    arr->items[j] = tmp;
}

int partition(Da_Jb_Dist *arr, int low, int high) {
    int pivot = arr->items[high].dist;
    int i = low - 1;

    for (int j = low; j <= high - 1; j++) {
        if (arr->items[j].dist < pivot) {
            i++;
            swap(arr, i, j);
        }
    }

    swap(arr, i + 1, high);
    return i + 1;
}

void quick_sort(Da_Jb_Dist *arr, int low, int high) {
    if (low >= high) return;

    int pi = partition(arr, low, high);

    quick_sort(arr, low, pi - 1);
    quick_sort(arr, pi + 1, high);
}

Dyn_Array_Make(Da_Id, int);

bool da_id_contains(Da_Id *arr, int id) {
    for (int i = 0; i < arr->size; i++) {
        if (arr->items[i] == id) return true;
    }
    return false;
}

bool da_id_cmp_merge(Da_Id *arr1, Da_Id *arr2) {
    bool merge = false;
    for (int i = 0; i < arr2->size; i++) {
        if (da_id_contains(arr1, arr2->items[i])) {
            merge = true;
            break;
        }
    }
    if (!merge) return false;

    for (int i = 0; i < arr2->size; i++) {
        if (!da_id_contains(arr1, arr2->items[i])) da_append(arr1, arr2->items[i]);
    }
    return true;
}

Dyn_Array_Make(Da_Da_Id, Da_Id);

void Da_Da_Id_Destroy(Da_Da_Id *arr) {
    for (int i = 0; i < arr->size; i++) da_destroy(&arr->items[i]);
    da_destroy(arr);
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    Iter iter = iter_lit(file);
    Da_Vec3 vectors = {0};
    while (iterate_file(&iter)) {
        da_append(&vectors, iter.vec3);
    }

    for (int i = 0; i < vectors.size; i++) {
        printf("%d: Vec3{ %d, %d, %d }\n", i, vectors.items[i].x, vectors.items[i].y, vectors.items[i].z);
    }

    Da_Jb_Dist jb_dists = {0};
    for (int i = 0; i < vectors.size; i++) {
        for (int j = i + 1; j < vectors.size; j++) {
            da_append(&jb_dists, jb_dist_lit(i, j, distance(vectors.items[i], vectors.items[j])));
        }
    }
    quick_sort(&jb_dists, 0, jb_dists.size-1);

    Da_Da_Id da_da_id = {0};
    for (int i = 0; i < jb_dists.size; i++) {
        int id1 = jb_dists.items[i].jb1;
        int id2 = jb_dists.items[i].jb2;
        bool found = false;

        for (int j = 0; j < da_da_id.size; j++) {
            Da_Id *da_id = &da_da_id.items[j];
            if (da_id_contains(da_id, id1)) {
                if (!da_id_contains(da_id, id2)) {
                    da_append(da_id, id2);
                }
                found = true;
                break;
            }
            if (da_id_contains(da_id, id2)) {
                da_append(da_id, id1);
                found = true;
                break;
            }
        }
        if (found) continue;

        Da_Id da_id = {0};
        da_append(&da_id, id1);
        da_append(&da_id, id2);

        da_append(&da_da_id, da_id);
    }

    for (int i = 0; i < da_da_id.size; i++) {
        Da_Id *da_id1 = &da_da_id.items[i];
        for (int j = i + 1; j < da_da_id.size; j++) {
            Da_Id *da_id2 = &da_da_id.items[j];
            if (da_id_cmp_merge(da_id1, da_id2)) {
                da_destroy(da_id2);
                da_da_id.items[j] = da_pop(&da_da_id);
            }
        }
    }

    for (int i = 0; i < da_da_id.size; i++) {
        printf("[%d", da_da_id.items[i].items[0]);
        for (int j = 1; j < da_da_id.items[i].size; j++) {
            printf(", %d", da_da_id.items[i].items[j]);
        }
        printf("]\n");
    }

    da_destroy(&vectors);
    da_destroy(&jb_dists);
    Da_Da_Id_Destroy(&da_da_id);
    return 0;
}
