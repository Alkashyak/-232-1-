#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    int n;
    cin >> n;
    map<string, int> aboba;
    
    for (int i = 0; i < n; i++) {
        string txx;
        cin >> txx;
        aboba[txx]++;
    }
    
    int maxCount = 0;
    vector<string> abobus;

    for (const auto& bobus : aboba) {
        if (bobus.second > maxCount) {
            maxCount = bobus.second;
            abobus.clear();
            abobus.push_back(bobus.first);
        } else if (bobus.second == maxCount) {
            abobus.push_back(bobus.first);
        }
    }
    
    sort(abobus.begin(), abobus.end());
    
    for (const string& boba : abobus) {
        cout << boba << " ";
    }
    
    return 0;
}
