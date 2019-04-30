#include <stdio.h>

int fib(int n)
{
    if (n==0) return 0;
    if (n==1) return 1;
    if (n>1) return fib(n-1)+fib(n-2);
}
int main(){
   int l = fib(100);
   printf("fib:%d\n",l);
}



        

