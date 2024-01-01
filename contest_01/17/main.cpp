#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    long n;
    cin >> n;
    vector<long> suka(n);
    for (long i = 0; i < n; i++) {
        cin >> suka[i];
    }
    vector<long> lM(n);
    vector<long> rm(n);
    lM[0] = suka[0];
    for (long i = 1; i < n; i++) {
        lM[i] = max(lM[i - 1], suka[i]);
    }
    rm[n - 1] = suka[n - 1];
    for (long i = n - 2; i >= 0; i--) {
        rm[i] = max(rm[i + 1], suka[i]);
    }
    long water = 0;
    for (long i = 0; i < n; i++) {
        water += min(lM[i], rm[i]) - suka[i];
    }
    cout << water << endl;
    return 0;
}
