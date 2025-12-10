#include <stdlib.h>
#include "../../base.h"
#include "../../read_file.c"

#define FILE_SIZE 21373
#define STR_NUM_LEN 20

typedef struct {
    u64 from, to;
} Range;
#define range_lit(F, T) ((Range){ F, T })

typedef struct {
    char *input;
    int idx;
    Range range;
} Iter_Input;
#define iter_input_lit(I) { I, 0 }

u64 str_to_u64(char *str) {
    return strtoull(str, NULL, 10);
}

bool get_ranges(Iter_Input *iter) {
    if (iter->input[iter->idx] == '\n') {
        iter->idx++;
        return false;
    }

    char from[STR_NUM_LEN] = {0};
    char to[STR_NUM_LEN] = {0};
    int idx = 0;
    bool is_from = true;

    while (iter->input[iter->idx] != '\n') {
        if (iter->input[iter->idx] == '-') {
            is_from = false;
            idx = 0;
            iter->idx++;
        }
        if (is_from) from[idx] = iter->input[iter->idx];
        else to[idx] = iter->input[iter->idx];
        idx++;
        iter->idx++;
    }
    iter->idx++;
    iter->range = range_lit(str_to_u64(from), str_to_u64(to));

    return true;
}

typedef struct node Node;
struct node {
    u64 from;
    u64 to;
    Node *next;
};
Node *node_create(u64 from, u64 to) {
    Node *node = malloc(sizeof(Node));
    node->from = from;
    node->to = to;
    node->next = NULL;
    return node;
}

Node *node_delete_lower(Node *node, u64 to) {
    while (node) {
        if (to < node->to) return node;

        Node *delete = node;
        node = node->next;
        free(delete);
    }
    return NULL;
}

void unique_ids(Node **ranges, Range range) {
    Node *node = *ranges;
    if (node && range.from < node->from) {
        if (range.to <= node->to) {
            node->from = range.from;
            return;
        }
        node->next = node_delete_lower(node->next, range.to);
        return;
    }

    Node *prev = NULL;
    while (node) {
        if (range.from >= node->from && range.from <= node->to) {
            if (range.to > node->to) {
                node->to = range.to;
                node->next = node_delete_lower(node->next, range.to);
                return;
            }
            return;
        }
        prev = node;
        node = node->next;
    }

    
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    u64 res = 0;
    u64 range_from[200];
    u64 range_to[200];
    int range_size = 0;

    Iter_Input iter = iter_input_lit(file);
    while (get_ranges(&iter)) {
        u64 from = iter.range.from;
        u64 to = iter.range.to;

        for (int i = 0; i < range_size; i++) {
            if (from >= range_from[i] && from <= range_to[i]) from = range_to[i] + 1;
            if (to >= range_from[i] && to <= range_to[i]) to = range_from[i] - 1;
        }
        if (to < from) continue;

        range_from[range_size] = from;
        range_to[range_size] = to;
        range_size++;

        res += to - from + 1;
    }

    printf("res: %llu\n", res);

    return 0;
}
