#include <iostream>

class Vector;
class Point{
    public:int x,y;
    Point(int x_new, int y_new){
        this->x = x_new;
        this->y = y_new;
    }
    void move(Vector vec);
    
};
class Vector{
    public:int x,y;
    Vector(Point p1, Point p2){
        this->x = p2.x - p1.x;
        this->y = p2.y - p1.y;
    }
};
void Point::move(Vector vec){
        this->x += vec.x;
        this->y += vec.y;
}

int main(){
    int x1, y1, x2, y2, x3, y3;
    std::cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3;
    
    Point p(x1, y1);
    Vector dir(Point(x2, y2), Point(x3, y3));
    p.move(dir);
    
    std::cout << p.x << ' ' << p.y;
}
