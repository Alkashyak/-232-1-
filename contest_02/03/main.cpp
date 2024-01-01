#include <iostream>
#include <algorithm>
#include <vector>
#include <string>

bool compare(std::string a,std::string b){ 
    int num1{0}; 
    int num2{0}; 
    for(char& number : a){ 
        if(number == '1'){ 
            num1 += 1; 
        } 
    } 
    for(char& number : b){ 
        if(number == '1'){ num2 += 1;} 
    } 
    if(num1 != num2){ 
        return(num1 > num2); 
    } else if(num1 == num2) { int num1 = stoi(a); int num2 = stoi(b); return(num1 < num2);} 
}
int main(){
    int count;
    std::cin >> count;
    
    std::vector<std::string> nums(count);
    for(auto& line : nums) std::cin >> line;
    
    std::sort(nums.begin(), nums.end(), compare);
    
    for(auto& line : nums) std::cout << line << ' ';
}
