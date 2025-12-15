#include <stdio.h>

int main(int argc, char* argv[])
{
  char* test[4][4] = {{"dat", "dep"}};
  // test[0][0];
  printf("test[0][1]: %s\n", test[0][1]);

  return 0;
}
