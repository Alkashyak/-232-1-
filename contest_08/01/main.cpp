#include <iostream>

class Keeper{
    int trash;
    int secret;
    
public:
    Keeper(){
        std::cin >> secret;
    }
};

int hack_it(Keeper keeper) {
    class Hacked_Kreeper {
    public:
        int trash;
        int secret;
    };
    Hacked_Kreeper* hacked_1 = reinterpret_cast<Hacked_Kreeper*>(&keeper);
    return hacked_1->secret;
}

int main(){
    Keeper keeper;
    int secret= hack_it(keeper);
    std::cout << secret;
}
