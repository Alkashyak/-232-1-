#include <iostream>
#include <string>
#include <vector>
#include <map>
using namespace std;

int main() {
    int n;
    cin >> n;

    vector<string> login;
    vector<string> bal;

    for (int i = 0; i < n; i++) {
        string a;
        string b = "";
        string c = "";

        cin >> a;

        size_t ic = a.find(';');
        for (size_t j = 0; j < a.length(); j++) {
            if (j < ic) {
                b = b + a[j];
            } else if (j > ic) {
                c = c + a[j];
            }
        }

        login.push_back(b);
        bal.push_back(c);
    }

    int n2;
    cin >> n2;

    map<string, string> result;

    for (int i = 0; i < n; i++) {
        result[login[i]] = bal[i];
    }

    for (int i = 0; i < n2; i++) {
        string resk;
        cin >> resk;
        cout << result[resk] << ' ';
    }

    return 0;
}
