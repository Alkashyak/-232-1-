class Cat
{
public:
	bool alive;
	bool is_alive() {
		return alive;
	}
	Cat(bool is_alivee) {
		alive = is_alivee;
	}
};
class Box
{
public:
	Cat open() {
		double random = (double)rand() / RAND_MAX;
		bool is_alive = (random > 0.5);
		return Cat(is_alive);
	}
};
