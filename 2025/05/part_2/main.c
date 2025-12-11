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

void node_destroy(Node **nodes) {
    Node *node = *nodes;
    while (node) {
        Node *destroy = node;
        node = node->next;
        free(destroy);
    }
    *nodes = NULL;
}

void node_merge_forwards(Node *node) {
    while (node->next) {
        if (node->to < node->next->to) {
            if (node->to >= node->next->from) node->to = node->next->to;
            else return;
        }

        Node *delete = node->next;
        node->next = node->next->next;
        free(delete);
    }
}

#define in_range(I, R) (((I).from >= (R).from && (I).from <= (R).to) || ((I).to >= (R).from && (I).to <= (R).to))

void unique_ids(Node **ranges, Range range) {
    Node *node = *ranges;

    if (!node || range.to < node->from) {
        Node *new_node = node_create(range.from, range.to);
        new_node->next = node;
        *ranges = new_node;
        return;
    }

    while (node) {
        if (!node->next) break;
        if (in_range(range, *node)) break;
        if (in_range(*node, range)) break;
        if (range.to < node->next->from) break;
        node = node->next;
    }

    if (range.from >= node->from && range.to <= node->to) return;

    if (range.from < node->from) node->from = range.from;
    if (range.from <= node->to) {
        if (range.to >= node->to) {
            node->to = range.to;
            node_merge_forwards(node);
        }
        return;
    }

    Node *new_node = node_create(range.from, range.to);
    new_node->next = node->next;
    node->next = new_node;
}

u64 node_count(Node *node) {
    u64 count = 0;
    while (node) {
        count += node->to - node->from + 1;
        node = node->next;
    }
    return count;
}

int main() {
    char file[FILE_SIZE];
    int file_size = read_entire_file("../input.txt", file, FILE_SIZE);
    //int file_size = read_entire_file("../example.txt", file, FILE_SIZE);

    Node *ranges = NULL;

    Iter_Input iter = iter_input_lit(file);
    while (get_ranges(&iter)) unique_ids(&ranges, iter.range);

    printf("res: %llu\n", node_count(ranges));
    node_destroy(&ranges);

    return 0;
}
