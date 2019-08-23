#include "common.h"


int main(int argc, char** argv) {

    show_cmd(argc, argv);

    return 0;
}

void show_cmd(int argc, char** argv){
    int i=0;

    printf("cmd is:\n");

    while(i<argc) {
        printf("%s ", argv[i]);
        i++;
    }
    printf("\n");

    return;
}