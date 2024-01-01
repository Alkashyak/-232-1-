#include <iostream>
#include <string> 
using namespace std;
int main() {
    string stringi;
    char bukva = ' ';
    cin >> stringi;
    stringi += " ";
    int c = 1;
    for (char bukva2: stringi){
    if (bukva == bukva2) {
        c += 1;
    } 
    else {
        if (c != 1 && bukva != ' ') {
            cout << bukva << c;
        }
        else if (bukva != ' ' ) {
            cout << bukva;
        }
        c = 1;
        bukva = bukva2;
    }    
    }
    return 0;
}
