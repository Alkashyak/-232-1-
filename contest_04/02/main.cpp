#include <iostream> 
#include <cmath> 
#include <fstream>
#include "json.hpp"
using json = nlohmann::json;
using namespace std; 
int main()
{
    ifstream fi("data.json");
    json j;
    fi >> j;
    int count = 0; 
    int num;
    cin >> num;
    for (auto& pr: j.items()){
        if (!pr.value()["tasks"].empty()){
            for(auto& task_n: pr.value()["tasks"].items()){
                if(task_n.value()["user_id"] == num && task_n.value()["completed"] == true){
                    count++;
                }
            }
        }
    }
    cout << count; 
    return 0;
}
