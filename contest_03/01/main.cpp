#include <iostream>
#include <vector>
using namespace std;
int main() {
    int n;
    cin >> n;
    vector<int> n1(n);
    for (int i = 0; i < n; i++) {
        cin >> n1[i];
    }
    int m;
    cin >> m; 
    vector<int> n2(m);
    for (int i = 0; i < m; i++) {
        cin >> n2[i];
    }
    vector<int> result;
    int i = 0, j = 0;
    while (i < n && j < m) {
        if (n1[i] <= n2[j]) {
            result.push_back(n1[i]);
            i++;
        } else {
            result.push_back(n2[j]);
            j++;
        }
    }
    while (i < n) {
        result.push_back(n1[i]);
        i++;
    }
    while (j < m) {
        result.push_back(n2[j]);
        j++;
    }
    for (int k = 0; k < n + m; k++) {
        cout << result[k] << " ";
    }
    return 0;
}
