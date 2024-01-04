#include <iostream>

class IntSharedPointer {
private:
    int* data;
    int* count;

public:
    //Конструктор
    IntSharedPointer(int* p) {
        count = new int{1};
        data = p;
    }

    //Деструктор
    ~IntSharedPointer() {
        if (--(*count) == 0) {
            delete data;
            delete count;
        }
    }

    //Оператор *
    int& operator*() {
        return *data;
    }
    
    //Конструктор копирования
    IntSharedPointer(const IntSharedPointer& other) : data(other.data), count(other.count) {
        (*count)++;
    }

    //Оператор =
    IntSharedPointer& operator=(IntSharedPointer other) {
        std::swap((*this).data, other.data);
        std::swap((*this).count, other.count);
        return *this;
    }

};


void print(IntSharedPointer p){
    std::cout << *p << std::endl;
}

void print_with_exception(IntSharedPointer p){
    std::cout << *p << std::endl;
    if (*p % 2) throw std::string("error");
}

int main(){
    IntSharedPointer p(new int);
    std::cin >> *p;
    p = p;
    print(p);

    IntSharedPointer p2(p);
    std::cin >> *p2;
    print(p);
    print(p2);
    
    {
        int value;
        std::cin >> value;
        IntSharedPointer p3(new int(value));
        p2 = p3;
    }
    print(p2);
    
    p = p2;
    try{
        print_with_exception(p);
    }
    catch (const std::string& e) {
        std::cout << e << std::endl;
    }
}
