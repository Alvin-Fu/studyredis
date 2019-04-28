#include <stdio.h>
#include "functest.h"
int add(int a , int b , int (*add_value)())
{
        return (*add_value)(a,b);
}
int add_ret(int a , int b)
{
        return a + b ;
}

int main(void)
{
    int sum = add(3,4,add_ret);
    printf("sum:%d\n",sum);
    return 0 ;
}
