#include <iostream>

#include <cmath>
bool istina = true;
int is_prime(long n) {
    for(long i = 2; i < (pow(n, 0.5)+1); i++) {
        if(n % i == 0){istina = false;}
    }
return int(istina);
}

int main(){
    int number;
    std::cin >> number;
    std::cout << (is_prime(number) ? "YES" : "NO") << std::endl;
}
