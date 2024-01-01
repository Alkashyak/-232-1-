#include <iostream>
#include <string>
using namespace std;
int main() {
    string num1, num2;
    int count;
    cin >> num1 >> num2;
    for (char chislo1: num1) {
        for (char chislo2: num2){
            if (chislo2 == chislo1) {chislo2 = '*'; chislo1 = 'H'; count++; continue;}
        }
    }
    if (size(num1) != size(num2)) {
        cout << "NO";
    }else {
        if (size(num1) == count) {
            cout << "YES";
        }else { cout << "NO";}
    }
    return 0;
}
