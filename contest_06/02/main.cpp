#include <string>
#include <cmath>
class Employee{
public:
    std::string name = "";
    std::string post = "";
    long bon = 0;
    long salI = 0;
    Employee(std::string name, std::string post, long salI){
        this->name = name;
        this->post = post;
        this->salI = salI;
    }
    virtual long bonus(double per){
        double ab = per * this->salI;
        this->bon = std::round(ab);
        return bon;
    }
    void salary(){
        this->salI = salI + bonus(bonuses[post]);
    }
    friend std::ostream& operator<<(std::ostream& os, Employee& c1) {
        c1.salary();
        os << c1.name << " (" << c1.post << "): " << c1.salI;
        return os;
    }
};
class Manager : public Employee{
public:
    std::string name2 = "";
    long sal2 = 16242;
    Manager(std::string nm, long sl) : Employee(nm, "manager", sl){
        this->name2 = nm;
        this->sal2 = sl;
    }
    long bonus(double per) override{
        if (per >= 0.1){
            return std::round(per * sal2);
        }else{ 
            return std::round(sal2 * 0.1);
        }
    }
};
