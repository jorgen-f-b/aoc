#include <stdio.h>
#include <stdbool.h>

int read_entire_file(const char *filename, char *buffer, int buffer_size) {
    FILE *f = fopen(filename, "rb");
    fseek(f, 0, SEEK_END);
    int file_size = ftell(f);
    fseek(f, 0, SEEK_SET);

    if (buffer_size <= file_size) {
        fclose(f);
        printf("Buffer size: %d, file size: %d\n", buffer_size, file_size);
        return 0;
    }

    fread(buffer, file_size, 1, f);
    fclose(f);
    buffer[file_size] = 0;

    return file_size;
}

typedef struct {
    char *file;
    char *line;
} File_Line_Iter;
#define file_line_iter_lit(F) (File_Line_Iter){ F }

bool iter_line(File_Line_Iter *iter) {
    if (!*iter->file) return false;

    iter->line = iter->file;
    while (*iter->file != '\n') iter->file++;
    *iter->file = 0;
    iter->file++;

    return true;
}
