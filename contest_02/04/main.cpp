#include <iostream>
#include <tuple>

#include <cmath>
std::tuple<long, long> reduce(long a, long b) {
    long del = 1;
    if (abs(a) >= abs(b)) {
        for (long i = 2; i <= b; i++){
            if (a % i == 0 && b % i == 0) {
                del = i;
            }
        }
    }else { 
        for (long i = 2; i <= a; i++){
            if (a % i == 0 && b % i == 0) {
                del = i;
            }

    }
    }
    return std::make_tuple(a/del, b/del);
}
std::tuple<long, long, long>find_lcm(long c, long d) {
    long del = 1;
    if (abs(c) >= abs(d)) {
        for (long i = 2; i <= d; i++){
            if (c % i == 0 && d % i == 0) {
                del = i;
            }
        }
    }else { 
        for (long i = 2; i <= c; i++){
            if (c % i == 0 && d % i == 0) {
                del = i;
            }
    }

}
    long lcm = (c*d/del);
    return std::make_tuple(lcm, lcm/c, lcm/d);
}


int main() {
    int m1, n1, m2, n2;
    char _;
    std::cin >> m1 >> _ >> n1
             >> m2 >> _ >> n2;
    std::tie(m1, n1) = reduce(m1, n1);
    std::tie(m2, n2) = reduce(m2, n2);

    auto[lcm, c1, c2] = find_lcm(n1, n2);
    auto[m, n] = reduce(m1 * c1 + m2 * c2, lcm);

    std::cout << m << '/' << n << std::endl;
}
