#include<stdio.h>
#include<time.h>
int sfib(int i, int* j){
    printf("j的值%d",*j);
    int t=i;
    i=*j;
    *j=t+*j;
    printf("i的值%d\n",i);
    return i;
}

int (*fib())(int,int*){
    return sfib;
}
int main(){
    int rues;
    int tmp=1;
    int (*p)(int,int*);
    p=fib();
    int i=0;
    while(i<30){
        rues=p(rues,&tmp);
        i++;  
    }
    printf("斐波那契数列结果%d\n",rues);
    return 0;
}

