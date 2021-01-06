#include <stdio.h>

// see https://golang.org/cmd/cgo/
int fibonacci(int x)
{
  if (x==0)
    return 0;
  return (x + fibonacci(x - 1));
}

void main()
{
  printf("fibonacci(10)=%d\n", fibonacci(10));
}

