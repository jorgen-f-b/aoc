#include <math.h>
#include "../../read_file.c"
#include "../../parse.c"
#include "../../d_array.h"

#define FILE_SIZE 17646

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

#define pow2(x) pow(cast(double, x), 2) 

int distance(Vec3 v1, Vec3 v2) {
    return sqrt(pow2(v1.x - v2.x) + pow2(v1.y - v2.y) + pow2(v1.z - v2.z));
}

bool iterate_file(Iter *iter) {
    if (!*iter->file) return false;

    int i = 0;
    char num[22] = {0};

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
#define jb_dist_dist(A, I) A->items[I].dist

Dyn_Array_Make(Jb_Dists, Jb_Dist);
QuickSort_Make(jb_dists, Jb_Dists, Jb_Dist, jb_dist_dist)

Dyn_Array_Make(Circuit, int);

bool circuit_contains(Circuit *circuit, int id) {
    for (int i = 0; i < circuit->size; i++) {
        if (circuit->items[i] == id) return true;
    }
    return false;
}

void circuit_print(Circuit *c) {
    printf("[%d", c->items[0]);
    for (int i = 1; i < c->size; i++) {
        printf(", %d", c->items[i]);
    }
    printf("]\n");
}

Dyn_Array_Make(Circuits, Circuit);
#define circuits_size(A, I) A->items[I].size
QuickSort_Make(circuits, Circuits, Circuit, circuits_size)

void circuits_print(Circuits *c) {
    for (int i = 0; i < c->size; i++) circuit_print(&c->items[i]);
}

void Circuits_Connect(Circuits *circuits, int id1, int id2) {
    size_t c1 = circuits->size;
    size_t c2 = circuits->size;

    for (size_t i = 0; i < circuits->size; i++) {
        Circuit *circuit = &circuits->items[i];
        if (c1 == circuits->size && circuit_contains(circuit, id1)) c1 = i;
        if (c2 == circuits->size && circuit_contains(circuit, id2)) c2 = i;
        if (c1 != circuits->size && c2 != circuits->size) break;
    }

    if (c1 == c2) return;

    Circuit *circuit1 = &circuits->items[c1];
    Circuit *circuit2 = &circuits->items[c2];
    for (int i = 0; i < circuit2->size; i++) {
        if (!circuit_contains(circuit1, circuit2->items[i])) da_append(circuit1, circuit2->items[i]);
    }
    da_destroy(circuit2);
    circuits->items[c2] = da_pop(circuits);
}

void Circuits_Destroy(Circuits *arr) {
    for (int i = 0; i < arr->size; i++) da_destroy(&arr->items[i]);
    da_destroy(arr);
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    Iter iter = iter_lit(file);
    Da_Vec3 vectors = {0};
    while (iterate_file(&iter)) {
        da_append(&vectors, iter.vec3);
    }

    Jb_Dists jb_dists = {0};
    for (int i = 0; i < vectors.size; i++) {
        for (int j = i + 1; j < vectors.size; j++) {
            da_append(&jb_dists, jb_dist_lit(i, j, distance(vectors.items[i], vectors.items[j])));
        }
    }
    jb_dists_quick_sort(&jb_dists, 0, jb_dists.size-1);

    Circuits circuits = {0};
    for (int i = 0; i < vectors.size; i++) {
        Circuit c = {0};
        da_append(&c, i);
        da_append(&circuits, c);
    }

    for (int i = 0; i < 1000; i++) {
        Circuits_Connect(&circuits, jb_dists.items[i].jb1, jb_dists.items[i].jb2);
    }
    circuits_quick_sort(&circuits, 0, circuits.size-1);

    u64 res = 1;
    for (int i = 0; i < 3; i++) {
        res *= circuits.items[circuits.size-1-i].size;
    }
    printf("res: %llu\n", res);

    da_destroy(&vectors);
    da_destroy(&jb_dists);
    Circuits_Destroy(&circuits);
    return 0;
}
