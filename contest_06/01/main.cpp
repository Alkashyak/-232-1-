#include <cmath>
class Shape {
public:
    virtual double getPerimeter() {
        return 0.0;
    }
};
class Rectangle : public Shape {
public:
    int a;
    int b; 
    Rectangle(int a, int b) {
        this->a = a;
        this->b = b;
    }

    double getPerimeter() override {
        return 2 * (this->a + this->b);
    }
};

class Circle : public Shape {
public:
    int r;
    Circle(int r) {
        this->r = r;
    }
    double getPerimeter() override {
        return 2 * M_PI * r;
    }
};
class Triangle : public Shape {
public:
    int a;
    int b;
    int c;
    Triangle(int a, int b, int c) {
        this->a = a;
        this->b = b;
        this->c = c;
    }

    double getPerimeter() override {
        if (this->a + this->b > this->c && this->a + this->c > this->b && this->b + this->c > this->a) {
            return this->a + this->b + this->c;
        } else {
            return 0.0;
        }
    }
};
