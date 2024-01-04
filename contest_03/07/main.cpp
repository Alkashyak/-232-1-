#include <iostream>
#include <fstream>
#include <vector>
#include <map>
#include <sstream>
#include <algorithm>
#include <iterator>
using namespace std;

map<string, int> nextWords;

bool sortAlgorithm(string a, string b) {
    if (nextWords[a] == nextWords[b]) {
        int i = distance(nextWords.begin(), nextWords.find(a));
        int j = distance(nextWords.begin(), nextWords.find(b));
        return i < j;
    }
    return nextWords[a] > nextWords[b];
}

int main() {
    ifstream file("data.txt");
    string text;
    getline(file, text);
    vector<string> words;
    istringstream newText(text);
    string word = "";

    while (word != "stopword") {
        getline(newText, word, ' ');
        words.push_back(word);
    }

    string target;
    cin >> target;
    int maxCount = 1;

    for (int i = 1; words[i] != "stopword"; i++) {
        if (words[i - 1] == target) {
            if (nextWords.count(words[i]) > 0) {
                nextWords[words[i]] += 1;
                maxCount = max(maxCount, nextWords[words[i]]);
            } else {
                nextWords[words[i]] = 1;
            }
        }
    }

    vector<string> finalStrings;
    int count = 0;

    while (count < 3) {
        for (auto it = nextWords.begin(); it != nextWords.end(); it++) {
            if (it->second == maxCount && count < 3) {
                count += 1;
                finalStrings.push_back(it->first);
            }
        }
        maxCount -= 1;

        if (maxCount == 0) {
            break;
        }
    }

    if (count == 0) {
        cout << "-";
    } else {
        sort(finalStrings.begin(), finalStrings.end(), sortAlgorithm);
    }

    for (string word : finalStrings) {
        cout << word << " ";
    }
}



