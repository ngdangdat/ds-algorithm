#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>


bool isCouponNameValid(char* code) {
  if (code[0] == '\0') return false;
  int idx = 0;
  while (code[idx] != '\0') {
    char c = code[idx];
    idx++;
    if ((c >= 'a' && c <= 'z')) continue;
    if ((c >= 'A' && c <= 'Z')) continue;
    if ((c >= '0' && c <= '9')) continue;
    if (c == '_') continue;
    return false;
  }
  return true;
}

// use the mechanism of insertion sort
void insertInOrder(char** arr, char* el) {
  if (*arr == NULL) {
    arr[0] = el;
    return;
  }
  int insertIdx = 0;
  for (int i = 0; arr[i] != NULL; i++) {
    if (strcmp(el, arr[i]) < 0) {
      insertIdx = i;
      break;
    }
    // make sure it keeps tracking the index
    insertIdx = i;
  }
  char* nextEl = el;
  for (int i = insertIdx; nextEl != NULL; i++) {
    char* tmp = arr[i];
    arr[i] = nextEl;
    nextEl = tmp;
  }
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
char** validateCoupons(char** code, int codeSize, char** businessLine, int businessLineSize, bool* isActive, int isActiveSize, int* returnSize) {
  const char* validLines[] = {"electronics", "grocery", "pharmacy", "restaurant"};
  char** result = calloc(codeSize, sizeof(char*));
  int idx = 0;
  const int validLinesSize = (int) (sizeof(validLines) / sizeof(char*));
  for (int i = 0; i < codeSize; i++) {
    char* cLine = businessLine[i];
    bool cActive = isActive[i];
    char* cCode = code[i];
    bool isLineValid = false;
    for (int j = 0; j < validLinesSize; j++) {
      if (strcmp(validLines[j], cLine) == 0) {
        isLineValid = true;
        break;
      }
    }
    if (isLineValid && cActive && isCouponNameValid(cCode)) {
      insertInOrder(result, cCode);
      idx++;
    }
  }
  *returnSize = idx;
  return result;
}

int main(int argc, char **argv)
{
  
  char* codes[] = {"SAVE20","","PHARMA5","SAVE@20"};
  char* lines[] = {"restaurant","grocery","pharmacy","restaurant"};
  bool actives[] = {true, true, true, true};
  int returnSize = -1;
  char** result = validateCoupons(codes, 4, lines, 4, actives, 4, &returnSize);
  printf("return size: %d\n", returnSize);
  for (int i = 0; i < returnSize; i++) {
    printf("%s\n", result[i]);
  }
  printf("test, dat > da: %d\n", strcmp("da", "dat"));
  return 0;
}
