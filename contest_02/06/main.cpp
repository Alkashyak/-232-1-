#include <iostream>
#include <vector>
#include <string>
#include <algorithm>

struct Student {
    std::string name;
    std::string group;
};

#include <sstream>
#include <cstring>
#include <stdio.h>
#include <stdlib.h>
using namespace std;
Student make_student(const string& str) {
    Student stud;
    istringstream iss(str);
    getline(iss, stud.name, ';');  
    getline(iss, stud.group);   
    return stud;
}
bool is_upper(Student s1, Student s2) {
    if (s1.group == s2.group)
        return s1.name < s2.name;
    return s1.group < s2.group;
}
void print(vector<Student> students) {
    if (students.size() == 0) {
        cout << "Empty list!";
    } else {
        string g = "";
        for (int i = 0; i < students.size(); i++) {
            if (g != students[i].group) {
                g = students[i].group;
                cout << g << endl;
                cout << "- " << students[i].name << endl;
            }
            else {
                cout << "- " << students[i].name << endl;
            }
        }
    }
}

int main()
{
    int count;
    std::cin >> count;
    std::cin.ignore(1);  // Убираем из потока символ \n для корректной работы getline
    
    std::vector<Student> students(count);
    for(auto& student: students) {
        std::string line;
        std::getline(std::cin, line);
        student = make_student(line);
    }
    
    std::sort(students.begin(), students.end(), is_upper);
    
    print(students);
}
