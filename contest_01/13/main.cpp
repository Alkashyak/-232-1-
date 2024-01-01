#include <iostream>
using namespace std;
int main() {
    int N;
    int Maxhui = 2, zaebalsya = 1, hui = 1;
    cin >> N;

    while (hui <= N) {
        for (zaebalsya; zaebalsya < Maxhui; zaebalsya++) {
            for (int i = 1; (i <= zaebalsya) && (hui <= N); i += 1) {
                cout << hui << " ";
                hui++;
            }
            cout << "\n";
        }

        for (zaebalsya; zaebalsya > 1; zaebalsya--) {
            for (int i = 1; (i <= zaebalsya) && (hui <= N); i += 1) {
                cout << hui << " ";
                hui++;
            }
            cout << "\n";
        }
        Maxhui += 1;
    }
    return 0;
}
