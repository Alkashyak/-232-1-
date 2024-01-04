class Water {
public:
    int temperature;
    Water(int temperature) {
        this->temperature = temperature;
    }
    Water() {
        this->temperature = 0;
    }
};
class Teapot {
public:
    int heat;
    Water water;
    Teapot(Water& water) {
        this->water = water;
        this->heat = 0;
    }
    int heat_up(int heat) {
        this->heat += heat;
        this->water.temperature += heat;
        return heat;
    }
    bool is_boiling() {
        return this->water.temperature >= 100;
    }
};

