#include <iostream>
#include <string>
#include <set>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    int n, m;
    cin >> n;
    set<string> n1;
    for (int i = 0; i < n; i++) {
        string input;
        cin >> input;
        n1.insert(input);
    }

    cin >> m;
    set<string> n2;
    for (int i = 0; i < m; i++) {
        string input;
        cin >> input;
        n2.insert(input);
    }

    vector<string> result;
    set_intersection(n1.begin(), n1.end(), n2.begin(), n2.end(), back_inserter(result));

    if (result.empty()) {
        cout << -1;
    } else {
        for (const string& word : result) {
            cout << word << " ";
        }
    }

    return 0;
}
