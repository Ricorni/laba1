// Массив. 
// создание, 
// добавление элемента (по индексу и в конец массива)
// получение элемента по индексу, 
// удаление элемента по индексу, 
// замена элемента по индексу, 
// длина массива, 
// чтение.
#pragma once

#include <cstddef>
#include <stdexcept>
#include <iostream>

template <typename T>
struct ArrayList 
{
    T* data;
    std::size_t size;
    std::size_t capacity;

    ArrayList() {
        data = nullptr;
        size = 0;
        capacity = 0;
    }

    ~ArrayList() {
        delete[] data;
    }

    void pushBack(const T& value) {
        if (size == capacity) {
            capacity = (capacity == 0) ? 1 : capacity * 2;
            T* newData = new T[capacity];
            for (std::size_t i = 0; i < size; i++) {
                newData[i] = data[i];
            }
            delete[] data;
            data = newData;
        }
        data[size] = value;
        size++;
    }

    void insert(const T& value, std::size_t index) {
        if (index > size) {
            throw std::out_of_range("index out of range");
        }
        
        if (index == size) {
            pushBack(value);
            return;
        }

        if (size == capacity) {
            capacity = (capacity == 0) ? 1 : capacity * 2;
            T* newData = new T[capacity];
            for (std::size_t i = 0; i < size; i++) {
                newData[i] = data[i];
            }
            delete[] data;
            data = newData;
        }

        for (std::size_t i = size; i > index; i--) {
            data[i] = data[i - 1];
        }

        data[index] = value;
        size++;
    }

    T& operator[](std::size_t index) {
        if (index >= size) {
            throw std::out_of_range("index out of range");
        }
        return data[index];
    }

    void remove(std::size_t index) {
        if (index >= size) {
            throw std::out_of_range("index out of range");
        }
        for (std::size_t i = index + 1; i < size; i++) {
            data[i - 1] = data[i];
        }
        size--;
    }

    std::size_t getSize() const {
        return size;
    }

    void print() const {
        for (std::size_t i = 0; i < size; i++) {
            std::cout << data[i] << " ";
        }
        std::cout << "\n";
    }

    void save(std::ostream& os) const {
        for (std::size_t i = 0; i < size; i++) {
            os << data[i] << "\n";
        }
    }
};