#include <iostream>
#include <string>

class Document{
    std::string data;
public:
    Document(std::string data):data(data){
    }

    const std::string& get() const{
        return data;
    }
};

class PrintStrategy {
public:
    virtual void print(const Document& document) = 0;
};


class Printer {
    PrintStrategy* strategy;
    
public:
    Printer(PrintStrategy* strategy):strategy(strategy){
    }

    void setStrategy(PrintStrategy* strategy) {
        delete this->strategy;
        this->strategy = strategy;
    }
    
    void print(const Document& doc) {
        if (strategy == nullptr) exit(1);
        
        strategy->print(doc);
    }
    
    ~Printer(){
        delete strategy;
    }
};

class ScreenPrintStrategy : public PrintStrategy {
public:
    void print(const Document& doc) override {
        std::cout << doc.get() << '\n';
    }
};

class StringPrintStrategy : public PrintStrategy {
private:
    std::string str;
    int count;
public:
    StringPrintStrategy() {
        this->count = 1;
    }
    void print(const Document& document) override {
        this->str += ("--- doc " + std::to_string(count) + " ---\n" + document.get() + "\n");
        count++;
    }
    std::string getPrintedDocuments() {
        return this->str;
    }
};

class MockPrintStrategy : public PrintStrategy {
private:
    int cnt;
public:
    MockPrintStrategy() {
        this->cnt = 0;
    }
    int getPrintedDocumentsCount() {
        return this->cnt;
    }
    void print(const Document& doc) override {
        this->cnt++;
    }
};

int main(){
    Printer printer(new ScreenPrintStrategy());

    int count;
    std::cin >> count; std::cin.ignore(1);
    while(count--){
        std::string data;
        std::getline(std::cin, data);
        printer.print(Document(data)); // Печать документа на экран
    }
    
    StringPrintStrategy* stringStrategy = new StringPrintStrategy();
    printer.setStrategy(stringStrategy);
    std::cin >> count; std::cin.ignore(1);
    while(count--){
        std::string data;
        std::getline(std::cin, data);
        printer.print(Document(data)); // Печать документа в строку
    }
    std::cout << stringStrategy->getPrintedDocuments() << std::endl;
    
    
    MockPrintStrategy* mockStrategy = new MockPrintStrategy();
    printer.setStrategy(mockStrategy);
    std::cin >> count; std::cin.ignore(1);
    while(count--){
        std::string data;
        std::getline(std::cin, data);
        printer.print(Document(data)); // Имитация печати, на самом деле просто подсчёт
    }
    std::cout << mockStrategy->getPrintedDocumentsCount() << std::endl;
}
