#include <iostream>

struct Complex{
  double re;  // Действительная часть
  double im;  // Мнимая часть
};

#include <sstream>  
#include <string>
#include <cctype>
#include <algorithm>
Complex make_complex(const std::string& num) {
    Complex complex;
    std::string str = num;
    str.erase(std::remove(str.begin(), str.end(), ' '), str.end());
    std::istringstream iss(str);
    iss >> complex.re >> complex.im;
    return complex;
}
void print(Complex c) {
    std::cout << c.re;
	if (c.im == -0) {
    	c.im = 0;
	}
    if (c.im >= 0.0) {
        std::cout << "+";
    }
    std::cout << c.im << "j" << std::endl;
}

Complex sum(Complex c1, Complex c2) {
    Complex c3;
    c3.re = c1.re + c2.re;
    c3.im = c1.im + c2.im;
    return c3;
}

Complex sub(Complex c1, Complex c2) {
    Complex c3;
    c3.re = c1.re - c2.re;
    c3.im = c1.im - c2.im;
    return c3;
}

Complex mul(Complex c1, Complex c2) {
    Complex c3;
    c3.re = (c1.re * c2.re - c1.im * c2.im);
    c3.im = (c1.im * c2.re + c2.im * c1.re);
    return c3;
}

Complex div(Complex c1, Complex c2) {
    Complex c3;
    c3.re = ((c1.re * c2.re + c1.im * c2.im) / (c2.re * c2.re + c2.im * c2.im));
    c3.im = ((c1.im * c2.re - c1.re * c2.im) / (c2.re * c2.re + c2.im * c2.im));
    return c3;
}


int main()
{
    std::string num;
    std::getline(std::cin, num, 'j');
    Complex c1 = make_complex(num);
    
    std::getline(std::cin, num, 'j');
    Complex c2 = make_complex(num);
    
    print(sum(c1, c2));
    print(sub(c1, c2));
    print(mul(c1, c2));
    print(div(c1, c2));
}
