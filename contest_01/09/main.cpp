#include <iostream> 
#include <cmath> 
using namespace std;
int main() {
    float a, b, c;
    cin >> a >> b;
    c = pow((pow(a, 2) + pow(b, 2)), 0.5);
    cout << c;
    return 0;
}
