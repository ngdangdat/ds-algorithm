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

int sortstring( const void *str1, const void *str2 )
{
    char *const *pp1 = str1;
    char *const *pp2 = str2;
    return strcmp(*pp1, *pp2);
}

// use the mechanism of insertion sort
void insertInOrder(char** arr, char* el) {
  if (*arr == NULL) {
    arr[0] = el;
    return;
  }
  int insertIdx = 0;
  for (int i = 0; arr[i] != NULL;) {
    if (strcmp(el, arr[i]) < 0) {
      insertIdx = i;
      break;
    }
    // make sure it keeps tracking the index
    i++;
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
  char* resultByCategory[4][codeSize] = {};
  int resultByCategoryCnt[4] = {};
  int idx = 0;
  const int validLinesSize = (int) (sizeof(validLines) / sizeof(char*));
  for (int i = 0; i < codeSize; i++) {
    int lineIndex = -1;
    char* cLine = businessLine[i];
    bool cActive = isActive[i];
    char* cCode = code[i];
    bool isLineValid = false;
    for (int j = 0; j < validLinesSize; j++) {
      if (strcmp(validLines[j], cLine) == 0) {
        isLineValid = true;
        lineIndex = j;
        break;
      }
    }
    if (isLineValid && cActive && isCouponNameValid(cCode)) {
      resultByCategory[lineIndex][resultByCategoryCnt[lineIndex]] = cCode;
      resultByCategoryCnt[lineIndex]++;
    }
  }
  for (int i = 0; i < validLinesSize; i++) {
    if (resultByCategoryCnt[i] > 1) {
      qsort(resultByCategory[i], resultByCategoryCnt[i], sizeof(char*), sortstring);
    }
    for (int j = 0; j < resultByCategoryCnt[i]; j++) {
      result[idx] = resultByCategory[i][j];
      idx++;
    }
  }
  *returnSize = idx;
  return result;
}

int main(int argc, char **argv)
{
  
  // char* codes[] = {"w4xOTEM20C", "Qf8NjqOTYp"};
  char* codes[] = {"Qf8NjqOTYp", "w4xOTEM20C"};
  char* lines[] = {"pharmacy", "pharmacy"};
  bool actives[] = {true, true};
  int returnSize = -1;
  char** result = validateCoupons(codes, 2, lines, 2, actives, 2, &returnSize);
  printf("return size: %d\n", returnSize);
  for (int i = 0; i < returnSize; i++) {
    printf("%s\n", result[i]);
  }
  free(result);
  return 0;
}
