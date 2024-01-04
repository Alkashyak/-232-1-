#include <cmath>
using namespace std;
    class Point{
      public:
      double x, y, r, a;
      Point(double x, double y){
          this->x = x;
          this->y = y;
          r = sqrt(pow(x, 2.0) + pow(y, 2.0));
          a = atan2(this->y, this->x);
      }
      void set_x(double x) {
          this->x = x;
          this->r = sqrt(pow(this->x, 2.0) + pow(this->y, 2.0));
          this->a = atan2(this->y, this->x);
      }
      void set_y(double y) {
          this->y = y;
          this->r = sqrt(pow(this->x, 2.0) + pow(this->y, 2.0));
          this->a = atan2(this->y, this->x);
      }
	  void set_a(double a){
          this->a = a;
          this->x = r * cos(this->a);
          this->y = r * sin(this->a);
      }
	  void set_r(double r) {
	      this->r = r;
    	  this->x = r * cos(this->a);
    	  this->y = r * sin(this->a);
	  }
      double get_x() {
          return this->x;
      }
      double get_y() {
          return this->y;
      }
      double get_r() {
          return this->r;
      }
      double get_a() {
          return this->a;
      }
    }; 
