#include <iostream>
#include <string>
#include <stack>
using namespace std;
int main() {
string aboba;
getline(cin, aboba, '!');
stack<char> abobus;
for (char c : aboba) {
    switch (c) {
        case '(': abobus.push(')'); break;
        case '[': abobus.push(']'); break;
        case '{': abobus.push('}'); break;
        case ')':
        case ']':
        case '}':
if (abobus.empty() || abobus.top() != c) {
    cout << "NO";
    return 0;
}
    abobus.pop();
    break;
    default:
    break;
}
}
cout << (abobus.empty() ? "YES" : "NO");
}
