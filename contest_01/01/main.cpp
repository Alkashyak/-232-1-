#include <iostream> 
#include <cmath> 
using namespace std; 
int main() {
    float a = pow(12.0, 0.5) * (1.0 - 1.0/(3.0*3.0) + 1.0/(5.0*pow(3.0,2.0)) - 1.0/(7.0*pow(3.0,3.0)) + 1.0/(9.0*pow(3.0,4.0)) - 1.0/(11.0*pow(3.0,5.0)));
    cout << a; 
    return 0;
}
