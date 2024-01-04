using namespace std;
#include <cctype>
#include <algorithm>
#include <sstream>  

class Complex {
public:
	double I_ch;
	double Rl;
	Complex(string num) {
		string str = num;
		str.erase(remove(str.begin(), str.end(), ' '), str.end());
		istringstream iss(str);
		iss >> Rl >> I_ch;
	}
	
	Complex(double real, double imag) : Rl(real), I_ch(imag) {}
	
	friend ostream& operator<<(ostream& out, const Complex& num) {
		out << (num.Rl == 0.0 ? 0.0 : num.Rl)
			<< (num.I_ch < 0 ? "" : "+")
			<< (num.I_ch == 0.0 ? 0.0 : num.I_ch) << 'j';
		return out;
	}
	
	friend Complex operator + (Complex c1, Complex c2) {
		return Complex(c1.Rl + c2.Rl, c1.I_ch + c2.I_ch);
	}
	
	friend Complex operator - (Complex c1, Complex c2) {
		return Complex(c1.Rl - c2.Rl, c1.I_ch - c2.I_ch);
	}
	
	friend Complex operator * (Complex c1, Complex c2) {
		return Complex(c1.Rl * c2.Rl - c1.I_ch * c2.I_ch, c1.I_ch * c2.Rl + c2.I_ch * c1.Rl);
	}
	
	friend Complex operator / (Complex c1, Complex c2) {
		return Complex((c1.Rl * c2.Rl + c1.I_ch * c2.I_ch) / (c2.Rl * c2.Rl + c2.I_ch * c2.I_ch), (c1.I_ch * c2.Rl - c1.Rl * c2.I_ch) / (c2.Rl * c2.Rl + c2.I_ch * c2.I_ch));
	}
	
};
